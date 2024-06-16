package main

import (
	"flag"
	"log"
	"log/slog"
	"net"

	"github.com/ayhonz/raft_kv_go/kv/server"
	pb "github.com/ayhonz/raft_kv_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr = flag.String("addr", ":8080", "server port")

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	grpcServer := grpc.NewServer()
	server := server.NewServer()

	pb.RegisterKvServer(grpcServer, server)
	reflection.Register(grpcServer)

	lisner, err := net.Listen("tcp", *addr)
	if err != nil {
		slog.Error("Failed to start tcp server: %v", err)
	}

	err = grpcServer.Serve(lisner)
	if err != nil {
		slog.Error("Failed to start grpc server: %v", err)
	}
}
