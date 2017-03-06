package battle_room

import "battle_core"

type RoomPlayerStatus byte

const (
	RPS_STAND RoomPlayerStatus = 0 // 旁观中
	RPS_SIT                    = 1 // 占座中
	RPS_READY                  = 2 // 准备中
)

type BattleRoomPlayer struct {
	playerId int64                 // 玩家ID
	sideInfo *battle_core.SideInfo // 角色信息
	status   RoomPlayerStatus      // 玩家状态
}
