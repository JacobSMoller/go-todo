package todo

import (
	"context"

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
	result := s.DB.Create(req.Item)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %v", result.Error)
	}
	return &todo.CreateTodoResponse{Id: req.Item.Id}, nil
}

func (s Service) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	result := s.DB.Create(&todo.Todo{Id: req.Id})
	if result.Error != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not delete the item: %v", result.Error)
	}
	return &todo.DeleteTodoResponse{}, nil
}

func (s Service) CreateOwner(ctx context.Context, req *todo.CreateOwnerRequest) (*todo.CreateOwnerResponse, error) {
	result := s.DB.Create(req.Owner)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not create owner: %v", result.Error)
	}
	return &todo.CreateOwnerResponse{Id: req.Owner.Id}, nil
}

// GetTodo retrieves a todo item from its ID
func (s Service) GetTodo(ctx context.Context, req *todo.GetTodoRequest) (*todo.GetTodoResponse, error) {
	var item todo.Todo
	result := s.DB.First(&item, req.Id)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", result.Error)
	}
	return &todo.GetTodoResponse{Item: &item}, nil
}
