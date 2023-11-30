// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/deployer/v1/topic/deployer.proto

package deployer_tpb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	deployer_pb "github.com/pentops/o5-go/deployer/v1/deployer_pb"
	_ "github.com/pentops/o5-go/messaging/v1/messaging_pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestDeploymentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId string                      `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Spec         *deployer_pb.DeploymentSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *RequestDeploymentMessage) Reset() {
	*x = RequestDeploymentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestDeploymentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestDeploymentMessage) ProtoMessage() {}

func (x *RequestDeploymentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestDeploymentMessage.ProtoReflect.Descriptor instead.
func (*RequestDeploymentMessage) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP(), []int{0}
}

func (x *RequestDeploymentMessage) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *RequestDeploymentMessage) GetSpec() *deployer_pb.DeploymentSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeploymentCompleteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	StackId         string `protobuf:"bytes,2,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	Version         string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	EnvironmentName string `protobuf:"bytes,4,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
	ApplicationName string `protobuf:"bytes,5,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
}

func (x *DeploymentCompleteMessage) Reset() {
	*x = DeploymentCompleteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentCompleteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentCompleteMessage) ProtoMessage() {}

func (x *DeploymentCompleteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentCompleteMessage.ProtoReflect.Descriptor instead.
func (*DeploymentCompleteMessage) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentCompleteMessage) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DeploymentCompleteMessage) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *DeploymentCompleteMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DeploymentCompleteMessage) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (x *DeploymentCompleteMessage) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

type DeploymentFailedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	StackId         string `protobuf:"bytes,2,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	Version         string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	EnvironmentName string `protobuf:"bytes,4,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
	ApplicationName string `protobuf:"bytes,5,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	Error           string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeploymentFailedMessage) Reset() {
	*x = DeploymentFailedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentFailedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentFailedMessage) ProtoMessage() {}

func (x *DeploymentFailedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentFailedMessage.ProtoReflect.Descriptor instead.
func (*DeploymentFailedMessage) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP(), []int{2}
}

func (x *DeploymentFailedMessage) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DeploymentFailedMessage) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *DeploymentFailedMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DeploymentFailedMessage) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (x *DeploymentFailedMessage) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *DeploymentFailedMessage) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type TriggerDeploymentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	StackId         string `protobuf:"bytes,2,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	Version         string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	EnvironmentName string `protobuf:"bytes,4,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
	ApplicationName string `protobuf:"bytes,5,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
}

func (x *TriggerDeploymentMessage) Reset() {
	*x = TriggerDeploymentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerDeploymentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerDeploymentMessage) ProtoMessage() {}

func (x *TriggerDeploymentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerDeploymentMessage.ProtoReflect.Descriptor instead.
func (*TriggerDeploymentMessage) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP(), []int{3}
}

func (x *TriggerDeploymentMessage) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *TriggerDeploymentMessage) GetStackId() string {
	if x != nil {
		return x.StackId
	}
	return ""
}

func (x *TriggerDeploymentMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TriggerDeploymentMessage) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (x *TriggerDeploymentMessage) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

type StackStatusChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StackId   *StackID                   `protobuf:"bytes,1,opt,name=stack_id,json=stackId,proto3" json:"stack_id,omitempty"`
	Status    string                     `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Outputs   []*deployer_pb.KeyValue    `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Lifecycle deployer_pb.StackLifecycle `protobuf:"varint,4,opt,name=lifecycle,proto3,enum=o5.deployer.v1.StackLifecycle" json:"lifecycle,omitempty"`
}

func (x *StackStatusChangedMessage) Reset() {
	*x = StackStatusChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackStatusChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackStatusChangedMessage) ProtoMessage() {}

func (x *StackStatusChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackStatusChangedMessage.ProtoReflect.Descriptor instead.
func (*StackStatusChangedMessage) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP(), []int{4}
}

func (x *StackStatusChangedMessage) GetStackId() *StackID {
	if x != nil {
		return x.StackId
	}
	return nil
}

func (x *StackStatusChangedMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StackStatusChangedMessage) GetOutputs() []*deployer_pb.KeyValue {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *StackStatusChangedMessage) GetLifecycle() deployer_pb.StackLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return deployer_pb.StackLifecycle(0)
}

type MigrationStatusChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MigrationId  string                              `protobuf:"bytes,1,opt,name=migration_id,json=migrationId,proto3" json:"migration_id,omitempty"`
	DeploymentId string                              `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Status       deployer_pb.DatabaseMigrationStatus `protobuf:"varint,3,opt,name=status,proto3,enum=o5.deployer.v1.DatabaseMigrationStatus" json:"status,omitempty"`
	Error        *string                             `protobuf:"bytes,4,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *MigrationStatusChangedMessage) Reset() {
	*x = MigrationStatusChangedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationStatusChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationStatusChangedMessage) ProtoMessage() {}

func (x *MigrationStatusChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_topic_deployer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationStatusChangedMessage.ProtoReflect.Descriptor instead.
func (*MigrationStatusChangedMessage) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP(), []int{5}
}

func (x *MigrationStatusChangedMessage) GetMigrationId() string {
	if x != nil {
		return x.MigrationId
	}
	return ""
}

func (x *MigrationStatusChangedMessage) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *MigrationStatusChangedMessage) GetStatus() deployer_pb.DatabaseMigrationStatus {
	if x != nil {
		return x.Status
	}
	return deployer_pb.DatabaseMigrationStatus(0)
}

func (x *MigrationStatusChangedMessage) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_o5_deployer_v1_topic_deployer_proto protoreflect.FileDescriptor

var file_o5_deployer_v1_topic_deployer_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2f, 0x61,
	0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6f, 0x35, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x18, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0xdf, 0x01, 0x0a, 0x19, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xf3, 0x01, 0x0a,
	0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0xde, 0x01, 0x0a, 0x18, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2d, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x49, 0x44, 0x52, 0x07, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x09, 0x6c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x35, 0x2e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x09, 0x6c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x1d, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xf2, 0x04, 0x0a, 0x0d, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x5d, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x6f,
	0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2f, 0x2e, 0x6f,
	0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x2d, 0x2e, 0x6f, 0x35,
	0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x11, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x2e, 0x6f, 0x35, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x2f, 0x2e, 0x6f, 0x35, 0x2e, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x16, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x33,
	0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x1a, 0x1b, 0xd2,
	0xa2, 0xf5, 0xe4, 0x02, 0x15, 0x12, 0x13, 0x0a, 0x11, 0x6f, 0x35, 0x2d, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73,
	0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_deployer_v1_topic_deployer_proto_rawDescOnce sync.Once
	file_o5_deployer_v1_topic_deployer_proto_rawDescData = file_o5_deployer_v1_topic_deployer_proto_rawDesc
)

func file_o5_deployer_v1_topic_deployer_proto_rawDescGZIP() []byte {
	file_o5_deployer_v1_topic_deployer_proto_rawDescOnce.Do(func() {
		file_o5_deployer_v1_topic_deployer_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_deployer_v1_topic_deployer_proto_rawDescData)
	})
	return file_o5_deployer_v1_topic_deployer_proto_rawDescData
}

var file_o5_deployer_v1_topic_deployer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_o5_deployer_v1_topic_deployer_proto_goTypes = []interface{}{
	(*RequestDeploymentMessage)(nil),         // 0: o5.deployer.v1.topic.RequestDeploymentMessage
	(*DeploymentCompleteMessage)(nil),        // 1: o5.deployer.v1.topic.DeploymentCompleteMessage
	(*DeploymentFailedMessage)(nil),          // 2: o5.deployer.v1.topic.DeploymentFailedMessage
	(*TriggerDeploymentMessage)(nil),         // 3: o5.deployer.v1.topic.TriggerDeploymentMessage
	(*StackStatusChangedMessage)(nil),        // 4: o5.deployer.v1.topic.StackStatusChangedMessage
	(*MigrationStatusChangedMessage)(nil),    // 5: o5.deployer.v1.topic.MigrationStatusChangedMessage
	(*deployer_pb.DeploymentSpec)(nil),       // 6: o5.deployer.v1.DeploymentSpec
	(*StackID)(nil),                          // 7: o5.deployer.v1.topic.StackID
	(*deployer_pb.KeyValue)(nil),             // 8: o5.deployer.v1.KeyValue
	(deployer_pb.StackLifecycle)(0),          // 9: o5.deployer.v1.StackLifecycle
	(deployer_pb.DatabaseMigrationStatus)(0), // 10: o5.deployer.v1.DatabaseMigrationStatus
	(*emptypb.Empty)(nil),                    // 11: google.protobuf.Empty
}
var file_o5_deployer_v1_topic_deployer_proto_depIdxs = []int32{
	6,  // 0: o5.deployer.v1.topic.RequestDeploymentMessage.spec:type_name -> o5.deployer.v1.DeploymentSpec
	7,  // 1: o5.deployer.v1.topic.StackStatusChangedMessage.stack_id:type_name -> o5.deployer.v1.topic.StackID
	8,  // 2: o5.deployer.v1.topic.StackStatusChangedMessage.outputs:type_name -> o5.deployer.v1.KeyValue
	9,  // 3: o5.deployer.v1.topic.StackStatusChangedMessage.lifecycle:type_name -> o5.deployer.v1.StackLifecycle
	10, // 4: o5.deployer.v1.topic.MigrationStatusChangedMessage.status:type_name -> o5.deployer.v1.DatabaseMigrationStatus
	0,  // 5: o5.deployer.v1.topic.DeployerTopic.RequestDeployment:input_type -> o5.deployer.v1.topic.RequestDeploymentMessage
	1,  // 6: o5.deployer.v1.topic.DeployerTopic.DeploymentComplete:input_type -> o5.deployer.v1.topic.DeploymentCompleteMessage
	2,  // 7: o5.deployer.v1.topic.DeployerTopic.DeploymentFailed:input_type -> o5.deployer.v1.topic.DeploymentFailedMessage
	3,  // 8: o5.deployer.v1.topic.DeployerTopic.TriggerDeployment:input_type -> o5.deployer.v1.topic.TriggerDeploymentMessage
	4,  // 9: o5.deployer.v1.topic.DeployerTopic.StackStatusChanged:input_type -> o5.deployer.v1.topic.StackStatusChangedMessage
	5,  // 10: o5.deployer.v1.topic.DeployerTopic.MigrationStatusChanged:input_type -> o5.deployer.v1.topic.MigrationStatusChangedMessage
	11, // 11: o5.deployer.v1.topic.DeployerTopic.RequestDeployment:output_type -> google.protobuf.Empty
	11, // 12: o5.deployer.v1.topic.DeployerTopic.DeploymentComplete:output_type -> google.protobuf.Empty
	11, // 13: o5.deployer.v1.topic.DeployerTopic.DeploymentFailed:output_type -> google.protobuf.Empty
	11, // 14: o5.deployer.v1.topic.DeployerTopic.TriggerDeployment:output_type -> google.protobuf.Empty
	11, // 15: o5.deployer.v1.topic.DeployerTopic.StackStatusChanged:output_type -> google.protobuf.Empty
	11, // 16: o5.deployer.v1.topic.DeployerTopic.MigrationStatusChanged:output_type -> google.protobuf.Empty
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_o5_deployer_v1_topic_deployer_proto_init() }
func file_o5_deployer_v1_topic_deployer_proto_init() {
	if File_o5_deployer_v1_topic_deployer_proto != nil {
		return
	}
	file_o5_deployer_v1_topic_aws_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_o5_deployer_v1_topic_deployer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestDeploymentMessage); i {
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
		file_o5_deployer_v1_topic_deployer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentCompleteMessage); i {
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
		file_o5_deployer_v1_topic_deployer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentFailedMessage); i {
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
		file_o5_deployer_v1_topic_deployer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerDeploymentMessage); i {
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
		file_o5_deployer_v1_topic_deployer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackStatusChangedMessage); i {
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
		file_o5_deployer_v1_topic_deployer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationStatusChangedMessage); i {
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
	file_o5_deployer_v1_topic_deployer_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_deployer_v1_topic_deployer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_o5_deployer_v1_topic_deployer_proto_goTypes,
		DependencyIndexes: file_o5_deployer_v1_topic_deployer_proto_depIdxs,
		MessageInfos:      file_o5_deployer_v1_topic_deployer_proto_msgTypes,
	}.Build()
	File_o5_deployer_v1_topic_deployer_proto = out.File
	file_o5_deployer_v1_topic_deployer_proto_rawDesc = nil
	file_o5_deployer_v1_topic_deployer_proto_goTypes = nil
	file_o5_deployer_v1_topic_deployer_proto_depIdxs = nil
}
