package gin

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Irfish/component/hash"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/component/token"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-login/orm"
	"github.com/Irfish/fantasy.server/service-login/server"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type LoginByAccount struct {
}

func NewLoginByAccount() LoginByAccount {
	p := LoginByAccount{}
	p.RedisParserInit()
	return p
}

func (p *LoginByAccount) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["status"] = false
			result["err"] = e.Error()
			log.Debug("%s",e.Error())
		}
		c.JSON(http.StatusOK, result)
	}()
	accountId, ok := c.GetPostForm("accountId")
	if !ok {
		e = fmt.Errorf("can not found accountId key")
		return
	}
	id, err := strconv.Atoi(accountId)
	if err != nil {
		e = fmt.Errorf("accountId can not convert to int")
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok {
		e = fmt.Errorf("can not found password key")
		return
	}
	u := orm.User{
		Id: int64(id),
	}
	exist, err := orm.UserXorm().Get(&u)
	if err != nil {
		e = fmt.Errorf("mysql get user err: %s", err.Error())
		return
	}
	if !exist {
		e = fmt.Errorf("user not exist ")
		return
	}
	decodePwd := hash.Md5WithBase64(password)
	if u.Pwd != decodePwd {
		e = fmt.Errorf("password not illegal")
		return
	}
	gwAddr := server.GetOneGwAddr()
	if gwAddr == "" {
		e = fmt.Errorf("can not found service gateway")
		return
	}
	expireTime := time.Now().Add(time.Second * 60 * 60 * 1).Unix()
	token := token.GenToken(expireTime, token.GetTokenKey(), u.Id)
	info := &pb.RedisUserLogin{
		Token:            token,
		TokenExpiredTime: expireTime,
	}
	_, e1 := redis.RedisHset("service.login.user.login", "info", info)
	if e1 != nil {
		e = fmt.Errorf("redis err: %s", e1.Error())
		return
	}

	result["userId"] = u.Id
	result["expireTime"] = expireTime
	result["token"] = base64.URLEncoding.EncodeToString(info.Token)
	result["status"] = true
	result["gw"] = strings.Split(gwAddr, ":")
	result["err"] = ""
}

func (p *LoginByAccount) RedisParserInit() {
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
				redisUserLogin := &pb.RedisUserLogin{}
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
