// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/environment/v1/environment.proto

package environment_pb

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

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// Types that are assignable to Provider:
	//
	//	*Environment_Aws
	Provider isEnvironment_Provider `protobuf_oneof:"provider"`
	Vars     []*CustomVariable      `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_o5_environment_v1_environment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{0}
}

func (x *Environment) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (m *Environment) GetProvider() isEnvironment_Provider {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (x *Environment) GetAws() *AWS {
	if x, ok := x.GetProvider().(*Environment_Aws); ok {
		return x.Aws
	}
	return nil
}

func (x *Environment) GetVars() []*CustomVariable {
	if x != nil {
		return x.Vars
	}
	return nil
}

type isEnvironment_Provider interface {
	isEnvironment_Provider()
}

type Environment_Aws struct {
	Aws *AWS `protobuf:"bytes,10,opt,name=aws,proto3,oneof"`
}

func (*Environment_Aws) isEnvironment_Provider() {}

type CustomVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// oneof, but logically, because repeated oneof isn't a thing
	//
	// Types that are assignable to Src:
	//
	//	*CustomVariable_Value
	//	*CustomVariable_Join_
	Src isCustomVariable_Src `protobuf_oneof:"src"`
}

func (x *CustomVariable) Reset() {
	*x = CustomVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomVariable) ProtoMessage() {}

func (x *CustomVariable) ProtoReflect() protoreflect.Message {
	mi := &file_o5_environment_v1_environment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomVariable.ProtoReflect.Descriptor instead.
func (*CustomVariable) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{1}
}

func (x *CustomVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *CustomVariable) GetSrc() isCustomVariable_Src {
	if m != nil {
		return m.Src
	}
	return nil
}

func (x *CustomVariable) GetValue() string {
	if x, ok := x.GetSrc().(*CustomVariable_Value); ok {
		return x.Value
	}
	return ""
}

func (x *CustomVariable) GetJoin() *CustomVariable_Join {
	if x, ok := x.GetSrc().(*CustomVariable_Join_); ok {
		return x.Join
	}
	return nil
}

type isCustomVariable_Src interface {
	isCustomVariable_Src()
}

type CustomVariable_Value struct {
	Value string `protobuf:"bytes,10,opt,name=value,proto3,oneof"`
}

type CustomVariable_Join_ struct {
	Join *CustomVariable_Join `protobuf:"bytes,11,opt,name=join,proto3,oneof"`
}

func (*CustomVariable_Value) isCustomVariable_Src() {}

func (*CustomVariable_Join_) isCustomVariable_Src() {}

type AWS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenerArn          string `protobuf:"bytes,1,opt,name=listener_arn,json=listenerArn,proto3" json:"listener_arn,omitempty"`
	EcsClusterName       string `protobuf:"bytes,2,opt,name=ecs_cluster_name,json=ecsClusterName,proto3" json:"ecs_cluster_name,omitempty"`
	EcsRepo              string `protobuf:"bytes,3,opt,name=ecs_repo,json=ecsRepo,proto3" json:"ecs_repo,omitempty"`
	EcsTaskExecutionRole string `protobuf:"bytes,4,opt,name=ecs_task_execution_role,json=ecsTaskExecutionRole,proto3" json:"ecs_task_execution_role,omitempty"`
	VpcId                string `protobuf:"bytes,5,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	// The role to assume when deploying to this environment
	O5DeployerAssumeRole string `protobuf:"bytes,9,opt,name=o5_deployer_assume_role,json=o5DeployerAssumeRole,proto3" json:"o5_deployer_assume_role,omitempty"`
	// The roles which a deployer service (e.g. o5 itself) should be allowed to
	// assume, when grant_meta_deploy_permissions is set in the application config
	O5MetaDeployerAssumeRoleArns []string `protobuf:"bytes,10,rep,name=o5_meta_deployer_assume_role_arns,json=o5MetaDeployerAssumeRoleArns,proto3" json:"o5_meta_deployer_assume_role_arns,omitempty"`
	// when set, adds a host header rule exactly as defined to all listener rules.
	// This allows re-use of the same listener for multiple (related) environments.
	HostHeader       *string    `protobuf:"bytes,6,opt,name=host_header,json=hostHeader,proto3,oneof" json:"host_header,omitempty"`
	RdsHosts         []*RDSHost `protobuf:"bytes,7,rep,name=rds_hosts,json=rdsHosts,proto3" json:"rds_hosts,omitempty"`
	EnvironmentLinks []*AWSLink `protobuf:"bytes,8,rep,name=environment_links,json=environmentLinks,proto3" json:"environment_links,omitempty"`
}

func (x *AWS) Reset() {
	*x = AWS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWS) ProtoMessage() {}

func (x *AWS) ProtoReflect() protoreflect.Message {
	mi := &file_o5_environment_v1_environment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWS.ProtoReflect.Descriptor instead.
func (*AWS) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{2}
}

func (x *AWS) GetListenerArn() string {
	if x != nil {
		return x.ListenerArn
	}
	return ""
}

func (x *AWS) GetEcsClusterName() string {
	if x != nil {
		return x.EcsClusterName
	}
	return ""
}

func (x *AWS) GetEcsRepo() string {
	if x != nil {
		return x.EcsRepo
	}
	return ""
}

func (x *AWS) GetEcsTaskExecutionRole() string {
	if x != nil {
		return x.EcsTaskExecutionRole
	}
	return ""
}

func (x *AWS) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *AWS) GetO5DeployerAssumeRole() string {
	if x != nil {
		return x.O5DeployerAssumeRole
	}
	return ""
}

func (x *AWS) GetO5MetaDeployerAssumeRoleArns() []string {
	if x != nil {
		return x.O5MetaDeployerAssumeRoleArns
	}
	return nil
}

func (x *AWS) GetHostHeader() string {
	if x != nil && x.HostHeader != nil {
		return *x.HostHeader
	}
	return ""
}

func (x *AWS) GetRdsHosts() []*RDSHost {
	if x != nil {
		return x.RdsHosts
	}
	return nil
}

func (x *AWS) GetEnvironmentLinks() []*AWSLink {
	if x != nil {
		return x.EnvironmentLinks
	}
	return nil
}

type AWSLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName  string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"` // used to match
	SnsPrefix string `protobuf:"bytes,2,opt,name=sns_prefix,json=snsPrefix,proto3" json:"sns_prefix,omitempty"`
}

func (x *AWSLink) Reset() {
	*x = AWSLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSLink) ProtoMessage() {}

func (x *AWSLink) ProtoReflect() protoreflect.Message {
	mi := &file_o5_environment_v1_environment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSLink.ProtoReflect.Descriptor instead.
func (*AWSLink) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{3}
}

func (x *AWSLink) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *AWSLink) GetSnsPrefix() string {
	if x != nil {
		return x.SnsPrefix
	}
	return ""
}

type RDSHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerGroup string `protobuf:"bytes,1,opt,name=server_group,json=serverGroup,proto3" json:"server_group,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
}

func (x *RDSHost) Reset() {
	*x = RDSHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RDSHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RDSHost) ProtoMessage() {}

func (x *RDSHost) ProtoReflect() protoreflect.Message {
	mi := &file_o5_environment_v1_environment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RDSHost.ProtoReflect.Descriptor instead.
func (*RDSHost) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{4}
}

func (x *RDSHost) GetServerGroup() string {
	if x != nil {
		return x.ServerGroup
	}
	return ""
}

func (x *RDSHost) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

type CustomVariable_Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Delimiter string   `protobuf:"bytes,1,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	Values    []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *CustomVariable_Join) Reset() {
	*x = CustomVariable_Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomVariable_Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomVariable_Join) ProtoMessage() {}

func (x *CustomVariable_Join) ProtoReflect() protoreflect.Message {
	mi := &file_o5_environment_v1_environment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomVariable_Join.ProtoReflect.Descriptor instead.
func (*CustomVariable_Join) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CustomVariable_Join) GetDelimiter() string {
	if x != nil {
		return x.Delimiter
	}
	return ""
}

func (x *CustomVariable_Join) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_o5_environment_v1_environment_proto protoreflect.FileDescriptor

var file_o5_environment_v1_environment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x35, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77,
	0x73, 0x12, 0x35, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x04, 0x76, 0x61, 0x72, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x22, 0xbf, 0x01, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x04, 0x6a, 0x6f, 0x69,
	0x6e, 0x1a, 0x3c, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42,
	0x05, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x22, 0xf3, 0x03, 0x0a, 0x03, 0x41, 0x57, 0x53, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x41, 0x72,
	0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x63, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x63, 0x73,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x63, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x63, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x35, 0x0a, 0x17, 0x65, 0x63, 0x73, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x63, 0x73, 0x54, 0x61, 0x73, 0x6b,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x70, 0x63, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x6f, 0x35, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6f, 0x35, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x21, 0x6f,
	0x35, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f,
	0x61, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1c, 0x6f, 0x35, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x41, 0x73, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x41, 0x72, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x68, 0x6f, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x64,
	0x73, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x44, 0x53, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x72, 0x64, 0x73, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x47, 0x0a, 0x11, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x10, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x07,
	0x41, 0x57, 0x53, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6e, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6e, 0x73, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x22, 0x4d, 0x0a, 0x07, 0x52, 0x44, 0x53, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_environment_v1_environment_proto_rawDescOnce sync.Once
	file_o5_environment_v1_environment_proto_rawDescData = file_o5_environment_v1_environment_proto_rawDesc
)

func file_o5_environment_v1_environment_proto_rawDescGZIP() []byte {
	file_o5_environment_v1_environment_proto_rawDescOnce.Do(func() {
		file_o5_environment_v1_environment_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_environment_v1_environment_proto_rawDescData)
	})
	return file_o5_environment_v1_environment_proto_rawDescData
}

var file_o5_environment_v1_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_o5_environment_v1_environment_proto_goTypes = []interface{}{
	(*Environment)(nil),         // 0: o5.environment.v1.Environment
	(*CustomVariable)(nil),      // 1: o5.environment.v1.CustomVariable
	(*AWS)(nil),                 // 2: o5.environment.v1.AWS
	(*AWSLink)(nil),             // 3: o5.environment.v1.AWSLink
	(*RDSHost)(nil),             // 4: o5.environment.v1.RDSHost
	(*CustomVariable_Join)(nil), // 5: o5.environment.v1.CustomVariable.Join
}
var file_o5_environment_v1_environment_proto_depIdxs = []int32{
	2, // 0: o5.environment.v1.Environment.aws:type_name -> o5.environment.v1.AWS
	1, // 1: o5.environment.v1.Environment.vars:type_name -> o5.environment.v1.CustomVariable
	5, // 2: o5.environment.v1.CustomVariable.join:type_name -> o5.environment.v1.CustomVariable.Join
	4, // 3: o5.environment.v1.AWS.rds_hosts:type_name -> o5.environment.v1.RDSHost
	3, // 4: o5.environment.v1.AWS.environment_links:type_name -> o5.environment.v1.AWSLink
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_o5_environment_v1_environment_proto_init() }
func file_o5_environment_v1_environment_proto_init() {
	if File_o5_environment_v1_environment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_o5_environment_v1_environment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_o5_environment_v1_environment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomVariable); i {
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
		file_o5_environment_v1_environment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWS); i {
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
		file_o5_environment_v1_environment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSLink); i {
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
		file_o5_environment_v1_environment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RDSHost); i {
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
		file_o5_environment_v1_environment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomVariable_Join); i {
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
	file_o5_environment_v1_environment_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Environment_Aws)(nil),
	}
	file_o5_environment_v1_environment_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CustomVariable_Value)(nil),
		(*CustomVariable_Join_)(nil),
	}
	file_o5_environment_v1_environment_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_environment_v1_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_o5_environment_v1_environment_proto_goTypes,
		DependencyIndexes: file_o5_environment_v1_environment_proto_depIdxs,
		MessageInfos:      file_o5_environment_v1_environment_proto_msgTypes,
	}.Build()
	File_o5_environment_v1_environment_proto = out.File
	file_o5_environment_v1_environment_proto_rawDesc = nil
	file_o5_environment_v1_environment_proto_goTypes = nil
	file_o5_environment_v1_environment_proto_depIdxs = nil
}
