// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: idl/comment.proto

package commentproto

import (
	context "context"
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

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type CommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId  int64  `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreateDate string `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"` //mm-dd
}

func (x *CommentInfo) Reset() {
	*x = CommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentInfo) ProtoMessage() {}

func (x *CommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentInfo.ProtoReflect.Descriptor instead.
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentInfo) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *CommentInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentInfo) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type CreateCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoId int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateCommentReq) Reset() {
	*x = CreateCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentReq) ProtoMessage() {}

func (x *CreateCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentReq.ProtoReflect.Descriptor instead.
func (*CreateCommentReq) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCommentReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCommentReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CreateCommentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp    *BaseResp    `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	CommentInfo *CommentInfo `protobuf:"bytes,2,opt,name=comment_info,json=commentInfo,proto3" json:"comment_info,omitempty"`
}

func (x *CreateCommentResp) Reset() {
	*x = CreateCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResp) ProtoMessage() {}

func (x *CreateCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResp.ProtoReflect.Descriptor instead.
func (*CreateCommentResp) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommentResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *CreateCommentResp) GetCommentInfo() *CommentInfo {
	if x != nil {
		return x.CommentInfo
	}
	return nil
}

type DeleteCommentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId int64 `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	VideoId   int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *DeleteCommentReq) Reset() {
	*x = DeleteCommentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentReq) ProtoMessage() {}

func (x *DeleteCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[4]
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
	return file_idl_comment_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCommentReq) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

func (x *DeleteCommentReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type DeleteCommentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *DeleteCommentResp) Reset() {
	*x = DeleteCommentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResp) ProtoMessage() {}

func (x *DeleteCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[5]
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
	return file_idl_comment_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCommentResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type GetCommentListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *GetCommentListReq) Reset() {
	*x = GetCommentListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentListReq) ProtoMessage() {}

func (x *GetCommentListReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentListReq.ProtoReflect.Descriptor instead.
func (*GetCommentListReq) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{6}
}

func (x *GetCommentListReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type GetCommentListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp     *BaseResp      `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	CommentInfos []*CommentInfo `protobuf:"bytes,2,rep,name=comment_infos,json=commentInfos,proto3" json:"comment_infos,omitempty"`
}

func (x *GetCommentListResp) Reset() {
	*x = GetCommentListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentListResp) ProtoMessage() {}

func (x *GetCommentListResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentListResp.ProtoReflect.Descriptor instead.
func (*GetCommentListResp) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{7}
}

func (x *GetCommentListResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *GetCommentListResp) GetCommentInfos() []*CommentInfo {
	if x != nil {
		return x.CommentInfos
	}
	return nil
}

var File_idl_comment_proto protoreflect.FileDescriptor

var file_idl_comment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x08,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x60, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7c, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e,
	0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2e,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x7f,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32,
	0xeb, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x28, 0x5a,
	0x26, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x67, 0x65, 0x6e,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_comment_proto_rawDescOnce sync.Once
	file_idl_comment_proto_rawDescData = file_idl_comment_proto_rawDesc
)

func file_idl_comment_proto_rawDescGZIP() []byte {
	file_idl_comment_proto_rawDescOnce.Do(func() {
		file_idl_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_comment_proto_rawDescData)
	})
	return file_idl_comment_proto_rawDescData
}

var file_idl_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_idl_comment_proto_goTypes = []interface{}{
	(*BaseResp)(nil),           // 0: comment.BaseResp
	(*CommentInfo)(nil),        // 1: comment.CommentInfo
	(*CreateCommentReq)(nil),   // 2: comment.CreateCommentReq
	(*CreateCommentResp)(nil),  // 3: comment.CreateCommentResp
	(*DeleteCommentReq)(nil),   // 4: comment.DeleteCommentReq
	(*DeleteCommentResp)(nil),  // 5: comment.DeleteCommentResp
	(*GetCommentListReq)(nil),  // 6: comment.GetCommentListReq
	(*GetCommentListResp)(nil), // 7: comment.GetCommentListResp
}
var file_idl_comment_proto_depIdxs = []int32{
	0, // 0: comment.CreateCommentResp.base_resp:type_name -> comment.BaseResp
	1, // 1: comment.CreateCommentResp.comment_info:type_name -> comment.CommentInfo
	0, // 2: comment.DeleteCommentResp.base_resp:type_name -> comment.BaseResp
	0, // 3: comment.GetCommentListResp.base_resp:type_name -> comment.BaseResp
	1, // 4: comment.GetCommentListResp.comment_infos:type_name -> comment.CommentInfo
	2, // 5: comment.CommentService.CreateComment:input_type -> comment.CreateCommentReq
	4, // 6: comment.CommentService.DeleteComment:input_type -> comment.DeleteCommentReq
	6, // 7: comment.CommentService.GetCommentList:input_type -> comment.GetCommentListReq
	3, // 8: comment.CommentService.CreateComment:output_type -> comment.CreateCommentResp
	5, // 9: comment.CommentService.DeleteComment:output_type -> comment.DeleteCommentResp
	7, // 10: comment.CommentService.GetCommentList:output_type -> comment.GetCommentListResp
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_idl_comment_proto_init() }
func file_idl_comment_proto_init() {
	if File_idl_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_idl_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentInfo); i {
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
		file_idl_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentReq); i {
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
		file_idl_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResp); i {
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
		file_idl_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentListReq); i {
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
		file_idl_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentListResp); i {
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
			RawDescriptor: file_idl_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_comment_proto_goTypes,
		DependencyIndexes: file_idl_comment_proto_depIdxs,
		MessageInfos:      file_idl_comment_proto_msgTypes,
	}.Build()
	File_idl_comment_proto = out.File
	file_idl_comment_proto_rawDesc = nil
	file_idl_comment_proto_goTypes = nil
	file_idl_comment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type CommentService interface {
	CreateComment(ctx context.Context, req *CreateCommentReq) (res *CreateCommentResp, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentReq) (res *DeleteCommentResp, err error)
	GetCommentList(ctx context.Context, req *GetCommentListReq) (res *GetCommentListResp, err error)
}