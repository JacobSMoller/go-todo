package todo

import (
	"context"

	todo "github.com/JacobSMoller/go-todo/api/todo/v1"
	"github.com/go-pg/pg"
	"github.com/gogo/protobuf/types"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Service is the service dealing with storing
// and retrieving todo items from the database.
type Service struct {
	DB *pg.DB
}

// CreateTodo creates a todo given a description
func (s Service) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	req.Item.Id = uuid.NewV4().String()
	err := s.DB.Insert(req.Item)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}
	return &todo.CreateTodoResponse{Id: req.Item.Id}, nil
}
