// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/rule/pb/group.proto

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

type Rules struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态
	// @gotags: bson:"status" json:"status"
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status" bson:"status"`
	// 数据
	// @gotags: bson:"data" json:"data"
	Data *Groups `protobuf:"bytes,2,opt,name=data,proto3" json:"data" bson:"data"`
}

func (x *Rules) Reset() {
	*x = Rules{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rules) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rules) ProtoMessage() {}

func (x *Rules) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rules.ProtoReflect.Descriptor instead.
func (*Rules) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_group_proto_rawDescGZIP(), []int{0}
}

func (x *Rules) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Rules) GetData() *Groups {
	if x != nil {
		return x.Data
	}
	return nil
}

type Groups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组ID
	// @gotags: bson:"groups" json:"groups"
	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups" bson:"groups"`
}

func (x *Groups) Reset() {
	*x = Groups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Groups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Groups) ProtoMessage() {}

func (x *Groups) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Groups.ProtoReflect.Descriptor instead.
func (*Groups) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_group_proto_rawDescGZIP(), []int{1}
}

func (x *Groups) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 组ID
	// @gotags: bson:"_id" json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id" validate:"required,lte=64"`
	// 组名称
	// @gotags: bson:"name" json:"name" validate:"required,lte=64"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=64"`
	// 组所在域
	// @gotags: bson:"domain" json:"domain" validate:"required,lte=64"
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain" bson:"domain" validate:"required,lte=64"`
	// 组所在名称空间
	// @gotags: bson:"namespace" json:"namespace" validate:"required,lte=64"
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace" bson:"namespace" validate:"required,lte=64"`
	// 组报警间隔时间
	// @gotags: bson:"interval" json:"interval"
	Interval uint64 `protobuf:"varint,6,opt,name=interval,proto3" json:"interval" bson:"interval"`
	// 组内报警规则
	// @gotags: bson:"-" json:"rules"
	Rules []*Rule `protobuf:"bytes,7,rep,name=rules,proto3" json:"rules" bson:"-"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_group_proto_rawDescGZIP(), []int{2}
}

func (x *Group) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Group) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Group) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Group) GetInterval() uint64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *Group) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type GroupSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数量
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 数据
	// @gotags: bson:"items" json:"items"
	Items []*Group `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *GroupSet) Reset() {
	*x = GroupSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rule_pb_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupSet) ProtoMessage() {}

func (x *GroupSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rule_pb_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupSet.ProtoReflect.Descriptor instead.
func (*GroupSet) Descriptor() ([]byte, []int) {
	return file_apps_rule_pb_group_proto_rawDescGZIP(), []int{3}
}

func (x *GroupSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GroupSet) GetItems() []*Group {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_apps_rule_pb_group_proto protoreflect.FileDescriptor

var file_apps_rule_pb_group_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x1a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x05, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0xa5, 0x01,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x26, 0x0a,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x08, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44,
	0x75, 0x6b, 0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_apps_rule_pb_group_proto_rawDescOnce sync.Once
	file_apps_rule_pb_group_proto_rawDescData = file_apps_rule_pb_group_proto_rawDesc
)

func file_apps_rule_pb_group_proto_rawDescGZIP() []byte {
	file_apps_rule_pb_group_proto_rawDescOnce.Do(func() {
		file_apps_rule_pb_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_rule_pb_group_proto_rawDescData)
	})
	return file_apps_rule_pb_group_proto_rawDescData
}

var file_apps_rule_pb_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_rule_pb_group_proto_goTypes = []interface{}{
	(*Rules)(nil),    // 0: robot.rule.Rules
	(*Groups)(nil),   // 1: robot.rule.Groups
	(*Group)(nil),    // 2: robot.rule.Group
	(*GroupSet)(nil), // 3: robot.rule.GroupSet
	(*Rule)(nil),     // 4: robot.rule.Rule
}
var file_apps_rule_pb_group_proto_depIdxs = []int32{
	1, // 0: robot.rule.Rules.data:type_name -> robot.rule.Groups
	2, // 1: robot.rule.Groups.groups:type_name -> robot.rule.Group
	4, // 2: robot.rule.Group.rules:type_name -> robot.rule.Rule
	2, // 3: robot.rule.GroupSet.items:type_name -> robot.rule.Group
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_rule_pb_group_proto_init() }
func file_apps_rule_pb_group_proto_init() {
	if File_apps_rule_pb_group_proto != nil {
		return
	}
	file_apps_rule_pb_rule_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_rule_pb_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rules); i {
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
		file_apps_rule_pb_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Groups); i {
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
		file_apps_rule_pb_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_apps_rule_pb_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupSet); i {
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
			RawDescriptor: file_apps_rule_pb_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_rule_pb_group_proto_goTypes,
		DependencyIndexes: file_apps_rule_pb_group_proto_depIdxs,
		MessageInfos:      file_apps_rule_pb_group_proto_msgTypes,
	}.Build()
	File_apps_rule_pb_group_proto = out.File
	file_apps_rule_pb_group_proto_rawDesc = nil
	file_apps_rule_pb_group_proto_goTypes = nil
	file_apps_rule_pb_group_proto_depIdxs = nil
}