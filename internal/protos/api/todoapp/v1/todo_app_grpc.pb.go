// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/todoapp/v1/todo_app.proto

package todoappv1

import (
	context "context"
	data "github.com/todo-app/internal/protos/api/todoapp/v1/data"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodoAppServiceClient is the client API for TodoAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoAppServiceClient interface {
	// Returns todo with the provided ids.
	ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*ListTodosResponse, error)
	// Create a todo
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*data.Todo, error)
	// Returns a todo by ID
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*data.Todo, error)
	// Update an existing todo
	UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*data.Todo, error)
	// Delete an existing todo
	DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*Empty, error)
}

type todoAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoAppServiceClient(cc grpc.ClientConnInterface) TodoAppServiceClient {
	return &todoAppServiceClient{cc}
}

func (c *todoAppServiceClient) ListTodos(ctx context.Context, in *ListTodosRequest, opts ...grpc.CallOption) (*ListTodosResponse, error) {
	out := new(ListTodosResponse)
	err := c.cc.Invoke(ctx, "/api.todoapp.v1.TodoAppService/ListTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*data.Todo, error) {
	out := new(data.Todo)
	err := c.cc.Invoke(ctx, "/api.todoapp.v1.TodoAppService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*data.Todo, error) {
	out := new(data.Todo)
	err := c.cc.Invoke(ctx, "/api.todoapp.v1.TodoAppService/GetTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceClient) UpdateTodo(ctx context.Context, in *UpdateTodoRequest, opts ...grpc.CallOption) (*data.Todo, error) {
	out := new(data.Todo)
	err := c.cc.Invoke(ctx, "/api.todoapp.v1.TodoAppService/UpdateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoAppServiceClient) DeleteTodo(ctx context.Context, in *DeleteTodoRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.todoapp.v1.TodoAppService/DeleteTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoAppServiceServer is the server API for TodoAppService service.
// All implementations must embed UnimplementedTodoAppServiceServer
// for forward compatibility
type TodoAppServiceServer interface {
	// Returns todo with the provided ids.
	ListTodos(context.Context, *ListTodosRequest) (*ListTodosResponse, error)
	// Create a todo
	CreateTodo(context.Context, *CreateTodoRequest) (*data.Todo, error)
	// Returns a todo by ID
	GetTodo(context.Context, *GetTodoRequest) (*data.Todo, error)
	// Update an existing todo
	UpdateTodo(context.Context, *UpdateTodoRequest) (*data.Todo, error)
	// Delete an existing todo
	DeleteTodo(context.Context, *DeleteTodoRequest) (*Empty, error)
	mustEmbedUnimplementedTodoAppServiceServer()
}

// UnimplementedTodoAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoAppServiceServer struct {
}

func (UnimplementedTodoAppServiceServer) ListTodos(context.Context, *ListTodosRequest) (*ListTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
func (UnimplementedTodoAppServiceServer) CreateTodo(context.Context, *CreateTodoRequest) (*data.Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoAppServiceServer) GetTodo(context.Context, *GetTodoRequest) (*data.Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}
func (UnimplementedTodoAppServiceServer) UpdateTodo(context.Context, *UpdateTodoRequest) (*data.Todo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (UnimplementedTodoAppServiceServer) DeleteTodo(context.Context, *DeleteTodoRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (UnimplementedTodoAppServiceServer) mustEmbedUnimplementedTodoAppServiceServer() {}

// UnsafeTodoAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoAppServiceServer will
// result in compilation errors.
type UnsafeTodoAppServiceServer interface {
	mustEmbedUnimplementedTodoAppServiceServer()
}

func RegisterTodoAppServiceServer(s grpc.ServiceRegistrar, srv TodoAppServiceServer) {
	s.RegisterService(&TodoAppService_ServiceDesc, srv)
}

func _TodoAppService_ListTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceServer).ListTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todoapp.v1.TodoAppService/ListTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceServer).ListTodos(ctx, req.(*ListTodosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todoapp.v1.TodoAppService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todoapp.v1.TodoAppService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppService_UpdateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceServer).UpdateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todoapp.v1.TodoAppService/UpdateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceServer).UpdateTodo(ctx, req.(*UpdateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoAppService_DeleteTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoAppServiceServer).DeleteTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.todoapp.v1.TodoAppService/DeleteTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoAppServiceServer).DeleteTodo(ctx, req.(*DeleteTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoAppService_ServiceDesc is the grpc.ServiceDesc for TodoAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.todoapp.v1.TodoAppService",
	HandlerType: (*TodoAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTodos",
			Handler:    _TodoAppService_ListTodos_Handler,
		},
		{
			MethodName: "CreateTodo",
			Handler:    _TodoAppService_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoAppService_GetTodo_Handler,
		},
		{
			MethodName: "UpdateTodo",
			Handler:    _TodoAppService_UpdateTodo_Handler,
		},
		{
			MethodName: "DeleteTodo",
			Handler:    _TodoAppService_DeleteTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/todoapp/v1/todo_app.proto",
}
