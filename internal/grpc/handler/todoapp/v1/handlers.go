package v1

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pb "github.com/todo-app/internal/protos/api/todoapp/v1"
	modelpb "github.com/todo-app/internal/protos/api/todoapp/v1/data"
	"google.golang.org/grpc"
)

type TodoAppServiceHandler struct {
	pb.UnimplementedTodoAppServiceServer

	todoHandler *TodoHandler
}

// NewTodoAppService creates TodoAppServiceCtrl instance.
func NewTodoAppService() *TodoAppServiceHandler {
	return &TodoAppServiceHandler{
		todoHandler: NewTodoHandler(),
	}
}

// RegisterService registers the GRPC service.
func (t *TodoAppServiceHandler) RegisterService(g *grpc.Server) {
	pb.RegisterTodoAppServiceServer(g, t)
}

// RegisterGateway starts the gateway (i.e. reverse proxy)
// that proxies HTTP requests to the appropriate gRPC endpoints.
func (t *TodoAppServiceHandler) RegisterGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return pb.RegisterTodoAppServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// ListTodos returns a todos set with the provided ids.
func (t *TodoAppServiceHandler) ListTodos(ctx context.Context, req *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	return t.todoHandler.ListTodos(ctx, req)
}

// CreateTodo allows to create a new todo.
func (t *TodoAppServiceHandler) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*modelpb.Todo, error) {
	return t.todoHandler.CreateTodo(ctx, req)
}

// GetTodo returns todo with the provided id.
func (t *TodoAppServiceHandler) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*modelpb.Todo, error) {
	return t.todoHandler.GetTodo(ctx, req)
}

// UpdateTodo allows to modify an existing todo.
func (t *TodoAppServiceHandler) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*modelpb.Todo, error) {
	return t.todoHandler.UpdateTodo(ctx, req)
}

// DeleteTodo allows to delete an existing todo.
func (t *TodoAppServiceHandler) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.Empty, error) {
	return t.todoHandler.DeleteTodo(ctx, req)
}
