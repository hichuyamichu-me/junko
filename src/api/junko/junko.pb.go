// Code generated by protoc-gen-go. DO NOT EDIT.
// source: junko.proto

package junko

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CommandsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandsRequest) Reset()         { *m = CommandsRequest{} }
func (m *CommandsRequest) String() string { return proto.CompactTextString(m) }
func (*CommandsRequest) ProtoMessage()    {}
func (*CommandsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c59d0c72b5ab08, []int{0}
}

func (m *CommandsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandsRequest.Unmarshal(m, b)
}
func (m *CommandsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandsRequest.Marshal(b, m, deterministic)
}
func (m *CommandsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandsRequest.Merge(m, src)
}
func (m *CommandsRequest) XXX_Size() int {
	return xxx_messageInfo_CommandsRequest.Size(m)
}
func (m *CommandsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandsRequest proto.InternalMessageInfo

type CommandsResponce struct {
	Commands             []*Command `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CommandsResponce) Reset()         { *m = CommandsResponce{} }
func (m *CommandsResponce) String() string { return proto.CompactTextString(m) }
func (*CommandsResponce) ProtoMessage()    {}
func (*CommandsResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c59d0c72b5ab08, []int{1}
}

func (m *CommandsResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandsResponce.Unmarshal(m, b)
}
func (m *CommandsResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandsResponce.Marshal(b, m, deterministic)
}
func (m *CommandsResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandsResponce.Merge(m, src)
}
func (m *CommandsResponce) XXX_Size() int {
	return xxx_messageInfo_CommandsResponce.Size(m)
}
func (m *CommandsResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandsResponce.DiscardUnknown(m)
}

var xxx_messageInfo_CommandsResponce proto.InternalMessageInfo

func (m *CommandsResponce) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

type Command struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Usage                string   `protobuf:"bytes,3,opt,name=usage,proto3" json:"usage,omitempty"`
	Examples             []string `protobuf:"bytes,4,rep,name=examples,proto3" json:"examples,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_41c59d0c72b5ab08, []int{2}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Command) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Command) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *Command) GetExamples() []string {
	if m != nil {
		return m.Examples
	}
	return nil
}

func init() {
	proto.RegisterType((*CommandsRequest)(nil), "junko.CommandsRequest")
	proto.RegisterType((*CommandsResponce)(nil), "junko.CommandsResponce")
	proto.RegisterType((*Command)(nil), "junko.Command")
}

func init() {
	proto.RegisterFile("junko.proto", fileDescriptor_41c59d0c72b5ab08)
}

var fileDescriptor_41c59d0c72b5ab08 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x2a, 0xcd, 0xcb,
	0xce, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x04, 0xb9, 0xf8, 0x9d,
	0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x8a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xec,
	0xb8, 0x04, 0x10, 0x42, 0xc5, 0x05, 0xf9, 0x79, 0xc9, 0xa9, 0x42, 0x5a, 0x5c, 0x1c, 0xc9, 0x50,
	0x31, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0x88, 0x69, 0x50, 0xa5, 0x41, 0x70,
	0x79, 0xa5, 0x42, 0x2e, 0x76, 0xa8, 0xa0, 0x90, 0x14, 0x17, 0x47, 0x72, 0x62, 0x49, 0x6a, 0x7a,
	0x7e, 0x51, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc1, 0xc5, 0x9e,
	0x9c, 0x9f, 0x57, 0x92, 0x9a, 0x57, 0x22, 0xc1, 0x04, 0x96, 0x82, 0x71, 0x85, 0x44, 0xb8, 0x58,
	0x4b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x98, 0xc1, 0xe2, 0x10, 0x0e, 0xc8, 0xac, 0xd4, 0x8a, 0xc4,
	0xdc, 0x82, 0x9c, 0xd4, 0x62, 0x09, 0x16, 0x05, 0x66, 0x90, 0x59, 0x30, 0xbe, 0x91, 0x27, 0x17,
	0xab, 0x17, 0xc8, 0x35, 0x42, 0x0e, 0x5c, 0xdc, 0xe9, 0xa9, 0x25, 0x30, 0xe7, 0x0b, 0x89, 0xa1,
	0x3a, 0x12, 0xe6, 0x45, 0x29, 0x71, 0x0c, 0x71, 0x88, 0x3f, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xc1,
	0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x1f, 0x9b, 0xaa, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JunkoClient is the client API for Junko service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JunkoClient interface {
	GetCommands(ctx context.Context, in *CommandsRequest, opts ...grpc.CallOption) (*CommandsResponce, error)
}

type junkoClient struct {
	cc grpc.ClientConnInterface
}

func NewJunkoClient(cc grpc.ClientConnInterface) JunkoClient {
	return &junkoClient{cc}
}

func (c *junkoClient) GetCommands(ctx context.Context, in *CommandsRequest, opts ...grpc.CallOption) (*CommandsResponce, error) {
	out := new(CommandsResponce)
	err := c.cc.Invoke(ctx, "/junko.Junko/getCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JunkoServer is the server API for Junko service.
type JunkoServer interface {
	GetCommands(context.Context, *CommandsRequest) (*CommandsResponce, error)
}

// UnimplementedJunkoServer can be embedded to have forward compatible implementations.
type UnimplementedJunkoServer struct {
}

func (*UnimplementedJunkoServer) GetCommands(ctx context.Context, req *CommandsRequest) (*CommandsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommands not implemented")
}

func RegisterJunkoServer(s *grpc.Server, srv JunkoServer) {
	s.RegisterService(&_Junko_serviceDesc, srv)
}

func _Junko_GetCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JunkoServer).GetCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/junko.Junko/GetCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JunkoServer).GetCommands(ctx, req.(*CommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Junko_serviceDesc = grpc.ServiceDesc{
	ServiceName: "junko.Junko",
	HandlerType: (*JunkoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCommands",
			Handler:    _Junko_GetCommands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junko.proto",
}
