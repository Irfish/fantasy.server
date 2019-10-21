package gin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Irfish/component/hash"
	"github.com/Irfish/component/token"
	"github.com/Irfish/fantasy.server/service-login/orm"
	"github.com/Irfish/fantasy.server/service-login/server"
	"github.com/gin-gonic/gin"
)

type LoginByAccount struct {
}

func NewLoginByAccount() LoginByAccount {
	return LoginByAccount{}
}

func (p *LoginByAccount) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["status"] = false
			result["err"] = e.Error()
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
		e = err
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

	result["userId"] = u.Id
	result["expireTime"] = expireTime
	result["token"] = token
	result["status"] = true
	result["gw"] = gwAddr
	result["err"] = ""
}
