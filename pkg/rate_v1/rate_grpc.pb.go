// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: rate.proto

package rate_v1

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

// RateV1Client is the client API for RateV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RateV1Client interface {
	FindRate(ctx context.Context, in *FindRateRequest, opts ...grpc.CallOption) (*FindRateResponse, error)
	FindSupRates(ctx context.Context, in *FindSupRatesRequest, opts ...grpc.CallOption) (*FindSupRatesResponse, error)
}

type rateV1Client struct {
	cc grpc.ClientConnInterface
}

func NewRateV1Client(cc grpc.ClientConnInterface) RateV1Client {
	return &rateV1Client{cc}
}

func (c *rateV1Client) FindRate(ctx context.Context, in *FindRateRequest, opts ...grpc.CallOption) (*FindRateResponse, error) {
	out := new(FindRateResponse)
	err := c.cc.Invoke(ctx, "/rate_v1.RateV1/FindRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateV1Client) FindSupRates(ctx context.Context, in *FindSupRatesRequest, opts ...grpc.CallOption) (*FindSupRatesResponse, error) {
	out := new(FindSupRatesResponse)
	err := c.cc.Invoke(ctx, "/rate_v1.RateV1/FindSupRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateV1Server is the server API for RateV1 service.
// All implementations must embed UnimplementedRateV1Server
// for forward compatibility
type RateV1Server interface {
	FindRate(context.Context, *FindRateRequest) (*FindRateResponse, error)
	FindSupRates(context.Context, *FindSupRatesRequest) (*FindSupRatesResponse, error)
	mustEmbedUnimplementedRateV1Server()
}

// UnimplementedRateV1Server must be embedded to have forward compatible implementations.
type UnimplementedRateV1Server struct {
}

func (UnimplementedRateV1Server) FindRate(context.Context, *FindRateRequest) (*FindRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRate not implemented")
}
func (UnimplementedRateV1Server) FindSupRates(context.Context, *FindSupRatesRequest) (*FindSupRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSupRates not implemented")
}
func (UnimplementedRateV1Server) mustEmbedUnimplementedRateV1Server() {}

// UnsafeRateV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RateV1Server will
// result in compilation errors.
type UnsafeRateV1Server interface {
	mustEmbedUnimplementedRateV1Server()
}

func RegisterRateV1Server(s grpc.ServiceRegistrar, srv RateV1Server) {
	s.RegisterService(&RateV1_ServiceDesc, srv)
}

func _RateV1_FindRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateV1Server).FindRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate_v1.RateV1/FindRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateV1Server).FindRate(ctx, req.(*FindRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateV1_FindSupRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSupRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateV1Server).FindSupRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate_v1.RateV1/FindSupRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateV1Server).FindSupRates(ctx, req.(*FindSupRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RateV1_ServiceDesc is the grpc.ServiceDesc for RateV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RateV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rate_v1.RateV1",
	HandlerType: (*RateV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindRate",
			Handler:    _RateV1_FindRate_Handler,
		},
		{
			MethodName: "FindSupRates",
			Handler:    _RateV1_FindSupRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
