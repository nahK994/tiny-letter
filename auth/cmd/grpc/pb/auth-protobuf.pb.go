// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.21.12
// source: auth-protobuf.proto

package pb_auth

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

type PublisherActionType int32

const (
	PublisherActionType_PUBLISHER_SUBSCRIBE   PublisherActionType = 0
	PublisherActionType_PUBLISHER_UNSUBSCRIBE PublisherActionType = 1
	PublisherActionType_PUBLISHER_CHANGE_PLAN PublisherActionType = 2
)

// Enum value maps for PublisherActionType.
var (
	PublisherActionType_name = map[int32]string{
		0: "PUBLISHER_SUBSCRIBE",
		1: "PUBLISHER_UNSUBSCRIBE",
		2: "PUBLISHER_CHANGE_PLAN",
	}
	PublisherActionType_value = map[string]int32{
		"PUBLISHER_SUBSCRIBE":   0,
		"PUBLISHER_UNSUBSCRIBE": 1,
		"PUBLISHER_CHANGE_PLAN": 2,
	}
)

func (x PublisherActionType) Enum() *PublisherActionType {
	p := new(PublisherActionType)
	*p = x
	return p
}

func (x PublisherActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PublisherActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_protobuf_proto_enumTypes[0].Descriptor()
}

func (PublisherActionType) Type() protoreflect.EnumType {
	return &file_auth_protobuf_proto_enumTypes[0]
}

func (x PublisherActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PublisherActionType.Descriptor instead.
func (PublisherActionType) EnumDescriptor() ([]byte, []int) {
	return file_auth_protobuf_proto_rawDescGZIP(), []int{0}
}

type SubscriberActionType int32

const (
	SubscriberActionType_SUBSCRIBER_JOIN        SubscriberActionType = 0
	SubscriberActionType_SUBSCRIBER_LEAVE       SubscriberActionType = 1
	SubscriberActionType_SUBSCRIBER_CHANGE_PLAN SubscriberActionType = 2
)

// Enum value maps for SubscriberActionType.
var (
	SubscriberActionType_name = map[int32]string{
		0: "SUBSCRIBER_JOIN",
		1: "SUBSCRIBER_LEAVE",
		2: "SUBSCRIBER_CHANGE_PLAN",
	}
	SubscriberActionType_value = map[string]int32{
		"SUBSCRIBER_JOIN":        0,
		"SUBSCRIBER_LEAVE":       1,
		"SUBSCRIBER_CHANGE_PLAN": 2,
	}
)

func (x SubscriberActionType) Enum() *SubscriberActionType {
	p := new(SubscriberActionType)
	*p = x
	return p
}

func (x SubscriberActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriberActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_protobuf_proto_enumTypes[1].Descriptor()
}

func (SubscriberActionType) Type() protoreflect.EnumType {
	return &file_auth_protobuf_proto_enumTypes[1]
}

func (x SubscriberActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriberActionType.Descriptor instead.
func (SubscriberActionType) EnumDescriptor() ([]byte, []int) {
	return file_auth_protobuf_proto_rawDescGZIP(), []int{1}
}

type PublisherActionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        PublisherActionType    `protobuf:"varint,1,opt,name=action,proto3,enum=PublisherActionType" json:"action,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlanId        int32                  `protobuf:"varint,3,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"` // Defaults to 0 if not set
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublisherActionRequest) Reset() {
	*x = PublisherActionRequest{}
	mi := &file_auth_protobuf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublisherActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherActionRequest) ProtoMessage() {}

func (x *PublisherActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_protobuf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherActionRequest.ProtoReflect.Descriptor instead.
func (*PublisherActionRequest) Descriptor() ([]byte, []int) {
	return file_auth_protobuf_proto_rawDescGZIP(), []int{0}
}

func (x *PublisherActionRequest) GetAction() PublisherActionType {
	if x != nil {
		return x.Action
	}
	return PublisherActionType_PUBLISHER_SUBSCRIBE
}

func (x *PublisherActionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublisherActionRequest) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

type SubscriberActionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        SubscriberActionType   `protobuf:"varint,1,opt,name=action,proto3,enum=SubscriberActionType" json:"action,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PublicationId int32                  `protobuf:"varint,3,opt,name=publication_id,json=publicationId,proto3" json:"publication_id,omitempty"`
	PlanId        int32                  `protobuf:"varint,4,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"` // Defaults to 0 if not set
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscriberActionRequest) Reset() {
	*x = SubscriberActionRequest{}
	mi := &file_auth_protobuf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriberActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberActionRequest) ProtoMessage() {}

func (x *SubscriberActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_protobuf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberActionRequest.ProtoReflect.Descriptor instead.
func (*SubscriberActionRequest) Descriptor() ([]byte, []int) {
	return file_auth_protobuf_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriberActionRequest) GetAction() SubscriberActionType {
	if x != nil {
		return x.Action
	}
	return SubscriberActionType_SUBSCRIBER_JOIN
}

func (x *SubscriberActionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SubscriberActionRequest) GetPublicationId() int32 {
	if x != nil {
		return x.PublicationId
	}
	return 0
}

func (x *SubscriberActionRequest) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_auth_protobuf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_auth_protobuf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_auth_protobuf_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_auth_protobuf_proto protoreflect.FileDescriptor

var file_auth_protobuf_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22,
	0xa1, 0x01, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x64,
	0x0a, 0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48,
	0x45, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x00, 0x12, 0x19,
	0x0a, 0x15, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x55,
	0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x55, 0x42,
	0x4c, 0x49, 0x53, 0x48, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x50, 0x4c,
	0x41, 0x4e, 0x10, 0x02, 0x2a, 0x5d, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f,
	0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x52, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x52, 0x5f,
	0x4c, 0x45, 0x41, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x42, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x50, 0x4c, 0x41,
	0x4e, 0x10, 0x02, 0x32, 0x85, 0x01, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0f, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e,
	0x3b, 0x70, 0x62, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_protobuf_proto_rawDescOnce sync.Once
	file_auth_protobuf_proto_rawDescData = file_auth_protobuf_proto_rawDesc
)

func file_auth_protobuf_proto_rawDescGZIP() []byte {
	file_auth_protobuf_proto_rawDescOnce.Do(func() {
		file_auth_protobuf_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_protobuf_proto_rawDescData)
	})
	return file_auth_protobuf_proto_rawDescData
}

var file_auth_protobuf_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auth_protobuf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auth_protobuf_proto_goTypes = []any{
	(PublisherActionType)(0),        // 0: PublisherActionType
	(SubscriberActionType)(0),       // 1: SubscriberActionType
	(*PublisherActionRequest)(nil),  // 2: PublisherActionRequest
	(*SubscriberActionRequest)(nil), // 3: SubscriberActionRequest
	(*Response)(nil),                // 4: Response
}
var file_auth_protobuf_proto_depIdxs = []int32{
	0, // 0: PublisherActionRequest.action:type_name -> PublisherActionType
	1, // 1: SubscriberActionRequest.action:type_name -> SubscriberActionType
	2, // 2: SubscriptionService.PublisherAction:input_type -> PublisherActionRequest
	3, // 3: SubscriptionService.SubscriberAction:input_type -> SubscriberActionRequest
	4, // 4: SubscriptionService.PublisherAction:output_type -> Response
	4, // 5: SubscriptionService.SubscriberAction:output_type -> Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_protobuf_proto_init() }
func file_auth_protobuf_proto_init() {
	if File_auth_protobuf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_protobuf_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_protobuf_proto_goTypes,
		DependencyIndexes: file_auth_protobuf_proto_depIdxs,
		EnumInfos:         file_auth_protobuf_proto_enumTypes,
		MessageInfos:      file_auth_protobuf_proto_msgTypes,
	}.Build()
	File_auth_protobuf_proto = out.File
	file_auth_protobuf_proto_rawDesc = nil
	file_auth_protobuf_proto_goTypes = nil
	file_auth_protobuf_proto_depIdxs = nil
}
