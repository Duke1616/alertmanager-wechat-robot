// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/policy/pb/rpc.proto

package policy

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

type DESCRIBE_BY int32

const (
	// 通过policy_ids id
	DESCRIBE_BY_id DESCRIBE_BY = 0
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "id",
	}
	DESCRIBE_BY_value = map[string]int32{
		"id": 0,
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
	return file_apps_policy_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_apps_policy_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type DescribePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询详情的方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=robot.policy.DESCRIBE_BY" json:"describe_by"`
	// 服务客户端Id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id"`
	// 策略类型
	// @gotags: json:"policy_type"
	PolicyType string `protobuf:"bytes,3,opt,name=policy_type,json=policyType,proto3" json:"policy_type"`
}

func (x *DescribePolicyRequest) Reset() {
	*x = DescribePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePolicyRequest) ProtoMessage() {}

func (x *DescribePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePolicyRequest.ProtoReflect.Descriptor instead.
func (*DescribePolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DescribePolicyRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_id
}

func (x *DescribePolicyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescribePolicyRequest) GetPolicyType() string {
	if x != nil {
		return x.PolicyType
	}
	return ""
}

type QueryPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页参数
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// target_id
	// @gotags: json:"target_name"
	TargetName string `protobuf:"bytes,2,opt,name=target_name,json=targetName,proto3" json:"target_name"`
	// policy_ids
	// @gotags: json:"policy_ids"
	PolicyIds []string `protobuf:"bytes,4,rep,name=policy_ids,json=policyIds,proto3" json:"policy_ids"`
	// 策略类型
	// @gotags: json:"policy_type"
	PolicyType string `protobuf:"bytes,5,opt,name=policy_type,json=policyType,proto3" json:"policy_type"`
	// 策略类型
	// @gotags: json:"create_type"
	CreateType string `protobuf:"bytes,6,opt,name=create_type,json=createType,proto3" json:"create_type"`
	// 排序方式
	// @gotags: json:"sort"
	Sort string `protobuf:"bytes,7,opt,name=sort,proto3" json:"sort"`
	// 快速检索
	// @gotags: json:"filter_name"
	FilterName string `protobuf:"bytes,8,opt,name=filter_name,json=filterName,proto3" json:"filter_name"`
}

func (x *QueryPolicyRequest) Reset() {
	*x = QueryPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPolicyRequest) ProtoMessage() {}

func (x *QueryPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPolicyRequest.ProtoReflect.Descriptor instead.
func (*QueryPolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *QueryPolicyRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPolicyRequest) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *QueryPolicyRequest) GetPolicyIds() []string {
	if x != nil {
		return x.PolicyIds
	}
	return nil
}

func (x *QueryPolicyRequest) GetPolicyType() string {
	if x != nil {
		return x.PolicyType
	}
	return ""
}

func (x *QueryPolicyRequest) GetCreateType() string {
	if x != nil {
		return x.CreateType
	}
	return ""
}

func (x *QueryPolicyRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *QueryPolicyRequest) GetFilterName() string {
	if x != nil {
		return x.FilterName
	}
	return ""
}

type UpdatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,1,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 群组id
	// @gotags: json:"target_id" validate:"required,lte=64"
	TargetId string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id" validate:"required,lte=64"`
	// 群组名称
	// @gotags: json:"name" validate:"required,lte=64"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name" validate:"required,lte=64"`
	// wechat url
	// @gotags: json:"url" validate:"required,lte=64"
	Url string `protobuf:"bytes,4,opt,name=url,proto3" json:"url" validate:"required,lte=64"`
	// robot 密钥信息
	// @gotags: json:"secret" validate:"required,lte=64"
	Secret string `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret" validate:"required,lte=64"`
	// 描述信息
	// @gotags: json:"description"
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description"`
}

func (x *UpdatePolicyRequest) Reset() {
	*x = UpdatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyRequest) ProtoMessage() {}

func (x *UpdatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePolicyRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdatePolicyRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *UpdatePolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePolicyRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdatePolicyRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *UpdatePolicyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeletePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户账号id
	// @gotags: json:"policy_ids" validate:"required,lte=60"
	PolicyIds []string `protobuf:"bytes,2,rep,name=policy_ids,json=policyIds,proto3" json:"policy_ids" validate:"required,lte=60"`
}

func (x *DeletePolicyRequest) Reset() {
	*x = DeletePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_policy_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePolicyRequest) ProtoMessage() {}

func (x *DeletePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_policy_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePolicyRequest.ProtoReflect.Descriptor instead.
func (*DeletePolicyRequest) Descriptor() ([]byte, []int) {
	return file_apps_policy_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePolicyRequest) GetPolicyIds() []string {
	if x != nil {
		return x.PolicyIds
	}
	return nil
}

var File_apps_policy_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_policy_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x84, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd9, 0x01, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x73, 0x2a, 0x15,
	0x0a, 0x0b, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x06, 0x0a,
	0x02, 0x69, 0x64, 0x10, 0x00, 0x32, 0xfa, 0x02, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x47, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x4b, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x23, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x48, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x20, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x4a, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x2e,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x21, 0x2e, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x75, 0x6b, 0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_policy_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_policy_pb_rpc_proto_rawDescData = file_apps_policy_pb_rpc_proto_rawDesc
)

func file_apps_policy_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_policy_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_policy_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_policy_pb_rpc_proto_rawDescData)
	})
	return file_apps_policy_pb_rpc_proto_rawDescData
}

var file_apps_policy_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_policy_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_policy_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),              // 0: robot.policy.DESCRIBE_BY
	(*DescribePolicyRequest)(nil), // 1: robot.policy.DescribePolicyRequest
	(*QueryPolicyRequest)(nil),    // 2: robot.policy.QueryPolicyRequest
	(*UpdatePolicyRequest)(nil),   // 3: robot.policy.UpdatePolicyRequest
	(*DeletePolicyRequest)(nil),   // 4: robot.policy.DeletePolicyRequest
	(*request.PageRequest)(nil),   // 5: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),      // 6: infraboard.mcube.request.UpdateMode
	(*CreatePolicyRequest)(nil),   // 7: robot.policy.CreatePolicyRequest
	(*Policy)(nil),                // 8: robot.policy.Policy
	(*PolicySet)(nil),             // 9: robot.policy.PolicySet
}
var file_apps_policy_pb_rpc_proto_depIdxs = []int32{
	0, // 0: robot.policy.DescribePolicyRequest.describe_by:type_name -> robot.policy.DESCRIBE_BY
	5, // 1: robot.policy.QueryPolicyRequest.page:type_name -> infraboard.mcube.page.PageRequest
	6, // 2: robot.policy.UpdatePolicyRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	7, // 3: robot.policy.RPC.CreatePolicy:input_type -> robot.policy.CreatePolicyRequest
	1, // 4: robot.policy.RPC.DescribePolicy:input_type -> robot.policy.DescribePolicyRequest
	2, // 5: robot.policy.RPC.QueryPolicy:input_type -> robot.policy.QueryPolicyRequest
	4, // 6: robot.policy.RPC.DeletePolicy:input_type -> robot.policy.DeletePolicyRequest
	3, // 7: robot.policy.RPC.UpdatePolicy:input_type -> robot.policy.UpdatePolicyRequest
	8, // 8: robot.policy.RPC.CreatePolicy:output_type -> robot.policy.Policy
	8, // 9: robot.policy.RPC.DescribePolicy:output_type -> robot.policy.Policy
	9, // 10: robot.policy.RPC.QueryPolicy:output_type -> robot.policy.PolicySet
	9, // 11: robot.policy.RPC.DeletePolicy:output_type -> robot.policy.PolicySet
	8, // 12: robot.policy.RPC.UpdatePolicy:output_type -> robot.policy.Policy
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_policy_pb_rpc_proto_init() }
func file_apps_policy_pb_rpc_proto_init() {
	if File_apps_policy_pb_rpc_proto != nil {
		return
	}
	file_apps_policy_pb_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_policy_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePolicyRequest); i {
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
		file_apps_policy_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPolicyRequest); i {
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
		file_apps_policy_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePolicyRequest); i {
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
		file_apps_policy_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePolicyRequest); i {
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
			RawDescriptor: file_apps_policy_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_policy_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_policy_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_policy_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_policy_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_policy_pb_rpc_proto = out.File
	file_apps_policy_pb_rpc_proto_rawDesc = nil
	file_apps_policy_pb_rpc_proto_goTypes = nil
	file_apps_policy_pb_rpc_proto_depIdxs = nil
}
