package battle_core

type AttackEvent byte

const (
	AE_HIT      AttackEvent = 1  // 命中
	AE_BLOCK                = 2  // 格挡
	AE_CRITICAL             = 4  // 暴击
	AE_BEATBACK             = 8  // 反击
	AE_DODGE                = 16 // 闪避
)

type RoundState byte

const (
	RS_NONE    RoundState = 0
	RS_ATK_WIN            = 1
	RS_DEF_WIN            = 2
)

type Round struct {
	RoundNum      int        // 回合数
	FighterNum    [2]int     // 剩余人数
	State         RoundState // 战场状态，用于判断输赢
	AttackerIndex int        // 攻击者索引
	Targets       []*Target
}

type Target struct {
	DefenderIndex int         // 被攻击者位置索引
	Event         AttackEvent // 攻击时产生哪些事件，使用位运算判断
	Hurt          int         // 伤害值
	BeatBackHurt  int         // 如果被攻击者反击，则有反击伤害
	f             *Fighter    // 被攻击者
}

func (round *Round) Reset() {
	round.AttackerIndex = -1
	round.Targets = round.Targets[0:0]
	round.State = RS_NONE
}

func (b *Battle) NextRound() {
	b.Round.Reset()

	// 按出手顺序找到下一个出手的角色
	for b.Round.FighterNum[AttackSide] != 0 && b.Round.FighterNum[DefendSide] != 0 {
		if b.defenderIndex >= 18 && b.attackerIndex >= 18 {
			b.defenderIndex = 1
			b.attackerIndex = 0
			b.actionSide = 0
			b.Round.RoundNum++
		}

		var attacker *Fighter

		if b.actionSide == 1 {
			if b.defenderIndex < 18 {
				attacker = &b.Fighters[b.defenderIndex]
				b.defenderIndex += 2
			} else {
				b.actionSide = 0
				continue
			}
		} else {
			if b.attackerIndex < 18 {
				attacker = &b.Fighters[b.attackerIndex]
				b.attackerIndex += 2
			} else {
				b.actionSide = 1
				continue
			}
		}

		if attacker.Type == FT_NONE {
			continue
		}

		b.fight(attacker, 0)

		b.actionSide = (b.actionSide + 1) % 2
		return
	}
}

func (b *Battle) fight(attacker *Fighter, target int) {
	b.Round.AttackerIndex = attacker.Index

	defender := findSingleTarget(b, attacker)[0]

	// 事件判断
	var event AttackEvent = AE_HIT

	hurt := attacker.Attack - defender.Defence

	hurt = b.updateHealth(nil, defender, -hurt)

	b.Round.Targets = append(b.Round.Targets, &Target{
		DefenderIndex: defender.Index,
		Event:         event,
		Hurt:          hurt,
		BeatBackHurt:  0,
		f:             defender,
	})

	b.checkWinLose()
}

func (b *Battle) checkWinLose() {
	if b.Round.FighterNum[DefendSide] == 0 {
		b.Round.State = RS_ATK_WIN
	} else if b.Round.FighterNum[AttackSide] == 0 {
		b.Round.State = RS_DEF_WIN
	}
}

func (b *Battle) updateHealth(attacker *Fighter, defender *Fighter, changeValue int) int {
	if defender.Health == 0 || changeValue == 0 {
		return 0
	}

	newHealth := defender.Health + changeValue
	if newHealth > defender.rawHealth {
		changeValue = defender.rawHealth - defender.Health
	}

	defender.Health += changeValue

	if defender.Health == 0 {
		b.Round.FighterNum[defender.Index%2]--
	}

	return changeValue
}
