// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/deployer/v1/stack.proto

package deployer_pb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/pentops/protostate/gen/v1/psm_pb"
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

type StackStatus int32

const (
	StackStatus_STACK_STATUS_UNSPECIFIED StackStatus = 0
	StackStatus_STACK_STATUS_CREATING    StackStatus = 1
	StackStatus_STACK_STATUS_STABLE      StackStatus = 2
	StackStatus_STACK_STATUS_AVAILABLE   StackStatus = 3
	StackStatus_STACK_STATUS_MIGRATING   StackStatus = 4
	StackStatus_STACK_STATUS_BROKEN      StackStatus = 5
)

// Enum value maps for StackStatus.
var (
	StackStatus_name = map[int32]string{
		0: "STACK_STATUS_UNSPECIFIED",
		1: "STACK_STATUS_CREATING",
		2: "STACK_STATUS_STABLE",
		3: "STACK_STATUS_AVAILABLE",
		4: "STACK_STATUS_MIGRATING",
		5: "STACK_STATUS_BROKEN",
	}
	StackStatus_value = map[string]int32{
		"STACK_STATUS_UNSPECIFIED": 0,
		"STACK_STATUS_CREATING":    1,
		"STACK_STATUS_STABLE":      2,
		"STACK_STATUS_AVAILABLE":   3,
		"STACK_STATUS_MIGRATING":   4,
		"STACK_STATUS_BROKEN":      5,
	}
)

func (x StackStatus) Enum() *StackStatus {
	p := new(StackStatus)
	*p = x
	return p
}

func (x StackStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StackStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_o5_deployer_v1_stack_proto_enumTypes[0].Descriptor()
}

func (StackStatus) Type() protoreflect.EnumType {
	return &file_o5_deployer_v1_stack_proto_enumTypes[0]
}

func (x StackStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StackStatus.Descriptor instead.
func (StackStatus) EnumDescriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{0}
}

type StackState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StackId           string             `protobuf:"bytes,1,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	Status            StackStatus        `protobuf:"varint,2,opt,name=status,proto3,enum=o5.deployer.v1.StackStatus" json:"status,omitempty"`
	CurrentDeployment *StackDeployment   `protobuf:"bytes,3,opt,name=current_deployment,json=currentDeployment,proto3,oneof" json:"current_deployment,omitempty"`
	ApplicationName   string             `protobuf:"bytes,4,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	EnvironmentName   string             `protobuf:"bytes,5,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
	QueuedDeployments []*StackDeployment `protobuf:"bytes,6,rep,name=queued_deployments,json=queuedDeployments,proto3" json:"queued_deployments,omitempty"`
}

func (x *StackState) Reset() {
	*x = StackState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackState) ProtoMessage() {}

func (x *StackState) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackState.ProtoReflect.Descriptor instead.
func (*StackState) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{0}
}

func (x *StackState) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *StackState) GetStatus() StackStatus {
	if x != nil {
		return x.Status
	}
	return StackStatus_STACK_STATUS_UNSPECIFIED
}

func (x *StackState) GetCurrentDeployment() *StackDeployment {
	if x != nil {
		return x.CurrentDeployment
	}
	return nil
}

func (x *StackState) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *StackState) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (x *StackState) GetQueuedDeployments() []*StackDeployment {
	if x != nil {
		return x.QueuedDeployments
	}
	return nil
}

type StackDeployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Version      string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *StackDeployment) Reset() {
	*x = StackDeployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackDeployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackDeployment) ProtoMessage() {}

func (x *StackDeployment) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackDeployment.ProtoReflect.Descriptor instead.
func (*StackDeployment) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{1}
}

func (x *StackDeployment) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *StackDeployment) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type StackEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *EventMetadata  `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	StackId  string          `protobuf:"bytes,2,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	Event    *StackEventType `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *StackEvent) Reset() {
	*x = StackEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackEvent) ProtoMessage() {}

func (x *StackEvent) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackEvent.ProtoReflect.Descriptor instead.
func (*StackEvent) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{2}
}

func (x *StackEvent) GetMetadata() *EventMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *StackEvent) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *StackEvent) GetEvent() *StackEventType {
	if x != nil {
		return x.Event
	}
	return nil
}

type StackEventType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*StackEventType_Triggered_
	//	*StackEventType_DeploymentCompleted_
	//	*StackEventType_DeploymentFailed_
	//	*StackEventType_Available_
	Type isStackEventType_Type `protobuf_oneof:"type"`
}

func (x *StackEventType) Reset() {
	*x = StackEventType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackEventType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackEventType) ProtoMessage() {}

func (x *StackEventType) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackEventType.ProtoReflect.Descriptor instead.
func (*StackEventType) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{3}
}

func (m *StackEventType) GetType() isStackEventType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *StackEventType) GetTriggered() *StackEventType_Triggered {
	if x, ok := x.GetType().(*StackEventType_Triggered_); ok {
		return x.Triggered
	}
	return nil
}

func (x *StackEventType) GetDeploymentCompleted() *StackEventType_DeploymentCompleted {
	if x, ok := x.GetType().(*StackEventType_DeploymentCompleted_); ok {
		return x.DeploymentCompleted
	}
	return nil
}

func (x *StackEventType) GetDeploymentFailed() *StackEventType_DeploymentFailed {
	if x, ok := x.GetType().(*StackEventType_DeploymentFailed_); ok {
		return x.DeploymentFailed
	}
	return nil
}

func (x *StackEventType) GetAvailable() *StackEventType_Available {
	if x, ok := x.GetType().(*StackEventType_Available_); ok {
		return x.Available
	}
	return nil
}

type isStackEventType_Type interface {
	isStackEventType_Type()
}

type StackEventType_Triggered_ struct {
	Triggered *StackEventType_Triggered `protobuf:"bytes,1,opt,name=triggered,proto3,oneof"`
}

type StackEventType_DeploymentCompleted_ struct {
	DeploymentCompleted *StackEventType_DeploymentCompleted `protobuf:"bytes,2,opt,name=deployment_completed,json=deploymentCompleted,proto3,oneof"`
}

type StackEventType_DeploymentFailed_ struct {
	DeploymentFailed *StackEventType_DeploymentFailed `protobuf:"bytes,3,opt,name=deployment_failed,json=deploymentFailed,proto3,oneof"`
}

type StackEventType_Available_ struct {
	Available *StackEventType_Available `protobuf:"bytes,4,opt,name=available,proto3,oneof"`
}

func (*StackEventType_Triggered_) isStackEventType_Type() {}

func (*StackEventType_DeploymentCompleted_) isStackEventType_Type() {}

func (*StackEventType_DeploymentFailed_) isStackEventType_Type() {}

func (*StackEventType_Available_) isStackEventType_Type() {}

type StackEventType_Triggered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment      *StackDeployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	ApplicationName string           `protobuf:"bytes,2,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	EnvironmentName string           `protobuf:"bytes,3,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
}

func (x *StackEventType_Triggered) Reset() {
	*x = StackEventType_Triggered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackEventType_Triggered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackEventType_Triggered) ProtoMessage() {}

func (x *StackEventType_Triggered) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackEventType_Triggered.ProtoReflect.Descriptor instead.
func (*StackEventType_Triggered) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{3, 0}
}

func (x *StackEventType_Triggered) GetDeployment() *StackDeployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *StackEventType_Triggered) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *StackEventType_Triggered) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

type StackEventType_DeploymentCompleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment *StackDeployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
}

func (x *StackEventType_DeploymentCompleted) Reset() {
	*x = StackEventType_DeploymentCompleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackEventType_DeploymentCompleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackEventType_DeploymentCompleted) ProtoMessage() {}

func (x *StackEventType_DeploymentCompleted) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackEventType_DeploymentCompleted.ProtoReflect.Descriptor instead.
func (*StackEventType_DeploymentCompleted) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{3, 1}
}

func (x *StackEventType_DeploymentCompleted) GetDeployment() *StackDeployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

type StackEventType_DeploymentFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment *StackDeployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Error      string           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StackEventType_DeploymentFailed) Reset() {
	*x = StackEventType_DeploymentFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackEventType_DeploymentFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackEventType_DeploymentFailed) ProtoMessage() {}

func (x *StackEventType_DeploymentFailed) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackEventType_DeploymentFailed.ProtoReflect.Descriptor instead.
func (*StackEventType_DeploymentFailed) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{3, 2}
}

func (x *StackEventType_DeploymentFailed) GetDeployment() *StackDeployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *StackEventType_DeploymentFailed) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type StackEventType_Available struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StackEventType_Available) Reset() {
	*x = StackEventType_Available{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_stack_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackEventType_Available) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackEventType_Available) ProtoMessage() {}

func (x *StackEventType_Available) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_stack_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackEventType_Available.ProtoReflect.Descriptor instead.
func (*StackEventType_Available) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_stack_proto_rawDescGZIP(), []int{3, 3}
}

var File_o5_deployer_v1_stack_proto protoreflect.FileDescriptor

var file_o5_deployer_v1_stack_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x35,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6f, 0x35, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x73, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x87, 0x03, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x53, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x12, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x5f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x11, 0x71, 0x75, 0x65, 0x75, 0x65, 0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x3a, 0x0d, 0xc2, 0xe8, 0x81, 0xd9, 0x02, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x0f, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x0d,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xd9, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0e, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0xda, 0xe8, 0x81,
	0xd9, 0x02, 0x02, 0x10, 0x01, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x10, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0xda, 0xe8, 0x81, 0xd9, 0x02,
	0x02, 0x18, 0x01, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x35,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0e, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0xda, 0xe8, 0x81, 0xd9, 0x02, 0x02, 0x08, 0x01, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x3a, 0x0d, 0xca, 0xe8, 0x81, 0xd9, 0x02, 0x07, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x22, 0xea, 0x05, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x09, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x12, 0x67,
	0x0a, 0x14, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6f,
	0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x13, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x5e, 0x0a, 0x11, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x48, 0x00, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x48, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x35, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x1a, 0xa2, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x3f, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x56, 0x0a, 0x13, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a,
	0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x69,
	0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x0b, 0x0a, 0x09, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xb0,
	0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x18, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x43, 0x4b,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x54, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x49, 0x47,
	0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x43,
	0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x52, 0x4f, 0x4b, 0x45, 0x4e, 0x10,
	0x05, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_deployer_v1_stack_proto_rawDescOnce sync.Once
	file_o5_deployer_v1_stack_proto_rawDescData = file_o5_deployer_v1_stack_proto_rawDesc
)

func file_o5_deployer_v1_stack_proto_rawDescGZIP() []byte {
	file_o5_deployer_v1_stack_proto_rawDescOnce.Do(func() {
		file_o5_deployer_v1_stack_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_deployer_v1_stack_proto_rawDescData)
	})
	return file_o5_deployer_v1_stack_proto_rawDescData
}

var file_o5_deployer_v1_stack_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_o5_deployer_v1_stack_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_o5_deployer_v1_stack_proto_goTypes = []interface{}{
	(StackStatus)(0),                           // 0: o5.deployer.v1.StackStatus
	(*StackState)(nil),                         // 1: o5.deployer.v1.StackState
	(*StackDeployment)(nil),                    // 2: o5.deployer.v1.StackDeployment
	(*StackEvent)(nil),                         // 3: o5.deployer.v1.StackEvent
	(*StackEventType)(nil),                     // 4: o5.deployer.v1.StackEventType
	(*StackEventType_Triggered)(nil),           // 5: o5.deployer.v1.StackEventType.Triggered
	(*StackEventType_DeploymentCompleted)(nil), // 6: o5.deployer.v1.StackEventType.DeploymentCompleted
	(*StackEventType_DeploymentFailed)(nil),    // 7: o5.deployer.v1.StackEventType.DeploymentFailed
	(*StackEventType_Available)(nil),           // 8: o5.deployer.v1.StackEventType.Available
	(*EventMetadata)(nil),                      // 9: o5.deployer.v1.EventMetadata
}
var file_o5_deployer_v1_stack_proto_depIdxs = []int32{
	0,  // 0: o5.deployer.v1.StackState.status:type_name -> o5.deployer.v1.StackStatus
	2,  // 1: o5.deployer.v1.StackState.current_deployment:type_name -> o5.deployer.v1.StackDeployment
	2,  // 2: o5.deployer.v1.StackState.queued_deployments:type_name -> o5.deployer.v1.StackDeployment
	9,  // 3: o5.deployer.v1.StackEvent.metadata:type_name -> o5.deployer.v1.EventMetadata
	4,  // 4: o5.deployer.v1.StackEvent.event:type_name -> o5.deployer.v1.StackEventType
	5,  // 5: o5.deployer.v1.StackEventType.triggered:type_name -> o5.deployer.v1.StackEventType.Triggered
	6,  // 6: o5.deployer.v1.StackEventType.deployment_completed:type_name -> o5.deployer.v1.StackEventType.DeploymentCompleted
	7,  // 7: o5.deployer.v1.StackEventType.deployment_failed:type_name -> o5.deployer.v1.StackEventType.DeploymentFailed
	8,  // 8: o5.deployer.v1.StackEventType.available:type_name -> o5.deployer.v1.StackEventType.Available
	2,  // 9: o5.deployer.v1.StackEventType.Triggered.deployment:type_name -> o5.deployer.v1.StackDeployment
	2,  // 10: o5.deployer.v1.StackEventType.DeploymentCompleted.deployment:type_name -> o5.deployer.v1.StackDeployment
	2,  // 11: o5.deployer.v1.StackEventType.DeploymentFailed.deployment:type_name -> o5.deployer.v1.StackDeployment
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_o5_deployer_v1_stack_proto_init() }
func file_o5_deployer_v1_stack_proto_init() {
	if File_o5_deployer_v1_stack_proto != nil {
		return
	}
	file_o5_deployer_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_o5_deployer_v1_stack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackDeployment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackEventType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackEventType_Triggered); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackEventType_DeploymentCompleted); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackEventType_DeploymentFailed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_o5_deployer_v1_stack_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackEventType_Available); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_o5_deployer_v1_stack_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_o5_deployer_v1_stack_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StackEventType_Triggered_)(nil),
		(*StackEventType_DeploymentCompleted_)(nil),
		(*StackEventType_DeploymentFailed_)(nil),
		(*StackEventType_Available_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_deployer_v1_stack_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_o5_deployer_v1_stack_proto_goTypes,
		DependencyIndexes: file_o5_deployer_v1_stack_proto_depIdxs,
		EnumInfos:         file_o5_deployer_v1_stack_proto_enumTypes,
		MessageInfos:      file_o5_deployer_v1_stack_proto_msgTypes,
	}.Build()
	File_o5_deployer_v1_stack_proto = out.File
	file_o5_deployer_v1_stack_proto_rawDesc = nil
	file_o5_deployer_v1_stack_proto_goTypes = nil
	file_o5_deployer_v1_stack_proto_depIdxs = nil
}
