// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: game_player_data/grpc.proto

package gameplayerdata

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
	GamePlayerDataService_GetGamePlayerData_FullMethodName         = "/emortal.grpc.gameplayerdata.GamePlayerDataService/GetGamePlayerData"
	GamePlayerDataService_GetMultipleGamePlayerData_FullMethodName = "/emortal.grpc.gameplayerdata.GamePlayerDataService/GetMultipleGamePlayerData"
)

// GamePlayerDataServiceClient is the client API for GamePlayerDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GamePlayerDataServiceClient interface {
	GetGamePlayerData(ctx context.Context, in *GetGamePlayerDataRequest, opts ...grpc.CallOption) (*GetGamePlayerDataResponse, error)
	GetMultipleGamePlayerData(ctx context.Context, in *GetMultipleGamePlayerDataRequest, opts ...grpc.CallOption) (*GetMultipleGamePlayerDataResponse, error)
}

type gamePlayerDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGamePlayerDataServiceClient(cc grpc.ClientConnInterface) GamePlayerDataServiceClient {
	return &gamePlayerDataServiceClient{cc}
}

func (c *gamePlayerDataServiceClient) GetGamePlayerData(ctx context.Context, in *GetGamePlayerDataRequest, opts ...grpc.CallOption) (*GetGamePlayerDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGamePlayerDataResponse)
	err := c.cc.Invoke(ctx, GamePlayerDataService_GetGamePlayerData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gamePlayerDataServiceClient) GetMultipleGamePlayerData(ctx context.Context, in *GetMultipleGamePlayerDataRequest, opts ...grpc.CallOption) (*GetMultipleGamePlayerDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMultipleGamePlayerDataResponse)
	err := c.cc.Invoke(ctx, GamePlayerDataService_GetMultipleGamePlayerData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GamePlayerDataServiceServer is the server API for GamePlayerDataService service.
// All implementations must embed UnimplementedGamePlayerDataServiceServer
// for forward compatibility.
type GamePlayerDataServiceServer interface {
	GetGamePlayerData(context.Context, *GetGamePlayerDataRequest) (*GetGamePlayerDataResponse, error)
	GetMultipleGamePlayerData(context.Context, *GetMultipleGamePlayerDataRequest) (*GetMultipleGamePlayerDataResponse, error)
	mustEmbedUnimplementedGamePlayerDataServiceServer()
}

// UnimplementedGamePlayerDataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGamePlayerDataServiceServer struct{}

func (UnimplementedGamePlayerDataServiceServer) GetGamePlayerData(context.Context, *GetGamePlayerDataRequest) (*GetGamePlayerDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamePlayerData not implemented")
}
func (UnimplementedGamePlayerDataServiceServer) GetMultipleGamePlayerData(context.Context, *GetMultipleGamePlayerDataRequest) (*GetMultipleGamePlayerDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleGamePlayerData not implemented")
}
func (UnimplementedGamePlayerDataServiceServer) mustEmbedUnimplementedGamePlayerDataServiceServer() {}
func (UnimplementedGamePlayerDataServiceServer) testEmbeddedByValue()                               {}

// UnsafeGamePlayerDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GamePlayerDataServiceServer will
// result in compilation errors.
type UnsafeGamePlayerDataServiceServer interface {
	mustEmbedUnimplementedGamePlayerDataServiceServer()
}

func RegisterGamePlayerDataServiceServer(s grpc.ServiceRegistrar, srv GamePlayerDataServiceServer) {
	// If the following call pancis, it indicates UnimplementedGamePlayerDataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GamePlayerDataService_ServiceDesc, srv)
}

func _GamePlayerDataService_GetGamePlayerData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGamePlayerDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamePlayerDataServiceServer).GetGamePlayerData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GamePlayerDataService_GetGamePlayerData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamePlayerDataServiceServer).GetGamePlayerData(ctx, req.(*GetGamePlayerDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GamePlayerDataService_GetMultipleGamePlayerData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleGamePlayerDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GamePlayerDataServiceServer).GetMultipleGamePlayerData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GamePlayerDataService_GetMultipleGamePlayerData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GamePlayerDataServiceServer).GetMultipleGamePlayerData(ctx, req.(*GetMultipleGamePlayerDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GamePlayerDataService_ServiceDesc is the grpc.ServiceDesc for GamePlayerDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GamePlayerDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.gameplayerdata.GamePlayerDataService",
	HandlerType: (*GamePlayerDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGamePlayerData",
			Handler:    _GamePlayerDataService_GetGamePlayerData_Handler,
		},
		{
			MethodName: "GetMultipleGamePlayerData",
			Handler:    _GamePlayerDataService_GetMultipleGamePlayerData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_player_data/grpc.proto",
}
