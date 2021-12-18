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
	var x1, x2 int
	var y1, y2 int
	fmt.Sscanf(line, "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	t1 := time.Now()
	fmt.Printf("Part 1: %d in %s\n", -y1*(-y1-1)/2, time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d in %s\n", arc(x1, y1, x2, y2), time.Since(t2))
	fmt.Printf("Time: %s\n", time.Since(start))
}

func arc(x1, y1, x2, y2 int) (res int) {
	startX, endX := utils.MinMax([]int{0, x2})
	startY, endY := utils.MinMax([]int{-y1, y1})
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			peak := checkPeak(x, y, x1, y1, x2, y2)
			if peak != math.MinInt {
				res++
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
