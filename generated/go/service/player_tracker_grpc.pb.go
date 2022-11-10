// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: service/player_tracker.proto

package service

import (
	context "context"
	model "github.com/towerdefence-cc/grpc-api-specs/generated/go/model"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlayerTrackerClient is the client API for PlayerTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerTrackerClient interface {
	ProxyPlayerLogin(ctx context.Context, in *PlayerLoginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ServerPlayerLogin(ctx context.Context, in *PlayerLoginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ProxyPlayerDisconnect(ctx context.Context, in *PlayerDisconnectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPlayerServer(ctx context.Context, in *model.PlayerRequest, opts ...grpc.CallOption) (*GetPlayerServerResponse, error)
	GetPlayerServers(ctx context.Context, in *model.PlayersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error)
	GetServerPlayerCount(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayerCountResponse, error)
	GetServerPlayers(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error)
}

type playerTrackerClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerTrackerClient(cc grpc.ClientConnInterface) PlayerTrackerClient {
	return &playerTrackerClient{cc}
}

func (c *playerTrackerClient) ProxyPlayerLogin(ctx context.Context, in *PlayerLoginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/ProxyPlayerLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) ServerPlayerLogin(ctx context.Context, in *PlayerLoginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/ServerPlayerLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) ProxyPlayerDisconnect(ctx context.Context, in *PlayerDisconnectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/ProxyPlayerDisconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetPlayerServer(ctx context.Context, in *model.PlayerRequest, opts ...grpc.CallOption) (*GetPlayerServerResponse, error) {
	out := new(GetPlayerServerResponse)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/GetPlayerServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetPlayerServers(ctx context.Context, in *model.PlayersRequest, opts ...grpc.CallOption) (*GetPlayerServersResponse, error) {
	out := new(GetPlayerServersResponse)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/GetPlayerServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerPlayerCount(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayerCountResponse, error) {
	out := new(GetServerPlayerCountResponse)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/GetServerPlayerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerTrackerClient) GetServerPlayers(ctx context.Context, in *ServerIdRequest, opts ...grpc.CallOption) (*GetServerPlayersResponse, error) {
	out := new(GetServerPlayersResponse)
	err := c.cc.Invoke(ctx, "/service.player_tracker.PlayerTracker/GetServerPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerTrackerServer is the server API for PlayerTracker service.
// All implementations must embed UnimplementedPlayerTrackerServer
// for forward compatibility
type PlayerTrackerServer interface {
	ProxyPlayerLogin(context.Context, *PlayerLoginRequest) (*emptypb.Empty, error)
	ServerPlayerLogin(context.Context, *PlayerLoginRequest) (*emptypb.Empty, error)
	ProxyPlayerDisconnect(context.Context, *PlayerDisconnectRequest) (*emptypb.Empty, error)
	GetPlayerServer(context.Context, *model.PlayerRequest) (*GetPlayerServerResponse, error)
	GetPlayerServers(context.Context, *model.PlayersRequest) (*GetPlayerServersResponse, error)
	GetServerPlayerCount(context.Context, *ServerIdRequest) (*GetServerPlayerCountResponse, error)
	GetServerPlayers(context.Context, *ServerIdRequest) (*GetServerPlayersResponse, error)
	mustEmbedUnimplementedPlayerTrackerServer()
}

// UnimplementedPlayerTrackerServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerTrackerServer struct {
}

func (UnimplementedPlayerTrackerServer) ProxyPlayerLogin(context.Context, *PlayerLoginRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyPlayerLogin not implemented")
}
func (UnimplementedPlayerTrackerServer) ServerPlayerLogin(context.Context, *PlayerLoginRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerPlayerLogin not implemented")
}
func (UnimplementedPlayerTrackerServer) ProxyPlayerDisconnect(context.Context, *PlayerDisconnectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProxyPlayerDisconnect not implemented")
}
func (UnimplementedPlayerTrackerServer) GetPlayerServer(context.Context, *model.PlayerRequest) (*GetPlayerServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServer not implemented")
}
func (UnimplementedPlayerTrackerServer) GetPlayerServers(context.Context, *model.PlayersRequest) (*GetPlayerServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerServers not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayerCount(context.Context, *ServerIdRequest) (*GetServerPlayerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayerCount not implemented")
}
func (UnimplementedPlayerTrackerServer) GetServerPlayers(context.Context, *ServerIdRequest) (*GetServerPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerPlayers not implemented")
}
func (UnimplementedPlayerTrackerServer) mustEmbedUnimplementedPlayerTrackerServer() {}

// UnsafePlayerTrackerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerTrackerServer will
// result in compilation errors.
type UnsafePlayerTrackerServer interface {
	mustEmbedUnimplementedPlayerTrackerServer()
}

func RegisterPlayerTrackerServer(s grpc.ServiceRegistrar, srv PlayerTrackerServer) {
	s.RegisterService(&PlayerTracker_ServiceDesc, srv)
}

func _PlayerTracker_ProxyPlayerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).ProxyPlayerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/ProxyPlayerLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).ProxyPlayerLogin(ctx, req.(*PlayerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_ServerPlayerLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).ServerPlayerLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/ServerPlayerLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).ServerPlayerLogin(ctx, req.(*PlayerLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_ProxyPlayerDisconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerDisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).ProxyPlayerDisconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/ProxyPlayerDisconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).ProxyPlayerDisconnect(ctx, req.(*PlayerDisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetPlayerServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/GetPlayerServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerServer(ctx, req.(*model.PlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetPlayerServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.PlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetPlayerServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/GetPlayerServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetPlayerServers(ctx, req.(*model.PlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerPlayerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerPlayerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/GetServerPlayerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerPlayerCount(ctx, req.(*ServerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerTracker_GetServerPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTrackerServer).GetServerPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.player_tracker.PlayerTracker/GetServerPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTrackerServer).GetServerPlayers(ctx, req.(*ServerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerTracker_ServiceDesc is the grpc.ServiceDesc for PlayerTracker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerTracker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.player_tracker.PlayerTracker",
	HandlerType: (*PlayerTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProxyPlayerLogin",
			Handler:    _PlayerTracker_ProxyPlayerLogin_Handler,
		},
		{
			MethodName: "ServerPlayerLogin",
			Handler:    _PlayerTracker_ServerPlayerLogin_Handler,
		},
		{
			MethodName: "ProxyPlayerDisconnect",
			Handler:    _PlayerTracker_ProxyPlayerDisconnect_Handler,
		},
		{
			MethodName: "GetPlayerServer",
			Handler:    _PlayerTracker_GetPlayerServer_Handler,
		},
		{
			MethodName: "GetPlayerServers",
			Handler:    _PlayerTracker_GetPlayerServers_Handler,
		},
		{
			MethodName: "GetServerPlayerCount",
			Handler:    _PlayerTracker_GetServerPlayerCount_Handler,
		},
		{
			MethodName: "GetServerPlayers",
			Handler:    _PlayerTracker_GetServerPlayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/player_tracker.proto",
}