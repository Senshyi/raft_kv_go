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

func (s *Server) Get(ctx context.Context, _ *pb.GetRequest) (*pb.GetResponse, error) {
	slog.Info("Got request from client")
	return &pb.GetResponse{
		Kv: &pb.KvPair{
			Key:   "1",
			Value: "I am value",
		},
	}, nil
}

func (s *Server) Set(ctx context.Context, request *pb.PutRequest) (*pb.PutResponse, error) {
	slog.Info("Got request from client")
	slog.Info("incomming set request %+v", request)
	return &pb.PutResponse{
		Key:   "2",
		Value: "value",
	}, nil
}

func (s *Server) Scan(ctx context.Context, _ *pb.ScanRequest) (*pb.ScanResponse, error) {
	slog.Info("Got request from client")
	return &pb.ScanResponse{
		Kvs: []*pb.KvPair{{Key: "1", Value: "I'm value"}},
	}, nil
}
