// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	storage.proto

It has these top-level messages:
	UtxoEntry
	Mainchain
*/
package storage

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

type UtxoEntry struct {
	IsCoinBase  bool   `protobuf:"varint,1,opt,name=isCoinBase" json:"isCoinBase,omitempty"`
	BlockHeight uint64 `protobuf:"varint,2,opt,name=blockHeight" json:"blockHeight,omitempty"`
	Spent       bool   `protobuf:"varint,3,opt,name=spent" json:"spent,omitempty"`
}

func (m *UtxoEntry) Reset()                    { *m = UtxoEntry{} }
func (m *UtxoEntry) String() string            { return proto.CompactTextString(m) }
func (*UtxoEntry) ProtoMessage()               {}
func (*UtxoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UtxoEntry) GetIsCoinBase() bool {
	if m != nil {
		return m.IsCoinBase
	}
	return false
}

func (m *UtxoEntry) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *UtxoEntry) GetSpent() bool {
	if m != nil {
		return m.Spent
	}
	return false
}

// Mainchain represents a mainchain of the blockchain
type Mainchain struct {
	Hashs []*Mainchain_Hash `protobuf:"bytes,1,rep,name=hashs" json:"hashs,omitempty"`
}

func (m *Mainchain) Reset()                    { *m = Mainchain{} }
func (m *Mainchain) String() string            { return proto.CompactTextString(m) }
func (*Mainchain) ProtoMessage()               {}
func (*Mainchain) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Mainchain) GetHashs() []*Mainchain_Hash {
	if m != nil {
		return m.Hashs
	}
	return nil
}

type Mainchain_Hash struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *Mainchain_Hash) Reset()                    { *m = Mainchain_Hash{} }
func (m *Mainchain_Hash) String() string            { return proto.CompactTextString(m) }
func (*Mainchain_Hash) ProtoMessage()               {}
func (*Mainchain_Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Mainchain_Hash) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*UtxoEntry)(nil), "chain.core.txdb.internal.storage.UtxoEntry")
	proto.RegisterType((*Mainchain)(nil), "chain.core.txdb.internal.storage.Mainchain")
	proto.RegisterType((*Mainchain_Hash)(nil), "chain.core.txdb.internal.storage.Mainchain.Hash")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x89, 0xdd, 0x55, 0x3b, 0xab, 0x20, 0xc1, 0x43, 0xf0, 0x20, 0x61, 0x4f, 0x3d, 0x0d,
	0xa2, 0xff, 0x60, 0x45, 0xd9, 0x8b, 0x97, 0x80, 0x17, 0x6f, 0x69, 0x0c, 0x9b, 0xb0, 0x35, 0x29,
	0xc9, 0x1c, 0xda, 0x7f, 0x2f, 0x4d, 0x8b, 0xf4, 0xe6, 0x6d, 0xde, 0x63, 0x3e, 0xf8, 0x78, 0x70,
	0x9b, 0x29, 0x26, 0x7d, 0xb2, 0xd8, 0xa7, 0x48, 0x91, 0x4b, 0xe3, 0xb4, 0x0f, 0x68, 0x62, 0xb2,
	0x48, 0xc3, 0x77, 0x8b, 0x3e, 0x90, 0x4d, 0x41, 0x77, 0xb8, 0xfc, 0xed, 0x0d, 0xd4, 0x9f, 0x34,
	0xc4, 0xb7, 0x40, 0x69, 0xe4, 0x8f, 0x00, 0x3e, 0xbf, 0x46, 0x1f, 0x0e, 0x3a, 0x5b, 0xc1, 0x24,
	0x6b, 0xae, 0xd5, 0xaa, 0xe1, 0x12, 0x76, 0x6d, 0x17, 0xcd, 0xf9, 0x68, 0xfd, 0xc9, 0x91, 0xb8,
	0x90, 0xac, 0xd9, 0xa8, 0x75, 0xc5, 0xef, 0x61, 0x9b, 0x7b, 0x1b, 0x48, 0x54, 0x05, 0x9e, 0xc3,
	0xfe, 0x07, 0xea, 0x0f, 0xed, 0x43, 0x91, 0xe1, 0xef, 0xb0, 0x75, 0x3a, 0xbb, 0x2c, 0x98, 0xac,
	0x9a, 0xdd, 0xf3, 0x13, 0xfe, 0xe7, 0x88, 0x7f, 0x2c, 0x1e, 0x75, 0x76, 0x6a, 0xc6, 0x1f, 0x04,
	0x6c, 0xa6, 0xc8, 0xef, 0xa0, 0x3a, 0xdb, 0xb1, 0xd8, 0xde, 0xa8, 0xe9, 0x3c, 0xd4, 0x5f, 0x57,
	0x0b, 0xda, 0x5e, 0x96, 0x1d, 0x5e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xca, 0xf0, 0xed, 0x5a,
	0x18, 0x01, 0x00, 0x00,
}
