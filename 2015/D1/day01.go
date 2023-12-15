package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	line := utils.ReadFileLineByLine(f)[0]
	p1, p2 := solve(line)
	fmt.Printf("Part 1: %d\nPart 2: %d\nTotal Time: %s", p1, p2, time.Since(start))
}

func solve(line string) (r1, r2 int) {
	for i, c := range line {
		if c == '(' {
			r1++
		} else if c == ')' {
			r1--
		}
		if r1 == -1 && r2 == 0 {
			r2 = i + 1
		}
	}
	return
}
