// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: user_interactions.proto

package user_interactions

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{0}
}

type UserInteraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	CompositionId   string `protobuf:"bytes,2,opt,name=compositionId,proto3" json:"compositionId,omitempty"`
	InteractionType string `protobuf:"bytes,3,opt,name=interactionType,proto3" json:"interactionType,omitempty"`
}

func (x *UserInteraction) Reset() {
	*x = UserInteraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInteraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInteraction) ProtoMessage() {}

func (x *UserInteraction) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInteraction.ProtoReflect.Descriptor instead.
func (*UserInteraction) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{1}
}

func (x *UserInteraction) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInteraction) GetCompositionId() string {
	if x != nil {
		return x.CompositionId
	}
	return ""
}

func (x *UserInteraction) GetInteractionType() string {
	if x != nil {
		return x.InteractionType
	}
	return ""
}

type UserInteractionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CompositionId   string `protobuf:"bytes,3,opt,name=compositionId,proto3" json:"compositionId,omitempty"`
	InteractionType string `protobuf:"bytes,4,opt,name=interactionType,proto3" json:"interactionType,omitempty"`
	CreatedAt       string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt       string `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *UserInteractionRes) Reset() {
	*x = UserInteractionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInteractionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInteractionRes) ProtoMessage() {}

func (x *UserInteractionRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInteractionRes.ProtoReflect.Descriptor instead.
func (*UserInteractionRes) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{2}
}

func (x *UserInteractionRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInteractionRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInteractionRes) GetCompositionId() string {
	if x != nil {
		return x.CompositionId
	}
	return ""
}

func (x *UserInteractionRes) GetInteractionType() string {
	if x != nil {
		return x.InteractionType
	}
	return ""
}

func (x *UserInteractionRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserInteractionRes) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserInteractionSRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInteraction []*UserInteractionRes `protobuf:"bytes,1,rep,name=userInteraction,proto3" json:"userInteraction,omitempty"`
}

func (x *UserInteractionSRes) Reset() {
	*x = UserInteractionSRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInteractionSRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInteractionSRes) ProtoMessage() {}

func (x *UserInteractionSRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInteractionSRes.ProtoReflect.Descriptor instead.
func (*UserInteractionSRes) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{3}
}

func (x *UserInteractionSRes) GetUserInteraction() []*UserInteractionRes {
	if x != nil {
		return x.UserInteraction
	}
	return nil
}

type UserInteractionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInteractionId string `protobuf:"bytes,1,opt,name=userInteractionId,proto3" json:"userInteractionId,omitempty"`
}

func (x *UserInteractionId) Reset() {
	*x = UserInteractionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_interactions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInteractionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInteractionId) ProtoMessage() {}

func (x *UserInteractionId) ProtoReflect() protoreflect.Message {
	mi := &file_user_interactions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInteractionId.ProtoReflect.Descriptor instead.
func (*UserInteractionId) Descriptor() ([]byte, []int) {
	return file_user_interactions_proto_rawDescGZIP(), []int{4}
}

func (x *UserInteractionId) GetUserInteractionId() string {
	if x != nil {
		return x.UserInteractionId
	}
	return ""
}

var File_user_interactions_proto protoreflect.FileDescriptor

var file_user_interactions_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x06, 0x0a, 0x04, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x79, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc8,
	0x01, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x65, 0x0a, 0x13, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x52, 0x65, 0x73,
	0x12, 0x4e, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52,
	0x0f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x41, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x32, 0x89, 0x03, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x1a, 0x25, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x52, 0x65, 0x73, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x46, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42,
	0x1c, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_interactions_proto_rawDescOnce sync.Once
	file_user_interactions_proto_rawDescData = file_user_interactions_proto_rawDesc
)

func file_user_interactions_proto_rawDescGZIP() []byte {
	file_user_interactions_proto_rawDescOnce.Do(func() {
		file_user_interactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_interactions_proto_rawDescData)
	})
	return file_user_interactions_proto_rawDescData
}

var file_user_interactions_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_interactions_proto_goTypes = []interface{}{
	(*Void)(nil),                // 0: userInteractions.Void
	(*UserInteraction)(nil),     // 1: userInteractions.UserInteraction
	(*UserInteractionRes)(nil),  // 2: userInteractions.UserInteractionRes
	(*UserInteractionSRes)(nil), // 3: userInteractions.UserInteractionSRes
	(*UserInteractionId)(nil),   // 4: userInteractions.UserInteractionId
}
var file_user_interactions_proto_depIdxs = []int32{
	2, // 0: userInteractions.UserInteractionSRes.userInteraction:type_name -> userInteractions.UserInteractionRes
	1, // 1: userInteractions.UserInteractionsService.Create:input_type -> userInteractions.UserInteraction
	0, // 2: userInteractions.UserInteractionsService.Get:input_type -> userInteractions.Void
	4, // 3: userInteractions.UserInteractionsService.GetById:input_type -> userInteractions.UserInteractionId
	2, // 4: userInteractions.UserInteractionsService.Update:input_type -> userInteractions.UserInteractionRes
	4, // 5: userInteractions.UserInteractionsService.Delete:input_type -> userInteractions.UserInteractionId
	0, // 6: userInteractions.UserInteractionsService.Create:output_type -> userInteractions.Void
	3, // 7: userInteractions.UserInteractionsService.Get:output_type -> userInteractions.UserInteractionSRes
	2, // 8: userInteractions.UserInteractionsService.GetById:output_type -> userInteractions.UserInteractionRes
	0, // 9: userInteractions.UserInteractionsService.Update:output_type -> userInteractions.Void
	0, // 10: userInteractions.UserInteractionsService.Delete:output_type -> userInteractions.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_interactions_proto_init() }
func file_user_interactions_proto_init() {
	if File_user_interactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_interactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_user_interactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInteraction); i {
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
		file_user_interactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInteractionRes); i {
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
		file_user_interactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInteractionSRes); i {
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
		file_user_interactions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInteractionId); i {
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
			RawDescriptor: file_user_interactions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_interactions_proto_goTypes,
		DependencyIndexes: file_user_interactions_proto_depIdxs,
		MessageInfos:      file_user_interactions_proto_msgTypes,
	}.Build()
	File_user_interactions_proto = out.File
	file_user_interactions_proto_rawDesc = nil
	file_user_interactions_proto_goTypes = nil
	file_user_interactions_proto_depIdxs = nil
}
