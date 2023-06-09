// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: apps/notify/pb/notify.proto

package notify

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

type NOTIFY_TYPE int32

const (
	NOTIFY_TYPE_WECHAT NOTIFY_TYPE = 0
	NOTIFY_TYPE_EMAIL  NOTIFY_TYPE = 1
	NOTIFY_TYPE_PHONE  NOTIFY_TYPE = 2
)

// Enum value maps for NOTIFY_TYPE.
var (
	NOTIFY_TYPE_name = map[int32]string{
		0: "WECHAT",
		1: "EMAIL",
		2: "PHONE",
	}
	NOTIFY_TYPE_value = map[string]int32{
		"WECHAT": 0,
		"EMAIL":  1,
		"PHONE":  2,
	}
)

func (x NOTIFY_TYPE) Enum() *NOTIFY_TYPE {
	p := new(NOTIFY_TYPE)
	*p = x
	return p
}

func (x NOTIFY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NOTIFY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_notify_pb_notify_proto_enumTypes[0].Descriptor()
}

func (NOTIFY_TYPE) Type() protoreflect.EnumType {
	return &file_apps_notify_pb_notify_proto_enumTypes[0]
}

func (x NOTIFY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NOTIFY_TYPE.Descriptor instead.
func (NOTIFY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{0}
}

type SendNotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 发送类型
	// @gotags: bson:"notify_type" json:"notify_type"
	NotifyType NOTIFY_TYPE `protobuf:"varint,1,opt,name=notify_type,json=notifyType,proto3,enum=robot.notify.NOTIFY_TYPE" json:"notify_type" bson:"notify_type"`
	// 消息标题
	// @gotags: json:"title" bson:"title" validate:"required"
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title" bson:"title" validate:"required"`
	// 用户
	// @gotags: json:"users" bson:"users"
	Users []string `protobuf:"bytes,3,rep,name=users,proto3" json:"users" bson:"users"`
	// 目标ID
	// @gotags: bson:"target_id" json:"target_id"
	TargetId string `protobuf:"bytes,4,opt,name=target_id,json=targetId,proto3" json:"target_id" bson:"target_id"`
	// 消息格式
	// @gotags: bson:"content_type" json:"content_type"
	ContentType string `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type" bson:"content_type"`
	// 消息内容
	// @gotags: json:"content" bson:"content"
	Content string `protobuf:"bytes,6,opt,name=content,proto3" json:"content" bson:"content"`
	// 企业微信格式
	// @gotags: bson:"wechat_request" json:"wechat_request"
	WechatRequest *WechatRequest `protobuf:"bytes,7,opt,name=wechat_request,json=wechatRequest,proto3" json:"wechat_request" bson:"wechat_request"`
}

func (x *SendNotifyRequest) Reset() {
	*x = SendNotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotifyRequest) ProtoMessage() {}

func (x *SendNotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotifyRequest.ProtoReflect.Descriptor instead.
func (*SendNotifyRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{0}
}

func (x *SendNotifyRequest) GetNotifyType() NOTIFY_TYPE {
	if x != nil {
		return x.NotifyType
	}
	return NOTIFY_TYPE_WECHAT
}

func (x *SendNotifyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendNotifyRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SendNotifyRequest) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *SendNotifyRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SendNotifyRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendNotifyRequest) GetWechatRequest() *WechatRequest {
	if x != nil {
		return x.WechatRequest
	}
	return nil
}

type WechatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否@所有人
	// @gotags: bson:"url" json:"url"
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url" bson:"url"`
	// 是否@所有人
	// @gotags: bson:"secret" json:"secret"
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret" bson:"secret"`
}

func (x *WechatRequest) Reset() {
	*x = WechatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WechatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WechatRequest) ProtoMessage() {}

func (x *WechatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WechatRequest.ProtoReflect.Descriptor instead.
func (*WechatRequest) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{1}
}

func (x *WechatRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WechatRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_notify_pb_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apps_notify_pb_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_apps_notify_pb_notify_proto_rawDescGZIP(), []int{2}
}

func (x *SendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_apps_notify_pb_notify_proto protoreflect.FileDescriptor

var file_apps_notify_pb_notify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x62,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x99, 0x02, 0x0a, 0x11,
	0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x0d, 0x57, 0x65, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x2f, 0x0a, 0x0b,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a, 0x06, 0x57,
	0x45, 0x43, 0x48, 0x41, 0x54, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x32, 0x50, 0x0a,
	0x03, 0x52, 0x50, 0x43, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x1f, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x75,
	0x6b, 0x65, 0x31, 0x36, 0x31, 0x36, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2d, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2d, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_notify_pb_notify_proto_rawDescOnce sync.Once
	file_apps_notify_pb_notify_proto_rawDescData = file_apps_notify_pb_notify_proto_rawDesc
)

func file_apps_notify_pb_notify_proto_rawDescGZIP() []byte {
	file_apps_notify_pb_notify_proto_rawDescOnce.Do(func() {
		file_apps_notify_pb_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_notify_pb_notify_proto_rawDescData)
	})
	return file_apps_notify_pb_notify_proto_rawDescData
}

var file_apps_notify_pb_notify_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_notify_pb_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_notify_pb_notify_proto_goTypes = []interface{}{
	(NOTIFY_TYPE)(0),          // 0: robot.notify.NOTIFY_TYPE
	(*SendNotifyRequest)(nil), // 1: robot.notify.SendNotifyRequest
	(*WechatRequest)(nil),     // 2: robot.notify.WechatRequest
	(*SendResponse)(nil),      // 3: robot.notify.SendResponse
}
var file_apps_notify_pb_notify_proto_depIdxs = []int32{
	0, // 0: robot.notify.SendNotifyRequest.notify_type:type_name -> robot.notify.NOTIFY_TYPE
	2, // 1: robot.notify.SendNotifyRequest.wechat_request:type_name -> robot.notify.WechatRequest
	1, // 2: robot.notify.RPC.SendNotify:input_type -> robot.notify.SendNotifyRequest
	3, // 3: robot.notify.RPC.SendNotify:output_type -> robot.notify.SendResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_notify_pb_notify_proto_init() }
func file_apps_notify_pb_notify_proto_init() {
	if File_apps_notify_pb_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_notify_pb_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotifyRequest); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WechatRequest); i {
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
		file_apps_notify_pb_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
			RawDescriptor: file_apps_notify_pb_notify_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_notify_pb_notify_proto_goTypes,
		DependencyIndexes: file_apps_notify_pb_notify_proto_depIdxs,
		EnumInfos:         file_apps_notify_pb_notify_proto_enumTypes,
		MessageInfos:      file_apps_notify_pb_notify_proto_msgTypes,
	}.Build()
	File_apps_notify_pb_notify_proto = out.File
	file_apps_notify_pb_notify_proto_rawDesc = nil
	file_apps_notify_pb_notify_proto_goTypes = nil
	file_apps_notify_pb_notify_proto_depIdxs = nil
}
