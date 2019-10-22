package service

import (
	"github.com/Irfish/component/etcd3"
	"github.com/Irfish/component/leaf"
	lconf "github.com/Irfish/component/leaf/conf"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/service-g001/base"
	"github.com/Irfish/fantasy.server/service-g001/server"
)

func Run() {
	log.Debug("fantasy service game running ")
	lconf.LogLevel = base.Server.LogLevel
	lconf.LogPath = base.Server.LogPath
	lconf.LogFlag = base.LogFlag
	lconf.ConsolePort = base.Server.ConsolePort
	lconf.ProfilePath = base.Server.ProfilePath
	log.Debug(base.Server.EtcdAddr)
	//连接etcd
	etcd3.Init([]string{base.Server.EtcdAddr}, 3)

	leaf.Run(server.Module)
}
