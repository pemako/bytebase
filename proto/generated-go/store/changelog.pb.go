// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: store/changelog.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChangelogTask_Status int32

const (
	ChangelogTask_STATUS_UNSPECIFIED ChangelogTask_Status = 0
	ChangelogTask_PENDING            ChangelogTask_Status = 1
	ChangelogTask_DONE               ChangelogTask_Status = 2
	ChangelogTask_FAILED             ChangelogTask_Status = 3
)

// Enum value maps for ChangelogTask_Status.
var (
	ChangelogTask_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "DONE",
		3: "FAILED",
	}
	ChangelogTask_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PENDING":            1,
		"DONE":               2,
		"FAILED":             3,
	}
)

func (x ChangelogTask_Status) Enum() *ChangelogTask_Status {
	p := new(ChangelogTask_Status)
	*p = x
	return p
}

func (x ChangelogTask_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangelogTask_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_store_changelog_proto_enumTypes[0].Descriptor()
}

func (ChangelogTask_Status) Type() protoreflect.EnumType {
	return &file_store_changelog_proto_enumTypes[0]
}

func (x ChangelogTask_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangelogTask_Status.Descriptor instead.
func (ChangelogTask_Status) EnumDescriptor() ([]byte, []int) {
	return file_store_changelog_proto_rawDescGZIP(), []int{1, 0}
}

type ChangelogTask_Type int32

const (
	ChangelogTask_TYPE_UNSPECIFIED ChangelogTask_Type = 0
	ChangelogTask_BASELINE         ChangelogTask_Type = 1
	ChangelogTask_MIGRATE          ChangelogTask_Type = 2
	ChangelogTask_MIGRATE_SDL      ChangelogTask_Type = 3
	ChangelogTask_MIGRATE_GHOST    ChangelogTask_Type = 4
	ChangelogTask_DATA             ChangelogTask_Type = 6
)

// Enum value maps for ChangelogTask_Type.
var (
	ChangelogTask_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "BASELINE",
		2: "MIGRATE",
		3: "MIGRATE_SDL",
		4: "MIGRATE_GHOST",
		6: "DATA",
	}
	ChangelogTask_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"BASELINE":         1,
		"MIGRATE":          2,
		"MIGRATE_SDL":      3,
		"MIGRATE_GHOST":    4,
		"DATA":             6,
	}
)

func (x ChangelogTask_Type) Enum() *ChangelogTask_Type {
	p := new(ChangelogTask_Type)
	*p = x
	return p
}

func (x ChangelogTask_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangelogTask_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_changelog_proto_enumTypes[1].Descriptor()
}

func (ChangelogTask_Type) Type() protoreflect.EnumType {
	return &file_store_changelog_proto_enumTypes[1]
}

func (x ChangelogTask_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangelogTask_Type.Descriptor instead.
func (ChangelogTask_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_changelog_proto_rawDescGZIP(), []int{1, 1}
}

type ChangelogRevision_Op int32

const (
	ChangelogRevision_OP_UNSPECIFIED ChangelogRevision_Op = 0
	ChangelogRevision_CREATE         ChangelogRevision_Op = 1
	ChangelogRevision_DELETE         ChangelogRevision_Op = 2
)

// Enum value maps for ChangelogRevision_Op.
var (
	ChangelogRevision_Op_name = map[int32]string{
		0: "OP_UNSPECIFIED",
		1: "CREATE",
		2: "DELETE",
	}
	ChangelogRevision_Op_value = map[string]int32{
		"OP_UNSPECIFIED": 0,
		"CREATE":         1,
		"DELETE":         2,
	}
)

func (x ChangelogRevision_Op) Enum() *ChangelogRevision_Op {
	p := new(ChangelogRevision_Op)
	*p = x
	return p
}

func (x ChangelogRevision_Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangelogRevision_Op) Descriptor() protoreflect.EnumDescriptor {
	return file_store_changelog_proto_enumTypes[2].Descriptor()
}

func (ChangelogRevision_Op) Type() protoreflect.EnumType {
	return &file_store_changelog_proto_enumTypes[2]
}

func (x ChangelogRevision_Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangelogRevision_Op.Descriptor instead.
func (ChangelogRevision_Op) EnumDescriptor() ([]byte, []int) {
	return file_store_changelog_proto_rawDescGZIP(), []int{2, 0}
}

type ChangelogPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task     *ChangelogTask     `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Revision *ChangelogRevision `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *ChangelogPayload) Reset() {
	*x = ChangelogPayload{}
	mi := &file_store_changelog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangelogPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogPayload) ProtoMessage() {}

func (x *ChangelogPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_changelog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogPayload.ProtoReflect.Descriptor instead.
func (*ChangelogPayload) Descriptor() ([]byte, []int) {
	return file_store_changelog_proto_rawDescGZIP(), []int{0}
}

func (x *ChangelogPayload) GetTask() *ChangelogTask {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ChangelogPayload) GetRevision() *ChangelogRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

type ChangelogTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}/rollouts/{rollout}/stages/{stage}/tasks/{task}/taskruns/{taskrun}
	TaskRun string `protobuf:"bytes,1,opt,name=task_run,json=taskRun,proto3" json:"task_run,omitempty"`
	// Format: projects/{project}/issues/{issue}
	Issue string `protobuf:"bytes,2,opt,name=issue,proto3" json:"issue,omitempty"`
	// The revision uid.
	// optional
	Revision          int64                `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	ChangedResources  *ChangedResources    `protobuf:"bytes,4,opt,name=changed_resources,json=changedResources,proto3" json:"changed_resources,omitempty"`
	Status            ChangelogTask_Status `protobuf:"varint,5,opt,name=status,proto3,enum=bytebase.store.ChangelogTask_Status" json:"status,omitempty"`
	PrevSyncHistoryId int64                `protobuf:"varint,6,opt,name=prev_sync_history_id,json=prevSyncHistoryId,proto3" json:"prev_sync_history_id,omitempty"`
	SyncHistoryId     int64                `protobuf:"varint,7,opt,name=sync_history_id,json=syncHistoryId,proto3" json:"sync_history_id,omitempty"`
	// The sheet that holds the content.
	// Format: projects/{project}/sheets/{sheet}
	Sheet   string             `protobuf:"bytes,8,opt,name=sheet,proto3" json:"sheet,omitempty"`
	Version string             `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	Type    ChangelogTask_Type `protobuf:"varint,10,opt,name=type,proto3,enum=bytebase.store.ChangelogTask_Type" json:"type,omitempty"`
}

func (x *ChangelogTask) Reset() {
	*x = ChangelogTask{}
	mi := &file_store_changelog_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangelogTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogTask) ProtoMessage() {}

func (x *ChangelogTask) ProtoReflect() protoreflect.Message {
	mi := &file_store_changelog_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogTask.ProtoReflect.Descriptor instead.
func (*ChangelogTask) Descriptor() ([]byte, []int) {
	return file_store_changelog_proto_rawDescGZIP(), []int{1}
}

func (x *ChangelogTask) GetTaskRun() string {
	if x != nil {
		return x.TaskRun
	}
	return ""
}

func (x *ChangelogTask) GetIssue() string {
	if x != nil {
		return x.Issue
	}
	return ""
}

func (x *ChangelogTask) GetRevision() int64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *ChangelogTask) GetChangedResources() *ChangedResources {
	if x != nil {
		return x.ChangedResources
	}
	return nil
}

func (x *ChangelogTask) GetStatus() ChangelogTask_Status {
	if x != nil {
		return x.Status
	}
	return ChangelogTask_STATUS_UNSPECIFIED
}

func (x *ChangelogTask) GetPrevSyncHistoryId() int64 {
	if x != nil {
		return x.PrevSyncHistoryId
	}
	return 0
}

func (x *ChangelogTask) GetSyncHistoryId() int64 {
	if x != nil {
		return x.SyncHistoryId
	}
	return 0
}

func (x *ChangelogTask) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *ChangelogTask) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ChangelogTask) GetType() ChangelogTask_Type {
	if x != nil {
		return x.Type
	}
	return ChangelogTask_TYPE_UNSPECIFIED
}

type ChangelogRevision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Marshalled revision for display
	Revision  string               `protobuf:"bytes,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Operation ChangelogRevision_Op `protobuf:"varint,2,opt,name=operation,proto3,enum=bytebase.store.ChangelogRevision_Op" json:"operation,omitempty"`
}

func (x *ChangelogRevision) Reset() {
	*x = ChangelogRevision{}
	mi := &file_store_changelog_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangelogRevision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangelogRevision) ProtoMessage() {}

func (x *ChangelogRevision) ProtoReflect() protoreflect.Message {
	mi := &file_store_changelog_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangelogRevision.ProtoReflect.Descriptor instead.
func (*ChangelogRevision) Descriptor() ([]byte, []int) {
	return file_store_changelog_proto_rawDescGZIP(), []int{2}
}

func (x *ChangelogRevision) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *ChangelogRevision) GetOperation() ChangelogRevision_Op {
	if x != nil {
		return x.Operation
	}
	return ChangelogRevision_OP_UNSPECIFIED
}

var File_store_changelog_proto protoreflect.FileDescriptor

var file_store_changelog_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x23, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xd6, 0x04, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x73, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52,
	0x10, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x54, 0x61, 0x73, 0x6b,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2f, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x70,
	0x72, 0x65, 0x76, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x43, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0x65, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41, 0x53, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x44, 0x4c, 0x10, 0x03, 0x12,
	0x11, 0x0a, 0x0d, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x48, 0x4f, 0x53, 0x54,
	0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x06, 0x22, 0xa5, 0x01, 0x0a,
	0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x02, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_store_changelog_proto_rawDescOnce sync.Once
	file_store_changelog_proto_rawDescData = file_store_changelog_proto_rawDesc
)

func file_store_changelog_proto_rawDescGZIP() []byte {
	file_store_changelog_proto_rawDescOnce.Do(func() {
		file_store_changelog_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_changelog_proto_rawDescData)
	})
	return file_store_changelog_proto_rawDescData
}

var file_store_changelog_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_store_changelog_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_changelog_proto_goTypes = []any{
	(ChangelogTask_Status)(0), // 0: bytebase.store.ChangelogTask.Status
	(ChangelogTask_Type)(0),   // 1: bytebase.store.ChangelogTask.Type
	(ChangelogRevision_Op)(0), // 2: bytebase.store.ChangelogRevision.Op
	(*ChangelogPayload)(nil),  // 3: bytebase.store.ChangelogPayload
	(*ChangelogTask)(nil),     // 4: bytebase.store.ChangelogTask
	(*ChangelogRevision)(nil), // 5: bytebase.store.ChangelogRevision
	(*ChangedResources)(nil),  // 6: bytebase.store.ChangedResources
}
var file_store_changelog_proto_depIdxs = []int32{
	4, // 0: bytebase.store.ChangelogPayload.task:type_name -> bytebase.store.ChangelogTask
	5, // 1: bytebase.store.ChangelogPayload.revision:type_name -> bytebase.store.ChangelogRevision
	6, // 2: bytebase.store.ChangelogTask.changed_resources:type_name -> bytebase.store.ChangedResources
	0, // 3: bytebase.store.ChangelogTask.status:type_name -> bytebase.store.ChangelogTask.Status
	1, // 4: bytebase.store.ChangelogTask.type:type_name -> bytebase.store.ChangelogTask.Type
	2, // 5: bytebase.store.ChangelogRevision.operation:type_name -> bytebase.store.ChangelogRevision.Op
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_store_changelog_proto_init() }
func file_store_changelog_proto_init() {
	if File_store_changelog_proto != nil {
		return
	}
	file_store_instance_change_history_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_changelog_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_changelog_proto_goTypes,
		DependencyIndexes: file_store_changelog_proto_depIdxs,
		EnumInfos:         file_store_changelog_proto_enumTypes,
		MessageInfos:      file_store_changelog_proto_msgTypes,
	}.Build()
	File_store_changelog_proto = out.File
	file_store_changelog_proto_rawDesc = nil
	file_store_changelog_proto_goTypes = nil
	file_store_changelog_proto_depIdxs = nil
}