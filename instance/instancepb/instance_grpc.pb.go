// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: instancepb/instance.proto

package instancepb

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
	Instance_GetInstancesByRegion_FullMethodName = "/instancepb.instance/GetInstancesByRegion"
	Instance_SendStatusUpdates_FullMethodName    = "/instancepb.instance/SendStatusUpdates"
)

// InstanceClient is the client API for Instance service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceClient interface {
	GetInstancesByRegion(ctx context.Context, in *GetInstancesByRegionRequest, opts ...grpc.CallOption) (Instance_GetInstancesByRegionClient, error)
	SendStatusUpdates(ctx context.Context, in *GetInstancesByRegionRequest, opts ...grpc.CallOption) (Instance_SendStatusUpdatesClient, error)
}

type instanceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceClient(cc grpc.ClientConnInterface) InstanceClient {
	return &instanceClient{cc}
}

func (c *instanceClient) GetInstancesByRegion(ctx context.Context, in *GetInstancesByRegionRequest, opts ...grpc.CallOption) (Instance_GetInstancesByRegionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Instance_ServiceDesc.Streams[0], Instance_GetInstancesByRegion_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &instanceGetInstancesByRegionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Instance_GetInstancesByRegionClient interface {
	Recv() (*GetInstancesByRegionResponse, error)
	grpc.ClientStream
}

type instanceGetInstancesByRegionClient struct {
	grpc.ClientStream
}

func (x *instanceGetInstancesByRegionClient) Recv() (*GetInstancesByRegionResponse, error) {
	m := new(GetInstancesByRegionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *instanceClient) SendStatusUpdates(ctx context.Context, in *GetInstancesByRegionRequest, opts ...grpc.CallOption) (Instance_SendStatusUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Instance_ServiceDesc.Streams[1], Instance_SendStatusUpdates_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &instanceSendStatusUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Instance_SendStatusUpdatesClient interface {
	Recv() (*StatusUpdate, error)
	grpc.ClientStream
}

type instanceSendStatusUpdatesClient struct {
	grpc.ClientStream
}

func (x *instanceSendStatusUpdatesClient) Recv() (*StatusUpdate, error) {
	m := new(StatusUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InstanceServer is the server API for Instance service.
// All implementations must embed UnimplementedInstanceServer
// for forward compatibility
type InstanceServer interface {
	GetInstancesByRegion(*GetInstancesByRegionRequest, Instance_GetInstancesByRegionServer) error
	SendStatusUpdates(*GetInstancesByRegionRequest, Instance_SendStatusUpdatesServer) error
	mustEmbedUnimplementedInstanceServer()
}

// UnimplementedInstanceServer must be embedded to have forward compatible implementations.
type UnimplementedInstanceServer struct {
}

func (UnimplementedInstanceServer) GetInstancesByRegion(*GetInstancesByRegionRequest, Instance_GetInstancesByRegionServer) error {
	return status.Errorf(codes.Unimplemented, "method GetInstancesByRegion not implemented")
}
func (UnimplementedInstanceServer) SendStatusUpdates(*GetInstancesByRegionRequest, Instance_SendStatusUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SendStatusUpdates not implemented")
}
func (UnimplementedInstanceServer) mustEmbedUnimplementedInstanceServer() {}

// UnsafeInstanceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServer will
// result in compilation errors.
type UnsafeInstanceServer interface {
	mustEmbedUnimplementedInstanceServer()
}

func RegisterInstanceServer(s grpc.ServiceRegistrar, srv InstanceServer) {
	s.RegisterService(&Instance_ServiceDesc, srv)
}

func _Instance_GetInstancesByRegion_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetInstancesByRegionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InstanceServer).GetInstancesByRegion(m, &instanceGetInstancesByRegionServer{stream})
}

type Instance_GetInstancesByRegionServer interface {
	Send(*GetInstancesByRegionResponse) error
	grpc.ServerStream
}

type instanceGetInstancesByRegionServer struct {
	grpc.ServerStream
}

func (x *instanceGetInstancesByRegionServer) Send(m *GetInstancesByRegionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Instance_SendStatusUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetInstancesByRegionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InstanceServer).SendStatusUpdates(m, &instanceSendStatusUpdatesServer{stream})
}

type Instance_SendStatusUpdatesServer interface {
	Send(*StatusUpdate) error
	grpc.ServerStream
}

type instanceSendStatusUpdatesServer struct {
	grpc.ServerStream
}

func (x *instanceSendStatusUpdatesServer) Send(m *StatusUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// Instance_ServiceDesc is the grpc.ServiceDesc for Instance service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Instance_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "instancepb.instance",
	HandlerType: (*InstanceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetInstancesByRegion",
			Handler:       _Instance_GetInstancesByRegion_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendStatusUpdates",
			Handler:       _Instance_SendStatusUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "instancepb/instance.proto",
}
