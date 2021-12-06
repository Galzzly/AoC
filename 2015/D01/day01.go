package main

import (
	"fmt"
	"os"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	line := utils.ReadFileLineByLine(f)
	fmt.Println(part1(line[0]))
	fmt.Println(part2(line[0]))
}

func part1(s string) (res int) {
	for _, c := range s {
		if c == '(' {
			res++
		} else if c == ')' {
			res--
		}
	}
	return
}

func part2(s string) (res int) {
	floor := 0
	for i, c := range s {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 {
			res = i + 1
			break
		}
	}
	return
}
