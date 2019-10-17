package server

import (
	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
)

var service *Server

func Run() {
	service = new(Server)
	service.Init()
}

type Server struct {
	*etcd3.ServiceNode
	GatewayNodes map[string]*etcd3.Node
}

func (s *Server) Init() {
	s.ServiceNode = new(etcd3.ServiceNode)
	s.GatewayNodes = make(map[string]*etcd3.Node, 0)
	//监控所有gateway服务状态
	etcd3.WatchNode(s)
	//第一次启动时尝试获取所有的gateway服务
	nodes := etcd3.GetNodes(pb.GwPrefix())
	for _, n := range nodes {
		s.GatewayNodes[n.Key] = n
	}
}

func (s *Server) GetNodePrefix() string {
	return pb.GwPrefix()
}

func (s *Server) OnNodeRegister(k string, v interface{}) {
	log.Debug("OnNodeRegister: %s = %v", k, v)
	node := v.(*etcd3.Node)
	s.GatewayNodes[node.Key] = node
}

func (s *Server) OnNodeUnregister(k string) {
	log.Debug("OnNodeUnregister: %s ", k)
	if _, ok := s.GatewayNodes[k]; ok {
		delete(s.GatewayNodes, k)
	}
}

func GetOneGwAddr() string {
	for _, n := range service.GatewayNodes {
		return n.Address
	}
	return ""
}
