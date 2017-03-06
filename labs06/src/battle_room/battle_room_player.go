package battle_room

import "battle_core"

type RoomPlayerStatus byte

const (
	RPS_HOLD  RoomPlayerStatus = 0 // 准备中
	RPS_LOOK                   = 1 // 旁观中
	RPS_READY                  = 2 // 占位中
)

type BattleRoomPlayer struct {
	playerId int64                 // 玩家ID
	sideInfo *battle_core.SideInfo // 角色信息
	status   RoomPlayerStatus      // 玩家状态
}
