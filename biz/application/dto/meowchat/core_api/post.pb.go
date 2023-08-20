// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: meowchat/core_api/post.proto

package core_api

import (
	basic "github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
	_ "github.com/xh-polaris/meowchat-core-api/biz/application/dto/http"
	user "github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
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

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	CreateAt   int64             `protobuf:"varint,2,opt,name=createAt,proto3" json:"createAt,omitempty" form:"createAt" query:"createAt"`
	UpdateAt   int64             `protobuf:"varint,3,opt,name=updateAt,proto3" json:"updateAt,omitempty" form:"updateAt" query:"updateAt"`
	Title      string            `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`
	Text       string            `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty" form:"text" query:"text"`
	CoverUrl   string            `protobuf:"bytes,7,opt,name=coverUrl,proto3" json:"coverUrl,omitempty" form:"coverUrl" query:"coverUrl"`
	Tags       []string          `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty" form:"tags" query:"tags"`
	User       *user.UserPreview `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty" form:"user" query:"user"`
	IsOfficial bool              `protobuf:"varint,10,opt,name=isOfficial,proto3" json:"isOfficial,omitempty" form:"isOfficial" query:"isOfficial"`
	Likes      int64             `protobuf:"varint,11,opt,name=likes,proto3" json:"likes,omitempty" form:"likes" query:"likes"`
	Comments   int64             `protobuf:"varint,12,opt,name=comments,proto3" json:"comments,omitempty" form:"comments" query:"comments"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Post) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Post) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Post) GetUser() *user.UserPreview {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Post) GetIsOfficial() bool {
	if x != nil {
		return x.IsOfficial
	}
	return false
}

func (x *Post) GetLikes() int64 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *Post) GetComments() int64 {
	if x != nil {
		return x.Comments
	}
	return 0
}

// 多选一
type SearchOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询全部字段
	Key *string `protobuf:"bytes,1,opt,name=key,proto3,oneof" json:"key,omitempty" form:"key" query:"key"`
	// 各自匹配对应字段
	Title *string `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty" form:"title" query:"title"`
	Text  *string `protobuf:"bytes,3,opt,name=text,proto3,oneof" json:"text,omitempty" form:"text" query:"text"`
	Tag   *string `protobuf:"bytes,4,opt,name=tag,proto3,oneof" json:"tag,omitempty" form:"tag" query:"tag"`
}

func (x *SearchOptions) Reset() {
	*x = SearchOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchOptions) ProtoMessage() {}

func (x *SearchOptions) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchOptions.ProtoReflect.Descriptor instead.
func (*SearchOptions) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{1}
}

func (x *SearchOptions) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *SearchOptions) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *SearchOptions) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *SearchOptions) GetTag() string {
	if x != nil && x.Tag != nil {
		return *x.Tag
	}
	return ""
}

type GetPostPreviewsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginationOption *basic.PaginationOptions `protobuf:"bytes,1,opt,name=paginationOption,proto3" json:"paginationOption,omitempty" form:"paginationOption" query:"paginationOption"`
	OnlyOfficial     *bool                    `protobuf:"varint,2,opt,name=onlyOfficial,proto3,oneof" json:"onlyOfficial,omitempty" form:"onlyOfficial" query:"onlyOfficial"`
	OnlyUserId       *string                  `protobuf:"bytes,3,opt,name=onlyUserId,proto3,oneof" json:"onlyUserId,omitempty" form:"onlyUserId" query:"onlyUserId"`
	SearchOptions    *SearchOptions           `protobuf:"bytes,4,opt,name=searchOptions,proto3,oneof" json:"searchOptions,omitempty" form:"searchOptions" query:"searchOptions"`
}

func (x *GetPostPreviewsReq) Reset() {
	*x = GetPostPreviewsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostPreviewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostPreviewsReq) ProtoMessage() {}

func (x *GetPostPreviewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostPreviewsReq.ProtoReflect.Descriptor instead.
func (*GetPostPreviewsReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostPreviewsReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

func (x *GetPostPreviewsReq) GetOnlyOfficial() bool {
	if x != nil && x.OnlyOfficial != nil {
		return *x.OnlyOfficial
	}
	return false
}

func (x *GetPostPreviewsReq) GetOnlyUserId() string {
	if x != nil && x.OnlyUserId != nil {
		return *x.OnlyUserId
	}
	return ""
}

func (x *GetPostPreviewsReq) GetSearchOptions() *SearchOptions {
	if x != nil {
		return x.SearchOptions
	}
	return nil
}

type GetPostPreviewsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty" form:"posts" query:"posts"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total" query:"total"`
	Token string  `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" form:"token" query:"token"`
}

func (x *GetPostPreviewsResp) Reset() {
	*x = GetPostPreviewsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostPreviewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostPreviewsResp) ProtoMessage() {}

func (x *GetPostPreviewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostPreviewsResp.ProtoReflect.Descriptor instead.
func (*GetPostPreviewsResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostPreviewsResp) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *GetPostPreviewsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetPostPreviewsResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetPostDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty" form:"postId" query:"postId"`
}

func (x *GetPostDetailReq) Reset() {
	*x = GetPostDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostDetailReq) ProtoMessage() {}

func (x *GetPostDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostDetailReq.ProtoReflect.Descriptor instead.
func (*GetPostDetailReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostDetailReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type GetPostDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty" form:"post" query:"post"`
}

func (x *GetPostDetailResp) Reset() {
	*x = GetPostDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostDetailResp) ProtoMessage() {}

func (x *GetPostDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostDetailResp.ProtoReflect.Descriptor instead.
func (*GetPostDetailResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{5}
}

func (x *GetPostDetailResp) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type NewPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *string  `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty" form:"id" query:"id"`
	Title    string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" form:"title" query:"title"`
	Text     string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty" form:"text" query:"text"`
	CoverUrl string   `protobuf:"bytes,4,opt,name=coverUrl,proto3" json:"coverUrl,omitempty" form:"coverUrl" query:"coverUrl"`
	Tags     []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty" form:"tags" query:"tags"`
}

func (x *NewPostReq) Reset() {
	*x = NewPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPostReq) ProtoMessage() {}

func (x *NewPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPostReq.ProtoReflect.Descriptor instead.
func (*NewPostReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{6}
}

func (x *NewPostReq) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *NewPostReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NewPostReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *NewPostReq) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *NewPostReq) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type NewPostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId       string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty" form:"postId" query:"postId"`
	GetFish      bool   `protobuf:"varint,2,opt,name=getFish,proto3" json:"getFish,omitempty" form:"getFish" query:"getFish"`
	GetFishTimes int64  `protobuf:"varint,3,opt,name=getFishTimes,proto3" json:"getFishTimes,omitempty" form:"getFishTimes" query:"getFishTimes"`
}

func (x *NewPostResp) Reset() {
	*x = NewPostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPostResp) ProtoMessage() {}

func (x *NewPostResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPostResp.ProtoReflect.Descriptor instead.
func (*NewPostResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{7}
}

func (x *NewPostResp) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *NewPostResp) GetGetFish() bool {
	if x != nil {
		return x.GetFish
	}
	return false
}

func (x *NewPostResp) GetGetFishTimes() int64 {
	if x != nil {
		return x.GetFishTimes
	}
	return 0
}

type DeletePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
}

func (x *DeletePostReq) Reset() {
	*x = DeletePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostReq) ProtoMessage() {}

func (x *DeletePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostReq.ProtoReflect.Descriptor instead.
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePostReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeletePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePostResp) Reset() {
	*x = DeletePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResp) ProtoMessage() {}

func (x *DeletePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResp.ProtoReflect.Descriptor instead.
func (*DeletePostResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{9}
}

type SetOfficialReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId,omitempty" form:"postId" query:"postId"`
	IsRemove *bool  `protobuf:"varint,2,opt,name=isRemove,proto3,oneof" json:"isRemove,omitempty" form:"isRemove" query:"isRemove"`
}

func (x *SetOfficialReq) Reset() {
	*x = SetOfficialReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOfficialReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOfficialReq) ProtoMessage() {}

func (x *SetOfficialReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOfficialReq.ProtoReflect.Descriptor instead.
func (*SetOfficialReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{10}
}

func (x *SetOfficialReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *SetOfficialReq) GetIsRemove() bool {
	if x != nil && x.IsRemove != nil {
		return *x.IsRemove
	}
	return false
}

type SetOfficialResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetOfficialResp) Reset() {
	*x = SetOfficialResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_post_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOfficialResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOfficialResp) ProtoMessage() {}

func (x *SetOfficialResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_post_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOfficialResp.ProtoReflect.Descriptor instead.
func (*SetOfficialResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_post_proto_rawDescGZIP(), []int{11}
}

var File_meowchat_core_api_post_proto protoreflect.FileDescriptor

var file_meowchat_core_api_post_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x15, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x03, 0x74, 0x61, 0x67, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x61, 0x67, 0x22, 0xa7, 0x02, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x0c, 0x6f, 0x6e, 0x6c, 0x79, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x4b, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x02,
	0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x69, 0x61, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x70, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x0b, 0x4e, 0x65, 0x77,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x65,
	0x74, 0x46, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x1f,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x56, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x69,
	0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x08, 0x69, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x69, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x42, 0x4f, 0x5a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x63,
	0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_post_proto_rawDescOnce sync.Once
	file_meowchat_core_api_post_proto_rawDescData = file_meowchat_core_api_post_proto_rawDesc
)

func file_meowchat_core_api_post_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_post_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_post_proto_rawDescData)
	})
	return file_meowchat_core_api_post_proto_rawDescData
}

var file_meowchat_core_api_post_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_meowchat_core_api_post_proto_goTypes = []interface{}{
	(*Post)(nil),                    // 0: meowchat.core_api.Post
	(*SearchOptions)(nil),           // 1: meowchat.core_api.SearchOptions
	(*GetPostPreviewsReq)(nil),      // 2: meowchat.core_api.GetPostPreviewsReq
	(*GetPostPreviewsResp)(nil),     // 3: meowchat.core_api.GetPostPreviewsResp
	(*GetPostDetailReq)(nil),        // 4: meowchat.core_api.GetPostDetailReq
	(*GetPostDetailResp)(nil),       // 5: meowchat.core_api.GetPostDetailResp
	(*NewPostReq)(nil),              // 6: meowchat.core_api.NewPostReq
	(*NewPostResp)(nil),             // 7: meowchat.core_api.NewPostResp
	(*DeletePostReq)(nil),           // 8: meowchat.core_api.DeletePostReq
	(*DeletePostResp)(nil),          // 9: meowchat.core_api.DeletePostResp
	(*SetOfficialReq)(nil),          // 10: meowchat.core_api.SetOfficialReq
	(*SetOfficialResp)(nil),         // 11: meowchat.core_api.SetOfficialResp
	(*user.UserPreview)(nil),        // 12: meowchat.user.UserPreview
	(*basic.PaginationOptions)(nil), // 13: basic.PaginationOptions
}
var file_meowchat_core_api_post_proto_depIdxs = []int32{
	12, // 0: meowchat.core_api.Post.user:type_name -> meowchat.user.UserPreview
	13, // 1: meowchat.core_api.GetPostPreviewsReq.paginationOption:type_name -> basic.PaginationOptions
	1,  // 2: meowchat.core_api.GetPostPreviewsReq.searchOptions:type_name -> meowchat.core_api.SearchOptions
	0,  // 3: meowchat.core_api.GetPostPreviewsResp.posts:type_name -> meowchat.core_api.Post
	0,  // 4: meowchat.core_api.GetPostDetailResp.post:type_name -> meowchat.core_api.Post
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func file_meowchat_core_api_post_proto_init() {
	if File_meowchat_core_api_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_meowchat_core_api_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchOptions); i {
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
		file_meowchat_core_api_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostPreviewsReq); i {
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
		file_meowchat_core_api_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostPreviewsResp); i {
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
		file_meowchat_core_api_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostDetailReq); i {
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
		file_meowchat_core_api_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostDetailResp); i {
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
		file_meowchat_core_api_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPostReq); i {
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
		file_meowchat_core_api_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPostResp); i {
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
		file_meowchat_core_api_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostReq); i {
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
		file_meowchat_core_api_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResp); i {
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
		file_meowchat_core_api_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOfficialReq); i {
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
		file_meowchat_core_api_post_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOfficialResp); i {
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
	file_meowchat_core_api_post_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_meowchat_core_api_post_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_meowchat_core_api_post_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_meowchat_core_api_post_proto_msgTypes[10].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_post_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_post_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_post_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_post_proto = out.File
	file_meowchat_core_api_post_proto_rawDesc = nil
	file_meowchat_core_api_post_proto_goTypes = nil
	file_meowchat_core_api_post_proto_depIdxs = nil
}
