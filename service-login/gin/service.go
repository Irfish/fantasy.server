package gin

import (
	g "github.com/Irfish/component/gin"
)

type Gin struct {
}

func Run() {
	server := new(Gin)
	g.Run(server)
}
