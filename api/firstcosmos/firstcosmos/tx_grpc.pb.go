// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: firstcosmos/firstcosmos/tx.proto

package firstcosmos

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
	Msg_UpdateParams_FullMethodName = "/firstcosmos.firstcosmos.Msg/UpdateParams"
	Msg_Hello_FullMethodName        = "/firstcosmos.firstcosmos.Msg/Hello"
	Msg_CreatePeople_FullMethodName = "/firstcosmos.firstcosmos.Msg/CreatePeople"
	Msg_UpdatePeople_FullMethodName = "/firstcosmos.firstcosmos.Msg/UpdatePeople"
	Msg_DeletePeople_FullMethodName = "/firstcosmos.firstcosmos.Msg/DeletePeople"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	Hello(ctx context.Context, in *MsgHello, opts ...grpc.CallOption) (*MsgHelloResponse, error)
	CreatePeople(ctx context.Context, in *MsgCreatePeople, opts ...grpc.CallOption) (*MsgCreatePeopleResponse, error)
	UpdatePeople(ctx context.Context, in *MsgUpdatePeople, opts ...grpc.CallOption) (*MsgUpdatePeopleResponse, error)
	DeletePeople(ctx context.Context, in *MsgDeletePeople, opts ...grpc.CallOption) (*MsgDeletePeopleResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Hello(ctx context.Context, in *MsgHello, opts ...grpc.CallOption) (*MsgHelloResponse, error) {
	out := new(MsgHelloResponse)
	err := c.cc.Invoke(ctx, Msg_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePeople(ctx context.Context, in *MsgCreatePeople, opts ...grpc.CallOption) (*MsgCreatePeopleResponse, error) {
	out := new(MsgCreatePeopleResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePeople_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePeople(ctx context.Context, in *MsgUpdatePeople, opts ...grpc.CallOption) (*MsgUpdatePeopleResponse, error) {
	out := new(MsgUpdatePeopleResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePeople_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePeople(ctx context.Context, in *MsgDeletePeople, opts ...grpc.CallOption) (*MsgDeletePeopleResponse, error) {
	out := new(MsgDeletePeopleResponse)
	err := c.cc.Invoke(ctx, Msg_DeletePeople_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	Hello(context.Context, *MsgHello) (*MsgHelloResponse, error)
	CreatePeople(context.Context, *MsgCreatePeople) (*MsgCreatePeopleResponse, error)
	UpdatePeople(context.Context, *MsgUpdatePeople) (*MsgUpdatePeopleResponse, error)
	DeletePeople(context.Context, *MsgDeletePeople) (*MsgDeletePeopleResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) Hello(context.Context, *MsgHello) (*MsgHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedMsgServer) CreatePeople(context.Context, *MsgCreatePeople) (*MsgCreatePeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeople not implemented")
}
func (UnimplementedMsgServer) UpdatePeople(context.Context, *MsgUpdatePeople) (*MsgUpdatePeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeople not implemented")
}
func (UnimplementedMsgServer) DeletePeople(context.Context, *MsgDeletePeople) (*MsgDeletePeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeople not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgHello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Hello(ctx, req.(*MsgHello))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePeople)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePeople_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePeople(ctx, req.(*MsgCreatePeople))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePeople)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePeople_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePeople(ctx, req.(*MsgUpdatePeople))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePeople)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeletePeople_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePeople(ctx, req.(*MsgDeletePeople))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "firstcosmos.firstcosmos.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _Msg_Hello_Handler,
		},
		{
			MethodName: "CreatePeople",
			Handler:    _Msg_CreatePeople_Handler,
		},
		{
			MethodName: "UpdatePeople",
			Handler:    _Msg_UpdatePeople_Handler,
		},
		{
			MethodName: "DeletePeople",
			Handler:    _Msg_DeletePeople_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "firstcosmos/firstcosmos/tx.proto",
}
