package main

import (
	// "context"
	// "fmt"
	// "net/http"
	// "os"
	// "path"
	// "runtime"
	"log"
	"net"
	"time"

	pb "github.com/JacobSMoller/go-todo/api/todo/v1"
	"github.com/JacobSMoller/go-todo/service/todo"
	"github.com/go-pg/pg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	// Listen on port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Connect to PostgreSQL
	db := pg.Connect(&pg.Options{
		User:                  "postgres",
		Password:              "mysecretpassword",
		Database:              "todo",
		Addr:                  "127.0.0.1:5432",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})
	db.CreateTable(&pb.Todo{}, nil)
	db.CreateTable(&pb.Owner{}, nil)
	// Setup a grpc server
	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server, &todo.Service{DB: db})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
