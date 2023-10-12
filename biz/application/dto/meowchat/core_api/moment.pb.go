// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: meowchat/core_api/moment.proto

package core_api

import (
	basic "github.com/xh-polaris/meowchat-core-api/biz/application/dto/basic"
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

type Moment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	CreateAt    int64             `protobuf:"varint,2,opt,name=createAt,proto3" json:"createAt" form:"createAt" query:"createAt"`
	CatId       string            `protobuf:"bytes,3,opt,name=catId,proto3" json:"catId" form:"catId" query:"catId"`
	Photos      []string          `protobuf:"bytes,4,rep,name=photos,proto3" json:"photos" form:"photos" query:"photos"` // 图片url
	Title       string            `protobuf:"bytes,5,opt,name=title,proto3" json:"title" form:"title" query:"title"`
	Text        string            `protobuf:"bytes,6,opt,name=text,proto3" json:"text" form:"text" query:"text"`
	CommunityId string            `protobuf:"bytes,7,opt,name=communityId,proto3" json:"communityId" form:"communityId" query:"communityId"`
	User        *user.UserPreview `protobuf:"bytes,8,opt,name=user,proto3" json:"user" form:"user" query:"user"`
}

func (x *Moment) Reset() {
	*x = Moment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Moment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moment) ProtoMessage() {}

func (x *Moment) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moment.ProtoReflect.Descriptor instead.
func (*Moment) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{0}
}

func (x *Moment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Moment) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Moment) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *Moment) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *Moment) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Moment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Moment) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *Moment) GetUser() *user.UserPreview {
	if x != nil {
		return x.User
	}
	return nil
}

type GetMomentPreviewsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId      *string                  `protobuf:"bytes,1,opt,name=communityId,proto3,oneof" json:"communityId" form:"communityId" query:"communityId"`
	IsParent         *bool                    `protobuf:"varint,2,opt,name=isParent,proto3,oneof" json:"isParent" form:"isParent" query:"isParent"`
	OnlyUserId       *string                  `protobuf:"bytes,3,opt,name=onlyUserId,proto3,oneof" json:"onlyUserId" form:"onlyUserId" query:"onlyUserId"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,4,opt,name=paginationOption,proto3" json:"paginationOption" form:"paginationOption" query:"paginationOption"`
}

func (x *GetMomentPreviewsReq) Reset() {
	*x = GetMomentPreviewsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMomentPreviewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMomentPreviewsReq) ProtoMessage() {}

func (x *GetMomentPreviewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMomentPreviewsReq.ProtoReflect.Descriptor instead.
func (*GetMomentPreviewsReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{1}
}

func (x *GetMomentPreviewsReq) GetCommunityId() string {
	if x != nil && x.CommunityId != nil {
		return *x.CommunityId
	}
	return ""
}

func (x *GetMomentPreviewsReq) GetIsParent() bool {
	if x != nil && x.IsParent != nil {
		return *x.IsParent
	}
	return false
}

func (x *GetMomentPreviewsReq) GetOnlyUserId() string {
	if x != nil && x.OnlyUserId != nil {
		return *x.OnlyUserId
	}
	return ""
}

func (x *GetMomentPreviewsReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

type GetMomentPreviewsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moments []*Moment `protobuf:"bytes,1,rep,name=moments,proto3" json:"moments" form:"moments" query:"moments"`
	Total   int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total" form:"total" query:"total"`
	Token   string    `protobuf:"bytes,3,opt,name=token,proto3" json:"token" form:"token" query:"token"`
}

func (x *GetMomentPreviewsResp) Reset() {
	*x = GetMomentPreviewsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMomentPreviewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMomentPreviewsResp) ProtoMessage() {}

func (x *GetMomentPreviewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMomentPreviewsResp.ProtoReflect.Descriptor instead.
func (*GetMomentPreviewsResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{2}
}

func (x *GetMomentPreviewsResp) GetMoments() []*Moment {
	if x != nil {
		return x.Moments
	}
	return nil
}

func (x *GetMomentPreviewsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetMomentPreviewsResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetMomentDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MomentId string `protobuf:"bytes,1,opt,name=momentId,proto3" json:"momentId" form:"momentId" query:"momentId"`
}

func (x *GetMomentDetailReq) Reset() {
	*x = GetMomentDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMomentDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMomentDetailReq) ProtoMessage() {}

func (x *GetMomentDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMomentDetailReq.ProtoReflect.Descriptor instead.
func (*GetMomentDetailReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{3}
}

func (x *GetMomentDetailReq) GetMomentId() string {
	if x != nil {
		return x.MomentId
	}
	return ""
}

type GetMomentDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moment *Moment `protobuf:"bytes,1,opt,name=moment,proto3" json:"moment" form:"moment" query:"moment"`
}

func (x *GetMomentDetailResp) Reset() {
	*x = GetMomentDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMomentDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMomentDetailResp) ProtoMessage() {}

func (x *GetMomentDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMomentDetailResp.ProtoReflect.Descriptor instead.
func (*GetMomentDetailResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{4}
}

func (x *GetMomentDetailResp) GetMoment() *Moment {
	if x != nil {
		return x.Moment
	}
	return nil
}

type DeleteMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MomentId string `protobuf:"bytes,1,opt,name=momentId,proto3" json:"momentId" form:"momentId" query:"momentId"`
}

func (x *DeleteMomentReq) Reset() {
	*x = DeleteMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMomentReq) ProtoMessage() {}

func (x *DeleteMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMomentReq.ProtoReflect.Descriptor instead.
func (*DeleteMomentReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteMomentReq) GetMomentId() string {
	if x != nil {
		return x.MomentId
	}
	return ""
}

type DeleteMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteMomentResp) Reset() {
	*x = DeleteMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMomentResp) ProtoMessage() {}

func (x *DeleteMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMomentResp.ProtoReflect.Descriptor instead.
func (*DeleteMomentResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{6}
}

type NewMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *string  `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id" form:"id" query:"id"`
	Title       *string  `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title" form:"title" query:"title"`
	CatId       *string  `protobuf:"bytes,3,opt,name=catId,proto3,oneof" json:"catId" form:"catId" query:"catId"`
	Text        *string  `protobuf:"bytes,4,opt,name=text,proto3,oneof" json:"text" form:"text" query:"text"`
	Photos      []string `protobuf:"bytes,5,rep,name=photos,proto3" json:"photos" form:"photos" query:"photos"`
	CommunityId *string  `protobuf:"bytes,6,opt,name=communityId,proto3,oneof" json:"communityId" form:"communityId" query:"communityId"`
}

func (x *NewMomentReq) Reset() {
	*x = NewMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMomentReq) ProtoMessage() {}

func (x *NewMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMomentReq.ProtoReflect.Descriptor instead.
func (*NewMomentReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{7}
}

func (x *NewMomentReq) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *NewMomentReq) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *NewMomentReq) GetCatId() string {
	if x != nil && x.CatId != nil {
		return *x.CatId
	}
	return ""
}

func (x *NewMomentReq) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *NewMomentReq) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *NewMomentReq) GetCommunityId() string {
	if x != nil && x.CommunityId != nil {
		return *x.CommunityId
	}
	return ""
}

type NewMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MomentId     string `protobuf:"bytes,1,opt,name=momentId,proto3" json:"momentId" form:"momentId" query:"momentId"`
	GetFish      bool   `protobuf:"varint,2,opt,name=getFish,proto3" json:"getFish" form:"getFish" query:"getFish"`
	GetFishTimes int64  `protobuf:"varint,3,opt,name=getFishTimes,proto3" json:"getFishTimes" form:"getFishTimes" query:"getFishTimes"`
}

func (x *NewMomentResp) Reset() {
	*x = NewMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewMomentResp) ProtoMessage() {}

func (x *NewMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewMomentResp.ProtoReflect.Descriptor instead.
func (*NewMomentResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{8}
}

func (x *NewMomentResp) GetMomentId() string {
	if x != nil {
		return x.MomentId
	}
	return ""
}

func (x *NewMomentResp) GetGetFish() bool {
	if x != nil {
		return x.GetFish
	}
	return false
}

func (x *NewMomentResp) GetGetFishTimes() int64 {
	if x != nil {
		return x.GetFishTimes
	}
	return 0
}

type SearchMomentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommunityId      *string                  `protobuf:"bytes,1,opt,name=communityId,proto3,oneof" json:"communityId" form:"communityId" query:"communityId"`
	IsParent         *bool                    `protobuf:"varint,2,opt,name=isParent,proto3,oneof" json:"isParent" form:"isParent" query:"isParent"`
	OnlyUserId       *string                  `protobuf:"bytes,3,opt,name=onlyUserId,proto3,oneof" json:"onlyUserId" form:"onlyUserId" query:"onlyUserId"`
	Keyword          *string                  `protobuf:"bytes,4,opt,name=keyword,proto3,oneof" json:"keyword" form:"keyword" query:"keyword"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,5,opt,name=paginationOption,proto3" json:"paginationOption" form:"paginationOption" query:"paginationOption"`
}

func (x *SearchMomentReq) Reset() {
	*x = SearchMomentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMomentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMomentReq) ProtoMessage() {}

func (x *SearchMomentReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMomentReq.ProtoReflect.Descriptor instead.
func (*SearchMomentReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{9}
}

func (x *SearchMomentReq) GetCommunityId() string {
	if x != nil && x.CommunityId != nil {
		return *x.CommunityId
	}
	return ""
}

func (x *SearchMomentReq) GetIsParent() bool {
	if x != nil && x.IsParent != nil {
		return *x.IsParent
	}
	return false
}

func (x *SearchMomentReq) GetOnlyUserId() string {
	if x != nil && x.OnlyUserId != nil {
		return *x.OnlyUserId
	}
	return ""
}

func (x *SearchMomentReq) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

func (x *SearchMomentReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

type SearchMomentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moments []*Moment `protobuf:"bytes,1,rep,name=moments,proto3" json:"moments" form:"moments" query:"moments"`
	Total   int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total" form:"total" query:"total"`
}

func (x *SearchMomentResp) Reset() {
	*x = SearchMomentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_moment_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMomentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMomentResp) ProtoMessage() {}

func (x *SearchMomentResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_moment_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMomentResp.ProtoReflect.Descriptor instead.
func (*SearchMomentResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_moment_proto_rawDescGZIP(), []int{10}
}

func (x *SearchMomentResp) GetMoments() []*Moment {
	if x != nil {
		return x.Moments
	}
	return nil
}

func (x *SearchMomentResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_meowchat_core_api_moment_proto protoreflect.FileDescriptor

var file_meowchat_core_api_moment_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x61, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xf5, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x08, 0x69, 0x73,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x44,
	0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x78, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0xe5, 0x01, 0x0a, 0x0c, 0x4e, 0x65,
	0x77, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x61,
	0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x63, 0x61, 0x74,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x63, 0x61, 0x74, 0x49, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x69, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x46,
	0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x67, 0x65, 0x74, 0x46, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x9b, 0x02, 0x0a,
	0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x69, 0x73, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x08, 0x69, 0x73, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x6e, 0x6c, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a,
	0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x10,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x73, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5d, 0x0a, 0x10, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33,
	0x0a, 0x07, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x85, 0x01, 0x0a, 0x25, 0x63, 0x6f,
	0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67,
	0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x61, 0x70, 0x69, 0x42, 0x0b, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68,
	0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_moment_proto_rawDescOnce sync.Once
	file_meowchat_core_api_moment_proto_rawDescData = file_meowchat_core_api_moment_proto_rawDesc
)

func file_meowchat_core_api_moment_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_moment_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_moment_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_moment_proto_rawDescData)
	})
	return file_meowchat_core_api_moment_proto_rawDescData
}

var file_meowchat_core_api_moment_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_meowchat_core_api_moment_proto_goTypes = []interface{}{
	(*Moment)(nil),                  // 0: meowchat.core_api.Moment
	(*GetMomentPreviewsReq)(nil),    // 1: meowchat.core_api.GetMomentPreviewsReq
	(*GetMomentPreviewsResp)(nil),   // 2: meowchat.core_api.GetMomentPreviewsResp
	(*GetMomentDetailReq)(nil),      // 3: meowchat.core_api.GetMomentDetailReq
	(*GetMomentDetailResp)(nil),     // 4: meowchat.core_api.GetMomentDetailResp
	(*DeleteMomentReq)(nil),         // 5: meowchat.core_api.DeleteMomentReq
	(*DeleteMomentResp)(nil),        // 6: meowchat.core_api.DeleteMomentResp
	(*NewMomentReq)(nil),            // 7: meowchat.core_api.NewMomentReq
	(*NewMomentResp)(nil),           // 8: meowchat.core_api.NewMomentResp
	(*SearchMomentReq)(nil),         // 9: meowchat.core_api.SearchMomentReq
	(*SearchMomentResp)(nil),        // 10: meowchat.core_api.SearchMomentResp
	(*user.UserPreview)(nil),        // 11: meowchat.user.UserPreview
	(*basic.PaginationOptions)(nil), // 12: basic.PaginationOptions
}
var file_meowchat_core_api_moment_proto_depIdxs = []int32{
	11, // 0: meowchat.core_api.Moment.user:type_name -> meowchat.user.UserPreview
	12, // 1: meowchat.core_api.GetMomentPreviewsReq.paginationOption:type_name -> basic.PaginationOptions
	0,  // 2: meowchat.core_api.GetMomentPreviewsResp.moments:type_name -> meowchat.core_api.Moment
	0,  // 3: meowchat.core_api.GetMomentDetailResp.moment:type_name -> meowchat.core_api.Moment
	12, // 4: meowchat.core_api.SearchMomentReq.paginationOption:type_name -> basic.PaginationOptions
	0,  // 5: meowchat.core_api.SearchMomentResp.moments:type_name -> meowchat.core_api.Moment
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func file_meowchat_core_api_moment_proto_init() {
	if File_meowchat_core_api_moment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_moment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Moment); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMomentPreviewsReq); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMomentPreviewsResp); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMomentDetailReq); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMomentDetailResp); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMomentReq); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMomentResp); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMomentReq); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewMomentResp); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMomentReq); i {
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
		file_meowchat_core_api_moment_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMomentResp); i {
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
	file_meowchat_core_api_moment_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_meowchat_core_api_moment_proto_msgTypes[7].OneofWrappers = []interface{}{}
	file_meowchat_core_api_moment_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_moment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_moment_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_moment_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_moment_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_moment_proto = out.File
	file_meowchat_core_api_moment_proto_rawDesc = nil
	file_meowchat_core_api_moment_proto_goTypes = nil
	file_meowchat_core_api_moment_proto_depIdxs = nil
}
