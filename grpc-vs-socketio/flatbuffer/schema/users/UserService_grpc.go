//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: user

package users

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for UserService service
type UserServiceClient interface{
  Add(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* AddResponse, error)  
}

type userServiceClient struct {
  cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
  return &userServiceClient{cc}
}

func (c *userServiceClient) Add(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* AddResponse, error) {
  out := new(AddResponse)
  err := grpc.Invoke(ctx, "/users.UserService/Add", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for UserService service
type UserServiceServer interface {
  Add(context.Context, *AddRequest) (*flatbuffers.Builder, error)  
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
  s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Add_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(AddRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(UserServiceServer).Add(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/users.UserService/Add",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(UserServiceServer).Add(ctx, req.(* AddRequest))
  }
  return interceptor(ctx, in, info, handler)
}


var _UserService_serviceDesc = grpc.ServiceDesc{
  ServiceName: "users.UserService",
  HandlerType: (*UserServiceServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Add",
      Handler: _UserService_Add_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}
