package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	instructions := utils.ReadFileLineByLine("prestwood")
	r1, r2 := checkMonad(instructions)
	fmt.Printf("Part 1: %d\n", r1)
	fmt.Printf("Part 2: %d\n", r2)
	fmt.Printf("Total Time: %s", time.Since(start))
}

var Monad []int64

func checkMonad(instructions []string) (r1, r2 int) {
	p1 := make([]int, 14)
	p2 := make([]int, 14)
	p3 := make([]int, 14)
	for i := 0; i < 14; i++ {
		p1[i] = utils.Atoi(instructions[i*18+4][6:])
		p2[i] = utils.Atoi(instructions[i*18+5][6:])
		p3[i] = utils.Atoi(instructions[i*18+15][6:])
	}

	pairs := []utils.Point{}
	for len(pairs) < 7 {
		x := -1
		for i, p := range p1 {
			if p == 1 {
				x = i
			} else if p == 26 {
				pairs = append(pairs, utils.Point{x, i})
				p1[x] = 2
				p1[i] = 27
				break
			}
		}
	}

	mmax := make([]int, 14)
	mmin := make([]int, 14)
	for _, pr := range pairs {
		diff := p3[pr.X] + p2[pr.Y]
		if diff > 0 {
			mmax[pr.X] = 9 - diff
			mmax[pr.Y] = 9
			mmin[pr.X] = 1
			mmin[pr.Y] = 1 + diff
		} else {
			mmax[pr.X] = 9
			mmax[pr.Y] = 9 + diff
			mmin[pr.X] = 1 - diff
			mmin[pr.Y] = 1
		}
	}

	for i := 0; i < 14; i++ {
		r1 *= 10
		r2 *= 10
		r1 += mmax[i]
		r2 += mmin[i]
	}
	return
}
