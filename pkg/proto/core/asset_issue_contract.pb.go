// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: core/contract/asset_issue_contract.proto

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

type AssetIssueContract struct {
	state                   protoimpl.MessageState             `protogen:"open.v1"`
	Id                      string                             `protobuf:"bytes,41,opt,name=id,proto3" json:"id,omitempty"`
	OwnerAddress            []byte                             `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Name                    []byte                             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbr                    []byte                             `protobuf:"bytes,3,opt,name=abbr,proto3" json:"abbr,omitempty"`
	TotalSupply             int64                              `protobuf:"varint,4,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty"`
	FrozenSupply            []*AssetIssueContract_FrozenSupply `protobuf:"bytes,5,rep,name=frozen_supply,json=frozenSupply,proto3" json:"frozen_supply,omitempty"`
	TrxNum                  int32                              `protobuf:"varint,6,opt,name=trx_num,json=trxNum,proto3" json:"trx_num,omitempty"`
	Precision               int32                              `protobuf:"varint,7,opt,name=precision,proto3" json:"precision,omitempty"`
	Num                     int32                              `protobuf:"varint,8,opt,name=num,proto3" json:"num,omitempty"`
	StartTime               int64                              `protobuf:"varint,9,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime                 int64                              `protobuf:"varint,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Order                   int64                              `protobuf:"varint,11,opt,name=order,proto3" json:"order,omitempty"` // useless
	VoteScore               int32                              `protobuf:"varint,16,opt,name=vote_score,json=voteScore,proto3" json:"vote_score,omitempty"`
	Description             []byte                             `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	Url                     []byte                             `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`
	FreeAssetNetLimit       int64                              `protobuf:"varint,22,opt,name=free_asset_net_limit,json=freeAssetNetLimit,proto3" json:"free_asset_net_limit,omitempty"`
	PublicFreeAssetNetLimit int64                              `protobuf:"varint,23,opt,name=public_free_asset_net_limit,json=publicFreeAssetNetLimit,proto3" json:"public_free_asset_net_limit,omitempty"`
	PublicFreeAssetNetUsage int64                              `protobuf:"varint,24,opt,name=public_free_asset_net_usage,json=publicFreeAssetNetUsage,proto3" json:"public_free_asset_net_usage,omitempty"`
	PublicLatestFreeNetTime int64                              `protobuf:"varint,25,opt,name=public_latest_free_net_time,json=publicLatestFreeNetTime,proto3" json:"public_latest_free_net_time,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *AssetIssueContract) Reset() {
	*x = AssetIssueContract{}
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetIssueContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetIssueContract) ProtoMessage() {}

func (x *AssetIssueContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetIssueContract.ProtoReflect.Descriptor instead.
func (*AssetIssueContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{0}
}

func (x *AssetIssueContract) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssetIssueContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *AssetIssueContract) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *AssetIssueContract) GetAbbr() []byte {
	if x != nil {
		return x.Abbr
	}
	return nil
}

func (x *AssetIssueContract) GetTotalSupply() int64 {
	if x != nil {
		return x.TotalSupply
	}
	return 0
}

func (x *AssetIssueContract) GetFrozenSupply() []*AssetIssueContract_FrozenSupply {
	if x != nil {
		return x.FrozenSupply
	}
	return nil
}

func (x *AssetIssueContract) GetTrxNum() int32 {
	if x != nil {
		return x.TrxNum
	}
	return 0
}

func (x *AssetIssueContract) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *AssetIssueContract) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *AssetIssueContract) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *AssetIssueContract) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *AssetIssueContract) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *AssetIssueContract) GetVoteScore() int32 {
	if x != nil {
		return x.VoteScore
	}
	return 0
}

func (x *AssetIssueContract) GetDescription() []byte {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *AssetIssueContract) GetUrl() []byte {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *AssetIssueContract) GetFreeAssetNetLimit() int64 {
	if x != nil {
		return x.FreeAssetNetLimit
	}
	return 0
}

func (x *AssetIssueContract) GetPublicFreeAssetNetLimit() int64 {
	if x != nil {
		return x.PublicFreeAssetNetLimit
	}
	return 0
}

func (x *AssetIssueContract) GetPublicFreeAssetNetUsage() int64 {
	if x != nil {
		return x.PublicFreeAssetNetUsage
	}
	return 0
}

func (x *AssetIssueContract) GetPublicLatestFreeNetTime() int64 {
	if x != nil {
		return x.PublicLatestFreeNetTime
	}
	return 0
}

type TransferAssetContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetName     []byte                 `protobuf:"bytes,1,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"` // this field is token name before the proposal ALLOW_SAME_TOKEN_NAME is active, otherwise it is token id and token is should be in string format.
	OwnerAddress  []byte                 `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress     []byte                 `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount        int64                  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferAssetContract) Reset() {
	*x = TransferAssetContract{}
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferAssetContract) ProtoMessage() {}

func (x *TransferAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferAssetContract.ProtoReflect.Descriptor instead.
func (*TransferAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{1}
}

func (x *TransferAssetContract) GetAssetName() []byte {
	if x != nil {
		return x.AssetName
	}
	return nil
}

func (x *TransferAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *TransferAssetContract) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *TransferAssetContract) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type UnfreezeAssetContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress  []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnfreezeAssetContract) Reset() {
	*x = UnfreezeAssetContract{}
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnfreezeAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfreezeAssetContract) ProtoMessage() {}

func (x *UnfreezeAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfreezeAssetContract.ProtoReflect.Descriptor instead.
func (*UnfreezeAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{2}
}

func (x *UnfreezeAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

type UpdateAssetContract struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress   []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Description    []byte                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url            []byte                 `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	NewLimit       int64                  `protobuf:"varint,4,opt,name=new_limit,json=newLimit,proto3" json:"new_limit,omitempty"`
	NewPublicLimit int64                  `protobuf:"varint,5,opt,name=new_public_limit,json=newPublicLimit,proto3" json:"new_public_limit,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateAssetContract) Reset() {
	*x = UpdateAssetContract{}
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAssetContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssetContract) ProtoMessage() {}

func (x *UpdateAssetContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssetContract.ProtoReflect.Descriptor instead.
func (*UpdateAssetContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAssetContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *UpdateAssetContract) GetDescription() []byte {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdateAssetContract) GetUrl() []byte {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *UpdateAssetContract) GetNewLimit() int64 {
	if x != nil {
		return x.NewLimit
	}
	return 0
}

func (x *UpdateAssetContract) GetNewPublicLimit() int64 {
	if x != nil {
		return x.NewPublicLimit
	}
	return 0
}

type ParticipateAssetIssueContract struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerAddress  []byte                 `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress     []byte                 `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	AssetName     []byte                 `protobuf:"bytes,3,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"` // this field is token name before the proposal ALLOW_SAME_TOKEN_NAME is active, otherwise it is token id and token is should be in string format.
	Amount        int64                  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`                       // the amount of drops
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParticipateAssetIssueContract) Reset() {
	*x = ParticipateAssetIssueContract{}
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParticipateAssetIssueContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipateAssetIssueContract) ProtoMessage() {}

func (x *ParticipateAssetIssueContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipateAssetIssueContract.ProtoReflect.Descriptor instead.
func (*ParticipateAssetIssueContract) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{4}
}

func (x *ParticipateAssetIssueContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ParticipateAssetIssueContract) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *ParticipateAssetIssueContract) GetAssetName() []byte {
	if x != nil {
		return x.AssetName
	}
	return nil
}

func (x *ParticipateAssetIssueContract) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AssetIssueContract_FrozenSupply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FrozenAmount  int64                  `protobuf:"varint,1,opt,name=frozen_amount,json=frozenAmount,proto3" json:"frozen_amount,omitempty"`
	FrozenDays    int64                  `protobuf:"varint,2,opt,name=frozen_days,json=frozenDays,proto3" json:"frozen_days,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetIssueContract_FrozenSupply) Reset() {
	*x = AssetIssueContract_FrozenSupply{}
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetIssueContract_FrozenSupply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetIssueContract_FrozenSupply) ProtoMessage() {}

func (x *AssetIssueContract_FrozenSupply) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_asset_issue_contract_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetIssueContract_FrozenSupply.ProtoReflect.Descriptor instead.
func (*AssetIssueContract_FrozenSupply) Descriptor() ([]byte, []int) {
	return file_core_contract_asset_issue_contract_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AssetIssueContract_FrozenSupply) GetFrozenAmount() int64 {
	if x != nil {
		return x.FrozenAmount
	}
	return 0
}

func (x *AssetIssueContract_FrozenSupply) GetFrozenDays() int64 {
	if x != nil {
		return x.FrozenDays
	}
	return 0
}

var File_core_contract_asset_issue_contract_proto protoreflect.FileDescriptor

const file_core_contract_asset_issue_contract_proto_rawDesc = "" +
	"\n" +
	"(core/contract/asset_issue_contract.proto\x12\bprotocol\"\x91\x06\n" +
	"\x12AssetIssueContract\x12\x0e\n" +
	"\x02id\x18) \x01(\tR\x02id\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\x12\x12\n" +
	"\x04name\x18\x02 \x01(\fR\x04name\x12\x12\n" +
	"\x04abbr\x18\x03 \x01(\fR\x04abbr\x12!\n" +
	"\ftotal_supply\x18\x04 \x01(\x03R\vtotalSupply\x12N\n" +
	"\rfrozen_supply\x18\x05 \x03(\v2).protocol.AssetIssueContract.FrozenSupplyR\ffrozenSupply\x12\x17\n" +
	"\atrx_num\x18\x06 \x01(\x05R\x06trxNum\x12\x1c\n" +
	"\tprecision\x18\a \x01(\x05R\tprecision\x12\x10\n" +
	"\x03num\x18\b \x01(\x05R\x03num\x12\x1d\n" +
	"\n" +
	"start_time\x18\t \x01(\x03R\tstartTime\x12\x19\n" +
	"\bend_time\x18\n" +
	" \x01(\x03R\aendTime\x12\x14\n" +
	"\x05order\x18\v \x01(\x03R\x05order\x12\x1d\n" +
	"\n" +
	"vote_score\x18\x10 \x01(\x05R\tvoteScore\x12 \n" +
	"\vdescription\x18\x14 \x01(\fR\vdescription\x12\x10\n" +
	"\x03url\x18\x15 \x01(\fR\x03url\x12/\n" +
	"\x14free_asset_net_limit\x18\x16 \x01(\x03R\x11freeAssetNetLimit\x12<\n" +
	"\x1bpublic_free_asset_net_limit\x18\x17 \x01(\x03R\x17publicFreeAssetNetLimit\x12<\n" +
	"\x1bpublic_free_asset_net_usage\x18\x18 \x01(\x03R\x17publicFreeAssetNetUsage\x12<\n" +
	"\x1bpublic_latest_free_net_time\x18\x19 \x01(\x03R\x17publicLatestFreeNetTime\x1aT\n" +
	"\fFrozenSupply\x12#\n" +
	"\rfrozen_amount\x18\x01 \x01(\x03R\ffrozenAmount\x12\x1f\n" +
	"\vfrozen_days\x18\x02 \x01(\x03R\n" +
	"frozenDays\"\x92\x01\n" +
	"\x15TransferAssetContract\x12\x1d\n" +
	"\n" +
	"asset_name\x18\x01 \x01(\fR\tassetName\x12#\n" +
	"\rowner_address\x18\x02 \x01(\fR\fownerAddress\x12\x1d\n" +
	"\n" +
	"to_address\x18\x03 \x01(\fR\ttoAddress\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\x03R\x06amount\"<\n" +
	"\x15UnfreezeAssetContract\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\"\xb5\x01\n" +
	"\x13UpdateAssetContract\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\x12 \n" +
	"\vdescription\x18\x02 \x01(\fR\vdescription\x12\x10\n" +
	"\x03url\x18\x03 \x01(\fR\x03url\x12\x1b\n" +
	"\tnew_limit\x18\x04 \x01(\x03R\bnewLimit\x12(\n" +
	"\x10new_public_limit\x18\x05 \x01(\x03R\x0enewPublicLimit\"\x9a\x01\n" +
	"\x1dParticipateAssetIssueContract\x12#\n" +
	"\rowner_address\x18\x01 \x01(\fR\fownerAddress\x12\x1d\n" +
	"\n" +
	"to_address\x18\x02 \x01(\fR\ttoAddress\x12\x1d\n" +
	"\n" +
	"asset_name\x18\x03 \x01(\fR\tassetName\x12\x16\n" +
	"\x06amount\x18\x04 \x01(\x03R\x06amountBN\n" +
	"\x18org.tron.protos.contractZ2github.com/TheByteArray/go-tron-sdk/pkg/proto/coreb\x06proto3"

var (
	file_core_contract_asset_issue_contract_proto_rawDescOnce sync.Once
	file_core_contract_asset_issue_contract_proto_rawDescData []byte
)

func file_core_contract_asset_issue_contract_proto_rawDescGZIP() []byte {
	file_core_contract_asset_issue_contract_proto_rawDescOnce.Do(func() {
		file_core_contract_asset_issue_contract_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_core_contract_asset_issue_contract_proto_rawDesc), len(file_core_contract_asset_issue_contract_proto_rawDesc)))
	})
	return file_core_contract_asset_issue_contract_proto_rawDescData
}

var file_core_contract_asset_issue_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_core_contract_asset_issue_contract_proto_goTypes = []any{
	(*AssetIssueContract)(nil),              // 0: protocol.AssetIssueContract
	(*TransferAssetContract)(nil),           // 1: protocol.TransferAssetContract
	(*UnfreezeAssetContract)(nil),           // 2: protocol.UnfreezeAssetContract
	(*UpdateAssetContract)(nil),             // 3: protocol.UpdateAssetContract
	(*ParticipateAssetIssueContract)(nil),   // 4: protocol.ParticipateAssetIssueContract
	(*AssetIssueContract_FrozenSupply)(nil), // 5: protocol.AssetIssueContract.FrozenSupply
}
var file_core_contract_asset_issue_contract_proto_depIdxs = []int32{
	5, // 0: protocol.AssetIssueContract.frozen_supply:type_name -> protocol.AssetIssueContract.FrozenSupply
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_core_contract_asset_issue_contract_proto_init() }
func file_core_contract_asset_issue_contract_proto_init() {
	if File_core_contract_asset_issue_contract_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_core_contract_asset_issue_contract_proto_rawDesc), len(file_core_contract_asset_issue_contract_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_contract_asset_issue_contract_proto_goTypes,
		DependencyIndexes: file_core_contract_asset_issue_contract_proto_depIdxs,
		MessageInfos:      file_core_contract_asset_issue_contract_proto_msgTypes,
	}.Build()
	File_core_contract_asset_issue_contract_proto = out.File
	file_core_contract_asset_issue_contract_proto_goTypes = nil
	file_core_contract_asset_issue_contract_proto_depIdxs = nil
}
