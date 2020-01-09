package server

import (
	"encoding/base64"
	"fmt"
	"github.com/Irfish/component/log"

	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/fantasy.server/pb"
)

func ctsUserAuthenticationHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserAuthentication)
	a := args[1].(gate.Agent)
	var e error
	defer func() {
		if e != nil {
			sendMessage(a, &pb.StcErrorNotice{
				Info: e.Error(),
			})
			log.Debug("ctsUserAuthenticationHandler err:%s",e.Error())
		}
	}()
	userData := a.UserData()
	sessionId := userData.(int64)
	e = UserManager.UserAuthentication(m.UserId, sessionId)
	if e != nil {
		return
	}
	r, err := redis.RedisHget("service.login.user.login", "info")
	if err != nil {
		e = fmt.Errorf("%s", err.Error())
		return
	}
	if r == nil {
		e = fmt.Errorf(" redis.RedisHget key :service.login.user.login  is nil")
		return
	}
	info := r.(*pb.RedisUserLogin)
	tokenStr := base64.URLEncoding.EncodeToString(info.Token)
	if tokenStr != string(m.Token) {
		e = fmt.Errorf("token illegal %s  %s", tokenStr, string(m.Token))
		return
	}
	sendMessage(a, &pb.StcUserAuthentication{
		Result:    "user authentication success",
		SessionId: sessionId,
	})
}
