package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	branches "server/genproto"
	pb "server/genproto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBranchServiceServer
	pb.UnimplementedStreamServiceServer
	branches []*pb.Branch
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBranchServiceServer(s, &server{})
	pb.RegisterStreamServiceServer(s, &server{})
	log.Printf("Server listening at: %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// func (s *server) Count(req *pb.Request, res pb.StreamService_CountServer) error {
// 	fmt.Println("request:", req.GetNumber())
// 	for i := 1; i < int(req.GetNumber()); i++ {
// 		err := res.Send(&pb.Response{Count: int32(i)})
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		fmt.Println("res sent:", i)
// 		time.Sleep(time.Second)
// 	}
// 	return nil
// }
// func (s *server) Sum(stream pb.StreamService_SumServer) error {
// 	var total int32

// 	for {
// 		value, err := stream.Recv()
// 		if err == io.EOF {
// 			fmt.Println("respond:", total)
// 			return stream.SendAndClose(&pb.Response{
// 				Count: total,
// 			})
// 		}

// 		if err != nil {
// 			return err
// 		}

// 		fmt.Println("received number:", value.GetNumber())
// 		total += value.GetNumber()
// 	}
// }

// func (s *server) Sqr(stream pb.StreamService_SqrServer) error {
// 	for {
// 		value, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println("received number:", value.GetNumber())

// 		if err := stream.Send(&pb.Response{
// 			Count: int32(math.Pow(float64(value.GetNumber()), 2)),
// 		}); err != nil {
// 			return err
// 		}

// 	}

// }

func (s *server) Fibonacci(req *pb.Request, res pb.StreamService_FibonacciServer) error {
	fmt.Println("request:", req.GetNumber())

	a, b := 0, 1
	for i := 0; a < int(req.GetNumber()); i++ {

		if err := res.Send(&pb.Response{Count: int32(a)}); err != nil {
			return err
		}

		fmt.Println("sent: ", a)
		a, b = b, a+b
		time.Sleep(time.Second)
	}
	return nil
}

func (s *server) Translate(stream pb.StreamService_TranslateServer) error {

	wordsMap := map[string]string{
		"yellow": "sariq",
		"blue":   "moviy",
		"red":    "qizil",
		"green":  "yashil",
		"white":  "oq",
		"black":  "qora",
		"brown":  "kulrang",
		"pink":   "pushti",
	}
	for _, word := range wordsMap {
		value, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("received word:", value.GetWord())

		if word == value.GetWord() {
			if err := stream.Send(&pb.RespondWords{
				Word: word,
			}); err != nil {
				return err
			}
		}
	}
	return nil
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
func (s *server) GetAllBranch(ctx context.Context, req *pb.GetAllBranchRequest) (*pb.GetAllBranchResponse, error) {
	log.Printf("Received: %v", req)

	page := req.Page
	limit := req.Limit
	searchQ := req.Search

	filteredBranches := make([]*branches.Branch, 0)
	for _, branch := range s.branches {
		if searchQ == "" || strings.Contains(strings.ToLower(branch.Name), strings.ToLower(searchQ)) {
			filteredBranches = append(filteredBranches, branch)
		}

	}

	totalBranches := int32(len(filteredBranches))

	if limit == 0 {
		limit = int32(len(filteredBranches))
	}

	if page == 0 {
		page = 1
	}
	startIndex := (page - 1) * limit
	endIndex := page * limit

	if startIndex >= int32(len(filteredBranches)) {
		return &branches.GetAllBranchResponse{
			Branches: nil,
			Count:    totalBranches,
		}, nil
	}

	if endIndex > int32(len(filteredBranches)) {
		endIndex = int32(len(filteredBranches))
	}
	// Get the branches within the specified range
	pagedBranches := filteredBranches[startIndex:endIndex]

	return &pb.GetAllBranchResponse{
		Branches: pagedBranches,
		Count:    totalBranches,
	}, nil

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
