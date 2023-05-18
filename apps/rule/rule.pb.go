// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/rule/pb/rule.proto

package rule

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/protobuf/timestamp.proto.

type Timestamp = timestamppb.Timestamp

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 规则ID
	// @gotags: bson:"_id" json:"id" validate:"required,lte=64"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id" validate:"required,lte=64"`
	// 组id
	// @gotags: bson:"group_id" json:"group_id"
	GroupId string `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id" bson:"group_id"`
	// 组id
	// @gotags: bson:"group_name" json:"group_name"
	GroupName string `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name" bson:"group_name"`
	// 规则名称
	// @gotags: bson:"name" json:"name" validate:"required,lte=64"
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=64"`
	// 规则查询语句
	// @gotags: bson:"query" json:"query"
	Query string `protobuf:"bytes,5,opt,name=query,proto3" json:"query" bson:"query"`
	// 规则查询语句
	// @gotags: bson:"level" json:"level"
	Level string `protobuf:"bytes,6,opt,name=level,proto3" json:"level" bson:"level"`
	// 规则查询语句
	// @gotags: bson:"service" json:"service"
	Service string `protobuf:"bytes,7,opt,name=service,proto3" json:"service" bson:"service"`
	// 规则标签
	// @gotags: bson:"labels" json:"labels"
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
	// 规则注释
	// @gotags: bson:"annotations" json:"annotations"
	Annotations map[string]string `protobuf:"bytes,9,rep,name=annotations,proto3" json:"annotations" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"annotations"`
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

func (x *Rule) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Rule) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *Rule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rule) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Rule) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Rule) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Rule) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Rule) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
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

var File_apps_rule_pb_rule_proto protoreflect.FileDescriptor

var file_apps_rule_pb_rule_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72,
	0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2e, 0x72, 0x75, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x03, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65,
	0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x07, 0x52, 0x75, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x72, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x75, 0x6b, 0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x75, 0x6c, 0x65, 0x50, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_apps_rule_pb_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_rule_pb_rule_proto_goTypes = []interface{}{
	(*Rule)(nil),    // 0: robot.rule.Rule
	(*RuleSet)(nil), // 1: robot.rule.RuleSet
	nil,             // 2: robot.rule.Rule.LabelsEntry
	nil,             // 3: robot.rule.Rule.AnnotationsEntry
}
var file_apps_rule_pb_rule_proto_depIdxs = []int32{
	2, // 0: robot.rule.Rule.labels:type_name -> robot.rule.Rule.LabelsEntry
	3, // 1: robot.rule.Rule.annotations:type_name -> robot.rule.Rule.AnnotationsEntry
	0, // 2: robot.rule.RuleSet.items:type_name -> robot.rule.Rule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_rule_pb_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_rule_pb_rule_proto_goTypes,
		DependencyIndexes: file_apps_rule_pb_rule_proto_depIdxs,
		MessageInfos:      file_apps_rule_pb_rule_proto_msgTypes,
	}.Build()
	File_apps_rule_pb_rule_proto = out.File
	file_apps_rule_pb_rule_proto_rawDesc = nil
	file_apps_rule_pb_rule_proto_goTypes = nil
	file_apps_rule_pb_rule_proto_depIdxs = nil
}
