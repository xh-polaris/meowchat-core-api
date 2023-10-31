// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: meowchat/core_api/comment.proto

package core_api

import (
	comment "github.com/xh-polaris/meowchat-core-api/biz/application/dto/platform/comment"
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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	FirstLevelId *string `protobuf:"bytes,2,opt,name=firstLevelId,proto3,oneof" json:"firstLevelId,omitempty" form:"firstLevelId" query:"firstLevelId"`
	Text         string  `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty" form:"text" query:"text"`
	User         *User   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty" form:"user" query:"user"`
	ReplyUser    *User   `protobuf:"bytes,5,opt,name=replyUser,proto3,oneof" json:"replyUser,omitempty" form:"replyUser" query:"replyUser"` // 这条评论回复的用户
	Comments     *int64  `protobuf:"varint,6,opt,name=comments,proto3,oneof" json:"comments,omitempty" form:"comments" query:"comments"`
	CreateAt     int64   `protobuf:"varint,7,opt,name=createAt,proto3" json:"createAt,omitempty" form:"createAt" query:"createAt"`
	LikeCount    *int64  `protobuf:"varint,8,opt,name=likeCount,proto3,oneof" json:"likeCount,omitempty" form:"likeCount" query:"likeCount"`
	IsLiked      *bool   `protobuf:"varint,9,opt,name=isLiked,proto3,oneof" json:"isLiked,omitempty" form:"isLiked" query:"isLiked"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetFirstLevelId() string {
	if x != nil && x.FirstLevelId != nil {
		return *x.FirstLevelId
	}
	return ""
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetReplyUser() *User {
	if x != nil {
		return x.ReplyUser
	}
	return nil
}

func (x *Comment) GetComments() int64 {
	if x != nil && x.Comments != nil {
		return *x.Comments
	}
	return 0
}

func (x *Comment) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Comment) GetLikeCount() int64 {
	if x != nil && x.LikeCount != nil {
		return *x.LikeCount
	}
	return 0
}

func (x *Comment) GetIsLiked() bool {
	if x != nil && x.IsLiked != nil {
		return *x.IsLiked
	}
	return false
}

type NewCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *string             `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty" form:"id" query:"id"`
	FirstLevelId *string             `protobuf:"bytes,2,opt,name=firstLevelId,proto3,oneof" json:"firstLevelId,omitempty" form:"firstLevelId" query:"firstLevelId"`
	Text         string              `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty" form:"text" query:"text"`
	Type         comment.CommentType `protobuf:"varint,4,opt,name=type,proto3,enum=platform.comment.CommentType" json:"type,omitempty" form:"type" query:"type"`
}

func (x *NewCommentReq) Reset() {
	*x = NewCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCommentReq) ProtoMessage() {}

func (x *NewCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCommentReq.ProtoReflect.Descriptor instead.
func (*NewCommentReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{1}
}

func (x *NewCommentReq) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *NewCommentReq) GetFirstLevelId() string {
	if x != nil && x.FirstLevelId != nil {
		return *x.FirstLevelId
	}
	return ""
}

func (x *NewCommentReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *NewCommentReq) GetType() comment.CommentType {
	if x != nil {
		return x.Type
	}
	return comment.CommentType(0)
}

type NewCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetFish      bool  `protobuf:"varint,1,opt,name=getFish,proto3" json:"getFish,omitempty" form:"getFish" query:"getFish"`
	GetFishTimes int64 `protobuf:"varint,2,opt,name=getFishTimes,proto3" json:"getFishTimes,omitempty" form:"getFishTimes" query:"getFishTimes"`
	GetFishNum   int64 `protobuf:"varint,3,opt,name=getFishNum,proto3" json:"getFishNum,omitempty" form:"getFishNum" query:"getFishNum"`
}

func (x *NewCommentResp) Reset() {
	*x = NewCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCommentResp) ProtoMessage() {}

func (x *NewCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCommentResp.ProtoReflect.Descriptor instead.
func (*NewCommentResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{2}
}

func (x *NewCommentResp) GetGetFish() bool {
	if x != nil {
		return x.GetFish
	}
	return false
}

func (x *NewCommentResp) GetGetFishTimes() int64 {
	if x != nil {
		return x.GetFishTimes
	}
	return 0
}

func (x *NewCommentResp) GetGetFishNum() int64 {
	if x != nil {
		return x.GetFishNum
	}
	return 0
}

type GetCommentsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Type comment.CommentType `protobuf:"varint,2,opt,name=type,proto3,enum=platform.comment.CommentType" json:"type,omitempty" form:"type" query:"type"`
	Page int64               `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty" form:"page" query:"page"`
}

func (x *GetCommentsReq) Reset() {
	*x = GetCommentsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsReq) ProtoMessage() {}

func (x *GetCommentsReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsReq.ProtoReflect.Descriptor instead.
func (*GetCommentsReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentsReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCommentsReq) GetType() comment.CommentType {
	if x != nil {
		return x.Type
	}
	return comment.CommentType(0)
}

func (x *GetCommentsReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetCommentsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty" form:"comments" query:"comments"`
	Total    int64      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total" query:"total"`
}

func (x *GetCommentsResp) Reset() {
	*x = GetCommentsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsResp) ProtoMessage() {}

func (x *GetCommentsResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsResp.ProtoReflect.Descriptor instead.
func (*GetCommentsResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentsResp) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *GetCommentsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=commentId,proto3" json:"commentId,omitempty" form:"commentId" query:"commentId"`
}

func (x *DeleteCommentReq) Reset() {
	*x = DeleteCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentReq) ProtoMessage() {}

func (x *DeleteCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentReq.ProtoReflect.Descriptor instead.
func (*DeleteCommentReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCommentReq) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

type DeleteCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentResp) Reset() {
	*x = DeleteCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResp) ProtoMessage() {}

func (x *DeleteCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResp.ProtoReflect.Descriptor instead.
func (*DeleteCommentResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_comment_proto_rawDescGZIP(), []int{6}
}

var File_meowchat_core_api_comment_proto protoreflect.FileDescriptor

var file_meowchat_core_api_comment_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x84, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x48, 0x01, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x21, 0x0a, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x69, 0x73, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x0d, 0x4e, 0x65, 0x77,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x27, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x31, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x74,
	0x46, 0x69, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x65, 0x74, 0x46,
	0x69, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x67, 0x65, 0x74, 0x46, 0x69,
	0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x46, 0x69,
	0x73, 0x68, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x74,
	0x46, 0x69, 0x73, 0x68, 0x4e, 0x75, 0x6d, 0x22, 0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x5f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x30, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x86, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d,
	0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65,
	0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x61,
	0x70, 0x69, 0x42, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68,
	0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_comment_proto_rawDescOnce sync.Once
	file_meowchat_core_api_comment_proto_rawDescData = file_meowchat_core_api_comment_proto_rawDesc
)

func file_meowchat_core_api_comment_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_comment_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_comment_proto_rawDescData)
	})
	return file_meowchat_core_api_comment_proto_rawDescData
}

var file_meowchat_core_api_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_meowchat_core_api_comment_proto_goTypes = []interface{}{
	(*Comment)(nil),           // 0: meowchat.core_api.Comment
	(*NewCommentReq)(nil),     // 1: meowchat.core_api.NewCommentReq
	(*NewCommentResp)(nil),    // 2: meowchat.core_api.NewCommentResp
	(*GetCommentsReq)(nil),    // 3: meowchat.core_api.GetCommentsReq
	(*GetCommentsResp)(nil),   // 4: meowchat.core_api.GetCommentsResp
	(*DeleteCommentReq)(nil),  // 5: meowchat.core_api.DeleteCommentReq
	(*DeleteCommentResp)(nil), // 6: meowchat.core_api.DeleteCommentResp
	(*User)(nil),              // 7: meowchat.core_api.User
	(comment.CommentType)(0),  // 8: platform.comment.CommentType
}
var file_meowchat_core_api_comment_proto_depIdxs = []int32{
	7, // 0: meowchat.core_api.Comment.user:type_name -> meowchat.core_api.User
	7, // 1: meowchat.core_api.Comment.replyUser:type_name -> meowchat.core_api.User
	8, // 2: meowchat.core_api.NewCommentReq.type:type_name -> platform.comment.CommentType
	8, // 3: meowchat.core_api.GetCommentsReq.type:type_name -> platform.comment.CommentType
	0, // 4: meowchat.core_api.GetCommentsResp.comments:type_name -> meowchat.core_api.Comment
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func file_meowchat_core_api_comment_proto_init() {
	if File_meowchat_core_api_comment_proto != nil {
		return
	}
	file_meowchat_core_api_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_meowchat_core_api_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCommentReq); i {
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
		file_meowchat_core_api_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCommentResp); i {
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
		file_meowchat_core_api_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsReq); i {
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
		file_meowchat_core_api_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsResp); i {
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
		file_meowchat_core_api_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentReq); i {
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
		file_meowchat_core_api_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentResp); i {
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
	file_meowchat_core_api_comment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_meowchat_core_api_comment_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_comment_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_comment_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_comment_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_comment_proto = out.File
	file_meowchat_core_api_comment_proto_rawDesc = nil
	file_meowchat_core_api_comment_proto_goTypes = nil
	file_meowchat_core_api_comment_proto_depIdxs = nil
}
