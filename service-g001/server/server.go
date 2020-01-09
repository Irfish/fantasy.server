package server

import (
	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/base"
	"github.com/Irfish/fantasy.server/service-g001/msg"
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
	skeletonCloseSig chan bool
	Node             etcd3.ServiceNode
}

func (s *Server) OnInit() {
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
	//启动leaf的skeleton
	s.skeletonCloseSig = make(chan bool, 1)
	go func() {
		skeleton.Run(s.skeletonCloseSig)
	}()
	//注册到etcd
	s.SetNodeData(&etcd3.Node{Key: pb.GetServerKey(pb.SERVICE_G001), Name: pb.GetName(pb.SERVICE_G001), Address: base.Server.TCPAddr, Version: "v1.0"})
	etcd3.RegisterNode(s)
}

func (s *Server) BeforeDestroy() {
	s.skeletonCloseSig <- true
}

func (s *Server) GetNodePrefix() string {
	return pb.ServicePrefix()
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
	a.WriteMsg(&pb.Message{Body: bytes, Header: &pb.Header{
		ServiceId0: int32(pb.SERVICE_GW),
		UserId:     1000,
	}})
}
