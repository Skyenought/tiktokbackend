// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: video.proto

package videorpc

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

type FavoriteActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID     int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	VideoID    int64 `protobuf:"varint,2,opt,name=videoID,proto3" json:"videoID,omitempty"`
	ActionType int32 `protobuf:"varint,3,opt,name=actionType,proto3" json:"actionType,omitempty"`
}

func (x *FavoriteActionReq) Reset() {
	*x = FavoriteActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionReq) ProtoMessage() {}

func (x *FavoriteActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionReq.ProtoReflect.Descriptor instead.
func (*FavoriteActionReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *FavoriteActionReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *FavoriteActionReq) GetVideoID() int64 {
	if x != nil {
		return x.VideoID
	}
	return 0
}

func (x *FavoriteActionReq) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type FavoriteActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *FavoriteActionResp) Reset() {
	*x = FavoriteActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteActionResp) ProtoMessage() {}

func (x *FavoriteActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteActionResp.ProtoReflect.Descriptor instead.
func (*FavoriteActionResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *FavoriteActionResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type FavoriteVideoListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *FavoriteVideoListReq) Reset() {
	*x = FavoriteVideoListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteVideoListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteVideoListReq) ProtoMessage() {}

func (x *FavoriteVideoListReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteVideoListReq.ProtoReflect.Descriptor instead.
func (*FavoriteVideoListReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *FavoriteVideoListReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type FavoriteVideoListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	VideoList  []*Video `protobuf:"bytes,2,rep,name=VideoList,proto3" json:"VideoList,omitempty"`
}

func (x *FavoriteVideoListResp) Reset() {
	*x = FavoriteVideoListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoriteVideoListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoriteVideoListResp) ProtoMessage() {}

func (x *FavoriteVideoListResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoriteVideoListResp.ProtoReflect.Descriptor instead.
func (*FavoriteVideoListResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *FavoriteVideoListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FavoriteVideoListResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type PublishActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *PublishActionReq) Reset() {
	*x = PublishActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishActionReq) ProtoMessage() {}

func (x *PublishActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishActionReq.ProtoReflect.Descriptor instead.
func (*PublishActionReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

func (x *PublishActionReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *PublishActionReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PublishActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *PublishActionResp) Reset() {
	*x = PublishActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishActionResp) ProtoMessage() {}

func (x *PublishActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishActionResp.ProtoReflect.Descriptor instead.
func (*PublishActionResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

func (x *PublishActionResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type PublishlistReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *PublishlistReq) Reset() {
	*x = PublishlistReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishlistReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishlistReq) ProtoMessage() {}

func (x *PublishlistReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishlistReq.ProtoReflect.Descriptor instead.
func (*PublishlistReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{6}
}

func (x *PublishlistReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type PublishlistResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	VideoList  []*Video `protobuf:"bytes,2,rep,name=videoList,proto3" json:"videoList,omitempty"`
}

func (x *PublishlistResp) Reset() {
	*x = PublishlistResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishlistResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishlistResp) ProtoMessage() {}

func (x *PublishlistResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishlistResp.ProtoReflect.Descriptor instead.
func (*PublishlistResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{7}
}

func (x *PublishlistResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PublishlistResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime int64 `protobuf:"varint,1,opt,name=latestTime,proto3" json:"latestTime,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{8}
}

func (x *FeedReq) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

type FeedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	VideoList  []*Video `protobuf:"bytes,2,rep,name=videoList,proto3" json:"videoList,omitempty"`
	NextTime   int64    `protobuf:"varint,3,opt,name=nextTime,proto3" json:"nextTime,omitempty"`
}

func (x *FeedResp) Reset() {
	*x = FeedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResp) ProtoMessage() {}

func (x *FeedResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResp.ProtoReflect.Descriptor instead.
func (*FeedResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{9}
}

func (x *FeedResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FeedResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *FeedResp) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID        int64  `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	PlayURL       string `protobuf:"bytes,3,opt,name=playURL,proto3" json:"playURL,omitempty"`
	CoverURL      string `protobuf:"bytes,4,opt,name=coverURL,proto3" json:"coverURL,omitempty"`
	FavoriteCount int32  `protobuf:"varint,5,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"`
	CommentCount  int32  `protobuf:"varint,6,opt,name=commentCount,proto3" json:"commentCount,omitempty"`
	Title         string `protobuf:"bytes,7,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{10}
}

func (x *Video) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Video) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Video) GetPlayURL() string {
	if x != nil {
		return x.PlayURL
	}
	return ""
}

func (x *Video) GetCoverURL() string {
	if x != nil {
		return x.CoverURL
	}
	return ""
}

func (x *Video) GetFavoriteCount() int32 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type UpdateCommentCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoID    int64 `protobuf:"varint,1,opt,name=videoID,proto3" json:"videoID,omitempty"`
	ActionType int32 `protobuf:"varint,2,opt,name=actionType,proto3" json:"actionType,omitempty"`
}

func (x *UpdateCommentCountReq) Reset() {
	*x = UpdateCommentCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentCountReq) ProtoMessage() {}

func (x *UpdateCommentCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentCountReq.ProtoReflect.Descriptor instead.
func (*UpdateCommentCountReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateCommentCountReq) GetVideoID() int64 {
	if x != nil {
		return x.VideoID
	}
	return 0
}

func (x *UpdateCommentCountReq) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type UpdateCommentCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *UpdateCommentCountResp) Reset() {
	*x = UpdateCommentCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentCountResp) ProtoMessage() {}

func (x *UpdateCommentCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentCountResp.ProtoReflect.Descriptor instead.
func (*UpdateCommentCountResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateCommentCountResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x22, 0x65, 0x0a, 0x11, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x34,
	0x0a, 0x12, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x66, 0x0a, 0x15, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a,
	0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x10,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x33,
	0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x60, 0x0a,
	0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x2d, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x29, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x75, 0x0a, 0x08, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0xc5, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x52, 0x4c, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x51, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x38, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xc7, 0x03, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70,
	0x63, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x11, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x0d, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72,
	0x70, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64,
	0x12, 0x11, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x57, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_video_proto_goTypes = []interface{}{
	(*FavoriteActionReq)(nil),      // 0: videorpc.FavoriteActionReq
	(*FavoriteActionResp)(nil),     // 1: videorpc.FavoriteActionResp
	(*FavoriteVideoListReq)(nil),   // 2: videorpc.FavoriteVideoListReq
	(*FavoriteVideoListResp)(nil),  // 3: videorpc.FavoriteVideoListResp
	(*PublishActionReq)(nil),       // 4: videorpc.PublishActionReq
	(*PublishActionResp)(nil),      // 5: videorpc.PublishActionResp
	(*PublishlistReq)(nil),         // 6: videorpc.PublishlistReq
	(*PublishlistResp)(nil),        // 7: videorpc.PublishlistResp
	(*FeedReq)(nil),                // 8: videorpc.FeedReq
	(*FeedResp)(nil),               // 9: videorpc.FeedResp
	(*Video)(nil),                  // 10: videorpc.Video
	(*UpdateCommentCountReq)(nil),  // 11: videorpc.UpdateCommentCountReq
	(*UpdateCommentCountResp)(nil), // 12: videorpc.UpdateCommentCountResp
}
var file_video_proto_depIdxs = []int32{
	10, // 0: videorpc.FavoriteVideoListResp.VideoList:type_name -> videorpc.Video
	10, // 1: videorpc.PublishlistResp.videoList:type_name -> videorpc.Video
	10, // 2: videorpc.FeedResp.videoList:type_name -> videorpc.Video
	0,  // 3: videorpc.VideoService.FavoriteAction:input_type -> videorpc.FavoriteActionReq
	2,  // 4: videorpc.VideoService.FavoriteVideoList:input_type -> videorpc.FavoriteVideoListReq
	4,  // 5: videorpc.VideoService.PublishAction:input_type -> videorpc.PublishActionReq
	6,  // 6: videorpc.VideoService.PublishList:input_type -> videorpc.PublishlistReq
	8,  // 7: videorpc.VideoService.Feed:input_type -> videorpc.FeedReq
	11, // 8: videorpc.VideoService.UpdateCommentCount:input_type -> videorpc.UpdateCommentCountReq
	1,  // 9: videorpc.VideoService.FavoriteAction:output_type -> videorpc.FavoriteActionResp
	3,  // 10: videorpc.VideoService.FavoriteVideoList:output_type -> videorpc.FavoriteVideoListResp
	5,  // 11: videorpc.VideoService.PublishAction:output_type -> videorpc.PublishActionResp
	7,  // 12: videorpc.VideoService.PublishList:output_type -> videorpc.PublishlistResp
	9,  // 13: videorpc.VideoService.Feed:output_type -> videorpc.FeedResp
	12, // 14: videorpc.VideoService.UpdateCommentCount:output_type -> videorpc.UpdateCommentCountResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionReq); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteActionResp); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteVideoListReq); i {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoriteVideoListResp); i {
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
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishActionReq); i {
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
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishActionResp); i {
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
		file_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishlistReq); i {
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
		file_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishlistResp); i {
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
		file_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResp); i {
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
		file_video_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_video_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentCountReq); i {
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
		file_video_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentCountResp); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}