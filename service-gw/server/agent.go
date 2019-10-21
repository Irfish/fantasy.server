package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/uuid"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/msg"
	"github.com/golang/protobuf/proto"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("SwitchRouterMsg", rpcSwitchRouterMsg)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	sessionId:= uuid.GenUid()
	a.SetUserData(sessionId)
	log.Debug("one client connect to gw")
	sendMessage(a, &pb.StcUserEnter{UserId:int32(12500),Result:"sddff"})
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	log.Debug("one client disconnect from gw")
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
	a.WriteMsg(&pb.Message{Body: bytes, Header: &pb.Header{UserId:1000}})
}
