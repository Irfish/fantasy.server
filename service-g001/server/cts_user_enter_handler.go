package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsUserEnterHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserEnter)
	a := args[1].(gate.Agent)
	log.Debug("user enter:%d", m.RoomId)
	p, e := logic.RoomManager.PlayerEnterRoom(m.UserId, m.RoomId)
	if e != nil {
		sendMessage(a, &pb.StcErrorNotice{
			Info: e.Error(),
		})
		return
	}
	sendMessage(a, &pb.StcUserEnter{
		UserId:  p.UserId,
		RoomId:  p.RoomId,
		ChairId: p.ChairId,
	})
}
