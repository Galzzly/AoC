package main

import (
	"fmt"
	"image"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Grid map[image.Point][]string

var movement = map[rune]image.Point{
	'U': {0, -1},
	'R': {1, 0},
	'D': {0, 1},
	'L': {-1, 0},
}

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	motions := utils.ReadFileLineByLine(f)
	p1, p2 := solve(motions)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println(time.Since(start))
}

func solve(motions []string) (r1, r2 int) {
	tvisit := map[image.Point]bool{}
	tail := map[image.Point]bool{}
	rope := make([]image.Point, 10)

	for _, motion := range motions {
		var d rune
		var n int
		fmt.Sscanf(motion, "%c %d", &d, &n)
		for i := 0; i < n; i++ {
			rope[0] = rope[0].Add(movement[d])
			for i := 1; i < len(rope); i++ {
				if dir := rope[i-1].Sub(rope[i]); utils.Abs(dir.X) > 1 || utils.Abs(dir.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{seeker(dir.X), seeker(dir.Y)})
				}
			}
			tvisit[rope[1]] = true
			tail[rope[len(rope)-1]] = true
		}
	}
	r1 = len(tvisit)
	r2 = len(tail)
	return
}

func seeker(n int) (r int) {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}
