// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package osapi

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
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

type ContainerDriver int32

const (
	ContainerDriver_CONTAINERD ContainerDriver = 0
	ContainerDriver_CRI        ContainerDriver = 1
)

var ContainerDriver_name = map[int32]string{
	0: "CONTAINERD",
	1: "CRI",
}

var ContainerDriver_value = map[string]int32{
	"CONTAINERD": 0,
	"CRI":        1,
}

func (x ContainerDriver) String() string {
	return proto.EnumName(ContainerDriver_name, int32(x))
}

func (ContainerDriver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

// The request message containing the containerd namespace.
type ContainersRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// driver might be default "containerd" or "cri"
	Driver               ContainerDriver `protobuf:"varint,2,opt,name=driver,proto3,enum=proto.ContainerDriver" json:"driver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ContainersRequest) Reset()         { *m = ContainersRequest{} }
func (m *ContainersRequest) String() string { return proto.CompactTextString(m) }
func (*ContainersRequest) ProtoMessage()    {}
func (*ContainersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *ContainersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainersRequest.Unmarshal(m, b)
}

func (m *ContainersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainersRequest.Marshal(b, m, deterministic)
}

func (m *ContainersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainersRequest.Merge(m, src)
}

func (m *ContainersRequest) XXX_Size() int {
	return xxx_messageInfo_ContainersRequest.Size(m)
}

func (m *ContainersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainersRequest proto.InternalMessageInfo

func (m *ContainersRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ContainersRequest) GetDriver() ContainerDriver {
	if m != nil {
		return m.Driver
	}
	return ContainerDriver_CONTAINERD
}

// The response message containing the requested containers.
type ContainersReply struct {
	Containers           []*Container `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContainersReply) Reset()         { *m = ContainersReply{} }
func (m *ContainersReply) String() string { return proto.CompactTextString(m) }
func (*ContainersReply) ProtoMessage()    {}
func (*ContainersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *ContainersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainersReply.Unmarshal(m, b)
}

func (m *ContainersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainersReply.Marshal(b, m, deterministic)
}

func (m *ContainersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainersReply.Merge(m, src)
}

func (m *ContainersReply) XXX_Size() int {
	return xxx_messageInfo_ContainersReply.Size(m)
}

func (m *ContainersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainersReply.DiscardUnknown(m)
}

var xxx_messageInfo_ContainersReply proto.InternalMessageInfo

func (m *ContainersReply) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

// The response message containing the requested containers.
type Container struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Pid                  uint32   `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	PodId                string   `protobuf:"bytes,6,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}

func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}

func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}

func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}

func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *Container) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Container) GetPodId() string {
	if m != nil {
		return m.PodId
	}
	return ""
}

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message containing the containerd namespace.
type StatsRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// driver might be default "containerd" or "cri"
	Driver               ContainerDriver `protobuf:"varint,2,opt,name=driver,proto3,enum=proto.ContainerDriver" json:"driver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatsRequest) Reset()         { *m = StatsRequest{} }
func (m *StatsRequest) String() string { return proto.CompactTextString(m) }
func (*StatsRequest) ProtoMessage()    {}
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *StatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsRequest.Unmarshal(m, b)
}

func (m *StatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsRequest.Marshal(b, m, deterministic)
}

func (m *StatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsRequest.Merge(m, src)
}

func (m *StatsRequest) XXX_Size() int {
	return xxx_messageInfo_StatsRequest.Size(m)
}

func (m *StatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsRequest proto.InternalMessageInfo

func (m *StatsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *StatsRequest) GetDriver() ContainerDriver {
	if m != nil {
		return m.Driver
	}
	return ContainerDriver_CONTAINERD
}

// The response message containing the requested stats.
type StatsReply struct {
	Stats                []*Stat  `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsReply) Reset()         { *m = StatsReply{} }
func (m *StatsReply) String() string { return proto.CompactTextString(m) }
func (*StatsReply) ProtoMessage()    {}
func (*StatsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *StatsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsReply.Unmarshal(m, b)
}

func (m *StatsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsReply.Marshal(b, m, deterministic)
}

func (m *StatsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsReply.Merge(m, src)
}

func (m *StatsReply) XXX_Size() int {
	return xxx_messageInfo_StatsReply.Size(m)
}

func (m *StatsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatsReply proto.InternalMessageInfo

func (m *StatsReply) GetStats() []*Stat {
	if m != nil {
		return m.Stats
	}
	return nil
}

// The response message containing the requested stat.
type Stat struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	MemoryUsage          uint64   `protobuf:"varint,4,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	CpuUsage             uint64   `protobuf:"varint,5,opt,name=cpu_usage,json=cpuUsage,proto3" json:"cpu_usage,omitempty"`
	PodId                string   `protobuf:"bytes,6,opt,name=pod_id,json=podId,proto3" json:"pod_id,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stat) Reset()         { *m = Stat{} }
func (m *Stat) String() string { return proto.CompactTextString(m) }
func (*Stat) ProtoMessage()    {}
func (*Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stat.Unmarshal(m, b)
}

func (m *Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stat.Marshal(b, m, deterministic)
}

func (m *Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stat.Merge(m, src)
}

func (m *Stat) XXX_Size() int {
	return xxx_messageInfo_Stat.Size(m)
}

func (m *Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Stat proto.InternalMessageInfo

func (m *Stat) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Stat) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Stat) GetMemoryUsage() uint64 {
	if m != nil {
		return m.MemoryUsage
	}
	return 0
}

func (m *Stat) GetCpuUsage() uint64 {
	if m != nil {
		return m.CpuUsage
	}
	return 0
}

func (m *Stat) GetPodId() string {
	if m != nil {
		return m.PodId
	}
	return ""
}

func (m *Stat) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message containing the process to restart.
type RestartRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// driver might be default "containerd" or "cri"
	Driver               ContainerDriver `protobuf:"varint,3,opt,name=driver,proto3,enum=proto.ContainerDriver" json:"driver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RestartRequest) Reset()         { *m = RestartRequest{} }
func (m *RestartRequest) String() string { return proto.CompactTextString(m) }
func (*RestartRequest) ProtoMessage()    {}
func (*RestartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *RestartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartRequest.Unmarshal(m, b)
}

func (m *RestartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartRequest.Marshal(b, m, deterministic)
}

func (m *RestartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartRequest.Merge(m, src)
}

func (m *RestartRequest) XXX_Size() int {
	return xxx_messageInfo_RestartRequest.Size(m)
}

func (m *RestartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestartRequest proto.InternalMessageInfo

func (m *RestartRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RestartRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RestartRequest) GetDriver() ContainerDriver {
	if m != nil {
		return m.Driver
	}
	return ContainerDriver_CONTAINERD
}

// The response message containing the restart status.
type RestartReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestartReply) Reset()         { *m = RestartReply{} }
func (m *RestartReply) String() string { return proto.CompactTextString(m) }
func (*RestartReply) ProtoMessage()    {}
func (*RestartReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *RestartReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartReply.Unmarshal(m, b)
}

func (m *RestartReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartReply.Marshal(b, m, deterministic)
}

func (m *RestartReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartReply.Merge(m, src)
}

func (m *RestartReply) XXX_Size() int {
	return xxx_messageInfo_RestartReply.Size(m)
}

func (m *RestartReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartReply.DiscardUnknown(m)
}

var xxx_messageInfo_RestartReply proto.InternalMessageInfo

// The request message containing the process name.
type LogsRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// driver might be default "containerd" or "cri"
	Driver               ContainerDriver `protobuf:"varint,3,opt,name=driver,proto3,enum=proto.ContainerDriver" json:"driver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LogsRequest) Reset()         { *m = LogsRequest{} }
func (m *LogsRequest) String() string { return proto.CompactTextString(m) }
func (*LogsRequest) ProtoMessage()    {}
func (*LogsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *LogsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogsRequest.Unmarshal(m, b)
}

func (m *LogsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogsRequest.Marshal(b, m, deterministic)
}

func (m *LogsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogsRequest.Merge(m, src)
}

func (m *LogsRequest) XXX_Size() int {
	return xxx_messageInfo_LogsRequest.Size(m)
}

func (m *LogsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogsRequest proto.InternalMessageInfo

func (m *LogsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LogsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogsRequest) GetDriver() ContainerDriver {
	if m != nil {
		return m.Driver
	}
	return ContainerDriver_CONTAINERD
}

// The response message containing the requested logs.
type Data struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}

func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}

func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}

func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}

func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type ProcessesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessesRequest) Reset()         { *m = ProcessesRequest{} }
func (m *ProcessesRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessesRequest) ProtoMessage()    {}
func (*ProcessesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *ProcessesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessesRequest.Unmarshal(m, b)
}

func (m *ProcessesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessesRequest.Marshal(b, m, deterministic)
}

func (m *ProcessesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessesRequest.Merge(m, src)
}

func (m *ProcessesRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessesRequest.Size(m)
}

func (m *ProcessesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessesRequest proto.InternalMessageInfo

type ProcessesReply struct {
	ProcessList          *ProcessList `protobuf:"bytes,1,opt,name=process_list,json=processList,proto3" json:"process_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProcessesReply) Reset()         { *m = ProcessesReply{} }
func (m *ProcessesReply) String() string { return proto.CompactTextString(m) }
func (*ProcessesReply) ProtoMessage()    {}
func (*ProcessesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *ProcessesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessesReply.Unmarshal(m, b)
}

func (m *ProcessesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessesReply.Marshal(b, m, deterministic)
}

func (m *ProcessesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessesReply.Merge(m, src)
}

func (m *ProcessesReply) XXX_Size() int {
	return xxx_messageInfo_ProcessesReply.Size(m)
}

func (m *ProcessesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessesReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessesReply proto.InternalMessageInfo

func (m *ProcessesReply) GetProcessList() *ProcessList {
	if m != nil {
		return m.ProcessList
	}
	return nil
}

type ProcessList struct {
	Bytes                []byte   `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessList) Reset()         { *m = ProcessList{} }
func (m *ProcessList) String() string { return proto.CompactTextString(m) }
func (*ProcessList) ProtoMessage()    {}
func (*ProcessList) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *ProcessList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessList.Unmarshal(m, b)
}

func (m *ProcessList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessList.Marshal(b, m, deterministic)
}

func (m *ProcessList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessList.Merge(m, src)
}

func (m *ProcessList) XXX_Size() int {
	return xxx_messageInfo_ProcessList.Size(m)
}

func (m *ProcessList) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessList.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessList proto.InternalMessageInfo

func (m *ProcessList) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ContainerDriver", ContainerDriver_name, ContainerDriver_value)
	proto.RegisterType((*ContainersRequest)(nil), "proto.ContainersRequest")
	proto.RegisterType((*ContainersReply)(nil), "proto.ContainersReply")
	proto.RegisterType((*Container)(nil), "proto.Container")
	proto.RegisterType((*StatsRequest)(nil), "proto.StatsRequest")
	proto.RegisterType((*StatsReply)(nil), "proto.StatsReply")
	proto.RegisterType((*Stat)(nil), "proto.Stat")
	proto.RegisterType((*RestartRequest)(nil), "proto.RestartRequest")
	proto.RegisterType((*RestartReply)(nil), "proto.RestartReply")
	proto.RegisterType((*LogsRequest)(nil), "proto.LogsRequest")
	proto.RegisterType((*Data)(nil), "proto.Data")
	proto.RegisterType((*ProcessesRequest)(nil), "proto.ProcessesRequest")
	proto.RegisterType((*ProcessesReply)(nil), "proto.ProcessesReply")
	proto.RegisterType((*ProcessList)(nil), "proto.ProcessList")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x27, 0x6d, 0xd3, 0x91, 0x97, 0x52, 0x3a, 0x8f, 0x55, 0xd1, 0xb6, 0x43, 0x17, 0x2e, 0x65,
	0x42, 0xe9, 0x28, 0xe2, 0xc2, 0x01, 0x69, 0x6b, 0x27, 0x34, 0x31, 0x6d, 0x93, 0x07, 0x17, 0x84,
	0x34, 0xb9, 0x8d, 0x57, 0x59, 0x34, 0xb5, 0x89, 0x1d, 0xa4, 0x7e, 0x1a, 0x2e, 0x7c, 0x44, 0x3e,
	0x00, 0xb2, 0x9d, 0xa6, 0x69, 0x35, 0x44, 0x77, 0x80, 0x53, 0xf2, 0x7e, 0xef, 0xf7, 0xfe, 0xf8,
	0xf7, 0xec, 0x07, 0x1e, 0x11, 0x2c, 0x12, 0x29, 0x57, 0x1c, 0xb9, 0xe6, 0xb3, 0xb7, 0x3f, 0xe1,
	0x7c, 0x32, 0xa5, 0x3d, 0x63, 0x8d, 0xb2, 0xbb, 0x1e, 0x4d, 0x84, 0x9a, 0x5b, 0x4e, 0x48, 0x60,
	0x7b, 0xc0, 0x67, 0x8a, 0xb0, 0x19, 0x4d, 0x25, 0xa6, 0xdf, 0x32, 0x2a, 0x15, 0x3a, 0x00, 0x6f,
	0x46, 0x12, 0x2a, 0x05, 0x19, 0xd3, 0xc0, 0xe9, 0x38, 0x5d, 0x0f, 0x2f, 0x01, 0x14, 0x41, 0x3d,
	0x4e, 0xd9, 0x77, 0x9a, 0x06, 0x95, 0x8e, 0xd3, 0x6d, 0xf6, 0xdb, 0x36, 0x55, 0x54, 0xe4, 0x19,
	0x1a, 0x2f, 0xce, 0x59, 0xe1, 0x00, 0x9e, 0x96, 0x4b, 0x88, 0xe9, 0x1c, 0x1d, 0x03, 0x8c, 0x0b,
	0x28, 0x70, 0x3a, 0xd5, 0xae, 0xdf, 0x6f, 0xad, 0xa7, 0xc1, 0x25, 0x4e, 0xf8, 0xd3, 0x01, 0xaf,
	0xf0, 0xfc, 0xa5, 0xc1, 0x26, 0x54, 0x58, 0x6c, 0x9a, 0xf3, 0x70, 0x85, 0xc5, 0xe8, 0x19, 0xb8,
	0x2c, 0x21, 0x13, 0x1a, 0x54, 0x0d, 0x64, 0x0d, 0xd4, 0x82, 0xaa, 0x60, 0x71, 0x50, 0xeb, 0x38,
	0xdd, 0x27, 0x58, 0xff, 0xa2, 0x36, 0xd4, 0xa5, 0x22, 0x2a, 0x93, 0x81, 0x6b, 0x88, 0xb9, 0x85,
	0x76, 0xa1, 0x2e, 0x78, 0x7c, 0xcb, 0xe2, 0xa0, 0x6e, 0x13, 0x08, 0x1e, 0x9f, 0xc7, 0x08, 0x41,
	0x4d, 0xd7, 0x0c, 0xb6, 0x0c, 0x68, 0xfe, 0xc3, 0x2f, 0xd0, 0xb8, 0x51, 0x44, 0xfd, 0x23, 0x25,
	0x7b, 0x00, 0x79, 0x76, 0x2d, 0xe2, 0x21, 0xb8, 0xba, 0xc1, 0x85, 0x7e, 0x7e, 0x1e, 0xac, 0x19,
	0xd8, 0x7a, 0xc2, 0x1f, 0x0e, 0xd4, 0xb4, 0xfd, 0x40, 0xc1, 0x0e, 0xa1, 0x91, 0xd0, 0x84, 0xa7,
	0xf3, 0xdb, 0x4c, 0x6a, 0xdd, 0xb4, 0x46, 0x35, 0xec, 0x5b, 0xec, 0x93, 0x86, 0xd0, 0x3e, 0x78,
	0x63, 0x91, 0xe5, 0x7e, 0xd7, 0xf8, 0x1f, 0x8f, 0x45, 0x66, 0x9d, 0x0f, 0x10, 0x6c, 0x06, 0x4d,
	0x4c, 0xa5, 0x22, 0xa9, 0xda, 0x4c, 0xb2, 0xf5, 0x56, 0x97, 0x12, 0x56, 0x37, 0x92, 0xb0, 0x09,
	0x8d, 0xa2, 0x9e, 0x98, 0xce, 0xc3, 0xaf, 0xe0, 0x5f, 0xf0, 0x89, 0xfc, 0x3f, 0xc5, 0x0f, 0xa0,
	0x36, 0x24, 0x8a, 0xe8, 0x0b, 0x39, 0x9a, 0x2b, 0x2a, 0x4d, 0x85, 0x06, 0xb6, 0x46, 0x88, 0xa0,
	0x75, 0x9d, 0xf2, 0x31, 0x95, 0x92, 0x2e, 0xfa, 0x09, 0xdf, 0x43, 0xb3, 0x84, 0xe9, 0xa9, 0xbf,
	0x81, 0x86, 0xb0, 0xc8, 0xed, 0x94, 0x49, 0x65, 0x52, 0xf8, 0x7d, 0x94, 0x57, 0xce, 0xc9, 0x17,
	0x4c, 0x2a, 0xec, 0x8b, 0xa5, 0x11, 0x3e, 0x07, 0xbf, 0xe4, 0xbb, 0xbf, 0x83, 0xa3, 0xa3, 0xd2,
	0x4b, 0xb5, 0xad, 0xa3, 0x26, 0xc0, 0xe0, 0xea, 0xf2, 0xe3, 0xc9, 0xf9, 0xe5, 0x19, 0x1e, 0xb6,
	0x1e, 0xa1, 0x2d, 0xa8, 0x0e, 0xf0, 0x79, 0xcb, 0xe9, 0xff, 0xaa, 0x40, 0xe5, 0xea, 0x06, 0xbd,
	0x04, 0x77, 0x98, 0x50, 0x39, 0x41, 0xed, 0xc8, 0xae, 0x99, 0x68, 0xb1, 0x66, 0xa2, 0x33, 0xbd,
	0x66, 0xf6, 0x16, 0xd7, 0xd2, 0x1c, 0xfc, 0x15, 0xc0, 0x87, 0x6c, 0x44, 0xc7, 0x7c, 0x76, 0xc7,
	0x36, 0x0c, 0x79, 0x01, 0x35, 0x3d, 0x20, 0xb4, 0x38, 0x61, 0x69, 0x5a, 0x2b, 0xc4, 0x63, 0x07,
	0xbd, 0x03, 0x58, 0x2e, 0x1a, 0x14, 0xac, 0x0f, 0xa3, 0x08, 0x6b, 0xdf, 0xe3, 0xb1, 0xd2, 0x6e,
	0xe5, 0x77, 0x03, 0xed, 0xe6, 0x94, 0xd5, 0xbb, 0xb9, 0xb7, 0xb3, 0x0e, 0xeb, 0xb0, 0x1e, 0xb8,
	0xe6, 0x55, 0xa2, 0x9d, 0xd2, 0x0b, 0x2c, 0x8a, 0x6d, 0xaf, 0x82, 0x3a, 0xe0, 0x2d, 0x78, 0xc5,
	0x50, 0xff, 0x28, 0xc2, 0xee, 0xea, 0x44, 0xf3, 0xf1, 0x9f, 0x1e, 0xe8, 0xcd, 0x99, 0x44, 0x5c,
	0x46, 0x44, 0xb0, 0x53, 0xf7, 0x4a, 0x9e, 0x08, 0x76, 0xed, 0x7c, 0x76, 0xb9, 0x24, 0x82, 0x8d,
	0xea, 0x26, 0xe6, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xa5, 0x03, 0x8b, 0x05, 0x06,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OSClient is the client API for OS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OSClient interface {
	Dmesg(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error)
	Kubeconfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error)
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (OS_LogsClient, error)
	Containers(ctx context.Context, in *ContainersRequest, opts ...grpc.CallOption) (*ContainersReply, error)
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartReply, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error)
	Processes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProcessesReply, error)
}

type oSClient struct {
	cc *grpc.ClientConn
}

func NewOSClient(cc *grpc.ClientConn) OSClient {
	return &oSClient{cc}
}

func (c *oSClient) Dmesg(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/proto.OS/Dmesg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSClient) Kubeconfig(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/proto.OS/Kubeconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (OS_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OS_serviceDesc.Streams[0], "/proto.OS/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &oSLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OS_LogsClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type oSLogsClient struct {
	grpc.ClientStream
}

func (x *oSLogsClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oSClient) Containers(ctx context.Context, in *ContainersRequest, opts ...grpc.CallOption) (*ContainersReply, error) {
	out := new(ContainersReply)
	err := c.cc.Invoke(ctx, "/proto.OS/Containers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartReply, error) {
	out := new(RestartReply)
	err := c.cc.Invoke(ctx, "/proto.OS/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsReply, error) {
	out := new(StatsReply)
	err := c.cc.Invoke(ctx, "/proto.OS/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSClient) Processes(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ProcessesReply, error) {
	out := new(ProcessesReply)
	err := c.cc.Invoke(ctx, "/proto.OS/Processes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OSServer is the server API for OS service.
type OSServer interface {
	Dmesg(context.Context, *empty.Empty) (*Data, error)
	Kubeconfig(context.Context, *empty.Empty) (*Data, error)
	Logs(*LogsRequest, OS_LogsServer) error
	Containers(context.Context, *ContainersRequest) (*ContainersReply, error)
	Restart(context.Context, *RestartRequest) (*RestartReply, error)
	Stats(context.Context, *StatsRequest) (*StatsReply, error)
	Processes(context.Context, *empty.Empty) (*ProcessesReply, error)
}

func RegisterOSServer(s *grpc.Server, srv OSServer) {
	s.RegisterService(&_OS_serviceDesc, srv)
}

func _OS_Dmesg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Dmesg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OS/Dmesg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Dmesg(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OS_Kubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Kubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OS/Kubeconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Kubeconfig(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OS_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OSServer).Logs(m, &oSLogsServer{stream})
}

type OS_LogsServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type oSLogsServer struct {
	grpc.ServerStream
}

func (x *oSLogsServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func _OS_Containers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Containers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OS/Containers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Containers(ctx, req.(*ContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OS_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OS/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OS_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OS/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OS_Processes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Processes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OS/Processes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Processes(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _OS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OS",
	HandlerType: (*OSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dmesg",
			Handler:    _OS_Dmesg_Handler,
		},
		{
			MethodName: "Kubeconfig",
			Handler:    _OS_Kubeconfig_Handler,
		},
		{
			MethodName: "Containers",
			Handler:    _OS_Containers_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _OS_Restart_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _OS_Stats_Handler,
		},
		{
			MethodName: "Processes",
			Handler:    _OS_Processes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _OS_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
