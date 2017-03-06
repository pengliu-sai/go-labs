package battle_room

import "battle_core"

var gBattleRooms map[int64]*BattleRoom

func init() {
	gBattleRooms = make(map[int64]*BattleRoom)
}

type NextRoundParam struct {
	PlayerId int64
}

type BattleRoom struct {
	players             map[int64]*BattleRoomPlayer // 房间内的玩家
	battlePlayers       []*BattleRoomPlayer         // 战斗选手
	maxBattlePlayerNums int                         // 最大战斗玩家个数
	nextRoundParams     []*NextRoundParam           // nextRound请求
	onRound             func(*battle_core.Battle)
	battle              *battle_core.Battle
}

func NewBattleRoom(playerId int64, sideInfo *battle_core.SideInfo, onRound func(*battle_core.Battle)) *BattleRoom {
	br := &BattleRoom{
		players:             make(map[int64]*BattleRoomPlayer),
		maxBattlePlayerNums: 2,
		nextRoundParams:     make([]*NextRoundParam, 0, 2),
		onRound:             onRound,
	}

	gBattleRooms[playerId] = br

	br.Join(playerId, sideInfo)

	return br
}

// 加入房间
func (br *BattleRoom) Join(playerId int64, sideInfo *battle_core.SideInfo) {
	var status RoomPlayerStatus
	if sideInfo != nil {
		status = RPS_SIT
	}
	br.players[playerId] = &BattleRoomPlayer{
		playerId: playerId,
		sideInfo: sideInfo,
		status:   status,
	}

	if status == RPS_SIT {
		br.battlePlayers = append(br.battlePlayers, br.players[playerId])
	}

	br.BroadcastJoinRoom(playerId)
}

// 离开房间
func (br *BattleRoom) Leave(playerId int64) {
	delete(br.players, playerId)

	// 当房间内的人都离开时， 关闭房间
	br.Close()

	br.BroadcaseLeaveRoom(playerId)
}

// 关闭房间
func (br *BattleRoom) Close() {

	br.BroadcastCloseRoom()
}

// 玩家准备
func (br *BattleRoom) Ready(playerId int64) {
	if player := br.players[playerId]; player != nil && player.status == RPS_SIT {
		player.status = RPS_READY
		br.BroadcastReady(playerId)
	}

	// 房间人准备的人满时，开始战斗
	if len(br.battlePlayers) == br.maxBattlePlayerNums {
		allReady := true
		for _, player := range br.battlePlayers {
			if player.status != RPS_READY {
				allReady = false
				break
			}
		}
		if allReady {
			br.StartBattle()
		}
	}
}

// 开始战斗
func (br *BattleRoom) StartBattle() {
	br.battle = battle_core.NewBattle(br.players[0].sideInfo, br.players[1].sideInfo, br.onRound)

	br.BroadcastStartBattle()
}

// 请求NextRound
func (br *BattleRoom) NextRound(nextRoundParam *NextRoundParam) {
	br.nextRoundParams = append(br.nextRoundParams, nextRoundParam)

}

// 执行doNextRound
func (br *BattleRoom) doNextRound() {
	if len(br.nextRoundParams) == 2 {
		br.BroadcastNextRound()
		br.nextRoundParams = br.nextRoundParams[0:0]
	}
}

// 广播 加入房间
func (br *BattleRoom) BroadcastJoinRoom(playerId int64) {

}

// 广播 退出房间
func (br *BattleRoom) BroadcaseLeaveRoom(playerId int64) {

}

// 广播 关闭房间
func (br *BattleRoom) BroadcastCloseRoom() {

}

// 广播 准备
func (br *BattleRoom) BroadcastReady(playerId int64) {

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
