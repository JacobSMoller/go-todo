package todo

import (
	"context"
	"time"

	todo "github.com/JacobSMoller/go-todo/api/todo/v1"
	"github.com/jinzhu/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Service is the service dealing with storing
// and retrieving todo items from the database.
type Service struct {
	DB *gorm.DB
}

// CreateTodo creates a todo given a description
func (s Service) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	epoch := time.Now().Unix()
	req.Item.CreatedAt = epoch
	err := s.DB.Create(req.Item)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}
	return &todo.CreateTodoResponse{Id: req.Item.Id}, nil
}

func (s Service) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	err := s.DB.Create(&todo.Todo{Id: req.Id})
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not delete the item: %s", err)
	}
	return &todo.DeleteTodoResponse{}, nil
}

func (s Service) CreateOwner(ctx context.Context, req *todo.CreateOwnerRequest) (*todo.CreateOwnerResponse, error) {
	err := s.DB.Create(req.Owner)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not create owner: %s", err)
	}
	return &todo.CreateOwnerResponse{Id: req.Owner.Id}, nil
}
