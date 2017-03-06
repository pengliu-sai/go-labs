package battle_core

type FighterType byte

const (
	FT_NONE    FighterType = 0 // 无
	FT_PLAYER              = 1 // 玩家角色
	FT_MONSTER             = 2 // 怪物角色
	FT_BUDDY               = 3 // 伙伴
	FT_SUMMON              = 4 // 召唤物
)

type Fighter struct {
	ID int // 自增ID

	RoleData // 角色数据

	rawHealth  int // （原值）生命
	rawAttack  int // （原值）攻击
	rawDefence int // （原值）防御
}

type RoleData struct {
	Index     int         // 站位
	Type      FighterType // 角色类型
	RoleID    int         // 角色ID
	RoleJob   int8        // 角色职业
	RoleLevel int         // 角色等级

	Power    int // 精气
	MaxPower int // 最大精气

	Health  int // 生命
	Attack  int // 攻击
	Defence int // 防御
}

func (f *Fighter) init(b *Battle, roleData *RoleData) {
	*f = Fighter{
		RoleData: *roleData,
	}

	b.autoFighterID++
	f.ID = b.autoFighterID

	f.rawHealth = f.Health
	f.rawAttack = f.Attack
	f.rawDefence = f.Defence
}

func (f *Fighter) canBeAttacked() bool {
	return f.Health > 0
}
