// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.0
// source: relationship/grpc.proto

package relationship

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

// RelationshipClient is the client API for Relationship service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationshipClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
	RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error)
	DenyFriendRequest(ctx context.Context, in *DenyFriendRequestRequest, opts ...grpc.CallOption) (*DenyFriendRequestResponse, error)
	MassDenyFriendRequest(ctx context.Context, in *MassDenyFriendRequestRequest, opts ...grpc.CallOption) (*MassDenyFriendRequestResponse, error)
	GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
	GetPendingFriendRequestList(ctx context.Context, in *GetPendingFriendRequestListRequest, opts ...grpc.CallOption) (*PendingFriendListResponse, error)
	CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*CreateBlockResponse, error)
	DeleteBlock(ctx context.Context, in *DeleteBlockRequest, opts ...grpc.CallOption) (*DeleteBlockResponse, error)
	IsBlocked(ctx context.Context, in *IsBlockedRequest, opts ...grpc.CallOption) (*IsBlockedResponse, error)
	GetBlockedList(ctx context.Context, in *GetBlockedListRequest, opts ...grpc.CallOption) (*BlockedListResponse, error)
}

type relationshipClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationshipClient(cc grpc.ClientConnInterface) RelationshipClient {
	return &relationshipClient{cc}
}

func (c *relationshipClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	out := new(AddFriendResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*RemoveFriendResponse, error) {
	out := new(RemoveFriendResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/RemoveFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) DenyFriendRequest(ctx context.Context, in *DenyFriendRequestRequest, opts ...grpc.CallOption) (*DenyFriendRequestResponse, error) {
	out := new(DenyFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/DenyFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) MassDenyFriendRequest(ctx context.Context, in *MassDenyFriendRequestRequest, opts ...grpc.CallOption) (*MassDenyFriendRequestResponse, error) {
	out := new(MassDenyFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/MassDenyFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	out := new(FriendListResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/GetFriendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) GetPendingFriendRequestList(ctx context.Context, in *GetPendingFriendRequestListRequest, opts ...grpc.CallOption) (*PendingFriendListResponse, error) {
	out := new(PendingFriendListResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/GetPendingFriendRequestList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) CreateBlock(ctx context.Context, in *CreateBlockRequest, opts ...grpc.CallOption) (*CreateBlockResponse, error) {
	out := new(CreateBlockResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/CreateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) DeleteBlock(ctx context.Context, in *DeleteBlockRequest, opts ...grpc.CallOption) (*DeleteBlockResponse, error) {
	out := new(DeleteBlockResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/DeleteBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) IsBlocked(ctx context.Context, in *IsBlockedRequest, opts ...grpc.CallOption) (*IsBlockedResponse, error) {
	out := new(IsBlockedResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/IsBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipClient) GetBlockedList(ctx context.Context, in *GetBlockedListRequest, opts ...grpc.CallOption) (*BlockedListResponse, error) {
	out := new(BlockedListResponse)
	err := c.cc.Invoke(ctx, "/emortal.grpc.relationship.Relationship/GetBlockedList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationshipServer is the server API for Relationship service.
// All implementations must embed UnimplementedRelationshipServer
// for forward compatibility
type RelationshipServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error)
	RemoveFriend(context.Context, *RemoveFriendRequest) (*RemoveFriendResponse, error)
	DenyFriendRequest(context.Context, *DenyFriendRequestRequest) (*DenyFriendRequestResponse, error)
	MassDenyFriendRequest(context.Context, *MassDenyFriendRequestRequest) (*MassDenyFriendRequestResponse, error)
	GetFriendList(context.Context, *GetFriendListRequest) (*FriendListResponse, error)
	GetPendingFriendRequestList(context.Context, *GetPendingFriendRequestListRequest) (*PendingFriendListResponse, error)
	CreateBlock(context.Context, *CreateBlockRequest) (*CreateBlockResponse, error)
	DeleteBlock(context.Context, *DeleteBlockRequest) (*DeleteBlockResponse, error)
	IsBlocked(context.Context, *IsBlockedRequest) (*IsBlockedResponse, error)
	GetBlockedList(context.Context, *GetBlockedListRequest) (*BlockedListResponse, error)
	mustEmbedUnimplementedRelationshipServer()
}

// UnimplementedRelationshipServer must be embedded to have forward compatible implementations.
type UnimplementedRelationshipServer struct {
}

func (UnimplementedRelationshipServer) AddFriend(context.Context, *AddFriendRequest) (*AddFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedRelationshipServer) RemoveFriend(context.Context, *RemoveFriendRequest) (*RemoveFriendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriend not implemented")
}
func (UnimplementedRelationshipServer) DenyFriendRequest(context.Context, *DenyFriendRequestRequest) (*DenyFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenyFriendRequest not implemented")
}
func (UnimplementedRelationshipServer) MassDenyFriendRequest(context.Context, *MassDenyFriendRequestRequest) (*MassDenyFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MassDenyFriendRequest not implemented")
}
func (UnimplementedRelationshipServer) GetFriendList(context.Context, *GetFriendListRequest) (*FriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendList not implemented")
}
func (UnimplementedRelationshipServer) GetPendingFriendRequestList(context.Context, *GetPendingFriendRequestListRequest) (*PendingFriendListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPendingFriendRequestList not implemented")
}
func (UnimplementedRelationshipServer) CreateBlock(context.Context, *CreateBlockRequest) (*CreateBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlock not implemented")
}
func (UnimplementedRelationshipServer) DeleteBlock(context.Context, *DeleteBlockRequest) (*DeleteBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlock not implemented")
}
func (UnimplementedRelationshipServer) IsBlocked(context.Context, *IsBlockedRequest) (*IsBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlocked not implemented")
}
func (UnimplementedRelationshipServer) GetBlockedList(context.Context, *GetBlockedListRequest) (*BlockedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockedList not implemented")
}
func (UnimplementedRelationshipServer) mustEmbedUnimplementedRelationshipServer() {}

// UnsafeRelationshipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationshipServer will
// result in compilation errors.
type UnsafeRelationshipServer interface {
	mustEmbedUnimplementedRelationshipServer()
}

func RegisterRelationshipServer(s grpc.ServiceRegistrar, srv RelationshipServer) {
	s.RegisterService(&Relationship_ServiceDesc, srv)
}

func _Relationship_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_RemoveFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).RemoveFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/RemoveFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).RemoveFriend(ctx, req.(*RemoveFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_DenyFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DenyFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).DenyFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/DenyFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).DenyFriendRequest(ctx, req.(*DenyFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_MassDenyFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MassDenyFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).MassDenyFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/MassDenyFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).MassDenyFriendRequest(ctx, req.(*MassDenyFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_GetFriendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).GetFriendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/GetFriendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).GetFriendList(ctx, req.(*GetFriendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_GetPendingFriendRequestList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPendingFriendRequestListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).GetPendingFriendRequestList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/GetPendingFriendRequestList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).GetPendingFriendRequestList(ctx, req.(*GetPendingFriendRequestListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_CreateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).CreateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/CreateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).CreateBlock(ctx, req.(*CreateBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_DeleteBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).DeleteBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/DeleteBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).DeleteBlock(ctx, req.(*DeleteBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_IsBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsBlockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).IsBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/IsBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).IsBlocked(ctx, req.(*IsBlockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relationship_GetBlockedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServer).GetBlockedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/emortal.grpc.relationship.Relationship/GetBlockedList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServer).GetBlockedList(ctx, req.(*GetBlockedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Relationship_ServiceDesc is the grpc.ServiceDesc for Relationship service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relationship_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emortal.grpc.relationship.Relationship",
	HandlerType: (*RelationshipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _Relationship_AddFriend_Handler,
		},
		{
			MethodName: "RemoveFriend",
			Handler:    _Relationship_RemoveFriend_Handler,
		},
		{
			MethodName: "DenyFriendRequest",
			Handler:    _Relationship_DenyFriendRequest_Handler,
		},
		{
			MethodName: "MassDenyFriendRequest",
			Handler:    _Relationship_MassDenyFriendRequest_Handler,
		},
		{
			MethodName: "GetFriendList",
			Handler:    _Relationship_GetFriendList_Handler,
		},
		{
			MethodName: "GetPendingFriendRequestList",
			Handler:    _Relationship_GetPendingFriendRequestList_Handler,
		},
		{
			MethodName: "CreateBlock",
			Handler:    _Relationship_CreateBlock_Handler,
		},
		{
			MethodName: "DeleteBlock",
			Handler:    _Relationship_DeleteBlock_Handler,
		},
		{
			MethodName: "IsBlocked",
			Handler:    _Relationship_IsBlocked_Handler,
		},
		{
			MethodName: "GetBlockedList",
			Handler:    _Relationship_GetBlockedList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relationship/grpc.proto",
}
