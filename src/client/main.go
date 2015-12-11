package main

import (
	"fmt"
	"strconv"

	pb "github.com/edward-limitless/goRPC/src/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address    = "localhost:50051"
	default_id = "cookie1234"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect to server: %v", err)
	} else {
		fmt.Println("Client connected to " + address)
	}
	defer conn.Close()
	c := pb.NewBidderClient(conn)

	name := default_id
	r, err := c.Bid(context.Background(), &pb.BidRequest{UserIdentifier: &name})

	if err != nil {
		fmt.Errorf("Request broke, yo. Here's error: %v", err)
	}

	fmt.Printf("Bid? %s\n", strconv.FormatBool(*r.Bid))
}
