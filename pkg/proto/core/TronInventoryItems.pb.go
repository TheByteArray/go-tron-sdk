// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: core/TronInventoryItems.proto

package core

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

type InventoryItems struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          int32                  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Items         [][]byte               `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InventoryItems) Reset() {
	*x = InventoryItems{}
	mi := &file_core_TronInventoryItems_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InventoryItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItems) ProtoMessage() {}

func (x *InventoryItems) ProtoReflect() protoreflect.Message {
	mi := &file_core_TronInventoryItems_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItems.ProtoReflect.Descriptor instead.
func (*InventoryItems) Descriptor() ([]byte, []int) {
	return file_core_TronInventoryItems_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryItems) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *InventoryItems) GetItems() [][]byte {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_core_TronInventoryItems_proto protoreflect.FileDescriptor

const file_core_TronInventoryItems_proto_rawDesc = "" +
	"\n" +
	"\x1dcore/TronInventoryItems.proto\x12\bprotocol\":\n" +
	"\x0eInventoryItems\x12\x12\n" +
	"\x04type\x18\x01 \x01(\x05R\x04type\x12\x14\n" +
	"\x05items\x18\x02 \x03(\fR\x05itemsBY\n" +
	"\x0forg.tron.protosB\x12TronInventoryItemsZ2github.com/TheByteArray/go-tron-sdk/pkg/proto/coreb\x06proto3"

var (
	file_core_TronInventoryItems_proto_rawDescOnce sync.Once
	file_core_TronInventoryItems_proto_rawDescData []byte
)

func file_core_TronInventoryItems_proto_rawDescGZIP() []byte {
	file_core_TronInventoryItems_proto_rawDescOnce.Do(func() {
		file_core_TronInventoryItems_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_core_TronInventoryItems_proto_rawDesc), len(file_core_TronInventoryItems_proto_rawDesc)))
	})
	return file_core_TronInventoryItems_proto_rawDescData
}

var file_core_TronInventoryItems_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_TronInventoryItems_proto_goTypes = []any{
	(*InventoryItems)(nil), // 0: protocol.InventoryItems
}
var file_core_TronInventoryItems_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_TronInventoryItems_proto_init() }
func file_core_TronInventoryItems_proto_init() {
	if File_core_TronInventoryItems_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_core_TronInventoryItems_proto_rawDesc), len(file_core_TronInventoryItems_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_TronInventoryItems_proto_goTypes,
		DependencyIndexes: file_core_TronInventoryItems_proto_depIdxs,
		MessageInfos:      file_core_TronInventoryItems_proto_msgTypes,
	}.Build()
	File_core_TronInventoryItems_proto = out.File
	file_core_TronInventoryItems_proto_goTypes = nil
	file_core_TronInventoryItems_proto_depIdxs = nil
}
