// Package postgres provides the resource for PostgreSQL server and utility packages.
package postgres

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
)

// start starts a postgres database instance.
// If port is 0, then it will choose a random unused port.
func start(port int, dataDir string, serverLog bool) (err error) {
	// See -p -k -h option definitions in the link below.
	// https://www.postgresql.org/docs/current/app-postgres.html
	// We also set max_connections to 500 for tests.
	p := exec.Command("pg_ctl", "start", "-w",
		"-D", dataDir,
		"-o", fmt.Sprintf(`-p %d -k %s -N 500 -h ""`, port, common.GetPostgresSocketDir()))

	uid, _, sameUser, err := shouldSwitchUser()
	if err != nil {
		return err
	}
	if !sameUser {
		p.SysProcAttr = &syscall.SysProcAttr{
			Setpgid:    true,
			Credential: &syscall.Credential{Uid: uint32(uid)},
		}
	}

	// It's useful to log the SQL statement errors from Postgres in developer environment.
	if serverLog {
		p.Stdout = os.Stdout
	}
	p.Stderr = os.Stderr
	if err := p.Run(); err != nil {
		return errors.Wrapf(err, "failed to start postgres %q", p.String())
	}

	return nil
}

// stop stops a postgres instance, outputs to stdout and stderr.
func stop(pgDataDir string) error {
	p := exec.Command("pg_ctl", "stop", "-w", "-D", pgDataDir)
	uid, _, sameUser, err := shouldSwitchUser()
	if err != nil {
		return err
	}
	if !sameUser {
		p.SysProcAttr = &syscall.SysProcAttr{
			Setpgid:    true,
			Credential: &syscall.Credential{Uid: uint32(uid)},
		}
	}

	// Suppress log spam
	p.Stdout = nil
	p.Stderr = os.Stderr
	return p.Run()
}

// initDB inits a postgres database if not yet.
func initDB(pgDataDir, pgUser string) error {
	versionPath := filepath.Join(pgDataDir, "PG_VERSION")
	_, err := os.Stat(versionPath)
	if err != nil && !os.IsNotExist(err) {
		return errors.Wrapf(err, "failed to check postgres version in data directory path %q", versionPath)
	}
	initDone := err == nil

	// Skip initDB if setup already.
	if initDone {
		// If file permission was mutated before, postgres cannot start up. We should change file permissions to 0700 for all pgdata files.
		if err := os.Chmod(pgDataDir, 0700); err != nil {
			return errors.Wrapf(err, "failed to chmod postgres data directory %q to 0700", pgDataDir)
		}
		return nil
	}

	// For pgDataDir and every intermediate to be created by MkdirAll, we need to prepare to chown
	// it to the bytebase user. Otherwise, initdb will complain file permission error.
	dirListToChown := []string{pgDataDir}
	path := filepath.Dir(pgDataDir)
	for path != "/" {
		_, err := os.Stat(path)
		if err == nil {
			break
		}
		if !os.IsNotExist(err) {
			return errors.Wrapf(err, "failed to check data directory path existence %q", path)
		}
		dirListToChown = append(dirListToChown, path)
		path = filepath.Dir(path)
	}
	slog.Debug("Data directory list to Chown", slog.Any("dirListToChown", dirListToChown))

	if err := os.MkdirAll(pgDataDir, 0700); err != nil {
		return errors.Wrapf(err, "failed to make postgres data directory %q", pgDataDir)
	}

	uid, gid, sameUser, err := shouldSwitchUser()
	if err != nil {
		return err
	}
	if !sameUser {
		slog.Info(fmt.Sprintf("Recursively change owner of data directory %q to bytebase...", pgDataDir))
		for _, dir := range dirListToChown {
			slog.Info(fmt.Sprintf("Change owner of %q to bytebase", dir))
			if err := os.Chown(dir, int(uid), int(gid)); err != nil {
				return errors.Wrapf(err, "failed to change owner of %q to bytebase", dir)
			}
		}
	}

	args := []string{
		"-U", pgUser,
		"-D", pgDataDir,
	}
	p := exec.Command("initdb", args...)
	p.Env = append(os.Environ(),
		"LC_ALL=en_US.UTF-8",
		"LC_CTYPE=en_US.UTF-8",
	)
	if !sameUser {
		p.SysProcAttr = &syscall.SysProcAttr{
			Setpgid:    true,
			Credential: &syscall.Credential{Uid: uint32(uid)},
		}
	}
	// Suppress log spam
	p.Stdout = nil
	p.Stderr = os.Stderr
	slog.Info("-----Postgres initdb BEGIN-----")
	if err := p.Run(); err != nil {
		return errors.Wrapf(err, "failed to initdb %q", p.String())
	}
	slog.Info("-----Postgres initdb END-----")

	return nil
}

func shouldSwitchUser() (int, int, bool, error) {
	sameUser := true
	bytebaseUser, err := user.Current()
	if err != nil {
		return 0, 0, true, errors.Wrap(err, "failed to get current user")
	}
	// If user runs Bytebase as root user, we will attempt to run as user `bytebase`.
	// https://www.postgresql.org/docs/14/app-initdb.html
	if bytebaseUser.Username == "root" {
		bytebaseUser, err = user.Lookup("bytebase")
		if err != nil {
			return 0, 0, false, errors.Errorf("please run Bytebase as non-root user. You can use the following command to create a dedicated \"bytebase\" user to run the application: addgroup --gid 113 --system bytebase && adduser --uid 113 --system bytebase && adduser bytebase bytebase")
		}
		sameUser = false
	}

	uid, err := strconv.ParseUint(bytebaseUser.Uid, 10, 32)
	if err != nil {
		return 0, 0, false, err
	}
	gid, err := strconv.ParseUint(bytebaseUser.Gid, 10, 32)
	if err != nil {
		return 0, 0, false, err
	}
	return int(uid), int(gid), sameUser, nil
}
