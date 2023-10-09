package main

import (
	"context"
	sale_service "example-grpc-client/genproto"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// c := sale_service.NewBranchServiceClient(conn)
	// Contact the server and print out its response.
	// for {
	// 	name := ""
	// 	address := ""
	// 	fmt.Scan(&name)
	// 	fmt.Scan(&address)
	// 	r, err := c.Create(context.Background(), &branch.CreateBranch{Name: name, Address: address})
	// 	if err != nil {
	// 		log.Fatalf("could not greet: %v", err)
	// 	}
	// 	log.Printf("new branch: %v", r)
	// }
	streamService := sale_service.NewStreamServiceClient(conn)
	//  server side streaming
	// resStream, err := streamService.Count(context.Background(), &sale_service.Request{Number: 12})
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// for {
	// 	resp, err := resStream.Recv()
	// 	if err == io.EOF { //End Of File
	// 		return
	// 	}
	// 	if err != nil {
	// 		log.Fatalln(err.Error())
	// 	}
	// 	fmt.Println("resp received: ", resp.GetCount())
	// }

	// ***** client side streaming
	// stream, err := streamService.Sum(context.Background())
	// if err != nil {
	// 	log.Fatalln("Consuming stream", err)
	// }

	// for i := 0; i < 10; i++ {
	// 	err := stream.Send(&sale_service.Request{Number: int32(i)})
	// 	if err != nil {
	// 		log.Fatalln("Sending value", err)
	// 	}
	// 	fmt.Println("sent:", i)
	// 	time.Sleep(time.Second)
	// }

	// res, err := stream.CloseAndRecv()
	// if err != nil {
	// 	log.Fatalln("Closing", err)
	// }
	// fmt.Println("Total sum", res.GetCount())

	// bidirectional streaming
	// stream, err := streamService.Sqr(context.Background())
	// if err != nil {
	// 	log.Fatalln("Opening stream", err)
	// }

	// for i := 0; i < 10; i++ {
	// 	err := stream.Send(&sale_service.Request{Number: int32(i)})
	// 	if err != nil {
	// 		log.Fatalln("Sending value", err)
	// 	}
	// 	fmt.Println("send:", i)

	// }
	// if err := stream.CloseSend(); err != nil {
	// 	log.Fatalln("CloseSend", err)
	// }
	// for {
	// 	res, err := stream.Recv()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatalln("Closing", err)
	// 	}
	// 	fmt.Println("Received:", res.GetCount())
	// 	time.Sleep(time.Second)
	// }

	// stream, err := streamService.Sqr(context.Background())
	// if err != nil {
	// 	log.Fatalln("Opening stream", err)
	// }

	// fibonacci streaming
	stream, err := streamService.Fibonacci(context.Background())
	num := 20
	err = stream.Send(&sale_service.Request{Number: int32(num)})
	if err != nil {
		log.Fatalln("Sending value", err)
	}
	fmt.Println("sent:", 20)

	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Closing", err)
		}
		fmt.Println("Received:", res.GetCount())
		time.Sleep(time.Second)
	}

}
