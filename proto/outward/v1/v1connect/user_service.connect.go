// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: outward/v1/user_service.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/invzhi/outward/proto/outward/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "outward.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceCreateUserProcedure is the fully-qualified name of the UserService's CreateUser RPC.
	UserServiceCreateUserProcedure = "/outward.v1.UserService/CreateUser"
	// UserServiceGetUserListProcedure is the fully-qualified name of the UserService's GetUserList RPC.
	UserServiceGetUserListProcedure = "/outward.v1.UserService/GetUserList"
	// UserServiceCreateWorkspaceMemberProcedure is the fully-qualified name of the UserService's
	// CreateWorkspaceMember RPC.
	UserServiceCreateWorkspaceMemberProcedure = "/outward.v1.UserService/CreateWorkspaceMember"
	// UserServiceLoginProcedure is the fully-qualified name of the UserService's Login RPC.
	UserServiceLoginProcedure = "/outward.v1.UserService/Login"
	// UserServiceLogoutProcedure is the fully-qualified name of the UserService's Logout RPC.
	UserServiceLogoutProcedure = "/outward.v1.UserService/Logout"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor                     = v1.File_outward_v1_user_service_proto.Services().ByName("UserService")
	userServiceCreateUserMethodDescriptor            = userServiceServiceDescriptor.Methods().ByName("CreateUser")
	userServiceGetUserListMethodDescriptor           = userServiceServiceDescriptor.Methods().ByName("GetUserList")
	userServiceCreateWorkspaceMemberMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("CreateWorkspaceMember")
	userServiceLoginMethodDescriptor                 = userServiceServiceDescriptor.Methods().ByName("Login")
	userServiceLogoutMethodDescriptor                = userServiceServiceDescriptor.Methods().ByName("Logout")
)

// UserServiceClient is a client for the outward.v1.UserService service.
type UserServiceClient interface {
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error)
	GetUserList(context.Context, *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error)
	CreateWorkspaceMember(context.Context, *connect.Request[v1.CreateWorkspaceMemberRequest]) (*connect.Response[v1.CreateWorkspaceMemberResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error)
}

// NewUserServiceClient constructs a client for the outward.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		createUser: connect.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+UserServiceCreateUserProcedure,
			connect.WithSchema(userServiceCreateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUserList: connect.NewClient[v1.GetUserListRequest, v1.GetUserListResponse](
			httpClient,
			baseURL+UserServiceGetUserListProcedure,
			connect.WithSchema(userServiceGetUserListMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createWorkspaceMember: connect.NewClient[v1.CreateWorkspaceMemberRequest, v1.CreateWorkspaceMemberResponse](
			httpClient,
			baseURL+UserServiceCreateWorkspaceMemberProcedure,
			connect.WithSchema(userServiceCreateWorkspaceMemberMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+UserServiceLoginProcedure,
			connect.WithSchema(userServiceLoginMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		logout: connect.NewClient[v1.LogoutRequest, v1.LogoutResponse](
			httpClient,
			baseURL+UserServiceLogoutProcedure,
			connect.WithSchema(userServiceLogoutMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	createUser            *connect.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	getUserList           *connect.Client[v1.GetUserListRequest, v1.GetUserListResponse]
	createWorkspaceMember *connect.Client[v1.CreateWorkspaceMemberRequest, v1.CreateWorkspaceMemberResponse]
	login                 *connect.Client[v1.LoginRequest, v1.LoginResponse]
	logout                *connect.Client[v1.LogoutRequest, v1.LogoutResponse]
}

// CreateUser calls outward.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// GetUserList calls outward.v1.UserService.GetUserList.
func (c *userServiceClient) GetUserList(ctx context.Context, req *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error) {
	return c.getUserList.CallUnary(ctx, req)
}

// CreateWorkspaceMember calls outward.v1.UserService.CreateWorkspaceMember.
func (c *userServiceClient) CreateWorkspaceMember(ctx context.Context, req *connect.Request[v1.CreateWorkspaceMemberRequest]) (*connect.Response[v1.CreateWorkspaceMemberResponse], error) {
	return c.createWorkspaceMember.CallUnary(ctx, req)
}

// Login calls outward.v1.UserService.Login.
func (c *userServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// Logout calls outward.v1.UserService.Logout.
func (c *userServiceClient) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return c.logout.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the outward.v1.UserService service.
type UserServiceHandler interface {
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error)
	GetUserList(context.Context, *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error)
	CreateWorkspaceMember(context.Context, *connect.Request[v1.CreateWorkspaceMemberRequest]) (*connect.Response[v1.CreateWorkspaceMemberResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceCreateUserHandler := connect.NewUnaryHandler(
		UserServiceCreateUserProcedure,
		svc.CreateUser,
		connect.WithSchema(userServiceCreateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetUserListHandler := connect.NewUnaryHandler(
		UserServiceGetUserListProcedure,
		svc.GetUserList,
		connect.WithSchema(userServiceGetUserListMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateWorkspaceMemberHandler := connect.NewUnaryHandler(
		UserServiceCreateWorkspaceMemberProcedure,
		svc.CreateWorkspaceMember,
		connect.WithSchema(userServiceCreateWorkspaceMemberMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceLoginHandler := connect.NewUnaryHandler(
		UserServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(userServiceLoginMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceLogoutHandler := connect.NewUnaryHandler(
		UserServiceLogoutProcedure,
		svc.Logout,
		connect.WithSchema(userServiceLogoutMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/outward.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceCreateUserProcedure:
			userServiceCreateUserHandler.ServeHTTP(w, r)
		case UserServiceGetUserListProcedure:
			userServiceGetUserListHandler.ServeHTTP(w, r)
		case UserServiceCreateWorkspaceMemberProcedure:
			userServiceCreateWorkspaceMemberHandler.ServeHTTP(w, r)
		case UserServiceLoginProcedure:
			userServiceLoginHandler.ServeHTTP(w, r)
		case UserServiceLogoutProcedure:
			userServiceLogoutHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.CreateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("outward.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUserList(context.Context, *connect.Request[v1.GetUserListRequest]) (*connect.Response[v1.GetUserListResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("outward.v1.UserService.GetUserList is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateWorkspaceMember(context.Context, *connect.Request[v1.CreateWorkspaceMemberRequest]) (*connect.Response[v1.CreateWorkspaceMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("outward.v1.UserService.CreateWorkspaceMember is not implemented"))
}

func (UnimplementedUserServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("outward.v1.UserService.Login is not implemented"))
}

func (UnimplementedUserServiceHandler) Logout(context.Context, *connect.Request[v1.LogoutRequest]) (*connect.Response[v1.LogoutResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("outward.v1.UserService.Logout is not implemented"))
}
