// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: fibonacci.proto

package proto

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

type FibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	mi := &file_fibonacci_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{0}
}

func (x *FibonacciRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

type FibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*FibonacciResponse_IntValue
	//	*FibonacciResponse_BigValue
	Value isFibonacciResponse_Value `protobuf_oneof:"value"`
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	mi := &file_fibonacci_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{1}
}

func (m *FibonacciResponse) GetValue() isFibonacciResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *FibonacciResponse) GetIntValue() int64 {
	if x, ok := x.GetValue().(*FibonacciResponse_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *FibonacciResponse) GetBigValue() string {
	if x, ok := x.GetValue().(*FibonacciResponse_BigValue); ok {
		return x.BigValue
	}
	return ""
}

type isFibonacciResponse_Value interface {
	isFibonacciResponse_Value()
}

type FibonacciResponse_IntValue struct {
	IntValue int64 `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3,oneof"`
}

type FibonacciResponse_BigValue struct {
	BigValue string `protobuf:"bytes,2,opt,name=big_value,json=bigValue,proto3,oneof"`
}

func (*FibonacciResponse_IntValue) isFibonacciResponse_Value() {}

func (*FibonacciResponse_BigValue) isFibonacciResponse_Value() {}

type FibonacciSequenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxIndex int64 `protobuf:"varint,1,opt,name=max_index,json=maxIndex,proto3" json:"max_index,omitempty"`
}

func (x *FibonacciSequenceRequest) Reset() {
	*x = FibonacciSequenceRequest{}
	mi := &file_fibonacci_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FibonacciSequenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciSequenceRequest) ProtoMessage() {}

func (x *FibonacciSequenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciSequenceRequest.ProtoReflect.Descriptor instead.
func (*FibonacciSequenceRequest) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{2}
}

func (x *FibonacciSequenceRequest) GetMaxIndex() int64 {
	if x != nil {
		return x.MaxIndex
	}
	return 0
}

type FibonacciSequenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence []string `protobuf:"bytes,1,rep,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *FibonacciSequenceResponse) Reset() {
	*x = FibonacciSequenceResponse{}
	mi := &file_fibonacci_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FibonacciSequenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciSequenceResponse) ProtoMessage() {}

func (x *FibonacciSequenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciSequenceResponse.ProtoReflect.Descriptor instead.
func (*FibonacciSequenceResponse) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{3}
}

func (x *FibonacciSequenceResponse) GetSequence() []string {
	if x != nil {
		return x.Sequence
	}
	return nil
}

var File_fibonacci_proto protoreflect.FileDescriptor

var file_fibonacci_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x22, 0x28, 0x0a, 0x10,
	0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x5a, 0x0a, 0x11, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61,
	0x63, 0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x08, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x69,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x62, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x37, 0x0a, 0x18, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x37, 0x0a, 0x19, 0x46,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x32, 0xc6, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63,
	0x63, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x46, 0x69, 0x62, 0x6f,
	0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x66,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63,
	0x63, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x23, 0x2e, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x46,
	0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61,
	0x63, 0x63, 0x69, 0x2e, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x13, 0x5a,
	0x11, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fibonacci_proto_rawDescOnce sync.Once
	file_fibonacci_proto_rawDescData = file_fibonacci_proto_rawDesc
)

func file_fibonacci_proto_rawDescGZIP() []byte {
	file_fibonacci_proto_rawDescOnce.Do(func() {
		file_fibonacci_proto_rawDescData = protoimpl.X.CompressGZIP(file_fibonacci_proto_rawDescData)
	})
	return file_fibonacci_proto_rawDescData
}

var file_fibonacci_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_fibonacci_proto_goTypes = []any{
	(*FibonacciRequest)(nil),          // 0: fibonacci.FibonacciRequest
	(*FibonacciResponse)(nil),         // 1: fibonacci.FibonacciResponse
	(*FibonacciSequenceRequest)(nil),  // 2: fibonacci.FibonacciSequenceRequest
	(*FibonacciSequenceResponse)(nil), // 3: fibonacci.FibonacciSequenceResponse
}
var file_fibonacci_proto_depIdxs = []int32{
	0, // 0: fibonacci.FibonacciService.GetFibonacciNumber:input_type -> fibonacci.FibonacciRequest
	2, // 1: fibonacci.FibonacciService.GetFibonacciSequence:input_type -> fibonacci.FibonacciSequenceRequest
	1, // 2: fibonacci.FibonacciService.GetFibonacciNumber:output_type -> fibonacci.FibonacciResponse
	3, // 3: fibonacci.FibonacciService.GetFibonacciSequence:output_type -> fibonacci.FibonacciSequenceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fibonacci_proto_init() }
func file_fibonacci_proto_init() {
	if File_fibonacci_proto != nil {
		return
	}
	file_fibonacci_proto_msgTypes[1].OneofWrappers = []any{
		(*FibonacciResponse_IntValue)(nil),
		(*FibonacciResponse_BigValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fibonacci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fibonacci_proto_goTypes,
		DependencyIndexes: file_fibonacci_proto_depIdxs,
		MessageInfos:      file_fibonacci_proto_msgTypes,
	}.Build()
	File_fibonacci_proto = out.File
	file_fibonacci_proto_rawDesc = nil
	file_fibonacci_proto_goTypes = nil
	file_fibonacci_proto_depIdxs = nil
}
