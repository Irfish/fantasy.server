package main

import (
	"flag"
	"fmt"

	"github.com/Irfish/component/pid"
	db "github.com/Irfish/fantasy.server/service-db/service"
	g "github.com/Irfish/fantasy.server/service-g001/service"
	gw "github.com/Irfish/fantasy.server/service-gw/service"
	logS "github.com/Irfish/fantasy.server/service-log/service"
	login "github.com/Irfish/fantasy.server/service-login/service"
	web "github.com/Irfish/fantasy.server/service-web/service"
)

var s = flag.String("s", "", "service name:g001,db,gw,login,web,log")

func main() {
	flag.Parse()
	if *s != "" {
		pid.Pid(*s)
	}
	switch *s {
	case "g001":
		g.Run()
	case "db":
		db.Run()
	case "gw":
		gw.Run()
	case "login":
		login.Run()
	case "web":
		web.Run()
	case "log":
		logS.Run()
	default:
		fmt.Println("place input args -s xxx to run service")
	}
}
