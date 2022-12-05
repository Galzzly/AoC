package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	var p1, p2 int
	for _, line := range lines {
		var e1start, e1end, e2start, e2end int
		fmt.Sscanf(line, "%d-%d,%d-%d", &e1start, &e1end, &e2start, &e2end)
		if e2start >= e1start && e2end <= e1end || e1start >= e2start && e1end <= e2end {
			p1++
		}
		if e2start <= e1end && e2end >= e1start || e1start <= e2end && e1end >= e2end {
			p2++
		}
	}
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println(time.Since(start))
}
