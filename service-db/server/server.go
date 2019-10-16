package server

import (
	"net"

	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-db/rpc"

	"google.golang.org/grpc"
)

func Run() {
	server := grpc.NewServer()
	pb.RegisterRpcTestServer(server, rpc.Server)
	listener, err := net.Listen("tcp", "192.168.0.130:5000")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}
