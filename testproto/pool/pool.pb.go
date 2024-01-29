// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: pool/pool.proto

package pool

import (
	_ "github.com/planetscale/vtprotobuf/vtproto"
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

type OptionalMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OptionalMessage) Reset() {
	*x = OptionalMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionalMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionalMessage) ProtoMessage() {}

func (x *OptionalMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionalMessage.ProtoReflect.Descriptor instead.
func (*OptionalMessage) Descriptor() ([]byte, []int) {
	return file_pool_pool_proto_rawDescGZIP(), []int{0}
}

type MemoryPoolExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo1 string           `protobuf:"bytes,1,opt,name=foo1,proto3" json:"foo1,omitempty"`
	Foo2 uint64           `protobuf:"varint,2,opt,name=foo2,proto3" json:"foo2,omitempty"`
	Foo3 *OptionalMessage `protobuf:"bytes,3,opt,name=foo3,proto3,oneof" json:"foo3,omitempty"`
}

func (x *MemoryPoolExtension) Reset() {
	*x = MemoryPoolExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pool_pool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryPoolExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryPoolExtension) ProtoMessage() {}

func (x *MemoryPoolExtension) ProtoReflect() protoreflect.Message {
	mi := &file_pool_pool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryPoolExtension.ProtoReflect.Descriptor instead.
func (*MemoryPoolExtension) Descriptor() ([]byte, []int) {
	return file_pool_pool_proto_rawDescGZIP(), []int{1}
}

func (x *MemoryPoolExtension) GetFoo1() string {
	if x != nil {
		return x.Foo1
	}
	return ""
}

func (x *MemoryPoolExtension) GetFoo2() uint64 {
	if x != nil {
		return x.Foo2
	}
	return 0
}

func (x *MemoryPoolExtension) GetFoo3() *OptionalMessage {
	if x != nil {
		return x.Foo3
	}
	return nil
}

var File_pool_pool_proto protoreflect.FileDescriptor

var file_pool_pool_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x0f, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x22,
	0x77, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x6f, 0x6c, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6f,
	0x6f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x32, 0x12, 0x29,
	0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x66, 0x6f, 0x6f, 0x33, 0x88, 0x01, 0x01, 0x3a, 0x04, 0xa8, 0xa6, 0x1f, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x66, 0x6f, 0x6f, 0x33, 0x42, 0x10, 0x5a, 0x0e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pool_pool_proto_rawDescOnce sync.Once
	file_pool_pool_proto_rawDescData = file_pool_pool_proto_rawDesc
)

func file_pool_pool_proto_rawDescGZIP() []byte {
	file_pool_pool_proto_rawDescOnce.Do(func() {
		file_pool_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_pool_pool_proto_rawDescData)
	})
	return file_pool_pool_proto_rawDescData
}

var file_pool_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pool_pool_proto_goTypes = []interface{}{
	(*OptionalMessage)(nil),     // 0: OptionalMessage
	(*MemoryPoolExtension)(nil), // 1: MemoryPoolExtension
}
var file_pool_pool_proto_depIdxs = []int32{
	0, // 0: MemoryPoolExtension.foo3:type_name -> OptionalMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pool_pool_proto_init() }
func file_pool_pool_proto_init() {
	if File_pool_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pool_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionalMessage); i {
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
		file_pool_pool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryPoolExtension); i {
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
	file_pool_pool_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pool_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pool_pool_proto_goTypes,
		DependencyIndexes: file_pool_pool_proto_depIdxs,
		MessageInfos:      file_pool_pool_proto_msgTypes,
	}.Build()
	File_pool_pool_proto = out.File
	file_pool_pool_proto_rawDesc = nil
	file_pool_pool_proto_goTypes = nil
	file_pool_pool_proto_depIdxs = nil
}
