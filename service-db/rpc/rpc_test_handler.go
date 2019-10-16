package rpc

import (
	"context"

	"github.com/Irfish/fantasy.server/pb"
)

func (s *RpcServer) Test(context.Context, *pb.RpcTestReq) (*pb.RpcTestRsp, error) {
	return &pb.RpcTestRsp{}, nil
}
