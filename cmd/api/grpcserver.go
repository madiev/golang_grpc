package api

import (
	"myapp/myapp"
	"context"
)

type GRPCServer struct {
	myapp.UnimplementedMyAppServer
}

func (s *GRPCServer) Sum(ctx context.Context, req *myapp.AddRequest) (*myapp.AddResponse, error) {
	return &myapp.AddResponse{Result: req.GetX() + req.GetY()}, nil
}