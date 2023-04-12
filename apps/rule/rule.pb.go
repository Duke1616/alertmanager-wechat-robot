// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/rule/pb/rule.proto

package rule

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

type ACTIVE int32

const (
	ACTIVE_NEWS ACTIVE = 0
	// 状态转换
	ACTIVE_STATUS_TRANS ACTIVE = 2
	// 丢弃报警
	ACTIVE_DROP ACTIVE = 1
)

// Enum value maps for ACTIVE.
var (
	ACTIVE_name = map[int32]string{
		0: "NEWS",
		2: "STATUS_TRANS",
		1: "DROP",
	}
	ACTIVE_value = map[string]int32{
		"NEWS":         0,
		"STATUS_TRANS": 2,
		"DROP":         1,
	}
)

func (x ACTIVE) Enum() *ACTIVE {
	p := new(ACTIVE)
	*p = x
	return p
}

func (x ACTIVE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACTIVE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_rule_pb_rule_proto_enumTypes[0].Descriptor()
}

func (ACTIVE) Type() protoreflect.EnumType {
	return &file_apps_rule_pb_rule_proto_enumTypes[0]
}

func (x ACTIVE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACTIVE.Descriptor instead.
func (ACTIVE) EnumDescriptor() ([]byte, []int) {
	return file_apps_rule_pb_rule_proto_rawDescGZIP(), []int{0}
}

type LABEL_TYPE int32

const (
	// 组策略标签匹配
	LABEL_TYPE_GROUP LABEL_TYPE = 0
	// 全局配置标签匹配
	LABEL_TYPE_COMMON LABEL_TYPE = 1
)

// Enum value maps for LABEL_TYPE.
var (
	LABEL_TYPE_name = map[int32]string{
		0: "GROUP",
		1: "COMMON",
	}
	LABEL_TYPE_value = map[string]int32{
		"GROUP":  0,
		"COMMON": 1,
	}
)

func (x LABEL_TYPE) Enum() *LABEL_TYPE {
	p := new(LABEL_TYPE)
	*p = x
	return p
}

func (x LABEL_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LABEL_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_rule_pb_rule_proto_enumTypes[1].Descriptor()
}

func (LABEL_TYPE) Type() protoreflect.EnumType {
	return &file_apps_rule_pb_rule_proto_enumTypes[1]
}

func (x LABEL_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LABEL_TYPE.Descriptor instead.
func (LABEL_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_rule_pb_rule_proto_rawDescGZIP(), []int{1}
}

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// rule ID
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 是否开启
	// @gotags: bson:"enabled" json:"enabled"
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled" bson:"enabled"`
	// 创建时间
	// @gotags: bson:"create_at" json:"create_at"
	CreateAt int64 `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at" bson:"create_at"`
	// 创建信息
	// @gotags: bson:"spec" json:"spec"
	Spec *CreateRuleRequest `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec" bson:"spec"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_rule_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Rule) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Rule) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Rule) GetSpec() *CreateRuleRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type RuleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 数据
	// @gotags: bson:"items" json:"items"
	Items []*Rule `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *RuleSet) Reset() {
	*x = RuleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleSet) ProtoMessage() {}

func (x *RuleSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleSet.ProtoReflect.Descriptor instead.
func (*RuleSet) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_rule_proto_rawDescGZIP(), []int{1}
}

func (x *RuleSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RuleSet) GetItems() []*Rule {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 关联target ID
	// @gotags: bson:"target_id" json:"target_id" validate:"required,lte=64"
	TargetId string `protobuf:"bytes,1,opt,name=target_id,json=targetId,proto3" json:"target_id" bson:"target_id" validate:"required,lte=64"`
	// 标签类型
	// @gotags: bson:"label_type" json:"label_type"
	LabelType LABEL_TYPE `protobuf:"varint,2,opt,name=label_type,json=labelType,proto3,enum=robot.rule.LABEL_TYPE" json:"label_type" bson:"label_type"`
	// 匹配标签
	// @gotags: bson:"label" json:"label" validate:"required,lte=64"
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label" bson:"label" validate:"required,lte=64"`
	// 匹配值
	// @gotags: bson:"value" json:"value"
	Value []string `protobuf:"bytes,4,rep,name=value,proto3" json:"value" bson:"value"`
	// 动作
	// @gotags: bson:"active" json:"active"
	Active ACTIVE `protobuf:"varint,5,opt,name=active,proto3,enum=robot.rule.ACTIVE" json:"active" bson:"active"`
	// @其他人
	// @gotags: bson:"mention" json:"mention"
	Mention *Mention `protobuf:"bytes,6,opt,name=mention,proto3" json:"mention" bson:"mention"`
}

func (x *CreateRuleRequest) Reset() {
	*x = CreateRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRuleRequest) ProtoMessage() {}

func (x *CreateRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRuleRequest.ProtoReflect.Descriptor instead.
func (*CreateRuleRequest) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_rule_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRuleRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *CreateRuleRequest) GetLabelType() LABEL_TYPE {
	if x != nil {
		return x.LabelType
	}
	return LABEL_TYPE_GROUP
}

func (x *CreateRuleRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateRuleRequest) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CreateRuleRequest) GetActive() ACTIVE {
	if x != nil {
		return x.Active
	}
	return ACTIVE_NEWS
}

func (x *CreateRuleRequest) GetMention() *Mention {
	if x != nil {
		return x.Mention
	}
	return nil
}

// 提到的人
type Mention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否@所有人
	// @gotags: bson:"all" json:"all"
	All bool `protobuf:"varint,1,opt,name=all,proto3" json:"all" bson:"all"`
	// @指定的人
	// @gotags: bson:"mobiles" json:"mobiles"
	Mobiles []string `protobuf:"bytes,2,rep,name=mobiles,proto3" json:"mobiles" bson:"mobiles"`
}

func (x *Mention) Reset() {
	*x = Mention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mention) ProtoMessage() {}

func (x *Mention) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mention.ProtoReflect.Descriptor instead.
func (*Mention) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_rule_proto_rawDescGZIP(), []int{3}
}

func (x *Mention) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

func (x *Mention) GetMobiles() []string {
	if x != nil {
		return x.Mobiles
	}
	return nil
}

var File_apps_rule_pb_rule_proto protoreflect.FileDescriptor

var file_apps_rule_pb_rule_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x47, 0x0a, 0x07, 0x52, 0x75, 0x6c, 0x65,
	0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xee, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x52, 0x09, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e,
	0x72, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c,
	0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6d, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x73, 0x2a, 0x2e, 0x0a, 0x06, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x45, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x57, 0x53, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50, 0x10, 0x01, 0x2a, 0x23, 0x0a, 0x0a, 0x4c, 0x41, 0x42,
	0x45, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x10, 0x01, 0x42, 0x39,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x75, 0x6b,
	0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_rule_pb_rule_proto_rawDescOnce sync.Once
	file_apps_rule_pb_rule_proto_rawDescData = file_apps_rule_pb_rule_proto_rawDesc
)

func file_apps_rule_pb_rule_proto_rawDescGZIP() []byte {
	file_apps_rule_pb_rule_proto_rawDescOnce.Do(func() {
		file_apps_rule_pb_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_rule_pb_rule_proto_rawDescData)
	})
	return file_apps_rule_pb_rule_proto_rawDescData
}

var file_apps_rule_pb_rule_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_rule_pb_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_rule_pb_rule_proto_goTypes = []interface{}{
	(ACTIVE)(0),               // 0: robot.rule.ACTIVE
	(LABEL_TYPE)(0),           // 1: robot.rule.LABEL_TYPE
	(*Rule)(nil),              // 2: robot.rule.Rule
	(*RuleSet)(nil),           // 3: robot.rule.RuleSet
	(*CreateRuleRequest)(nil), // 4: robot.rule.CreateRuleRequest
	(*Mention)(nil),           // 5: robot.rule.Mention
}
var file_apps_rule_pb_rule_proto_depIdxs = []int32{
	4, // 0: robot.rule.Rule.spec:type_name -> robot.rule.CreateRuleRequest
	2, // 1: robot.rule.RuleSet.items:type_name -> robot.rule.Rule
	1, // 2: robot.rule.CreateRuleRequest.label_type:type_name -> robot.rule.LABEL_TYPE
	0, // 3: robot.rule.CreateRuleRequest.active:type_name -> robot.rule.ACTIVE
	5, // 4: robot.rule.CreateRuleRequest.mention:type_name -> robot.rule.Mention
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_rule_pb_rule_proto_init() }
func file_apps_rule_pb_rule_proto_init() {
	if File_apps_rule_pb_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_rule_pb_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_apps_rule_pb_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleSet); i {
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
		file_apps_rule_pb_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRuleRequest); i {
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
		file_apps_rule_pb_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mention); i {
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
			RawDescriptor: file_apps_rule_pb_rule_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_rule_pb_rule_proto_goTypes,
		DependencyIndexes: file_apps_rule_pb_rule_proto_depIdxs,
		EnumInfos:         file_apps_rule_pb_rule_proto_enumTypes,
		MessageInfos:      file_apps_rule_pb_rule_proto_msgTypes,
	}.Build()
	File_apps_rule_pb_rule_proto = out.File
	file_apps_rule_pb_rule_proto_rawDesc = nil
	file_apps_rule_pb_rule_proto_goTypes = nil
	file_apps_rule_pb_rule_proto_depIdxs = nil
}