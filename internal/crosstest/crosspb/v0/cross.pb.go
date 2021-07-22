// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/crosstest/crosspb/cross.proto

package crosspb

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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_internal_crosstest_crosspb_cross_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_internal_crosstest_crosspb_cross_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *FailRequest) Reset() {
	*x = FailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailRequest) ProtoMessage() {}

func (x *FailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailRequest.ProtoReflect.Descriptor instead.
func (*FailRequest) Descriptor() ([]byte, []int) {
	return file_internal_crosstest_crosspb_cross_proto_rawDescGZIP(), []int{2}
}

func (x *FailRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type FailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FailResponse) Reset() {
	*x = FailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailResponse) ProtoMessage() {}

func (x *FailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_crosstest_crosspb_cross_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailResponse.ProtoReflect.Descriptor instead.
func (*FailResponse) Descriptor() ([]byte, []int) {
	return file_internal_crosstest_crosspb_cross_proto_rawDescGZIP(), []int{3}
}

var File_internal_crosstest_crosspb_cross_proto protoreflect.FileDescriptor

var file_internal_crosstest_crosspb_cross_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x70, 0x62, 0x2f, 0x63, 0x72, 0x6f,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x72, 0x65, 0x72, 0x70, 0x63, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x30, 0x22, 0x25, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x0e, 0x0a,
	0x0c, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe1, 0x01,
	0x0a, 0x09, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x65, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x2e, 0x2e, 0x72, 0x65, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x63,
	0x72, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x72, 0x65, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x63,
	0x72, 0x6f, 0x73, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x2e,
	0x2e, 0x72, 0x65, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x2e,
	0x76, 0x30, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f,
	0x2e, 0x72, 0x65, 0x72, 0x70, 0x63, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x2e,
	0x76, 0x30, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6b, 0x73, 0x68, 0x61, 0x79, 0x6a, 0x73, 0x68, 0x61, 0x68, 0x2f, 0x72, 0x65, 0x72, 0x70,
	0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x70, 0x62, 0x2f, 0x76, 0x30, 0x3b,
	0x63, 0x72, 0x6f, 0x73, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_crosstest_crosspb_cross_proto_rawDescOnce sync.Once
	file_internal_crosstest_crosspb_cross_proto_rawDescData = file_internal_crosstest_crosspb_cross_proto_rawDesc
)

func file_internal_crosstest_crosspb_cross_proto_rawDescGZIP() []byte {
	file_internal_crosstest_crosspb_cross_proto_rawDescOnce.Do(func() {
		file_internal_crosstest_crosspb_cross_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_crosstest_crosspb_cross_proto_rawDescData)
	})
	return file_internal_crosstest_crosspb_cross_proto_rawDescData
}

var file_internal_crosstest_crosspb_cross_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_crosstest_crosspb_cross_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: rerpc.internal.crosstest.cross.v0.PingRequest
	(*PingResponse)(nil), // 1: rerpc.internal.crosstest.cross.v0.PingResponse
	(*FailRequest)(nil),  // 2: rerpc.internal.crosstest.cross.v0.FailRequest
	(*FailResponse)(nil), // 3: rerpc.internal.crosstest.cross.v0.FailResponse
}
var file_internal_crosstest_crosspb_cross_proto_depIdxs = []int32{
	0, // 0: rerpc.internal.crosstest.cross.v0.Crosstest.Ping:input_type -> rerpc.internal.crosstest.cross.v0.PingRequest
	2, // 1: rerpc.internal.crosstest.cross.v0.Crosstest.Fail:input_type -> rerpc.internal.crosstest.cross.v0.FailRequest
	1, // 2: rerpc.internal.crosstest.cross.v0.Crosstest.Ping:output_type -> rerpc.internal.crosstest.cross.v0.PingResponse
	3, // 3: rerpc.internal.crosstest.cross.v0.Crosstest.Fail:output_type -> rerpc.internal.crosstest.cross.v0.FailResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_crosstest_crosspb_cross_proto_init() }
func file_internal_crosstest_crosspb_cross_proto_init() {
	if File_internal_crosstest_crosspb_cross_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_crosstest_crosspb_cross_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_internal_crosstest_crosspb_cross_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_internal_crosstest_crosspb_cross_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailRequest); i {
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
		file_internal_crosstest_crosspb_cross_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailResponse); i {
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
			RawDescriptor: file_internal_crosstest_crosspb_cross_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_crosstest_crosspb_cross_proto_goTypes,
		DependencyIndexes: file_internal_crosstest_crosspb_cross_proto_depIdxs,
		MessageInfos:      file_internal_crosstest_crosspb_cross_proto_msgTypes,
	}.Build()
	File_internal_crosstest_crosspb_cross_proto = out.File
	file_internal_crosstest_crosspb_cross_proto_rawDesc = nil
	file_internal_crosstest_crosspb_cross_proto_goTypes = nil
	file_internal_crosstest_crosspb_cross_proto_depIdxs = nil
}
