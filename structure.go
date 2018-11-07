package matcher

import (
	"errors"

)

var (
	// 游戏的每桌最大人数
	maxP int
)

type Matcher interface {
	MatchRoom(userId int, gameType string) (roomId int, err error)

	Move2AnotherRoom(userId int) (roomId int, err error)

	ExitRoom(userId int) (err error)
}

//---------------------------------------------------------------

var queueMap  map[string]*queue


type matcher struct {
	playerRoomMap map[int]*room
	gameName string
}

func NewMatcher(maxPlayer int, gameName string) Matcher {
	var m = new(matcher)
	maxP = maxPlayer
	queueMap = make(map[string]*queue)
	m.playerRoomMap = make(map[int]*room)
	m.gameName = gameName
	return m
}

func (m *matcher) MatchRoom(playerId int, gameType string) (roomId int, err error) {
	var q *queue
	if v, ok := queueMap[gameType]; ok {
		q = v
	} else {
		q = CreateNewQueue(gameType)
		queueMap[gameType] = q
	}

	// 分配房间
	r := q.GetRoom()
	// 加入房间，返回房间号
	roomId = r.InsertPlayer(playerId)

	if roomId != 0 {
		m.playerRoomMap[playerId] = r
		return roomId, nil
	} else {
		return 0, errors.New(" 加入房间失败 ")
	}

}

func (m *matcher) Move2AnotherRoom(userId int) (roomId int, err error) {
	// 通过user id 查到 room， 进行换房操作
	if r, ok := m.playerRoomMap[userId]; !ok {
		return 0, errors.New(" 玩家没有加入任何房间，不能换桌")
	} else {
		q := queueMap[r.gameType]
		err := m.ExitRoom(userId)
		if err!= nil{
			return 0, err
		}
		newRoom := q.ChangeRoom(r.id)
		return newRoom.GetRoomId(), nil
	}
}

func (m *matcher) ExitRoom(userId int) (err error) {
	if r, ok := m.playerRoomMap[userId]; !ok {
		return errors.New(" 玩家没有加入任何房间，退出失败")
	} else {
		r.ExitRoom(userId)
		delete(m.playerRoomMap, userId)
		return nil
	}
}
