package main

import (
	"context"
	"log"

	pb "asg4/protofile"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	user, err := c.AddUser(context.Background(), &pb.User{
		Id:    210109015,
		Name:  "John",
		Email: "john.doe@example.com",
	})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("User added successfully: %v", user)

	getUser, err := c.GetUser(context.Background(), &pb.UserId{Id: 1})
	if err != nil {
		log.Fatalf("failed to retrieve user: %v", err)
	}
	log.Printf("User retrieved successfully: %v", getUser)

	listUsers, err := c.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("failed to list users: %v", err)
	}
	log.Printf("Users listed successfully: %v", listUsers.GetUsers())
}
