package server

import (
	"reflect"

	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-gw/msg"
)

func handler(m interface{}, h interface{}) {
	ChanRpc.Register(reflect.TypeOf(m), h)
}

func init() {
	router()
	registerHandler()
}

func router() {
	msg.Processor.SetRouter(&pb.CtsUserAuthentication{}, ChanRpc)
}

func registerHandler() {
	handler(&pb.CtsUserAuthentication{}, ctsUserAuthenticationHandler)
}
