// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: discovery.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*Node `protobuf:"bytes,1,rep,name=Peers,proto3" json:"Peers,omitempty"`
}

func (x *NodeList) Reset() {
	*x = NodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeList) ProtoMessage() {}

func (x *NodeList) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeList.ProtoReflect.Descriptor instead.
func (*NodeList) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *NodeList) GetPeers() []*Node {
	if x != nil {
		return x.Peers
	}
	return nil
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Node) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_discovery_proto protoreflect.FileDescriptor

var file_discovery_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x08, 0x4e, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x32, 0x6a, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x0f, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x13,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x45, 0x78, 0x69, 0x74, 0x12, 0x0f, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discovery_proto_rawDescOnce sync.Once
	file_discovery_proto_rawDescData = file_discovery_proto_rawDesc
)

func file_discovery_proto_rawDescGZIP() []byte {
	file_discovery_proto_rawDescOnce.Do(func() {
		file_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_discovery_proto_rawDescData)
	})
	return file_discovery_proto_rawDescData
}

var file_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_discovery_proto_goTypes = []interface{}{
	(*NodeList)(nil),      // 0: discovery.NodeList
	(*Node)(nil),          // 1: discovery.Node
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_discovery_proto_depIdxs = []int32{
	1, // 0: discovery.NodeList.Peers:type_name -> discovery.Node
	1, // 1: discovery.Discovery.Join:input_type -> discovery.Node
	1, // 2: discovery.Discovery.Exit:input_type -> discovery.Node
	0, // 3: discovery.Discovery.Join:output_type -> discovery.NodeList
	2, // 4: discovery.Discovery.Exit:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_discovery_proto_init() }
func file_discovery_proto_init() {
	if File_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeList); i {
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
		file_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
			RawDescriptor: file_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discovery_proto_goTypes,
		DependencyIndexes: file_discovery_proto_depIdxs,
		MessageInfos:      file_discovery_proto_msgTypes,
	}.Build()
	File_discovery_proto = out.File
	file_discovery_proto_rawDesc = nil
	file_discovery_proto_goTypes = nil
	file_discovery_proto_depIdxs = nil
}
