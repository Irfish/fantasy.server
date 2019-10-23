package logic

import "github.com/Irfish/component/leaf/module"

//when the game end, the winner can be challenge by the looser in an another game scene of rpg
const (
	ROOM_STATUS_CROWD = iota
	ROOM_STATUS_FREE
)

type Room struct {
	Skeleton         *module.Skeleton
	Id               int64
	level            int32
	ChairIdToPlayer map[int32]*Player
	Status           int
	PlayerLimit      int32
	CurrentId        int32
	CloseSign        chan bool
}

func NewRoom(playerLimit int32) *Room {
	r := new(Room)
	r.ChairIdToPlayer = make(map[int32]*Player, 0)
	r.CurrentId = 0
	r.PlayerLimit = playerLimit
	return r
}

func (r *Room) PlayerEnter(player *Player) {
	if r.PlayerLimit == r.CurrentId+1 {
		r.Status = ROOM_STATUS_CROWD
	}
	if r.Status == ROOM_STATUS_CROWD {
		return
	}
	player.ChairId = r.CurrentId
	player.Status = PLAYER_STATUS_ONLINE
	r.ChairIdToPlayer[player.ChairId] = player
	r.CurrentId = r.CurrentId + 1
}

func (r *Room) PlayerLeave(id int32) {
	delete(r.ChairIdToPlayer, id)
	r.Status = ROOM_STATUS_FREE
}

func (r *Room) PlayerOffLine(id int32) {
	r.ChairIdToPlayer[id].Status = PLAYER_STATUS_OFFLINE
}

func (r *Room) PlayerOnline(id int32) {
	r.ChairIdToPlayer[id].Status = PLAYER_STATUS_ONLINE
}
