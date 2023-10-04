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

type isEnvironment_Provider interface {
	isEnvironment_Provider()
}

type Environment_Aws struct {
	Aws *AWS `protobuf:"bytes,10,opt,name=aws,proto3,oneof"`
}

func (*Environment_Aws) isEnvironment_Provider() {}

type AWS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenerArn          string `protobuf:"bytes,1,opt,name=listener_arn,json=listenerArn,proto3" json:"listener_arn,omitempty"`
	EcsClusterName       string `protobuf:"bytes,2,opt,name=ecs_cluster_name,json=ecsClusterName,proto3" json:"ecs_cluster_name,omitempty"`
	EcsRepo              string `protobuf:"bytes,3,opt,name=ecs_repo,json=ecsRepo,proto3" json:"ecs_repo,omitempty"`
	EcsTaskExecutionRole string `protobuf:"bytes,4,opt,name=ecs_task_execution_role,json=ecsTaskExecutionRole,proto3" json:"ecs_task_execution_role,omitempty"`
	VpcId                string `protobuf:"bytes,5,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
}

func (x *AWS) Reset() {
	*x = AWS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_environment_v1_environment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWS) ProtoMessage() {}

func (x *AWS) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AWS.ProtoReflect.Descriptor instead.
func (*AWS) Descriptor() ([]byte, []int) {
	return file_o5_environment_v1_environment_proto_rawDescGZIP(), []int{1}
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

var File_o5_environment_v1_environment_proto protoreflect.FileDescriptor

var file_o5_environment_v1_environment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6f, 0x35, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x62, 0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77, 0x73,
	0x42, 0x0a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xbb, 0x01, 0x0a,
	0x03, 0x41, 0x57, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x41, 0x72, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x63, 0x73, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x65, 0x63, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x63, 0x73, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x63, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x35, 0x0a, 0x17,
	0x65, 0x63, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65,
	0x63, 0x73, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73,
	0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_o5_environment_v1_environment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_o5_environment_v1_environment_proto_goTypes = []interface{}{
	(*Environment)(nil), // 0: o5.environment.v1.Environment
	(*AWS)(nil),         // 1: o5.environment.v1.AWS
}
var file_o5_environment_v1_environment_proto_depIdxs = []int32{
	1, // 0: o5.environment.v1.Environment.aws:type_name -> o5.environment.v1.AWS
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	file_o5_environment_v1_environment_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Environment_Aws)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_environment_v1_environment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
