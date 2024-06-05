// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: mirage_auth_service/v1.proto

package mirage_auth_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AuthService_Registration_FullMethodName = "/auth_v1.AuthService/registration"
	AuthService_Login_FullMethodName        = "/auth_v1.AuthService/login"
	AuthService_Refresh_FullMethodName      = "/auth_v1.AuthService/refresh"
	AuthService_Logout_FullMethodName       = "/auth_v1.AuthService/logout"
	AuthService_CheckAuth_FullMethodName    = "/auth_v1.AuthService/checkAuth"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Refresh(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*AuthResponse, error)
	Logout(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*Empty, error)
	CheckAuth(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*JwtUser, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Registration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Refresh(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Refresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *RefreshToken, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuthService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) CheckAuth(ctx context.Context, in *AccessToken, opts ...grpc.CallOption) (*JwtUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JwtUser)
	err := c.cc.Invoke(ctx, AuthService_CheckAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Registration(context.Context, *RegistrationRequest) (*AuthResponse, error)
	Login(context.Context, *LoginRequest) (*AuthResponse, error)
	Refresh(context.Context, *RefreshToken) (*AuthResponse, error)
	Logout(context.Context, *RefreshToken) (*Empty, error)
	CheckAuth(context.Context, *AccessToken) (*JwtUser, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Registration(context.Context, *RegistrationRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Refresh(context.Context, *RefreshToken) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAuthServiceServer) Logout(context.Context, *RefreshToken) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceServer) CheckAuth(context.Context, *AccessToken) (*JwtUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
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

func _AuthService_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Registration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Refresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Refresh(ctx, req.(*RefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*RefreshToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_CheckAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckAuth(ctx, req.(*AccessToken))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registration",
			Handler:    _AuthService_Registration_Handler,
		},
		{
			MethodName: "login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "refresh",
			Handler:    _AuthService_Refresh_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "checkAuth",
			Handler:    _AuthService_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mirage_auth_service/v1.proto",
}

const (
	UserService_GetUserById_FullMethodName           = "/auth_v1.UserService/getUserById"
	UserService_GetUsersByQuery_FullMethodName       = "/auth_v1.UserService/getUsersByQuery"
	UserService_UpdateUserRole_FullMethodName        = "/auth_v1.UserService/updateUserRole"
	UserService_DeleteUserById_FullMethodName        = "/auth_v1.UserService/deleteUserById"
	UserService_ChangePassword_FullMethodName        = "/auth_v1.UserService/changePassword"
	UserService_CheckRole_FullMethodName             = "/auth_v1.UserService/checkRole"
	UserService_ConfirmChangePassword_FullMethodName = "/auth_v1.UserService/confirmChangePassword"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseUser, error)
	GetUsersByQuery(ctx context.Context, in *PaginationQuery, opts ...grpc.CallOption) (*ManyResponseUser, error)
	UpdateUserRole(ctx context.Context, in *UpdateUserRole, opts ...grpc.CallOption) (*ResponseUser, error)
	DeleteUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Empty, error)
	CheckRole(ctx context.Context, in *CheckRoleRequest, opts ...grpc.CallOption) (*Bool, error)
	ConfirmChangePassword(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUser)
	err := c.cc.Invoke(ctx, UserService_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersByQuery(ctx context.Context, in *PaginationQuery, opts ...grpc.CallOption) (*ManyResponseUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ManyResponseUser)
	err := c.cc.Invoke(ctx, UserService_GetUsersByQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserRole(ctx context.Context, in *UpdateUserRole, opts ...grpc.CallOption) (*ResponseUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUser)
	err := c.cc.Invoke(ctx, UserService_UpdateUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckRole(ctx context.Context, in *CheckRoleRequest, opts ...grpc.CallOption) (*Bool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Bool)
	err := c.cc.Invoke(ctx, UserService_CheckRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmChangePassword(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UserService_ConfirmChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserById(context.Context, *Id) (*ResponseUser, error)
	GetUsersByQuery(context.Context, *PaginationQuery) (*ManyResponseUser, error)
	UpdateUserRole(context.Context, *UpdateUserRole) (*ResponseUser, error)
	DeleteUserById(context.Context, *Id) (*Empty, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*Empty, error)
	CheckRole(context.Context, *CheckRoleRequest) (*Bool, error)
	ConfirmChangePassword(context.Context, *Link) (*Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserById(context.Context, *Id) (*ResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUsersByQuery(context.Context, *PaginationQuery) (*ManyResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByQuery not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserRole(context.Context, *UpdateUserRole) (*ResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserById(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserById not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) CheckRole(context.Context, *CheckRoleRequest) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRole not implemented")
}
func (UnimplementedUserServiceServer) ConfirmChangePassword(context.Context, *Link) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmChangePassword not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaginationQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsersByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUsersByQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsersByQuery(ctx, req.(*PaginationQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserRole(ctx, req.(*UpdateUserRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckRole(ctx, req.(*CheckRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ConfirmChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmChangePassword(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "getUsersByQuery",
			Handler:    _UserService_GetUsersByQuery_Handler,
		},
		{
			MethodName: "updateUserRole",
			Handler:    _UserService_UpdateUserRole_Handler,
		},
		{
			MethodName: "deleteUserById",
			Handler:    _UserService_DeleteUserById_Handler,
		},
		{
			MethodName: "changePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "checkRole",
			Handler:    _UserService_CheckRole_Handler,
		},
		{
			MethodName: "confirmChangePassword",
			Handler:    _UserService_ConfirmChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mirage_auth_service/v1.proto",
}

const (
	UtilityService_LinkActivate_FullMethodName   = "/auth_v1.UtilityService/linkActivate"
	UtilityService_IsUserActivate_FullMethodName = "/auth_v1.UtilityService/isUserActivate"
	UtilityService_BanUser_FullMethodName        = "/auth_v1.UtilityService/banUser"
	UtilityService_UnbanUser_FullMethodName      = "/auth_v1.UtilityService/unbanUser"
)

// UtilityServiceClient is the client API for UtilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UtilityServiceClient interface {
	LinkActivate(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error)
	IsUserActivate(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Bool, error)
	BanUser(ctx context.Context, in *BanUser, opts ...grpc.CallOption) (*ResponseUser, error)
	UnbanUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseUser, error)
}

type utilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUtilityServiceClient(cc grpc.ClientConnInterface) UtilityServiceClient {
	return &utilityServiceClient{cc}
}

func (c *utilityServiceClient) LinkActivate(ctx context.Context, in *Link, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, UtilityService_LinkActivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilityServiceClient) IsUserActivate(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Bool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Bool)
	err := c.cc.Invoke(ctx, UtilityService_IsUserActivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilityServiceClient) BanUser(ctx context.Context, in *BanUser, opts ...grpc.CallOption) (*ResponseUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUser)
	err := c.cc.Invoke(ctx, UtilityService_BanUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilityServiceClient) UnbanUser(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ResponseUser, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseUser)
	err := c.cc.Invoke(ctx, UtilityService_UnbanUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UtilityServiceServer is the server API for UtilityService service.
// All implementations must embed UnimplementedUtilityServiceServer
// for forward compatibility
type UtilityServiceServer interface {
	LinkActivate(context.Context, *Link) (*Empty, error)
	IsUserActivate(context.Context, *Id) (*Bool, error)
	BanUser(context.Context, *BanUser) (*ResponseUser, error)
	UnbanUser(context.Context, *Id) (*ResponseUser, error)
	mustEmbedUnimplementedUtilityServiceServer()
}

// UnimplementedUtilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUtilityServiceServer struct {
}

func (UnimplementedUtilityServiceServer) LinkActivate(context.Context, *Link) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkActivate not implemented")
}
func (UnimplementedUtilityServiceServer) IsUserActivate(context.Context, *Id) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUserActivate not implemented")
}
func (UnimplementedUtilityServiceServer) BanUser(context.Context, *BanUser) (*ResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BanUser not implemented")
}
func (UnimplementedUtilityServiceServer) UnbanUser(context.Context, *Id) (*ResponseUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbanUser not implemented")
}
func (UnimplementedUtilityServiceServer) mustEmbedUnimplementedUtilityServiceServer() {}

// UnsafeUtilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UtilityServiceServer will
// result in compilation errors.
type UnsafeUtilityServiceServer interface {
	mustEmbedUnimplementedUtilityServiceServer()
}

func RegisterUtilityServiceServer(s grpc.ServiceRegistrar, srv UtilityServiceServer) {
	s.RegisterService(&UtilityService_ServiceDesc, srv)
}

func _UtilityService_LinkActivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServiceServer).LinkActivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UtilityService_LinkActivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServiceServer).LinkActivate(ctx, req.(*Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _UtilityService_IsUserActivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServiceServer).IsUserActivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UtilityService_IsUserActivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServiceServer).IsUserActivate(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _UtilityService_BanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServiceServer).BanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UtilityService_BanUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServiceServer).BanUser(ctx, req.(*BanUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UtilityService_UnbanUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilityServiceServer).UnbanUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UtilityService_UnbanUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilityServiceServer).UnbanUser(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// UtilityService_ServiceDesc is the grpc.ServiceDesc for UtilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UtilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.UtilityService",
	HandlerType: (*UtilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "linkActivate",
			Handler:    _UtilityService_LinkActivate_Handler,
		},
		{
			MethodName: "isUserActivate",
			Handler:    _UtilityService_IsUserActivate_Handler,
		},
		{
			MethodName: "banUser",
			Handler:    _UtilityService_BanUser_Handler,
		},
		{
			MethodName: "unbanUser",
			Handler:    _UtilityService_UnbanUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mirage_auth_service/v1.proto",
}