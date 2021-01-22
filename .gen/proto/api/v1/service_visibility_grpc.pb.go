// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiv1

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

// VisibilityAPIClient is the client API for VisibilityAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VisibilityAPIClient interface {
	// ListWorkflowExecutions is a visibility API to list workflow executions in a specific domain.
	ListWorkflowExecutions(ctx context.Context, in *ListWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListWorkflowExecutionsResponse, error)
	// ListOpenWorkflowExecutions is a visibility API to list the open executions in a specific domain.
	ListOpenWorkflowExecutions(ctx context.Context, in *ListOpenWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListOpenWorkflowExecutionsResponse, error)
	// ListClosedWorkflowExecutions is a visibility API to list the closed executions in a specific domain.
	ListClosedWorkflowExecutions(ctx context.Context, in *ListClosedWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListClosedWorkflowExecutionsResponse, error)
	// ListArchivedWorkflowExecutions is a visibility API to list archived workflow executions in a specific domain.
	ListArchivedWorkflowExecutions(ctx context.Context, in *ListArchivedWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListArchivedWorkflowExecutionsResponse, error)
	// ScanWorkflowExecutions is a visibility API to list large amount of workflow executions in a specific domain without order.
	ScanWorkflowExecutions(ctx context.Context, in *ScanWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ScanWorkflowExecutionsResponse, error)
	// CountWorkflowExecutions is a visibility API to count of workflow executions in a specific domain.
	CountWorkflowExecutions(ctx context.Context, in *CountWorkflowExecutionsRequest, opts ...grpc.CallOption) (*CountWorkflowExecutionsResponse, error)
	// GetSearchAttributes is a visibility API to get all legal keys that could be used in list APIs.
	GetSearchAttributes(ctx context.Context, in *GetSearchAttributesRequest, opts ...grpc.CallOption) (*GetSearchAttributesResponse, error)
}

type visibilityAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewVisibilityAPIClient(cc grpc.ClientConnInterface) VisibilityAPIClient {
	return &visibilityAPIClient{cc}
}

func (c *visibilityAPIClient) ListWorkflowExecutions(ctx context.Context, in *ListWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListWorkflowExecutionsResponse, error) {
	out := new(ListWorkflowExecutionsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/ListWorkflowExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visibilityAPIClient) ListOpenWorkflowExecutions(ctx context.Context, in *ListOpenWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListOpenWorkflowExecutionsResponse, error) {
	out := new(ListOpenWorkflowExecutionsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/ListOpenWorkflowExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visibilityAPIClient) ListClosedWorkflowExecutions(ctx context.Context, in *ListClosedWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListClosedWorkflowExecutionsResponse, error) {
	out := new(ListClosedWorkflowExecutionsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/ListClosedWorkflowExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visibilityAPIClient) ListArchivedWorkflowExecutions(ctx context.Context, in *ListArchivedWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ListArchivedWorkflowExecutionsResponse, error) {
	out := new(ListArchivedWorkflowExecutionsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/ListArchivedWorkflowExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visibilityAPIClient) ScanWorkflowExecutions(ctx context.Context, in *ScanWorkflowExecutionsRequest, opts ...grpc.CallOption) (*ScanWorkflowExecutionsResponse, error) {
	out := new(ScanWorkflowExecutionsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/ScanWorkflowExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visibilityAPIClient) CountWorkflowExecutions(ctx context.Context, in *CountWorkflowExecutionsRequest, opts ...grpc.CallOption) (*CountWorkflowExecutionsResponse, error) {
	out := new(CountWorkflowExecutionsResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/CountWorkflowExecutions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visibilityAPIClient) GetSearchAttributes(ctx context.Context, in *GetSearchAttributesRequest, opts ...grpc.CallOption) (*GetSearchAttributesResponse, error) {
	out := new(GetSearchAttributesResponse)
	err := c.cc.Invoke(ctx, "/uber.cadence.api.v1.VisibilityAPI/GetSearchAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisibilityAPIServer is the server API for VisibilityAPI service.
// All implementations must embed UnimplementedVisibilityAPIServer
// for forward compatibility
type VisibilityAPIServer interface {
	// ListWorkflowExecutions is a visibility API to list workflow executions in a specific domain.
	ListWorkflowExecutions(context.Context, *ListWorkflowExecutionsRequest) (*ListWorkflowExecutionsResponse, error)
	// ListOpenWorkflowExecutions is a visibility API to list the open executions in a specific domain.
	ListOpenWorkflowExecutions(context.Context, *ListOpenWorkflowExecutionsRequest) (*ListOpenWorkflowExecutionsResponse, error)
	// ListClosedWorkflowExecutions is a visibility API to list the closed executions in a specific domain.
	ListClosedWorkflowExecutions(context.Context, *ListClosedWorkflowExecutionsRequest) (*ListClosedWorkflowExecutionsResponse, error)
	// ListArchivedWorkflowExecutions is a visibility API to list archived workflow executions in a specific domain.
	ListArchivedWorkflowExecutions(context.Context, *ListArchivedWorkflowExecutionsRequest) (*ListArchivedWorkflowExecutionsResponse, error)
	// ScanWorkflowExecutions is a visibility API to list large amount of workflow executions in a specific domain without order.
	ScanWorkflowExecutions(context.Context, *ScanWorkflowExecutionsRequest) (*ScanWorkflowExecutionsResponse, error)
	// CountWorkflowExecutions is a visibility API to count of workflow executions in a specific domain.
	CountWorkflowExecutions(context.Context, *CountWorkflowExecutionsRequest) (*CountWorkflowExecutionsResponse, error)
	// GetSearchAttributes is a visibility API to get all legal keys that could be used in list APIs.
	GetSearchAttributes(context.Context, *GetSearchAttributesRequest) (*GetSearchAttributesResponse, error)
	mustEmbedUnimplementedVisibilityAPIServer()
}

// UnimplementedVisibilityAPIServer must be embedded to have forward compatible implementations.
type UnimplementedVisibilityAPIServer struct {
}

func (UnimplementedVisibilityAPIServer) ListWorkflowExecutions(context.Context, *ListWorkflowExecutionsRequest) (*ListWorkflowExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowExecutions not implemented")
}
func (UnimplementedVisibilityAPIServer) ListOpenWorkflowExecutions(context.Context, *ListOpenWorkflowExecutionsRequest) (*ListOpenWorkflowExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOpenWorkflowExecutions not implemented")
}
func (UnimplementedVisibilityAPIServer) ListClosedWorkflowExecutions(context.Context, *ListClosedWorkflowExecutionsRequest) (*ListClosedWorkflowExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClosedWorkflowExecutions not implemented")
}
func (UnimplementedVisibilityAPIServer) ListArchivedWorkflowExecutions(context.Context, *ListArchivedWorkflowExecutionsRequest) (*ListArchivedWorkflowExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArchivedWorkflowExecutions not implemented")
}
func (UnimplementedVisibilityAPIServer) ScanWorkflowExecutions(context.Context, *ScanWorkflowExecutionsRequest) (*ScanWorkflowExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanWorkflowExecutions not implemented")
}
func (UnimplementedVisibilityAPIServer) CountWorkflowExecutions(context.Context, *CountWorkflowExecutionsRequest) (*CountWorkflowExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountWorkflowExecutions not implemented")
}
func (UnimplementedVisibilityAPIServer) GetSearchAttributes(context.Context, *GetSearchAttributesRequest) (*GetSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchAttributes not implemented")
}
func (UnimplementedVisibilityAPIServer) mustEmbedUnimplementedVisibilityAPIServer() {}

// UnsafeVisibilityAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VisibilityAPIServer will
// result in compilation errors.
type UnsafeVisibilityAPIServer interface {
	mustEmbedUnimplementedVisibilityAPIServer()
}

func RegisterVisibilityAPIServer(s grpc.ServiceRegistrar, srv VisibilityAPIServer) {
	s.RegisterService(&VisibilityAPI_ServiceDesc, srv)
}

func _VisibilityAPI_ListWorkflowExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).ListWorkflowExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/ListWorkflowExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).ListWorkflowExecutions(ctx, req.(*ListWorkflowExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisibilityAPI_ListOpenWorkflowExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOpenWorkflowExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).ListOpenWorkflowExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/ListOpenWorkflowExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).ListOpenWorkflowExecutions(ctx, req.(*ListOpenWorkflowExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisibilityAPI_ListClosedWorkflowExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClosedWorkflowExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).ListClosedWorkflowExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/ListClosedWorkflowExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).ListClosedWorkflowExecutions(ctx, req.(*ListClosedWorkflowExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisibilityAPI_ListArchivedWorkflowExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArchivedWorkflowExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).ListArchivedWorkflowExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/ListArchivedWorkflowExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).ListArchivedWorkflowExecutions(ctx, req.(*ListArchivedWorkflowExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisibilityAPI_ScanWorkflowExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanWorkflowExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).ScanWorkflowExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/ScanWorkflowExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).ScanWorkflowExecutions(ctx, req.(*ScanWorkflowExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisibilityAPI_CountWorkflowExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountWorkflowExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).CountWorkflowExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/CountWorkflowExecutions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).CountWorkflowExecutions(ctx, req.(*CountWorkflowExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisibilityAPI_GetSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisibilityAPIServer).GetSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.cadence.api.v1.VisibilityAPI/GetSearchAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisibilityAPIServer).GetSearchAttributes(ctx, req.(*GetSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VisibilityAPI_ServiceDesc is the grpc.ServiceDesc for VisibilityAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VisibilityAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "uber.cadence.api.v1.VisibilityAPI",
	HandlerType: (*VisibilityAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkflowExecutions",
			Handler:    _VisibilityAPI_ListWorkflowExecutions_Handler,
		},
		{
			MethodName: "ListOpenWorkflowExecutions",
			Handler:    _VisibilityAPI_ListOpenWorkflowExecutions_Handler,
		},
		{
			MethodName: "ListClosedWorkflowExecutions",
			Handler:    _VisibilityAPI_ListClosedWorkflowExecutions_Handler,
		},
		{
			MethodName: "ListArchivedWorkflowExecutions",
			Handler:    _VisibilityAPI_ListArchivedWorkflowExecutions_Handler,
		},
		{
			MethodName: "ScanWorkflowExecutions",
			Handler:    _VisibilityAPI_ScanWorkflowExecutions_Handler,
		},
		{
			MethodName: "CountWorkflowExecutions",
			Handler:    _VisibilityAPI_CountWorkflowExecutions_Handler,
		},
		{
			MethodName: "GetSearchAttributes",
			Handler:    _VisibilityAPI_GetSearchAttributes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uber/cadence/api/v1/service_visibility.proto",
}
