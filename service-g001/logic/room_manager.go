package logic

import (
	"fmt"

	"github.com/Irfish/component/uuid"
)

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

func (m *Manager) CreateRoom(userId int64) (int64, error) {
	room := NewRoom(2)
	room.Id = uuid.GenUid()
	room.MasterId = userId
	m.RoomIdToRoom[room.Id] = room

	go func() {
		room.OnInit()
		room.Run(room.CloseSign)
	}()

	return room.Id,nil
}

func (m *Manager) PlayerEnterRoom(userId, roomId int64) (*Player, error) {
	room, ok := m.RoomIdToRoom[roomId]
	if !ok {
		return nil, fmt.Errorf("room not exist (id:%d)", roomId)
	}
	p := &Player{
		RoomId: room.Id,
		UserId: userId,
	}
	e := room.PlayerEnter(p)
	return p, e
}

func (m *Manager) PlayerLeaveRoom(chairId int32, roomId int64) (*Player, error) {
	room, ok := m.RoomIdToRoom[roomId]
	if !ok {
		return nil, fmt.Errorf("room not exist (id:%d)", roomId)
	}
	p := &Player{}
	room.PlayerLeave(chairId)
	return p, nil
}
