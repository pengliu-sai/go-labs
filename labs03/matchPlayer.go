package labs03

import "sort"

var (
	roundPlayers     []*matchPlayer
	idleTimesFactors []float64
)

type matchPlayer struct {
	playerID  int  // 玩家ID
	score     int  // 积分
	idleTimes int  // 轮空次数
	isRobot   bool // 机器人
}

type matchPlayerSlice []*matchPlayer

type matchPlayerSliceBy func(p1, p2 *matchPlayer) bool

func (by matchPlayerSliceBy) Sort(slice matchPlayerSlice) {
	sorter := &matchPlayerSliceSorter{
		Slice: slice,
		By:    by,
	}
	sort.Sort(sorter)
}

type matchPlayerSliceSorter struct {
	Slice matchPlayerSlice
	By    matchPlayerSliceBy
}

func (s *matchPlayerSliceSorter) Len() int {
	return len(s.Slice)
}

func (s *matchPlayerSliceSorter) Swap(i, j int) {
	s.Slice[i], s.Slice[j] = s.Slice[j], s.Slice[i]
}

func (s *matchPlayerSliceSorter) Less(i, j int) bool {
	return s.By(s.Slice[i], s.Slice[j])
}

func matchOnceRound(matchCallBack func(player1, player2 *matchPlayer)) {
	matchNextPlayer := func(currIndex, maxIndex int, curr *matchPlayer, matchCallBack func(player1, player2 *matchPlayer)) (succ bool) {
		nextIndex := currIndex + 1
		if nextIndex <= maxIndex {
			next := roundPlayers[nextIndex]
			succ = isMatch(curr, next)
			if succ {
				roundPlayers = append(roundPlayers[:currIndex], roundPlayers[currIndex+2:]...)
				matchCallBack(curr, next)
			}
		}
		return succ
	}

	matchPrevPlayer := func(currIndex int, curr *matchPlayer, matchCallBack func(player1, player2 *matchPlayer)) (succ bool) {
		prevIndex := currIndex - 1
		if prevIndex >= 0 {
			prev := roundPlayers[prevIndex]
			succ = isMatch(curr, prev)
			if succ {
				roundPlayers = append(roundPlayers[:currIndex-1], roundPlayers[currIndex+1:]...)
				matchCallBack(curr, prev)
			}
		}
		return succ
	}

	var curr *matchPlayer
	var currIndex int

	for {
		curr = nil
		succ := false
		maxIndex := len(roundPlayers) - 1

		if currIndex > maxIndex {
			break
		}

		curr = roundPlayers[currIndex]

		if curr.idleTimes == 2 || curr.idleTimes == 3 {
			succ = matchNextPlayer(currIndex, maxIndex, curr, matchCallBack)
			if !succ {
				succ = matchPrevPlayer(currIndex, curr, matchCallBack)

				if !succ {
					succ = curr.playerID%2 == 0
					if succ {

						roundPlayers = append(roundPlayers[:currIndex], roundPlayers[currIndex+1:]...)
						// matchCallBack(curPlayer, &matchPlayer{playerID: 99999, score: 99999})

					} else {
						curr.idleTimes++
						currIndex++
					}
				} else {
					currIndex--
				}
			}
		} else if curr.idleTimes == 4 {
			roundPlayers = append(roundPlayers[:currIndex], roundPlayers[currIndex+1:]...)
		} else {
			succ = matchNextPlayer(currIndex, maxIndex, curr, matchCallBack)
			if !succ {
				succ = matchPrevPlayer(currIndex, curr, matchCallBack)
				if !succ {
					curr.idleTimes++
					currIndex++
				} else {
					currIndex--
				}
			}
		}
	}
}

func isMatch(p1, p2 *matchPlayer) bool {
	p1IdleTimes := p1.idleTimes
	p2IdleTimes := p2.idleTimes

	if !p2.isRobot {
		if p2IdleTimes > 1 {
			p2IdleTimes = 1
		}
		if p1IdleTimes > 1 {
			p1IdleTimes = 1
		}
	}

	p1DiffScore := fixDiffScore(p1IdleTimes, int(float64(p1.score)*idleTimesFactors[p1IdleTimes]))
	p1Min := p1.score - p1DiffScore
	p1Max := p1.score + p1DiffScore

	p2DiffScore := fixDiffScore(p2IdleTimes, int(float64(p2.score)*idleTimesFactors[p2IdleTimes]))
	p2Min := p2.score - p2DiffScore
	p2Max := p2.score + p2DiffScore

	return p1.score >= p2Min && p1.score <= p2Max && p2.score >= p1Min && p2.score <= p1Max
}

func fixDiffScore(idleTimes, score int) int {
	switch idleTimes {
	case 0:
		if score < 50 {
			score = 50
		}
	case 1:
		if score < 90 {
			score = 90
		}
	case 2:
		if score < 70 {
			score = 70
		}
	case 4:
		if score < 150 {
			score = 150
		}
	}
	return score
}
