package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		s, _ := strconv.Unquote(line)
		res += len(line) - len(s)
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		s := strconv.Quote(line)
		res += len(s) - len(line)
	}
	return
}
