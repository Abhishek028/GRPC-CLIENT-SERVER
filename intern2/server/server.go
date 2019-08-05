package main

import (
	"context"
	"net"
	"fmt"
	pb "proto"
	"google.golang.org/grpc"
)

type server struct{}
var count int64 = 1
func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingsServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (server) SayHello(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	s := request.GetS() 
	fmt.Printf("%s%d\n",s,count)
	count++
	return &pb.Response{}, nil
	
}