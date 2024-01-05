// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/deployer/v1/service/command.proto

package deployer_spb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/pentops/o5-go/deployer/v1/deployer_pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TriggerDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId    string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	EnvironmentName string `protobuf:"bytes,2,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
	// Types that are assignable to Source:
	//
	//	*TriggerDeploymentRequest_Github
	Source isTriggerDeploymentRequest_Source `protobuf_oneof:"source"`
}

func (x *TriggerDeploymentRequest) Reset() {
	*x = TriggerDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_service_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerDeploymentRequest) ProtoMessage() {}

func (x *TriggerDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_service_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerDeploymentRequest.ProtoReflect.Descriptor instead.
func (*TriggerDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_service_command_proto_rawDescGZIP(), []int{0}
}

func (x *TriggerDeploymentRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *TriggerDeploymentRequest) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (m *TriggerDeploymentRequest) GetSource() isTriggerDeploymentRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *TriggerDeploymentRequest) GetGithub() *TriggerDeploymentRequest_GithubSource {
	if x, ok := x.GetSource().(*TriggerDeploymentRequest_Github); ok {
		return x.Github
	}
	return nil
}

type isTriggerDeploymentRequest_Source interface {
	isTriggerDeploymentRequest_Source()
}

type TriggerDeploymentRequest_Github struct {
	Github *TriggerDeploymentRequest_GithubSource `protobuf:"bytes,10,opt,name=github,proto3,oneof"`
}

func (*TriggerDeploymentRequest_Github) isTriggerDeploymentRequest_Source() {}

type TriggerDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TriggerDeploymentResponse) Reset() {
	*x = TriggerDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_service_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerDeploymentResponse) ProtoMessage() {}

func (x *TriggerDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_service_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerDeploymentResponse.ProtoReflect.Descriptor instead.
func (*TriggerDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_service_command_proto_rawDescGZIP(), []int{1}
}

type TriggerDeploymentRequest_GithubSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Repo   string `protobuf:"bytes,3,opt,name=repo,proto3" json:"repo,omitempty"`
	Commit string `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *TriggerDeploymentRequest_GithubSource) Reset() {
	*x = TriggerDeploymentRequest_GithubSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_deployer_v1_service_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriggerDeploymentRequest_GithubSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriggerDeploymentRequest_GithubSource) ProtoMessage() {}

func (x *TriggerDeploymentRequest_GithubSource) ProtoReflect() protoreflect.Message {
	mi := &file_o5_deployer_v1_service_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriggerDeploymentRequest_GithubSource.ProtoReflect.Descriptor instead.
func (*TriggerDeploymentRequest_GithubSource) Descriptor() ([]byte, []int) {
	return file_o5_deployer_v1_service_command_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TriggerDeploymentRequest_GithubSource) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *TriggerDeploymentRequest_GithubSource) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *TriggerDeploymentRequest_GithubSource) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

var File_o5_deployer_v1_service_command_proto protoreflect.FileDescriptor

var file_o5_deployer_v1_service_command_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6f, 0x35, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6f, 0x35, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x18, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x31, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x1a, 0x68,
	0x0a, 0x0c, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x1e, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xcc, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xaf, 0x01, 0x0a,
	0x11, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x30, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a,
	0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e,
	0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f,
	0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_deployer_v1_service_command_proto_rawDescOnce sync.Once
	file_o5_deployer_v1_service_command_proto_rawDescData = file_o5_deployer_v1_service_command_proto_rawDesc
)

func file_o5_deployer_v1_service_command_proto_rawDescGZIP() []byte {
	file_o5_deployer_v1_service_command_proto_rawDescOnce.Do(func() {
		file_o5_deployer_v1_service_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_deployer_v1_service_command_proto_rawDescData)
	})
	return file_o5_deployer_v1_service_command_proto_rawDescData
}

var file_o5_deployer_v1_service_command_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_o5_deployer_v1_service_command_proto_goTypes = []interface{}{
	(*TriggerDeploymentRequest)(nil),              // 0: o5.deployer.v1.service.TriggerDeploymentRequest
	(*TriggerDeploymentResponse)(nil),             // 1: o5.deployer.v1.service.TriggerDeploymentResponse
	(*TriggerDeploymentRequest_GithubSource)(nil), // 2: o5.deployer.v1.service.TriggerDeploymentRequest.GithubSource
}
var file_o5_deployer_v1_service_command_proto_depIdxs = []int32{
	2, // 0: o5.deployer.v1.service.TriggerDeploymentRequest.github:type_name -> o5.deployer.v1.service.TriggerDeploymentRequest.GithubSource
	0, // 1: o5.deployer.v1.service.DeploymentCommandService.TriggerDeployment:input_type -> o5.deployer.v1.service.TriggerDeploymentRequest
	1, // 2: o5.deployer.v1.service.DeploymentCommandService.TriggerDeployment:output_type -> o5.deployer.v1.service.TriggerDeploymentResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_o5_deployer_v1_service_command_proto_init() }
func file_o5_deployer_v1_service_command_proto_init() {
	if File_o5_deployer_v1_service_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_o5_deployer_v1_service_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerDeploymentRequest); i {
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
		file_o5_deployer_v1_service_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerDeploymentResponse); i {
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
		file_o5_deployer_v1_service_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriggerDeploymentRequest_GithubSource); i {
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
	file_o5_deployer_v1_service_command_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TriggerDeploymentRequest_Github)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_deployer_v1_service_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_o5_deployer_v1_service_command_proto_goTypes,
		DependencyIndexes: file_o5_deployer_v1_service_command_proto_depIdxs,
		MessageInfos:      file_o5_deployer_v1_service_command_proto_msgTypes,
	}.Build()
	File_o5_deployer_v1_service_command_proto = out.File
	file_o5_deployer_v1_service_command_proto_rawDesc = nil
	file_o5_deployer_v1_service_command_proto_goTypes = nil
	file_o5_deployer_v1_service_command_proto_depIdxs = nil
}
