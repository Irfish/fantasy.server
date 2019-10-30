package gin

import (
	"net/http"
	"time"

	"github.com/Irfish/component/log"
	"github.com/gin-gonic/gin"
)

//获取服务器时间
type ServiceTime struct {
}

func NewServiceTime() ServiceTime {
	return ServiceTime{}
}

func (p *ServiceTime) handle(c *gin.Context) {
	s := gin.H{"time": time.Now().Unix()}
	c.JSON(http.StatusOK, s)
	log.Debug("ServiceTime %v", s)
}
