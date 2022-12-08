package main

import (
	"fmt"
	"image"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Treemap map[image.Point]int

func main() {
	start := time.Now()
	f := "input.txt"
	treemap := utils.MakeIntImagePoint(utils.ReadFileLineByLine(f))
	p1, p2 := solve(treemap)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Time taken:", time.Since(start))
}

func solve(treemap Treemap) (r1, r2 int) {
	for k, v := range treemap {
		vis, score := treemap.lookaround(k, v)
		r1 += vis
		if score > r2 {
			r2 = score
		}
	}
	return
}

func (t Treemap) lookaround(k image.Point, v int) (r1, r2 int) {
	r2 = 1
	for _, p := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		for np, i := k.Add(p), 0; ; np, i = np.Add(p), i+1 {
			if _, ok := t[np]; !ok {
				r1, r2 = 1, r2*i
				break
			}
			if t[np] >= v {
				r2 *= i + 1
				break
			}
		}
	}
	return
}
