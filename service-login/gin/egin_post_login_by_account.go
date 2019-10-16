package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
)

type LoginByAccount struct {
}

func NewLoginByAccount() LoginByAccount {
	return LoginByAccount{}
}

func (p *LoginByAccount) handle(c *gin.Context) {
	log.Logf("post login")
}
