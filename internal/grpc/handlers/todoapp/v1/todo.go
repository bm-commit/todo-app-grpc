package v1

import (
	"context"
	"errors"

	uuid "github.com/satori/go.uuid"
	pb "github.com/todo-app/internal/protos/api/todoapp/v1"
	modelpb "github.com/todo-app/internal/protos/api/todoapp/v1/data"
	storage "github.com/todo-app/internal/storage/sql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	todoModel, err := t.s.GetTodo(req.GetId())
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "Todo not found")
		}

		return nil, handleError(err)
	}

	return t.toTodoMsg(todoModel), nil
}

// ListTodos return a list of todos with pagination
func (t *TodoHandler) ListTodos(_ context.Context, req *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	list, err := t.s.ListTodos(req.GetIds())
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "Todos not found")
		}
		return nil, handleError(err)
	}

	res := make([]*modelpb.Todo, len(list))
	for i, item := range list {
		res[i] = t.toTodoMsg(item)
	}

	return &pb.ListTodosResponse{
		Todos:         res,
		NextPageToken: "",
	}, nil
}

// CreateTodo creates a todo
func (t *TodoHandler) CreateTodo(_ context.Context, req *pb.CreateTodoRequest) (*modelpb.Todo, error) {
	newTodo := t.toTodoStore(uuid.NewV1().String(), req.GetTodo())
	if err := t.s.CreateEntry(newTodo); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return t.toTodoMsg(newTodo), nil
}

// UpdateTodo updates a todo
func (t *TodoHandler) UpdateTodo(_ context.Context, req *pb.UpdateTodoRequest) (*modelpb.Todo, error) {
	if _, err := t.s.GetTodo(req.GetId()); err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "Todo not found")
		}
		return nil, handleError(err)
	}

	updated := t.toTodoStore(req.GetId(), req.GetTodo())
	if err := t.s.UpdateEntryWithNulls(updated); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return t.toTodoMsg(updated), nil
}

// DeleteTodo deletes a todo
func (t *TodoHandler) DeleteTodo(_ context.Context, req *pb.DeleteTodoRequest) (*pb.Empty, error) {
	todoModel, err := t.s.GetTodo(req.GetId())
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "todo not found")
		}
		return nil, handleError(err)
	}
	if err := t.s.DeleteEntry(todoModel); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Empty{}, nil
}

func (t *TodoHandler) toTodoStore(id string, req *modelpb.Todo) *storage.Todo {
	return &storage.Todo{
		Common:      storage.Common{ID: storage.NewUUID(id)},
		Description: req.GetDescription(),
		Completed:   req.GetCompleted(),
	}
}

func (t *TodoHandler) toTodoMsg(todo *storage.Todo) *modelpb.Todo {
	return &modelpb.Todo{
		Id:          todo.ID.String(),
		Description: todo.Description,
		Completed:   todo.Completed,
	}
}
