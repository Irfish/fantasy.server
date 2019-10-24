package server

import (
	"fmt"

	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	client2 "github.com/Irfish/fantasy.server/service-gw/client"
)

var ServiceNode = NewNode()

func NewNode() *Node {
	n := new(Node)
	n.ServiceToClient = make(map[string]*client2.Client, 0)
	return n
}

type Node struct {
	*etcd3.ServiceNode
	ServiceToClient map[string]*client2.Client
}

func (s *Node) GetNodePrefix() string {
	return pb.ServicePrefix()
}

func (s *Node) OnNodeRegister(key string, d interface{}) {
	log.Debug("OnNodeRegister: %s = %v", key, d)
	node := d.(*etcd3.Node)
	c := CreateServiceClient(node.Name, node.Address, false)
	s.ServiceToClient[key] = c
}

func (s *Node) OnNodeUnregister(key string) {
	log.Debug("OnNodeUnregister: %s ", key)
	delete(s.ServiceToClient, key)
}

func CreateServiceClient(name, addr string, autoConnect bool) *client2.Client {
	gameClient := client2.NewClient(addr, name, autoConnect)
	gameClient.AppendOuterChanRpc(pb.GetName(pb.SERVICE_GW), ChanRpc)
	go func() {
		destroyChan := make(chan bool, 1)
		go gameClient.Run(destroyChan)
		<-destroyChan
		gameClient = nil
	}()
	return gameClient
}

func GetService(service pb.SERVICE) (*client2.Client, error) {
	key := pb.GetServerKey(service)
	if c, ok := ServiceNode.ServiceToClient[key]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("can not found service: %s", key)
}
