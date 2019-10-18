package gin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Irfish/component/hash"
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
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	accountId, ok := c.GetPostForm("accountId")
	if !ok {
		e = fmt.Errorf("can not found accountId")
		return
	}
	id, err := strconv.Atoi(accountId)
	if err != nil {
		e = fmt.Errorf("accountId can not convert to int")
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok {
		e = fmt.Errorf("can not found password")
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
		e = fmt.Errorf("user not found")
		return
	}
	decodePwd := hash.Md5WithBase64(password)
	if u.Pwd != decodePwd {
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
