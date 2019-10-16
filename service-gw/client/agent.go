package client

import (
	"github.com/Irfish/component/leaf/tcpclient"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/msg"
)

func (c *Client) InitRegister() {
	c.Skeleton.RegisterChanRPC("ClientConnected", rpcClientConnected)
	c.Skeleton.RegisterChanRPC("ClientClosed", rpcClientClosed)
	c.Skeleton.RegisterChanRPC("SwitchRouterMsg", rpcSwitchRouterMsg)
}

func rpcClientConnected(args []interface{}) {
	a := args[0].(tcpclient.Agent)
	serviceToAgent[a.GetClientId()] = a
	log.Debug("connected to service: %s", a.GetClientId())
	sendMessage(a, &pb.CtsUserEnter{UserId: 10000, Token: []byte(a.GetClientId())})
	_ = a
}

func rpcClientClosed(args []interface{}) {
	a := args[0].(tcpclient.Agent)
	delete(serviceToAgent, a.GetClientId())
	log.Debug("disconnected from service: %s", a.GetClientId())
	_ = a
}

func rpcSwitchRouterMsg(args []interface{}) {
	log.Debug("switchRoute message")
	a := args[1].(tcpclient.Agent)
	message := args[0].(*pb.Message)
	m, err := msg.Processor.Unmarshal(message.Body)
	if err != nil {
		log.Debug("unmarshal message error: %v", err)
	}
	switch message.Header.ServiceId0 {
	default:
		err := msg.Processor.Route(m, a)
		if err != nil {
			log.Debug("route message error: %v", err)
			break
		}
	}
}
