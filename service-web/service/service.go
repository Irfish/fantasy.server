package service

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Irfish/component/log"
	gin1 "github.com/Irfish/fantasy.server/service-web/gin"
)

func Run() {
	log.Debug("fantasy service web running ")
	gin1.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}
