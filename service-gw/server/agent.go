package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/uuid"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/msg"
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
		c.SendToService(message.Body,message.Header.Id)
	default:
		m, err := msg.Processor.Unmarshal(message.Body)
		if err != nil {
			log.Debug("unmarshal message error: %v", err)
		}
		//m1 := m.(proto.Message)
		err = msg.Processor.Route(m, a)
		if err != nil {
			log.Debug("route message error: %v", err)
			break
		}
	}
}

func onServiceMessage(args []interface{}) {
	message := args[0].(*pb.Message)
	if message.Header.Broadcast {
		for _, user := range UserManager.UserIdToUser {
			user.SendMessage(message.Body)
		}
	} else {
		user := UserManager.GetAgentByUserId(message.Header.UserId)
		if user != nil {
			user.SendMessage(message.Body)
		}
	}
}
