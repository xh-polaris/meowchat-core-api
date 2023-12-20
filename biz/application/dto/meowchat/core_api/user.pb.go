// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: meowchat/core_api/user.proto

package core_api

import (
	basic "github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	_ "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
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

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *string `protobuf:"bytes,1,opt,name=userId,proto3,oneof" json:"userId" form:"userId" query:"userId"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserInfoReq) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user" form:"user" query:"user"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserInfoResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarUrl *string `protobuf:"bytes,1,opt,name=avatarUrl,proto3,oneof" json:"avatarUrl" form:"avatarUrl" query:"avatarUrl"`
	Nickname  *string `protobuf:"bytes,2,opt,name=nickname,proto3,oneof" json:"nickname" form:"nickname" query:"nickname"`
	Motto     *string `protobuf:"bytes,3,opt,name=motto,proto3,oneof" json:"motto" form:"motto" query:"motto"`
}

func (x *UpdateUserInfoReq) Reset() {
	*x = UpdateUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoReq) ProtoMessage() {}

func (x *UpdateUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserInfoReq) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

func (x *UpdateUserInfoReq) GetNickname() string {
	if x != nil && x.Nickname != nil {
		return *x.Nickname
	}
	return ""
}

func (x *UpdateUserInfoReq) GetMotto() string {
	if x != nil && x.Motto != nil {
		return *x.Motto
	}
	return ""
}

type UpdateUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserInfoResp) Reset() {
	*x = UpdateUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoResp) ProtoMessage() {}

func (x *UpdateUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoResp.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{3}
}

type SearchUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword          string                   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword" form:"keyword" query:"keyword"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,2,opt,name=paginationOption,proto3" json:"paginationOption" form:"paginationOption" query:"paginationOption"`
}

func (x *SearchUserReq) Reset() {
	*x = SearchUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserReq) ProtoMessage() {}

func (x *SearchUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserReq.ProtoReflect.Descriptor instead.
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{4}
}

func (x *SearchUserReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchUserReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

type SearchUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users" form:"users" query:"users"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total" form:"total" query:"total"`
	Token string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token" form:"token" query:"token"`
}

func (x *SearchUserResp) Reset() {
	*x = SearchUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResp) ProtoMessage() {}

func (x *SearchUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResp.ProtoReflect.Descriptor instead.
func (*SearchUserResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_user_proto_rawDescGZIP(), []int{5}
}

func (x *SearchUserResp) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SearchUserResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchUserResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_meowchat_core_api_user_proto protoreflect.FileDescriptor

var file_meowchat_core_api_user_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x3e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x97, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x6f, 0x74,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x74, 0x74,
	0x6f, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55,
	0x72, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x6f, 0x74, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x6f, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x6b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x83, 0x01,
	0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e,
	0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x61, 0x70, 0x69, 0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69,
	0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74,
	0x6f, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_user_proto_rawDescOnce sync.Once
	file_meowchat_core_api_user_proto_rawDescData = file_meowchat_core_api_user_proto_rawDesc
)

func file_meowchat_core_api_user_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_user_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_user_proto_rawDescData)
	})
	return file_meowchat_core_api_user_proto_rawDescData
}

var file_meowchat_core_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_meowchat_core_api_user_proto_goTypes = []interface{}{
	(*GetUserInfoReq)(nil),          // 0: meowchat.core_api.GetUserInfoReq
	(*GetUserInfoResp)(nil),         // 1: meowchat.core_api.GetUserInfoResp
	(*UpdateUserInfoReq)(nil),       // 2: meowchat.core_api.UpdateUserInfoReq
	(*UpdateUserInfoResp)(nil),      // 3: meowchat.core_api.UpdateUserInfoResp
	(*SearchUserReq)(nil),           // 4: meowchat.core_api.SearchUserReq
	(*SearchUserResp)(nil),          // 5: meowchat.core_api.SearchUserResp
	(*User)(nil),                    // 6: meowchat.core_api.User
	(*basic.PaginationOptions)(nil), // 7: basic.PaginationOptions
}
var file_meowchat_core_api_user_proto_depIdxs = []int32{
	6, // 0: meowchat.core_api.GetUserInfoResp.user:type_name -> meowchat.core_api.User
	7, // 1: meowchat.core_api.SearchUserReq.paginationOption:type_name -> basic.PaginationOptions
	6, // 2: meowchat.core_api.SearchUserResp.users:type_name -> meowchat.core_api.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}


func file_meowchat_core_api_user_proto_init() {
	if File_meowchat_core_api_user_proto != nil {
		return
	}
	file_meowchat_core_api_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
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
		file_meowchat_core_api_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoResp); i {
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
		file_meowchat_core_api_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserReq); i {
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
		file_meowchat_core_api_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserResp); i {
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
	file_meowchat_core_api_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_meowchat_core_api_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_user_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_user_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_user_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_user_proto = out.File
	file_meowchat_core_api_user_proto_rawDesc = nil
	file_meowchat_core_api_user_proto_goTypes = nil
	file_meowchat_core_api_user_proto_depIdxs = nil
}
