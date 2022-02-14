package v1

import (
	"context"

	pb "github.com/todo-app/internal/protos/api/todoapp/v1"
	modelpb "github.com/todo-app/internal/protos/api/todoapp/v1/data"
	storage "github.com/todo-app/internal/storage/sql"
)

// TodoHandler handles all todo request.
type TodoHandler struct {
	s *storage.SQLStorage
}

// NewTodoHandler returns a circuit handler.
func NewTodoHandler(store *storage.SQLStorage) *TodoHandler {
	return &TodoHandler{
		s: store,
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
