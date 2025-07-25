// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: core/contract/witness_contract.proto

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

type WitnessCreateContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress  []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Url           []byte                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WitnessCreateContract) Reset() {
	*x = WitnessCreateContract{}
	mi := &file_core_contract_witness_contract_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WitnessCreateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessCreateContract) ProtoMessage() {}

func (x *WitnessCreateContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_witness_contract_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessCreateContract.ProtoReflect.Descriptor instead.
func (*WitnessCreateContract) Descriptor() ([]byte, []int) {
	return file_core_contract_witness_contract_proto_rawDescGZIP(), []int{0}
}

func (x *WitnessCreateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *WitnessCreateContract) GetUrl() []byte {
	if x != nil {
		return x.Url
	}
	return nil
}

type WitnessUpdateContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress  []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	UpdateUrl     []byte                 `protobuf:"bytes,12,opt,name=update_url,json=updateUrl,proto3" json:"update_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WitnessUpdateContract) Reset() {
	*x = WitnessUpdateContract{}
	mi := &file_core_contract_witness_contract_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WitnessUpdateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessUpdateContract) ProtoMessage() {}

func (x *WitnessUpdateContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_witness_contract_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessUpdateContract.ProtoReflect.Descriptor instead.
func (*WitnessUpdateContract) Descriptor() ([]byte, []int) {
	return file_core_contract_witness_contract_proto_rawDescGZIP(), []int{1}
}

func (x *WitnessUpdateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *WitnessUpdateContract) GetUpdateUrl() []byte {
	if x != nil {
		return x.UpdateUrl
	}
	return nil
}

type VoteWitnessContract struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	OwnerAddress  []byte                      `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Votes         []*VoteWitnessContract_Vote `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	Support       bool                        `protobuf:"varint,3,opt,name=support,proto3" json:"support,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoteWitnessContract) Reset() {
	*x = VoteWitnessContract{}
	mi := &file_core_contract_witness_contract_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteWitnessContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteWitnessContract) ProtoMessage() {}

func (x *VoteWitnessContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_witness_contract_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteWitnessContract.ProtoReflect.Descriptor instead.
func (*VoteWitnessContract) Descriptor() ([]byte, []int) {
	return file_core_contract_witness_contract_proto_rawDescGZIP(), []int{2}
}

func (x *VoteWitnessContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *VoteWitnessContract) GetVotes() []*VoteWitnessContract_Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

func (x *VoteWitnessContract) GetSupport() bool {
	if x != nil {
		return x.Support
	}
	return false
}

type VoteWitnessContract_Vote struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	VoteAddress   []byte                 `protobuf:"bytes,1,opt,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	VoteCount     int64                  `protobuf:"varint,2,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoteWitnessContract_Vote) Reset() {
	*x = VoteWitnessContract_Vote{}
	mi := &file_core_contract_witness_contract_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteWitnessContract_Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteWitnessContract_Vote) ProtoMessage() {}

func (x *VoteWitnessContract_Vote) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_witness_contract_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteWitnessContract_Vote.ProtoReflect.Descriptor instead.
func (*VoteWitnessContract_Vote) Descriptor() ([]byte, []int) {
	return file_core_contract_witness_contract_proto_rawDescGZIP(), []int{2, 0}
}

func (x *VoteWitnessContract_Vote) GetVoteAddress() []byte {
	if x != nil {
		return x.VoteAddress
	}
	return nil
}

func (x *VoteWitnessContract_Vote) GetVoteCount() int64 {
	if x != nil {
		return x.VoteCount
	}
	return 0
}

var File_core_contract_witness_contract_proto protoreflect.FileDescriptor

const file_core_contract_witness_contract_proto_rawDesc = "" +
	"\n" +
	"$core/contract/witness_contract.proto\x12\bprotocol\"N\n" +
	"\x15WitnessCreateContract\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\x12\x10\n" +
	"\x03url\x18\x02 \x01(\fR\x03url\"[\n" +
	"\x15WitnessUpdateContract\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\x12\x1d\n" +
	"\n" +
	"update_url\x18\f \x01(\fR\tupdateUrl\"\xd8\x01\n" +
	"\x13VoteWitnessContract\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\x128\n" +
	"\x05votes\x18\x02 \x03(\v2\".protocol.VoteWitnessContract.VoteR\x05votes\x12\x18\n" +
	"\asupport\x18\x03 \x01(\bR\asupport\x1aH\n" +
	"\x04Vote\x12!\n" +
	"\fvote_address\x18\x01 \x01(\fR\vvoteAddress\x12\x1d\n" +
	"\n" +
	"vote_count\x18\x02 \x01(\x03R\tvoteCountBN\n" +
	"\x18org.tron.protos.contractZ2github.com/TheByteArray/go-tron-sdk/pkg/proto/coreb\x06proto3"

var (
	file_core_contract_witness_contract_proto_rawDescOnce sync.Once
	file_core_contract_witness_contract_proto_rawDescData []byte
)

func file_core_contract_witness_contract_proto_rawDescGZIP() []byte {
	file_core_contract_witness_contract_proto_rawDescOnce.Do(func() {
		file_core_contract_witness_contract_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_core_contract_witness_contract_proto_rawDesc), len(file_core_contract_witness_contract_proto_rawDesc)))
	})
	return file_core_contract_witness_contract_proto_rawDescData
}

var file_core_contract_witness_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_core_contract_witness_contract_proto_goTypes = []any{
	(*WitnessCreateContract)(nil),    // 0: protocol.WitnessCreateContract
	(*WitnessUpdateContract)(nil),    // 1: protocol.WitnessUpdateContract
	(*VoteWitnessContract)(nil),      // 2: protocol.VoteWitnessContract
	(*VoteWitnessContract_Vote)(nil), // 3: protocol.VoteWitnessContract.Vote
}
var file_core_contract_witness_contract_proto_depIdxs = []int32{
	3, // 0: protocol.VoteWitnessContract.votes:type_name -> protocol.VoteWitnessContract.Vote
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_contract_witness_contract_proto_init() }
func file_core_contract_witness_contract_proto_init() {
	if File_core_contract_witness_contract_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_core_contract_witness_contract_proto_rawDesc), len(file_core_contract_witness_contract_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_contract_witness_contract_proto_goTypes,
		DependencyIndexes: file_core_contract_witness_contract_proto_depIdxs,
		MessageInfos:      file_core_contract_witness_contract_proto_msgTypes,
	}.Build()
	File_core_contract_witness_contract_proto = out.File
	file_core_contract_witness_contract_proto_goTypes = nil
	file_core_contract_witness_contract_proto_depIdxs = nil
}
