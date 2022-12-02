package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var results = map[string][]int{
	"A X": {4, 3},
	"A Y": {8, 4},
	"A Z": {3, 8},
	"B X": {1, 1},
	"B Y": {5, 5},
	"B Z": {9, 9},
	"C X": {7, 2},
	"C Y": {2, 6},
	"C Z": {6, 7},
}

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)

	r1, r2 := game(lines)
	fmt.Println("Part 1:", r1)
	fmt.Println("Part 2:", r2)

	fmt.Println("Total Time:", time.Since(start))
}

func game(lines []string) (r1, r2 int) {
	for _, l := range lines {
		r1 += results[l][0]
		r2 += results[l][1]
	}
	return
}
