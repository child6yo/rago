// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: generator/generator.proto

package generator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Query struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Query          string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	CollectionName string                 `protobuf:"bytes,2,opt,name=collectionName,proto3" json:"collectionName,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Query) Reset() {
	*x = Query{}
	mi := &file_generator_generator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_generator_generator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_generator_generator_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Query) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

type ResponseChunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chunk         string                 `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseChunk) Reset() {
	*x = ResponseChunk{}
	mi := &file_generator_generator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseChunk) ProtoMessage() {}

func (x *ResponseChunk) ProtoReflect() protoreflect.Message {
	mi := &file_generator_generator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseChunk.ProtoReflect.Descriptor instead.
func (*ResponseChunk) Descriptor() ([]byte, []int) {
	return file_generator_generator_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseChunk) GetChunk() string {
	if x != nil {
		return x.Chunk
	}
	return ""
}

var File_generator_generator_proto protoreflect.FileDescriptor

const file_generator_generator_proto_rawDesc = "" +
	"\n" +
	"\x19generator/generator.proto\x12\tgenerator\"E\n" +
	"\x05Query\x12\x14\n" +
	"\x05query\x18\x01 \x01(\tR\x05query\x12&\n" +
	"\x0ecollectionName\x18\x02 \x01(\tR\x0ecollectionName\"%\n" +
	"\rResponseChunk\x12\x14\n" +
	"\x05chunk\x18\x01 \x01(\tR\x05chunk2L\n" +
	"\x10GeneratorService\x128\n" +
	"\bGenerate\x12\x10.generator.Query\x1a\x18.generator.ResponseChunk0\x01B*Z(github.com/child6yo/rago/proto/generatorb\x06proto3"

var (
	file_generator_generator_proto_rawDescOnce sync.Once
	file_generator_generator_proto_rawDescData []byte
)

func file_generator_generator_proto_rawDescGZIP() []byte {
	file_generator_generator_proto_rawDescOnce.Do(func() {
		file_generator_generator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_generator_generator_proto_rawDesc), len(file_generator_generator_proto_rawDesc)))
	})
	return file_generator_generator_proto_rawDescData
}

var file_generator_generator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_generator_generator_proto_goTypes = []any{
	(*Query)(nil),         // 0: generator.Query
	(*ResponseChunk)(nil), // 1: generator.ResponseChunk
}
var file_generator_generator_proto_depIdxs = []int32{
	0, // 0: generator.GeneratorService.Generate:input_type -> generator.Query
	1, // 1: generator.GeneratorService.Generate:output_type -> generator.ResponseChunk
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_generator_generator_proto_init() }
func file_generator_generator_proto_init() {
	if File_generator_generator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_generator_generator_proto_rawDesc), len(file_generator_generator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generator_generator_proto_goTypes,
		DependencyIndexes: file_generator_generator_proto_depIdxs,
		MessageInfos:      file_generator_generator_proto_msgTypes,
	}.Build()
	File_generator_generator_proto = out.File
	file_generator_generator_proto_goTypes = nil
	file_generator_generator_proto_depIdxs = nil
}
