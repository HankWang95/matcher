package test

import (
	"fmt"
	"github.com/HankWang95/matcher"
	"testing"
	"time"
)

var (
	userIdList = make([]int, 0, 20)
	gameType = make([]string, 0 ,5)
)

// 生成userid
func initUserList() {
	var i int
	for i < 20 {
		id := matcher.GetRandomId()
		time.Sleep(time.Second / 20)
		userIdList = append(userIdList, id)
		i++
	}
}

// 初始化gametype
func initGameType() {
	for i:=0; i<5; i++{
		gameType = append(gameType, fmt.Sprintf("type%d", i))
	}
}

var m matcher.Matcher

// 生成matcher
func initMatcher() {
	m = matcher.NewMatcher(4)
}

func init() {
	initUserList()
	initGameType()
	initMatcher()
}
func getUserId() int {
	if len(userIdList) > 0{
		id := userIdList[0]
		userIdList = userIdList[1:]
		return id
	}
	return 0
}
func getGameType() string {
	if len(gameType) > 0{
		t := gameType[0]
		return t
	}
	return ""
}

// ----------------------------------------
func TestMatchRoom(t *testing.T) {
	var i int
	for i< 20{
		userId :=getUserId()
		gType :=getGameType()
		roomId, err := m.MatchRoom(userId, gType)
		fmt.Println("用户 ", userId, " 进入房间 :", roomId, err)
		if i % 3 == 0{
			newRoomId, err := m.Move2AnotherRoom(userId)
			fmt.Println("用户 ", userId,"从房间退出 ", roomId, "换到房间", newRoomId, err)
		}
		i++
	}
	fmt.Println(m)
	time.Sleep(time.Second)
	fmt.Println("======================================================")

}
//
func TestMatchRoom2(t *testing.T) {
	initUserList()
	var i int
	for i< 20{
		userId :=getUserId()
		gType :="type1"
		roomId, err := m.MatchRoom(userId, gType)
		fmt.Println("用户 ", userId, " 进入房间 :", roomId, err)
		if i % 3 == 0{
			newRoomId, err := m.Move2AnotherRoom(userId)
			fmt.Println("用户 ", userId,"从房间退出 ", roomId, "换到房间", newRoomId, err)
		}
		i++
	}
	fmt.Println(m)
	time.Sleep(time.Second)
}

//func TestExitRoom(t *testing.T) {
//	m.ExitRoom()
//}
//
//
//func TestGetRandomId(t *testing.T) {
//	var i int
//	for i < 20 {
//		id := matcher.GetRandomId()
//		if id < 100000 {
//			t.Errorf("random function error! id:%d", id)
//			break
//		}
//		i++
//	}
//}
