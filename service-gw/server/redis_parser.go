package server

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/golang/protobuf/proto"
)

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
