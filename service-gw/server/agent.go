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
	skeleton.RegisterChanRPC("OnServiceMessage", onServiceMessage)
}

func rpcNewAgent(args []interface{}) {
	log.Debug("one client connect to gw")
	a := args[0].(gate.Agent)
	sessionId := uuid.GenUid()
	a.SetUserData(sessionId)
	e := UserManager.UserConnect(sessionId, a)
	if e != nil {
		log.Debug(e.Error())
	}
}

func rpcCloseAgent(args []interface{}) {
	log.Debug("one client disconnect from gw")
	a := args[0].(gate.Agent)
	id := a.UserData()
	sessionId := id.(int64)
	UserManager.UserDisconnect(sessionId)
}

func rpcSwitchRouterMsg(args []interface{}) {
	//log.Debug("switchRoute message")
	a := args[1].(gate.Agent)
	message := args[0].(*pb.Message)
	m, err := msg.Processor.Unmarshal(message.Body)
	if err != nil {
		log.Debug("unmarshal message error: %v", err)
	}
	m1 := m.(proto.Message)
	//log.Debug("switchRoute m1: %s", m1.String())
	switch message.Header.ServiceId0 {
	case int32(pb.SERVICE_G001):
		e := UserManager.CheckMessage(message.Header.SessionId)
		if e != nil {
			sendMessage(a, &pb.StcErrorNotice{
				Info: e.Error(),
			})
			return
		}
		c, e := GetService(pb.SERVICE_G001)
		if e != nil {
			sendMessage(a, &pb.StcErrorNotice{
				Info: e.Error(),
			})
			return
		}
		c.SendToService(m1)
	default:
		err := msg.Processor.Route(m, a)
		if err != nil {
			log.Debug("route message error: %v", err)
			break
		}
	}
}

func onServiceMessage(args []interface{}) {
	//a := args[1].(tcpclient.Agent)
	//log.Debug("%s", a.GetClientId())
	message := args[0].(*pb.Message)
	m, err := msg.Processor.Unmarshal(message.Body)
	if err != nil {
		log.Debug("unmarshal message error: %v", err)
	}
	if message.Header.Broadcast {
		for _, user := range UserManager.UserIdToUser {
			user.SendMessage(m)
		}
	} else {
		user := UserManager.GetAgentByUserId(message.Header.UserId)
		if user != nil {
			user.SendMessage(m)
		}
	}
}
