// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: proto/movie.proto

package movie

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MovieCrud_Insert_FullMethodName  = "/api.v1.MovieCrud/Insert"
	MovieCrud_Update_FullMethodName  = "/api.v1.MovieCrud/Update"
	MovieCrud_List_FullMethodName    = "/api.v1.MovieCrud/List"
	MovieCrud_Delete_FullMethodName  = "/api.v1.MovieCrud/Delete"
	MovieCrud_GetById_FullMethodName = "/api.v1.MovieCrud/GetById"
)

// MovieCrudClient is the client API for MovieCrud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieCrudClient interface {
	Insert(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*ResponseMovie, error)
	Update(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*ResponseMovie, error)
	List(ctx context.Context, in *ListRequestMovie, opts ...grpc.CallOption) (*ResponseMovie, error)
	Delete(ctx context.Context, in *RequestByIdMovie, opts ...grpc.CallOption) (*ResponseMovie, error)
	GetById(ctx context.Context, in *RequestByIdMovie, opts ...grpc.CallOption) (*ResponseMovie, error)
}

type movieCrudClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieCrudClient(cc grpc.ClientConnInterface) MovieCrudClient {
	return &movieCrudClient{cc}
}

func (c *movieCrudClient) Insert(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*ResponseMovie, error) {
	out := new(ResponseMovie)
	err := c.cc.Invoke(ctx, MovieCrud_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCrudClient) Update(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*ResponseMovie, error) {
	out := new(ResponseMovie)
	err := c.cc.Invoke(ctx, MovieCrud_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCrudClient) List(ctx context.Context, in *ListRequestMovie, opts ...grpc.CallOption) (*ResponseMovie, error) {
	out := new(ResponseMovie)
	err := c.cc.Invoke(ctx, MovieCrud_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCrudClient) Delete(ctx context.Context, in *RequestByIdMovie, opts ...grpc.CallOption) (*ResponseMovie, error) {
	out := new(ResponseMovie)
	err := c.cc.Invoke(ctx, MovieCrud_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCrudClient) GetById(ctx context.Context, in *RequestByIdMovie, opts ...grpc.CallOption) (*ResponseMovie, error) {
	out := new(ResponseMovie)
	err := c.cc.Invoke(ctx, MovieCrud_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieCrudServer is the server API for MovieCrud service.
// All implementations must embed UnimplementedMovieCrudServer
// for forward compatibility
type MovieCrudServer interface {
	Insert(context.Context, *Movie) (*ResponseMovie, error)
	Update(context.Context, *Movie) (*ResponseMovie, error)
	List(context.Context, *ListRequestMovie) (*ResponseMovie, error)
	Delete(context.Context, *RequestByIdMovie) (*ResponseMovie, error)
	GetById(context.Context, *RequestByIdMovie) (*ResponseMovie, error)
	mustEmbedUnimplementedMovieCrudServer()
}

// UnimplementedMovieCrudServer must be embedded to have forward compatible implementations.
type UnimplementedMovieCrudServer struct {
}

func (UnimplementedMovieCrudServer) Insert(context.Context, *Movie) (*ResponseMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedMovieCrudServer) Update(context.Context, *Movie) (*ResponseMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMovieCrudServer) List(context.Context, *ListRequestMovie) (*ResponseMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMovieCrudServer) Delete(context.Context, *RequestByIdMovie) (*ResponseMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMovieCrudServer) GetById(context.Context, *RequestByIdMovie) (*ResponseMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedMovieCrudServer) mustEmbedUnimplementedMovieCrudServer() {}

// UnsafeMovieCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieCrudServer will
// result in compilation errors.
type UnsafeMovieCrudServer interface {
	mustEmbedUnimplementedMovieCrudServer()
}

func RegisterMovieCrudServer(s grpc.ServiceRegistrar, srv MovieCrudServer) {
	s.RegisterService(&MovieCrud_ServiceDesc, srv)
}

func _MovieCrud_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCrudServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieCrud_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCrudServer).Insert(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCrud_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCrudServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieCrud_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCrudServer).Update(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCrud_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequestMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCrudServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieCrud_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCrudServer).List(ctx, req.(*ListRequestMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCrud_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIdMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCrudServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieCrud_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCrudServer).Delete(ctx, req.(*RequestByIdMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCrud_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestByIdMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCrudServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MovieCrud_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCrudServer).GetById(ctx, req.(*RequestByIdMovie))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieCrud_ServiceDesc is the grpc.ServiceDesc for MovieCrud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieCrud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.MovieCrud",
	HandlerType: (*MovieCrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _MovieCrud_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _MovieCrud_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _MovieCrud_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MovieCrud_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _MovieCrud_GetById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/movie.proto",
}
