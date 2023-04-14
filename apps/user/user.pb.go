// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/user/pb/user.proto

package user

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
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

type USER_TYPE int32

const (
	// 报警通知用户，不可以登陆
	USER_TYPE_STAFF USER_TYPE = 0
	// 系统登陆用户
	USER_TYPE_SYSTEM USER_TYPE = 1
)

// Enum value maps for USER_TYPE.
var (
	USER_TYPE_name = map[int32]string{
		0: "STAFF",
		1: "SYSTEM",
	}
	USER_TYPE_value = map[string]int32{
		"STAFF":  0,
		"SYSTEM": 1,
	}
)

func (x USER_TYPE) Enum() *USER_TYPE {
	p := new(USER_TYPE)
	*p = x
	return p
}

func (x USER_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (USER_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_user_pb_user_proto_enumTypes[0].Descriptor()
}

func (USER_TYPE) Type() protoreflect.EnumType {
	return &file_apps_user_pb_user_proto_enumTypes[0]
}

func (x USER_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use USER_TYPE.Descriptor instead.
func (USER_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{0}
}

type DESCRIBE_BY int32

const (
	// 根据用户ID
	DESCRIBE_BY_ID DESCRIBE_BY = 0
	// 根据企业微信ID
	DESCRIBE_BY_WECHAT_ID DESCRIBE_BY = 1
	// 根据用户名称
	DESCRIBE_BY_NAME DESCRIBE_BY = 2
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "ID",
		1: "WECHAT_ID",
		2: "NAME",
	}
	DESCRIBE_BY_value = map[string]int32{
		"ID":        0,
		"WECHAT_ID": 1,
		"NAME":      2,
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
	return file_apps_user_pb_user_proto_enumTypes[1].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_apps_user_pb_user_proto_enumTypes[1]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{1}
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id信息
	// @gotags: bson:"name" json:"name" validate:"required,lte=80"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=80"`
	// id信息
	// @gotags: bson:"password" json:"password" validate:"required,lte=80"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password" bson:"password" validate:"required,lte=80"`
	// 用户ID，wechat_id
	// @gotags: bson:"wechat_id" json:"wechat_id"
	WechatId string `protobuf:"bytes,3,opt,name=wechat_id,json=wechatId,proto3" json:"wechat_id" bson:"wechat_id"`
	// target_names
	// @gotags: bson:"target_names" json:"target_names"
	TargetNames []string `protobuf:"bytes,4,rep,name=target_names,json=targetNames,proto3" json:"target_names" bson:"target_names"`
	// 用户类型
	// @gotags: bson:"user_type" json:"user_type"
	UserType USER_TYPE `protobuf:"varint,5,opt,name=user_type,json=userType,proto3,enum=robot.user.USER_TYPE" json:"user_type" bson:"user_type"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetWechatId() string {
	if x != nil {
		return x.WechatId
	}
	return ""
}

func (x *CreateUserRequest) GetTargetNames() []string {
	if x != nil {
		return x.TargetNames
	}
	return nil
}

func (x *CreateUserRequest) GetUserType() USER_TYPE {
	if x != nil {
		return x.UserType
	}
	return USER_TYPE_STAFF
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 电话号
	// @gotags: bson:"phone" json:"phone"
	Phone int64 `protobuf:"varint,1,opt,name=phone,proto3" json:"phone" bson:"phone"`
	// 职位
	// @gotags: bson:"position" json:"position"
	Position string `protobuf:"bytes,2,opt,name=position,proto3" json:"position" bson:"position"`
	// 邮箱
	// @gotags: bson:"email" json:"email"
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email" bson:"email"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

func (x *Profile) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id信息
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,2,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 创建时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,3,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 详细信息
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateUserRequest `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec" bson:"spec"`
	// 用户详细信息
	// @gotags: bson:"profile" json:"profile"
	Profile *Profile `protobuf:"bytes,5,opt,name=profile,proto3" json:"profile" bson:"profile"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *User) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *User) GetSpec() *CreateUserRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *User) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type UserSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 数据
	// @gotags: bson:"items" json:"items"
	Items []*User `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *UserSet) Reset() {
	*x = UserSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSet) ProtoMessage() {}

func (x *UserSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSet.ProtoReflect.Descriptor instead.
func (*UserSet) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserSet) GetItems() []*User {
	if x != nil {
		return x.Items
	}
	return nil
}

type DescribeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=robot.user.DESCRIBE_BY" json:"describe_by"`
	// 通过Id查询
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 通过name查询
	// @gotags: json:"wechat_id"
	WechatId string `protobuf:"bytes,3,opt,name=wechat_id,json=wechatId,proto3" json:"wechat_id"`
	// 通过name查询
	// @gotags: json:"name"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`
}

func (x *DescribeUserRequest) Reset() {
	*x = DescribeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserRequest) ProtoMessage() {}

func (x *DescribeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserRequest.ProtoReflect.Descriptor instead.
func (*DescribeUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeUserRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_ID
}

func (x *DescribeUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribeUserRequest) GetWechatId() string {
	if x != nil {
		return x.WechatId
	}
	return ""
}

func (x *DescribeUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type QueryUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// user_ids信息
	// @gotags: json:"user_ids"
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids"`
	// target_names信息
	// @gotags: json:"target_names"
	TargetNames []string `protobuf:"bytes,3,rep,name=target_names,json=targetNames,proto3" json:"target_names"`
}

func (x *QueryUserRequest) Reset() {
	*x = QueryUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserRequest) ProtoMessage() {}

func (x *QueryUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserRequest.ProtoReflect.Descriptor instead.
func (*QueryUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *QueryUserRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryUserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *QueryUserRequest) GetTargetNames() []string {
	if x != nil {
		return x.TargetNames
	}
	return nil
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名称
	// @gotags: json:"name"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// 登陆密码
	// @gotags: json:"password"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ValidateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 用户Id
	// @gotags: json:"user_id" validate:"required,lte=120"
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id" validate:"required,lte=120"`
	// profile 用户信息
	// @gotags: json:"profile"
	Profile *Profile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile"`
	// profile 用户信息
	// @gotags: json:"wechat_id"
	WechatId string `protobuf:"bytes,4,opt,name=wechat_id,json=wechatId,proto3" json:"wechat_id"`
	// target_id
	// @gotags: bson:"target_names" json:"target_names"
	TargetNames []string `protobuf:"bytes,5,rep,name=target_names,json=targetNames,proto3" json:"target_names" bson:"target_names"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *UpdateUserRequest) GetWechatId() string {
	if x != nil {
		return x.WechatId
	}
	return ""
}

func (x *UpdateUserRequest) GetTargetNames() []string {
	if x != nil {
		return x.TargetNames
	}
	return nil
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户账号id
	// @gotags: json:"user_ids" validate:"required,lte=60"
	UserIds []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids" validate:"required,lte=60"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_user_pb_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_user_pb_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_apps_user_pb_user_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

var File_apps_user_pb_user_proto protoreflect.FileDescriptor

var file_apps_user_pb_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x38, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb7, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xb2, 0x01, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x31, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x47, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59,
	0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x88, 0x01,
	0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xe2, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x22, 0x2e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x2a, 0x22, 0x0a, 0x09, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x54, 0x41, 0x46, 0x46, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x01, 0x2a, 0x2e, 0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45,
	0x5f, 0x42, 0x59, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x57,
	0x45, 0x43, 0x48, 0x41, 0x54, 0x5f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x02, 0x32, 0x87, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x3d, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3e,
	0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x3d,
	0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x40, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12,
	0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x75, 0x6b,
	0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_user_pb_user_proto_rawDescOnce sync.Once
	file_apps_user_pb_user_proto_rawDescData = file_apps_user_pb_user_proto_rawDesc
)

func file_apps_user_pb_user_proto_rawDescGZIP() []byte {
	file_apps_user_pb_user_proto_rawDescOnce.Do(func() {
		file_apps_user_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_user_pb_user_proto_rawDescData)
	})
	return file_apps_user_pb_user_proto_rawDescData
}

var file_apps_user_pb_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_user_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_apps_user_pb_user_proto_goTypes = []interface{}{
	(USER_TYPE)(0),              // 0: robot.user.USER_TYPE
	(DESCRIBE_BY)(0),            // 1: robot.user.DESCRIBE_BY
	(*CreateUserRequest)(nil),   // 2: robot.user.CreateUserRequest
	(*Profile)(nil),             // 3: robot.user.Profile
	(*User)(nil),                // 4: robot.user.User
	(*UserSet)(nil),             // 5: robot.user.UserSet
	(*DescribeUserRequest)(nil), // 6: robot.user.DescribeUserRequest
	(*QueryUserRequest)(nil),    // 7: robot.user.QueryUserRequest
	(*ValidateRequest)(nil),     // 8: robot.user.ValidateRequest
	(*UpdateUserRequest)(nil),   // 9: robot.user.UpdateUserRequest
	(*DeleteUserRequest)(nil),   // 10: robot.user.DeleteUserRequest
	(*request.PageRequest)(nil), // 11: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),    // 12: infraboard.mcube.request.UpdateMode
}
var file_apps_user_pb_user_proto_depIdxs = []int32{
	0,  // 0: robot.user.CreateUserRequest.user_type:type_name -> robot.user.USER_TYPE
	2,  // 1: robot.user.User.spec:type_name -> robot.user.CreateUserRequest
	3,  // 2: robot.user.User.profile:type_name -> robot.user.Profile
	4,  // 3: robot.user.UserSet.items:type_name -> robot.user.User
	1,  // 4: robot.user.DescribeUserRequest.describe_by:type_name -> robot.user.DESCRIBE_BY
	11, // 5: robot.user.QueryUserRequest.page:type_name -> infraboard.mcube.page.PageRequest
	12, // 6: robot.user.UpdateUserRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	3,  // 7: robot.user.UpdateUserRequest.profile:type_name -> robot.user.Profile
	2,  // 8: robot.user.RPC.CreateUser:input_type -> robot.user.CreateUserRequest
	6,  // 9: robot.user.RPC.DescribeUser:input_type -> robot.user.DescribeUserRequest
	7,  // 10: robot.user.RPC.QueryUser:input_type -> robot.user.QueryUserRequest
	8,  // 11: robot.user.RPC.ValidateUser:input_type -> robot.user.ValidateRequest
	10, // 12: robot.user.RPC.DeleteUser:input_type -> robot.user.DeleteUserRequest
	9,  // 13: robot.user.RPC.UpdateUser:input_type -> robot.user.UpdateUserRequest
	4,  // 14: robot.user.RPC.CreateUser:output_type -> robot.user.User
	4,  // 15: robot.user.RPC.DescribeUser:output_type -> robot.user.User
	5,  // 16: robot.user.RPC.QueryUser:output_type -> robot.user.UserSet
	4,  // 17: robot.user.RPC.ValidateUser:output_type -> robot.user.User
	5,  // 18: robot.user.RPC.DeleteUser:output_type -> robot.user.UserSet
	4,  // 19: robot.user.RPC.UpdateUser:output_type -> robot.user.User
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_apps_user_pb_user_proto_init() }
func file_apps_user_pb_user_proto_init() {
	if File_apps_user_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_user_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_apps_user_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_apps_user_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSet); i {
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
		file_apps_user_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUserRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_apps_user_pb_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
			RawDescriptor: file_apps_user_pb_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_user_pb_user_proto_goTypes,
		DependencyIndexes: file_apps_user_pb_user_proto_depIdxs,
		EnumInfos:         file_apps_user_pb_user_proto_enumTypes,
		MessageInfos:      file_apps_user_pb_user_proto_msgTypes,
	}.Build()
	File_apps_user_pb_user_proto = out.File
	file_apps_user_pb_user_proto_rawDesc = nil
	file_apps_user_pb_user_proto_goTypes = nil
	file_apps_user_pb_user_proto_depIdxs = nil
}
