// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blogpb/blog.proto

package blogpb

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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerType         string   `protobuf:"bytes,2,opt,name=customerType,proto3" json:"customerType,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	MiddleInitial        string   `protobuf:"bytes,5,opt,name=middleInitial,proto3" json:"middleInitial,omitempty"`
	OrganizationName     string   `protobuf:"bytes,6,opt,name=organizationName,proto3" json:"organizationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetCustomerType() string {
	if m != nil {
		return m.CustomerType
	}
	return ""
}

func (m *Blog) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Blog) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Blog) GetMiddleInitial() string {
	if m != nil {
		return m.MiddleInitial
	}
	return ""
}

func (m *Blog) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRequest) Reset()         { *m = ReadBlogRequest{} }
func (m *ReadBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRequest) ProtoMessage()    {}
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{3}
}

func (m *ReadBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRequest.Unmarshal(m, b)
}
func (m *ReadBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRequest.Marshal(b, m, deterministic)
}
func (m *ReadBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRequest.Merge(m, src)
}
func (m *ReadBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRequest.Size(m)
}
func (m *ReadBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRequest proto.InternalMessageInfo

func (m *ReadBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ReadBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogResponse) Reset()         { *m = ReadBlogResponse{} }
func (m *ReadBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ReadBlogResponse) ProtoMessage()    {}
func (*ReadBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{4}
}

func (m *ReadBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogResponse.Unmarshal(m, b)
}
func (m *ReadBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogResponse.Marshal(b, m, deterministic)
}
func (m *ReadBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogResponse.Merge(m, src)
}
func (m *ReadBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ReadBlogResponse.Size(m)
}
func (m *ReadBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogResponse proto.InternalMessageInfo

func (m *ReadBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ListBlogRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogRequest) Reset()         { *m = ListBlogRequest{} }
func (m *ListBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ListBlogRequest) ProtoMessage()    {}
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{5}
}

func (m *ListBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogRequest.Unmarshal(m, b)
}
func (m *ListBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogRequest.Marshal(b, m, deterministic)
}
func (m *ListBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogRequest.Merge(m, src)
}
func (m *ListBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ListBlogRequest.Size(m)
}
func (m *ListBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogRequest proto.InternalMessageInfo

type ListBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogResponse) Reset()         { *m = ListBlogResponse{} }
func (m *ListBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ListBlogResponse) ProtoMessage()    {}
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{6}
}

func (m *ListBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogResponse.Unmarshal(m, b)
}
func (m *ListBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogResponse.Marshal(b, m, deterministic)
}
func (m *ListBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogResponse.Merge(m, src)
}
func (m *ListBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ListBlogResponse.Size(m)
}
func (m *ListBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogResponse proto.InternalMessageInfo

func (m *ListBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.CreateBlogResponse")
	proto.RegisterType((*ReadBlogRequest)(nil), "blog.ReadBlogRequest")
	proto.RegisterType((*ReadBlogResponse)(nil), "blog.ReadBlogResponse")
	proto.RegisterType((*ListBlogRequest)(nil), "blog.ListBlogRequest")
	proto.RegisterType((*ListBlogResponse)(nil), "blog.ListBlogResponse")
}

func init() { proto.RegisterFile("blog/blogpb/blog.proto", fileDescriptor_a4b0406114889fe6) }

var fileDescriptor_a4b0406114889fe6 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x4e, 0x11, 0x6b, 0x19, 0x54, 0x60, 0x12, 0xa1, 0x21, 0xc6, 0x98, 0xc6, 0x83, 0xe1, 0x80,
	0x06, 0xbc, 0x18, 0x0f, 0x26, 0x78, 0x22, 0x31, 0x1e, 0xd0, 0x93, 0x17, 0xb3, 0xd0, 0x91, 0x6c,
	0x52, 0xba, 0xb5, 0x5d, 0x4c, 0xf4, 0x01, 0x79, 0x2e, 0xb3, 0x53, 0x4a, 0xa1, 0x3d, 0xc8, 0xa5,
	0xed, 0x7c, 0x7f, 0xdb, 0x6f, 0xb2, 0xd0, 0x9e, 0x06, 0x6a, 0x7e, 0x63, 0x1e, 0xd1, 0x94, 0x5f,
	0xfd, 0x28, 0x56, 0x5a, 0x61, 0xd5, 0x7c, 0x7b, 0x2b, 0x0b, 0xaa, 0xa3, 0x40, 0xcd, 0xf1, 0x14,
	0x2a, 0xd2, 0x77, 0xad, 0x4b, 0xeb, 0xba, 0x36, 0xa9, 0x48, 0x1f, 0x3d, 0x38, 0x9e, 0x2d, 0x13,
	0xad, 0x16, 0x14, 0xbf, 0xfd, 0x44, 0xe4, 0x56, 0x98, 0xd9, 0xc1, 0xf0, 0x1c, 0x6a, 0x9f, 0x32,
	0x4e, 0xf4, 0x8b, 0x58, 0x90, 0x7b, 0xc0, 0x82, 0x1c, 0xc0, 0x2e, 0x38, 0x81, 0x58, 0x93, 0x55,
	0x26, 0x37, 0x33, 0x5e, 0xc1, 0xc9, 0x42, 0xfa, 0x7e, 0x40, 0xe3, 0x50, 0x6a, 0x29, 0x02, 0xf7,
	0x90, 0x05, 0xbb, 0x20, 0xf6, 0xa0, 0xa9, 0xe2, 0xb9, 0x08, 0xe5, 0xaf, 0xd0, 0x52, 0x85, 0x9c,
	0x64, 0xb3, 0xb0, 0x84, 0x7b, 0x43, 0x68, 0x3d, 0xc5, 0x24, 0x34, 0x99, 0x36, 0x13, 0xfa, 0x5a,
	0x52, 0xa2, 0xf1, 0x02, 0xb8, 0x25, 0xd7, 0xaa, 0x0f, 0xa0, 0xcf, 0xf5, 0x59, 0x90, 0xb6, 0xbf,
	0x03, 0xdc, 0x36, 0x25, 0x91, 0x0a, 0x13, 0xfa, 0xd7, 0xd5, 0x83, 0xc6, 0x84, 0x84, 0xbf, 0x7d,
	0x50, 0x07, 0x8e, 0x0c, 0xf5, 0xb1, 0x59, 0xa1, 0x6d, 0xc6, 0xb1, 0xef, 0x0d, 0xa0, 0x99, 0x6b,
	0xf7, 0xcc, 0x6f, 0x41, 0xe3, 0x59, 0x26, 0x7a, 0x2b, 0xdf, 0xc4, 0xe4, 0xd0, 0x7e, 0x31, 0x83,
	0x95, 0x05, 0x75, 0x33, 0xbe, 0x52, 0xfc, 0x2d, 0x67, 0x84, 0x8f, 0x00, 0x79, 0x59, 0xec, 0xa4,
	0xfa, 0xd2, 0xce, 0xba, 0x6e, 0x99, 0x58, 0x1f, 0x78, 0x0f, 0x4e, 0xd6, 0x05, 0xcf, 0x52, 0x55,
	0x61, 0x0f, 0xdd, 0x76, 0x11, 0x5e, 0x5b, 0x1f, 0xc0, 0xc9, 0xfe, 0x3f, 0xb3, 0x16, 0x2a, 0x66,
	0xd6, 0x62, 0xcd, 0x5b, 0x6b, 0xe4, 0xbc, 0xdb, 0xe9, 0xf5, 0x9d, 0xda, 0x7c, 0x75, 0x87, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x05, 0x38, 0xdb, 0x42, 0xd4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error)
	ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListBlogClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error) {
	out := new(ReadBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/ListBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogClient interface {
	Recv() (*ListBlogResponse, error)
	grpc.ClientStream
}

type blogServiceListBlogClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogClient) Recv() (*ListBlogResponse, error) {
	m := new(ListBlogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
	ListBlog(*ListBlogRequest, BlogService_ListBlogServer) error
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*CreateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(ctx context.Context, req *ReadBlogRequest) (*ReadBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ListBlog(req *ListBlogRequest, srv BlogService_ListBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlog(m, &blogServiceListBlogServer{stream})
}

type BlogService_ListBlogServer interface {
	Send(*ListBlogResponse) error
	grpc.ServerStream
}

type blogServiceListBlogServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogServer) Send(m *ListBlogResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlog",
			Handler:       _BlogService_ListBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog/blogpb/blog.proto",
}
