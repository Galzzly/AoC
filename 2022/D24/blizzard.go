package main

import (
	"fmt"
	"image"
	"time"

	"github.com/Galzzly/AoC/utils"
	"github.com/gammazero/deque"
)

var D = map[rune]image.Point{
	'#': {0, 0},
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
}

type Status struct {
	P image.Point
	T int
}

var valley map[image.Point]rune
var blizzard image.Rectangle

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"

	lines := utils.ReadFileLineByLine(f)
	valley = utils.MakeImagePointMap(lines)
	for p := range valley {
		blizzard = blizzard.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
	}
	blizzard.Min, blizzard.Max = blizzard.Min.Add(image.Point{1, 1}), blizzard.Max.Sub(image.Point{1, 1})
	startP := blizzard.Min.Sub(image.Point{0, 1})
	endP := blizzard.Max.Sub(image.Point{1, 0})
	t1 := time.Now()
	p1 := solve(startP, endP, 0)
	fmt.Println("Part 1:", p1, "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(startP, endP, solve(endP, startP, p1)), "- Took:", time.Since(t2))

	fmt.Println("Total Time: ", time.Since(start))
}

func solve(startP, endP image.Point, T int) int {
	var Q deque.Deque[Status]
	Q.PushBack(Status{startP, T})
	S := map[Status]bool{Q.Front(): true}
	for Q.Len() != 0 {
		current := Q.PopFront()
	path:
		for _, d := range D {
			next := Status{current.P.Add(d), current.T + 1}
			if next.P == endP {
				return next.T
			}

			if S[next] {
				continue
			}
			if r, ok := valley[next.P]; !ok || r == '#' {
				continue
			}

			if next.P.In(blizzard) {
				for r, d := range D {
					if valley[next.P.Sub(d.Mul(next.T)).Mod(blizzard)] == r {
						continue path
					}
				}
			}
			S[next] = true
			Q.PushBack(next)
		}
	}
	return -1
}
