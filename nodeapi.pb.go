// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.1
// source: nodeapi.proto

package synerex_nodeapi

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NodeType int32

const (
	NodeType_PROVIDER NodeType = 0 // most of node is normal provider
	NodeType_SERVER   NodeType = 1 // node is synerex server.
	NodeType_GATEWAY  NodeType = 2 // node is synerex gateway. (for inter synerex/outer synerex)
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "PROVIDER",
		1: "SERVER",
		2: "GATEWAY",
	}
	NodeType_value = map[string]int32{
		"PROVIDER": 0,
		"SERVER":   1,
		"GATEWAY":  2,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeapi_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_nodeapi_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{0}
}

type KeepAliveCommand int32

const (
	KeepAliveCommand_NONE                KeepAliveCommand = 0
	KeepAliveCommand_RECONNECT           KeepAliveCommand = 1 // node server is now restarted. please reconnect.
	KeepAliveCommand_SERVER_CHANGE       KeepAliveCommand = 2 // There is change in Synerex Server. Obtain server information.
	KeepAliveCommand_PROVIDER_DISCONNECT KeepAliveCommand = 3 // only for Synerex-Server. Provider is disconnected. Remove Channels
)

// Enum value maps for KeepAliveCommand.
var (
	KeepAliveCommand_name = map[int32]string{
		0: "NONE",
		1: "RECONNECT",
		2: "SERVER_CHANGE",
		3: "PROVIDER_DISCONNECT",
	}
	KeepAliveCommand_value = map[string]int32{
		"NONE":                0,
		"RECONNECT":           1,
		"SERVER_CHANGE":       2,
		"PROVIDER_DISCONNECT": 3,
	}
)

func (x KeepAliveCommand) Enum() *KeepAliveCommand {
	p := new(KeepAliveCommand)
	*p = x
	return p
}

func (x KeepAliveCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeepAliveCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_nodeapi_proto_enumTypes[1].Descriptor()
}

func (KeepAliveCommand) Type() protoreflect.EnumType {
	return &file_nodeapi_proto_enumTypes[1]
}

func (x KeepAliveCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeepAliveCommand.Descriptor instead.
func (KeepAliveCommand) EnumDescriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{1}
}

// information for synerex servers and providers, gateways (nodes)
type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName         string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NodeType         NodeType `protobuf:"varint,2,opt,name=node_type,json=nodeType,proto3,enum=nodeapi.NodeType" json:"node_type,omitempty"`    // node is provider/server/gateway?
	ServerInfo       string   `protobuf:"bytes,3,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`                     // for synerex servers/gateways
	NodePbaseVersion string   `protobuf:"bytes,4,opt,name=node_pbase_version,json=nodePbaseVersion,proto3" json:"node_pbase_version,omitempty"` // which protocol base version
	WithNodeId       int32    `protobuf:"varint,5,opt,name=with_node_id,json=withNodeId,proto3" json:"with_node_id,omitempty"`                  // for reconnection with previous node_id (usually -1)
	ClusterId        int32    `protobuf:"varint,6,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`                       //
	AreaId           string   `protobuf:"bytes,7,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`                                 // for area definition
	ChannelTypes     []uint32 `protobuf:"varint,8,rep,packed,name=channelTypes,proto3" json:"channelTypes,omitempty"`                           // used channel list
	GwInfo           string   `protobuf:"bytes,9,opt,name=gw_info,json=gwInfo,proto3" json:"gw_info,omitempty"`                                 // for gateway information
	// for information for controller
	BinVersion    string                 `protobuf:"bytes,10,opt,name=bin_version,json=binVersion,proto3" json:"bin_version,omitempty"` // version of binary
	Count         int32                  `protobuf:"varint,11,opt,name=count,proto3" json:"count,omitempty"`                            // keepalive update count
	LastAliveTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=last_alive_time,json=lastAliveTime,proto3" json:"last_alive_time,omitempty"`
	KeepaliveArg  string                 `protobuf:"bytes,13,opt,name=keepalive_arg,json=keepaliveArg,proto3" json:"keepalive_arg,omitempty"` // keepalive argument
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nodeapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{0}
}

func (x *NodeInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *NodeInfo) GetNodeType() NodeType {
	if x != nil {
		return x.NodeType
	}
	return NodeType_PROVIDER
}

func (x *NodeInfo) GetServerInfo() string {
	if x != nil {
		return x.ServerInfo
	}
	return ""
}

func (x *NodeInfo) GetNodePbaseVersion() string {
	if x != nil {
		return x.NodePbaseVersion
	}
	return ""
}

func (x *NodeInfo) GetWithNodeId() int32 {
	if x != nil {
		return x.WithNodeId
	}
	return 0
}

func (x *NodeInfo) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (x *NodeInfo) GetAreaId() string {
	if x != nil {
		return x.AreaId
	}
	return ""
}

func (x *NodeInfo) GetChannelTypes() []uint32 {
	if x != nil {
		return x.ChannelTypes
	}
	return nil
}

func (x *NodeInfo) GetGwInfo() string {
	if x != nil {
		return x.GwInfo
	}
	return ""
}

func (x *NodeInfo) GetBinVersion() string {
	if x != nil {
		return x.BinVersion
	}
	return ""
}

func (x *NodeInfo) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *NodeInfo) GetLastAliveTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastAliveTime
	}
	return nil
}

func (x *NodeInfo) GetKeepaliveArg() string {
	if x != nil {
		return x.KeepaliveArg
	}
	return ""
}

type NodeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId            int32  `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`                                  // unique id for each node in current node server.
	Secret            uint64 `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`                                               // secret id with node_server (Not used for Query)
	ServerInfo        string `protobuf:"bytes,3,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`                       // synerex server address (only for registration of Server/Gateway)
	KeepaliveDuration int32  `protobuf:"varint,4,opt,name=keepalive_duration,json=keepaliveDuration,proto3" json:"keepalive_duration,omitempty"` // at least make keep alive less than this time.
}

func (x *NodeID) Reset() {
	*x = NodeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeID) ProtoMessage() {}

func (x *NodeID) ProtoReflect() protoreflect.Message {
	mi := &file_nodeapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeID.ProtoReflect.Descriptor instead.
func (*NodeID) Descriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{1}
}

func (x *NodeID) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeID) GetSecret() uint64 {
	if x != nil {
		return x.Secret
	}
	return 0
}

func (x *NodeID) GetServerInfo() string {
	if x != nil {
		return x.ServerInfo
	}
	return ""
}

func (x *NodeID) GetKeepaliveDuration() int32 {
	if x != nil {
		return x.KeepaliveDuration
	}
	return 0
}

type ServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu      float64 `protobuf:"fixed64,1,opt,name=cpu,proto3" json:"cpu,omitempty"`                          // cpu load average
	Memory   float64 `protobuf:"fixed64,2,opt,name=memory,proto3" json:"memory,omitempty"`                    // memory usage rate
	MsgCount uint64  `protobuf:"varint,3,opt,name=msg_count,json=msgCount,proto3" json:"msg_count,omitempty"` // message count
}

func (x *ServerStatus) Reset() {
	*x = ServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus) ProtoMessage() {}

func (x *ServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nodeapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus.ProtoReflect.Descriptor instead.
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStatus) GetCpu() float64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *ServerStatus) GetMemory() float64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *ServerStatus) GetMsgCount() uint64 {
	if x != nil {
		return x.MsgCount
	}
	return 0
}

type NodeUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId      int32         `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`                // unique id for each node in current node server.
	Secret      uint64        `protobuf:"fixed64,2,opt,name=secret,proto3" json:"secret,omitempty"`                             // secret number for each provider (for auth)
	UpdateCount int32         `protobuf:"varint,3,opt,name=update_count,json=updateCount,proto3" json:"update_count,omitempty"` // sequential counter for nodes
	NodeStatus  int32         `protobuf:"varint,4,opt,name=node_status,json=nodeStatus,proto3" json:"node_status,omitempty"`    // running state (health check)
	NodeArg     string        `protobuf:"bytes,5,opt,name=node_arg,json=nodeArg,proto3" json:"node_arg,omitempty"`
	Status      *ServerStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"` // server status (load average)
}

func (x *NodeUpdate) Reset() {
	*x = NodeUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeUpdate) ProtoMessage() {}

func (x *NodeUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_nodeapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeUpdate.ProtoReflect.Descriptor instead.
func (*NodeUpdate) Descriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{3}
}

func (x *NodeUpdate) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeUpdate) GetSecret() uint64 {
	if x != nil {
		return x.Secret
	}
	return 0
}

func (x *NodeUpdate) GetUpdateCount() int32 {
	if x != nil {
		return x.UpdateCount
	}
	return 0
}

func (x *NodeUpdate) GetNodeStatus() int32 {
	if x != nil {
		return x.NodeStatus
	}
	return 0
}

func (x *NodeUpdate) GetNodeArg() string {
	if x != nil {
		return x.NodeArg
	}
	return ""
}

func (x *NodeUpdate) GetStatus() *ServerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool             `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Command KeepAliveCommand `protobuf:"varint,2,opt,name=command,proto3,enum=nodeapi.KeepAliveCommand" json:"command,omitempty"`
	Err     string           `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodeapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_nodeapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_nodeapi_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Response) GetCommand() KeepAliveCommand {
	if x != nil {
		return x.Command
	}
	return KeepAliveCommand_NONE
}

func (x *Response) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_nodeapi_proto protoreflect.FileDescriptor

var file_nodeapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03, 0x0a, 0x08, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x50, 0x62, 0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x77, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x69, 0x6e,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x6c, 0x69, 0x76, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x61, 0x72, 0x67, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x65,
	0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x41, 0x72, 0x67, 0x22, 0x89, 0x01, 0x0a, 0x06, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x0a, 0x12, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x11, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x55, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x73, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xcb, 0x01, 0x0a,
	0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x67, 0x12, 0x2d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x61, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x2a, 0x31, 0x0a,
	0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f,
	0x56, 0x49, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0x02,
	0x2a, 0x57, 0x0a, 0x10, 0x4b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x52, 0x45, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x02,
	0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10, 0x03, 0x32, 0xde, 0x01, 0x0a, 0x04, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x34, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x11, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x4b,
	0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x55, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x65, 0x78,
	0x2f, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x65, 0x78, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodeapi_proto_rawDescOnce sync.Once
	file_nodeapi_proto_rawDescData = file_nodeapi_proto_rawDesc
)

func file_nodeapi_proto_rawDescGZIP() []byte {
	file_nodeapi_proto_rawDescOnce.Do(func() {
		file_nodeapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodeapi_proto_rawDescData)
	})
	return file_nodeapi_proto_rawDescData
}

var file_nodeapi_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_nodeapi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nodeapi_proto_goTypes = []interface{}{
	(NodeType)(0),                 // 0: nodeapi.NodeType
	(KeepAliveCommand)(0),         // 1: nodeapi.KeepAliveCommand
	(*NodeInfo)(nil),              // 2: nodeapi.NodeInfo
	(*NodeID)(nil),                // 3: nodeapi.NodeID
	(*ServerStatus)(nil),          // 4: nodeapi.ServerStatus
	(*NodeUpdate)(nil),            // 5: nodeapi.NodeUpdate
	(*Response)(nil),              // 6: nodeapi.Response
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_nodeapi_proto_depIdxs = []int32{
	0, // 0: nodeapi.NodeInfo.node_type:type_name -> nodeapi.NodeType
	7, // 1: nodeapi.NodeInfo.last_alive_time:type_name -> google.protobuf.Timestamp
	4, // 2: nodeapi.NodeUpdate.status:type_name -> nodeapi.ServerStatus
	1, // 3: nodeapi.Response.command:type_name -> nodeapi.KeepAliveCommand
	2, // 4: nodeapi.Node.RegisterNode:input_type -> nodeapi.NodeInfo
	3, // 5: nodeapi.Node.QueryNode:input_type -> nodeapi.NodeID
	5, // 6: nodeapi.Node.KeepAlive:input_type -> nodeapi.NodeUpdate
	3, // 7: nodeapi.Node.UnRegisterNode:input_type -> nodeapi.NodeID
	3, // 8: nodeapi.Node.RegisterNode:output_type -> nodeapi.NodeID
	2, // 9: nodeapi.Node.QueryNode:output_type -> nodeapi.NodeInfo
	6, // 10: nodeapi.Node.KeepAlive:output_type -> nodeapi.Response
	6, // 11: nodeapi.Node.UnRegisterNode:output_type -> nodeapi.Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nodeapi_proto_init() }
func file_nodeapi_proto_init() {
	if File_nodeapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodeapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nodeapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nodeapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nodeapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_nodeapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nodeapi_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodeapi_proto_goTypes,
		DependencyIndexes: file_nodeapi_proto_depIdxs,
		EnumInfos:         file_nodeapi_proto_enumTypes,
		MessageInfos:      file_nodeapi_proto_msgTypes,
	}.Build()
	File_nodeapi_proto = out.File
	file_nodeapi_proto_rawDesc = nil
	file_nodeapi_proto_goTypes = nil
	file_nodeapi_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	RegisterNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeID, error)
	QueryNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*NodeInfo, error)
	KeepAlive(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*Response, error)
	UnRegisterNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*Response, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) RegisterNode(ctx context.Context, in *NodeInfo, opts ...grpc.CallOption) (*NodeID, error) {
	out := new(NodeID)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) QueryNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*NodeInfo, error) {
	out := new(NodeInfo)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/QueryNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) KeepAlive(ctx context.Context, in *NodeUpdate, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) UnRegisterNode(ctx context.Context, in *NodeID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/nodeapi.Node/UnRegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	RegisterNode(context.Context, *NodeInfo) (*NodeID, error)
	QueryNode(context.Context, *NodeID) (*NodeInfo, error)
	KeepAlive(context.Context, *NodeUpdate) (*Response, error)
	UnRegisterNode(context.Context, *NodeID) (*Response, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) RegisterNode(context.Context, *NodeInfo) (*NodeID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedNodeServer) QueryNode(context.Context, *NodeID) (*NodeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNode not implemented")
}
func (*UnimplementedNodeServer) KeepAlive(context.Context, *NodeUpdate) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (*UnimplementedNodeServer) UnRegisterNode(context.Context, *NodeID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegisterNode not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).RegisterNode(ctx, req.(*NodeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_QueryNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).QueryNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/QueryNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).QueryNode(ctx, req.(*NodeID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).KeepAlive(ctx, req.(*NodeUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_UnRegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).UnRegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodeapi.Node/UnRegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).UnRegisterNode(ctx, req.(*NodeID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodeapi.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _Node_RegisterNode_Handler,
		},
		{
			MethodName: "QueryNode",
			Handler:    _Node_QueryNode_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Node_KeepAlive_Handler,
		},
		{
			MethodName: "UnRegisterNode",
			Handler:    _Node_UnRegisterNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodeapi.proto",
}
