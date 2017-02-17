package labs03

import (
	"math/rand"
	"testing"
)

var (
	okPlayers [][2]*matchPlayer
)

func init() {
	idleTimesFactors = make([]float64, 4)
	idleTimesFactors[0] = 0.1
	idleTimesFactors[1] = 0.15
	idleTimesFactors[2] = 0.1
	idleTimesFactors[3] = 0.2
}

func addPlayer(playerID int, score int) {
	roundPlayers = append(roundPlayers, &matchPlayer{playerID: playerID, score: score})
}

func initRandData() {
	roundPlayers = roundPlayers[0:0]

	for i := 0; i < 10000; i++ {
		addPlayer(i, rand.Intn(2000-1200)+1200)
	}

	f := func(p1, p2 *matchPlayer) bool {
		return p1.score < p2.score
	}
	matchPlayerSliceBy(f).Sort(roundPlayers)

	times := 10

	for times > 0 {
		matchOnceRound(func(player1, player2 *matchPlayer) {
			okPlayers = append(okPlayers, [2]*matchPlayer{
				player1, player2,
			})
		})
		times--
	}
}

var readyPlayer = []*matchPlayer{
	&matchPlayer{playerID: 1, score: 0},
	&matchPlayer{playerID: 2, score: 0},
	&matchPlayer{playerID: 3, score: 20},
	&matchPlayer{playerID: 4, score: 50},
	&matchPlayer{playerID: 5, score: 80},
	&matchPlayer{playerID: 6, score: 100},
	&matchPlayer{playerID: 7, score: 150},
	&matchPlayer{playerID: 8, score: 200},
	&matchPlayer{playerID: 9, score: 220},
	&matchPlayer{playerID: 10, score: 300},
	&matchPlayer{playerID: 11, score: 330},
	&matchPlayer{playerID: 12, score: 380},
	&matchPlayer{playerID: 13, score: 420},
	&matchPlayer{playerID: 14, score: 500},
	&matchPlayer{playerID: 15, score: 550},
	&matchPlayer{playerID: 16, score: 600},
	&matchPlayer{playerID: 17, score: 800},
	&matchPlayer{playerID: 18, score: 1000},
	&matchPlayer{playerID: 19, score: 1400},
	&matchPlayer{playerID: 20, score: 1800},
}

var matchDatas = [][]*matchPlayer{
	{
		&matchPlayer{playerID: 9, score: 220, idleTimes: 1},
		&matchPlayer{playerID: 16, score: 600, idleTimes: 1},
		&matchPlayer{playerID: 17, score: 800, idleTimes: 1},
		&matchPlayer{playerID: 18, score: 1000, idleTimes: 1},
		&matchPlayer{playerID: 19, score: 1400, idleTimes: 1},
		&matchPlayer{playerID: 20, score: 1800, idleTimes: 1},
	}, {
		&matchPlayer{playerID: 9, score: 220, idleTimes: 2},
		&matchPlayer{playerID: 16, score: 600, idleTimes: 2},
		&matchPlayer{playerID: 17, score: 800, idleTimes: 2},
		&matchPlayer{playerID: 18, score: 1000, idleTimes: 2},
	}, {
		&matchPlayer{playerID: 9, score: 220, idleTimes: 3},
		&matchPlayer{playerID: 23, score: 300, idleTimes: 1},
		&matchPlayer{playerID: 17, score: 800, idleTimes: 3},
	}, {
		&matchPlayer{playerID: 17, score: 800, idleTimes: 4},
	},
}

func initNormalData() {
	roundPlayers = roundPlayers[0:0]

	for _, readyPlayer := range readyPlayer {
		addPlayer(readyPlayer.playerID, readyPlayer.score)
	}

}

func Test_okPlayer(t *testing.T) {
	initRandData()
	var ids []int

	for _, v := range okPlayers {
		p1 := v[0]
		p2 := v[1]

		matchSucc := isMatch(p1, p2)

		idRepeat := false
		for _, id := range ids {
			if id == p1.playerID || id == p2.playerID || p1.playerID == p2.playerID {
				t.FailNow()
			}
		}

		if !idRepeat {
			ids = append(ids, p1.playerID, p2.playerID)
		}

		if !matchSucc {
			t.FailNow()
		}
	}
}

func Test_NormalData(t *testing.T) {
	initNormalData()
	okPlayers = okPlayers[0:0]

	times := 0

	//-----------第一次匹配-----------------
	matchOnceRound(func(player1, player2 *matchPlayer) {
		okPlayers = append(okPlayers, [2]*matchPlayer{
			player1, player2,
		})
	})

	if len(roundPlayers) != 6 {
		t.FailNow()
	}

	for i, roundPlayer := range roundPlayers {
		if matchDatas[times][i].playerID != roundPlayer.playerID || matchDatas[times][i].score != roundPlayer.score || matchDatas[times][i].idleTimes != roundPlayer.idleTimes {
			t.FailNow()
		}
	}

	addPlayer(21, 1300)
	addPlayer(22, 1600)

	f := func(p1, p2 *matchPlayer) bool {
		return p1.score < p2.score
	}
	matchPlayerSliceBy(f).Sort(roundPlayers)

	times++
	//-----------第二次匹配-----------------
	matchOnceRound(func(player1, player2 *matchPlayer) {
		okPlayers = append(okPlayers, [2]*matchPlayer{
			player1, player2,
		})
	})

	if len(roundPlayers) != 4 {
		t.FailNow()
	}
	for i, roundPlayer := range roundPlayers {
		if matchDatas[times][i].playerID != roundPlayer.playerID || matchDatas[times][i].score != roundPlayer.score || matchDatas[times][i].idleTimes != roundPlayer.idleTimes {
			t.FailNow()
		}
	}

	addPlayer(23, 300)
	addPlayer(24, 600)

	matchPlayerSliceBy(f).Sort(roundPlayers)

	times++

	//-----------第三次匹配-----------------
	matchOnceRound(func(player1, player2 *matchPlayer) {
		okPlayers = append(okPlayers, [2]*matchPlayer{
			player1, player2,
		})
	})

	if len(roundPlayers) != 3 {
		t.FailNow()
	}
	for i, roundPlayer := range roundPlayers {
		if matchDatas[times][i].playerID != roundPlayer.playerID || matchDatas[times][i].score != roundPlayer.score || matchDatas[times][i].idleTimes != roundPlayer.idleTimes {
			t.FailNow()
		}
	}

	times++
	//----------第四次匹配------------------
	matchOnceRound(func(player1, player2 *matchPlayer) {
		okPlayers = append(okPlayers, [2]*matchPlayer{
			player1, player2,
		})
	})

	for i, roundPlayer := range roundPlayers {
		if matchDatas[times][i].playerID != roundPlayer.playerID || matchDatas[times][i].score != roundPlayer.score || matchDatas[times][i].idleTimes != roundPlayer.idleTimes {
			t.FailNow()
		}
	}

	times++

	//----------第五次匹配------------------
	matchOnceRound(func(player1, player2 *matchPlayer) {
		okPlayers = append(okPlayers, [2]*matchPlayer{
			player1, player2,
		})
	})

	if len(roundPlayers) != 0 {
		t.FailNow()
	}
}
