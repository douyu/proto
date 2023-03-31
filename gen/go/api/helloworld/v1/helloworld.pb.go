// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: helloworld/v1/helloworld.proto

package helloworldv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Sex
type Sex int32

const (
	// SEX_UNSPECIFIED ...
	Sex_SEX_UNSPECIFIED Sex = 0
	// SEX_MALE ...
	Sex_SEX_MALE Sex = 1
	// SEX_FEMALE ...
	Sex_SEX_FEMALE Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "SEX_UNSPECIFIED",
		1: "SEX_MALE",
		2: "SEX_FEMALE",
	}
	Sex_value = map[string]int32{
		"SEX_UNSPECIFIED": 0,
		"SEX_MALE":        1,
		"SEX_FEMALE":      2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld_v1_helloworld_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_helloworld_v1_helloworld_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}

// The request message containing the user's name.
type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name ...
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name" param:"name"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error
	Error uint32 `protobuf:"varint,1,opt,name=error,proto3" json:"error"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// data ...
	Data *SayHelloResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SayHelloResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SayHelloResponse) GetData() *SayHelloResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// The request message containing the user's name.
type SayHiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name ...
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name" form:"name"`
}

func (x *SayHiRequest) Reset() {
	*x = SayHiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHiRequest) ProtoMessage() {}

func (x *SayHiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHiRequest.ProtoReflect.Descriptor instead.
func (*SayHiRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *SayHiRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type SayHiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// error
	Error uint32 `protobuf:"varint,1,opt,name=error,proto3" json:"error"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	// data ...
	Data *SayHiResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data"`
}

func (x *SayHiResponse) Reset() {
	*x = SayHiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHiResponse) ProtoMessage() {}

func (x *SayHiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHiResponse.ProtoReflect.Descriptor instead.
func (*SayHiResponse) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{3}
}

func (x *SayHiResponse) GetError() uint32 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *SayHiResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SayHiResponse) GetData() *SayHiResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// Data is the data to be sent.
type SayHelloResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// age_number is the age number.
	AgeNumber uint64 `protobuf:"varint,2,opt,name=age_number,json=ageNumber,proto3" json:"ageNumber"`
	// sex is the user's sex
	Sex Sex `protobuf:"varint,3,opt,name=sex,proto3,enum=helloworld.v1.Sex" json:"sex"`
	// metadata is the user's metadata
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SayHelloResponse_Data) Reset() {
	*x = SayHelloResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse_Data) ProtoMessage() {}

func (x *SayHelloResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse_Data.ProtoReflect.Descriptor instead.
func (*SayHelloResponse_Data) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SayHelloResponse_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayHelloResponse_Data) GetAgeNumber() uint64 {
	if x != nil {
		return x.AgeNumber
	}
	return 0
}

func (x *SayHelloResponse_Data) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_SEX_UNSPECIFIED
}

func (x *SayHelloResponse_Data) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Data is the data to be sent.
type SayHiResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the user
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// age_number is the age number.
	AgeNumber uint64 `protobuf:"varint,2,opt,name=age_number,json=ageNumber,proto3" json:"ageNumber"`
}

func (x *SayHiResponse_Data) Reset() {
	*x = SayHiResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_v1_helloworld_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHiResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHiResponse_Data) ProtoMessage() {}

func (x *SayHiResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_v1_helloworld_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHiResponse_Data.ProtoReflect.Descriptor instead.
func (*SayHiResponse_Data) Descriptor() ([]byte, []int) {
	return file_helloworld_v1_helloworld_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SayHiResponse_Data) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayHiResponse_Data) GetAgeNumber() uint64 {
	if x != nil {
		return x.AgeNumber
	}
	return 0
}

var File_helloworld_v1_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_v1_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x1d, 0x9a, 0x84, 0x9e, 0x03, 0x18, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf5, 0x02, 0x0a, 0x10, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a,
	0xfe, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x03, 0x73, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03,
	0x73, 0x65, 0x78, 0x12, 0x4e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x34, 0x0a, 0x0c, 0x53, 0x61, 0x79, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10,
	0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x79, 0x48, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x79, 0x48, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2a, 0x38, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x58,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x45, 0x58, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a,
	0x53, 0x45, 0x58, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x32, 0xa8, 0x02, 0x0a,
	0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xcf, 0x01, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1e, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x81, 0x01,
	0x92, 0x41, 0x2a, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x30, 0x30, 0x31,
	0x72, 0x1b, 0x0a, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x20, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x28, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x4e, 0x3a, 0x01, 0x2a, 0x5a, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x2f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x44, 0x0a, 0x05, 0x53, 0x61, 0x79, 0x48, 0x69, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x69,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xba, 0x02, 0x92, 0x41, 0xca, 0x01, 0x12, 0x4e,
	0x0a, 0x1e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x20, 0x64, 0x65, 0x6d,
	0x6f, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x27, 0x0a, 0x05, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x12, 0x1e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f,
	0x75, 0x79, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x32, 0x21,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x77, 0x77,
	0x77, 0x2d, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x75, 0x72, 0x6c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x6a, 0x55, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x20, 0x30, 0x30, 0x31,
	0x12, 0x15, 0x54, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x68, 0x6f, 0x77, 0x20, 0x77, 0x65,
	0x20, 0x64, 0x6f, 0x20, 0x69, 0x74, 0x2e, 0x1a, 0x2f, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x20,
	0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75,
	0x79, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f,
	0x75, 0x79, 0x75, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x42, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld_v1_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_v1_helloworld_proto_rawDescData = file_helloworld_v1_helloworld_proto_rawDesc
)

func file_helloworld_v1_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_v1_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_v1_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_v1_helloworld_proto_rawDescData)
	})
	return file_helloworld_v1_helloworld_proto_rawDescData
}

var file_helloworld_v1_helloworld_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_helloworld_v1_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_helloworld_v1_helloworld_proto_goTypes = []interface{}{
	(Sex)(0),                      // 0: helloworld.v1.Sex
	(*SayHelloRequest)(nil),       // 1: helloworld.v1.SayHelloRequest
	(*SayHelloResponse)(nil),      // 2: helloworld.v1.SayHelloResponse
	(*SayHiRequest)(nil),          // 3: helloworld.v1.SayHiRequest
	(*SayHiResponse)(nil),         // 4: helloworld.v1.SayHiResponse
	(*SayHelloResponse_Data)(nil), // 5: helloworld.v1.SayHelloResponse.Data
	nil,                           // 6: helloworld.v1.SayHelloResponse.Data.MetadataEntry
	(*SayHiResponse_Data)(nil),    // 7: helloworld.v1.SayHiResponse.Data
}
var file_helloworld_v1_helloworld_proto_depIdxs = []int32{
	5, // 0: helloworld.v1.SayHelloResponse.data:type_name -> helloworld.v1.SayHelloResponse.Data
	7, // 1: helloworld.v1.SayHiResponse.data:type_name -> helloworld.v1.SayHiResponse.Data
	0, // 2: helloworld.v1.SayHelloResponse.Data.sex:type_name -> helloworld.v1.Sex
	6, // 3: helloworld.v1.SayHelloResponse.Data.metadata:type_name -> helloworld.v1.SayHelloResponse.Data.MetadataEntry
	1, // 4: helloworld.v1.GreeterService.SayHello:input_type -> helloworld.v1.SayHelloRequest
	3, // 5: helloworld.v1.GreeterService.SayHi:input_type -> helloworld.v1.SayHiRequest
	2, // 6: helloworld.v1.GreeterService.SayHello:output_type -> helloworld.v1.SayHelloResponse
	4, // 7: helloworld.v1.GreeterService.SayHi:output_type -> helloworld.v1.SayHiResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_helloworld_v1_helloworld_proto_init() }
func file_helloworld_v1_helloworld_proto_init() {
	if File_helloworld_v1_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_v1_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHiRequest); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHiResponse); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse_Data); i {
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
		file_helloworld_v1_helloworld_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHiResponse_Data); i {
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
			RawDescriptor: file_helloworld_v1_helloworld_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_v1_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_v1_helloworld_proto_depIdxs,
		EnumInfos:         file_helloworld_v1_helloworld_proto_enumTypes,
		MessageInfos:      file_helloworld_v1_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_v1_helloworld_proto = out.File
	file_helloworld_v1_helloworld_proto_rawDesc = nil
	file_helloworld_v1_helloworld_proto_goTypes = nil
	file_helloworld_v1_helloworld_proto_depIdxs = nil
}
