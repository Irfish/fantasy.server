package service

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/service-db/server"
)

func Run() {
	log.Debug("fantasy service db running ")
	server.Run()
}
