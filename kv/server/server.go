package server

import (
	"context"
	"log/slog"

	pb "github.com/ayhonz/raft_kv_go/proto"
)

type Server struct {
	pb.KvServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) KvGet(ctx context.Context, _ *pb.None) (*pb.GetResponse, error) {
	slog.Info("Got request from client")
	return &pb.GetResponse{
		Msg: "Hello there buddy",
	}, nil
}
