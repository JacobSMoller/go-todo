package main

import (
	"context"
	"log"
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
	item := &pb.Todo{
		Title:       "Foo task",
		Description: "A task created from my client",
	}
	req := &pb.CreateTodoRequest{Item: item}
	r, err := c.CreateTodo(ctx, req)
	if err != nil {
		log.Fatalf("Could not create item %v", err)
	}
	log.Printf("Created todo with id: %d", r.Id)

	// del_req := &pb.DeleteTodoRequest{Id: 4}
	// del_res, err := c.DeleteTodo(ctx, del_req)
	// if err != nil {
	// 	log.Fatalf("Could not create item %v", err)
	// }
	// log.Println(del_res)
	// log.Printf("Deleted todo with id: %d", del_req.Id)

	// owner := &pb.Owner{
	// 	Firstname: "John",
	// 	Lastname:  "Doe",
	// }
	// req := &pb.CreateOwnerRequest{Owner: owner}
	// r, err := c.CreateOwner(ctx, req)
	// if err != nil {
	// 	log.Fatalf("Could not create item %v", err)
	// }
	// log.Printf("Created owner with id: %d", r.Id)
}
