// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zgc/dasigners/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	TokensPerVote     string `protobuf:"bytes,1,opt,name=tokens_per_vote,json=tokensPerVote,proto3" json:"tokens_per_vote,omitempty"`
	MaxVotesPerSigner uint64 `protobuf:"varint,2,opt,name=max_votes_per_signer,json=maxVotesPerSigner,proto3" json:"max_votes_per_signer,omitempty"`
	MaxQuorums        uint64 `protobuf:"varint,3,opt,name=max_quorums,json=maxQuorums,proto3" json:"max_quorums,omitempty"`
	EpochBlocks       uint64 `protobuf:"varint,4,opt,name=epoch_blocks,json=epochBlocks,proto3" json:"epoch_blocks,omitempty"`
	EncodedSlices     uint64 `protobuf:"varint,5,opt,name=encoded_slices,json=encodedSlices,proto3" json:"encoded_slices,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetTokensPerVote() string {
	if m != nil {
		return m.TokensPerVote
	}
	return ""
}

func (m *Params) GetMaxVotesPerSigner() uint64 {
	if m != nil {
		return m.MaxVotesPerSigner
	}
	return 0
}

func (m *Params) GetMaxQuorums() uint64 {
	if m != nil {
		return m.MaxQuorums
	}
	return 0
}

func (m *Params) GetEpochBlocks() uint64 {
	if m != nil {
		return m.EpochBlocks
	}
	return 0
}

func (m *Params) GetEncodedSlices() uint64 {
	if m != nil {
		return m.EncodedSlices
	}
	return 0
}

// GenesisState defines the dasigners module's genesis state.
type GenesisState struct {
	// params defines all the parameters of related to deposit.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// params epoch_number the epoch number
	EpochNumber uint64 `protobuf:"varint,2,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	// signers defines all signers information
	Signers []*Signer `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
	// quorums_by_epoch defines chosen quorums by epoch
	QuorumsByEpoch []*Quorums `protobuf:"bytes,4,rep,name=quorums_by_epoch,json=quorumsByEpoch,proto3" json:"quorums_by_epoch,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_896efa766aaca3be, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetEpochNumber() uint64 {
	if m != nil {
		return m.EpochNumber
	}
	return 0
}

func (m *GenesisState) GetSigners() []*Signer {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *GenesisState) GetQuorumsByEpoch() []*Quorums {
	if m != nil {
		return m.QuorumsByEpoch
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "zgc.dasigners.v1.Params")
	proto.RegisterType((*GenesisState)(nil), "zgc.dasigners.v1.GenesisState")
}

func init() { proto.RegisterFile("zgc/dasigners/v1/genesis.proto", fileDescriptor_896efa766aaca3be) }

var fileDescriptor_896efa766aaca3be = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x40, 0xb3, 0x24, 0x04, 0xe1, 0xb4, 0xa5, 0x58, 0x3d, 0x38, 0x3d, 0x6c, 0x42, 0x25, 0x50,
	0x2f, 0xac, 0xdb, 0x22, 0xf1, 0x01, 0x41, 0x08, 0x71, 0x41, 0x65, 0x23, 0x71, 0xe0, 0xb2, 0xf2,
	0x3a, 0xc6, 0x59, 0x35, 0xde, 0x59, 0xd6, 0xde, 0x28, 0xe9, 0x57, 0xf0, 0x59, 0x3d, 0x70, 0xe8,
	0x91, 0x13, 0x42, 0xc9, 0x8f, 0xa0, 0x1d, 0x9b, 0x56, 0xb4, 0xbd, 0xd9, 0x6f, 0xde, 0x8c, 0x66,
	0xc6, 0x26, 0xf1, 0xa5, 0x96, 0x7c, 0x26, 0x6c, 0xa1, 0x4b, 0x55, 0x5b, 0xbe, 0x3c, 0xe5, 0x5a,
	0x95, 0xca, 0x16, 0x36, 0xa9, 0x6a, 0x70, 0x40, 0xf7, 0x2f, 0xb5, 0x4c, 0x6e, 0xe2, 0xc9, 0xf2,
	0xf4, 0x70, 0x28, 0xc1, 0x1a, 0xb0, 0x19, 0xc6, 0xb9, 0xbf, 0x78, 0xf9, 0xf0, 0x40, 0x83, 0x06,
	0xcf, 0xdb, 0x53, 0xa0, 0x43, 0x0d, 0xa0, 0x17, 0x8a, 0xe3, 0x2d, 0x6f, 0xbe, 0x71, 0x51, 0xae,
	0x43, 0x68, 0x74, 0x37, 0xe4, 0x0a, 0xa3, 0xac, 0x13, 0xa6, 0x0a, 0xc2, 0xf8, 0x5e, 0x7b, 0xb7,
	0xbd, 0xa0, 0x71, 0xf4, 0x33, 0x22, 0xfd, 0x73, 0x51, 0x0b, 0x63, 0xe9, 0x2b, 0xf2, 0xcc, 0xc1,
	0x85, 0x2a, 0x6d, 0x56, 0xa9, 0x3a, 0x5b, 0x82, 0x53, 0x2c, 0x1a, 0x47, 0xc7, 0x4f, 0xd3, 0x5d,
	0x8f, 0xcf, 0x55, 0xfd, 0x05, 0x9c, 0xa2, 0x9c, 0x1c, 0x18, 0xb1, 0x42, 0xc1, 0xab, 0xbe, 0x22,
	0x7b, 0x34, 0x8e, 0x8e, 0x7b, 0xe9, 0x73, 0x23, 0x56, 0xad, 0xd6, 0xea, 0x53, 0x0c, 0xd0, 0x11,
	0x19, 0xb4, 0x09, 0xdf, 0x1b, 0xa8, 0x1b, 0x63, 0x59, 0x17, 0x3d, 0x62, 0xc4, 0xea, 0xb3, 0x27,
	0xf4, 0x05, 0xd9, 0x51, 0x15, 0xc8, 0x79, 0x96, 0x2f, 0x40, 0x5e, 0x58, 0xd6, 0x43, 0x63, 0x80,
	0x6c, 0x82, 0x88, 0xbe, 0x24, 0x7b, 0xaa, 0x94, 0x30, 0x53, 0xb3, 0xcc, 0x2e, 0x0a, 0xa9, 0x2c,
	0x7b, 0x8c, 0xd2, 0x6e, 0xa0, 0x53, 0x84, 0x47, 0x9b, 0x88, 0xec, 0x7c, 0xf0, 0x2f, 0x30, 0x75,
	0xc2, 0x29, 0xfa, 0x96, 0xf4, 0x2b, 0x1c, 0x0f, 0x67, 0x19, 0x9c, 0xb1, 0xe4, 0xee, 0x8b, 0x24,
	0x7e, 0xfc, 0x49, 0xef, 0xea, 0xf7, 0xa8, 0x93, 0x06, 0xfb, 0xb6, 0xa5, 0xb2, 0x31, 0xf9, 0xcd,
	0x70, 0xbe, 0xa5, 0x4f, 0x88, 0xe8, 0x19, 0x79, 0x12, 0xaa, 0xb0, 0xee, 0xb8, 0xfb, 0x70, 0x6d,
	0xbf, 0x81, 0xf4, 0x9f, 0x48, 0xdf, 0x91, 0xfd, 0xb0, 0x86, 0x2c, 0x5f, 0x67, 0x58, 0x8d, 0xf5,
	0x30, 0x79, 0x78, 0x3f, 0x39, 0xac, 0x27, 0xdd, 0x0b, 0x29, 0x93, 0xf5, 0x7b, 0xdc, 0xc8, 0xc7,
	0xab, 0x4d, 0x1c, 0x5d, 0x6f, 0xe2, 0xe8, 0xcf, 0x26, 0x8e, 0x7e, 0x6c, 0xe3, 0xce, 0xf5, 0x36,
	0xee, 0xfc, 0xda, 0xc6, 0x9d, 0xaf, 0x5c, 0x17, 0x6e, 0xde, 0xe4, 0x89, 0x04, 0xc3, 0x4f, 0xf4,
	0x42, 0xe4, 0x96, 0x9f, 0xe8, 0xd7, 0x72, 0x2e, 0x8a, 0x92, 0xaf, 0xfe, 0xff, 0x08, 0x6e, 0x5d,
	0x29, 0x9b, 0xf7, 0xf1, 0x17, 0xbc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x79, 0x90, 0xf6, 0x00,
	0xc8, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EncodedSlices != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EncodedSlices))
		i--
		dAtA[i] = 0x28
	}
	if m.EpochBlocks != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochBlocks))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxQuorums != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxQuorums))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxVotesPerSigner != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxVotesPerSigner))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TokensPerVote) > 0 {
		i -= len(m.TokensPerVote)
		copy(dAtA[i:], m.TokensPerVote)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokensPerVote)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QuorumsByEpoch) > 0 {
		for iNdEx := len(m.QuorumsByEpoch) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QuorumsByEpoch[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Signers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.EpochNumber != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochNumber))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokensPerVote)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxVotesPerSigner != 0 {
		n += 1 + sovGenesis(uint64(m.MaxVotesPerSigner))
	}
	if m.MaxQuorums != 0 {
		n += 1 + sovGenesis(uint64(m.MaxQuorums))
	}
	if m.EpochBlocks != 0 {
		n += 1 + sovGenesis(uint64(m.EpochBlocks))
	}
	if m.EncodedSlices != 0 {
		n += 1 + sovGenesis(uint64(m.EncodedSlices))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochNumber != 0 {
		n += 1 + sovGenesis(uint64(m.EpochNumber))
	}
	if len(m.Signers) > 0 {
		for _, e := range m.Signers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.QuorumsByEpoch) > 0 {
		for _, e := range m.QuorumsByEpoch {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensPerVote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokensPerVote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxVotesPerSigner", wireType)
			}
			m.MaxVotesPerSigner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxVotesPerSigner |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxQuorums", wireType)
			}
			m.MaxQuorums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxQuorums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocks", wireType)
			}
			m.EpochBlocks = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocks |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncodedSlices", wireType)
			}
			m.EncodedSlices = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EncodedSlices |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, &Signer{})
			if err := m.Signers[len(m.Signers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuorumsByEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuorumsByEpoch = append(m.QuorumsByEpoch, &Quorums{})
			if err := m.QuorumsByEpoch[len(m.QuorumsByEpoch)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
