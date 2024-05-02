package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "asg4/protofile"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	s.users = append(s.users, in)
	return in, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	for _, user := range s.users {
		if user.Id == in.Id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *server) ListUsers(ctx context.Context, in *pb.Empty) (*pb.UserList, error) {
	return &pb.UserList{Users: s.users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
