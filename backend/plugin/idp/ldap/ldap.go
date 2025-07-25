// Package ldap is the plugin for LDAP Identity Provider.
package ldap

import (
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/pkg/errors"

	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
)

// IdentityProvider represents an LDAP Identity Provider.
type IdentityProvider struct {
	config IdentityProviderConfig
}

// IdentityProviderConfig is the configuration to be consumed by the LDAP
// Identity Provider.
type IdentityProviderConfig struct {
	// Host is the hostname or IP address of the LDAP server, e.g.
	// "ldap.example.com".
	Host string `json:"host"`
	// Port is the port number of the LDAP server, e.g. 389. When not set, the
	// default port of the corresponding security protocol will be used, i.e. 389
	// for StartTLS and 636 for LDAPS.
	Port int `json:"port"`
	// SkipTLSVerify controls whether to skip TLS certificate verification.
	SkipTLSVerify bool `json:"skipTlsVerify"`
	// BindDN is the DN of the user to bind as a service account to perform
	// search requests.
	BindDN string `json:"bindDn"`
	// BindPassword is the password of the user to bind as a service account.
	BindPassword string `json:"bindPassword"`
	// BaseDN is the base DN to search for users, e.g. "ou=users,dc=example,dc=com".
	BaseDN string `json:"baseDn"`
	// UserFilter is the filter to search for users, e.g. "(uid=%s)".
	UserFilter string `json:"userFilter"`
	// SecurityProtocol is the security protocol to be used for establishing
	// connections with the LDAP server.
	SecurityProtocol storepb.LDAPIdentityProviderConfig_SecurityProtocol `json:"securityProtocol"`
	// FieldMapping is the mapping of the user attributes returned by the LDAP
	// server.
	FieldMapping *storepb.FieldMapping `json:"fieldMapping"`
}

// NewIdentityProvider initializes a new LDAP Identity Provider with the given
// configuration.
func NewIdentityProvider(config IdentityProviderConfig) (*IdentityProvider, error) {
	for v, field := range map[string]string{
		config.Host:                    "host",
		config.BindDN:                  "bindDn",
		config.BindPassword:            "bindPassword",
		config.BaseDN:                  "baseDn",
		config.UserFilter:              "userFilter",
		config.FieldMapping.Identifier: "fieldMapping.identifier",
	} {
		if v == "" {
			return nil, errors.Errorf("the field %q is empty but required", field)
		}
	}

	if config.Port <= 0 {
		if config.SecurityProtocol == storepb.LDAPIdentityProviderConfig_LDAPS {
			config.Port = 636
		} else {
			config.Port = 389
		}
	}

	return &IdentityProvider{
		config: config,
	}, nil
}

func (p *IdentityProvider) dial() (*ldap.Conn, error) {
	tlsConfig := &tls.Config{
		ServerName:         p.config.Host,
		InsecureSkipVerify: p.config.SkipTLSVerify,
	}
	switch p.config.SecurityProtocol {
	case storepb.LDAPIdentityProviderConfig_LDAPS:
		url := fmt.Sprintf("ldaps://%s:%d", p.config.Host, p.config.Port)
		conn, err := ldap.DialURL(url, ldap.DialWithTLSConfig(tlsConfig))
		if err != nil {
			return nil, errors.Errorf("dial TLS: %v", err)
		}
		return conn, nil
	case storepb.LDAPIdentityProviderConfig_START_TLS:
		url := fmt.Sprintf("ldap://%s:%d", p.config.Host, p.config.Port)
		conn, err := ldap.DialURL(url)
		if err != nil {
			return nil, errors.Errorf("dial: %v", err)
		}
		if err := conn.StartTLS(tlsConfig); err != nil {
			_ = conn.Close()
			return nil, errors.Errorf("start TLS: %v", err)
		}
		return conn, nil
	default:
		url := fmt.Sprintf("ldap://%s:%d", p.config.Host, p.config.Port)
		conn, err := ldap.DialURL(url)
		if err != nil {
			return nil, errors.Errorf("dial: %v", err)
		}
		return conn, nil
	}
}

// Connect establishes a connection using the bind DN and bind password.
func (p *IdentityProvider) Connect() (*ldap.Conn, error) {
	conn, err := p.dial()
	if err != nil {
		return nil, err
	}

	// Bind with a system account
	err = conn.Bind(p.config.BindDN, p.config.BindPassword)
	if err != nil {
		_ = conn.Close()
		return nil, errors.Errorf("bind: %v", err)
	}
	return conn, nil
}

// Authenticate authenticates the user with the given username and password.
func (p *IdentityProvider) Authenticate(username, password string) (*storepb.IdentityProviderUserInfo, error) {
	conn, err := p.Connect()
	if err != nil {
		return nil, errors.Errorf("connect: %v", err)
	}
	defer func() { _ = conn.Close() }()

	sr, err := conn.Search(
		ldap.NewSearchRequest(
			p.config.BaseDN,
			ldap.ScopeWholeSubtree,
			ldap.NeverDerefAliases,
			0,
			0,
			false,
			strings.ReplaceAll(p.config.UserFilter, "%s", username),
			[]string{"dn", p.config.FieldMapping.Identifier, p.config.FieldMapping.DisplayName},
			nil,
		),
	)
	if err != nil {
		return nil, errors.Errorf("search user DN: %v", err)
	} else if len(sr.Entries) != 1 {
		return nil, errors.Errorf("expect 1 user DN but got %d", len(sr.Entries))
	}
	entry := sr.Entries[0]

	// Bind as the user to verify their password
	err = conn.Bind(entry.DN, password)
	if err != nil {
		return nil, errors.Errorf("bind user: %v", err)
	}

	identifier := entry.GetAttributeValue(p.config.FieldMapping.Identifier)
	return &storepb.IdentityProviderUserInfo{
		Identifier:  identifier,
		DisplayName: entry.GetAttributeValue(p.config.FieldMapping.DisplayName),
	}, nil
}
