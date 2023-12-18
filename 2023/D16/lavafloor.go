package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Pos struct {
	R, C, D int
}

type RC struct {
	R, C int
}

var (
	RowDir = []int{-1, 0, 1, 0}
	ColDir = []int{0, 1, 0, -1}
)

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	G := makeGrid(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(G, Pos{0, 0, 1}), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(G), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part2(G [][]rune) (res int) {
	for R := 0; R < len(G); R++ {
		res = utils.Biggest(res, solve(G, Pos{R, 0, 1}))
		res = utils.Biggest(res, solve(G, Pos{R, len(G[0]) - 1, 3}))
	}
	for C := 0; C < len(G[0]); C++ {
		res = utils.Biggest(res, solve(G, Pos{0, C, 2}))
		res = utils.Biggest(res, solve(G, Pos{len(G) - 1, C, 0}))
	}
	return
}

func solve(G [][]rune, pos Pos) int {
	POS := []Pos{pos}
	Seen := map[RC]bool{}
	Seen2 := map[Pos]bool{}
	for {
		NP := []Pos{}
		if len(POS) == 0 {
			break
		}
		for _, P := range POS {
			R, C := P.R, P.C
			if (0 <= R && R < len(G)) && (0 <= C && C < len(G[0])) {
				Seen[RC{R, C}] = true
				if _, ok := Seen2[P]; ok {
					continue
				}
				Seen2[P] = true
				switch G[R][C] {
				case '.':
					NP = append(NP, step(P))
				case '/':
					NP = append(NP, step(Pos{P.R, P.C, []int{1, 0, 3, 2}[P.D]}))
				case '\\':
					NP = append(NP, step(Pos{P.R, P.C, []int{3, 2, 1, 0}[P.D]}))
				case '|':
					if P.D == 0 || P.D == 2 {
						NP = append(NP, step(P))
						continue
					}
					NP = append(NP, step(Pos{P.R, P.C, 0}), step(Pos{P.R, P.C, 2}))
				case '-':
					if P.D == 1 || P.D == 3 {
						NP = append(NP, step(P))
						continue
					}
					NP = append(NP, step(Pos{P.R, P.C, 1}), step(Pos{P.R, P.C, 3}))
				}
			}
		}
		POS = NP
	}
	return len(Seen)
}

func step(P Pos) Pos {
	return Pos{P.R + RowDir[P.D], P.C + ColDir[P.D], P.D}
}

func makeGrid(lines []string) (res [][]rune) {
	res = make([][]rune, len(lines))
	for Y, line := range lines {
		res[Y] = []rune(line)
	}
	return
}
