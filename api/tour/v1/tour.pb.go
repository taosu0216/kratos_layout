// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.1
// source: tour.proto

package v1

import (
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

type HiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HiRequest) Reset() {
	*x = HiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tour_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiRequest) ProtoMessage() {}

func (x *HiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tour_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiRequest.ProtoReflect.Descriptor instead.
func (*HiRequest) Descriptor() ([]byte, []int) {
	return file_tour_proto_rawDescGZIP(), []int{0}
}

type HiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HiReply) Reset() {
	*x = HiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tour_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HiReply) ProtoMessage() {}

func (x *HiReply) ProtoReflect() protoreflect.Message {
	mi := &file_tour_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HiReply.ProtoReflect.Descriptor instead.
func (*HiReply) Descriptor() ([]byte, []int) {
	return file_tour_proto_rawDescGZIP(), []int{1}
}

func (x *HiReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AiRequest) Reset() {
	*x = AiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tour_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiRequest) ProtoMessage() {}

func (x *AiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tour_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiRequest.ProtoReflect.Descriptor instead.
func (*AiRequest) Descriptor() ([]byte, []int) {
	return file_tour_proto_rawDescGZIP(), []int{2}
}

func (x *AiRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AiReply) Reset() {
	*x = AiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tour_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AiReply) ProtoMessage() {}

func (x *AiReply) ProtoReflect() protoreflect.Message {
	mi := &file_tour_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AiReply.ProtoReflect.Descriptor instead.
func (*AiReply) Descriptor() ([]byte, []int) {
	return file_tour_proto_rawDescGZIP(), []int{3}
}

func (x *AiReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_tour_proto protoreflect.FileDescriptor

var file_tour_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x6f,
	0x75, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x23, 0x0a, 0x07, 0x48, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x09, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x1b, 0x0a, 0x07, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x8a, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x75, 0x72, 0x12, 0x44, 0x0a, 0x08, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x6f, 0x75,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x02, 0x41, 0x69, 0x12, 0x12, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x6f, 0x75,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x69, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x2f, 0x61, 0x69, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tour_proto_rawDescOnce sync.Once
	file_tour_proto_rawDescData = file_tour_proto_rawDesc
)

func file_tour_proto_rawDescGZIP() []byte {
	file_tour_proto_rawDescOnce.Do(func() {
		file_tour_proto_rawDescData = protoimpl.X.CompressGZIP(file_tour_proto_rawDescData)
	})
	return file_tour_proto_rawDescData
}

var file_tour_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tour_proto_goTypes = []interface{}{
	(*HiRequest)(nil), // 0: tour.v1.HiRequest
	(*HiReply)(nil),   // 1: tour.v1.HiReply
	(*AiRequest)(nil), // 2: tour.v1.AiRequest
	(*AiReply)(nil),   // 3: tour.v1.AiReply
}
var file_tour_proto_depIdxs = []int32{
	0, // 0: tour.v1.Tour.SayHello:input_type -> tour.v1.HiRequest
	2, // 1: tour.v1.Tour.Ai:input_type -> tour.v1.AiRequest
	1, // 2: tour.v1.Tour.SayHello:output_type -> tour.v1.HiReply
	3, // 3: tour.v1.Tour.Ai:output_type -> tour.v1.AiReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tour_proto_init() }
func file_tour_proto_init() {
	if File_tour_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tour_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiRequest); i {
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
		file_tour_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HiReply); i {
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
		file_tour_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiRequest); i {
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
		file_tour_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AiReply); i {
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
			RawDescriptor: file_tour_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tour_proto_goTypes,
		DependencyIndexes: file_tour_proto_depIdxs,
		MessageInfos:      file_tour_proto_msgTypes,
	}.Build()
	File_tour_proto = out.File
	file_tour_proto_rawDesc = nil
	file_tour_proto_goTypes = nil
	file_tour_proto_depIdxs = nil
}
