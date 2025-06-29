// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: generator/generator.proto

package generator

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
	GeneratorService_Generate_FullMethodName = "/generator.GeneratorService/Generate"
)

// GeneratorServiceClient is the client API for GeneratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneratorServiceClient interface {
	Generate(ctx context.Context, in *Query, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ResponseChunk], error)
}

type generatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneratorServiceClient(cc grpc.ClientConnInterface) GeneratorServiceClient {
	return &generatorServiceClient{cc}
}

func (c *generatorServiceClient) Generate(ctx context.Context, in *Query, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ResponseChunk], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GeneratorService_ServiceDesc.Streams[0], GeneratorService_Generate_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Query, ResponseChunk]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GeneratorService_GenerateClient = grpc.ServerStreamingClient[ResponseChunk]

// GeneratorServiceServer is the server API for GeneratorService service.
// All implementations must embed UnimplementedGeneratorServiceServer
// for forward compatibility.
type GeneratorServiceServer interface {
	Generate(*Query, grpc.ServerStreamingServer[ResponseChunk]) error
	mustEmbedUnimplementedGeneratorServiceServer()
}

// UnimplementedGeneratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGeneratorServiceServer struct{}

func (UnimplementedGeneratorServiceServer) Generate(*Query, grpc.ServerStreamingServer[ResponseChunk]) error {
	return status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedGeneratorServiceServer) mustEmbedUnimplementedGeneratorServiceServer() {}
func (UnimplementedGeneratorServiceServer) testEmbeddedByValue()                          {}

// UnsafeGeneratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneratorServiceServer will
// result in compilation errors.
type UnsafeGeneratorServiceServer interface {
	mustEmbedUnimplementedGeneratorServiceServer()
}

func RegisterGeneratorServiceServer(s grpc.ServiceRegistrar, srv GeneratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedGeneratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GeneratorService_ServiceDesc, srv)
}

func _GeneratorService_Generate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeneratorServiceServer).Generate(m, &grpc.GenericServerStream[Query, ResponseChunk]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GeneratorService_GenerateServer = grpc.ServerStreamingServer[ResponseChunk]

// GeneratorService_ServiceDesc is the grpc.ServiceDesc for GeneratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeneratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generator.GeneratorService",
	HandlerType: (*GeneratorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Generate",
			Handler:       _GeneratorService_Generate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "generator/generator.proto",
}
