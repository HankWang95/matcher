package matcher

import (
	"math/rand"
	"sync"
	"time"
)

// 创建新队列
func CreateNewQueue(gameType string) *queue {
	if v, ok := queueMap[gameType]; ok {
		return v
	}
	q := &queue{}
	q.RWMutex = &sync.RWMutex{}
	q.gameType = gameType
	q.roomMap = make(map[int]*room)
	queueMap[gameType] = q
	return q
}

type queue struct {
	*sync.RWMutex
	gameType string
	roomMap  map[int]*room
	fl string
}

// 向队列中加入房间
func (q *queue) insertRoom(room *room) {

	q.RLock()
	if _, ok := q.roomMap[room.GetRoomId()]; !ok {
		q.roomMap[room.GetRoomId()] = room
	}
	q.RUnlock()
}

// 从队列中移除房间
func (q *queue) dropRoom(roomId int) {
	q.RLock()
	delete(q.roomMap, roomId)
	q.RUnlock()
	return
}

func (q *queue) GetRoom() *room {
	for _, v := range q.roomMap {
		if v != nil {
			return v
		}
	}
	// TODO 调用db 得到一个room id
	var id int
	id = GetRandomId()
	newRoom := CreateNewRoom(id, q.gameType)
	q.insertRoom(newRoom)
	return newRoom
}

func (q *queue) ChangeRoom(roomId int) *room {
	for k, v := range q.roomMap {
		if v != nil && k != roomId {
			return v
		}
	}

	// TODO 调用db 得到一个room id
	var id int
	id = GetRandomId()
	newRoom := CreateNewRoom(id, q.gameType)
	q.insertRoom(newRoom)
	return newRoom
}


// ===========================

// 模拟从数据库取房间号
func GetRandomId() (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(999999)
	if  result < 100000{
		result = GetRandomId()
	}
	return
}