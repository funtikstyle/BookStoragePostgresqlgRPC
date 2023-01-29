// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/main.proto

package proto

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	GetBooksStorage(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*BooksReply, error)
	CreateBookStorage(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookReply, error)
	DeleteBookStorage(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookReply, error)
	TakeABookStorage(ctx context.Context, in *TakeABookRequest, opts ...grpc.CallOption) (*TakeABookReply, error)
	GetBookStorage(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookReply, error)
	ReturnABookStorage(ctx context.Context, in *ReturnABookRequest, opts ...grpc.CallOption) (*ReturnABookReply, error)
	UpdateBookStorage(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookReply, error)
	GetNotTakenBookByIds(ctx context.Context, in *GetNotTakenBookByIdsRequest, opts ...grpc.CallOption) (*GetNotTakenBookByIdsReply, error)
	CreateClientStorage(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientReply, error)
	GetClientsStorage(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsReply, error)
	DeleteClientStorage(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientReply, error)
	GetClientStorage(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*GetClientReply, error)
	UpdateClientStorage(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*UpdateClientReply, error)
	GetBooksByClientIdStorage(ctx context.Context, in *GetBooksByClientIdRequest, opts ...grpc.CallOption) (*GetBooksByClientIdReply, error)
	StatusClientByBooks(ctx context.Context, in *StatusClientByBooksRequest, opts ...grpc.CallOption) (*StatusClientByBooksReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetBooksStorage(ctx context.Context, in *GetBooksRequest, opts ...grpc.CallOption) (*BooksReply, error) {
	out := new(BooksReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/GetBooksStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) CreateBookStorage(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookReply, error) {
	out := new(CreateBookReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/CreateBookStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) DeleteBookStorage(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookReply, error) {
	out := new(DeleteBookReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/DeleteBookStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) TakeABookStorage(ctx context.Context, in *TakeABookRequest, opts ...grpc.CallOption) (*TakeABookReply, error) {
	out := new(TakeABookReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/TakeABookStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetBookStorage(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookReply, error) {
	out := new(GetBookReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/GetBookStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ReturnABookStorage(ctx context.Context, in *ReturnABookRequest, opts ...grpc.CallOption) (*ReturnABookReply, error) {
	out := new(ReturnABookReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/ReturnABookStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateBookStorage(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*UpdateBookReply, error) {
	out := new(UpdateBookReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/UpdateBookStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetNotTakenBookByIds(ctx context.Context, in *GetNotTakenBookByIdsRequest, opts ...grpc.CallOption) (*GetNotTakenBookByIdsReply, error) {
	out := new(GetNotTakenBookByIdsReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/GetNotTakenBookByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) CreateClientStorage(ctx context.Context, in *CreateClientRequest, opts ...grpc.CallOption) (*CreateClientReply, error) {
	out := new(CreateClientReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/CreateClientStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetClientsStorage(ctx context.Context, in *GetClientsRequest, opts ...grpc.CallOption) (*GetClientsReply, error) {
	out := new(GetClientsReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/GetClientsStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) DeleteClientStorage(ctx context.Context, in *DeleteClientRequest, opts ...grpc.CallOption) (*DeleteClientReply, error) {
	out := new(DeleteClientReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/DeleteClientStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetClientStorage(ctx context.Context, in *GetClientRequest, opts ...grpc.CallOption) (*GetClientReply, error) {
	out := new(GetClientReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/GetClientStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateClientStorage(ctx context.Context, in *UpdateClientRequest, opts ...grpc.CallOption) (*UpdateClientReply, error) {
	out := new(UpdateClientReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/UpdateClientStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetBooksByClientIdStorage(ctx context.Context, in *GetBooksByClientIdRequest, opts ...grpc.CallOption) (*GetBooksByClientIdReply, error) {
	out := new(GetBooksByClientIdReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/GetBooksByClientIdStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) StatusClientByBooks(ctx context.Context, in *StatusClientByBooksRequest, opts ...grpc.CallOption) (*StatusClientByBooksReply, error) {
	out := new(StatusClientByBooksReply)
	err := c.cc.Invoke(ctx, "/main.Greeter/StatusClientByBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	GetBooksStorage(context.Context, *GetBooksRequest) (*BooksReply, error)
	CreateBookStorage(context.Context, *CreateBookRequest) (*CreateBookReply, error)
	DeleteBookStorage(context.Context, *DeleteBookRequest) (*DeleteBookReply, error)
	TakeABookStorage(context.Context, *TakeABookRequest) (*TakeABookReply, error)
	GetBookStorage(context.Context, *GetBookRequest) (*GetBookReply, error)
	ReturnABookStorage(context.Context, *ReturnABookRequest) (*ReturnABookReply, error)
	UpdateBookStorage(context.Context, *UpdateBookRequest) (*UpdateBookReply, error)
	GetNotTakenBookByIds(context.Context, *GetNotTakenBookByIdsRequest) (*GetNotTakenBookByIdsReply, error)
	CreateClientStorage(context.Context, *CreateClientRequest) (*CreateClientReply, error)
	GetClientsStorage(context.Context, *GetClientsRequest) (*GetClientsReply, error)
	DeleteClientStorage(context.Context, *DeleteClientRequest) (*DeleteClientReply, error)
	GetClientStorage(context.Context, *GetClientRequest) (*GetClientReply, error)
	UpdateClientStorage(context.Context, *UpdateClientRequest) (*UpdateClientReply, error)
	GetBooksByClientIdStorage(context.Context, *GetBooksByClientIdRequest) (*GetBooksByClientIdReply, error)
	StatusClientByBooks(context.Context, *StatusClientByBooksRequest) (*StatusClientByBooksReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) GetBooksStorage(context.Context, *GetBooksRequest) (*BooksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksStorage not implemented")
}
func (UnimplementedGreeterServer) CreateBookStorage(context.Context, *CreateBookRequest) (*CreateBookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBookStorage not implemented")
}
func (UnimplementedGreeterServer) DeleteBookStorage(context.Context, *DeleteBookRequest) (*DeleteBookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookStorage not implemented")
}
func (UnimplementedGreeterServer) TakeABookStorage(context.Context, *TakeABookRequest) (*TakeABookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeABookStorage not implemented")
}
func (UnimplementedGreeterServer) GetBookStorage(context.Context, *GetBookRequest) (*GetBookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookStorage not implemented")
}
func (UnimplementedGreeterServer) ReturnABookStorage(context.Context, *ReturnABookRequest) (*ReturnABookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnABookStorage not implemented")
}
func (UnimplementedGreeterServer) UpdateBookStorage(context.Context, *UpdateBookRequest) (*UpdateBookReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBookStorage not implemented")
}
func (UnimplementedGreeterServer) GetNotTakenBookByIds(context.Context, *GetNotTakenBookByIdsRequest) (*GetNotTakenBookByIdsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotTakenBookByIds not implemented")
}
func (UnimplementedGreeterServer) CreateClientStorage(context.Context, *CreateClientRequest) (*CreateClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientStorage not implemented")
}
func (UnimplementedGreeterServer) GetClientsStorage(context.Context, *GetClientsRequest) (*GetClientsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientsStorage not implemented")
}
func (UnimplementedGreeterServer) DeleteClientStorage(context.Context, *DeleteClientRequest) (*DeleteClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientStorage not implemented")
}
func (UnimplementedGreeterServer) GetClientStorage(context.Context, *GetClientRequest) (*GetClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientStorage not implemented")
}
func (UnimplementedGreeterServer) UpdateClientStorage(context.Context, *UpdateClientRequest) (*UpdateClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientStorage not implemented")
}
func (UnimplementedGreeterServer) GetBooksByClientIdStorage(context.Context, *GetBooksByClientIdRequest) (*GetBooksByClientIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksByClientIdStorage not implemented")
}
func (UnimplementedGreeterServer) StatusClientByBooks(context.Context, *StatusClientByBooksRequest) (*StatusClientByBooksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusClientByBooks not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_GetBooksStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBooksStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/GetBooksStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBooksStorage(ctx, req.(*GetBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_CreateBookStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateBookStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/CreateBookStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateBookStorage(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_DeleteBookStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).DeleteBookStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/DeleteBookStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).DeleteBookStorage(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_TakeABookStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeABookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).TakeABookStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/TakeABookStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).TakeABookStorage(ctx, req.(*TakeABookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetBookStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBookStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/GetBookStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBookStorage(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ReturnABookStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReturnABookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ReturnABookStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/ReturnABookStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ReturnABookStorage(ctx, req.(*ReturnABookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateBookStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateBookStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/UpdateBookStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateBookStorage(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetNotTakenBookByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotTakenBookByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetNotTakenBookByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/GetNotTakenBookByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetNotTakenBookByIds(ctx, req.(*GetNotTakenBookByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_CreateClientStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateClientStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/CreateClientStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateClientStorage(ctx, req.(*CreateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetClientsStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetClientsStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/GetClientsStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetClientsStorage(ctx, req.(*GetClientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_DeleteClientStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).DeleteClientStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/DeleteClientStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).DeleteClientStorage(ctx, req.(*DeleteClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetClientStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetClientStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/GetClientStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetClientStorage(ctx, req.(*GetClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateClientStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateClientStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/UpdateClientStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateClientStorage(ctx, req.(*UpdateClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetBooksByClientIdStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBooksByClientIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBooksByClientIdStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/GetBooksByClientIdStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBooksByClientIdStorage(ctx, req.(*GetBooksByClientIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_StatusClientByBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusClientByBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).StatusClientByBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.Greeter/StatusClientByBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).StatusClientByBooks(ctx, req.(*StatusClientByBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBooksStorage",
			Handler:    _Greeter_GetBooksStorage_Handler,
		},
		{
			MethodName: "CreateBookStorage",
			Handler:    _Greeter_CreateBookStorage_Handler,
		},
		{
			MethodName: "DeleteBookStorage",
			Handler:    _Greeter_DeleteBookStorage_Handler,
		},
		{
			MethodName: "TakeABookStorage",
			Handler:    _Greeter_TakeABookStorage_Handler,
		},
		{
			MethodName: "GetBookStorage",
			Handler:    _Greeter_GetBookStorage_Handler,
		},
		{
			MethodName: "ReturnABookStorage",
			Handler:    _Greeter_ReturnABookStorage_Handler,
		},
		{
			MethodName: "UpdateBookStorage",
			Handler:    _Greeter_UpdateBookStorage_Handler,
		},
		{
			MethodName: "GetNotTakenBookByIds",
			Handler:    _Greeter_GetNotTakenBookByIds_Handler,
		},
		{
			MethodName: "CreateClientStorage",
			Handler:    _Greeter_CreateClientStorage_Handler,
		},
		{
			MethodName: "GetClientsStorage",
			Handler:    _Greeter_GetClientsStorage_Handler,
		},
		{
			MethodName: "DeleteClientStorage",
			Handler:    _Greeter_DeleteClientStorage_Handler,
		},
		{
			MethodName: "GetClientStorage",
			Handler:    _Greeter_GetClientStorage_Handler,
		},
		{
			MethodName: "UpdateClientStorage",
			Handler:    _Greeter_UpdateClientStorage_Handler,
		},
		{
			MethodName: "GetBooksByClientIdStorage",
			Handler:    _Greeter_GetBooksByClientIdStorage_Handler,
		},
		{
			MethodName: "StatusClientByBooks",
			Handler:    _Greeter_StatusClientByBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/main.proto",
}
