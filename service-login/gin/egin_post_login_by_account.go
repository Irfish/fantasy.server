package gin

import (
	"fmt"
	"net/http"

	"github.com/Irfish/component/hash"
	"github.com/Irfish/component/log"
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
	log.Debug("post login %v", c.Keys)
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	accountId, ok := c.Keys["accountId"]
	if !ok {
		e = fmt.Errorf("can not found accountId")
		return
	}
	password, ok := c.Keys["password"]
	if !ok {
		e = fmt.Errorf("can not found password")
		return
	}
	u := orm.User{
		UserName: accountId.(string),
	}
	exist, err := orm.UserXorm().Get(&u)
	if err != nil {
		e = err
		return
	}
	if !exist {
		e = fmt.Errorf("user not found")
		return
	}
	pwd := password.(string)
	decodePwd := hash.Md5WithBase64(pwd)
	if u.Pwd == decodePwd {
		e = fmt.Errorf("password not illegal")
		return
	}
	gwAddr := server.GetOneGwAddr()
	if gwAddr == "" {
		e = fmt.Errorf("can not found service gw")
		return
	}
	result["status"] = true
	result["gw"] = gwAddr
}
