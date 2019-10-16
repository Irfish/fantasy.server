package service

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/component/xorm"
	gin1 "github.com/Irfish/fantasy.server/service-login/gin"
	"github.com/Irfish/fantasy.server/service-login/server"
)

func Run() {
	log.Debug("fantasy service login running ")
	etcd3.Init([]string{"192.168.0.131:2888"}, 3)
	redis.Run()
	xorm.Run()
	gin1.Run()
	server.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-ch
}
