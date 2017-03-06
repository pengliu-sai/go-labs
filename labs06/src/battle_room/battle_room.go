package battle_room

import "battle_core"

var gBattleRooms map[int64]*BattleRoom

func init() {
	gBattleRooms = make(map[int64]*BattleRoom)
}

type BattleRoom struct {
	players             map[int64]*BattleRoomPlayer // 房间内的玩家
	maxBattlePlayerNums int                         // 最大战斗玩家个数
	onRound             func(*battle_core.Battle)
}

func NewBattleRoom(playerId int64, sideInfo *battle_core.SideInfo, onRound func(*battle_core.Battle)) *BattleRoom {
	br := &BattleRoom{
		players:             make(map[int64]*BattleRoomPlayer),
		maxBattlePlayerNums: 2,
		onRound:             onRound,
	}

	gBattleRooms[playerId] = br

	br.Join(playerId, sideInfo)

	return br
}

// 加入房间
func (br *BattleRoom) Join(playerId int64, sideInfo *battle_core.SideInfo) {
	br.players[playerId] = &BattleRoomPlayer{
		playerId: playerId,
		sideInfo: sideInfo,
	}

	// 房间人准备的人满时，开始战斗
	if len(br.players) == br.maxBattlePlayerNums {
		br.StartBattle()
	}
}

// 离开房间
func (br *BattleRoom) Leave(playerId int64) {
	delete(br.players, playerId)

	// 当房间内的人都离开时， 关闭房间
	br.Close()
}

// 关闭房间
func (br *BattleRoom) Close() {

}

// 开始战斗
func (br *BattleRoom) StartBattle() {
	// b := battle_core.NewBattle(br.players[0].sideInfo, br.players[1].sideInfo, br.onRound)
}

// 广播 加入房间
func (br *BattleRoom) BroadcastJoinRoom() {

}

// 广播 退出房间
func (br *BattleRoom) BroadcaseLeaveRoom() {

}

// 广播 关闭房间
func (br *BattleRoom) BroadcastCloseRoom() {

}

// 广播 StartBattle
func (br *BattleRoom) BroadcastStartBattle() {

}

// 广播 NextRound
func (br *BattleRoom) BroadcastNextRound() {

}

// 广播 Continue
func (br *BattleRoom) BroadcastContinue() {

}

// type RoomState byte // 战斗房间状态

// const (
// 	ROOM_STATE_READY RoomState = 0 // 准备中
// 	ROOM_STATE_LOCK            = 1 // 锁定中
// 	ROOM_STATE_END             = 2 // 战斗结束
// )

// type BattleRoom struct {
// 	battle  *battle_core.Battle         // 战斗基础类
// 	players map[int64]*BattleRoomPlayer // 房间内的玩家
// 	state   RoomState                   // 战场房间状态

// 	playerNum int // 参加战斗人数(冗余)
// }

// type NextRoundParam struct {
// }

// func NewBattleRoom(attacker, defender *battle_core.SideInfo, onRound func(*battle_core.Battle)) *BattleRoom {

// 	playerNum := len(attacker.Players) + len(defender.Players)
// 	players := make(map[int64]*BattleRoomPlayer, playerNum)

// 	for _, playerId := range attacker.Players {
// 		players[playerId] = &BattleRoomPlayer{
// 			playerId: playerId,
// 		}
// 	}

// 	for _, playerId := range defender.Players {
// 		players[playerId] = &BattleRoomPlayer{
// 			playerId: playerId,
// 		}
// 	}

// 	b := battle_core.NewBattle(attacker, defender, onRound)

// 	br := &BattleRoom{
// 		battle:    b,
// 		players:   players,
// 		playerNum: playerNum,
// 	}

// 	return br
// }
