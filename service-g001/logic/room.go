package logic

import (
	"fmt"
	"time"

	"github.com/Irfish/component/leaf/module"
	"github.com/Irfish/component/log"
	"github.com/Irfish/fantasy.server/pb"
)

//when the game end, the winner can be challenge by the looser in an another game scene of rpg
const (
	RoomStatusCrowd = 1
	RoomStatusFree  = 2

	GameStatusIdle    = 1 //空闲
	GameStatusRunning = 2 //游戏中...
)

type Room struct {
	Skeleton          *module.Skeleton
	MasterId          int64
	Id                int64
	level             int32
	ChairIdToPlayer   map[int32]*Player
	UserIdToPlayer    map[int64]*Player
	Status            int
	PlayerLimit       int32
	CloseSign         chan bool
	TimeCounterStatus int
	ChairIdSeed       []int32
	PlayerCount       int32
	GameStatus        int32
}

func NewRoom(playerLimit int32) *Room {
	r := new(Room)
	r.Status = RoomStatusFree
	r.ChairIdToPlayer = make(map[int32]*Player, 0)
	r.UserIdToPlayer = make(map[int64]*Player, 0)
	r.PlayerLimit = playerLimit
	r.ChairIdSeed = make([]int32, playerLimit)
	r.PlayerCount = 0
	r.TimeCounterStatus = 0
	r.GameStatus = GameStatusIdle
	InitTable(10, 10)
	return r
}

func (r *Room) InitChairIdSeed() {
	for i := int32(0); i < r.PlayerLimit; i++ {
		r.ChairIdSeed = append(r.ChairIdSeed, 0)
	}
}

func (r *Room) RandChairId() int32 {
	for k, v := range r.ChairIdSeed {
		if v == 0 {
			return int32(k)
		}
	}
	return -1
}

func (r *Room) PlayerEnter(player *Player) (p *Player, e error) {
	if _, ok := r.UserIdToPlayer[player.UserId]; ok {
		e = fmt.Errorf("user aready exist(%d)", player.UserId)
		return
	}
	if r.Status == RoomStatusCrowd {
		e = fmt.Errorf("room is crowd")
		return
	}
	chairId := r.RandChairId()
	if chairId == -1 {
		e = fmt.Errorf("room is crowd")
		return
	}
	player.ChairId = chairId
	player.Status = PlayerStatusOnline
	if (chairId+1)%2 == 1 {
		player.PieceColor = PieceValueBlack
	} else {
		player.PieceColor = PieceValueWhite
	}
	r.ChairIdToPlayer[chairId] = player
	r.UserIdToPlayer[player.UserId] = player
	r.ChairIdSeed[chairId] = 1
	r.PlayerCount++
	p = player
	return
}

func (r *Room) PlayerLeave(chairId int32) (player *Player, e error) {
	if r.GameStatus == GameStatusRunning {
		e = fmt.Errorf("game is running,can not leave(%d)", chairId)
		return
	}
	p, ok := r.ChairIdToPlayer[chairId]
	if !ok {
		e = fmt.Errorf("user not exist(%d)", chairId)
		return
	}
	player = p
	delete(r.UserIdToPlayer, p.UserId)
	delete(r.ChairIdToPlayer, chairId)
	r.ChairIdSeed[chairId] = 0
	r.PlayerCount--
	r.Status = RoomStatusFree
	return
}

func (r *Room) PlayerOffLine(chairId int32) {
	r.ChairIdToPlayer[chairId].Status = PlayerStatusOffline
}

func (r *Room) PlayerOnline(chairId int32) {
	r.ChairIdToPlayer[chairId].Status = PlayerStatusOnline
}

func (r *Room) PlayerPlayPiece(chairId int32, x, y int32) (list []*pb.Piece, e error) {
	player, ok := r.ChairIdToPlayer[chairId]
	if !ok {
		e = fmt.Errorf("user not exist(%d)", chairId)
		return
	}
	if r.GameStatus != GameStatusRunning {
		e = fmt.Errorf("game not running")
		return
	}
	list, e = player.Play(x, y)
	if len(list) >= WinCount {
		r.GameEnd()
	}
	return
}

func (r *Room) PlayerReady(chairId int32, status bool) (e error) {
	player, ok := r.ChairIdToPlayer[chairId]
	if !ok {
		e = fmt.Errorf("user not exist(%d)", chairId)
		return
	}
	player.Ready(status)
	r.CheckStartGame()
	return
}

func (r *Room) CheckStartGame() {
	count := 0
	for _, p := range r.ChairIdToPlayer {
		if p.ReadyStatus {
			count++
		}
	}
	if count == 2 {
		if r.TimeCounterStatus == 0 {
			r.GameStatus = GameStatusRunning
			r.StartTimer()
		}
	}
}

func (r *Room) GameEnd() {
	r.GameStatus = GameStatusIdle
	r.StopTimer()
}

func (r *Room) TimeCounter() {
	log.Debug("time out calling....%d", r.Id)
	if r.TimeCounterStatus == 1 {
		r.Skeleton.AfterFunc(time.Second*5, r.TimeCounter)
	}
}

func (r *Room) StartTimer() {
	r.TimeCounterStatus = 1
	r.TimeCounter()
}

func (r *Room) StopTimer() {
	r.TimeCounterStatus = 0
}
