package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
)

type Login struct {
}

func NewLogin() Login {
	return Login{}
}

func (p *Login) handle(c *gin.Context) {
	log.Logf("post login")
}
