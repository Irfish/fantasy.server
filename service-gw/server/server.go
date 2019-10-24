package server

import (
	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/base"
	"github.com/Irfish/fantasy.server/service-gw/msg"
	"github.com/golang/protobuf/proto"
)

var (
	skeleton = base.NewSkeleton()
	ChanRpc  = skeleton.ChanRPCServer
	Module   = new(Server)
)

type Server struct {
	*gate.Gate
	*etcd3.ServiceNode
	SkeletonCloseSig chan bool
}

func (s *Server) OnInit() {
	RedisParserInit()
	s.Gate = &gate.Gate{
		MaxConnNum:      base.Server.MaxConnNum,
		PendingWriteNum: base.PendingWriteNum,
		MaxMsgLen:       base.MaxMsgLen,
		WSAddr:          base.Server.WSAddr,
		HTTPTimeout:     base.HTTPTimeout,
		CertFile:        base.Server.CertFile,
		KeyFile:         base.Server.KeyFile,
		TCPAddr:         base.Server.TCPAddr,
		LenMsgLen:       base.LenMsgLen,
		LittleEndian:    true,
		Processor:       msg.Processor,
		AgentChanRPC:    ChanRpc,
	}
	s.ServiceNode = new(etcd3.ServiceNode)

	//启动leaf的skeleton  处理通讯
	s.SkeletonCloseSig = make(chan bool, 1)
	go func() {
		skeleton.Run(s.SkeletonCloseSig)
	}()

	//注册服务
	s.SetNodeData(&etcd3.Node{Key: pb.GetGwKey(pb.SERVICE_GW), Name: pb.GetName(pb.SERVICE_GW), Address: base.Server.TCPAddr, Version: "v1.0"})
	etcd3.RegisterNode(s)

	//监控service下的node
	etcd3.WatchNode(ServiceNode)

	//第一次启动尝试连接所有服务
	nodes := etcd3.GetNodes(pb.ServicePrefix())
	for _, n := range nodes {
		cli := CreateServiceClient(n.Name, n.Address, false)
		ServiceNode.ServiceToClient[n.Key] = cli
	}

}

func (s *Server) BeforeDestroy() {
	s.SkeletonCloseSig <- true
}

func (s *Server) GetNodePrefix() string {
	return pb.GwPrefix()
}

func (s *Server) OnNodeRegister(k string, v interface{}) {}

func (s *Server) OnNodeUnregister(k string) {}

func sendMessage(a gate.Agent, m interface{}) {
	m1 := m.(proto.Message)
	body, err := msg.Processor.Marshal(m1)
	if err != nil {
		log.Error("SendToService proto.Marshal message err:%s", err.Error())
	}
	bytes := make([]byte, 0)
	for _, b := range body {
		bytes = append(bytes, b...)
	}
	a.WriteMsg(&pb.Message{Body: bytes, Header: &pb.Header{UserId: 1000}})
}
