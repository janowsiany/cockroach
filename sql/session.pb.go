// Code generated by protoc-gen-gogo.
// source: cockroach/sql/session.proto
// DO NOT EDIT!

package sql

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/roachpb"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import github_com_cockroachdb_cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Session struct {
	Database string `protobuf:"bytes,1,opt,name=database" json:"database"`
	Syntax   int32  `protobuf:"varint,2,opt,name=syntax" json:"syntax"`
	// Info about the open transaction (if any).
	Txn Session_Transaction `protobuf:"bytes,3,opt,name=txn" json:"txn"`
	// Types that are valid to be assigned to Timezone:
	//	*Session_Location
	//	*Session_Offset
	Timezone              isSession_Timezone               `protobuf_oneof:"timezone"`
	DefaultIsolationLevel cockroach_roachpb1.IsolationType `protobuf:"varint,7,opt,name=default_isolation_level,json=defaultIsolationLevel,enum=cockroach.roachpb.IsolationType" json:"default_isolation_level"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{0} }

type isSession_Timezone interface {
	isSession_Timezone()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Session_Location struct {
	Location string `protobuf:"bytes,5,opt,name=location,oneof"`
}
type Session_Offset struct {
	Offset int64 `protobuf:"varint,6,opt,name=offset,oneof"`
}

func (*Session_Location) isSession_Timezone() {}
func (*Session_Offset) isSession_Timezone()   {}

func (m *Session) GetTimezone() isSession_Timezone {
	if m != nil {
		return m.Timezone
	}
	return nil
}

func (m *Session) GetLocation() string {
	if x, ok := m.GetTimezone().(*Session_Location); ok {
		return x.Location
	}
	return ""
}

func (m *Session) GetOffset() int64 {
	if x, ok := m.GetTimezone().(*Session_Offset); ok {
		return x.Offset
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Session) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Session_OneofMarshaler, _Session_OneofUnmarshaler, _Session_OneofSizer, []interface{}{
		(*Session_Location)(nil),
		(*Session_Offset)(nil),
	}
}

func _Session_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Session)
	// timezone
	switch x := m.Timezone.(type) {
	case *Session_Location:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Location)
	case *Session_Offset:
		_ = b.EncodeVarint(6<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Offset))
	case nil:
	default:
		return fmt.Errorf("Session.Timezone has unexpected type %T", x)
	}
	return nil
}

func _Session_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Session)
	switch tag {
	case 5: // timezone.location
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Timezone = &Session_Location{x}
		return true, err
	case 6: // timezone.offset
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Timezone = &Session_Offset{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Session_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Session)
	// timezone
	switch x := m.Timezone.(type) {
	case *Session_Location:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Location)))
		n += len(x.Location)
	case *Session_Offset:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Offset))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Timestamp represents an absolute timestamp devoid of time-zone.
type Session_Timestamp struct {
	// The time in seconds since, January 1, 1970 UTC (Unix time).
	Sec int64 `protobuf:"varint,1,opt,name=sec" json:"sec"`
	// nsec specifies a non-negative nanosecond offset within sec.
	// It must be in the range [0, 999999999].
	Nsec uint32 `protobuf:"varint,2,opt,name=nsec" json:"nsec"`
}

func (m *Session_Timestamp) Reset()                    { *m = Session_Timestamp{} }
func (m *Session_Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Session_Timestamp) ProtoMessage()               {}
func (*Session_Timestamp) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{0, 0} }

type Session_Transaction struct {
	// If missing, it means we're not inside a (KV) txn.
	Txn *cockroach_roachpb1.Transaction `protobuf:"bytes,1,opt,name=txn" json:"txn,omitempty"`
	// txnAborted is set once executing a statement returned an error from KV.
	// While in this state, every subsequent statement must be rejected until
	// a COMMIT/ROLLBACK is seen.
	TxnAborted   bool                                                  `protobuf:"varint,2,opt,name=txnAborted" json:"txnAborted"`
	UserPriority github_com_cockroachdb_cockroach_roachpb.UserPriority `protobuf:"fixed64,3,opt,name=user_priority,json=userPriority,casttype=github.com/cockroachdb/cockroach/roachpb.UserPriority" json:"user_priority"`
	// Indicates that the transaction is mutating keys in the
	// SystemConfig span.
	MutatesSystemConfig bool `protobuf:"varint,4,opt,name=mutates_system_config,json=mutatesSystemConfig" json:"mutates_system_config"`
}

func (m *Session_Transaction) Reset()                    { *m = Session_Transaction{} }
func (m *Session_Transaction) String() string            { return proto.CompactTextString(m) }
func (*Session_Transaction) ProtoMessage()               {}
func (*Session_Transaction) Descriptor() ([]byte, []int) { return fileDescriptorSession, []int{0, 1} }

func init() {
	proto.RegisterType((*Session)(nil), "cockroach.sql.Session")
	proto.RegisterType((*Session_Timestamp)(nil), "cockroach.sql.Session.Timestamp")
	proto.RegisterType((*Session_Transaction)(nil), "cockroach.sql.Session.Transaction")
}
func (m *Session) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Session) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintSession(data, i, uint64(len(m.Database)))
	i += copy(data[i:], m.Database)
	data[i] = 0x10
	i++
	i = encodeVarintSession(data, i, uint64(m.Syntax))
	data[i] = 0x1a
	i++
	i = encodeVarintSession(data, i, uint64(m.Txn.Size()))
	n1, err := m.Txn.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Timezone != nil {
		nn2, err := m.Timezone.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	data[i] = 0x38
	i++
	i = encodeVarintSession(data, i, uint64(m.DefaultIsolationLevel))
	return i, nil
}

func (m *Session_Location) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x2a
	i++
	i = encodeVarintSession(data, i, uint64(len(m.Location)))
	i += copy(data[i:], m.Location)
	return i, nil
}
func (m *Session_Offset) MarshalTo(data []byte) (int, error) {
	i := 0
	data[i] = 0x30
	i++
	i = encodeVarintSession(data, i, uint64(m.Offset))
	return i, nil
}
func (m *Session_Timestamp) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Session_Timestamp) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSession(data, i, uint64(m.Sec))
	data[i] = 0x10
	i++
	i = encodeVarintSession(data, i, uint64(m.Nsec))
	return i, nil
}

func (m *Session_Transaction) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Session_Transaction) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Txn != nil {
		data[i] = 0xa
		i++
		i = encodeVarintSession(data, i, uint64(m.Txn.Size()))
		n3, err := m.Txn.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	data[i] = 0x10
	i++
	if m.TxnAborted {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	data[i] = 0x19
	i++
	i = encodeFixed64Session(data, i, uint64(math.Float64bits(float64(m.UserPriority))))
	data[i] = 0x20
	i++
	if m.MutatesSystemConfig {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func encodeFixed64Session(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Session(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSession(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Session) Size() (n int) {
	var l int
	_ = l
	l = len(m.Database)
	n += 1 + l + sovSession(uint64(l))
	n += 1 + sovSession(uint64(m.Syntax))
	l = m.Txn.Size()
	n += 1 + l + sovSession(uint64(l))
	if m.Timezone != nil {
		n += m.Timezone.Size()
	}
	n += 1 + sovSession(uint64(m.DefaultIsolationLevel))
	return n
}

func (m *Session_Location) Size() (n int) {
	var l int
	_ = l
	l = len(m.Location)
	n += 1 + l + sovSession(uint64(l))
	return n
}
func (m *Session_Offset) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSession(uint64(m.Offset))
	return n
}
func (m *Session_Timestamp) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSession(uint64(m.Sec))
	n += 1 + sovSession(uint64(m.Nsec))
	return n
}

func (m *Session_Transaction) Size() (n int) {
	var l int
	_ = l
	if m.Txn != nil {
		l = m.Txn.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	n += 2
	n += 9
	n += 2
	return n
}

func sovSession(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Session) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Database = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Syntax", wireType)
			}
			m.Syntax = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Syntax |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Txn.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timezone = &Session_Location{string(data[iNdEx:postIndex])}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Timezone = &Session_Offset{v}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultIsolationLevel", wireType)
			}
			m.DefaultIsolationLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DefaultIsolationLevel |= (cockroach_roachpb1.IsolationType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSession(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
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
func (m *Session_Timestamp) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sec", wireType)
			}
			m.Sec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Sec |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nsec", wireType)
			}
			m.Nsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Nsec |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSession(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
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
func (m *Session_Transaction) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Txn == nil {
				m.Txn = &cockroach_roachpb1.Transaction{}
			}
			if err := m.Txn.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnAborted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TxnAborted = bool(v != 0)
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserPriority", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(data[iNdEx-8])
			v |= uint64(data[iNdEx-7]) << 8
			v |= uint64(data[iNdEx-6]) << 16
			v |= uint64(data[iNdEx-5]) << 24
			v |= uint64(data[iNdEx-4]) << 32
			v |= uint64(data[iNdEx-3]) << 40
			v |= uint64(data[iNdEx-2]) << 48
			v |= uint64(data[iNdEx-1]) << 56
			m.UserPriority = github_com_cockroachdb_cockroach_roachpb.UserPriority(math.Float64frombits(v))
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutatesSystemConfig", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MutatesSystemConfig = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSession(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSession
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
func skipSession(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowSession
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSession
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSession
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipSession(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthSession = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorSession = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x6f, 0xd3, 0x30,
	0x1c, 0x6d, 0x96, 0xae, 0xeb, 0x7e, 0xa3, 0x1c, 0x0c, 0x03, 0xab, 0x8c, 0xae, 0x9a, 0x38, 0x70,
	0x4a, 0xd0, 0x24, 0x24, 0x84, 0xb4, 0x03, 0xe5, 0x02, 0x12, 0x07, 0xb4, 0x15, 0x09, 0x71, 0x20,
	0x72, 0x52, 0x27, 0x8b, 0x48, 0xe2, 0x10, 0x3b, 0x68, 0xe1, 0xca, 0x17, 0xe0, 0x63, 0xf5, 0xc8,
	0x91, 0xd3, 0xc4, 0x9f, 0x3b, 0x1f, 0x80, 0x13, 0x3f, 0x3b, 0x6e, 0x08, 0x62, 0x07, 0xb7, 0xf6,
	0x7b, 0xcf, 0xbf, 0xf7, 0xfa, 0x6a, 0xb8, 0x13, 0x89, 0xe8, 0x5d, 0x25, 0x58, 0x74, 0xee, 0xcb,
	0xf7, 0x99, 0x2f, 0xb9, 0x94, 0xa9, 0x28, 0xbc, 0xb2, 0x12, 0x4a, 0x90, 0x49, 0x47, 0x7a, 0x48,
	0x4e, 0x0f, 0xfe, 0x6a, 0xcd, 0x67, 0x19, 0xfa, 0x2b, 0xa6, 0x58, 0x2b, 0x9e, 0xde, 0x4c, 0x44,
	0x22, 0xcc, 0xd6, 0xd7, 0xbb, 0x16, 0x3d, 0xfa, 0x35, 0x84, 0x9d, 0xb3, 0x76, 0x28, 0x99, 0xc3,
	0x58, 0xeb, 0x43, 0x26, 0x39, 0x75, 0xe6, 0xce, 0xfd, 0xdd, 0xc5, 0x70, 0x7d, 0x79, 0x38, 0x38,
	0xed, 0x50, 0x72, 0x00, 0x23, 0xd9, 0x14, 0x8a, 0x5d, 0xd0, 0x2d, 0xe4, 0xb7, 0x2d, 0x6f, 0x31,
	0xf2, 0x18, 0x5c, 0x75, 0x51, 0x50, 0x17, 0xa9, 0xbd, 0xe3, 0x23, 0xef, 0x9f, 0x70, 0x9e, 0x35,
	0xf1, 0x96, 0x15, 0x2b, 0x24, 0x8b, 0x14, 0xee, 0xed, 0x75, 0x7d, 0x09, 0x27, 0x8f, 0x33, 0x11,
	0x31, 0x0d, 0xd3, 0x6d, 0xed, 0xfd, 0x0c, 0x7d, 0x37, 0x08, 0xa1, 0x30, 0x12, 0x71, 0x2c, 0xb9,
	0xa2, 0x23, 0xe4, 0x5c, 0xe4, 0xec, 0x99, 0xbc, 0x85, 0xdb, 0x2b, 0x1e, 0xb3, 0x3a, 0x53, 0x41,
	0x2a, 0x45, 0x66, 0xe4, 0x41, 0xc6, 0x3f, 0xf0, 0x8c, 0xee, 0xa0, 0xf4, 0xfa, 0xf1, 0xbc, 0x97,
	0xc3, 0xb6, 0xe2, 0x3d, 0xdf, 0x28, 0x97, 0x4d, 0xc9, 0x6d, 0x8a, 0x7d, 0x3b, 0xa6, 0xe3, 0x5e,
	0xe8, 0x21, 0xd3, 0x13, 0xd8, 0x5d, 0xa6, 0x39, 0x97, 0x8a, 0xe5, 0x25, 0xb9, 0x05, 0xae, 0xe4,
	0x91, 0xe9, 0xc6, 0xdd, 0x84, 0x47, 0x00, 0xe3, 0x0d, 0x0b, 0x4d, 0xe8, 0x52, 0x26, 0x96, 0x30,
	0xc8, 0xf4, 0xd3, 0x16, 0xec, 0xf5, 0x7e, 0x31, 0x79, 0xd0, 0x56, 0xe4, 0x98, 0x8a, 0x66, 0x57,
	0x44, 0xeb, 0x89, 0xdb, 0x62, 0xee, 0x01, 0xe0, 0xd7, 0x93, 0x50, 0x54, 0x8a, 0xaf, 0x8c, 0xc3,
	0xd8, 0x3a, 0xf4, 0x70, 0x12, 0xc2, 0xa4, 0x96, 0xbc, 0x0a, 0xca, 0x2a, 0x15, 0x55, 0xaa, 0x1a,
	0xf3, 0x27, 0x38, 0x8b, 0x13, 0x2d, 0xfc, 0x7d, 0x79, 0xf8, 0x30, 0x49, 0xd5, 0x79, 0x1d, 0xa2,
	0x5f, 0xee, 0x77, 0x9e, 0xab, 0xd0, 0xff, 0xef, 0xc1, 0x78, 0xaf, 0x70, 0xca, 0x4b, 0x3b, 0xe4,
	0xf4, 0x5a, 0xdd, 0x3b, 0x91, 0x47, 0xb0, 0x9f, 0xd7, 0x8a, 0x29, 0x2e, 0x03, 0xd9, 0x48, 0xc5,
	0xf3, 0x20, 0x12, 0x45, 0x9c, 0x26, 0x74, 0xd8, 0x0b, 0x75, 0xc3, 0x4a, 0xce, 0x8c, 0xe2, 0xa9,
	0x11, 0x2c, 0x00, 0xc6, 0x0a, 0x4b, 0xfc, 0x28, 0x0a, 0xbe, 0xb8, 0xbb, 0xfe, 0x3e, 0x1b, 0xac,
	0x7f, 0xcc, 0x9c, 0x2f, 0xb8, 0xbe, 0xe2, 0xfa, 0x86, 0xeb, 0xf3, 0xcf, 0xd9, 0xe0, 0x8d, 0x8b,
	0xcf, 0xe4, 0xb5, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x41, 0x31, 0xdf, 0xbb, 0xf2, 0x02, 0x00,
	0x00,
}
