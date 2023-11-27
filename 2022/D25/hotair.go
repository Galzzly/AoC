package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var SNAFU = map[rune]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
var rSNAFU = []string{0: "=", 1: "-", 2: "0", 3: "1", 4: "2"}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(lines), "- Took:", time.Since(t1))
	fmt.Println("Total Time: ", time.Since(start))
}

func solve(lines []string) (res string) {
	output := 0
	for _, line := range lines {
		linelen := len(line)
		var out int
		for i := 1; i <= linelen; i++ {
			c := line[linelen-i]
			v := SNAFU[rune(c)]
			for j := 1; j < i; j++ {
				v *= 5
			}
			out += v
		}
		output += out
	}

	for output > 0 {
		res = rSNAFU[(output+2)%5] + res
		output = (output + 2) / 5
	}

	return res
}
