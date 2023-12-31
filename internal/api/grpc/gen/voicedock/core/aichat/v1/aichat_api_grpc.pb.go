// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: voicedock/core/aichat/v1/aichat_api.proto

package aichatv1

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
	AichatAPI_Generate_FullMethodName      = "/voicedock.core.aichat.v1.AichatAPI/Generate"
	AichatAPI_GetModels_FullMethodName     = "/voicedock.core.aichat.v1.AichatAPI/GetModels"
	AichatAPI_DownloadModel_FullMethodName = "/voicedock.core.aichat.v1.AichatAPI/DownloadModel"
)

// AichatAPIClient is the client API for AichatAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AichatAPIClient interface {
	// Generate response text by prompt.
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (AichatAPI_GenerateClient, error)
	// Returns available ai chat models.
	GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error)
	// Downloads selected ai model.
	DownloadModel(ctx context.Context, in *DownloadModelRequest, opts ...grpc.CallOption) (*DownloadModelResponse, error)
}

type aichatAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAichatAPIClient(cc grpc.ClientConnInterface) AichatAPIClient {
	return &aichatAPIClient{cc}
}

func (c *aichatAPIClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (AichatAPI_GenerateClient, error) {
	stream, err := c.cc.NewStream(ctx, &AichatAPI_ServiceDesc.Streams[0], AichatAPI_Generate_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &aichatAPIGenerateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AichatAPI_GenerateClient interface {
	Recv() (*GenerateResponse, error)
	grpc.ClientStream
}

type aichatAPIGenerateClient struct {
	grpc.ClientStream
}

func (x *aichatAPIGenerateClient) Recv() (*GenerateResponse, error) {
	m := new(GenerateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aichatAPIClient) GetModels(ctx context.Context, in *GetModelsRequest, opts ...grpc.CallOption) (*GetModelsResponse, error) {
	out := new(GetModelsResponse)
	err := c.cc.Invoke(ctx, AichatAPI_GetModels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aichatAPIClient) DownloadModel(ctx context.Context, in *DownloadModelRequest, opts ...grpc.CallOption) (*DownloadModelResponse, error) {
	out := new(DownloadModelResponse)
	err := c.cc.Invoke(ctx, AichatAPI_DownloadModel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AichatAPIServer is the server API for AichatAPI service.
// All implementations must embed UnimplementedAichatAPIServer
// for forward compatibility
type AichatAPIServer interface {
	// Generate response text by prompt.
	Generate(*GenerateRequest, AichatAPI_GenerateServer) error
	// Returns available ai chat models.
	GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error)
	// Downloads selected ai model.
	DownloadModel(context.Context, *DownloadModelRequest) (*DownloadModelResponse, error)
	mustEmbedUnimplementedAichatAPIServer()
}

// UnimplementedAichatAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAichatAPIServer struct {
}

func (UnimplementedAichatAPIServer) Generate(*GenerateRequest, AichatAPI_GenerateServer) error {
	return status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedAichatAPIServer) GetModels(context.Context, *GetModelsRequest) (*GetModelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModels not implemented")
}
func (UnimplementedAichatAPIServer) DownloadModel(context.Context, *DownloadModelRequest) (*DownloadModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadModel not implemented")
}
func (UnimplementedAichatAPIServer) mustEmbedUnimplementedAichatAPIServer() {}

// UnsafeAichatAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AichatAPIServer will
// result in compilation errors.
type UnsafeAichatAPIServer interface {
	mustEmbedUnimplementedAichatAPIServer()
}

func RegisterAichatAPIServer(s grpc.ServiceRegistrar, srv AichatAPIServer) {
	s.RegisterService(&AichatAPI_ServiceDesc, srv)
}

func _AichatAPI_Generate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GenerateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AichatAPIServer).Generate(m, &aichatAPIGenerateServer{stream})
}

type AichatAPI_GenerateServer interface {
	Send(*GenerateResponse) error
	grpc.ServerStream
}

type aichatAPIGenerateServer struct {
	grpc.ServerStream
}

func (x *aichatAPIGenerateServer) Send(m *GenerateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AichatAPI_GetModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AichatAPIServer).GetModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AichatAPI_GetModels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AichatAPIServer).GetModels(ctx, req.(*GetModelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AichatAPI_DownloadModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AichatAPIServer).DownloadModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AichatAPI_DownloadModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AichatAPIServer).DownloadModel(ctx, req.(*DownloadModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AichatAPI_ServiceDesc is the grpc.ServiceDesc for AichatAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AichatAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voicedock.core.aichat.v1.AichatAPI",
	HandlerType: (*AichatAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModels",
			Handler:    _AichatAPI_GetModels_Handler,
		},
		{
			MethodName: "DownloadModel",
			Handler:    _AichatAPI_DownloadModel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Generate",
			Handler:       _AichatAPI_Generate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "voicedock/core/aichat/v1/aichat_api.proto",
}
