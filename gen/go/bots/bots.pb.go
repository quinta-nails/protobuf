// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: bots/bots.proto

package bots

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	StudioId  int64  `protobuf:"varint,4,opt,name=studio_id,json=studioId,proto3" json:"studio_id,omitempty"`
}

func (x *Bot) Reset() {
	*x = Bot{}
	mi := &file_bots_bots_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bot) ProtoMessage() {}

func (x *Bot) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bot.ProtoReflect.Descriptor instead.
func (*Bot) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{0}
}

func (x *Bot) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Bot) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Bot) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Bot) GetStudioId() int64 {
	if x != nil {
		return x.StudioId
	}
	return 0
}

// AuthByInitData --------------------------------------------------------------------
type AuthByInitDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InitData string `protobuf:"bytes,1,opt,name=init_data,json=initData,proto3" json:"init_data,omitempty"`
}

func (x *AuthByInitDataRequest) Reset() {
	*x = AuthByInitDataRequest{}
	mi := &file_bots_bots_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthByInitDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByInitDataRequest) ProtoMessage() {}

func (x *AuthByInitDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByInitDataRequest.ProtoReflect.Descriptor instead.
func (*AuthByInitDataRequest) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{1}
}

func (x *AuthByInitDataRequest) GetInitData() string {
	if x != nil {
		return x.InitData
	}
	return ""
}

type AuthByInitDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *AuthByInitDataResponse) Reset() {
	*x = AuthByInitDataResponse{}
	mi := &file_bots_bots_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthByInitDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthByInitDataResponse) ProtoMessage() {}

func (x *AuthByInitDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthByInitDataResponse.ProtoReflect.Descriptor instead.
func (*AuthByInitDataResponse) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{2}
}

func (x *AuthByInitDataResponse) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

// AddBot ----------------------------------------------------------------------------
type AddBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudioId int64  `protobuf:"varint,2,opt,name=studio_id,json=studioId,proto3" json:"studio_id,omitempty"`
}

func (x *AddBotRequest) Reset() {
	*x = AddBotRequest{}
	mi := &file_bots_bots_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBotRequest) ProtoMessage() {}

func (x *AddBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBotRequest.ProtoReflect.Descriptor instead.
func (*AddBotRequest) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{3}
}

func (x *AddBotRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AddBotRequest) GetStudioId() int64 {
	if x != nil {
		return x.StudioId
	}
	return 0
}

type AddBotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Bot `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddBotResponse) Reset() {
	*x = AddBotResponse{}
	mi := &file_bots_bots_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddBotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBotResponse) ProtoMessage() {}

func (x *AddBotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBotResponse.ProtoReflect.Descriptor instead.
func (*AddBotResponse) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{4}
}

func (x *AddBotResponse) GetResult() *Bot {
	if x != nil {
		return x.Result
	}
	return nil
}

// SetupBot ----------------------------------------------------------------------------
type SetupBotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SetupBotRequest) Reset() {
	*x = SetupBotRequest{}
	mi := &file_bots_bots_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetupBotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupBotRequest) ProtoMessage() {}

func (x *SetupBotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupBotRequest.ProtoReflect.Descriptor instead.
func (*SetupBotRequest) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{5}
}

func (x *SetupBotRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SetupBotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *SetupBotResponse) Reset() {
	*x = SetupBotResponse{}
	mi := &file_bots_bots_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetupBotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupBotResponse) ProtoMessage() {}

func (x *SetupBotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupBotResponse.ProtoReflect.Descriptor instead.
func (*SetupBotResponse) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{6}
}

func (x *SetupBotResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Reserve ----------------------------------------------------------------------------
type ReserveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotId     int64                  `protobuf:"varint,1,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	UserId    int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ServiceId int64                  `protobuf:"varint,3,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Datetime  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=datetime,proto3" json:"datetime,omitempty"`
}

func (x *ReserveRequest) Reset() {
	*x = ReserveRequest{}
	mi := &file_bots_bots_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReserveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveRequest) ProtoMessage() {}

func (x *ReserveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveRequest.ProtoReflect.Descriptor instead.
func (*ReserveRequest) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{7}
}

func (x *ReserveRequest) GetBotId() int64 {
	if x != nil {
		return x.BotId
	}
	return 0
}

func (x *ReserveRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReserveRequest) GetServiceId() int64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *ReserveRequest) GetDatetime() *timestamppb.Timestamp {
	if x != nil {
		return x.Datetime
	}
	return nil
}

type ReserveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReserveResponse) Reset() {
	*x = ReserveResponse{}
	mi := &file_bots_bots_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReserveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveResponse) ProtoMessage() {}

func (x *ReserveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bots_bots_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveResponse.ProtoReflect.Descriptor instead.
func (*ReserveResponse) Descriptor() ([]byte, []int) {
	return file_bots_bots_proto_rawDescGZIP(), []int{8}
}

var File_bots_bots_proto protoreflect.FileDescriptor

var file_bots_bots_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x74, 0x73, 0x2f, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x6f, 0x74, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x03, 0x42, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x42, 0x79, 0x49, 0x6e,
	0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a, 0x16, 0x41, 0x75,
	0x74, 0x68, 0x42, 0x79, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x22, 0x55, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x2e,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22,
	0x02, 0x20, 0x00, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x49, 0x64, 0x22, 0x33, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x42, 0x6f, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x2a, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x75, 0x70, 0x42, 0x6f, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24,
	0x0a, 0x10, 0x53, 0x65, 0x74, 0x75, 0x70, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x22, 0xbc, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba, 0x48, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xba,
	0x48, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x40, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x08, 0xba, 0x48, 0x05, 0xb2, 0x01, 0x02, 0x40, 0x01, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x85, 0x01, 0x0a, 0x13, 0x54, 0x65, 0x6c, 0x65, 0x67,
	0x72, 0x61, 0x6d, 0x42, 0x6f, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x06, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x74, 0x73, 0x2e,
	0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x62, 0x6f, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x75, 0x70, 0x42, 0x6f, 0x74, 0x12,
	0x15, 0x2e, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x42, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x74, 0x73, 0x2e, 0x53, 0x65,
	0x74, 0x75, 0x70, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6c,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x6f, 0x74, 0x73, 0x42, 0x09, 0x42, 0x6f, 0x74, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x75, 0x69, 0x6e, 0x74, 0x61, 0x2d, 0x6e, 0x61, 0x69, 0x6c, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f, 0x74, 0x73, 0xa2, 0x02,
	0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x42, 0x6f, 0x74, 0x73, 0xca, 0x02, 0x04, 0x42, 0x6f,
	0x74, 0x73, 0xe2, 0x02, 0x10, 0x42, 0x6f, 0x74, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x42, 0x6f, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bots_bots_proto_rawDescOnce sync.Once
	file_bots_bots_proto_rawDescData = file_bots_bots_proto_rawDesc
)

func file_bots_bots_proto_rawDescGZIP() []byte {
	file_bots_bots_proto_rawDescOnce.Do(func() {
		file_bots_bots_proto_rawDescData = protoimpl.X.CompressGZIP(file_bots_bots_proto_rawDescData)
	})
	return file_bots_bots_proto_rawDescData
}

var file_bots_bots_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bots_bots_proto_goTypes = []any{
	(*Bot)(nil),                    // 0: bots.Bot
	(*AuthByInitDataRequest)(nil),  // 1: bots.AuthByInitDataRequest
	(*AuthByInitDataResponse)(nil), // 2: bots.AuthByInitDataResponse
	(*AddBotRequest)(nil),          // 3: bots.AddBotRequest
	(*AddBotResponse)(nil),         // 4: bots.AddBotResponse
	(*SetupBotRequest)(nil),        // 5: bots.SetupBotRequest
	(*SetupBotResponse)(nil),       // 6: bots.SetupBotResponse
	(*ReserveRequest)(nil),         // 7: bots.ReserveRequest
	(*ReserveResponse)(nil),        // 8: bots.ReserveResponse
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_bots_bots_proto_depIdxs = []int32{
	0, // 0: bots.AddBotResponse.result:type_name -> bots.Bot
	9, // 1: bots.ReserveRequest.datetime:type_name -> google.protobuf.Timestamp
	3, // 2: bots.TelegramBotsService.AddBot:input_type -> bots.AddBotRequest
	5, // 3: bots.TelegramBotsService.SetupBot:input_type -> bots.SetupBotRequest
	4, // 4: bots.TelegramBotsService.AddBot:output_type -> bots.AddBotResponse
	6, // 5: bots.TelegramBotsService.SetupBot:output_type -> bots.SetupBotResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bots_bots_proto_init() }
func file_bots_bots_proto_init() {
	if File_bots_bots_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bots_bots_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bots_bots_proto_goTypes,
		DependencyIndexes: file_bots_bots_proto_depIdxs,
		MessageInfos:      file_bots_bots_proto_msgTypes,
	}.Build()
	File_bots_bots_proto = out.File
	file_bots_bots_proto_rawDesc = nil
	file_bots_bots_proto_goTypes = nil
	file_bots_bots_proto_depIdxs = nil
}
