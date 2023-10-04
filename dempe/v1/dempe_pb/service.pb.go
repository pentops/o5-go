// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/dempe/v1/service.proto

package dempe_pb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	auth_pb "github.com/pentops/o5-go/auth/v1/auth_pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMessagesRequest) Reset() {
	*x = ListMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesRequest) ProtoMessage() {}

func (x *ListMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{0}
}

type ListMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*CapturedMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"` // TODO: Paging
}

func (x *ListMessagesResponse) Reset() {
	*x = ListMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesResponse) ProtoMessage() {}

func (x *ListMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListMessagesResponse) GetMessages() []*CapturedMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// when not set, returns the next unhandled message
	MessageId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequest) GetMessageId() *wrapperspb.StringValue {
	if x != nil {
		return x.MessageId
	}
	return nil
}

type CapturedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cause *DeadMessage `protobuf:"bytes,1,opt,name=cause,proto3" json:"cause,omitempty"`
}

func (x *CapturedMessage) Reset() {
	*x = CapturedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapturedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapturedMessage) ProtoMessage() {}

func (x *CapturedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapturedMessage.ProtoReflect.Descriptor instead.
func (*CapturedMessage) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *CapturedMessage) GetCause() *DeadMessage {
	if x != nil {
		return x.Cause
	}
	return nil
}

type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *CapturedMessage                       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Actions []*GetMessageResponse_GetMessageAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessageResponse) GetMessage() *CapturedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *GetMessageResponse) GetActions() []*GetMessageResponse_GetMessageAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

type MessageAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Note string `protobuf:"bytes,2,opt,name=note,proto3" json:"note,omitempty"`
	// Types that are assignable to Action:
	//
	//	*MessageAction_Delete
	//	*MessageAction_Requeue
	//	*MessageAction_Edit
	Action isMessageAction_Action `protobuf_oneof:"action"`
}

func (x *MessageAction) Reset() {
	*x = MessageAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAction) ProtoMessage() {}

func (x *MessageAction) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAction.ProtoReflect.Descriptor instead.
func (*MessageAction) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *MessageAction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageAction) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (m *MessageAction) GetAction() isMessageAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *MessageAction) GetDelete() *ActionDelete {
	if x, ok := x.GetAction().(*MessageAction_Delete); ok {
		return x.Delete
	}
	return nil
}

func (x *MessageAction) GetRequeue() *ActionRequeue {
	if x, ok := x.GetAction().(*MessageAction_Requeue); ok {
		return x.Requeue
	}
	return nil
}

func (x *MessageAction) GetEdit() *ActionEdit {
	if x, ok := x.GetAction().(*MessageAction_Edit); ok {
		return x.Edit
	}
	return nil
}

type isMessageAction_Action interface {
	isMessageAction_Action()
}

type MessageAction_Delete struct {
	// optional
	Delete *ActionDelete `protobuf:"bytes,10,opt,name=delete,proto3,oneof"`
}

type MessageAction_Requeue struct {
	Requeue *ActionRequeue `protobuf:"bytes,11,opt,name=requeue,proto3,oneof"`
}

type MessageAction_Edit struct {
	Edit *ActionEdit `protobuf:"bytes,12,opt,name=edit,proto3,oneof"` // Will return an error if more than one id is provided
}

func (*MessageAction_Delete) isMessageAction_Action() {}

func (*MessageAction_Requeue) isMessageAction_Action() {}

func (*MessageAction_Edit) isMessageAction_Action() {}

type MessagesActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageIds []string       `protobuf:"bytes,1,rep,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
	Action     *MessageAction `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *MessagesActionRequest) Reset() {
	*x = MessagesActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagesActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagesActionRequest) ProtoMessage() {}

func (x *MessagesActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagesActionRequest.ProtoReflect.Descriptor instead.
func (*MessagesActionRequest) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *MessagesActionRequest) GetMessageIds() []string {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

func (x *MessagesActionRequest) GetAction() *MessageAction {
	if x != nil {
		return x.Action
	}
	return nil
}

type ActionDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActionDelete) Reset() {
	*x = ActionDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionDelete) ProtoMessage() {}

func (x *ActionDelete) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionDelete.ProtoReflect.Descriptor instead.
func (*ActionDelete) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{7}
}

type ActionRequeue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActionRequeue) Reset() {
	*x = ActionRequeue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequeue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequeue) ProtoMessage() {}

func (x *ActionRequeue) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequeue.ProtoReflect.Descriptor instead.
func (*ActionRequeue) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{8}
}

type ActionEdit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewMessageJson []byte `protobuf:"bytes,1,opt,name=new_message_json,json=newMessageJson,proto3" json:"new_message_json,omitempty"`
}

func (x *ActionEdit) Reset() {
	*x = ActionEdit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionEdit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionEdit) ProtoMessage() {}

func (x *ActionEdit) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionEdit.ProtoReflect.Descriptor instead.
func (*ActionEdit) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *ActionEdit) GetNewMessageJson() []byte {
	if x != nil {
		return x.NewMessageJson
	}
	return nil
}

type MessagesActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessagesActionResponse) Reset() {
	*x = MessagesActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagesActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagesActionResponse) ProtoMessage() {}

func (x *MessagesActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagesActionResponse.ProtoReflect.Descriptor instead.
func (*MessagesActionResponse) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{10}
}

type GetMessageResponse_GetMessageAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    *MessageAction         `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Actor     *auth_pb.Actor         `protobuf:"bytes,2,opt,name=actor,proto3" json:"actor,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GetMessageResponse_GetMessageAction) Reset() {
	*x = GetMessageResponse_GetMessageAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dempe_v1_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse_GetMessageAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse_GetMessageAction) ProtoMessage() {}

func (x *GetMessageResponse_GetMessageAction) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dempe_v1_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse_GetMessageAction.ProtoReflect.Descriptor instead.
func (*GetMessageResponse_GetMessageAction) Descriptor() ([]byte, []int) {
	return file_o5_dempe_v1_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetMessageResponse_GetMessageAction) GetAction() *MessageAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *GetMessageResponse_GetMessageAction) GetActor() *auth_pb.Actor {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *GetMessageResponse_GetMessageAction) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_o5_dempe_v1_service_proto protoreflect.FileDescriptor

var file_o5_dempe_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x35, 0x2e,
	0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6f, 0x35, 0x2f, 0x64, 0x65, 0x6d,
	0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6f, 0x35, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x22, 0x41, 0x0a, 0x0f, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x63,
	0x61, 0x75, 0x73, 0x65, 0x22, 0xc4, 0x02, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f,
	0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4a, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0xa9, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x35, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xe3, 0x01, 0x0a, 0x0d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x35,
	0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x36, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x65, 0x64, 0x69, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x65, 0x64, 0x69, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x76, 0x0a, 0x15, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0e, 0x0a, 0x0c, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x65, 0x77, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x73,
	0x6f, 0x6e, 0x22, 0x18, 0x0a, 0x16, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf7, 0x02, 0x0a,
	0x0c, 0x44, 0x65, 0x6d, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x20, 0x2e,
	0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x64, 0x65, 0x6d,
	0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x75,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x6f,
	0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f,
	0x35, 0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7f, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x65, 0x6d,
	0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f, 0x35,
	0x2e, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x64, 0x65,
	0x6d, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d,
	0x67, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6d, 0x70,
	0x65, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_dempe_v1_service_proto_rawDescOnce sync.Once
	file_o5_dempe_v1_service_proto_rawDescData = file_o5_dempe_v1_service_proto_rawDesc
)

func file_o5_dempe_v1_service_proto_rawDescGZIP() []byte {
	file_o5_dempe_v1_service_proto_rawDescOnce.Do(func() {
		file_o5_dempe_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_dempe_v1_service_proto_rawDescData)
	})
	return file_o5_dempe_v1_service_proto_rawDescData
}

var file_o5_dempe_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_o5_dempe_v1_service_proto_goTypes = []interface{}{
	(*ListMessagesRequest)(nil),                 // 0: o5.dempe.v1.ListMessagesRequest
	(*ListMessagesResponse)(nil),                // 1: o5.dempe.v1.ListMessagesResponse
	(*GetMessageRequest)(nil),                   // 2: o5.dempe.v1.GetMessageRequest
	(*CapturedMessage)(nil),                     // 3: o5.dempe.v1.CapturedMessage
	(*GetMessageResponse)(nil),                  // 4: o5.dempe.v1.GetMessageResponse
	(*MessageAction)(nil),                       // 5: o5.dempe.v1.MessageAction
	(*MessagesActionRequest)(nil),               // 6: o5.dempe.v1.MessagesActionRequest
	(*ActionDelete)(nil),                        // 7: o5.dempe.v1.ActionDelete
	(*ActionRequeue)(nil),                       // 8: o5.dempe.v1.ActionRequeue
	(*ActionEdit)(nil),                          // 9: o5.dempe.v1.ActionEdit
	(*MessagesActionResponse)(nil),              // 10: o5.dempe.v1.MessagesActionResponse
	(*GetMessageResponse_GetMessageAction)(nil), // 11: o5.dempe.v1.GetMessageResponse.GetMessageAction
	(*wrapperspb.StringValue)(nil),              // 12: google.protobuf.StringValue
	(*DeadMessage)(nil),                         // 13: o5.dempe.v1.DeadMessage
	(*auth_pb.Actor)(nil),                       // 14: o5.auth.v1.Actor
	(*timestamppb.Timestamp)(nil),               // 15: google.protobuf.Timestamp
}
var file_o5_dempe_v1_service_proto_depIdxs = []int32{
	3,  // 0: o5.dempe.v1.ListMessagesResponse.messages:type_name -> o5.dempe.v1.CapturedMessage
	12, // 1: o5.dempe.v1.GetMessageRequest.message_id:type_name -> google.protobuf.StringValue
	13, // 2: o5.dempe.v1.CapturedMessage.cause:type_name -> o5.dempe.v1.DeadMessage
	3,  // 3: o5.dempe.v1.GetMessageResponse.message:type_name -> o5.dempe.v1.CapturedMessage
	11, // 4: o5.dempe.v1.GetMessageResponse.actions:type_name -> o5.dempe.v1.GetMessageResponse.GetMessageAction
	7,  // 5: o5.dempe.v1.MessageAction.delete:type_name -> o5.dempe.v1.ActionDelete
	8,  // 6: o5.dempe.v1.MessageAction.requeue:type_name -> o5.dempe.v1.ActionRequeue
	9,  // 7: o5.dempe.v1.MessageAction.edit:type_name -> o5.dempe.v1.ActionEdit
	5,  // 8: o5.dempe.v1.MessagesActionRequest.action:type_name -> o5.dempe.v1.MessageAction
	5,  // 9: o5.dempe.v1.GetMessageResponse.GetMessageAction.action:type_name -> o5.dempe.v1.MessageAction
	14, // 10: o5.dempe.v1.GetMessageResponse.GetMessageAction.actor:type_name -> o5.auth.v1.Actor
	15, // 11: o5.dempe.v1.GetMessageResponse.GetMessageAction.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 12: o5.dempe.v1.DempeService.ListMessages:input_type -> o5.dempe.v1.ListMessagesRequest
	2,  // 13: o5.dempe.v1.DempeService.GetMessage:input_type -> o5.dempe.v1.GetMessageRequest
	6,  // 14: o5.dempe.v1.DempeService.MessagesAction:input_type -> o5.dempe.v1.MessagesActionRequest
	1,  // 15: o5.dempe.v1.DempeService.ListMessages:output_type -> o5.dempe.v1.ListMessagesResponse
	4,  // 16: o5.dempe.v1.DempeService.GetMessage:output_type -> o5.dempe.v1.GetMessageResponse
	10, // 17: o5.dempe.v1.DempeService.MessagesAction:output_type -> o5.dempe.v1.MessagesActionResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_o5_dempe_v1_service_proto_init() }
func file_o5_dempe_v1_service_proto_init() {
	if File_o5_dempe_v1_service_proto != nil {
		return
	}
	file_o5_dempe_v1_dead_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_o5_dempe_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesRequest); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesResponse); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapturedMessage); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAction); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagesActionRequest); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionDelete); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequeue); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionEdit); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagesActionResponse); i {
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
		file_o5_dempe_v1_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse_GetMessageAction); i {
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
	file_o5_dempe_v1_service_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*MessageAction_Delete)(nil),
		(*MessageAction_Requeue)(nil),
		(*MessageAction_Edit)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_dempe_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_o5_dempe_v1_service_proto_goTypes,
		DependencyIndexes: file_o5_dempe_v1_service_proto_depIdxs,
		MessageInfos:      file_o5_dempe_v1_service_proto_msgTypes,
	}.Build()
	File_o5_dempe_v1_service_proto = out.File
	file_o5_dempe_v1_service_proto_rawDesc = nil
	file_o5_dempe_v1_service_proto_goTypes = nil
	file_o5_dempe_v1_service_proto_depIdxs = nil
}
