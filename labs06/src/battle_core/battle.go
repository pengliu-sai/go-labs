package battle_core

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type Side byte

const (
	AttackSide Side = 0
	DefendSide Side = 1
)

var globalBattleID uint64

type Battle struct {
	id            uint64 // 全局唯一ID
	rand          *rand.Rand
	Fighters      [18]Fighter
	actionSide    Side          // 当前轮到哪一方出手
	attackerIndex int           // 攻方轮到谁出手
	defenderIndex int           // 守方轮到谁出手
	autoFighterID int           // 自增的Fighter ID，用来判断同一个位置上的Fighter是否已经换人
	Sides         [2]*SideInfo  // 攻击方 = 0, 守方 = 1
	onRound       func(*Battle) // 每回合回调
	Round         Round
}

type SideInfo struct {
	GroupIndex    int           // 当前参与战斗的组索引
	MaxGroupIndex int           // 最大的组索引
	Groups        [][]*RoleData // 战斗小组
	Players       []int64       // 参与战斗的玩家ID
}

func NewBattle(attackSide, defendSide *SideInfo, onRound func(*Battle)) *Battle {
	battleID := atomic.AddUint64(&globalBattleID, 1)

	b := &Battle{
		id:      battleID,
		rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Sides:   [2]*SideInfo{attackSide, defendSide},
		onRound: onRound,
	}

	attackSide.GroupIndex = -1
	defendSide.GroupIndex = -1

	b.nextGroup(AttackSide)
	b.nextGroup(DefendSide)

	return b
}

// 加载并初始化下一批数据
func (b *Battle) nextGroup(side Side) {
	sideInfo := b.Sides[side]

	b.Sides[side].GroupIndex++
	group := sideInfo.Groups[sideInfo.GroupIndex]

	b.Round.FighterNum[side] = 0
	for _, role := range group {
		b.Fighters[role.Index].init(b, role)
		b.Round.FighterNum[side]++
	}

	// 重置出手顺序
	b.actionSide = AttackSide
	b.attackerIndex = 0
	b.defenderIndex = 1
}

func (b *Battle) ID() uint64 {
	return b.id
}

func (b *Battle) CanContinue() bool {
	if b.Round.State == RS_ATK_WIN && b.Sides[DefendSide].GroupIndex < b.Sides[DefendSide].MaxGroupIndex {
		return true
	} else if b.Round.State == RS_DEF_WIN && b.Sides[AttackSide].GroupIndex < b.Sides[AttackSide].MaxGroupIndex {
		return true
	}
	return false
}
