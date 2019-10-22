package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
)

func ctsUserAuthenticationHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserAuthentication)
	a := args[1].(gate.Agent)
	//log.Debug("received message form client:%d", m.UserId)
	userData := a.UserData()
	sessionId := userData.(int64)
	e := UserManager.UserAuthentication(m.UserId, sessionId)
	if e != nil {
		log.Debug(e.Error())
		return
	}
	sendMessage(a, &pb.StcUserAuthentication{
		Result:    "user authentication success",
		SessionId: sessionId,
	})
}
