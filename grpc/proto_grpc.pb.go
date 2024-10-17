// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: grpc/proto.proto

package proto

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
	PatientShareSendingService_SendShare_FullMethodName = "/main.PatientShareSendingService/SendShare"
)

// PatientShareSendingServiceClient is the client API for PatientShareSendingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientShareSendingServiceClient interface {
	SendShare(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknowledge, error)
}

type patientShareSendingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientShareSendingServiceClient(cc grpc.ClientConnInterface) PatientShareSendingServiceClient {
	return &patientShareSendingServiceClient{cc}
}

func (c *patientShareSendingServiceClient) SendShare(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknowledge, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Acknowledge)
	err := c.cc.Invoke(ctx, PatientShareSendingService_SendShare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientShareSendingServiceServer is the server API for PatientShareSendingService service.
// All implementations must embed UnimplementedPatientShareSendingServiceServer
// for forward compatibility.
type PatientShareSendingServiceServer interface {
	SendShare(context.Context, *Share) (*Acknowledge, error)
	mustEmbedUnimplementedPatientShareSendingServiceServer()
}

// UnimplementedPatientShareSendingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPatientShareSendingServiceServer struct{}

func (UnimplementedPatientShareSendingServiceServer) SendShare(context.Context, *Share) (*Acknowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShare not implemented")
}
func (UnimplementedPatientShareSendingServiceServer) mustEmbedUnimplementedPatientShareSendingServiceServer() {
}
func (UnimplementedPatientShareSendingServiceServer) testEmbeddedByValue() {}

// UnsafePatientShareSendingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientShareSendingServiceServer will
// result in compilation errors.
type UnsafePatientShareSendingServiceServer interface {
	mustEmbedUnimplementedPatientShareSendingServiceServer()
}

func RegisterPatientShareSendingServiceServer(s grpc.ServiceRegistrar, srv PatientShareSendingServiceServer) {
	// If the following call pancis, it indicates UnimplementedPatientShareSendingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PatientShareSendingService_ServiceDesc, srv)
}

func _PatientShareSendingService_SendShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Share)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientShareSendingServiceServer).SendShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PatientShareSendingService_SendShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientShareSendingServiceServer).SendShare(ctx, req.(*Share))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientShareSendingService_ServiceDesc is the grpc.ServiceDesc for PatientShareSendingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientShareSendingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.PatientShareSendingService",
	HandlerType: (*PatientShareSendingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendShare",
			Handler:    _PatientShareSendingService_SendShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}

const (
	AggregationSendingService_SendAggregation_FullMethodName = "/main.AggregationSendingService/SendAggregation"
)

// AggregationSendingServiceClient is the client API for AggregationSendingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregationSendingServiceClient interface {
	SendAggregation(ctx context.Context, in *Aggregation, opts ...grpc.CallOption) (*Acknowledge, error)
}

type aggregationSendingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregationSendingServiceClient(cc grpc.ClientConnInterface) AggregationSendingServiceClient {
	return &aggregationSendingServiceClient{cc}
}

func (c *aggregationSendingServiceClient) SendAggregation(ctx context.Context, in *Aggregation, opts ...grpc.CallOption) (*Acknowledge, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Acknowledge)
	err := c.cc.Invoke(ctx, AggregationSendingService_SendAggregation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregationSendingServiceServer is the server API for AggregationSendingService service.
// All implementations must embed UnimplementedAggregationSendingServiceServer
// for forward compatibility.
type AggregationSendingServiceServer interface {
	SendAggregation(context.Context, *Aggregation) (*Acknowledge, error)
	mustEmbedUnimplementedAggregationSendingServiceServer()
}

// UnimplementedAggregationSendingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAggregationSendingServiceServer struct{}

func (UnimplementedAggregationSendingServiceServer) SendAggregation(context.Context, *Aggregation) (*Acknowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAggregation not implemented")
}
func (UnimplementedAggregationSendingServiceServer) mustEmbedUnimplementedAggregationSendingServiceServer() {
}
func (UnimplementedAggregationSendingServiceServer) testEmbeddedByValue() {}

// UnsafeAggregationSendingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregationSendingServiceServer will
// result in compilation errors.
type UnsafeAggregationSendingServiceServer interface {
	mustEmbedUnimplementedAggregationSendingServiceServer()
}

func RegisterAggregationSendingServiceServer(s grpc.ServiceRegistrar, srv AggregationSendingServiceServer) {
	// If the following call pancis, it indicates UnimplementedAggregationSendingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AggregationSendingService_ServiceDesc, srv)
}

func _AggregationSendingService_SendAggregation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Aggregation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregationSendingServiceServer).SendAggregation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregationSendingService_SendAggregation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregationSendingServiceServer).SendAggregation(ctx, req.(*Aggregation))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregationSendingService_ServiceDesc is the grpc.ServiceDesc for AggregationSendingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregationSendingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.AggregationSendingService",
	HandlerType: (*AggregationSendingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendAggregation",
			Handler:    _AggregationSendingService_SendAggregation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
