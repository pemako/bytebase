// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/instance_role_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetInstanceRoleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the role to retrieve.
	// Format: instances/{instance}/roles/{role name}
	// The role name is the unique name for the role.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInstanceRoleRequest) Reset() {
	*x = GetInstanceRoleRequest{}
	mi := &file_v1_instance_role_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInstanceRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInstanceRoleRequest) ProtoMessage() {}

func (x *GetInstanceRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_instance_role_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInstanceRoleRequest.ProtoReflect.Descriptor instead.
func (*GetInstanceRoleRequest) Descriptor() ([]byte, []int) {
	return file_v1_instance_role_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetInstanceRoleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListInstanceRolesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The parent, which owns this collection of roles.
	// Format: instances/{instance}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Not used.
	// The maximum number of roles to return. The service may return fewer than
	// this value.
	// If unspecified, at most 10 roles will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Not used.
	// A page token, received from a previous `ListInstanceRoles` call.
	// Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to `ListInstanceRoles` must match
	// the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Refresh will refresh and return the latest data.
	Refresh       bool `protobuf:"varint,4,opt,name=refresh,proto3" json:"refresh,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInstanceRolesRequest) Reset() {
	*x = ListInstanceRolesRequest{}
	mi := &file_v1_instance_role_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstanceRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceRolesRequest) ProtoMessage() {}

func (x *ListInstanceRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_instance_role_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceRolesRequest.ProtoReflect.Descriptor instead.
func (*ListInstanceRolesRequest) Descriptor() ([]byte, []int) {
	return file_v1_instance_role_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListInstanceRolesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListInstanceRolesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListInstanceRolesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListInstanceRolesRequest) GetRefresh() bool {
	if x != nil {
		return x.Refresh
	}
	return false
}

type ListInstanceRolesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The roles from the specified request.
	Roles []*InstanceRole `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListInstanceRolesResponse) Reset() {
	*x = ListInstanceRolesResponse{}
	mi := &file_v1_instance_role_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListInstanceRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInstanceRolesResponse) ProtoMessage() {}

func (x *ListInstanceRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_instance_role_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInstanceRolesResponse.ProtoReflect.Descriptor instead.
func (*ListInstanceRolesResponse) Descriptor() ([]byte, []int) {
	return file_v1_instance_role_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListInstanceRolesResponse) GetRoles() []*InstanceRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *ListInstanceRolesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// InstanceRole is the API message for instance role.
type InstanceRole struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the role.
	// Format: instances/{instance}/roles/{role}
	// The role name is the unique name for the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The role name. It's unique within the instance.
	RoleName string `protobuf:"bytes,2,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	// The role password.
	Password *string `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
	// The connection count limit for this role.
	ConnectionLimit *int32 `protobuf:"varint,4,opt,name=connection_limit,json=connectionLimit,proto3,oneof" json:"connection_limit,omitempty"`
	// The expiration for the role's password.
	ValidUntil *string `protobuf:"bytes,5,opt,name=valid_until,json=validUntil,proto3,oneof" json:"valid_until,omitempty"`
	// The role attribute.
	// For PostgreSQL, it containt super_user, no_inherit, create_role, create_db, can_login, replication and bypass_rls. Docs: https://www.postgresql.org/docs/current/role-attributes.html
	// For MySQL, it's the global privileges as GRANT statements, which means it only contains "GRANT ... ON *.* TO ...". Docs: https://dev.mysql.com/doc/refman/8.0/en/grant.html
	Attribute     *string `protobuf:"bytes,6,opt,name=attribute,proto3,oneof" json:"attribute,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InstanceRole) Reset() {
	*x = InstanceRole{}
	mi := &file_v1_instance_role_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstanceRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceRole) ProtoMessage() {}

func (x *InstanceRole) ProtoReflect() protoreflect.Message {
	mi := &file_v1_instance_role_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceRole.ProtoReflect.Descriptor instead.
func (*InstanceRole) Descriptor() ([]byte, []int) {
	return file_v1_instance_role_service_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceRole) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstanceRole) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *InstanceRole) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *InstanceRole) GetConnectionLimit() int32 {
	if x != nil && x.ConnectionLimit != nil {
		return *x.ConnectionLimit
	}
	return 0
}

func (x *InstanceRole) GetValidUntil() string {
	if x != nil && x.ValidUntil != nil {
		return *x.ValidUntil
	}
	return ""
}

func (x *InstanceRole) GetAttribute() string {
	if x != nil && x.Attribute != nil {
		return *x.Attribute
	}
	return ""
}

var File_v1_instance_role_service_proto protoreflect.FileDescriptor

const file_v1_instance_role_service_proto_rawDesc = "" +
	"\n" +
	"\x1ev1/instance_role_service.proto\x12\vbytebase.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x13v1/annotation.proto\"O\n" +
	"\x16GetInstanceRoleRequest\x125\n" +
	"\x04name\x18\x01 \x01(\tB!\xe0A\x02\xfaA\x1b\n" +
	"\x19bytebase.com/InstanceRoleR\x04name\"\xa7\x01\n" +
	"\x18ListInstanceRolesRequest\x125\n" +
	"\x06parent\x18\x01 \x01(\tB\x1d\xe0A\x02\xfaA\x17\n" +
	"\x15bytebase.com/InstanceR\x06parent\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\x12\x18\n" +
	"\arefresh\x18\x04 \x01(\bR\arefresh\"t\n" +
	"\x19ListInstanceRolesResponse\x12/\n" +
	"\x05roles\x18\x01 \x03(\v2\x19.bytebase.v1.InstanceRoleR\x05roles\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\xe1\x02\n" +
	"\fInstanceRole\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1b\n" +
	"\trole_name\x18\x02 \x01(\tR\broleName\x12$\n" +
	"\bpassword\x18\x03 \x01(\tB\x03\xe0A\x04H\x00R\bpassword\x88\x01\x01\x12.\n" +
	"\x10connection_limit\x18\x04 \x01(\x05H\x01R\x0fconnectionLimit\x88\x01\x01\x12$\n" +
	"\vvalid_until\x18\x05 \x01(\tH\x02R\n" +
	"validUntil\x88\x01\x01\x12!\n" +
	"\tattribute\x18\x06 \x01(\tH\x03R\tattribute\x88\x01\x01:A\xeaA>\n" +
	"\x19bytebase.com/InstanceRole\x12!instances/{instance}/roles/{role}B\v\n" +
	"\t_passwordB\x13\n" +
	"\x11_connection_limitB\x0e\n" +
	"\f_valid_untilB\f\n" +
	"\n" +
	"_attribute2\xe7\x02\n" +
	"\x13InstanceRoleService\x12\x9c\x01\n" +
	"\x0fGetInstanceRole\x12#.bytebase.v1.GetInstanceRoleRequest\x1a\x19.bytebase.v1.InstanceRole\"I\xdaA\x04name\x8a\xea0\x14bb.instanceRoles.get\x90\xea0\x01\x82\xd3\xe4\x93\x02 \x12\x1e/v1/{name=instances/*/roles/*}\x12\xb0\x01\n" +
	"\x11ListInstanceRoles\x12%.bytebase.v1.ListInstanceRolesRequest\x1a&.bytebase.v1.ListInstanceRolesResponse\"L\xdaA\x06parent\x8a\xea0\x15bb.instanceRoles.list\x90\xea0\x01\x82\xd3\xe4\x93\x02 \x12\x1e/v1/{parent=instances/*}/rolesB6Z4github.com/bytebase/bytebase/backend/generated-go/v1b\x06proto3"

var (
	file_v1_instance_role_service_proto_rawDescOnce sync.Once
	file_v1_instance_role_service_proto_rawDescData []byte
)

func file_v1_instance_role_service_proto_rawDescGZIP() []byte {
	file_v1_instance_role_service_proto_rawDescOnce.Do(func() {
		file_v1_instance_role_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_instance_role_service_proto_rawDesc), len(file_v1_instance_role_service_proto_rawDesc)))
	})
	return file_v1_instance_role_service_proto_rawDescData
}

var file_v1_instance_role_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_instance_role_service_proto_goTypes = []any{
	(*GetInstanceRoleRequest)(nil),    // 0: bytebase.v1.GetInstanceRoleRequest
	(*ListInstanceRolesRequest)(nil),  // 1: bytebase.v1.ListInstanceRolesRequest
	(*ListInstanceRolesResponse)(nil), // 2: bytebase.v1.ListInstanceRolesResponse
	(*InstanceRole)(nil),              // 3: bytebase.v1.InstanceRole
}
var file_v1_instance_role_service_proto_depIdxs = []int32{
	3, // 0: bytebase.v1.ListInstanceRolesResponse.roles:type_name -> bytebase.v1.InstanceRole
	0, // 1: bytebase.v1.InstanceRoleService.GetInstanceRole:input_type -> bytebase.v1.GetInstanceRoleRequest
	1, // 2: bytebase.v1.InstanceRoleService.ListInstanceRoles:input_type -> bytebase.v1.ListInstanceRolesRequest
	3, // 3: bytebase.v1.InstanceRoleService.GetInstanceRole:output_type -> bytebase.v1.InstanceRole
	2, // 4: bytebase.v1.InstanceRoleService.ListInstanceRoles:output_type -> bytebase.v1.ListInstanceRolesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_instance_role_service_proto_init() }
func file_v1_instance_role_service_proto_init() {
	if File_v1_instance_role_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_instance_role_service_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_instance_role_service_proto_rawDesc), len(file_v1_instance_role_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_instance_role_service_proto_goTypes,
		DependencyIndexes: file_v1_instance_role_service_proto_depIdxs,
		MessageInfos:      file_v1_instance_role_service_proto_msgTypes,
	}.Build()
	File_v1_instance_role_service_proto = out.File
	file_v1_instance_role_service_proto_goTypes = nil
	file_v1_instance_role_service_proto_depIdxs = nil
}
