// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.6.1
// source: subscriber-auth.proto

package pb

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

type ManagePublicationSubscriptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PublicationId int32                  `protobuf:"varint,2,opt,name=publication_id,json=publicationId,proto3" json:"publication_id,omitempty"`
	PlanId        int32                  `protobuf:"varint,3,opt,name=planId,proto3" json:"planId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManagePublicationSubscriptionRequest) Reset() {
	*x = ManagePublicationSubscriptionRequest{}
	mi := &file_subscriber_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManagePublicationSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagePublicationSubscriptionRequest) ProtoMessage() {}

func (x *ManagePublicationSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagePublicationSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ManagePublicationSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ManagePublicationSubscriptionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ManagePublicationSubscriptionRequest) GetPublicationId() int32 {
	if x != nil {
		return x.PublicationId
	}
	return 0
}

func (x *ManagePublicationSubscriptionRequest) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

type PublisherSubscriptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PlanId        int32                  `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublisherSubscriptionRequest) Reset() {
	*x = PublisherSubscriptionRequest{}
	mi := &file_subscriber_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublisherSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherSubscriptionRequest) ProtoMessage() {}

func (x *PublisherSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*PublisherSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_auth_proto_rawDescGZIP(), []int{1}
}

func (x *PublisherSubscriptionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PublisherSubscriptionRequest) GetPlanId() int32 {
	if x != nil {
		return x.PlanId
	}
	return 0
}

type PublisherUnsubscriptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublisherUnsubscriptionRequest) Reset() {
	*x = PublisherUnsubscriptionRequest{}
	mi := &file_subscriber_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublisherUnsubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherUnsubscriptionRequest) ProtoMessage() {}

func (x *PublisherUnsubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherUnsubscriptionRequest.ProtoReflect.Descriptor instead.
func (*PublisherUnsubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_auth_proto_rawDescGZIP(), []int{2}
}

func (x *PublisherUnsubscriptionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ChangePublisherPlanRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChangedPlanId int32                  `protobuf:"varint,2,opt,name=changed_plan_id,json=changedPlanId,proto3" json:"changed_plan_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePublisherPlanRequest) Reset() {
	*x = ChangePublisherPlanRequest{}
	mi := &file_subscriber_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePublisherPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePublisherPlanRequest) ProtoMessage() {}

func (x *ChangePublisherPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePublisherPlanRequest.ProtoReflect.Descriptor instead.
func (*ChangePublisherPlanRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ChangePublisherPlanRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChangePublisherPlanRequest) GetChangedPlanId() int32 {
	if x != nil {
		return x.ChangedPlanId
	}
	return 0
}

type ChangePublicationSubscriptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChangedPlanId int32                  `protobuf:"varint,2,opt,name=changed_plan_id,json=changedPlanId,proto3" json:"changed_plan_id,omitempty"`
	PublicationId int32                  `protobuf:"varint,3,opt,name=publication_id,json=publicationId,proto3" json:"publication_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePublicationSubscriptionRequest) Reset() {
	*x = ChangePublicationSubscriptionRequest{}
	mi := &file_subscriber_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePublicationSubscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePublicationSubscriptionRequest) ProtoMessage() {}

func (x *ChangePublicationSubscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePublicationSubscriptionRequest.ProtoReflect.Descriptor instead.
func (*ChangePublicationSubscriptionRequest) Descriptor() ([]byte, []int) {
	return file_subscriber_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePublicationSubscriptionRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChangePublicationSubscriptionRequest) GetChangedPlanId() int32 {
	if x != nil {
		return x.ChangedPlanId
	}
	return 0
}

func (x *ChangePublicationSubscriptionRequest) GetPublicationId() int32 {
	if x != nil {
		return x.PublicationId
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccess     bool                   `protobuf:"varint,1,opt,name=IsSuccess,proto3" json:"IsSuccess,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_subscriber_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_subscriber_auth_proto_msgTypes[5]
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
	return file_subscriber_auth_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_subscriber_auth_proto protoreflect.FileDescriptor

var file_subscriber_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2d, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x24, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x1c, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x1e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x24, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc7,
	0x03, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1d, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x18, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1f,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x1d, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x1f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e,
	0x12, 0x1b, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscriber_auth_proto_rawDescOnce sync.Once
	file_subscriber_auth_proto_rawDescData = file_subscriber_auth_proto_rawDesc
)

func file_subscriber_auth_proto_rawDescGZIP() []byte {
	file_subscriber_auth_proto_rawDescOnce.Do(func() {
		file_subscriber_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscriber_auth_proto_rawDescData)
	})
	return file_subscriber_auth_proto_rawDescData
}

var file_subscriber_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_subscriber_auth_proto_goTypes = []any{
	(*ManagePublicationSubscriptionRequest)(nil), // 0: ManagePublicationSubscriptionRequest
	(*PublisherSubscriptionRequest)(nil),         // 1: PublisherSubscriptionRequest
	(*PublisherUnsubscriptionRequest)(nil),       // 2: PublisherUnsubscriptionRequest
	(*ChangePublisherPlanRequest)(nil),           // 3: ChangePublisherPlanRequest
	(*ChangePublicationSubscriptionRequest)(nil), // 4: ChangePublicationSubscriptionRequest
	(*Response)(nil),                             // 5: Response
}
var file_subscriber_auth_proto_depIdxs = []int32{
	0, // 0: SubscriptionAuth.JoinPublication:input_type -> ManagePublicationSubscriptionRequest
	0, // 1: SubscriptionAuth.LeavePublication:input_type -> ManagePublicationSubscriptionRequest
	1, // 2: SubscriptionAuth.SubscribePublisherPlan:input_type -> PublisherSubscriptionRequest
	2, // 3: SubscriptionAuth.UnsubscribePublisherPlan:input_type -> PublisherUnsubscriptionRequest
	4, // 4: SubscriptionAuth.ChangePublicationSubscription:input_type -> ChangePublicationSubscriptionRequest
	3, // 5: SubscriptionAuth.ChangePublisherSubscriptionPlan:input_type -> ChangePublisherPlanRequest
	5, // 6: SubscriptionAuth.JoinPublication:output_type -> Response
	5, // 7: SubscriptionAuth.LeavePublication:output_type -> Response
	5, // 8: SubscriptionAuth.SubscribePublisherPlan:output_type -> Response
	5, // 9: SubscriptionAuth.UnsubscribePublisherPlan:output_type -> Response
	5, // 10: SubscriptionAuth.ChangePublicationSubscription:output_type -> Response
	5, // 11: SubscriptionAuth.ChangePublisherSubscriptionPlan:output_type -> Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_subscriber_auth_proto_init() }
func file_subscriber_auth_proto_init() {
	if File_subscriber_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subscriber_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subscriber_auth_proto_goTypes,
		DependencyIndexes: file_subscriber_auth_proto_depIdxs,
		MessageInfos:      file_subscriber_auth_proto_msgTypes,
	}.Build()
	File_subscriber_auth_proto = out.File
	file_subscriber_auth_proto_rawDesc = nil
	file_subscriber_auth_proto_goTypes = nil
	file_subscriber_auth_proto_depIdxs = nil
}
