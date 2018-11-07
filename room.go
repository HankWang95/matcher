package matcher

import (
	"fmt"
	"sync"
)

func CreateNewRoom(roomId int, gameType string) *room {
	r := new(room)
	r.id = roomId
	r.players = make(map[int]bool)
	r.gameType = gameType
	return r
}

type room struct {
	sync.RWMutex
	id       int
	players  map[int]bool
	gameType string
}

func (r *room) InsertPlayer(playerId int) (roomId int) {

	r.RLock()
	r.players[playerId] = true
	r.RUnlock()

	r.checkSelf()
	return r.id
}

func (r *room) isFull() bool {
	r.Lock()
	defer r.Unlock()
	return len(r.players) == maxP
}

func (r *room) isEmpty() bool {
	r.Lock()
	defer r.Unlock()
	return len(r.players) == 0
}

var iJu int
// 判断房间状态（人数、是否在队列中） 并做出相应动作
func (r *room) checkSelf() {
	// 人满 从队列中移除
	if r.isFull() {
		fmt.Println("房间id ", r.id ,"房间人数", r.players)
		iJu++
		fmt.Printf("人满开局 ---- 第 %d 局 \n", iJu)
		queueMap[r.gameType].dropRoom(r.id)
		return
	}

	// 人空 还回房间->database
	if r.isEmpty() {
		queueMap[r.gameType].dropRoom(r.id)
		return
	}

	// 向队列中添加房间
	queueMap[r.gameType].insertRoom(r)

	return
}

func (r *room) GetRoomId() int {
	return r.id
}

func (r *room) ExitRoom(playerId int) {
	r.RLock()
	delete(r.players, playerId)
	r.RUnlock()
	r.checkSelf()
}
