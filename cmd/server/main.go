package main

import (
	"myapp/cmd/api"
	"myapp/myapp"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	s := grpc.NewServer()
	srv := &api.GRPCServer{}
	myapp.RegisterMyAppServer(s, srv)

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}