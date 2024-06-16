package main

import (
	"log"
	"net"

	pb "github.com/ayhonz/raft_kv_go/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type KvServer struct {
	pb.KvServer
}

func (s *KvServer) KvGet(ctx context.Context, _ *pb.None) (*pb.GetResponse, error) {
	log.Printf("Got request from client")
	return &pb.GetResponse{
		Msg: "Hello there buddy",
	}, nil
}

func main() {
	lisner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to start tcp server: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterKvServer(grpcServer, &KvServer{})
	reflection.Register(grpcServer)

	err = grpcServer.Serve(lisner)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
}
