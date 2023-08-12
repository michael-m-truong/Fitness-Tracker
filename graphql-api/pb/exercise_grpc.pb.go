// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: pb/exercise.proto

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

// ExerciseServiceClient is the client API for ExerciseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExerciseServiceClient interface {
	CreateExercise(ctx context.Context, in *NewExercise, opts ...grpc.CallOption) (*Exercise, error)
}

type exerciseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExerciseServiceClient(cc grpc.ClientConnInterface) ExerciseServiceClient {
	return &exerciseServiceClient{cc}
}

func (c *exerciseServiceClient) CreateExercise(ctx context.Context, in *NewExercise, opts ...grpc.CallOption) (*Exercise, error) {
	out := new(Exercise)
	err := c.cc.Invoke(ctx, "/pb.ExerciseService/CreateExercise", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExerciseServiceServer is the server API for ExerciseService service.
// All implementations must embed UnimplementedExerciseServiceServer
// for forward compatibility
type ExerciseServiceServer interface {
	CreateExercise(context.Context, *NewExercise) (*Exercise, error)
	mustEmbedUnimplementedExerciseServiceServer()
}

// UnimplementedExerciseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExerciseServiceServer struct {
}

func (UnimplementedExerciseServiceServer) CreateExercise(context.Context, *NewExercise) (*Exercise, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExercise not implemented")
}
func (UnimplementedExerciseServiceServer) mustEmbedUnimplementedExerciseServiceServer() {}

// UnsafeExerciseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExerciseServiceServer will
// result in compilation errors.
type UnsafeExerciseServiceServer interface {
	mustEmbedUnimplementedExerciseServiceServer()
}

func RegisterExerciseServiceServer(s grpc.ServiceRegistrar, srv ExerciseServiceServer) {
	s.RegisterService(&ExerciseService_ServiceDesc, srv)
}

func _ExerciseService_CreateExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewExercise)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExerciseServiceServer).CreateExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExerciseService/CreateExercise",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExerciseServiceServer).CreateExercise(ctx, req.(*NewExercise))
	}
	return interceptor(ctx, in, info, handler)
}

// ExerciseService_ServiceDesc is the grpc.ServiceDesc for ExerciseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExerciseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ExerciseService",
	HandlerType: (*ExerciseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExercise",
			Handler:    _ExerciseService_CreateExercise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/exercise.proto",
}

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*AccessToken, error)
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*NewUser, error)
	ParseToken(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*User, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*AccessToken, error) {
	out := new(AccessToken)
	err := c.cc.Invoke(ctx, "/pb.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*NewUser, error) {
	out := new(NewUser)
	err := c.cc.Invoke(ctx, "/pb.AuthService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ParseToken(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.AuthService/ParseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *User) (*AccessToken, error)
	CreateUser(context.Context, *User) (*NewUser, error)
	ParseToken(context.Context, *AccessToken) (*User, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *User) (*AccessToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) CreateUser(context.Context, *User) (*NewUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthServiceServer) ParseToken(context.Context, *AccessToken) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseToken not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ParseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ParseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AuthService/ParseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ParseToken(ctx, req.(*AccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _AuthService_CreateUser_Handler,
		},
		{
			MethodName: "ParseToken",
			Handler:    _AuthService_ParseToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/exercise.proto",
}