// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/pytorch.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Custom proto for torch elastic config for distributed training using
// https://github.com/kubeflow/training-operator/blob/master/pkg/apis/kubeflow.org/v1/pytorch_types.go
type ElasticConfig struct {
	RdzvBackend          string   `protobuf:"bytes,1,opt,name=rdzv_backend,json=rdzvBackend,proto3" json:"rdzv_backend,omitempty"`
	MinReplicas          int32    `protobuf:"varint,2,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	MaxReplicas          int32    `protobuf:"varint,3,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	NprocPerNode         int32    `protobuf:"varint,4,opt,name=nproc_per_node,json=nprocPerNode,proto3" json:"nproc_per_node,omitempty"`
	MaxRestarts          int32    `protobuf:"varint,5,opt,name=max_restarts,json=maxRestarts,proto3" json:"max_restarts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElasticConfig) Reset()         { *m = ElasticConfig{} }
func (m *ElasticConfig) String() string { return proto.CompactTextString(m) }
func (*ElasticConfig) ProtoMessage()    {}
func (*ElasticConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4df8a9374b28b766, []int{0}
}

func (m *ElasticConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElasticConfig.Unmarshal(m, b)
}
func (m *ElasticConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElasticConfig.Marshal(b, m, deterministic)
}
func (m *ElasticConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElasticConfig.Merge(m, src)
}
func (m *ElasticConfig) XXX_Size() int {
	return xxx_messageInfo_ElasticConfig.Size(m)
}
func (m *ElasticConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ElasticConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ElasticConfig proto.InternalMessageInfo

func (m *ElasticConfig) GetRdzvBackend() string {
	if m != nil {
		return m.RdzvBackend
	}
	return ""
}

func (m *ElasticConfig) GetMinReplicas() int32 {
	if m != nil {
		return m.MinReplicas
	}
	return 0
}

func (m *ElasticConfig) GetMaxReplicas() int32 {
	if m != nil {
		return m.MaxReplicas
	}
	return 0
}

func (m *ElasticConfig) GetNprocPerNode() int32 {
	if m != nil {
		return m.NprocPerNode
	}
	return 0
}

func (m *ElasticConfig) GetMaxRestarts() int32 {
	if m != nil {
		return m.MaxRestarts
	}
	return 0
}

// Custom proto for plugin that enables distributed training using https://github.com/kubeflow/pytorch-operator
type DistributedPyTorchTrainingTask struct {
	// number of worker replicas spawned in the cluster for this job
	Workers int32 `protobuf:"varint,1,opt,name=workers,proto3" json:"workers,omitempty"`
	// config for an elastic pytorch job
	//
	ElasticConfig        *ElasticConfig `protobuf:"bytes,2,opt,name=elastic_config,json=elasticConfig,proto3" json:"elastic_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DistributedPyTorchTrainingTask) Reset()         { *m = DistributedPyTorchTrainingTask{} }
func (m *DistributedPyTorchTrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedPyTorchTrainingTask) ProtoMessage()    {}
func (*DistributedPyTorchTrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_4df8a9374b28b766, []int{1}
}

func (m *DistributedPyTorchTrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedPyTorchTrainingTask.Unmarshal(m, b)
}
func (m *DistributedPyTorchTrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedPyTorchTrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedPyTorchTrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedPyTorchTrainingTask.Merge(m, src)
}
func (m *DistributedPyTorchTrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedPyTorchTrainingTask.Size(m)
}
func (m *DistributedPyTorchTrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedPyTorchTrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedPyTorchTrainingTask proto.InternalMessageInfo

func (m *DistributedPyTorchTrainingTask) GetWorkers() int32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func (m *DistributedPyTorchTrainingTask) GetElasticConfig() *ElasticConfig {
	if m != nil {
		return m.ElasticConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ElasticConfig)(nil), "flyteidl.plugins.ElasticConfig")
	proto.RegisterType((*DistributedPyTorchTrainingTask)(nil), "flyteidl.plugins.DistributedPyTorchTrainingTask")
}

func init() { proto.RegisterFile("flyteidl/plugins/pytorch.proto", fileDescriptor_4df8a9374b28b766) }

var fileDescriptor_4df8a9374b28b766 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xbd, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0xa0, 0x20, 0xdc, 0x0f, 0xa1, 0x4c, 0x99, 0x4a, 0xa9, 0x18, 0xba, 0x90, 0x48,
	0x30, 0x20, 0xd6, 0xf2, 0x31, 0xa2, 0x2a, 0xea, 0xc4, 0x12, 0x39, 0xf6, 0xd5, 0x3d, 0x35, 0xb5,
	0xad, 0xb3, 0x0b, 0x2d, 0x23, 0xff, 0x19, 0xff, 0x19, 0xaa, 0x9b, 0x7e, 0xd0, 0xf1, 0xde, 0xfd,
	0xee, 0xa4, 0xf7, 0x1e, 0xeb, 0x4e, 0xaa, 0x95, 0x07, 0x94, 0x55, 0x66, 0xab, 0x85, 0x42, 0xed,
	0x32, 0xbb, 0xf2, 0x86, 0xc4, 0x34, 0xb5, 0x64, 0xbc, 0x89, 0xaf, 0xb6, 0xfb, 0xb4, 0xde, 0xf7,
	0x7f, 0x23, 0xd6, 0x7e, 0xad, 0xb8, 0xf3, 0x28, 0x9e, 0x8d, 0x9e, 0xa0, 0x8a, 0x6f, 0x58, 0x8b,
	0xe4, 0xf7, 0x67, 0x51, 0x72, 0x31, 0x03, 0x2d, 0x93, 0xa8, 0x17, 0x0d, 0x2e, 0xf3, 0xe6, 0x5a,
	0x1b, 0x6e, 0xa4, 0x35, 0x32, 0x47, 0x5d, 0x10, 0xd8, 0x0a, 0x05, 0x77, 0xc9, 0x49, 0x2f, 0x1a,
	0x34, 0xf2, 0xe6, 0x1c, 0x75, 0x5e, 0x4b, 0x01, 0xe1, 0xcb, 0x3d, 0x72, 0x5a, 0x23, 0x7c, 0xb9,
	0x43, 0x6e, 0x59, 0x47, 0x5b, 0x32, 0xa2, 0xb0, 0x40, 0x85, 0x36, 0x12, 0x92, 0xb3, 0x00, 0xb5,
	0x82, 0x3a, 0x02, 0x7a, 0x37, 0x12, 0xf6, 0x8f, 0x9c, 0xe7, 0xe4, 0x5d, 0xd2, 0x38, 0x78, 0xb4,
	0x91, 0xfa, 0x3f, 0x11, 0xeb, 0xbe, 0xa0, 0xf3, 0x84, 0xe5, 0xc2, 0x83, 0x1c, 0xad, 0xc6, 0x6b,
	0xcb, 0x63, 0xe2, 0xa8, 0x51, 0xab, 0x31, 0x77, 0xb3, 0x38, 0x61, 0x17, 0x5f, 0x86, 0x66, 0x40,
	0x2e, 0xf8, 0x69, 0xe4, 0xdb, 0x31, 0x7e, 0x63, 0x1d, 0xd8, 0xf8, 0x2f, 0x44, 0x08, 0x20, 0xb8,
	0x69, 0xde, 0x5f, 0xa7, 0xc7, 0x59, 0xa5, 0xff, 0x72, 0xca, 0xdb, 0x70, 0x38, 0x0e, 0x9f, 0x3e,
	0x1e, 0x15, 0xfa, 0xe9, 0xa2, 0x4c, 0x85, 0x99, 0x67, 0xe1, 0xd6, 0x90, 0xca, 0x76, 0x85, 0x28,
	0xd0, 0x99, 0x2d, 0xef, 0x94, 0xc9, 0x8e, 0x3b, 0x2a, 0xcf, 0x43, 0x39, 0x0f, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0x80, 0x2c, 0x15, 0xbe, 0x01, 0x00, 0x00,
}