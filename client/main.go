package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/JacobSMoller/go-todo/api/todo/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if os.Args[1] == strings.ToLower("create") {
		item := &pb.Todo{
			Title:       "My Foo task",
			Description: "A task created from my client",
		}
		req := &pb.CreateTodoRequest{Item: item}
		r, err := c.CreateTodo(ctx, req)
		if err != nil {
			log.Fatalf("Could not create item %v", err)
		}
		log.Printf("Created todo with id: %d", r.Id)
	}

	if os.Args[1] == strings.ToLower("delete") {
		del_req := &pb.DeleteTodoRequest{Id: 1}
		del_res, err := c.DeleteTodo(ctx, del_req)
		if err != nil {
			log.Fatalf("Could not create item %v", err)
		}
		log.Println(del_res)
		log.Printf("Deleted todo with id: %d", del_req.Id)
	}

	if os.Args[1] == strings.ToLower("select") {
		req := &pb.GetTodoRequest{Id: 2}
		r, err := c.GetTodo(ctx, req)
		if err != nil {
			log.Fatalf("Could not retrive item with id: %d %v", req.Id, err)
		}
		log.Printf("Found todo: %v", r.Item)
	}
}
