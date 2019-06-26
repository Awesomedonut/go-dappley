// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/metrics/pb/datastore.proto

package metricspb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pb "github.com/dappley/go-dappley/network/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Stat struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Stat_TransactionPoolSize
	//	*Stat_MemoryStats
	//	*Stat_CpuPercentage
	//	*Stat_ForkStats
	//	*Stat_BlockStats
	Value                isStat_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{0}
}
func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}
func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}
func (dst *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(dst, src)
}
func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}
func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type isStat_Value interface {
	isStat_Value()
}

type Stat_TransactionPoolSize struct {
	TransactionPoolSize int64 `protobuf:"varint,2,opt,name=transaction_pool_size,json=transactionPoolSize,proto3,oneof"`
}

type Stat_MemoryStats struct {
	MemoryStats *MemoryStats `protobuf:"bytes,3,opt,name=memory_stats,json=memoryStats,proto3,oneof"`
}

type Stat_CpuPercentage struct {
	CpuPercentage float64 `protobuf:"fixed64,4,opt,name=cpu_percentage,json=cpuPercentage,proto3,oneof"`
}

type Stat_ForkStats struct {
	ForkStats *ForkStats `protobuf:"bytes,5,opt,name=fork_stats,json=forkStats,proto3,oneof"`
}

type Stat_BlockStats struct {
	BlockStats *BlockStats `protobuf:"bytes,6,opt,name=block_stats,json=blockStats,proto3,oneof"`
}

func (*Stat_TransactionPoolSize) isStat_Value() {}

func (*Stat_MemoryStats) isStat_Value() {}

func (*Stat_CpuPercentage) isStat_Value() {}

func (*Stat_ForkStats) isStat_Value() {}

func (*Stat_BlockStats) isStat_Value() {}

func (m *Stat) GetValue() isStat_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Stat) GetTransactionPoolSize() int64 {
	if x, ok := m.GetValue().(*Stat_TransactionPoolSize); ok {
		return x.TransactionPoolSize
	}
	return 0
}

func (m *Stat) GetMemoryStats() *MemoryStats {
	if x, ok := m.GetValue().(*Stat_MemoryStats); ok {
		return x.MemoryStats
	}
	return nil
}

func (m *Stat) GetCpuPercentage() float64 {
	if x, ok := m.GetValue().(*Stat_CpuPercentage); ok {
		return x.CpuPercentage
	}
	return 0
}

func (m *Stat) GetForkStats() *ForkStats {
	if x, ok := m.GetValue().(*Stat_ForkStats); ok {
		return x.ForkStats
	}
	return nil
}

func (m *Stat) GetBlockStats() *BlockStats {
	if x, ok := m.GetValue().(*Stat_BlockStats); ok {
		return x.BlockStats
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Stat) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Stat_OneofMarshaler, _Stat_OneofUnmarshaler, _Stat_OneofSizer, []interface{}{
		(*Stat_TransactionPoolSize)(nil),
		(*Stat_MemoryStats)(nil),
		(*Stat_CpuPercentage)(nil),
		(*Stat_ForkStats)(nil),
		(*Stat_BlockStats)(nil),
	}
}

func _Stat_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Stat)
	// value
	switch x := m.Value.(type) {
	case *Stat_TransactionPoolSize:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.TransactionPoolSize))
	case *Stat_MemoryStats:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MemoryStats); err != nil {
			return err
		}
	case *Stat_CpuPercentage:
		b.EncodeVarint(4<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.CpuPercentage))
	case *Stat_ForkStats:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ForkStats); err != nil {
			return err
		}
	case *Stat_BlockStats:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BlockStats); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Stat.Value has unexpected type %T", x)
	}
	return nil
}

func _Stat_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Stat)
	switch tag {
	case 2: // value.transaction_pool_size
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &Stat_TransactionPoolSize{int64(x)}
		return true, err
	case 3: // value.memory_stats
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MemoryStats)
		err := b.DecodeMessage(msg)
		m.Value = &Stat_MemoryStats{msg}
		return true, err
	case 4: // value.cpu_percentage
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Value = &Stat_CpuPercentage{math.Float64frombits(x)}
		return true, err
	case 5: // value.fork_stats
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ForkStats)
		err := b.DecodeMessage(msg)
		m.Value = &Stat_ForkStats{msg}
		return true, err
	case 6: // value.block_stats
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BlockStats)
		err := b.DecodeMessage(msg)
		m.Value = &Stat_BlockStats{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Stat_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Stat)
	// value
	switch x := m.Value.(type) {
	case *Stat_TransactionPoolSize:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.TransactionPoolSize))
	case *Stat_MemoryStats:
		s := proto.Size(x.MemoryStats)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Stat_CpuPercentage:
		n += 1 // tag and wire
		n += 8
	case *Stat_ForkStats:
		s := proto.Size(x.ForkStats)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Stat_BlockStats:
		s := proto.Size(x.BlockStats)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type MemoryStats struct {
	HeapInUse            uint64   `protobuf:"varint,1,opt,name=heap_in_use,json=heapInUse,proto3" json:"heap_in_use,omitempty"`
	HeapSys              uint64   `protobuf:"varint,2,opt,name=heap_sys,json=heapSys,proto3" json:"heap_sys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemoryStats) Reset()         { *m = MemoryStats{} }
func (m *MemoryStats) String() string { return proto.CompactTextString(m) }
func (*MemoryStats) ProtoMessage()    {}
func (*MemoryStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{1}
}
func (m *MemoryStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryStats.Unmarshal(m, b)
}
func (m *MemoryStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryStats.Marshal(b, m, deterministic)
}
func (dst *MemoryStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryStats.Merge(dst, src)
}
func (m *MemoryStats) XXX_Size() int {
	return xxx_messageInfo_MemoryStats.Size(m)
}
func (m *MemoryStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryStats.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryStats proto.InternalMessageInfo

func (m *MemoryStats) GetHeapInUse() uint64 {
	if m != nil {
		return m.HeapInUse
	}
	return 0
}

func (m *MemoryStats) GetHeapSys() uint64 {
	if m != nil {
		return m.HeapSys
	}
	return 0
}

type ForkStats struct {
	NumForks             int64    `protobuf:"varint,1,opt,name=num_forks,json=numForks,proto3" json:"num_forks,omitempty"`
	LongestFork          int64    `protobuf:"varint,2,opt,name=longest_fork,json=longestFork,proto3" json:"longest_fork,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForkStats) Reset()         { *m = ForkStats{} }
func (m *ForkStats) String() string { return proto.CompactTextString(m) }
func (*ForkStats) ProtoMessage()    {}
func (*ForkStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{2}
}
func (m *ForkStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForkStats.Unmarshal(m, b)
}
func (m *ForkStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForkStats.Marshal(b, m, deterministic)
}
func (dst *ForkStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForkStats.Merge(dst, src)
}
func (m *ForkStats) XXX_Size() int {
	return xxx_messageInfo_ForkStats.Size(m)
}
func (m *ForkStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ForkStats.DiscardUnknown(m)
}

var xxx_messageInfo_ForkStats proto.InternalMessageInfo

func (m *ForkStats) GetNumForks() int64 {
	if m != nil {
		return m.NumForks
	}
	return 0
}

func (m *ForkStats) GetLongestFork() int64 {
	if m != nil {
		return m.LongestFork
	}
	return 0
}

type BlockStats struct {
	NumTransactions      uint64   `protobuf:"varint,1,opt,name=num_transactions,json=numTransactions,proto3" json:"num_transactions,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockStats) Reset()         { *m = BlockStats{} }
func (m *BlockStats) String() string { return proto.CompactTextString(m) }
func (*BlockStats) ProtoMessage()    {}
func (*BlockStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{3}
}
func (m *BlockStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockStats.Unmarshal(m, b)
}
func (m *BlockStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockStats.Marshal(b, m, deterministic)
}
func (dst *BlockStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStats.Merge(dst, src)
}
func (m *BlockStats) XXX_Size() int {
	return xxx_messageInfo_BlockStats.Size(m)
}
func (m *BlockStats) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStats.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStats proto.InternalMessageInfo

func (m *BlockStats) GetNumTransactions() uint64 {
	if m != nil {
		return m.NumTransactions
	}
	return 0
}

func (m *BlockStats) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Metric struct {
	Stats                []*Stat  `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{4}
}
func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (dst *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(dst, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetStats() []*Stat {
	if m != nil {
		return m.Stats
	}
	return nil
}

type DataStore struct {
	Metrics              map[string]*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataStore) Reset()         { *m = DataStore{} }
func (m *DataStore) String() string { return proto.CompactTextString(m) }
func (*DataStore) ProtoMessage()    {}
func (*DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{5}
}
func (m *DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStore.Unmarshal(m, b)
}
func (m *DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStore.Marshal(b, m, deterministic)
}
func (dst *DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStore.Merge(dst, src)
}
func (m *DataStore) XXX_Size() int {
	return xxx_messageInfo_DataStore.Size(m)
}
func (m *DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_DataStore proto.InternalMessageInfo

func (m *DataStore) GetMetrics() map[string]*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Metrics struct {
	DataStore            *DataStore     `protobuf:"bytes,1,opt,name=data_store,json=dataStore,proto3" json:"data_store,omitempty"`
	Peers                []*pb.PeerInfo `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	BlockStats           []*BlockStats  `protobuf:"bytes,3,rep,name=block_stats,json=blockStats,proto3" json:"block_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Metrics) Reset()         { *m = Metrics{} }
func (m *Metrics) String() string { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()    {}
func (*Metrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_datastore_6b07dd97cd0e6b91, []int{6}
}
func (m *Metrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metrics.Unmarshal(m, b)
}
func (m *Metrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metrics.Marshal(b, m, deterministic)
}
func (dst *Metrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metrics.Merge(dst, src)
}
func (m *Metrics) XXX_Size() int {
	return xxx_messageInfo_Metrics.Size(m)
}
func (m *Metrics) XXX_DiscardUnknown() {
	xxx_messageInfo_Metrics.DiscardUnknown(m)
}

var xxx_messageInfo_Metrics proto.InternalMessageInfo

func (m *Metrics) GetDataStore() *DataStore {
	if m != nil {
		return m.DataStore
	}
	return nil
}

func (m *Metrics) GetPeers() []*pb.PeerInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *Metrics) GetBlockStats() []*BlockStats {
	if m != nil {
		return m.BlockStats
	}
	return nil
}

func init() {
	proto.RegisterType((*Stat)(nil), "metricspb.Stat")
	proto.RegisterType((*MemoryStats)(nil), "metricspb.MemoryStats")
	proto.RegisterType((*ForkStats)(nil), "metricspb.ForkStats")
	proto.RegisterType((*BlockStats)(nil), "metricspb.BlockStats")
	proto.RegisterType((*Metric)(nil), "metricspb.Metric")
	proto.RegisterType((*DataStore)(nil), "metricspb.DataStore")
	proto.RegisterMapType((map[string]*Metric)(nil), "metricspb.DataStore.MetricsEntry")
	proto.RegisterType((*Metrics)(nil), "metricspb.Metrics")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/metrics/pb/datastore.proto", fileDescriptor_datastore_6b07dd97cd0e6b91)
}

var fileDescriptor_datastore_6b07dd97cd0e6b91 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0xd4, 0x3c,
	0x10, 0x6d, 0xba, 0x7f, 0xcd, 0xa4, 0xdf, 0xd7, 0xe2, 0xd2, 0x6a, 0x29, 0x08, 0x6d, 0x23, 0xa1,
	0x6e, 0x2f, 0xc8, 0x4a, 0x2d, 0xa0, 0x8a, 0xde, 0x55, 0x80, 0xb6, 0x42, 0x15, 0x95, 0x17, 0xae,
	0x23, 0x27, 0x75, 0x77, 0xa3, 0x4d, 0x6c, 0xcb, 0x76, 0x40, 0xe9, 0x8b, 0x70, 0xcd, 0x6b, 0xf0,
	0x74, 0xc8, 0x8e, 0x93, 0xa6, 0x12, 0xe2, 0xce, 0x33, 0x67, 0xce, 0xcc, 0xf1, 0xfc, 0xc0, 0xf9,
	0x32, 0xd3, 0xab, 0x32, 0x89, 0x52, 0x5e, 0xcc, 0x6e, 0x89, 0x10, 0x39, 0xad, 0x66, 0x4b, 0xfe,
	0xba, 0x79, 0x16, 0x54, 0xcb, 0x2c, 0x55, 0x33, 0x91, 0xcc, 0x6e, 0x89, 0x26, 0x4a, 0x73, 0x49,
	0x23, 0x21, 0xb9, 0xe6, 0xc8, 0x77, 0x98, 0x48, 0x0e, 0xcf, 0xfe, 0x9d, 0x84, 0x51, 0xfd, 0x83,
	0xcb, 0xb5, 0x49, 0x22, 0x28, 0x95, 0x35, 0x3f, 0xfc, 0xbd, 0x09, 0xfd, 0x85, 0x26, 0x1a, 0xbd,
	0x00, 0x5f, 0x67, 0x05, 0x55, 0x9a, 0x14, 0x62, 0xec, 0x4d, 0xbc, 0x69, 0x0f, 0x3f, 0x38, 0xd0,
	0x1b, 0xd8, 0xd7, 0x92, 0x30, 0x45, 0x52, 0x9d, 0x71, 0x16, 0x0b, 0xce, 0xf3, 0x58, 0x65, 0xf7,
	0x74, 0xbc, 0x69, 0x22, 0xe7, 0x1b, 0x78, 0xaf, 0x03, 0xdf, 0x70, 0x9e, 0x2f, 0xb2, 0x7b, 0x8a,
	0x2e, 0x60, 0xbb, 0xa0, 0x05, 0x97, 0x55, 0xac, 0x34, 0xd1, 0x6a, 0xdc, 0x9b, 0x78, 0xd3, 0xe0,
	0xf4, 0x20, 0x6a, 0x35, 0x47, 0xd7, 0x16, 0x36, 0x02, 0xd4, 0x7c, 0x03, 0x07, 0xc5, 0x83, 0x89,
	0x8e, 0xe1, 0xff, 0x54, 0x94, 0xb1, 0xa0, 0x32, 0xa5, 0x4c, 0x93, 0x25, 0x1d, 0xf7, 0x27, 0xde,
	0xd4, 0x9b, 0x6f, 0xe0, 0xff, 0x52, 0x51, 0xde, 0xb4, 0x6e, 0xf4, 0x16, 0xe0, 0x8e, 0xcb, 0xb5,
	0xab, 0x31, 0xb0, 0x35, 0x9e, 0x76, 0x6a, 0x7c, 0xe2, 0x72, 0xdd, 0x54, 0xf0, 0xef, 0x1a, 0x03,
	0x9d, 0x43, 0x90, 0xe4, 0x3c, 0x6d, 0x78, 0x43, 0xcb, 0xdb, 0xef, 0xf0, 0x2e, 0x0d, 0xda, 0x10,
	0x21, 0x69, 0xad, 0xcb, 0x11, 0x0c, 0xbe, 0x93, 0xbc, 0xa4, 0xe1, 0x1c, 0x82, 0xce, 0x07, 0xd0,
	0x4b, 0x08, 0x56, 0x94, 0x88, 0x38, 0x63, 0x71, 0xa9, 0xa8, 0x6d, 0x62, 0x1f, 0xfb, 0xc6, 0x75,
	0xc5, 0xbe, 0x29, 0x8a, 0x9e, 0xc1, 0x96, 0xc5, 0x55, 0xa5, 0x6c, 0xdf, 0xfa, 0x78, 0x64, 0xec,
	0x45, 0xa5, 0xc2, 0xcf, 0xe0, 0xb7, 0x32, 0xd1, 0x73, 0xf0, 0x59, 0x59, 0xc4, 0x46, 0xaa, 0x72,
	0xa3, 0xd8, 0x62, 0x65, 0x61, 0x02, 0x14, 0x3a, 0x82, 0xed, 0x9c, 0xb3, 0x25, 0x55, 0xda, 0x06,
	0xd4, 0x03, 0xc0, 0x81, 0xf3, 0x99, 0x98, 0xf0, 0x0b, 0xc0, 0x83, 0x76, 0x74, 0x02, 0xbb, 0x26,
	0x5b, 0x67, 0x3e, 0xca, 0x49, 0xdb, 0x61, 0x65, 0xf1, 0xb5, 0xe3, 0x46, 0x07, 0x30, 0x5c, 0xd1,
	0x6c, 0xb9, 0xd2, 0x4e, 0x9e, 0xb3, 0xc2, 0x19, 0x0c, 0xaf, 0x6d, 0x5b, 0xd0, 0x2b, 0x18, 0xd4,
	0xed, 0xf2, 0x26, 0xbd, 0x69, 0x70, 0xba, 0xd3, 0x69, 0x97, 0xa9, 0x86, 0x6b, 0x34, 0xfc, 0xe9,
	0x81, 0xff, 0x81, 0x68, 0xb2, 0x30, 0x9b, 0x8a, 0x2e, 0x60, 0xe4, 0xc2, 0x1c, 0xed, 0xa8, 0x43,
	0x6b, 0xc3, 0xa2, 0xba, 0x84, 0xfa, 0xc8, 0xb4, 0xac, 0x70, 0xc3, 0x38, 0xbc, 0x86, 0xed, 0x2e,
	0x80, 0x76, 0xa1, 0xb7, 0xa6, 0x95, 0xfd, 0x81, 0x8f, 0xcd, 0x13, 0x1d, 0xbb, 0x71, 0x58, 0xd1,
	0xc1, 0xe9, 0x93, 0x47, 0xeb, 0x65, 0x5e, 0xb8, 0xc6, 0xdf, 0x6f, 0x9e, 0x7b, 0xe1, 0x2f, 0x0f,
	0x46, 0x2e, 0x1f, 0x3a, 0x03, 0x30, 0xe7, 0x14, 0xdb, 0x7b, 0xb2, 0x19, 0x1f, 0x2f, 0x4e, 0x2b,
	0x0d, 0xfb, 0xb7, 0xed, 0x67, 0x4e, 0x60, 0x60, 0xce, 0xc7, 0x4c, 0xd0, 0x7c, 0x65, 0x2f, 0x72,
	0x77, 0x25, 0x92, 0xe8, 0x86, 0x52, 0x79, 0xc5, 0xee, 0x38, 0xae, 0x23, 0xd0, 0xbb, 0xc7, 0x1b,
	0xd6, 0xb3, 0x84, 0xbf, 0x6f, 0x58, 0x77, 0xbf, 0x92, 0xa1, 0x3d, 0xcd, 0xb3, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf8, 0xda, 0x32, 0xe1, 0x16, 0x04, 0x00, 0x00,
}
