package logic

import "github.com/Irfish/component/uuid"

var (
	RoomManager = newManager()
)

type Manager struct {
	RoomIdToRoom map[int64]*Room
}

func newManager() *Manager {
	m := new(Manager)
	m.RoomIdToRoom = make(map[int64]*Room, 0)
	return m
}

func (m *Manager) CreateRoom(userId int64) Player {
	room := NewRoom(2)
	room.Id = uuid.GenUid()
	m.RoomIdToRoom[room.Id] = room
	go func() {
		room.OnInit()
		room.Run(room.CloseSign)
	}()
	return m.PlayerEnterRoom(userId)
}

func (m *Manager) PlayerEnterRoom(userId int64) Player {
	p := Player{}
	return p
}

func (m *Manager) PlayerLeaveRoom(userId int64) Player {
	p := Player{}
	return p
}
