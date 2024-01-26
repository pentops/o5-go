// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/messaging/v1/annotations.proto

package messaging_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*Config_Broadcast
	//	*Config_Unicast
	//	*Config_Request
	//	*Config_Reply
	Type isConfig_Type `protobuf_oneof:"type"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_messaging_v1_annotations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_o5_messaging_v1_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_o5_messaging_v1_annotations_proto_rawDescGZIP(), []int{0}
}

func (m *Config) GetType() isConfig_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Config) GetBroadcast() *BroadcastConfig {
	if x, ok := x.GetType().(*Config_Broadcast); ok {
		return x.Broadcast
	}
	return nil
}

func (x *Config) GetUnicast() *UnicastConfig {
	if x, ok := x.GetType().(*Config_Unicast); ok {
		return x.Unicast
	}
	return nil
}

func (x *Config) GetRequest() *RequestConfig {
	if x, ok := x.GetType().(*Config_Request); ok {
		return x.Request
	}
	return nil
}

func (x *Config) GetReply() *ReplyConfig {
	if x, ok := x.GetType().(*Config_Reply); ok {
		return x.Reply
	}
	return nil
}

type isConfig_Type interface {
	isConfig_Type()
}

type Config_Broadcast struct {
	Broadcast *BroadcastConfig `protobuf:"bytes,1,opt,name=broadcast,proto3,oneof"`
}

type Config_Unicast struct {
	Unicast *UnicastConfig `protobuf:"bytes,2,opt,name=unicast,proto3,oneof"`
}

type Config_Request struct {
	Request *RequestConfig `protobuf:"bytes,3,opt,name=request,proto3,oneof"`
}

type Config_Reply struct {
	Reply *ReplyConfig `protobuf:"bytes,4,opt,name=reply,proto3,oneof"`
}

func (*Config_Broadcast) isConfig_Type() {}

func (*Config_Unicast) isConfig_Type() {}

func (*Config_Request) isConfig_Type() {}

func (*Config_Reply) isConfig_Type() {}

type BroadcastConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BroadcastConfig) Reset() {
	*x = BroadcastConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_messaging_v1_annotations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastConfig) ProtoMessage() {}

func (x *BroadcastConfig) ProtoReflect() protoreflect.Message {
	mi := &file_o5_messaging_v1_annotations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastConfig.ProtoReflect.Descriptor instead.
func (*BroadcastConfig) Descriptor() ([]byte, []int) {
	return file_o5_messaging_v1_annotations_proto_rawDescGZIP(), []int{1}
}

func (x *BroadcastConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UnicastConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UnicastConfig) Reset() {
	*x = UnicastConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_messaging_v1_annotations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnicastConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnicastConfig) ProtoMessage() {}

func (x *UnicastConfig) ProtoReflect() protoreflect.Message {
	mi := &file_o5_messaging_v1_annotations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnicastConfig.ProtoReflect.Descriptor instead.
func (*UnicastConfig) Descriptor() ([]byte, []int) {
	return file_o5_messaging_v1_annotations_proto_rawDescGZIP(), []int{2}
}

func (x *UnicastConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RequestConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the queue group (i.e. prefix), should match ReplyConfig.name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RequestConfig) Reset() {
	*x = RequestConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_messaging_v1_annotations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestConfig) ProtoMessage() {}

func (x *RequestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_o5_messaging_v1_annotations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestConfig.ProtoReflect.Descriptor instead.
func (*RequestConfig) Descriptor() ([]byte, []int) {
	return file_o5_messaging_v1_annotations_proto_rawDescGZIP(), []int{3}
}

func (x *RequestConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ReplyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the queue group (i.e. prefix), should match RequestConfig.name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ReplyConfig) Reset() {
	*x = ReplyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_messaging_v1_annotations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyConfig) ProtoMessage() {}

func (x *ReplyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_o5_messaging_v1_annotations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyConfig.ProtoReflect.Descriptor instead.
func (*ReplyConfig) Descriptor() ([]byte, []int) {
	return file_o5_messaging_v1_annotations_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var file_o5_messaging_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*Config)(nil),
		Field:         93563434,
		Name:          "o5.messaging.v1.config",
		Tag:           "bytes,93563434,opt,name=config",
		Filename:      "o5/messaging/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional o5.messaging.v1.Config config = 93563434;
	E_Config = &file_o5_messaging_v1_annotations_proto_extTypes[0]
)

var File_o5_messaging_v1_annotations_proto protoreflect.FileDescriptor

var file_o5_messaging_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x35, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6f, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x40, 0x0a, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x09, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6f, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x35, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x0f, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x23, 0x0a, 0x0d, 0x55, 0x6e, 0x69, 0x63, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x53, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xaa, 0xd4, 0xce, 0x2c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_messaging_v1_annotations_proto_rawDescOnce sync.Once
	file_o5_messaging_v1_annotations_proto_rawDescData = file_o5_messaging_v1_annotations_proto_rawDesc
)

func file_o5_messaging_v1_annotations_proto_rawDescGZIP() []byte {
	file_o5_messaging_v1_annotations_proto_rawDescOnce.Do(func() {
		file_o5_messaging_v1_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_messaging_v1_annotations_proto_rawDescData)
	})
	return file_o5_messaging_v1_annotations_proto_rawDescData
}

var file_o5_messaging_v1_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_o5_messaging_v1_annotations_proto_goTypes = []interface{}{
	(*Config)(nil),                      // 0: o5.messaging.v1.Config
	(*BroadcastConfig)(nil),             // 1: o5.messaging.v1.BroadcastConfig
	(*UnicastConfig)(nil),               // 2: o5.messaging.v1.UnicastConfig
	(*RequestConfig)(nil),               // 3: o5.messaging.v1.RequestConfig
	(*ReplyConfig)(nil),                 // 4: o5.messaging.v1.ReplyConfig
	(*descriptorpb.ServiceOptions)(nil), // 5: google.protobuf.ServiceOptions
}
var file_o5_messaging_v1_annotations_proto_depIdxs = []int32{
	1, // 0: o5.messaging.v1.Config.broadcast:type_name -> o5.messaging.v1.BroadcastConfig
	2, // 1: o5.messaging.v1.Config.unicast:type_name -> o5.messaging.v1.UnicastConfig
	3, // 2: o5.messaging.v1.Config.request:type_name -> o5.messaging.v1.RequestConfig
	4, // 3: o5.messaging.v1.Config.reply:type_name -> o5.messaging.v1.ReplyConfig
	5, // 4: o5.messaging.v1.config:extendee -> google.protobuf.ServiceOptions
	0, // 5: o5.messaging.v1.config:type_name -> o5.messaging.v1.Config
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	5, // [5:6] is the sub-list for extension type_name
	4, // [4:5] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_o5_messaging_v1_annotations_proto_init() }
func file_o5_messaging_v1_annotations_proto_init() {
	if File_o5_messaging_v1_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_o5_messaging_v1_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_o5_messaging_v1_annotations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastConfig); i {
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
		file_o5_messaging_v1_annotations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnicastConfig); i {
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
		file_o5_messaging_v1_annotations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestConfig); i {
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
		file_o5_messaging_v1_annotations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyConfig); i {
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
	file_o5_messaging_v1_annotations_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Config_Broadcast)(nil),
		(*Config_Unicast)(nil),
		(*Config_Request)(nil),
		(*Config_Reply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_messaging_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_o5_messaging_v1_annotations_proto_goTypes,
		DependencyIndexes: file_o5_messaging_v1_annotations_proto_depIdxs,
		MessageInfos:      file_o5_messaging_v1_annotations_proto_msgTypes,
		ExtensionInfos:    file_o5_messaging_v1_annotations_proto_extTypes,
	}.Build()
	File_o5_messaging_v1_annotations_proto = out.File
	file_o5_messaging_v1_annotations_proto_rawDesc = nil
	file_o5_messaging_v1_annotations_proto_goTypes = nil
	file_o5_messaging_v1_annotations_proto_depIdxs = nil
}
