package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	t := utils.Atoi(lines[0])
	b := make(map[int]int)
	for i, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		n := utils.Atoi(v)
		b[i] = n
	}
	t1 := time.Now()
	fmt.Println("Part 1:", part1(t, b), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(t, b), "Took:", time.Since(t2))
	fmt.Println("Total Time: ", time.Since(start))
}

func part1(t int, b map[int]int) (res int) {
	minT, minB := 10000, -1
	for _, bus := range b {
		next := bus - (t % bus)
		if next < minT {
			minT = next
			minB = bus
		}
	}

	res = minT * minB
	return
}

func part2(t int, b map[int]int) (res int) {
	res, step := 0, 1
	for i, bus := range b {
		for (res+i)%bus != 0 {
			res += step
		}
		step *= bus
	}
	return
}
