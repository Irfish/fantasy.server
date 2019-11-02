package gin

import (
	g "github.com/Irfish/component/gin"
	"github.com/Irfish/fantasy.server/service-login/base"
)

type Gin struct {
	Address string
}
func (p *Gin)Addr() string {
	return p.Address
}

func Run() {
	server := new(Gin)
	server.Address=base.Server.GinAddr
	g.Run(server)
}
