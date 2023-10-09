// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: stream.proto

package sale_service

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

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	// server-side streaming
	Count(ctx context.Context, in *Request, opts ...grpc.CallOption) (StreamService_CountClient, error)
	// client-side streaming
	Sum(ctx context.Context, opts ...grpc.CallOption) (StreamService_SumClient, error)
	// bidirectional stream
	Sqr(ctx context.Context, opts ...grpc.CallOption) (StreamService_SqrClient, error)
	// fibonacci
	Fibonacci(ctx context.Context, opts ...grpc.CallOption) (StreamService_FibonacciClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) Count(ctx context.Context, in *Request, opts ...grpc.CallOption) (StreamService_CountClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[0], "/sale.StreamService/Count", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceCountClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_CountClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceCountClient struct {
	grpc.ClientStream
}

func (x *streamServiceCountClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Sum(ctx context.Context, opts ...grpc.CallOption) (StreamService_SumClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[1], "/sale.StreamService/Sum", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSumClient{stream}
	return x, nil
}

type StreamService_SumClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type streamServiceSumClient struct {
	grpc.ClientStream
}

func (x *streamServiceSumClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSumClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Sqr(ctx context.Context, opts ...grpc.CallOption) (StreamService_SqrClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[2], "/sale.StreamService/Sqr", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSqrClient{stream}
	return x, nil
}

type StreamService_SqrClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceSqrClient struct {
	grpc.ClientStream
}

func (x *streamServiceSqrClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSqrClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Fibonacci(ctx context.Context, opts ...grpc.CallOption) (StreamService_FibonacciClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamService_ServiceDesc.Streams[3], "/sale.StreamService/Fibonacci", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceFibonacciClient{stream}
	return x, nil
}

type StreamService_FibonacciClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type streamServiceFibonacciClient struct {
	grpc.ClientStream
}

func (x *streamServiceFibonacciClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceFibonacciClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	// server-side streaming
	Count(*Request, StreamService_CountServer) error
	// client-side streaming
	Sum(StreamService_SumServer) error
	// bidirectional stream
	Sqr(StreamService_SqrServer) error
	// fibonacci
	Fibonacci(StreamService_FibonacciServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) Count(*Request, StreamService_CountServer) error {
	return status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedStreamServiceServer) Sum(StreamService_SumServer) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedStreamServiceServer) Sqr(StreamService_SqrServer) error {
	return status.Errorf(codes.Unimplemented, "method Sqr not implemented")
}
func (UnimplementedStreamServiceServer) Fibonacci(StreamService_FibonacciServer) error {
	return status.Errorf(codes.Unimplemented, "method Fibonacci not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_Count_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).Count(m, &streamServiceCountServer{stream})
}

type StreamService_CountServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type streamServiceCountServer struct {
	grpc.ServerStream
}

func (x *streamServiceCountServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Sum(&streamServiceSumServer{stream})
}

type StreamService_SumServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type streamServiceSumServer struct {
	grpc.ServerStream
}

func (x *streamServiceSumServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSumServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Sqr_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Sqr(&streamServiceSqrServer{stream})
}

type StreamService_SqrServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type streamServiceSqrServer struct {
	grpc.ServerStream
}

func (x *streamServiceSqrServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSqrServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Fibonacci_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Fibonacci(&streamServiceFibonacciServer{stream})
}

type StreamService_FibonacciServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type streamServiceFibonacciServer struct {
	grpc.ServerStream
}

func (x *streamServiceFibonacciServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceFibonacciServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sale.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Count",
			Handler:       _StreamService_Count_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Sum",
			Handler:       _StreamService_Sum_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Sqr",
			Handler:       _StreamService_Sqr_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Fibonacci",
			Handler:       _StreamService_Fibonacci_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
