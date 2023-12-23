// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: platform/comment/common.proto

package comment

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

type CommentType int32

const (
	CommentType_CommentType_Unknown CommentType = 0
	CommentType_CommentType_Comment CommentType = 1
	CommentType_CommentType_Post    CommentType = 2
	CommentType_CommentType_Moment  CommentType = 3
)

// Enum value maps for CommentType.
var (
	CommentType_name = map[int32]string{
		0: "CommentType_Unknown",
		1: "CommentType_Comment",
		2: "CommentType_Post",
		3: "CommentType_Moment",
	}
	CommentType_value = map[string]int32{
		"CommentType_Unknown": 0,
		"CommentType_Comment": 1,
		"CommentType_Post":    2,
		"CommentType_Moment":  3,
	}
)

func (x CommentType) Enum() *CommentType {
	p := new(CommentType)
	*p = x
	return p
}

func (x CommentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentType) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_comment_common_proto_enumTypes[0].Descriptor()
}

func (CommentType) Type() protoreflect.EnumType {
	return &file_platform_comment_common_proto_enumTypes[0]
}

func (x CommentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentType.Descriptor instead.
func (CommentType) EnumDescriptor() ([]byte, []int) {
	return file_platform_comment_common_proto_rawDescGZIP(), []int{0}
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id" form:"id" query:"id"`
	FirstLevelId string      `protobuf:"bytes,2,opt,name=firstLevelId,proto3" json:"firstLevelId" form:"firstLevelId" query:"firstLevelId"`
	Text         string      `protobuf:"bytes,3,opt,name=text,proto3" json:"text" form:"text" query:"text"`
	AuthorId     string      `protobuf:"bytes,4,opt,name=authorId,proto3" json:"authorId" form:"authorId" query:"authorId"`
	ReplyTo      string      `protobuf:"bytes,5,opt,name=replyTo,proto3" json:"replyTo" form:"replyTo" query:"replyTo"`
	Type         CommentType `protobuf:"varint,6,opt,name=type,proto3,enum=platform.comment.CommentType" json:"type" form:"type" query:"type"`
	ParentId     string      `protobuf:"bytes,7,opt,name=parentId,proto3" json:"parentId" form:"parentId" query:"parentId"`
	UpdateAt     int64       `protobuf:"varint,8,opt,name=updateAt,proto3" json:"updateAt" form:"updateAt" query:"updateAt"`
	CreateAt     int64       `protobuf:"varint,9,opt,name=createAt,proto3" json:"createAt" form:"createAt" query:"createAt"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_comment_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_platform_comment_common_proto_msgTypes[0]
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
	return file_platform_comment_common_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Comment) GetFirstLevelId() string {
	if x != nil {
		return x.FirstLevelId
	}
	return ""
}

func (x *Comment) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Comment) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Comment) GetReplyTo() string {
	if x != nil {
		return x.ReplyTo
	}
	return ""
}

func (x *Comment) GetType() CommentType {
	if x != nil {
		return x.Type
	}
	return CommentType_CommentType_Unknown
}

func (x *Comment) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Comment) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Comment) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

var File_platform_comment_common_proto protoreflect.FileDescriptor

var file_platform_comment_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x8e, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x12, 0x31, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x2a, 0x6d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x50, 0x6f, 0x73, 0x74, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x10,
	0x03, 0x42, 0x84, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61,
	0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_comment_common_proto_rawDescOnce sync.Once
	file_platform_comment_common_proto_rawDescData = file_platform_comment_common_proto_rawDesc
)

func file_platform_comment_common_proto_rawDescGZIP() []byte {
	file_platform_comment_common_proto_rawDescOnce.Do(func() {
		file_platform_comment_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_comment_common_proto_rawDescData)
	})
	return file_platform_comment_common_proto_rawDescData
}

var file_platform_comment_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_platform_comment_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_platform_comment_common_proto_goTypes = []interface{}{
	(CommentType)(0), // 0: platform.comment.CommentType
	(*Comment)(nil),  // 1: platform.comment.Comment
}
var file_platform_comment_common_proto_depIdxs = []int32{
	0, // 0: platform.comment.Comment.type:type_name -> platform.comment.CommentType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func file_platform_comment_common_proto_init() {
	if File_platform_comment_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_comment_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_comment_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_comment_common_proto_goTypes,
		DependencyIndexes: file_platform_comment_common_proto_depIdxs,
		EnumInfos:         file_platform_comment_common_proto_enumTypes,
		MessageInfos:      file_platform_comment_common_proto_msgTypes,
	}.Build()
	File_platform_comment_common_proto = out.File
	file_platform_comment_common_proto_rawDesc = nil
	file_platform_comment_common_proto_goTypes = nil
	file_platform_comment_common_proto_depIdxs = nil
}
