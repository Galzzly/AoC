package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	line := utils.ReadFileSingleLine("input")
	fmt.Println(line)
	var x1, x2 int
	var y1, y2 int
	fmt.Sscanf(line, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	fmt.Printf("%d,%d && %d,%d\n", x1, y1, x2, y2)
	r1, r2 := arc(x1, y1, x2, y2)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", r1, r2)
	fmt.Printf("Time: %s\n", time.Since(start))
}

func arc(x1, y1, x2, y2 int) (p1res, p2res int) {
	p1res = math.MinInt
	startX, endX := utils.MinMax([]int{0, x2})
	startY, endY := utils.MinMax([]int{-y1, y1})
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			peak := checkPeak(x, y, x1, y1, x2, y2)
			if peak != math.MinInt {
				p2res++
			}
			if peak > p1res {
				p1res = peak
			}
		}
	}
	return
}

func checkPeak(px, py, x1, y1, x2, y2 int) (res int) {
	x, y := 0, 0
	res = math.MinInt
	for {
		if y > res {
			res = y
		}
		if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
			return
		}
		x += px
		y += py
		if px < 0 {
			px++
		} else if px > 0 {
			px--
		}
		py--

		if py < 0 && y < y1 ||
			px == 0 && (x < x1 || x > x2) {
			return math.MinInt
		}
	}
}
