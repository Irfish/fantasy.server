package client

import (
	"github.com/Irfish/component/leaf/tcpclient"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
)

func stcUserEnterHandler(args []interface{}) {
	m := args[0].(*pb.StcUserEnter)
	log.Debug("received message form server:%s", m.UserId)
	a := args[1].(tcpclient.Agent)
	_ = a
}
