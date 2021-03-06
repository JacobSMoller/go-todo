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

	pb "github.com/JacobSMoller/go-todo/api/todo/v1"
	"github.com/JacobSMoller/go-todo/service/todo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	//Connect to database
	db, err := gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=todo password=mysecretpassword sslmode=disable")
	//close db at end of this function
	defer db.Close()
	// Setup a grpc server
	server := grpc.NewServer()
	pb.RegisterTodoServiceServer(server, &todo.Service{DB: db})
	// Register reflection service on gRPC server.
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
