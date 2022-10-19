// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: parcel-service.proto

package pb

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

// ParcelServiceClient is the client API for ParcelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParcelServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	GetParcels(ctx context.Context, in *GetParcelsRequest, opts ...grpc.CallOption) (*GetParcelsResponse, error)
	StudentGetParcels(ctx context.Context, in *StudentGetParcelsRequest, opts ...grpc.CallOption) (*StudentGetParcelsResponse, error)
	GetParcel(ctx context.Context, in *GetParcelRequest, opts ...grpc.CallOption) (*GetParcelResponse, error)
	CreateParcel(ctx context.Context, in *CreateParcelRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateParcel(ctx context.Context, in *UpdateParcelRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteParcel(ctx context.Context, in *DeleteParcelRequest, opts ...grpc.CallOption) (*Empty, error)
	StaffAcceptDelivery(ctx context.Context, in *StaffAcceptDeliveryRequest, opts ...grpc.CallOption) (*Empty, error)
	StudentClaimParcel(ctx context.Context, in *StudentClaimParcelRequest, opts ...grpc.CallOption) (*Empty, error)
}

type parcelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParcelServiceClient(cc grpc.ClientConnInterface) ParcelServiceClient {
	return &parcelServiceClient{cc}
}

func (c *parcelServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) GetParcels(ctx context.Context, in *GetParcelsRequest, opts ...grpc.CallOption) (*GetParcelsResponse, error) {
	out := new(GetParcelsResponse)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/GetParcels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) StudentGetParcels(ctx context.Context, in *StudentGetParcelsRequest, opts ...grpc.CallOption) (*StudentGetParcelsResponse, error) {
	out := new(StudentGetParcelsResponse)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/StudentGetParcels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) GetParcel(ctx context.Context, in *GetParcelRequest, opts ...grpc.CallOption) (*GetParcelResponse, error) {
	out := new(GetParcelResponse)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/GetParcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) CreateParcel(ctx context.Context, in *CreateParcelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/CreateParcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) UpdateParcel(ctx context.Context, in *UpdateParcelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/UpdateParcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) DeleteParcel(ctx context.Context, in *DeleteParcelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/DeleteParcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) StaffAcceptDelivery(ctx context.Context, in *StaffAcceptDeliveryRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/StaffAcceptDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelServiceClient) StudentClaimParcel(ctx context.Context, in *StudentClaimParcelRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ParcelService/StudentClaimParcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParcelServiceServer is the server API for ParcelService service.
// All implementations must embed UnimplementedParcelServiceServer
// for forward compatibility
type ParcelServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	GetParcels(context.Context, *GetParcelsRequest) (*GetParcelsResponse, error)
	StudentGetParcels(context.Context, *StudentGetParcelsRequest) (*StudentGetParcelsResponse, error)
	GetParcel(context.Context, *GetParcelRequest) (*GetParcelResponse, error)
	CreateParcel(context.Context, *CreateParcelRequest) (*Empty, error)
	UpdateParcel(context.Context, *UpdateParcelRequest) (*Empty, error)
	DeleteParcel(context.Context, *DeleteParcelRequest) (*Empty, error)
	StaffAcceptDelivery(context.Context, *StaffAcceptDeliveryRequest) (*Empty, error)
	StudentClaimParcel(context.Context, *StudentClaimParcelRequest) (*Empty, error)
	mustEmbedUnimplementedParcelServiceServer()
}

// UnimplementedParcelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedParcelServiceServer struct {
}

func (UnimplementedParcelServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedParcelServiceServer) GetParcels(context.Context, *GetParcelsRequest) (*GetParcelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParcels not implemented")
}
func (UnimplementedParcelServiceServer) StudentGetParcels(context.Context, *StudentGetParcelsRequest) (*StudentGetParcelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentGetParcels not implemented")
}
func (UnimplementedParcelServiceServer) GetParcel(context.Context, *GetParcelRequest) (*GetParcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParcel not implemented")
}
func (UnimplementedParcelServiceServer) CreateParcel(context.Context, *CreateParcelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParcel not implemented")
}
func (UnimplementedParcelServiceServer) UpdateParcel(context.Context, *UpdateParcelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParcel not implemented")
}
func (UnimplementedParcelServiceServer) DeleteParcel(context.Context, *DeleteParcelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParcel not implemented")
}
func (UnimplementedParcelServiceServer) StaffAcceptDelivery(context.Context, *StaffAcceptDeliveryRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StaffAcceptDelivery not implemented")
}
func (UnimplementedParcelServiceServer) StudentClaimParcel(context.Context, *StudentClaimParcelRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentClaimParcel not implemented")
}
func (UnimplementedParcelServiceServer) mustEmbedUnimplementedParcelServiceServer() {}

// UnsafeParcelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParcelServiceServer will
// result in compilation errors.
type UnsafeParcelServiceServer interface {
	mustEmbedUnimplementedParcelServiceServer()
}

func RegisterParcelServiceServer(s grpc.ServiceRegistrar, srv ParcelServiceServer) {
	s.RegisterService(&ParcelService_ServiceDesc, srv)
}

func _ParcelService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_GetParcels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParcelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).GetParcels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/GetParcels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).GetParcels(ctx, req.(*GetParcelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_StudentGetParcels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentGetParcelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).StudentGetParcels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/StudentGetParcels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).StudentGetParcels(ctx, req.(*StudentGetParcelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_GetParcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).GetParcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/GetParcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).GetParcel(ctx, req.(*GetParcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_CreateParcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateParcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).CreateParcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/CreateParcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).CreateParcel(ctx, req.(*CreateParcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_UpdateParcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).UpdateParcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/UpdateParcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).UpdateParcel(ctx, req.(*UpdateParcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_DeleteParcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteParcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).DeleteParcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/DeleteParcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).DeleteParcel(ctx, req.(*DeleteParcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_StaffAcceptDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffAcceptDeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).StaffAcceptDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/StaffAcceptDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).StaffAcceptDelivery(ctx, req.(*StaffAcceptDeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelService_StudentClaimParcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentClaimParcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelServiceServer).StudentClaimParcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ParcelService/StudentClaimParcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelServiceServer).StudentClaimParcel(ctx, req.(*StudentClaimParcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParcelService_ServiceDesc is the grpc.ServiceDesc for ParcelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParcelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ParcelService",
	HandlerType: (*ParcelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _ParcelService_Hello_Handler,
		},
		{
			MethodName: "GetParcels",
			Handler:    _ParcelService_GetParcels_Handler,
		},
		{
			MethodName: "StudentGetParcels",
			Handler:    _ParcelService_StudentGetParcels_Handler,
		},
		{
			MethodName: "GetParcel",
			Handler:    _ParcelService_GetParcel_Handler,
		},
		{
			MethodName: "CreateParcel",
			Handler:    _ParcelService_CreateParcel_Handler,
		},
		{
			MethodName: "UpdateParcel",
			Handler:    _ParcelService_UpdateParcel_Handler,
		},
		{
			MethodName: "DeleteParcel",
			Handler:    _ParcelService_DeleteParcel_Handler,
		},
		{
			MethodName: "StaffAcceptDelivery",
			Handler:    _ParcelService_StaffAcceptDelivery_Handler,
		},
		{
			MethodName: "StudentClaimParcel",
			Handler:    _ParcelService_StudentClaimParcel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parcel-service.proto",
}
