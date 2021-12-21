package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var statcache map[[5]int][2]int64

func main() {
	start := time.Now()
	p1, p2 := getPlayerStart(utils.ReadFileLineByLine("input"))
	t1 := time.Now()
	fmt.Printf("Part 1: %d in %s\n", part1(p1, p2, 1000), time.Since(t1))
	t2 := time.Now()
	statcache = map[[5]int][2]int64{}
	fmt.Printf("Part 2: %d in %s\n", max(dirac([2]int{p1, p2}, [2]int{0, 0}, 0)), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func part1(p1, p2 int, maxScore int) (res int) {
	p1score, p2score := 0, 0
	p1pos, p2pos := p1, p2
	dicecount, dierolls := 0, 0
	turn := 1
	var losing int
	for {
		var turnscore int
		for i := 0; i < 3; i++ {
			dicecount++
			dierolls++
			turnscore += dicecount
			if dicecount == 100 {
				dicecount = 0
			}
		}
		if turn%2 == 0 {
			p2pos = movePlayer(p2pos, turnscore)
			p2score += p2pos
			if p2score >= maxScore {
				losing = p1score
				goto finish
			}
		} else {
			p1pos = movePlayer(p1pos, turnscore)
			p1score += p1pos
			if p1score >= maxScore {
				losing = p2score
				goto finish
			}
		}
		turn++
	}
finish:
	res = losing * dierolls
	return
}

func dirac(pos [2]int, score [2]int, turn int) (res [2]int64) {
	if score[0] > 20 {
		return [2]int64{1, 0}
	}
	if score[1] > 20 {
		return [2]int64{0, 1}
	}

	iter := [5]int{pos[0], pos[1], score[0], score[1], turn}
	if v, ok := statcache[iter]; ok {
		return v
	}

	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				newPos := pos
				newScore := score
				newPos[turn] = movePlayer(pos[turn], (r1 + r2 + r3))
				newScore[turn] = score[turn] + newPos[turn]
				newTurn := 1
				if turn == 1 {
					newTurn = 0
				}
				win := dirac(newPos, newScore, newTurn)
				res[0] += win[0]
				res[1] += win[1]
			}
		}
	}
	statcache[iter] = res

	return
}

func movePlayer(pos, score int) int {
	for i := 0; i < score; i++ {
		if pos == 10 {
			pos = 0
		}
		pos++
	}
	return pos
}

func getPlayerStart(lines []string) (p1, p2 int) {
	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &p1)
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &p2)
	return
}

func max(in [2]int64) int64 {
	if in[0] > in[1] {
		return in[0]
	}
	return in[1]
}
