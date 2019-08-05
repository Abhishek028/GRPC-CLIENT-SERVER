package main

import (
	"context"
	"time"
	pb "proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewGreetingsClient(conn)

	var s string = "Hello"
	for i:=1;i<=100;i++ {
		req := &pb.Request{S: s}
		if _, err := client.SayHello(context.Background(), req); err == nil {
			
		}
		time.Sleep(5*time.Second)
	}

	
}