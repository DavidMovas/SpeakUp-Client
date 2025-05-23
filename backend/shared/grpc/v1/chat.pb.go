// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: chat.proto

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type CreateChatRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*CreateChatRequest_PrivateChat_
	//	*CreateChatRequest_GroupChat_
	Payload       isCreateChatRequest_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	mi := &file_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *CreateChatRequest) GetPayload() isCreateChatRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CreateChatRequest) GetPrivateChat() *CreateChatRequest_PrivateChat {
	if x != nil {
		if x, ok := x.Payload.(*CreateChatRequest_PrivateChat_); ok {
			return x.PrivateChat
		}
	}
	return nil
}

func (x *CreateChatRequest) GetGroupChat() *CreateChatRequest_GroupChat {
	if x != nil {
		if x, ok := x.Payload.(*CreateChatRequest_GroupChat_); ok {
			return x.GroupChat
		}
	}
	return nil
}

type isCreateChatRequest_Payload interface {
	isCreateChatRequest_Payload()
}

type CreateChatRequest_PrivateChat_ struct {
	PrivateChat *CreateChatRequest_PrivateChat `protobuf:"bytes,1,opt,name=private_chat,json=privateChat,proto3,oneof"`
}

type CreateChatRequest_GroupChat_ struct {
	GroupChat *CreateChatRequest_GroupChat `protobuf:"bytes,2,opt,name=group_chat,json=groupChat,proto3,oneof"`
}

func (*CreateChatRequest_PrivateChat_) isCreateChatRequest_Payload() {}

func (*CreateChatRequest_GroupChat_) isCreateChatRequest_Payload() {}

type CreateChatResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateChatResponse) Reset() {
	*x = CreateChatResponse{}
	mi := &file_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResponse) ProtoMessage() {}

func (x *CreateChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatResponse.ProtoReflect.Descriptor instead.
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatResponse) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type ConnectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*ConnectRequest_JoinChat_
	//	*ConnectRequest_Message
	Payload       isConnectRequest_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	mi := &file_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectRequest) GetPayload() isConnectRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ConnectRequest) GetJoinChat() *ConnectRequest_JoinChat {
	if x != nil {
		if x, ok := x.Payload.(*ConnectRequest_JoinChat_); ok {
			return x.JoinChat
		}
	}
	return nil
}

func (x *ConnectRequest) GetMessage() *Message {
	if x != nil {
		if x, ok := x.Payload.(*ConnectRequest_Message); ok {
			return x.Message
		}
	}
	return nil
}

type isConnectRequest_Payload interface {
	isConnectRequest_Payload()
}

type ConnectRequest_JoinChat_ struct {
	JoinChat *ConnectRequest_JoinChat `protobuf:"bytes,1,opt,name=join_chat,json=joinChat,proto3,oneof"`
}

type ConnectRequest_Message struct {
	Message *Message `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

func (*ConnectRequest_JoinChat_) isConnectRequest_Payload() {}

func (*ConnectRequest_Message) isConnectRequest_Payload() {}

type ConnectResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*ConnectResponse_Message
	//	*ConnectResponse_MessageList
	Payload       isConnectResponse_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	mi := &file_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectResponse) GetPayload() isConnectResponse_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ConnectResponse) GetMessage() *Message {
	if x != nil {
		if x, ok := x.Payload.(*ConnectResponse_Message); ok {
			return x.Message
		}
	}
	return nil
}

func (x *ConnectResponse) GetMessageList() *MessageList {
	if x != nil {
		if x, ok := x.Payload.(*ConnectResponse_MessageList); ok {
			return x.MessageList
		}
	}
	return nil
}

type isConnectResponse_Payload interface {
	isConnectResponse_Payload()
}

type ConnectResponse_Message struct {
	Message *Message `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

type ConnectResponse_MessageList struct {
	MessageList *MessageList `protobuf:"bytes,2,opt,name=message_list,json=messageList,proto3,oneof"`
}

func (*ConnectResponse_Message) isConnectResponse_Payload() {}

func (*ConnectResponse_MessageList) isConnectResponse_Payload() {}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	SenderId      string                 `protobuf:"bytes,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type MessageList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageList) Reset() {
	*x = MessageList{}
	mi := &file_chat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageList) ProtoMessage() {}

func (x *MessageList) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageList.ProtoReflect.Descriptor instead.
func (*MessageList) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *MessageList) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GetChatHistoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LastReadAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_read_at,json=lastReadAt,proto3" json:"last_read_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChatHistoryRequest) Reset() {
	*x = GetChatHistoryRequest{}
	mi := &file_chat_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChatHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatHistoryRequest) ProtoMessage() {}

func (x *GetChatHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetChatHistoryRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *GetChatHistoryRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *GetChatHistoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetChatHistoryRequest) GetLastReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReadAt
	}
	return nil
}

type GetChatHistoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChatHistoryResponse) Reset() {
	*x = GetChatHistoryResponse{}
	mi := &file_chat_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChatHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatHistoryResponse) ProtoMessage() {}

func (x *GetChatHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetChatHistoryResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{7}
}

func (x *GetChatHistoryResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type CreateChatRequest_PrivateChat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InitiatorId   string                 `protobuf:"bytes,1,opt,name=initiator_id,json=initiatorId,proto3" json:"initiator_id,omitempty"`
	MemberId      string                 `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateChatRequest_PrivateChat) Reset() {
	*x = CreateChatRequest_PrivateChat{}
	mi := &file_chat_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChatRequest_PrivateChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest_PrivateChat) ProtoMessage() {}

func (x *CreateChatRequest_PrivateChat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest_PrivateChat.ProtoReflect.Descriptor instead.
func (*CreateChatRequest_PrivateChat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateChatRequest_PrivateChat) GetInitiatorId() string {
	if x != nil {
		return x.InitiatorId
	}
	return ""
}

func (x *CreateChatRequest_PrivateChat) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

type CreateChatRequest_GroupChat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InitiatorId   string                 `protobuf:"bytes,1,opt,name=initiator_id,json=initiatorId,proto3" json:"initiator_id,omitempty"`
	MembersIds    []string               `protobuf:"bytes,2,rep,name=members_ids,json=membersIds,proto3" json:"members_ids,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateChatRequest_GroupChat) Reset() {
	*x = CreateChatRequest_GroupChat{}
	mi := &file_chat_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateChatRequest_GroupChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest_GroupChat) ProtoMessage() {}

func (x *CreateChatRequest_GroupChat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest_GroupChat.ProtoReflect.Descriptor instead.
func (*CreateChatRequest_GroupChat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CreateChatRequest_GroupChat) GetInitiatorId() string {
	if x != nil {
		return x.InitiatorId
	}
	return ""
}

func (x *CreateChatRequest_GroupChat) GetMembersIds() []string {
	if x != nil {
		return x.MembersIds
	}
	return nil
}

func (x *CreateChatRequest_GroupChat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ConnectRequest_JoinChat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChatId        string                 `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LastReadAt    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_read_at,json=lastReadAt,proto3" json:"last_read_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequest_JoinChat) Reset() {
	*x = ConnectRequest_JoinChat{}
	mi := &file_chat_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest_JoinChat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest_JoinChat) ProtoMessage() {}

func (x *ConnectRequest_JoinChat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest_JoinChat.ProtoReflect.Descriptor instead.
func (*ConnectRequest_JoinChat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ConnectRequest_JoinChat) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ConnectRequest_JoinChat) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConnectRequest_JoinChat) GetLastReadAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastReadAt
	}
	return nil
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0c,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43,
	0x68, 0x61, 0x74, 0x48, 0x00, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74,
	0x1a, 0x4d, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x1a,
	0x63, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x49, 0x64, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x86,
	0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x43, 0x68,
	0x61, 0x74, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x7a, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c,
	0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x94, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3b, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x61, 0x64, 0x41, 0x74, 0x22, 0x46, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xe9, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x51,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData []byte
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_proto_rawDesc), len(file_chat_proto_rawDesc)))
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chat_proto_goTypes = []any{
	(*CreateChatRequest)(nil),             // 0: chat.v1.CreateChatRequest
	(*CreateChatResponse)(nil),            // 1: chat.v1.CreateChatResponse
	(*ConnectRequest)(nil),                // 2: chat.v1.ConnectRequest
	(*ConnectResponse)(nil),               // 3: chat.v1.ConnectResponse
	(*Message)(nil),                       // 4: chat.v1.Message
	(*MessageList)(nil),                   // 5: chat.v1.MessageList
	(*GetChatHistoryRequest)(nil),         // 6: chat.v1.GetChatHistoryRequest
	(*GetChatHistoryResponse)(nil),        // 7: chat.v1.GetChatHistoryResponse
	(*CreateChatRequest_PrivateChat)(nil), // 8: chat.v1.CreateChatRequest.PrivateChat
	(*CreateChatRequest_GroupChat)(nil),   // 9: chat.v1.CreateChatRequest.GroupChat
	(*ConnectRequest_JoinChat)(nil),       // 10: chat.v1.ConnectRequest.JoinChat
	(*timestamppb.Timestamp)(nil),         // 11: google.protobuf.Timestamp
}
var file_chat_proto_depIdxs = []int32{
	8,  // 0: chat.v1.CreateChatRequest.private_chat:type_name -> chat.v1.CreateChatRequest.PrivateChat
	9,  // 1: chat.v1.CreateChatRequest.group_chat:type_name -> chat.v1.CreateChatRequest.GroupChat
	10, // 2: chat.v1.ConnectRequest.join_chat:type_name -> chat.v1.ConnectRequest.JoinChat
	4,  // 3: chat.v1.ConnectRequest.message:type_name -> chat.v1.Message
	4,  // 4: chat.v1.ConnectResponse.message:type_name -> chat.v1.Message
	5,  // 5: chat.v1.ConnectResponse.message_list:type_name -> chat.v1.MessageList
	11, // 6: chat.v1.Message.created_at:type_name -> google.protobuf.Timestamp
	4,  // 7: chat.v1.MessageList.messages:type_name -> chat.v1.Message
	11, // 8: chat.v1.GetChatHistoryRequest.last_read_at:type_name -> google.protobuf.Timestamp
	4,  // 9: chat.v1.GetChatHistoryResponse.messages:type_name -> chat.v1.Message
	11, // 10: chat.v1.ConnectRequest.JoinChat.last_read_at:type_name -> google.protobuf.Timestamp
	0,  // 11: chat.v1.ChatService.CreateChat:input_type -> chat.v1.CreateChatRequest
	2,  // 12: chat.v1.ChatService.connect:input_type -> chat.v1.ConnectRequest
	6,  // 13: chat.v1.ChatService.GetChatHistory:input_type -> chat.v1.GetChatHistoryRequest
	1,  // 14: chat.v1.ChatService.CreateChat:output_type -> chat.v1.CreateChatResponse
	3,  // 15: chat.v1.ChatService.connect:output_type -> chat.v1.ConnectResponse
	7,  // 16: chat.v1.ChatService.GetChatHistory:output_type -> chat.v1.GetChatHistoryResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	file_chat_proto_msgTypes[0].OneofWrappers = []any{
		(*CreateChatRequest_PrivateChat_)(nil),
		(*CreateChatRequest_GroupChat_)(nil),
	}
	file_chat_proto_msgTypes[2].OneofWrappers = []any{
		(*ConnectRequest_JoinChat_)(nil),
		(*ConnectRequest_Message)(nil),
	}
	file_chat_proto_msgTypes[3].OneofWrappers = []any{
		(*ConnectResponse_Message)(nil),
		(*ConnectResponse_MessageList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_proto_rawDesc), len(file_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
