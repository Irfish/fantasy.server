package server

import (
	"encoding/base64"
	"fmt"

	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/golang/protobuf/proto"
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
		e = fmt.Errorf("token illegal")
		return
	}
	sendMessage(a, &pb.StcUserAuthentication{
		Result:    "user authentication success",
		SessionId: sessionId,
	})
}

func RedisParserInit() {
	redis.AppendRedisMarshal(func(keys []string, i1 interface{}) (i interface{}, b bool) {
		if keys[0] == "hset" && keys[1] == "service.login.user.login" && keys[2] == "info" {
			i0, ok := i1.(*pb.RedisUserLogin)
			if ok {
				i2, e := proto.Marshal(i0)
				if e != nil {
					log.Debug("UserRegister RedisParserInit AppendRedisMarshal RedisUserLogin error:", e.Error())
					return
				}
				i = i2
				b = true
			}
		}
		return
	})
	redis.AppendRedisUnmarshal(func(keys []string, i1 interface{}) (i interface{}, b bool) {
		if keys[0] == "hget" && keys[1] == "service.login.user.login" && keys[2] == "info" {
			i0, ok := i1.([]byte)
			if ok {
				redisUserLogin := new(pb.RedisUserLogin)
				e := proto.Unmarshal(i0, redisUserLogin)
				if e != nil {
					log.Debug("UserRegister RedisParserInit AppendRedisUnmarshal RedisUserLogin error:", e.Error())
					return
				}
				i = redisUserLogin
				b = true
			}
		}
		return
	})

}
