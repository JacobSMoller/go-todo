// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/todo/v1/todo.proto

package todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Todo struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// @inject_tag: sql:",notnull,default:false"
	Completed bool `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	// @inject_tag: sql:"type:timestamptz,default:now()"
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: sql:"type:timestamptz"
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_993c5f6855264bd0, []int{0}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

func (m *Todo) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Todo) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateTodoRequest struct {
	Item                 *Todo    `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_993c5f6855264bd0, []int{1}
}
func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(dst, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetItem() *Todo {
	if m != nil {
		return m.Item
	}
	return nil
}

type CreateTodoResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_993c5f6855264bd0, []int{2}
}
func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(dst, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.v1.Todo")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.v1.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.v1.CreateTodoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.v1.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo/v1/todo.proto",
}

func init() { proto.RegisterFile("api/todo/v1/todo.proto", fileDescriptor_todo_993c5f6855264bd0) }

var fileDescriptor_todo_993c5f6855264bd0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0xfd, 0x6d, 0x7f, 0x69, 0xb5, 0x13, 0x14, 0x5c, 0x44, 0x42, 0x14, 0x8c, 0xc1, 0x43, 0x4e,
	0x1b, 0x1a, 0x41, 0xf0, 0x58, 0x3d, 0x78, 0x8f, 0xc5, 0x83, 0x17, 0x49, 0xb3, 0x63, 0x59, 0x48,
	0xba, 0x6b, 0x32, 0xc9, 0x07, 0xf6, 0x93, 0x48, 0x36, 0x89, 0x2d, 0x2a, 0x78, 0xda, 0xdd, 0xf7,
	0x67, 0xf6, 0xbd, 0x81, 0xb3, 0xcc, 0xa8, 0x98, 0xb4, 0xd4, 0x71, 0xbb, 0xb0, 0xa7, 0x30, 0x95,
	0x26, 0xcd, 0x0f, 0xec, 0xbd, 0x5d, 0xf8, 0x97, 0x1b, 0xad, 0x37, 0x05, 0xc6, 0x16, 0x5e, 0x37,
	0x6f, 0x31, 0xa9, 0x12, 0x6b, 0xca, 0x4a, 0xd3, 0x2b, 0xc3, 0x0f, 0x06, 0xce, 0x4a, 0x4b, 0xcd,
	0x8f, 0x61, 0xa2, 0xa4, 0xc7, 0x02, 0x16, 0x39, 0xe9, 0x44, 0x49, 0x7e, 0x0a, 0x53, 0x52, 0x54,
	0xa0, 0x37, 0x09, 0x58, 0x34, 0x4f, 0xfb, 0x07, 0x0f, 0xc0, 0x95, 0x58, 0xe7, 0x95, 0x32, 0xa4,
	0xf4, 0xd6, 0xfb, 0x6f, 0xb9, 0x7d, 0x88, 0x5f, 0xc0, 0x3c, 0xd7, 0xa5, 0x29, 0x90, 0x50, 0x7a,
	0x4e, 0xc0, 0xa2, 0xc3, 0x74, 0x07, 0xf0, 0x3b, 0x80, 0xbc, 0xc2, 0x8c, 0x50, 0xbe, 0x66, 0xe4,
	0x4d, 0x03, 0x16, 0xb9, 0x89, 0x2f, 0xfa, 0x90, 0x62, 0x0c, 0x29, 0x56, 0x63, 0xc8, 0x74, 0x3e,
	0xa8, 0x97, 0xd4, 0x59, 0x1b, 0x23, 0x47, 0xeb, 0xec, 0x6f, 0xeb, 0xa0, 0x5e, 0x52, 0x78, 0x0b,
	0x27, 0x0f, 0x76, 0x4e, 0xd7, 0x34, 0xc5, 0xf7, 0x06, 0x6b, 0xe2, 0x57, 0xe0, 0x28, 0xc2, 0xd2,
	0x56, 0x76, 0x93, 0x23, 0x31, 0xac, 0x4c, 0x58, 0x8d, 0xa5, 0xc2, 0x6b, 0xe0, 0xfb, 0xbe, 0xda,
	0xe8, 0x6d, 0x8d, 0xdf, 0x37, 0x95, 0x3c, 0x83, 0xdb, 0xf1, 0x4f, 0x58, 0xb5, 0x2a, 0x47, 0xfe,
	0x08, 0xb0, 0x33, 0x71, 0xff, 0x6b, 0xee, 0x8f, 0x04, 0xfe, 0xf9, 0xaf, 0x5c, 0xff, 0x4b, 0xf8,
	0xef, 0x7e, 0xf6, 0xe2, 0x74, 0xfc, 0x7a, 0x66, 0xcb, 0xdd, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x5c, 0xff, 0x0a, 0x0c, 0xed, 0x01, 0x00, 0x00,
}