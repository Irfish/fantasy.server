package server

import (
	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
	"github.com/Irfish/fantasy.server/service-g001/logic"
)

func ctsUserReadyHandler(args []interface{}) {
	m := args[0].(*pb.CtsUserReady)
	a := args[1].(gate.Agent)
	log.Debug("create room:%d", m.ChairId)
	e := logic.RoomManager.PlayerReady(m.ChairId, m.RoomId, m.Status)
	if e != nil {
		sendMessage(a, &pb.StcErrorNotice{
			Info: e.Error(),
		})
		return
	}
	sendMessage(a, &pb.StcUserReady{
		RoomId:  m.RoomId,
		ChairId: m.ChairId,
		Status:  m.Status,
	})
}
