// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dappley/go-dappley/config/pb/config.proto

package configpb

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

type Config struct {
	ConsensusConfig      *ConsensusConfig `protobuf:"bytes,1,opt,name=consensus_config,json=consensusConfig,proto3" json:"consensus_config,omitempty"`
	NodeConfig           *NodeConfig      `protobuf:"bytes,2,opt,name=node_config,json=nodeConfig,proto3" json:"node_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_0475c0df06877262, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetConsensusConfig() *ConsensusConfig {
	if m != nil {
		return m.ConsensusConfig
	}
	return nil
}

func (m *Config) GetNodeConfig() *NodeConfig {
	if m != nil {
		return m.NodeConfig
	}
	return nil
}

type ConsensusConfig struct {
	MinerAddress         string   `protobuf:"bytes,1,opt,name=miner_address,json=minerAddress,proto3" json:"miner_address,omitempty"`
	PrivateKey           string   `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsensusConfig) Reset()         { *m = ConsensusConfig{} }
func (m *ConsensusConfig) String() string { return proto.CompactTextString(m) }
func (*ConsensusConfig) ProtoMessage()    {}
func (*ConsensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_0475c0df06877262, []int{1}
}
func (m *ConsensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusConfig.Unmarshal(m, b)
}
func (m *ConsensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusConfig.Marshal(b, m, deterministic)
}
func (dst *ConsensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusConfig.Merge(dst, src)
}
func (m *ConsensusConfig) XXX_Size() int {
	return xxx_messageInfo_ConsensusConfig.Size(m)
}
func (m *ConsensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusConfig proto.InternalMessageInfo

func (m *ConsensusConfig) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *ConsensusConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

type NodeConfig struct {
	Port                   uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Seed                   []string `protobuf:"bytes,2,rep,name=seed,proto3" json:"seed,omitempty"`
	DbPath                 string   `protobuf:"bytes,3,opt,name=db_path,json=dbPath,proto3" json:"db_path,omitempty"`
	RpcPort                uint32   `protobuf:"varint,4,opt,name=rpc_port,json=rpcPort,proto3" json:"rpc_port,omitempty"`
	KeyPath                string   `protobuf:"bytes,5,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty"`
	TxPoolLimit            uint32   `protobuf:"varint,6,opt,name=tx_pool_limit,json=txPoolLimit,proto3" json:"tx_pool_limit,omitempty"`
	BlkSizeLimit           uint32   `protobuf:"varint,7,opt,name=blk_size_limit,json=blkSizeLimit,proto3" json:"blk_size_limit,omitempty"`
	NodeAddress            string   `protobuf:"bytes,8,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty"`
	GenesisPath            string   `protobuf:"bytes,9,opt,name=genesis_path,json=genesisPath,proto3" json:"genesis_path,omitempty"`
	MetricsPollingInterval int64    `protobuf:"varint,12,opt,name=metrics_polling_interval,json=metricsPollingInterval,proto3" json:"metrics_polling_interval,omitempty"`
	MetricsInterval        int64    `protobuf:"varint,13,opt,name=metrics_interval,json=metricsInterval,proto3" json:"metrics_interval,omitempty"`
	MaxConnectionOut       uint32   `protobuf:"varint,14,opt,name=max_connection_out,json=maxConnectionOut,proto3" json:"max_connection_out,omitempty"`
	MaxConnectionIn        uint32   `protobuf:"varint,15,opt,name=max_connection_in,json=maxConnectionIn,proto3" json:"max_connection_in,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *NodeConfig) Reset()         { *m = NodeConfig{} }
func (m *NodeConfig) String() string { return proto.CompactTextString(m) }
func (*NodeConfig) ProtoMessage()    {}
func (*NodeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_0475c0df06877262, []int{2}
}
func (m *NodeConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeConfig.Unmarshal(m, b)
}
func (m *NodeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeConfig.Marshal(b, m, deterministic)
}
func (dst *NodeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeConfig.Merge(dst, src)
}
func (m *NodeConfig) XXX_Size() int {
	return xxx_messageInfo_NodeConfig.Size(m)
}
func (m *NodeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NodeConfig proto.InternalMessageInfo

func (m *NodeConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *NodeConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NodeConfig) GetDbPath() string {
	if m != nil {
		return m.DbPath
	}
	return ""
}

func (m *NodeConfig) GetRpcPort() uint32 {
	if m != nil {
		return m.RpcPort
	}
	return 0
}

func (m *NodeConfig) GetKeyPath() string {
	if m != nil {
		return m.KeyPath
	}
	return ""
}

func (m *NodeConfig) GetTxPoolLimit() uint32 {
	if m != nil {
		return m.TxPoolLimit
	}
	return 0
}

func (m *NodeConfig) GetBlkSizeLimit() uint32 {
	if m != nil {
		return m.BlkSizeLimit
	}
	return 0
}

func (m *NodeConfig) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *NodeConfig) GetGenesisPath() string {
	if m != nil {
		return m.GenesisPath
	}
	return ""
}

func (m *NodeConfig) GetMetricsPollingInterval() int64 {
	if m != nil {
		return m.MetricsPollingInterval
	}
	return 0
}

func (m *NodeConfig) GetMetricsInterval() int64 {
	if m != nil {
		return m.MetricsInterval
	}
	return 0
}

func (m *NodeConfig) GetMaxConnectionOut() uint32 {
	if m != nil {
		return m.MaxConnectionOut
	}
	return 0
}

func (m *NodeConfig) GetMaxConnectionIn() uint32 {
	if m != nil {
		return m.MaxConnectionIn
	}
	return 0
}

type DynastyConfig struct {
	Producers            []string `protobuf:"bytes,1,rep,name=producers,proto3" json:"producers,omitempty"`
	MaxProducers         uint32   `protobuf:"varint,2,opt,name=max_producers,json=maxProducers,proto3" json:"max_producers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DynastyConfig) Reset()         { *m = DynastyConfig{} }
func (m *DynastyConfig) String() string { return proto.CompactTextString(m) }
func (*DynastyConfig) ProtoMessage()    {}
func (*DynastyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_0475c0df06877262, []int{3}
}
func (m *DynastyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynastyConfig.Unmarshal(m, b)
}
func (m *DynastyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynastyConfig.Marshal(b, m, deterministic)
}
func (dst *DynastyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynastyConfig.Merge(dst, src)
}
func (m *DynastyConfig) XXX_Size() int {
	return xxx_messageInfo_DynastyConfig.Size(m)
}
func (m *DynastyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DynastyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DynastyConfig proto.InternalMessageInfo

func (m *DynastyConfig) GetProducers() []string {
	if m != nil {
		return m.Producers
	}
	return nil
}

func (m *DynastyConfig) GetMaxProducers() uint32 {
	if m != nil {
		return m.MaxProducers
	}
	return 0
}

type CliConfig struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CliConfig) Reset()         { *m = CliConfig{} }
func (m *CliConfig) String() string { return proto.CompactTextString(m) }
func (*CliConfig) ProtoMessage()    {}
func (*CliConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_0475c0df06877262, []int{4}
}
func (m *CliConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CliConfig.Unmarshal(m, b)
}
func (m *CliConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CliConfig.Marshal(b, m, deterministic)
}
func (dst *CliConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CliConfig.Merge(dst, src)
}
func (m *CliConfig) XXX_Size() int {
	return xxx_messageInfo_CliConfig.Size(m)
}
func (m *CliConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CliConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CliConfig proto.InternalMessageInfo

func (m *CliConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *CliConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "configpb.Config")
	proto.RegisterType((*ConsensusConfig)(nil), "configpb.ConsensusConfig")
	proto.RegisterType((*NodeConfig)(nil), "configpb.NodeConfig")
	proto.RegisterType((*DynastyConfig)(nil), "configpb.DynastyConfig")
	proto.RegisterType((*CliConfig)(nil), "configpb.CliConfig")
}

func init() {
	proto.RegisterFile("github.com/dappley/go-dappley/config/pb/config.proto", fileDescriptor_config_0475c0df06877262)
}

var fileDescriptor_config_0475c0df06877262 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xdb, 0xa5, 0x6d, 0xa6, 0xed, 0xb6, 0x58, 0x08, 0x52, 0x84, 0x44, 0x37, 0x70, 0x28,
	0x08, 0x5a, 0x89, 0x0f, 0x09, 0x89, 0x13, 0xea, 0x5e, 0x56, 0x20, 0xa8, 0xc2, 0x81, 0xa3, 0xe5,
	0x24, 0xa6, 0xb5, 0x9a, 0xd8, 0x96, 0xed, 0x2e, 0xc9, 0xde, 0xf9, 0x83, 0xfc, 0x22, 0x94, 0xc9,
	0x47, 0xb5, 0x7b, 0xd8, 0xdb, 0xcc, 0x9b, 0xf7, 0xde, 0x4c, 0xc6, 0x13, 0xf8, 0xb0, 0x13, 0x6e,
	0x7f, 0x8c, 0x56, 0xb1, 0xca, 0xd6, 0x09, 0xd3, 0x3a, 0xe5, 0xc5, 0x7a, 0xa7, 0xde, 0x36, 0x61,
	0xac, 0xe4, 0x6f, 0xb1, 0x5b, 0xeb, 0xa8, 0x8e, 0x56, 0xda, 0x28, 0xa7, 0xc8, 0xb0, 0xca, 0x74,
	0x14, 0xfc, 0xed, 0x40, 0x7f, 0x83, 0x09, 0xb9, 0x84, 0x59, 0xac, 0xa4, 0xe5, 0xd2, 0x1e, 0x2d,
	0xad, 0x08, 0x7e, 0x67, 0xd1, 0x59, 0x8e, 0xde, 0xcd, 0x57, 0x0d, 0x7f, 0xb5, 0x69, 0x18, 0x95,
	0x28, 0x9c, 0xc6, 0xb7, 0x01, 0xf2, 0x11, 0x46, 0x52, 0x25, 0xbc, 0x31, 0xe8, 0xa2, 0xc1, 0xa3,
	0x93, 0xc1, 0x77, 0x95, 0xf0, 0x5a, 0x0b, 0xb2, 0x8d, 0x83, 0x5f, 0x30, 0xbd, 0x63, 0x4d, 0x5e,
	0xc0, 0x24, 0x13, 0x92, 0x1b, 0xca, 0x92, 0xc4, 0x70, 0x6b, 0x71, 0x18, 0x2f, 0x1c, 0x23, 0xf8,
	0xa5, 0xc2, 0xc8, 0x73, 0x18, 0x69, 0x23, 0xae, 0x99, 0xe3, 0xf4, 0xc0, 0x0b, 0x6c, 0xe7, 0x85,
	0x50, 0x43, 0x5f, 0x79, 0x11, 0xfc, 0xeb, 0x01, 0x9c, 0x7a, 0x12, 0x02, 0x67, 0x5a, 0x19, 0x87,
	0x5e, 0x93, 0x10, 0xe3, 0x12, 0xb3, 0x9c, 0x27, 0x7e, 0x77, 0xd1, 0x5b, 0x7a, 0x21, 0xc6, 0xe4,
	0x09, 0x0c, 0x92, 0x88, 0x6a, 0xe6, 0xf6, 0x7e, 0x0f, 0x3d, 0xfb, 0x49, 0xb4, 0x65, 0x6e, 0x4f,
	0xe6, 0x30, 0x34, 0x3a, 0xa6, 0x68, 0x72, 0x86, 0x26, 0x03, 0xa3, 0xe3, 0x6d, 0xe9, 0x33, 0x87,
	0xe1, 0x81, 0x17, 0x95, 0xe8, 0x01, 0x8a, 0x06, 0x07, 0x5e, 0xa0, 0x2a, 0x80, 0x89, 0xcb, 0xa9,
	0x56, 0x2a, 0xa5, 0xa9, 0xc8, 0x84, 0xf3, 0xfb, 0x28, 0x1d, 0xb9, 0x7c, 0xab, 0x54, 0xfa, 0xad,
	0x84, 0xc8, 0x4b, 0x38, 0x8f, 0xd2, 0x03, 0xb5, 0xe2, 0x86, 0xd7, 0xa4, 0x01, 0x92, 0xc6, 0x51,
	0x7a, 0xf8, 0x29, 0x6e, 0x78, 0xc5, 0xba, 0x80, 0x31, 0xee, 0xb7, 0x59, 0xca, 0x10, 0x1b, 0xe1,
	0xce, 0x9b, 0x9d, 0x5c, 0xc0, 0x78, 0xc7, 0x25, 0xb7, 0xc2, 0x56, 0xb3, 0x78, 0x15, 0xa5, 0xc6,
	0x70, 0x9e, 0x4f, 0xe0, 0x67, 0xdc, 0x19, 0x11, 0x5b, 0xaa, 0x55, 0x9a, 0x0a, 0xb9, 0xa3, 0x42,
	0x3a, 0x6e, 0xae, 0x59, 0xea, 0x8f, 0x17, 0x9d, 0x65, 0x2f, 0x7c, 0x5c, 0xd7, 0xb7, 0x55, 0xf9,
	0xaa, 0xae, 0x92, 0x57, 0x30, 0x6b, 0x94, 0xad, 0x62, 0x82, 0x8a, 0x69, 0x8d, 0xb7, 0xd4, 0x37,
	0x40, 0x32, 0x96, 0x97, 0x97, 0x20, 0x79, 0xec, 0x84, 0x92, 0x54, 0x1d, 0x9d, 0x7f, 0x8e, 0x1f,
	0x35, 0xcb, 0x58, 0xbe, 0x69, 0x0b, 0x3f, 0x8e, 0x8e, 0xbc, 0x86, 0x87, 0x77, 0xd8, 0x42, 0xfa,
	0x53, 0x24, 0x4f, 0x6f, 0x91, 0xaf, 0x64, 0x10, 0xc2, 0xe4, 0xb2, 0x90, 0xcc, 0xba, 0xa2, 0x7e,
	0xd6, 0x67, 0xe0, 0x69, 0xa3, 0x92, 0x63, 0xcc, 0x4d, 0x79, 0x27, 0xe5, 0x3b, 0x9e, 0x00, 0xbc,
	0x24, 0x96, 0xd3, 0x13, 0xa3, 0x5b, 0x2d, 0x36, 0x63, 0xf9, 0xb6, 0xc1, 0x82, 0xcf, 0xe0, 0x6d,
	0x52, 0x71, 0xcf, 0x99, 0x3c, 0x85, 0xa1, 0x66, 0xd6, 0xfe, 0x51, 0x26, 0xa9, 0xef, 0xac, 0xcd,
	0xa3, 0x3e, 0xfe, 0x57, 0xef, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x2b, 0xf9, 0xfe, 0x8f,
	0x03, 0x00, 0x00,
}
