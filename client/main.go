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
	// future_time, err := time.Parse(time.RFC3339, "2020-01-10T10:44:00+00:00")
	// if err != nil {
	// 	log.Fatalf("Failed to parse future time %v", err)
	// }
	// proto_future_time, err := ptypes.TimestampProto(future_time)
	// item := &pb.Todo{
	// 	Title:       "Foo task",
	// 	Description: "A task created from my client",
	// 	CreatedAt:   proto_future_time,
	// }
	// req := &pb.CreateTodoRequest{Item: item}
	// r, err := c.CreateTodo(ctx, req)
	// if err != nil {
	// 	log.Fatalf("Could not create item %v", err)
	// }
	// log.Printf("Created todo with id: %d", r.Id)

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

	req := &pb.GetTodoRequest{Id: 1}

}
