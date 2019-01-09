package todo

import (
	"context"
	"time"

	todo "github.com/JacobSMoller/go-todo/api/todo/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Service is the service dealing with storing
// and retrieving todo items from the database.
type Service struct {
	DB *gorm.DB
}

type DbItem struct {
	Id          uint64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateTodo creates a todo given a description
func (s Service) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest) (*todo.CreateTodoResponse, error) {
	// Convert proto timestamps to time.Time TODO(JS): Figure out how to handle cases where they are not present.
	// db_created_at, err := ptypes.Timestamp(req.Item.CreatedAt)
	// log.Print(err)
	// if err != nil {
	// 	return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database, failed to convert proto timestamp to Time: %v", err)
	// }
	// db_updated_at, err := ptypes.Timestamp(req.Item.UpdatedAt)
	// if err != nil {
	// 	return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database, failed to convert proto timestamp to Time: %v", err)
	// }

	dbitem := DbItem{
		Title:       req.Item.Title,
		Description: req.Item.Description,
	}
	result := s.DB.Table("todo").Create(&dbitem)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %v", result.Error)
	}
	return &todo.CreateTodoResponse{Id: dbitem.Id}, nil
}

func (s Service) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	result := s.DB.Delete(&todo.Todo{Id: req.Id})
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
	var dbitem DbItem
	result := s.DB.Table("todo").First(&dbitem, req.Id)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", result.Error)
	}

	// Convert time.Time to protobuf timestamps
	created_at, err := ptypes.TimestampProto(dbitem.CreatedAt)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not retrieve item from the database failed to convert CreatedAt to proto timestmap : %v", err)
	}
	updated_at, err := ptypes.TimestampProto(dbitem.CreatedAt)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not retrieve item from the database failed to convert UpdatedAt to proto timestmap : %v", err)
	}

	item := todo.Todo{
		Id:          dbitem.Id,
		Title:       dbitem.Title,
		Description: dbitem.Description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}
	return &todo.GetTodoResponse{Item: &item}, nil
}

// GetTodos retrieves all todo items.
func (s Service) GetTodos(ctx context.Context, req *todo.GetTodosRequest) (*todo.GetTodosResponse, error) {
	var dbitems []*DbItem
	result := s.DB.Table("todo").Find(&dbitems)
	if result.Error != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", result.Error)
	}
	var items []*todo.Todo
	for _, dbitem := range dbitems {
		created_at, err := ptypes.TimestampProto(dbitem.CreatedAt)
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "Could not retrieve item from the database failed to convert CreatedAt to proto timestmap : %v", err)
		}
		updated_at, err := ptypes.TimestampProto(dbitem.CreatedAt)
		if err != nil {
			return nil, grpc.Errorf(codes.Internal, "Could not retrieve item from the database failed to convert UpdatedAt to proto timestmap : %v", err)
		}
		item := todo.Todo{
			Id:          dbitem.Id,
			Title:       dbitem.Title,
			Description: dbitem.Description,
			CreatedAt:   created_at,
			UpdatedAt:   updated_at,
		}
		items = append(items, &item)
	}
	return &todo.GetTodosResponse{Items: items}, nil
}
