package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNum = "9001"

func main() {
	lis, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		log.Panicf("cannot listen")
	}

	grpcServer := grpc.NewServer()

}
