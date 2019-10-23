package server

import (
	"reflect"

	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/msg"
)

func handler(m interface{}, h interface{}) {
	ChanRpc.Register(reflect.TypeOf(m), h)
}

func init() {
	router()
	registerHandler()
}

func router() {
	msg.Processor.SetRouter(&pb.CtsUserEnter{}, ChanRpc)
	msg.Processor.SetRouter(&pb.CtsUserLeave{}, ChanRpc)
	msg.Processor.SetRouter(&pb.CtsCreateRoom{}, ChanRpc)
}

func registerHandler() {
	handler(&pb.CtsUserEnter{}, ctsUserEnterHandler)
	handler(&pb.CtsUserLeave{}, ctsUserLeaveHandler)
	handler(&pb.CtsCreateRoom{}, ctsCreateRoomHandler)
}
