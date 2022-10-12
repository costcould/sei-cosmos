// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/accesscontrol.proto

package accesscontrol

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type AccessOperation struct {
	AccessType         AccessType   `protobuf:"varint,1,opt,name=access_type,json=accessType,proto3,enum=cosmos.accesscontrol.v1beta1.AccessType" json:"access_type,omitempty"`
	ResourceType       ResourceType `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=cosmos.accesscontrol.v1beta1.ResourceType" json:"resource_type,omitempty"`
	IdentifierTemplate string       `protobuf:"bytes,3,opt,name=identifier_template,json=identifierTemplate,proto3" json:"identifier_template,omitempty"`
}

func (m *AccessOperation) Reset()         { *m = AccessOperation{} }
func (m *AccessOperation) String() string { return proto.CompactTextString(m) }
func (*AccessOperation) ProtoMessage()    {}
func (*AccessOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{0}
}
func (m *AccessOperation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessOperation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessOperation.Merge(m, src)
}
func (m *AccessOperation) XXX_Size() int {
	return m.Size()
}
func (m *AccessOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AccessOperation proto.InternalMessageInfo

func (m *AccessOperation) GetAccessType() AccessType {
	if m != nil {
		return m.AccessType
	}
	return AccessType_UNKNOWN
}

func (m *AccessOperation) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_ANY
}

func (m *AccessOperation) GetIdentifierTemplate() string {
	if m != nil {
		return m.IdentifierTemplate
	}
	return ""
}

type MessageDependencyMapping struct {
	MessageKey     string            `protobuf:"bytes,1,opt,name=message_key,json=messageKey,proto3" json:"message_key,omitempty"`
	AccessOps      []AccessOperation `protobuf:"bytes,2,rep,name=access_ops,json=accessOps,proto3" json:"access_ops"`
	DynamicEnabled bool              `protobuf:"varint,3,opt,name=dynamic_enabled,json=dynamicEnabled,proto3" json:"dynamic_enabled,omitempty"`
}

func (m *MessageDependencyMapping) Reset()         { *m = MessageDependencyMapping{} }
func (m *MessageDependencyMapping) String() string { return proto.CompactTextString(m) }
func (*MessageDependencyMapping) ProtoMessage()    {}
func (*MessageDependencyMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{1}
}
func (m *MessageDependencyMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageDependencyMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageDependencyMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageDependencyMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageDependencyMapping.Merge(m, src)
}
func (m *MessageDependencyMapping) XXX_Size() int {
	return m.Size()
}
func (m *MessageDependencyMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageDependencyMapping.DiscardUnknown(m)
}

var xxx_messageInfo_MessageDependencyMapping proto.InternalMessageInfo

func (m *MessageDependencyMapping) GetMessageKey() string {
	if m != nil {
		return m.MessageKey
	}
	return ""
}

func (m *MessageDependencyMapping) GetAccessOps() []AccessOperation {
	if m != nil {
		return m.AccessOps
	}
	return nil
}

func (m *MessageDependencyMapping) GetDynamicEnabled() bool {
	if m != nil {
		return m.DynamicEnabled
	}
	return false
}

type WasmFunctionDependencyMapping struct {
	WasmFunction string            `protobuf:"bytes,1,opt,name=wasm_function,json=wasmFunction,proto3" json:"wasm_function,omitempty"`
	Enabled      bool              `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AccessOps    []AccessOperation `protobuf:"bytes,3,rep,name=access_ops,json=accessOps,proto3" json:"access_ops"`
}

func (m *WasmFunctionDependencyMapping) Reset()         { *m = WasmFunctionDependencyMapping{} }
func (m *WasmFunctionDependencyMapping) String() string { return proto.CompactTextString(m) }
func (*WasmFunctionDependencyMapping) ProtoMessage()    {}
func (*WasmFunctionDependencyMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{2}
}
func (m *WasmFunctionDependencyMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WasmFunctionDependencyMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WasmFunctionDependencyMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WasmFunctionDependencyMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmFunctionDependencyMapping.Merge(m, src)
}
func (m *WasmFunctionDependencyMapping) XXX_Size() int {
	return m.Size()
}
func (m *WasmFunctionDependencyMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmFunctionDependencyMapping.DiscardUnknown(m)
}

var xxx_messageInfo_WasmFunctionDependencyMapping proto.InternalMessageInfo

func (m *WasmFunctionDependencyMapping) GetWasmFunction() string {
	if m != nil {
		return m.WasmFunction
	}
	return ""
}

func (m *WasmFunctionDependencyMapping) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *WasmFunctionDependencyMapping) GetAccessOps() []AccessOperation {
	if m != nil {
		return m.AccessOps
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessOperation)(nil), "cosmos.accesscontrol.v1beta1.AccessOperation")
	proto.RegisterType((*MessageDependencyMapping)(nil), "cosmos.accesscontrol.v1beta1.MessageDependencyMapping")
	proto.RegisterType((*WasmFunctionDependencyMapping)(nil), "cosmos.accesscontrol.v1beta1.WasmFunctionDependencyMapping")
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/accesscontrol.proto", fileDescriptor_d636a082612ba091)
}

var fileDescriptor_d636a082612ba091 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xdd, 0x8a, 0x13, 0x31,
	0x14, 0x9e, 0xb4, 0xa2, 0x36, 0xdd, 0x1f, 0x88, 0x5e, 0x0c, 0x8b, 0xce, 0x96, 0x2a, 0x38, 0x08,
	0x3b, 0xe3, 0xae, 0x4f, 0xe0, 0xa2, 0x82, 0xca, 0xb2, 0x30, 0x2c, 0x08, 0xde, 0x0c, 0x99, 0xcc,
	0xd9, 0x71, 0x68, 0x27, 0x09, 0x93, 0xd4, 0x32, 0x6f, 0xe1, 0xb3, 0x78, 0xe5, 0x23, 0xf4, 0xb2,
	0x97, 0x5e, 0x15, 0x69, 0x5f, 0x44, 0x9a, 0xa4, 0xb5, 0xad, 0x52, 0xbc, 0xf0, 0x2a, 0x39, 0xe7,
	0x7c, 0xdf, 0x97, 0xf3, 0x85, 0x0f, 0x87, 0x4c, 0xa8, 0x4a, 0xa8, 0x98, 0x32, 0x06, 0x4a, 0x31,
	0xc1, 0x75, 0x2d, 0x86, 0xdb, 0x55, 0x24, 0x6b, 0xa1, 0x05, 0x79, 0x64, 0x91, 0xd1, 0xf6, 0xec,
	0xcb, 0x79, 0x06, 0x9a, 0x9e, 0x9f, 0x3c, 0x2c, 0x44, 0x21, 0x0c, 0x30, 0x5e, 0xde, 0x2c, 0xe7,
	0xe4, 0xe9, 0x5f, 0xd5, 0x99, 0xe0, 0x4a, 0x53, 0xae, 0x95, 0x45, 0xf5, 0x67, 0x08, 0x1f, 0xbf,
	0x32, 0x88, 0x6b, 0x09, 0x35, 0xd5, 0xa5, 0xe0, 0xe4, 0x1d, 0xee, 0x5a, 0x52, 0xaa, 0x1b, 0x09,
	0x3e, 0xea, 0xa1, 0xf0, 0xe8, 0x22, 0x8c, 0xf6, 0xed, 0x10, 0x59, 0x8d, 0x9b, 0x46, 0x42, 0x82,
	0xe9, 0xfa, 0x4e, 0xae, 0xf1, 0x61, 0x0d, 0x4a, 0x8c, 0x6a, 0x06, 0x56, 0xac, 0x65, 0xc4, 0x9e,
	0xef, 0x17, 0x4b, 0x1c, 0xc5, 0xc8, 0x1d, 0xd4, 0x1b, 0x15, 0x89, 0xf1, 0x83, 0x32, 0x07, 0xae,
	0xcb, 0xdb, 0x12, 0xea, 0x54, 0x43, 0x25, 0x87, 0x54, 0x83, 0xdf, 0xee, 0xa1, 0xb0, 0x93, 0x90,
	0xdf, 0xa3, 0x1b, 0x37, 0xe9, 0x7f, 0x47, 0xd8, 0xbf, 0x02, 0xa5, 0x68, 0x01, 0xaf, 0x41, 0x02,
	0xcf, 0x81, 0xb3, 0xe6, 0x8a, 0x4a, 0x59, 0xf2, 0x82, 0x9c, 0xe2, 0x6e, 0x65, 0x67, 0xe9, 0x00,
	0x1a, 0xe3, 0xb4, 0x93, 0x60, 0xd7, 0xfa, 0x00, 0x0d, 0x49, 0xb0, 0x73, 0x93, 0x0a, 0xa9, 0xfc,
	0x56, 0xaf, 0x1d, 0x76, 0x2f, 0xce, 0xfe, 0xe5, 0x27, 0xd6, 0xbf, 0x79, 0x79, 0x67, 0x32, 0x3b,
	0xf5, 0x92, 0x0e, 0x75, 0x6d, 0x45, 0x9e, 0xe1, 0xe3, 0xbc, 0xe1, 0xb4, 0x2a, 0x59, 0x0a, 0x9c,
	0x66, 0x43, 0xc8, 0xcd, 0xfa, 0xf7, 0x93, 0x23, 0xd7, 0x7e, 0x63, 0xbb, 0xfd, 0x6f, 0x08, 0x3f,
	0xfe, 0x48, 0x55, 0xf5, 0x76, 0xc4, 0xd9, 0x52, 0xea, 0xcf, 0xfd, 0x9f, 0xe0, 0xc3, 0x31, 0x55,
	0x55, 0x7a, 0xeb, 0x10, 0xce, 0xc1, 0xc1, 0x78, 0x83, 0x45, 0x7c, 0x7c, 0x6f, 0xf5, 0x4e, 0xcb,
	0xbc, 0xb3, 0x2a, 0x77, 0xdc, 0xb5, 0xff, 0x87, 0xbb, 0xcb, 0xf7, 0x93, 0x79, 0x80, 0xa6, 0xf3,
	0x00, 0xfd, 0x9c, 0x07, 0xe8, 0xeb, 0x22, 0xf0, 0xa6, 0x8b, 0xc0, 0xfb, 0xb1, 0x08, 0xbc, 0x4f,
	0x2f, 0x8a, 0x52, 0x7f, 0x1e, 0x65, 0x11, 0x13, 0x55, 0xec, 0xb2, 0x69, 0x8f, 0x33, 0x95, 0x0f,
	0xe2, 0x65, 0x3a, 0x76, 0xc2, 0x9a, 0xdd, 0x35, 0x19, 0x7d, 0xf9, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xf2, 0xdc, 0x22, 0xf3, 0x29, 0x03, 0x00, 0x00,
}

func (m *AccessOperation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessOperation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessOperation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdentifierTemplate) > 0 {
		i -= len(m.IdentifierTemplate)
		copy(dAtA[i:], m.IdentifierTemplate)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.IdentifierTemplate)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ResourceType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.ResourceType))
		i--
		dAtA[i] = 0x10
	}
	if m.AccessType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.AccessType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MessageDependencyMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageDependencyMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageDependencyMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DynamicEnabled {
		i--
		if m.DynamicEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.AccessOps) > 0 {
		for iNdEx := len(m.AccessOps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccessOps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccesscontrol(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MessageKey) > 0 {
		i -= len(m.MessageKey)
		copy(dAtA[i:], m.MessageKey)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.MessageKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WasmFunctionDependencyMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WasmFunctionDependencyMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WasmFunctionDependencyMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccessOps) > 0 {
		for iNdEx := len(m.AccessOps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccessOps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccesscontrol(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.WasmFunction) > 0 {
		i -= len(m.WasmFunction)
		copy(dAtA[i:], m.WasmFunction)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.WasmFunction)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccesscontrol(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccesscontrol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessOperation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.AccessType))
	}
	if m.ResourceType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.ResourceType))
	}
	l = len(m.IdentifierTemplate)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	return n
}

func (m *MessageDependencyMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageKey)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	if len(m.AccessOps) > 0 {
		for _, e := range m.AccessOps {
			l = e.Size()
			n += 1 + l + sovAccesscontrol(uint64(l))
		}
	}
	if m.DynamicEnabled {
		n += 2
	}
	return n
}

func (m *WasmFunctionDependencyMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WasmFunction)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.AccessOps) > 0 {
		for _, e := range m.AccessOps {
			l = e.Size()
			n += 1 + l + sovAccesscontrol(uint64(l))
		}
	}
	return n
}

func sovAccesscontrol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccesscontrol(x uint64) (n int) {
	return sovAccesscontrol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessOperation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: AccessOperation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessOperation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessType", wireType)
			}
			m.AccessType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessType |= AccessType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			m.ResourceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceType |= ResourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentifierTemplate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentifierTemplate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func (m *MessageDependencyMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: MessageDependencyMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageDependencyMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessOps = append(m.AccessOps, AccessOperation{})
			if err := m.AccessOps[len(m.AccessOps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DynamicEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DynamicEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func (m *WasmFunctionDependencyMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: WasmFunctionDependencyMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WasmFunctionDependencyMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmFunction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WasmFunction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessOps = append(m.AccessOps, AccessOperation{})
			if err := m.AccessOps[len(m.AccessOps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func skipAccesscontrol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccesscontrol
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
					return 0, ErrIntOverflowAccesscontrol
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
					return 0, ErrIntOverflowAccesscontrol
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
				return 0, ErrInvalidLengthAccesscontrol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccesscontrol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccesscontrol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccesscontrol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccesscontrol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccesscontrol = fmt.Errorf("proto: unexpected end of group")
)