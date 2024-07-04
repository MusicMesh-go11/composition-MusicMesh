// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: composition_metadata.proto

package composition_metadata

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
		mi := &file_composition_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_composition_metadata_proto_msgTypes[0]
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
	return file_composition_metadata_proto_rawDescGZIP(), []int{0}
}

type CompositionMetadataId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataId string `protobuf:"bytes,1,opt,name=metadataId,proto3" json:"metadataId,omitempty"`
}

func (x *CompositionMetadataId) Reset() {
	*x = CompositionMetadataId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composition_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositionMetadataId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositionMetadataId) ProtoMessage() {}

func (x *CompositionMetadataId) ProtoReflect() protoreflect.Message {
	mi := &file_composition_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositionMetadataId.ProtoReflect.Descriptor instead.
func (*CompositionMetadataId) Descriptor() ([]byte, []int) {
	return file_composition_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *CompositionMetadataId) GetMetadataId() string {
	if x != nil {
		return x.MetadataId
	}
	return ""
}

type GenreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre string `protobuf:"bytes,1,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *GenreRequest) Reset() {
	*x = GenreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composition_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenreRequest) ProtoMessage() {}

func (x *GenreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_composition_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenreRequest.ProtoReflect.Descriptor instead.
func (*GenreRequest) Descriptor() ([]byte, []int) {
	return file_composition_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *GenreRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type CompositionMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Genre string `protobuf:"bytes,1,opt,name=genre,proto3" json:"genre,omitempty"`
	Tags  string `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CompositionMetadata) Reset() {
	*x = CompositionMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composition_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositionMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositionMetadata) ProtoMessage() {}

func (x *CompositionMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_composition_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositionMetadata.ProtoReflect.Descriptor instead.
func (*CompositionMetadata) Descriptor() ([]byte, []int) {
	return file_composition_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *CompositionMetadata) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *CompositionMetadata) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type CompositionsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Compositions []*CompositionRes `protobuf:"bytes,1,rep,name=compositions,proto3" json:"compositions,omitempty"`
}

func (x *CompositionsRes) Reset() {
	*x = CompositionsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composition_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositionsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositionsRes) ProtoMessage() {}

func (x *CompositionsRes) ProtoReflect() protoreflect.Message {
	mi := &file_composition_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositionsRes.ProtoReflect.Descriptor instead.
func (*CompositionsRes) Descriptor() ([]byte, []int) {
	return file_composition_metadata_proto_rawDescGZIP(), []int{4}
}

func (x *CompositionsRes) GetCompositions() []*CompositionRes {
	if x != nil {
		return x.Compositions
	}
	return nil
}

type CompositionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompositionId string `protobuf:"bytes,1,opt,name=composition_id,json=compositionId,proto3" json:"composition_id,omitempty"`
	Genre         string `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	Tags          string `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	ListenCount   int64  `protobuf:"varint,4,opt,name=listen_count,json=listenCount,proto3" json:"listen_count,omitempty"`
	LikeCount     int64  `protobuf:"varint,5,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
}

func (x *CompositionRes) Reset() {
	*x = CompositionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_composition_metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositionRes) ProtoMessage() {}

func (x *CompositionRes) ProtoReflect() protoreflect.Message {
	mi := &file_composition_metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositionRes.ProtoReflect.Descriptor instead.
func (*CompositionRes) Descriptor() ([]byte, []int) {
	return file_composition_metadata_proto_rawDescGZIP(), []int{5}
}

func (x *CompositionRes) GetCompositionId() string {
	if x != nil {
		return x.CompositionId
	}
	return ""
}

func (x *CompositionRes) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *CompositionRes) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

func (x *CompositionRes) GetListenCount() int64 {
	if x != nil {
		return x.ListenCount
	}
	return 0
}

func (x *CompositionRes) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

var File_composition_metadata_proto protoreflect.FileDescriptor

var file_composition_metadata_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x15, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x49, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x5a, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xfe, 0x04, 0x0a, 0x1a,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x6e, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x12, 0x48, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x4f, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49,
	0x64, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x1f, 0x5a, 0x1d,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_composition_metadata_proto_rawDescOnce sync.Once
	file_composition_metadata_proto_rawDescData = file_composition_metadata_proto_rawDesc
)

func file_composition_metadata_proto_rawDescGZIP() []byte {
	file_composition_metadata_proto_rawDescOnce.Do(func() {
		file_composition_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_composition_metadata_proto_rawDescData)
	})
	return file_composition_metadata_proto_rawDescData
}

var file_composition_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_composition_metadata_proto_goTypes = []interface{}{
	(*Void)(nil),                  // 0: compositionMetadata.Void
	(*CompositionMetadataId)(nil), // 1: compositionMetadata.CompositionMetadataId
	(*GenreRequest)(nil),          // 2: compositionMetadata.GenreRequest
	(*CompositionMetadata)(nil),   // 3: compositionMetadata.CompositionMetadata
	(*CompositionsRes)(nil),       // 4: compositionMetadata.CompositionsRes
	(*CompositionRes)(nil),        // 5: compositionMetadata.CompositionRes
}
var file_composition_metadata_proto_depIdxs = []int32{
	5, // 0: compositionMetadata.CompositionsRes.compositions:type_name -> compositionMetadata.CompositionRes
	3, // 1: compositionMetadata.CompositionMetadataService.Create:input_type -> compositionMetadata.CompositionMetadata
	0, // 2: compositionMetadata.CompositionMetadataService.GetTrending:input_type -> compositionMetadata.Void
	0, // 3: compositionMetadata.CompositionMetadataService.GetRecommended:input_type -> compositionMetadata.Void
	2, // 4: compositionMetadata.CompositionMetadataService.GetByGenre:input_type -> compositionMetadata.GenreRequest
	0, // 5: compositionMetadata.CompositionMetadataService.Like:input_type -> compositionMetadata.Void
	0, // 6: compositionMetadata.CompositionMetadataService.UnLike:input_type -> compositionMetadata.Void
	5, // 7: compositionMetadata.CompositionMetadataService.Update:input_type -> compositionMetadata.CompositionRes
	1, // 8: compositionMetadata.CompositionMetadataService.Delete:input_type -> compositionMetadata.CompositionMetadataId
	0, // 9: compositionMetadata.CompositionMetadataService.Create:output_type -> compositionMetadata.Void
	4, // 10: compositionMetadata.CompositionMetadataService.GetTrending:output_type -> compositionMetadata.CompositionsRes
	4, // 11: compositionMetadata.CompositionMetadataService.GetRecommended:output_type -> compositionMetadata.CompositionsRes
	4, // 12: compositionMetadata.CompositionMetadataService.GetByGenre:output_type -> compositionMetadata.CompositionsRes
	0, // 13: compositionMetadata.CompositionMetadataService.Like:output_type -> compositionMetadata.Void
	0, // 14: compositionMetadata.CompositionMetadataService.UnLike:output_type -> compositionMetadata.Void
	0, // 15: compositionMetadata.CompositionMetadataService.Update:output_type -> compositionMetadata.Void
	0, // 16: compositionMetadata.CompositionMetadataService.Delete:output_type -> compositionMetadata.Void
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_composition_metadata_proto_init() }
func file_composition_metadata_proto_init() {
	if File_composition_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_composition_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_composition_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositionMetadataId); i {
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
		file_composition_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenreRequest); i {
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
		file_composition_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositionMetadata); i {
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
		file_composition_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositionsRes); i {
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
		file_composition_metadata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositionRes); i {
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
			RawDescriptor: file_composition_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_composition_metadata_proto_goTypes,
		DependencyIndexes: file_composition_metadata_proto_depIdxs,
		MessageInfos:      file_composition_metadata_proto_msgTypes,
	}.Build()
	File_composition_metadata_proto = out.File
	file_composition_metadata_proto_rawDesc = nil
	file_composition_metadata_proto_goTypes = nil
	file_composition_metadata_proto_depIdxs = nil
}