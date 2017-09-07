// Code generated by protoc-gen-go.
// source: bc.proto
// DO NOT EDIT!

/*
Package bc is a generated protocol buffer package.

It is generated from these files:
	bc.proto

It has these top-level messages:
	Hash
	Program
	AssetID
	AssetAmount
	AssetDefinition
	ValueSource
	ValueDestination
	BlockHeader
	TxHeader
	Mux
	Nonce
	Output
	Retirement
	Issuance
	Spend
*/
package bc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Hash struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hash) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *Hash) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *Hash) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *Hash) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type Program struct {
	VmVersion uint64 `protobuf:"varint,1,opt,name=vm_version,json=vmVersion" json:"vm_version,omitempty"`
	Code      []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *Program) Reset()                    { *m = Program{} }
func (m *Program) String() string            { return proto.CompactTextString(m) }
func (*Program) ProtoMessage()               {}
func (*Program) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Program) GetVmVersion() uint64 {
	if m != nil {
		return m.VmVersion
	}
	return 0
}

func (m *Program) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

// This message type duplicates Hash, above. One alternative is to
// embed a Hash inside an AssetID. But it's useful for AssetID to be
// plain old data (without pointers). Another alternative is use Hash
// in any protobuf types where an AssetID is called for, but it's
// preferable to have type safety.
type AssetID struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *AssetID) Reset()                    { *m = AssetID{} }
func (m *AssetID) String() string            { return proto.CompactTextString(m) }
func (*AssetID) ProtoMessage()               {}
func (*AssetID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AssetID) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *AssetID) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *AssetID) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *AssetID) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

type AssetAmount struct {
	AssetId *AssetID `protobuf:"bytes,1,opt,name=asset_id,json=assetId" json:"asset_id,omitempty"`
	Amount  uint64   `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *AssetAmount) Reset()                    { *m = AssetAmount{} }
func (m *AssetAmount) String() string            { return proto.CompactTextString(m) }
func (*AssetAmount) ProtoMessage()               {}
func (*AssetAmount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AssetAmount) GetAssetId() *AssetID {
	if m != nil {
		return m.AssetId
	}
	return nil
}

func (m *AssetAmount) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AssetDefinition struct {
	InitialBlockId  *Hash    `protobuf:"bytes,1,opt,name=initial_block_id,json=initialBlockId" json:"initial_block_id,omitempty"`
	IssuanceProgram *Program `protobuf:"bytes,2,opt,name=issuance_program,json=issuanceProgram" json:"issuance_program,omitempty"`
	Data            *Hash    `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *AssetDefinition) Reset()                    { *m = AssetDefinition{} }
func (m *AssetDefinition) String() string            { return proto.CompactTextString(m) }
func (*AssetDefinition) ProtoMessage()               {}
func (*AssetDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AssetDefinition) GetInitialBlockId() *Hash {
	if m != nil {
		return m.InitialBlockId
	}
	return nil
}

func (m *AssetDefinition) GetIssuanceProgram() *Program {
	if m != nil {
		return m.IssuanceProgram
	}
	return nil
}

func (m *AssetDefinition) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

type ValueSource struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueSource) Reset()                    { *m = ValueSource{} }
func (m *ValueSource) String() string            { return proto.CompactTextString(m) }
func (*ValueSource) ProtoMessage()               {}
func (*ValueSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ValueSource) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueSource) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueSource) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type ValueDestination struct {
	Ref      *Hash        `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	Value    *AssetAmount `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Position uint64       `protobuf:"varint,3,opt,name=position" json:"position,omitempty"`
}

func (m *ValueDestination) Reset()                    { *m = ValueDestination{} }
func (m *ValueDestination) String() string            { return proto.CompactTextString(m) }
func (*ValueDestination) ProtoMessage()               {}
func (*ValueDestination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ValueDestination) GetRef() *Hash {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *ValueDestination) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ValueDestination) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

type BlockHeader struct {
	Version          uint64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SerializedSize   uint64 `protobuf:"varint,2,opt,name=serialized_size,json=serializedSize" json:"serialized_size,omitempty"`
	Height           uint64 `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	PreviousBlockId  *Hash  `protobuf:"bytes,4,opt,name=previous_block_id,json=previousBlockId" json:"previous_block_id,omitempty"`
	TimestampMs      uint64 `protobuf:"varint,5,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	TransactionsRoot *Hash  `protobuf:"bytes,6,opt,name=transactions_root,json=transactionsRoot" json:"transactions_root,omitempty"`
	AssetsRoot       *Hash  `protobuf:"bytes,7,opt,name=assets_root,json=assetsRoot" json:"assets_root,omitempty"`
	Nonce            uint64 `protobuf:"varint,8,opt,name=nonce" json:"nonce,omitempty"`
	Bits             uint64 `protobuf:"varint,9,opt,name=bits" json:"bits,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BlockHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockId() *Hash {
	if m != nil {
		return m.PreviousBlockId
	}
	return nil
}

func (m *BlockHeader) GetTimestampMs() uint64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

func (m *BlockHeader) GetTransactionsRoot() *Hash {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *BlockHeader) GetAssetsRoot() *Hash {
	if m != nil {
		return m.AssetsRoot
	}
	return nil
}

type TxHeader struct {
	Version        uint64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	SerializedSize uint64  `protobuf:"varint,2,opt,name=serialized_size,json=serializedSize" json:"serialized_size,omitempty"`
	ResultIds      []*Hash `protobuf:"bytes,3,rep,name=result_ids,json=resultIds" json:"result_ids,omitempty"`
	Data           *Hash   `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
	MinTimeMs      uint64  `protobuf:"varint,5,opt,name=min_time_ms,json=minTimeMs" json:"min_time_ms,omitempty"`
	MaxTimeMs      uint64  `protobuf:"varint,6,opt,name=max_time_ms,json=maxTimeMs" json:"max_time_ms,omitempty"`
	ExtHash        *Hash   `protobuf:"bytes,7,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
}

func (m *TxHeader) Reset()                    { *m = TxHeader{} }
func (m *TxHeader) String() string            { return proto.CompactTextString(m) }
func (*TxHeader) ProtoMessage()               {}
func (*TxHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *TxHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TxHeader) GetResultIds() []*Hash {
	if m != nil {
		return m.ResultIds
	}
	return nil
}

func (m *TxHeader) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TxHeader) GetMinTimeMs() uint64 {
	if m != nil {
		return m.MinTimeMs
	}
	return 0
}

func (m *TxHeader) GetMaxTimeMs() uint64 {
	if m != nil {
		return m.MaxTimeMs
	}
	return 0
}

func (m *TxHeader) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

type Mux struct {
	Sources             []*ValueSource      `protobuf:"bytes,1,rep,name=sources" json:"sources,omitempty"`
	Program             *Program            `protobuf:"bytes,2,opt,name=program" json:"program,omitempty"`
	ExtHash             *Hash               `protobuf:"bytes,3,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessDestinations []*ValueDestination `protobuf:"bytes,4,rep,name=witness_destinations,json=witnessDestinations" json:"witness_destinations,omitempty"`
	WitnessArguments    [][]byte            `protobuf:"bytes,5,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
}

func (m *Mux) Reset()                    { *m = Mux{} }
func (m *Mux) String() string            { return proto.CompactTextString(m) }
func (*Mux) ProtoMessage()               {}
func (*Mux) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Mux) GetSources() []*ValueSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Mux) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Mux) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Mux) GetWitnessDestinations() []*ValueDestination {
	if m != nil {
		return m.WitnessDestinations
	}
	return nil
}

func (m *Mux) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

type Nonce struct {
	Program           *Program `protobuf:"bytes,1,opt,name=program" json:"program,omitempty"`
	ExtHash           *Hash    `protobuf:"bytes,2,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessArguments  [][]byte `protobuf:"bytes,3,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	WitnessAnchoredId *Hash    `protobuf:"bytes,4,opt,name=witness_anchored_id,json=witnessAnchoredId" json:"witness_anchored_id,omitempty"`
}

func (m *Nonce) Reset()                    { *m = Nonce{} }
func (m *Nonce) String() string            { return proto.CompactTextString(m) }
func (*Nonce) ProtoMessage()               {}
func (*Nonce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Nonce) GetProgram() *Program {
	if m != nil {
		return m.Program
	}
	return nil
}

func (m *Nonce) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Nonce) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Nonce) GetWitnessAnchoredId() *Hash {
	if m != nil {
		return m.WitnessAnchoredId
	}
	return nil
}

type Coinbase struct {
	WitnessDestination *ValueDestination `protobuf:"bytes,1,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
}

func (m *Coinbase) Reset()                    { *m = Coinbase{} }
func (m *Coinbase) String() string            { return proto.CompactTextString(m) }
func (*Coinbase) ProtoMessage()               {}
func (*Coinbase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Coinbase) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

type Output struct {
	Source         *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	ControlProgram *Program     `protobuf:"bytes,2,opt,name=control_program,json=controlProgram" json:"control_program,omitempty"`
	Data           *Hash        `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	ExtHash        *Hash        `protobuf:"bytes,4,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	Ordinal        uint64       `protobuf:"varint,5,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Output) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Output) GetControlProgram() *Program {
	if m != nil {
		return m.ControlProgram
	}
	return nil
}

func (m *Output) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Output) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Output) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Retirement struct {
	Source  *ValueSource `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Data    *Hash        `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	ExtHash *Hash        `protobuf:"bytes,3,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	Ordinal uint64       `protobuf:"varint,4,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Retirement) Reset()                    { *m = Retirement{} }
func (m *Retirement) String() string            { return proto.CompactTextString(m) }
func (*Retirement) ProtoMessage()               {}
func (*Retirement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Retirement) GetSource() *ValueSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Retirement) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Retirement) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Retirement) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Issuance struct {
	AnchorId               *Hash             `protobuf:"bytes,1,opt,name=anchor_id,json=anchorId" json:"anchor_id,omitempty"`
	Value                  *AssetAmount      `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Data                   *Hash             `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	ExtHash                *Hash             `protobuf:"bytes,4,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessDestination     *ValueDestination `protobuf:"bytes,5,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessAssetDefinition *AssetDefinition  `protobuf:"bytes,6,opt,name=witness_asset_definition,json=witnessAssetDefinition" json:"witness_asset_definition,omitempty"`
	WitnessArguments       [][]byte          `protobuf:"bytes,7,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	WitnessAnchoredId      *Hash             `protobuf:"bytes,8,opt,name=witness_anchored_id,json=witnessAnchoredId" json:"witness_anchored_id,omitempty"`
	Ordinal                uint64            `protobuf:"varint,9,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Issuance) Reset()                    { *m = Issuance{} }
func (m *Issuance) String() string            { return proto.CompactTextString(m) }
func (*Issuance) ProtoMessage()               {}
func (*Issuance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Issuance) GetAnchorId() *Hash {
	if m != nil {
		return m.AnchorId
	}
	return nil
}

func (m *Issuance) GetValue() *AssetAmount {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Issuance) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Issuance) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Issuance) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Issuance) GetWitnessAssetDefinition() *AssetDefinition {
	if m != nil {
		return m.WitnessAssetDefinition
	}
	return nil
}

func (m *Issuance) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Issuance) GetWitnessAnchoredId() *Hash {
	if m != nil {
		return m.WitnessAnchoredId
	}
	return nil
}

func (m *Issuance) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

type Spend struct {
	SpentOutputId      *Hash             `protobuf:"bytes,1,opt,name=spent_output_id,json=spentOutputId" json:"spent_output_id,omitempty"`
	Data               *Hash             `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	ExtHash            *Hash             `protobuf:"bytes,3,opt,name=ext_hash,json=extHash" json:"ext_hash,omitempty"`
	WitnessDestination *ValueDestination `protobuf:"bytes,4,opt,name=witness_destination,json=witnessDestination" json:"witness_destination,omitempty"`
	WitnessArguments   [][]byte          `protobuf:"bytes,5,rep,name=witness_arguments,json=witnessArguments,proto3" json:"witness_arguments,omitempty"`
	WitnessAnchoredId  *Hash             `protobuf:"bytes,6,opt,name=witness_anchored_id,json=witnessAnchoredId" json:"witness_anchored_id,omitempty"`
	Ordinal            uint64            `protobuf:"varint,7,opt,name=ordinal" json:"ordinal,omitempty"`
}

func (m *Spend) Reset()                    { *m = Spend{} }
func (m *Spend) String() string            { return proto.CompactTextString(m) }
func (*Spend) ProtoMessage()               {}
func (*Spend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Spend) GetSpentOutputId() *Hash {
	if m != nil {
		return m.SpentOutputId
	}
	return nil
}

func (m *Spend) GetData() *Hash {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Spend) GetExtHash() *Hash {
	if m != nil {
		return m.ExtHash
	}
	return nil
}

func (m *Spend) GetWitnessDestination() *ValueDestination {
	if m != nil {
		return m.WitnessDestination
	}
	return nil
}

func (m *Spend) GetWitnessArguments() [][]byte {
	if m != nil {
		return m.WitnessArguments
	}
	return nil
}

func (m *Spend) GetWitnessAnchoredId() *Hash {
	if m != nil {
		return m.WitnessAnchoredId
	}
	return nil
}

func (m *Spend) GetOrdinal() uint64 {
	if m != nil {
		return m.Ordinal
	}
	return 0
}

func init() {
	proto.RegisterType((*Hash)(nil), "bc.Hash")
	proto.RegisterType((*Program)(nil), "bc.Program")
	proto.RegisterType((*AssetID)(nil), "bc.AssetID")
	proto.RegisterType((*AssetAmount)(nil), "bc.AssetAmount")
	proto.RegisterType((*AssetDefinition)(nil), "bc.AssetDefinition")
	proto.RegisterType((*ValueSource)(nil), "bc.ValueSource")
	proto.RegisterType((*ValueDestination)(nil), "bc.ValueDestination")
	proto.RegisterType((*BlockHeader)(nil), "bc.BlockHeader")
	proto.RegisterType((*TxHeader)(nil), "bc.TxHeader")
	proto.RegisterType((*Mux)(nil), "bc.Mux")
	proto.RegisterType((*Nonce)(nil), "bc.Nonce")
	proto.RegisterType((*Output)(nil), "bc.Output")
	proto.RegisterType((*Retirement)(nil), "bc.Retirement")
	proto.RegisterType((*Coinbase)(nil), "bc.Coinbase")
	proto.RegisterType((*Issuance)(nil), "bc.Issuance")
	proto.RegisterType((*Spend)(nil), "bc.Spend")
}

func init() { proto.RegisterFile("bc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdf, 0x6e, 0xe3, 0xc4,
	0x17, 0x96, 0x13, 0x27, 0x76, 0x4e, 0xba, 0x4d, 0x3a, 0xad, 0x2a, 0x6b, 0xf5, 0xfb, 0xa1, 0x62,
	0x54, 0x76, 0x57, 0xa0, 0xaa, 0xdb, 0x16, 0xc4, 0x05, 0x37, 0x85, 0x02, 0xeb, 0x8b, 0x00, 0xf2,
	0x56, 0x7b, 0x6b, 0x4d, 0xec, 0xd9, 0xc6, 0x22, 0x9e, 0x09, 0x9e, 0x71, 0xc8, 0x35, 0x8f, 0xc0,
	0x2d, 0x6f, 0xc1, 0x33, 0xec, 0x03, 0xf0, 0x18, 0x5c, 0xf3, 0x04, 0x68, 0x8e, 0xc7, 0xce, 0x9f,
	0x26, 0x69, 0x2a, 0x96, 0xbb, 0x9c, 0x39, 0x67, 0xce, 0x9f, 0xef, 0x7c, 0x9f, 0x33, 0xe0, 0x0e,
	0xe3, 0xb3, 0x49, 0x2e, 0x94, 0x20, 0x8d, 0x61, 0xec, 0x7f, 0x0b, 0xf6, 0x2b, 0x2a, 0x47, 0x64,
	0x1f, 0x1a, 0xd3, 0x73, 0xcf, 0x3a, 0xb1, 0x9e, 0xb7, 0xc3, 0xc6, 0xf4, 0x1c, 0xed, 0x97, 0x5e,
	0xc3, 0xd8, 0x2f, 0xd1, 0xbe, 0xf0, 0x9a, 0xc6, 0xbe, 0x40, 0xfb, 0xd2, 0xb3, 0x8d, 0x7d, 0xe9,
	0x7f, 0x09, 0xce, 0x8f, 0xb9, 0xb8, 0xcb, 0x69, 0x46, 0xfe, 0x0f, 0x30, 0xcd, 0xa2, 0x29, 0xcb,
	0x65, 0x2a, 0x38, 0xa6, 0xb4, 0xc3, 0xce, 0x34, 0x7b, 0x53, 0x1e, 0x10, 0x02, 0x76, 0x2c, 0x12,
	0x86, 0xb9, 0xf7, 0x42, 0xfc, 0xed, 0x07, 0xe0, 0x5c, 0x4b, 0xc9, 0x54, 0x70, 0xf3, 0xaf, 0x1b,
	0x19, 0x40, 0x17, 0x53, 0x5d, 0x67, 0xa2, 0xe0, 0x8a, 0x7c, 0x0c, 0x2e, 0xd5, 0x66, 0x94, 0x26,
	0x98, 0xb4, 0x7b, 0xd1, 0x3d, 0x1b, 0xc6, 0x67, 0xa6, 0x5a, 0xe8, 0xa0, 0x33, 0x48, 0xc8, 0x31,
	0xb4, 0x29, 0xde, 0xc0, 0x52, 0x76, 0x68, 0x2c, 0xff, 0x77, 0x0b, 0x7a, 0x18, 0x7c, 0xc3, 0xde,
	0xa6, 0x3c, 0x55, 0x7a, 0x82, 0x0b, 0xe8, 0xe3, 0x4f, 0x3a, 0x8e, 0x86, 0x63, 0x11, 0xff, 0x34,
	0xcf, 0xed, 0xea, 0xdc, 0x1a, 0xcf, 0x70, 0xdf, 0x44, 0x7c, 0xa5, 0x03, 0x82, 0x84, 0x7c, 0x0e,
	0xfd, 0x54, 0xca, 0x82, 0xf2, 0x98, 0x45, 0x93, 0x12, 0x28, 0xac, 0x64, 0xfa, 0x31, 0xd8, 0x85,
	0xbd, 0x2a, 0xa8, 0x02, 0xf3, 0x7f, 0x60, 0x27, 0x54, 0x51, 0x1c, 0x78, 0x31, 0x3f, 0x9e, 0xfa,
	0x63, 0xe8, 0xbe, 0xa1, 0xe3, 0x82, 0xbd, 0x16, 0x45, 0x1e, 0x33, 0xf2, 0x14, 0x9a, 0x39, 0x7b,
	0x7b, 0xaf, 0x17, 0x7d, 0x48, 0x4e, 0xa1, 0x35, 0xd5, 0xa1, 0xa6, 0x6a, 0xaf, 0x46, 0xa1, 0x04,
	0x2a, 0x2c, 0xbd, 0xe4, 0x29, 0xb8, 0x13, 0x21, 0x71, 0x4e, 0xac, 0x69, 0x87, 0xb5, 0xed, 0xff,
	0x0c, 0x7d, 0xac, 0x76, 0xc3, 0xa4, 0x4a, 0x39, 0x45, 0x2c, 0xfe, 0xe3, 0x92, 0xbf, 0x36, 0xa1,
	0x8b, 0x10, 0xbe, 0x62, 0x34, 0x61, 0x39, 0xf1, 0xc0, 0x59, 0x26, 0x56, 0x65, 0xea, 0x05, 0x8e,
	0x58, 0x7a, 0x37, 0xaa, 0x17, 0x58, 0x5a, 0xe4, 0x0a, 0x0e, 0x26, 0x39, 0x9b, 0xa6, 0xa2, 0x90,
	0xf3, 0x6d, 0xad, 0xa2, 0xd9, 0xab, 0x42, 0xaa, 0x75, 0x7d, 0x08, 0x7b, 0x2a, 0xcd, 0x98, 0x54,
	0x34, 0x9b, 0x44, 0x99, 0x44, 0x7e, 0xd9, 0x61, 0xb7, 0x3e, 0x1b, 0x48, 0xf2, 0x19, 0x1c, 0xa8,
	0x9c, 0x72, 0x49, 0x63, 0xdd, 0xa9, 0x8c, 0x72, 0x21, 0x94, 0xd7, 0x5a, 0x49, 0xdc, 0x5f, 0x0c,
	0x09, 0x85, 0x50, 0xe4, 0x05, 0x74, 0x91, 0x73, 0xe6, 0x42, 0x7b, 0xe5, 0x02, 0x94, 0x4e, 0x0c,
	0xbd, 0x82, 0x63, 0xce, 0x66, 0x2a, 0x8a, 0x05, 0x97, 0x8c, 0xcb, 0x42, 0xd6, 0xcc, 0x71, 0x50,
	0x3b, 0x47, 0xda, 0xfb, 0x75, 0xe5, 0xac, 0x18, 0xf3, 0x11, 0xb8, 0xfa, 0xd2, 0x88, 0xca, 0x91,
	0xe7, 0xae, 0x64, 0x77, 0xd8, 0x4c, 0xa1, 0xdc, 0x3f, 0x81, 0x83, 0x5f, 0x52, 0xc5, 0x99, 0x94,
	0x11, 0xcd, 0xef, 0x8a, 0x8c, 0x71, 0x25, 0xbd, 0xce, 0x49, 0xf3, 0xf9, 0x5e, 0xd8, 0x37, 0x8e,
	0xeb, 0xea, 0xdc, 0xff, 0xd3, 0x02, 0xf7, 0x76, 0xf6, 0xe0, 0x06, 0x9e, 0x01, 0xe4, 0x4c, 0x16,
	0x63, 0xad, 0x35, 0xe9, 0x35, 0x4e, 0x9a, 0x4b, 0xa5, 0x3b, 0xa5, 0x2f, 0x48, 0xe4, 0x76, 0x4e,
	0x93, 0x0f, 0xa0, 0x9b, 0xa5, 0x3c, 0xd2, 0x50, 0xcf, 0x91, 0xef, 0x64, 0x29, 0xbf, 0x4d, 0x33,
	0x36, 0x90, 0xe8, 0xa7, 0xb3, 0xda, 0xdf, 0x32, 0x7e, 0x3a, 0x33, 0xfe, 0xc5, 0xf9, 0xdb, 0x1b,
	0xe6, 0xf7, 0xff, 0xb6, 0xa0, 0x39, 0x28, 0x66, 0xe4, 0x05, 0x38, 0x12, 0xb5, 0x23, 0x3d, 0x0b,
	0x1b, 0x46, 0x92, 0x2e, 0x68, 0x2a, 0xac, 0xfc, 0xe4, 0x14, 0x9c, 0x2d, 0xc2, 0xad, 0x7c, 0x4b,
	0xe5, 0x9b, 0x9b, 0xe0, 0xff, 0x0e, 0x8e, 0x2a, 0xf8, 0x93, 0xb9, 0x98, 0xf4, 0xb0, 0xba, 0x87,
	0xa3, 0xba, 0x87, 0x05, 0xa5, 0x85, 0x87, 0xe6, 0xc6, 0xc2, 0x99, 0x5c, 0xbf, 0xc7, 0xd6, 0x86,
	0x3d, 0xfe, 0x65, 0x41, 0xeb, 0x7b, 0xc1, 0x63, 0xb6, 0x38, 0x8b, 0xb5, 0x65, 0x96, 0x4f, 0xe1,
	0x09, 0xc2, 0x9c, 0x53, 0x7e, 0xc7, 0xb4, 0x6e, 0x1a, 0x2b, 0x03, 0xa1, 0x20, 0x42, 0xed, 0x0d,
	0x92, 0xdd, 0x26, 0x5f, 0xdb, 0xb0, 0xbd, 0xbe, 0x61, 0xf2, 0x05, 0x1c, 0xd6, 0xc1, 0x3c, 0x1e,
	0x89, 0x9c, 0x25, 0xba, 0x8b, 0x55, 0x91, 0x55, 0x19, 0xaf, 0x4d, 0x4c, 0x90, 0xf8, 0xef, 0x2c,
	0x68, 0xff, 0x50, 0xa8, 0x49, 0xa1, 0xc8, 0x33, 0x68, 0x97, 0x2b, 0x34, 0xa3, 0xde, 0xdb, 0xb0,
	0x71, 0x93, 0x2b, 0xe8, 0xc5, 0x82, 0xab, 0x5c, 0x8c, 0xb7, 0x7d, 0xa1, 0xf7, 0x4d, 0xcc, 0x4e,
	0x1f, 0xe8, 0x25, 0x4c, 0xec, 0x4d, 0x98, 0x78, 0xe0, 0x88, 0x3c, 0x49, 0x39, 0x1d, 0x1b, 0x36,
	0x57, 0xa6, 0xff, 0x9b, 0x05, 0x10, 0x32, 0x95, 0xe6, 0x4c, 0x03, 0xb2, 0xfb, 0x28, 0x55, 0x53,
	0x8d, 0x07, 0x9b, 0x6a, 0xee, 0xd0, 0x94, 0xbd, 0xdc, 0xd4, 0x04, 0x3a, 0xb7, 0xd5, 0xda, 0x57,
	0xd5, 0x6a, 0x3d, 0xa0, 0xd6, 0xc6, 0x36, 0xb5, 0x6e, 0xea, 0xc5, 0xff, 0xa3, 0x09, 0x6e, 0x60,
	0xfe, 0x18, 0xc9, 0x29, 0x74, 0x4a, 0x32, 0xac, 0xfb, 0xdb, 0x75, 0x4b, 0x57, 0x90, 0xec, 0xfa,
	0xe7, 0xf3, 0x1e, 0xd6, 0xf7, 0xcd, 0x9c, 0xa5, 0x0b, 0x62, 0x36, 0x2c, 0x5d, 0xaf, 0x65, 0x72,
	0x5f, 0xcb, 0x64, 0x00, 0x5e, 0x4d, 0x76, 0x7c, 0xb1, 0x24, 0xf5, 0x8b, 0xc3, 0x7c, 0xc7, 0x0e,
	0xeb, 0x19, 0xe6, 0x8f, 0x91, 0xf0, 0xb8, 0x22, 0xff, 0xca, 0x23, 0x65, 0xad, 0xd0, 0x9c, 0xc7,
	0x09, 0xcd, 0x7d, 0x50, 0x68, 0x8b, 0x34, 0xe9, 0x2c, 0xd3, 0xe4, 0x5d, 0x03, 0x5a, 0xaf, 0x27,
	0x8c, 0x27, 0xe4, 0x1c, 0x7a, 0x72, 0xc2, 0xb8, 0x8a, 0x04, 0x2a, 0x72, 0xdd, 0xde, 0x9e, 0x60,
	0x40, 0xa9, 0xd8, 0x20, 0x79, 0x1f, 0xfc, 0xdd, 0xb0, 0x15, 0xfb, 0x91, 0x5b, 0x79, 0xcc, 0x07,
	0x76, 0x13, 0x8c, 0xed, 0x47, 0xc1, 0xe8, 0x2c, 0xc1, 0x38, 0x6c, 0xe3, 0x5b, 0xfd, 0xf2, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0x32, 0x78, 0x25, 0xb7, 0x0b, 0x00, 0x00,
}
