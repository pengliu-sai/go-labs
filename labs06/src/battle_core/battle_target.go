package battle_core

/*
	12 6  0  |  1 7  13
	14 8  2  |  3 9  15
	16 10 4  |  5 11 17
*/

// 单体攻击索引，角色站位%6以后进行索引
var singleTargetIndex = [][]int{
	0: []int{1, 3, 5, 7, 9, 11, 13, 15, 17}, // 0 6 12
	1: []int{0, 2, 4, 6, 8, 10, 12, 14, 16}, // 1 7 13
	2: []int{3, 5, 1, 9, 11, 7, 15, 17, 13}, // 2 8 14
	3: []int{2, 0, 4, 8, 6, 10, 14, 12, 16}, // 3 9 15
	4: []int{5, 3, 1, 11, 9, 7, 17, 15, 13}, // 4 10 16
	5: []int{4, 2, 0, 10, 8, 6, 16, 14, 12}, // 5 11 17
}

func findSingleTargetBase(b *Battle, f *Fighter) *Fighter {
	for _, i := range singleTargetIndex[f.Index%6] {
		if b.Fighters[i].canBeAttacked() {
			return &b.Fighters[i]
		}
	}
	return nil
}

func findSingleTarget(b *Battle, f *Fighter) []*Fighter {
	target := findSingleTargetBase(b, f)
	if target == nil {
		return nil
	}
	return []*Fighter{target}
}
