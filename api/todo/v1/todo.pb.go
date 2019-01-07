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
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{0}
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

type Owner struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{1}
}
func (m *Owner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Owner.Unmarshal(m, b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
}
func (dst *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(dst, src)
}
func (m *Owner) XXX_Size() int {
	return xxx_messageInfo_Owner.Size(m)
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Owner) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Owner) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
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
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{2}
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
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{3}
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

type DeleteTodoRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoRequest) Reset()         { *m = DeleteTodoRequest{} }
func (m *DeleteTodoRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoRequest) ProtoMessage()    {}
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{4}
}
func (m *DeleteTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoRequest.Unmarshal(m, b)
}
func (m *DeleteTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoRequest.Merge(dst, src)
}
func (m *DeleteTodoRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoRequest.Size(m)
}
func (m *DeleteTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoRequest proto.InternalMessageInfo

func (m *DeleteTodoRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteTodoResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTodoResponse) Reset()         { *m = DeleteTodoResponse{} }
func (m *DeleteTodoResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTodoResponse) ProtoMessage()    {}
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{5}
}
func (m *DeleteTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTodoResponse.Unmarshal(m, b)
}
func (m *DeleteTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTodoResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTodoResponse.Merge(dst, src)
}
func (m *DeleteTodoResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTodoResponse.Size(m)
}
func (m *DeleteTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTodoResponse proto.InternalMessageInfo

type CreateOwnerRequest struct {
	Owner                *Owner   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOwnerRequest) Reset()         { *m = CreateOwnerRequest{} }
func (m *CreateOwnerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOwnerRequest) ProtoMessage()    {}
func (*CreateOwnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{6}
}
func (m *CreateOwnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOwnerRequest.Unmarshal(m, b)
}
func (m *CreateOwnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOwnerRequest.Marshal(b, m, deterministic)
}
func (dst *CreateOwnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOwnerRequest.Merge(dst, src)
}
func (m *CreateOwnerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOwnerRequest.Size(m)
}
func (m *CreateOwnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOwnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOwnerRequest proto.InternalMessageInfo

func (m *CreateOwnerRequest) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

type CreateOwnerResponse struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOwnerResponse) Reset()         { *m = CreateOwnerResponse{} }
func (m *CreateOwnerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOwnerResponse) ProtoMessage()    {}
func (*CreateOwnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_95c0e673bb0c2a0b, []int{7}
}
func (m *CreateOwnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOwnerResponse.Unmarshal(m, b)
}
func (m *CreateOwnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOwnerResponse.Marshal(b, m, deterministic)
}
func (dst *CreateOwnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOwnerResponse.Merge(dst, src)
}
func (m *CreateOwnerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOwnerResponse.Size(m)
}
func (m *CreateOwnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOwnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOwnerResponse proto.InternalMessageInfo

func (m *CreateOwnerResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Todo)(nil), "todo.v1.Todo")
	proto.RegisterType((*Owner)(nil), "todo.v1.Owner")
	proto.RegisterType((*CreateTodoRequest)(nil), "todo.v1.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "todo.v1.CreateTodoResponse")
	proto.RegisterType((*DeleteTodoRequest)(nil), "todo.v1.DeleteTodoRequest")
	proto.RegisterType((*DeleteTodoResponse)(nil), "todo.v1.DeleteTodoResponse")
	proto.RegisterType((*CreateOwnerRequest)(nil), "todo.v1.CreateOwnerRequest")
	proto.RegisterType((*CreateOwnerResponse)(nil), "todo.v1.CreateOwnerResponse")
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
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error)
	CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*CreateOwnerResponse, error)
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

func (c *todoServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*DeleteTodoResponse, error) {
	out := new(DeleteTodoResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*CreateOwnerResponse, error) {
	out := new(CreateOwnerResponse)
	err := c.cc.Invoke(ctx, "/todo.v1.TodoService/CreateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	DeleteTodo(context.Context, *DeleteTodoRequest) (*DeleteTodoResponse, error)
	CreateOwner(context.Context, *CreateOwnerRequest) (*CreateOwnerResponse, error)
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

func _TodoService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.v1.TodoService/CreateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateOwner(ctx, req.(*CreateOwnerRequest))
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
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoService_DeleteTodo_Handler,
		},
		{
			MethodName: "CreateOwner",
			Handler:    _TodoService_CreateOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todo/v1/todo.proto",
}

func init() { proto.RegisterFile("api/todo/v1/todo.proto", fileDescriptor_todo_95c0e673bb0c2a0b) }

var fileDescriptor_todo_95c0e673bb0c2a0b = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0xcf, 0x93, 0x40,
	0x10, 0x0d, 0x15, 0xea, 0xc7, 0x10, 0x9b, 0x74, 0x6d, 0x0c, 0xc1, 0x26, 0x22, 0xd6, 0x84, 0x13,
	0xa4, 0x35, 0x31, 0xd1, 0x5b, 0xd5, 0x26, 0xde, 0x8c, 0xd8, 0x93, 0x17, 0x43, 0xd9, 0x69, 0xb3,
	0x09, 0xb0, 0x08, 0xdb, 0xfa, 0x7f, 0x4d, 0xfc, 0x1f, 0x86, 0x5d, 0x28, 0x6b, 0x69, 0xe2, 0x09,
	0x66, 0xde, 0x9b, 0x99, 0xf7, 0x1e, 0xc0, 0xb3, 0xb4, 0x62, 0xb1, 0xe0, 0x94, 0xc7, 0x97, 0xb5,
	0x7c, 0x46, 0x55, 0xcd, 0x05, 0x27, 0x8f, 0xe5, 0xfb, 0x65, 0xed, 0xbd, 0x38, 0x71, 0x7e, 0xca,
	0x31, 0x96, 0xed, 0xc3, 0xf9, 0x18, 0x0b, 0x56, 0x60, 0x23, 0xd2, 0xa2, 0x52, 0xcc, 0xe0, 0xb7,
	0x01, 0xe6, 0x9e, 0x53, 0x4e, 0x66, 0x30, 0x61, 0xd4, 0x35, 0x7c, 0x23, 0x34, 0x93, 0x09, 0xa3,
	0x64, 0x01, 0x96, 0x60, 0x22, 0x47, 0x77, 0xe2, 0x1b, 0xa1, 0x9d, 0xa8, 0x82, 0xf8, 0xe0, 0x50,
	0x6c, 0xb2, 0x9a, 0x55, 0x82, 0xf1, 0xd2, 0x7d, 0x24, 0x31, 0xbd, 0x45, 0x96, 0x60, 0x67, 0xbc,
	0xa8, 0x72, 0x14, 0x48, 0x5d, 0xd3, 0x37, 0xc2, 0x87, 0x64, 0x68, 0x90, 0x77, 0x00, 0x59, 0x8d,
	0xa9, 0x40, 0xfa, 0x23, 0x15, 0xae, 0xe5, 0x1b, 0xa1, 0xb3, 0xf1, 0x22, 0x25, 0x32, 0xea, 0x45,
	0x46, 0xfb, 0x5e, 0x64, 0x62, 0x77, 0xec, 0xad, 0x68, 0x47, 0xcf, 0x15, 0xed, 0x47, 0xa7, 0xff,
	0x1f, 0xed, 0xd8, 0x5b, 0x11, 0x7c, 0x05, 0xeb, 0xcb, 0xaf, 0x12, 0xeb, 0x91, 0xc9, 0x25, 0xd8,
	0x47, 0x56, 0x37, 0xa2, 0x4c, 0x8b, 0xde, 0xe8, 0xd0, 0x20, 0x1e, 0x3c, 0xe4, 0x69, 0x07, 0x2a,
	0xa7, 0xd7, 0x3a, 0x78, 0x0b, 0xf3, 0x8f, 0x52, 0x5a, 0x1b, 0x5e, 0x82, 0x3f, 0xcf, 0xd8, 0x08,
	0xf2, 0x12, 0x4c, 0x26, 0xb0, 0x90, 0x07, 0x9c, 0xcd, 0x93, 0xa8, 0xfb, 0x0a, 0x91, 0xe4, 0x48,
	0x28, 0x58, 0x01, 0xd1, 0xe7, 0x9a, 0x8a, 0x97, 0x0d, 0xde, 0xea, 0x0a, 0x5e, 0xc1, 0xfc, 0x13,
	0xb6, 0x89, 0xe9, 0xdb, 0x6f, 0x49, 0x0b, 0x20, 0x3a, 0x49, 0xad, 0x0a, 0xde, 0xf7, 0x07, 0xa4,
	0xe3, 0x7e, 0x76, 0x05, 0x16, 0x6f, 0xeb, 0x4e, 0xda, 0xec, 0x2a, 0x4d, 0xb1, 0x14, 0x18, 0xbc,
	0x86, 0xa7, 0xff, 0xcc, 0xde, 0x57, 0xb7, 0xf9, 0x63, 0x80, 0xd3, 0xde, 0xfc, 0x86, 0xf5, 0x85,
	0x65, 0x48, 0x76, 0x00, 0x83, 0x27, 0xe2, 0x5d, 0x77, 0x8f, 0x02, 0xf2, 0x9e, 0xdf, 0xc5, 0xba,
	0x33, 0x3b, 0x80, 0xc1, 0x8f, 0xb6, 0x66, 0x94, 0x84, 0xb6, 0x66, 0x1c, 0x00, 0xf9, 0x0c, 0x8e,
	0x66, 0x82, 0xdc, 0x9e, 0xd4, 0x63, 0xf1, 0x96, 0xf7, 0x41, 0xb5, 0xe9, 0xc3, 0xf4, 0xbb, 0xd9,
	0xc2, 0x87, 0xa9, 0xfc, 0xbb, 0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xa3, 0x4a, 0x26,
	0x6e, 0x03, 0x00, 0x00,
}
