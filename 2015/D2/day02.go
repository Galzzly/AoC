package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	p1, p2 := solve(lines)
	fmt.Printf("Part 1: %d\nPart 2: %d\nTotal Time: %s", p1, p2, time.Since(start))
}

func solve(lines []string) (r1, r2 int) {
	for _, line := range lines {
		sides := []int{}
		for _, c := range strings.Split(strings.TrimSpace(line), "x") {
			sides = append(sides, utils.Atoi(c))
		}
		LW := sides[0] * sides[1]
		WH := sides[1] * sides[2]
		HL := sides[2] * sides[0]
		r1 += (2*(LW) + 2*(WH) + 2*(HL)) + int(math.Min(math.Min(float64(LW), float64(WH)), float64(HL)))
		slices.Sort(sides)
		r2 += ((2 * sides[0]) * (2 * sides[1])) + (utils.SumArray(sides))
	}
	return
}
