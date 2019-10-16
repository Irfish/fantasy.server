package server

import (
	"github.com/golang/protobuf/proto"
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/msg"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("SwitchRouterMsg", rpcSwitchRouterMsg)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcSwitchRouterMsg(args []interface{}) {
	log.Debug("switchRoute message")
	a := args[1].(gate.Agent)
	message := args[0].(*pb.Message)
	m, err := msg.Processor.Unmarshal(message.Body)
	if err != nil {
		log.Debug("unmarshal message error: %v", err)
	}
	m1 := m.(proto.Message)
	log.Debug("switchRoute m1: %s", m1.String())
	switch message.Header.ServiceId0 {
	default:
		err := msg.Processor.Route(m, a)
		if err != nil {
			log.Debug("route message error: %v", err)
			break
		}
	}
}
