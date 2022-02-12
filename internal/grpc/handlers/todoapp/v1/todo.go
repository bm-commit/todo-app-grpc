package v1

import (
	"context"

	pb "github.com/todo-app/internal/protos/api/todoapp/v1"
	modelpb "github.com/todo-app/internal/protos/api/todoapp/v1/data"
)

// TodoHandler handles all todo request.
type TodoHandler struct {
	//ts *timingstore.TimingStore
}

// NewTodoHandler returns a circuit handler.
func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		//ts: ts,
	}
}

// GetTodo returns the requested todo.
func (t *TodoHandler) GetTodo(_ context.Context, req *pb.GetTodoRequest) (*modelpb.Todo, error) {
	return &modelpb.Todo{}, nil
}

// ListTodos return a list of todos with pagination
func (t *TodoHandler) ListTodos(_ context.Context, req *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	return &pb.ListTodosResponse{}, nil
}

// CreateTodo creates a todo
func (t *TodoHandler) CreateTodo(_ context.Context, req *pb.CreateTodoRequest) (*modelpb.Todo, error) {
	return &modelpb.Todo{}, nil
}

// UpdateTodo updates a todo
func (t *TodoHandler) UpdateTodo(_ context.Context, req *pb.UpdateTodoRequest) (*modelpb.Todo, error) {
	return &modelpb.Todo{}, nil
}

// DeleteTodo deletes a todo
func (t *TodoHandler) DeleteTodo(_ context.Context, req *pb.DeleteTodoRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
