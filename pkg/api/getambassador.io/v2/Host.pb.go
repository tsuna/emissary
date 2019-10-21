// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Host.proto

package Host

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// The Host resource will usually be a Kubernetes CRD, but it could
// appear in other forms. The HostSpec is the part of the Host resource
// that doesn't change, no matter what form it's in -- when it's a CRD,
// this is the part in the "spec" dictionary.
type HostSpec struct {
	// Common to all Ambassador objects (and optional).
	AmbassadorId []string `protobuf:"bytes,1,rep,name=ambassador_id,json=ambassadorId,proto3" json:"ambassador_id,omitempty"`
	// Common to all Ambassador objects (and optional).
	Generation int32 `protobuf:"varint,2,opt,name=generation,proto3" json:"generation,omitempty"`
	// Hostname by which the Ambassador can be reached.
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Selector by which we can find further configuration. Defaults to hostname=$hostname
	Selector *v1.LabelSelector `protobuf:"bytes,4,opt,name=selector,proto3" json:"selector,omitempty"`
	// Specifies who to talk ACME with to get certs. Defaults to Let's Encrypt; if "none", do
	// not try to do TLS for this Host.
	AcmeProvider string `protobuf:"bytes,5,opt,name=acme_provider,json=acmeProvider,proto3" json:"acme_provider,omitempty"`
	// Name of the Kubernetes secret into which to save generated certificates. Defaults
	// to $hostname
	SecretName           string   `protobuf:"bytes,6,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostSpec) Reset()         { *m = HostSpec{} }
func (m *HostSpec) String() string { return proto.CompactTextString(m) }
func (*HostSpec) ProtoMessage()    {}
func (*HostSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d8e6496365725b, []int{0}
}
func (m *HostSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostSpec.Merge(m, src)
}
func (m *HostSpec) XXX_Size() int {
	return m.Size()
}
func (m *HostSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_HostSpec.DiscardUnknown(m)
}

var xxx_messageInfo_HostSpec proto.InternalMessageInfo

func (m *HostSpec) GetAmbassadorId() []string {
	if m != nil {
		return m.AmbassadorId
	}
	return nil
}

func (m *HostSpec) GetGeneration() int32 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *HostSpec) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostSpec) GetSelector() *v1.LabelSelector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *HostSpec) GetAcmeProvider() string {
	if m != nil {
		return m.AcmeProvider
	}
	return ""
}

func (m *HostSpec) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func init() {
	proto.RegisterType((*HostSpec)(nil), "HostSpec")
}

func init() { proto.RegisterFile("Host.proto", fileDescriptor_88d8e6496365725b) }

var fileDescriptor_88d8e6496365725b = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0xd6, 0x96, 0x76, 0x1a, 0x2f, 0x39, 0x85, 0x1e, 0x62, 0xd0, 0x4b, 0x4e, 0x1b,
	0x6a, 0x3d, 0x78, 0xf6, 0xa4, 0x20, 0x2a, 0xe9, 0x03, 0x84, 0x49, 0x32, 0xa4, 0x4b, 0xba, 0xd9,
	0x65, 0x77, 0x09, 0x78, 0xf5, 0xe9, 0x3c, 0xfa, 0x08, 0x92, 0x27, 0x91, 0xfc, 0xa1, 0xf6, 0x38,
	0x3f, 0xe6, 0xfb, 0xe6, 0xc7, 0x00, 0x3c, 0x2b, 0xeb, 0xb8, 0x36, 0xca, 0xa9, 0xed, 0x43, 0xfd,
	0x68, 0xb9, 0x50, 0x09, 0x6a, 0x21, 0xb1, 0x38, 0x8a, 0x86, 0xcc, 0x67, 0xa2, 0xeb, 0xaa, 0x07,
	0x36, 0x91, 0xe4, 0x30, 0x69, 0x77, 0x49, 0x45, 0x0d, 0x19, 0x74, 0x54, 0x8e, 0xa9, 0xdb, 0xaf,
	0x19, 0xac, 0xfa, 0x92, 0x83, 0xa6, 0xc2, 0xbf, 0x83, 0x6b, 0x94, 0x39, 0x5a, 0x8b, 0xa5, 0x32,
	0x99, 0x28, 0x03, 0x16, 0xcd, 0xe3, 0x75, 0xea, 0xfd, 0xc3, 0x97, 0xd2, 0x0f, 0x01, 0xa6, 0x12,
	0xa1, 0x9a, 0x60, 0x16, 0xb1, 0x78, 0x91, 0x5e, 0x10, 0x7f, 0x0b, 0xab, 0xa3, 0xb2, 0xae, 0x41,
	0x49, 0xc1, 0x3c, 0x62, 0xf1, 0x3a, 0x3d, 0xcf, 0xfe, 0x3b, 0xac, 0x2c, 0x9d, 0xa8, 0x70, 0xca,
	0x04, 0x57, 0x11, 0x8b, 0x37, 0xf7, 0x7b, 0x3e, 0x6a, 0xf3, 0x4b, 0x6d, 0xae, 0xeb, 0xaa, 0x07,
	0x96, 0xf7, 0xda, 0xbc, 0xdd, 0xf1, 0x57, 0xcc, 0xe9, 0x74, 0x98, 0xa2, 0xe9, 0xb9, 0x64, 0x30,
	0x2e, 0x24, 0x65, 0xda, 0xa8, 0x56, 0x94, 0x64, 0x82, 0xc5, 0x70, 0xd1, 0xeb, 0xe1, 0xc7, 0xc4,
	0xfc, 0x1b, 0xd8, 0x58, 0x2a, 0x0c, 0xb9, 0x6c, 0x90, 0x5a, 0x0e, 0x2b, 0x30, 0xa2, 0x37, 0x94,
	0xf4, 0xe4, 0x7d, 0x77, 0x21, 0xfb, 0xe9, 0x42, 0xf6, 0xdb, 0x85, 0x2c, 0x5f, 0x0e, 0x9f, 0xd9,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x90, 0x73, 0x45, 0x5d, 0x01, 0x00, 0x00,
}

func (m *HostSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.SecretName) > 0 {
		i -= len(m.SecretName)
		copy(dAtA[i:], m.SecretName)
		i = encodeVarintHost(dAtA, i, uint64(len(m.SecretName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.AcmeProvider) > 0 {
		i -= len(m.AcmeProvider)
		copy(dAtA[i:], m.AcmeProvider)
		i = encodeVarintHost(dAtA, i, uint64(len(m.AcmeProvider)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Selector != nil {
		{
			size, err := m.Selector.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHost(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Hostname) > 0 {
		i -= len(m.Hostname)
		copy(dAtA[i:], m.Hostname)
		i = encodeVarintHost(dAtA, i, uint64(len(m.Hostname)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Generation != 0 {
		i = encodeVarintHost(dAtA, i, uint64(m.Generation))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AmbassadorId) > 0 {
		for iNdEx := len(m.AmbassadorId) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AmbassadorId[iNdEx])
			copy(dAtA[i:], m.AmbassadorId[iNdEx])
			i = encodeVarintHost(dAtA, i, uint64(len(m.AmbassadorId[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	offset -= sovHost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HostSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AmbassadorId) > 0 {
		for _, s := range m.AmbassadorId {
			l = len(s)
			n += 1 + l + sovHost(uint64(l))
		}
	}
	if m.Generation != 0 {
		n += 1 + sovHost(uint64(m.Generation))
	}
	l = len(m.Hostname)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.AcmeProvider)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.SecretName)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmbassadorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmbassadorId = append(m.AmbassadorId, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Generation", wireType)
			}
			m.Generation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Generation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hostname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hostname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &v1.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcmeProvider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AcmeProvider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecretName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecretName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHost
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHost
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHost(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHost
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHost = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost   = fmt.Errorf("proto: integer overflow")
)
