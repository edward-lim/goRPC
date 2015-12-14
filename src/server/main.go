package main

import (
	"fmt"
	"net"

	pb "github.com/edward-limitless/goRPC/src/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Bid(ctx context.Context, in *pb.BidRequest) (*pb.BidResponse, error) {
	fmt.Printf("Request came in for: %s\n", in.UserIdentifier)
	fmt.Println("Received request")
	return &pb.BidResponse{Bid: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Received error listening to BidRequests: %v", err)
	} else {
		fmt.Println("Server up and running, listening on port " + port)
	}

	s := grpc.NewServer()
	pb.RegisterBidderServer(s, &server{})
	s.Serve(lis)
}
