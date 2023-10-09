package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "server/genproto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBranchServiceServer
	branches []*pb.Branch
}

func (s *server) Create(ctx context.Context, req *pb.CreateBranch) (*pb.IdResponse, error) {
	log.Printf("Received: %v", req)
	id := uuid.NewString()
	s.branches = append(s.branches, &pb.Branch{
		Id:      id,
		Name:    req.Name,
		Address: req.Address,
	})
	return &pb.IdResponse{Id: id}, nil
}

func (s *server) GetBranch(ctx context.Context, req *pb.IdRequest) (*pb.Branch, error) {
	log.Printf("Received: %v", req)

	for _, br := range s.branches {
		if br.GetId() == req.GetId() {
			return br, nil
		}
	}

	return nil, fmt.Errorf("Branch not found")
}

func (s *server) UpdateBranch(ctx context.Context, req *pb.UpdateBranchRequest) (*pb.IdResponse, error) {
	log.Printf("Received: %v", req)
	for _, br := range s.branches {
		if br.GetId() == req.GetId() {
			br.Name = req.Name
			br.Address = req.Address
			return &pb.IdResponse{Id: br.Id}, nil
		}
	}
	return nil, fmt.Errorf("Branch cannot found")
}

func (s *server) DeleteBranch(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	log.Printf("Received: %v", req)

	for i, br := range s.branches {
		if br.GetId() == req.GetId() {
			// Remove the branch from the slice
			s.branches = append(s.branches[:i], s.branches[i+1:]...)
			return &pb.IdResponse{Id: req.GetId()}, nil
		}
	}

	return nil, fmt.Errorf("Branch not found")
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBranchServiceServer(s, &server{})

	log.Printf("Server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
