// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: pkg/pb/manga.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MangaService_CreateManga_FullMethodName = "/manga.MangaService/CreateManga"
	MangaService_GetManga_FullMethodName    = "/manga.MangaService/GetManga"
	MangaService_ListManga_FullMethodName   = "/manga.MangaService/ListManga"
	MangaService_UpdateManga_FullMethodName = "/manga.MangaService/UpdateManga"
	MangaService_DeleteManga_FullMethodName = "/manga.MangaService/DeleteManga"
)

// MangaServiceClient is the client API for MangaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MangaServiceClient interface {
	CreateManga(ctx context.Context, in *CreateMangaRequest, opts ...grpc.CallOption) (*CreateMangaResponse, error)
	GetManga(ctx context.Context, in *GetMangaRequest, opts ...grpc.CallOption) (*GetMangaResponse, error)
	ListManga(ctx context.Context, in *ListMangaRequest, opts ...grpc.CallOption) (*ListMangaResponse, error)
	UpdateManga(ctx context.Context, in *UpdateMangaRequest, opts ...grpc.CallOption) (*UpdateMangaResponse, error)
	DeleteManga(ctx context.Context, in *DeleteMangaRequest, opts ...grpc.CallOption) (*DeleteMangaResponse, error)
}

type mangaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMangaServiceClient(cc grpc.ClientConnInterface) MangaServiceClient {
	return &mangaServiceClient{cc}
}

func (c *mangaServiceClient) CreateManga(ctx context.Context, in *CreateMangaRequest, opts ...grpc.CallOption) (*CreateMangaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMangaResponse)
	err := c.cc.Invoke(ctx, MangaService_CreateManga_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mangaServiceClient) GetManga(ctx context.Context, in *GetMangaRequest, opts ...grpc.CallOption) (*GetMangaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMangaResponse)
	err := c.cc.Invoke(ctx, MangaService_GetManga_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mangaServiceClient) ListManga(ctx context.Context, in *ListMangaRequest, opts ...grpc.CallOption) (*ListMangaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMangaResponse)
	err := c.cc.Invoke(ctx, MangaService_ListManga_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mangaServiceClient) UpdateManga(ctx context.Context, in *UpdateMangaRequest, opts ...grpc.CallOption) (*UpdateMangaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMangaResponse)
	err := c.cc.Invoke(ctx, MangaService_UpdateManga_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mangaServiceClient) DeleteManga(ctx context.Context, in *DeleteMangaRequest, opts ...grpc.CallOption) (*DeleteMangaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMangaResponse)
	err := c.cc.Invoke(ctx, MangaService_DeleteManga_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MangaServiceServer is the server API for MangaService service.
// All implementations must embed UnimplementedMangaServiceServer
// for forward compatibility.
type MangaServiceServer interface {
	CreateManga(context.Context, *CreateMangaRequest) (*CreateMangaResponse, error)
	GetManga(context.Context, *GetMangaRequest) (*GetMangaResponse, error)
	ListManga(context.Context, *ListMangaRequest) (*ListMangaResponse, error)
	UpdateManga(context.Context, *UpdateMangaRequest) (*UpdateMangaResponse, error)
	DeleteManga(context.Context, *DeleteMangaRequest) (*DeleteMangaResponse, error)
	mustEmbedUnimplementedMangaServiceServer()
}

// UnimplementedMangaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMangaServiceServer struct{}

func (UnimplementedMangaServiceServer) CreateManga(context.Context, *CreateMangaRequest) (*CreateMangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateManga not implemented")
}
func (UnimplementedMangaServiceServer) GetManga(context.Context, *GetMangaRequest) (*GetMangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManga not implemented")
}
func (UnimplementedMangaServiceServer) ListManga(context.Context, *ListMangaRequest) (*ListMangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManga not implemented")
}
func (UnimplementedMangaServiceServer) UpdateManga(context.Context, *UpdateMangaRequest) (*UpdateMangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateManga not implemented")
}
func (UnimplementedMangaServiceServer) DeleteManga(context.Context, *DeleteMangaRequest) (*DeleteMangaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManga not implemented")
}
func (UnimplementedMangaServiceServer) mustEmbedUnimplementedMangaServiceServer() {}
func (UnimplementedMangaServiceServer) testEmbeddedByValue()                      {}

// UnsafeMangaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MangaServiceServer will
// result in compilation errors.
type UnsafeMangaServiceServer interface {
	mustEmbedUnimplementedMangaServiceServer()
}

func RegisterMangaServiceServer(s grpc.ServiceRegistrar, srv MangaServiceServer) {
	// If the following call pancis, it indicates UnimplementedMangaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MangaService_ServiceDesc, srv)
}

func _MangaService_CreateManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MangaServiceServer).CreateManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MangaService_CreateManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MangaServiceServer).CreateManga(ctx, req.(*CreateMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MangaService_GetManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MangaServiceServer).GetManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MangaService_GetManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MangaServiceServer).GetManga(ctx, req.(*GetMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MangaService_ListManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MangaServiceServer).ListManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MangaService_ListManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MangaServiceServer).ListManga(ctx, req.(*ListMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MangaService_UpdateManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MangaServiceServer).UpdateManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MangaService_UpdateManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MangaServiceServer).UpdateManga(ctx, req.(*UpdateMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MangaService_DeleteManga_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMangaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MangaServiceServer).DeleteManga(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MangaService_DeleteManga_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MangaServiceServer).DeleteManga(ctx, req.(*DeleteMangaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MangaService_ServiceDesc is the grpc.ServiceDesc for MangaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MangaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manga.MangaService",
	HandlerType: (*MangaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateManga",
			Handler:    _MangaService_CreateManga_Handler,
		},
		{
			MethodName: "GetManga",
			Handler:    _MangaService_GetManga_Handler,
		},
		{
			MethodName: "ListManga",
			Handler:    _MangaService_ListManga_Handler,
		},
		{
			MethodName: "UpdateManga",
			Handler:    _MangaService_UpdateManga_Handler,
		},
		{
			MethodName: "DeleteManga",
			Handler:    _MangaService_DeleteManga_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/manga.proto",
}