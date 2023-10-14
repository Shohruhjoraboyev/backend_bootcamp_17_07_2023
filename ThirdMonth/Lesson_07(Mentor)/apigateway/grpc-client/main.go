package main

import (
	branch "apigateway/genproto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := branch.NewBranchServiceClient(conn)

	CreateBranch(c)
	GetBranch(c)
	UpdateBranch(c)
	DeletetBranch(c)
	GetAllBranch(c)

}
func CreateBranch(c branch.BranchServiceClient) {
	name := ""
	address := ""
	fmt.Print("Enter name: ")
	fmt.Scan(&name)
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	req := &branch.CreateBranch{
		Name:    name,
		Address: address,
	}

	r, err := c.Create(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to create branch: %v", err)
	}
	log.Printf("BranchInfo: %s", r)
}

func GetBranch(c branch.BranchServiceClient) {
	BranchID := ""
	fmt.Print("Enter Branch ID To GetBranch: ")
	fmt.Scan(&BranchID)

	req := &branch.IdRequest{
		Id: BranchID,
	}

	resp, err := c.GetBranch(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get branch: %v", err)
	}

	// Print the branch information
	log.Printf("\nBranchInfo:\n%s\n", resp.String())
}
func GetAllBranch(c branch.BranchServiceClient) {
	limit := 0
	fmt.Println("Enter Limit: ")
	fmt.Scan(&limit)
	page := 0
	fmt.Print("Enter page: ")
	fmt.Scan(&page)
	search := ""
	fmt.Print("Enter branch name to search: ")
	fmt.Scan(&search)

	r, err := c.GetAllBranch(context.Background(), &branch.GetAllBranchRequest{
		Limit:  int32(limit),
		Page:   int32(page),
		Search: search,
	})
	if err != nil {
		log.Fatalf("Couldnt find branches: %v", err)
	}
	log.Printf("All Branches: %v", r)
}

func UpdateBranch(c branch.BranchServiceClient) {
	var id, name, address string
	fmt.Print("Enter Branch ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter updated name: ")
	fmt.Scan(&name)
	fmt.Print("Enter updated address: ")
	fmt.Scan(&address)

	// Create the request message
	req := &branch.UpdateBranchRequest{
		Id:      id,
		Name:    name,
		Address: address,
	}

	// Send the request to the server and receive the response
	resp, err := c.UpdateBranch(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to update branch: %v", err)
	}

	// Print the response
	log.Printf("Updated Branch ID: %s", resp.Id)
}

func DeletetBranch(c branch.BranchServiceClient) {
	BranchID := ""
	fmt.Print("Enter Branch ID To Delete: ")
	fmt.Scan(&BranchID)

	req := &branch.IdRequest{
		Id: BranchID,
	}
	resp, err := c.DeleteBranch(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to Delete branch: %v", err)
	}

	// Print the branch information
	log.Printf("Deleted Branch ID: %s", resp.Id)
}
