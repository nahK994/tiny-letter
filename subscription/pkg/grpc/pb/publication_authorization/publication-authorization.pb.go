// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.21.12
// source: publication-authorization.proto

package pb_publication_authorization

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

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsAuthorized  bool                   `protobuf:"varint,1,opt,name=is_authorized,json=isAuthorized,proto3" json:"is_authorized,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_publication_authorization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_publication_authorization_proto_msgTypes[0]
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
	return file_publication_authorization_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetIsAuthorized() bool {
	if x != nil {
		return x.IsAuthorized
	}
	return false
}

type IsAuthorizedPublisherRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PublicationId int32                  `protobuf:"varint,1,opt,name=publication_id,json=publicationId,proto3" json:"publication_id,omitempty"`
	UserId        int32                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsAuthorizedPublisherRequest) Reset() {
	*x = IsAuthorizedPublisherRequest{}
	mi := &file_publication_authorization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsAuthorizedPublisherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAuthorizedPublisherRequest) ProtoMessage() {}

func (x *IsAuthorizedPublisherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_publication_authorization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAuthorizedPublisherRequest.ProtoReflect.Descriptor instead.
func (*IsAuthorizedPublisherRequest) Descriptor() ([]byte, []int) {
	return file_publication_authorization_proto_rawDescGZIP(), []int{1}
}

func (x *IsAuthorizedPublisherRequest) GetPublicationId() int32 {
	if x != nil {
		return x.PublicationId
	}
	return 0
}

func (x *IsAuthorizedPublisherRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_publication_authorization_proto protoreflect.FileDescriptor

var file_publication_authorization_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x22, 0x5e, 0x0a, 0x1c, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x32, 0x5d, 0x0a, 0x18, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41,
	0x0a, 0x15, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x49, 0x73, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3a, 0x5a, 0x38, 0x2e, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x3b, 0x70, 0x62, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_publication_authorization_proto_rawDescOnce sync.Once
	file_publication_authorization_proto_rawDescData = file_publication_authorization_proto_rawDesc
)

func file_publication_authorization_proto_rawDescGZIP() []byte {
	file_publication_authorization_proto_rawDescOnce.Do(func() {
		file_publication_authorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_publication_authorization_proto_rawDescData)
	})
	return file_publication_authorization_proto_rawDescData
}

var file_publication_authorization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_publication_authorization_proto_goTypes = []any{
	(*Response)(nil),                     // 0: Response
	(*IsAuthorizedPublisherRequest)(nil), // 1: IsAuthorizedPublisherRequest
}
var file_publication_authorization_proto_depIdxs = []int32{
	1, // 0: PublicationAuthorization.IsAuthorizedPublisher:input_type -> IsAuthorizedPublisherRequest
	0, // 1: PublicationAuthorization.IsAuthorizedPublisher:output_type -> Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_publication_authorization_proto_init() }
func file_publication_authorization_proto_init() {
	if File_publication_authorization_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_publication_authorization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_publication_authorization_proto_goTypes,
		DependencyIndexes: file_publication_authorization_proto_depIdxs,
		MessageInfos:      file_publication_authorization_proto_msgTypes,
	}.Build()
	File_publication_authorization_proto = out.File
	file_publication_authorization_proto_rawDesc = nil
	file_publication_authorization_proto_goTypes = nil
	file_publication_authorization_proto_depIdxs = nil
}
