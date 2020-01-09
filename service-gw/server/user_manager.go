package server

import (
	"fmt"
	"github.com/Irfish/fantasy.server/pb"

	"github.com/Irfish/component/leaf/gate"
	"github.com/Irfish/component/log"
)

var (
	UserManager = NewManager()
)

type User struct {
	UserId    int32
	SessionId int64
	Agent     gate.Agent
}

func (u *User) SendMessage(bytes []byte) {
	u.Agent.WriteMsg(&pb.Message{Body: bytes, Header: &pb.Header{UserId: 1000}})
}

type Manager struct {
	SessionIdToUser map[int64]*User
	UserIdToUser    map[int32]*User
}

func NewManager() *Manager {
	m := new(Manager)
	m.SessionIdToUser = make(map[int64]*User, 0)
	m.UserIdToUser = make(map[int32]*User, 0)
	return m
}

func (m *Manager) UserConnect(sessionId int64, agent gate.Agent) error {
	if _, ok := m.SessionIdToUser[sessionId]; ok {
		return fmt.Errorf("UserConnect: sessionId(%d) has exist", sessionId)
	}
	user := new(User)
	user.SessionId = sessionId
	user.Agent = agent
	m.SessionIdToUser[sessionId] = user
	return nil
}

func (m *Manager) CheckMessage(sessionId int64) (e error) {
	if _, ok := m.SessionIdToUser[sessionId]; !ok {
		e = fmt.Errorf("message not illegal sessionId(%d)",sessionId)
	}
	return
}

func (m *Manager) UserDisconnect(sessionId int64) {
	if user, ok := m.SessionIdToUser[sessionId]; ok {
		userId := user.UserId
		if _, ok1 := m.UserIdToUser[userId]; ok1 {
			delete(m.UserIdToUser, userId)
		}
		delete(m.SessionIdToUser, sessionId)
	}
}

func (m *Manager) UserAuthentication(userId int32, sessionId int64) error {
	if _, ok := m.UserIdToUser[userId]; ok {
		return fmt.Errorf("UserAuthentication:userId(%d) has exist", userId)
	}
	user, ok := m.SessionIdToUser[sessionId]
	if !ok {
		return fmt.Errorf("UserAuthentication:user(%d) not exist", sessionId)
	}
	user.UserId = userId
	m.UserIdToUser[userId] = user
	return nil
}

func (m *Manager) GetAgentBySessionId(sessionId int64) *User {
	user, ok := m.SessionIdToUser[sessionId]
	if !ok {
		log.Debug("GetAgentBySessionId:user(%d) not found", sessionId)
		return nil
	}
	return user
}

func (m *Manager) GetAgentByUserId(userId int32) *User {
	user, ok := m.UserIdToUser[userId]
	if !ok {
		log.Debug("GetAgentByUserId:user(%d) not found", userId)
		return nil
	}
	return user
}
