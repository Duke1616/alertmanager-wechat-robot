// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/token/pb/token.proto

package token

import (
	user "github.com/Duke1616/alertmanager-wechat-robot/apps/user"
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

// 令牌类型
type PLATFORM int32

const (
	// Web 登陆授权
	PLATFORM_WEB PLATFORM = 0
	// TODO 接收报警数据需要
	PLATFORM_API PLATFORM = 1
)

// Enum value maps for PLATFORM.
var (
	PLATFORM_name = map[int32]string{
		0: "WEB",
		1: "API",
	}
	PLATFORM_value = map[string]int32{
		"WEB": 0,
		"API": 1,
	}
)

func (x PLATFORM) Enum() *PLATFORM {
	p := new(PLATFORM)
	*p = x
	return p
}

func (x PLATFORM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PLATFORM) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_token_pb_token_proto_enumTypes[0].Descriptor()
}

func (PLATFORM) Type() protoreflect.EnumType {
	return &file_apps_token_pb_token_proto_enumTypes[0]
}

func (x PLATFORM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PLATFORM.Descriptor instead.
func (PLATFORM) EnumDescriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{0}
}

type DESCRIBE_BY int32

const (
	// 通过access token查看
	DESCRIBE_BY_ACCESS_TOKEN DESCRIBE_BY = 0
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "ACCESS_TOKEN",
	}
	DESCRIBE_BY_value = map[string]int32{
		"ACCESS_TOKEN": 0,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_token_pb_token_proto_enumTypes[1].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_apps_token_pb_token_proto_enumTypes[1]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{1}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 颁发平台, 根据授权方式判断
	// @gotags: bson:"platform" json:"platform"
	Platform PLATFORM `protobuf:"varint,1,opt,name=platform,proto3,enum=robot.token.PLATFORM" json:"platform" bson:"platform"`
	// 访问令牌
	// @gotags: bson:"_id" json:"access_token"
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"_id"`
	// 颁发时间
	// @gotags: bson:"issue_at" json:"issue_at"
	IssueAt int64 `protobuf:"varint,3,opt,name=issue_at,json=issueAt,proto3" json:"issue_at" bson:"issue_at"`
	// 访问令牌过期时间
	// @gotags: bson:"access_expired_at" json:"access_expired_at"
	AccessExpiredAt int64 `protobuf:"varint,4,opt,name=access_expired_at,json=accessExpiredAt,proto3" json:"access_expired_at" bson:"access_expired_at"`
	// 用户类型
	// @gotags: bson:"user_type" json:"user_type"
	UserType user.USER_TYPE `protobuf:"varint,5,opt,name=user_type,json=userType,proto3,enum=robot.user.USER_TYPE" json:"user_type" bson:"user_type"`
	// 用户名
	// @gotags: bson:"username" json:"username"
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username" bson:"username"`
	// 用户Id
	// @gotags: bson:"user_id" json:"user_id"
	UserId string `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id" bson:"user_id"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetPlatform() PLATFORM {
	if x != nil {
		return x.Platform
	}
	return PLATFORM_WEB
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetIssueAt() int64 {
	if x != nil {
		return x.IssueAt
	}
	return 0
}

func (x *Token) GetAccessExpiredAt() int64 {
	if x != nil {
		return x.AccessExpiredAt
	}
	return 0
}

func (x *Token) GetUserType() user.USER_TYPE {
	if x != nil {
		return x.UserType
	}
	return user.USER_TYPE(0)
}

func (x *Token) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Token) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type IssueTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PASSWORD授权时, 用户名
	// @gotags: json:"username,omitempty"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// PASSWORD授权时, 用户密码
	// @gotags: json:"password,omitempty"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *IssueTokenRequest) Reset() {
	*x = IssueTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenRequest) ProtoMessage() {}

func (x *IssueTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{1}
}

func (x *IssueTokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RemoveTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 令牌
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
}

func (x *RemoveTokenRequest) Reset() {
	*x = RemoveTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTokenRequest) ProtoMessage() {}

func (x *RemoveTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTokenRequest.ProtoReflect.Descriptor instead.
func (*RemoveTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type DescribeTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 参数类型
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=robot.token.DESCRIBE_BY" json:"describe_by"`
	// 参数值
	// @gotags: json:"access_token"
	AccessToken string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token"`
}

func (x *DescribeTokenRequest) Reset() {
	*x = DescribeTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTokenRequest) ProtoMessage() {}

func (x *DescribeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTokenRequest.ProtoReflect.Descriptor instead.
func (*DescribeTokenRequest) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeTokenRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_ACCESS_TOKEN
}

func (x *DescribeTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_apps_token_pb_token_proto protoreflect.FileDescriptor

var file_apps_token_pb_token_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4b, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x39,
	0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x37, 0x0a, 0x12, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x74, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0x1c, 0x0a, 0x08, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x50, 0x49, 0x10, 0x01, 0x2a, 0x1f, 0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49,
	0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x00, 0x32, 0x9b, 0x02, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x40, 0x0a, 0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x46, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x42, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1f, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x75, 0x6b, 0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x2d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_token_pb_token_proto_rawDescOnce sync.Once
	file_apps_token_pb_token_proto_rawDescData = file_apps_token_pb_token_proto_rawDesc
)

func file_apps_token_pb_token_proto_rawDescGZIP() []byte {
	file_apps_token_pb_token_proto_rawDescOnce.Do(func() {
		file_apps_token_pb_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_token_pb_token_proto_rawDescData)
	})
	return file_apps_token_pb_token_proto_rawDescData
}

var file_apps_token_pb_token_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_token_pb_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_apps_token_pb_token_proto_goTypes = []interface{}{
	(PLATFORM)(0),                // 0: robot.token.PLATFORM
	(DESCRIBE_BY)(0),             // 1: robot.token.DESCRIBE_BY
	(*Token)(nil),                // 2: robot.token.Token
	(*IssueTokenRequest)(nil),    // 3: robot.token.IssueTokenRequest
	(*ValidateTokenRequest)(nil), // 4: robot.token.ValidateTokenRequest
	(*RemoveTokenRequest)(nil),   // 5: robot.token.RemoveTokenRequest
	(*DescribeTokenRequest)(nil), // 6: robot.token.DescribeTokenRequest
	(user.USER_TYPE)(0),          // 7: robot.user.USER_TYPE
}
var file_apps_token_pb_token_proto_depIdxs = []int32{
	0, // 0: robot.token.Token.platform:type_name -> robot.token.PLATFORM
	7, // 1: robot.token.Token.user_type:type_name -> robot.user.USER_TYPE
	1, // 2: robot.token.DescribeTokenRequest.describe_by:type_name -> robot.token.DESCRIBE_BY
	3, // 3: robot.token.RPC.IssueToken:input_type -> robot.token.IssueTokenRequest
	4, // 4: robot.token.RPC.ValidateToken:input_type -> robot.token.ValidateTokenRequest
	6, // 5: robot.token.RPC.DescribeToken:input_type -> robot.token.DescribeTokenRequest
	5, // 6: robot.token.RPC.RemoveToken:input_type -> robot.token.RemoveTokenRequest
	2, // 7: robot.token.RPC.IssueToken:output_type -> robot.token.Token
	2, // 8: robot.token.RPC.ValidateToken:output_type -> robot.token.Token
	2, // 9: robot.token.RPC.DescribeToken:output_type -> robot.token.Token
	2, // 10: robot.token.RPC.RemoveToken:output_type -> robot.token.Token
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_token_pb_token_proto_init() }
func file_apps_token_pb_token_proto_init() {
	if File_apps_token_pb_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_token_pb_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_apps_token_pb_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueTokenRequest); i {
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
		file_apps_token_pb_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenRequest); i {
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
		file_apps_token_pb_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTokenRequest); i {
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
		file_apps_token_pb_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTokenRequest); i {
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
			RawDescriptor: file_apps_token_pb_token_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_token_pb_token_proto_goTypes,
		DependencyIndexes: file_apps_token_pb_token_proto_depIdxs,
		EnumInfos:         file_apps_token_pb_token_proto_enumTypes,
		MessageInfos:      file_apps_token_pb_token_proto_msgTypes,
	}.Build()
	File_apps_token_pb_token_proto = out.File
	file_apps_token_pb_token_proto_rawDesc = nil
	file_apps_token_pb_token_proto_goTypes = nil
	file_apps_token_pb_token_proto_depIdxs = nil
}
