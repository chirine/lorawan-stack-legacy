// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: discovery.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type Announcement struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceVersion       string   `protobuf:"bytes,3,opt,name=service_version,json=serviceVersion,proto3" json:"service_version,omitempty"`
	NetAddress           string   `protobuf:"bytes,11,opt,name=net_address,json=netAddress,proto3" json:"net_address,omitempty"`
	PublicKey            string   `protobuf:"bytes,12,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Certificate          string   `protobuf:"bytes,13,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announcement) Reset()         { *m = Announcement{} }
func (m *Announcement) String() string { return proto.CompactTextString(m) }
func (*Announcement) ProtoMessage()    {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{1}
}
func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(m, src)
}
func (m *Announcement) XXX_Size() int {
	return m.Size()
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Announcement) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Announcement) GetServiceVersion() string {
	if m != nil {
		return m.ServiceVersion
	}
	return ""
}

func (m *Announcement) GetNetAddress() string {
	if m != nil {
		return m.NetAddress
	}
	return ""
}

func (m *Announcement) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Announcement) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "discovery.GetRequest")
	proto.RegisterType((*Announcement)(nil), "discovery.Announcement")
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor_1e7ff60feb39c8d0) }

var fileDescriptor_1e7ff60feb39c8d0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd1, 0xb1, 0x6a, 0x32, 0x41,
	0x14, 0x05, 0x60, 0x56, 0xe1, 0x87, 0xbd, 0xeb, 0xaf, 0x30, 0x10, 0x32, 0x04, 0x62, 0x8c, 0x4d,
	0xd2, 0xb8, 0x42, 0x24, 0x75, 0x50, 0x02, 0x16, 0x81, 0x14, 0x16, 0x29, 0xd2, 0xc8, 0x38, 0x7b,
	0xb3, 0x0e, 0xab, 0x77, 0x36, 0x33, 0x57, 0x65, 0xdf, 0x30, 0x45, 0x8a, 0x3c, 0x42, 0xf0, 0x49,
	0x02, 0xbb, 0x1a, 0xb7, 0x4e, 0xfb, 0x9d, 0xc3, 0x30, 0x9c, 0x0b, 0x9d, 0xc4, 0x78, 0x6d, 0xb7,
	0xe8, 0x8a, 0x38, 0x77, 0x96, 0xad, 0x08, 0x7f, 0xa1, 0xff, 0x00, 0x30, 0x45, 0x9e, 0xe1, 0xfb,
	0x06, 0x3d, 0x8b, 0x36, 0x34, 0x4c, 0x22, 0x83, 0x5e, 0x70, 0x1b, 0xce, 0x1a, 0x26, 0x11, 0xd7,
	0xd0, 0xf2, 0xe8, 0xb6, 0x46, 0xe3, 0x9c, 0xd4, 0x1a, 0x65, 0xa3, 0x4c, 0xa2, 0x83, 0x3d, 0xab,
	0x35, 0xf6, 0x3f, 0x03, 0x68, 0x8d, 0x89, 0xec, 0x86, 0x34, 0xae, 0x91, 0xfe, 0xf2, 0x86, 0xb8,
	0x81, 0xce, 0xb1, 0xb2, 0x45, 0xe7, 0x8d, 0x25, 0xd9, 0x2c, 0x5b, 0xed, 0x03, 0xbf, 0x54, 0x2a,
	0xae, 0x20, 0x22, 0xe4, 0xb9, 0x4a, 0x12, 0x87, 0xde, 0xcb, 0xa8, 0x2c, 0x01, 0x21, 0x8f, 0x2b,
	0x11, 0x97, 0x00, 0xf9, 0x66, 0xb1, 0x32, 0x7a, 0x9e, 0x61, 0x21, 0x5b, 0x65, 0x1e, 0x56, 0xf2,
	0x84, 0x85, 0xe8, 0x41, 0xa4, 0xd1, 0xb1, 0x79, 0x33, 0x5a, 0x31, 0xca, 0xff, 0xd5, 0x57, 0x6a,
	0x74, 0x37, 0x81, 0xf0, 0xf1, 0x38, 0x8e, 0xb8, 0x87, 0xe6, 0x14, 0x59, 0x9c, 0xc5, 0xa7, 0x01,
	0x4f, 0x63, 0x5d, 0x9c, 0xd7, 0xb8, 0xbe, 0xc0, 0x64, 0xfc, 0xb1, 0xef, 0x06, 0x5f, 0xfb, 0x6e,
	0xf0, 0xbd, 0xef, 0x06, 0xaf, 0xa3, 0xd4, 0xc6, 0xbc, 0x44, 0x5e, 0x1a, 0x4a, 0x7d, 0x4c, 0xc8,
	0x3b, 0xeb, 0xb2, 0xe1, 0xca, 0x3a, 0xb5, 0x53, 0x34, 0xf0, 0xac, 0x74, 0x36, 0x58, 0x61, 0xaa,
	0x74, 0x31, 0xcc, 0xb3, 0x74, 0xc8, 0x4c, 0xf9, 0x62, 0xf1, 0xaf, 0x3c, 0xd4, 0xe8, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x5d, 0x57, 0x09, 0xbb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryClient is the client API for Discovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Announcement, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Announcement, error) {
	out := new(Announcement)
	err := c.cc.Invoke(ctx, "/discovery.Discovery/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServer is the server API for Discovery service.
type DiscoveryServer interface {
	Get(context.Context, *GetRequest) (*Announcement, error)
}

// UnimplementedDiscoveryServer can be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServer struct {
}

func (*UnimplementedDiscoveryServer) Get(ctx context.Context, req *GetRequest) (*Announcement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Discovery_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

func (m *GetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Announcement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Announcement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Announcement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.NetAddress) > 0 {
		i -= len(m.NetAddress)
		copy(dAtA[i:], m.NetAddress)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.NetAddress)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ServiceVersion) > 0 {
		i -= len(m.ServiceVersion)
		copy(dAtA[i:], m.ServiceVersion)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ServiceVersion)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDiscovery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDiscovery(dAtA []byte, offset int, v uint64) int {
	offset -= sovDiscovery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Announcement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.ServiceVersion)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.NetAddress)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovDiscovery(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDiscovery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDiscovery(x uint64) (n int) {
	return sovDiscovery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
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
			return fmt.Errorf("proto: GetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiscovery
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
func (m *Announcement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDiscovery
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
			return fmt.Errorf("proto: Announcement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Announcement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDiscovery
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
				return ErrInvalidLengthDiscovery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDiscovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDiscovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDiscovery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDiscovery
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
func skipDiscovery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDiscovery
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
					return 0, ErrIntOverflowDiscovery
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
					return 0, ErrIntOverflowDiscovery
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
				return 0, ErrInvalidLengthDiscovery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDiscovery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDiscovery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDiscovery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDiscovery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDiscovery = fmt.Errorf("proto: unexpected end of group")
)
