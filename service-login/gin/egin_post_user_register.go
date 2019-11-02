package gin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Irfish/component/hash"
	"github.com/Irfish/fantasy.server/service-login/orm"
	"github.com/gin-gonic/gin"
)

type UserRegister struct {
}

func NewUserRegister() UserRegister {
	p := UserRegister{}
	return p
}

func (p *UserRegister) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	userName, ok := c.GetPostForm("userName")
	if !ok {
		e = fmt.Errorf("the key userName not found")
		return
	}
	pwd, ok := c.GetPostForm("pwd")
	if !ok {
		e = fmt.Errorf("the key pwd not found")
		return
	}

	u := orm.User{
		UserName:   userName,
		Pwd:        hash.Md5WithBase64(pwd),
		CreateTime: time.Now().Unix(),
		Level:      1,
	}
	effectCount, err := orm.UserXorm().Insert(&u)
	if err != nil {
		e = err
		return
	}
	if effectCount == 0 {
		e = fmt.Errorf("create a new user failed")
		return
	}
	result["status"] = true
}
