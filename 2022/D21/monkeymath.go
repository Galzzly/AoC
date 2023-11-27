package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var monkeys map[string][]string

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	p1, p2, t1, t2 := solve(lines)
	fmt.Println("Part 1:", p1, "- Took:", t1)
	fmt.Println("Part 2:", p2, "- Took:", t2)
	fmt.Println("Total Time: ", time.Since(start))
}

func solve(lines []string) (r1, r2 int, t1, t2 time.Duration) {
	monkeys = make(map[string][]string, len(lines))
	for _, line := range lines {
		s := strings.Split(line, ": ")
		monkeys[s[0]] = append(monkeys[s[0]], strings.Split(s[1], " ")...)
	}
	// Part 1
	s1 := time.Now()
	r1, t1 = int(solver("root", float64(-1))), time.Since(s1)

	// Part 2
	s2 := time.Now()
	// Need to go down each path connected to root
	p1 := monkeys["root"][0]
	p2 := monkeys["root"][2]

	// Work out which one contains the human, want it to be p1
	if solver(p2, 0) != solver(p2, 1) {
		p1, p2 = p2, p1
	}

	// Check that this is the case, again
	if solver(p1, 0) == solver(p1, 1) || solver(p2, 0) != solver(p2, 1) {
		fmt.Println("Values aren't right")
		return
	}

	// P2 is correct, so need to get the target here.
	target := solver(p2, 0) // 28379346560301
	// Set our starting limits
	var min, max float64
	min, max = 0, 1e20
	for min < max {
		mid := math.Floor((min + max) / 2)
		score := target - solver(p1, mid)
		if score < 0 {
			min = mid
		} else if score == 0 {
			r2, t2 = int(mid), time.Since(s2)
			break
		} else {
			max = mid
		}
	}
	return
}

func solver(monkey string, n float64) (res float64) {
	words := monkeys[monkey]
	if monkey == "humn" && n >= 0 {
		return n
	}
	if len(words) == 1 {
		return float64(utils.Atoi(words[0]))
	}
	e1 := solver(words[0], n)
	e2 := solver(words[2], n)
	switch words[1] {
	case "+":
		res = e1 + e2
	case "-":
		res = e1 - e2
	case "*":
		res = e1 * e2
	case "/":
		res = e1 / e2
	}
	return
}
