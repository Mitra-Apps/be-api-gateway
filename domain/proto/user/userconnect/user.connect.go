// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/user/user.proto

package userconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	user "github.com/Mitra-Apps/be-user-service/domain/proto/user"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "proto.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetUsersProcedure is the fully-qualified name of the UserService's GetUsers RPC.
	UserServiceGetUsersProcedure = "/proto.UserService/GetUsers"
)

// UserServiceClient is a client for the proto.UserService service.
type UserServiceClient interface {
	GetUsers(context.Context, *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error)
}

// NewUserServiceClient constructs a client for the proto.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUsers: connect.NewClient[user.GetUsersRequest, user.GetUsersResponse](
			httpClient,
			baseURL+UserServiceGetUsersProcedure,
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUsers *connect.Client[user.GetUsersRequest, user.GetUsersResponse]
}

// GetUsers calls proto.UserService.GetUsers.
func (c *userServiceClient) GetUsers(ctx context.Context, req *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error) {
	return c.getUsers.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the proto.UserService service.
type UserServiceHandler interface {
	GetUsers(context.Context, *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceGetUsersHandler := connect.NewUnaryHandler(
		UserServiceGetUsersProcedure,
		svc.GetUsers,
		opts...,
	)
	return "/proto.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetUsersProcedure:
			userServiceGetUsersHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUsers(context.Context, *connect.Request[user.GetUsersRequest]) (*connect.Response[user.GetUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.UserService.GetUsers is not implemented"))
}
