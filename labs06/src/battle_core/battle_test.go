package battle_core

import (
	"fmt"
	"testing"
)

func (r *Round) TextReport() string {
	var s string
	s = fmt.Sprintf("出手方：%d，技能：普攻", r.AttackerIndex)
	for _, target := range r.Targets {
		if target.Event|AE_HIT == AE_HIT {
			s = fmt.Sprintf("%s，{命中：%d，伤害：%d}", s, target.DefenderIndex, target.Hurt)
		}
	}
	return s
}

func Test_OneVsOne(t *testing.T) {
	fighters := make([]Fighter, 18)

	fighters[0].Health = 200
	fighters[0].Attack = 100
	fighters[0].Defence = 50
	fighters[0].Type = FT_PLAYER
	fighters[0].Index = 8
	fighters[0].rawHealth = fighters[0].Health

	fighters[1].Health = 200
	fighters[1].Attack = 150
	fighters[1].Defence = 50
	fighters[1].Type = FT_PLAYER
	fighters[1].Index = 9
	fighters[1].rawHealth = fighters[1].Health

	b := &Battle{}
	b.attackerIndex = 0
	b.defenderIndex = 1

	b.Fighters[0] = fighters[0]
	b.Fighters[1] = fighters[1]

	b.Round.FighterNum[0] = 1
	b.Round.FighterNum[1] = 1

	for {
		b.NextRound()
		t.Log(b.Round.TextReport())
		if b.Round.State != RS_NONE {
			break
		}
	}
}
