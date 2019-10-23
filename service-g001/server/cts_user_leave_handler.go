package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsUserLeaveHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserLeave)
	a := args[1].(gate.Agent)
	log.Debug("user leave:%d", m.RoomId)
	p, e := logic.RoomManager.PlayerLeaveRoom(m.ChairId, m.RoomId)
	if e != nil {
		sendMessage(a, &pb.StcErrorNotice{
			Info: e.Error(),
		})
		return
	}
	sendMessage(a, &pb.StcUserLeave{
		UserId:  p.UserId,
		RoomId:  p.RoomId,
		ChairId: p.ChairId,
	})
}
