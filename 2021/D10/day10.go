package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var (
	opening = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	closing = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	scoring = map[rune]int{
		'(': 1,
		')': 3,
		'[': 2,
		']': 57,
		'{': 3,
		'}': 1197,
		'<': 4,
		'>': 25137,
	}
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	r1, r2 := both(lines)
	fmt.Println("Part 1:", r1)
	fmt.Println("Part 2:", r2)
	fmt.Println("Time Taken:", time.Since(start))
}

func both(lines []string) (r1, r2 int) {
	r2res := []int{}
	for _, line := range lines {
		char := []rune{}
		valid := true
		for _, c := range line {
			if _, ok := opening[c]; ok {
				char = append(char, c)
			}
			if _, ok := closing[c]; ok {
				last := char[len(char)-1]
				if last != closing[c] {
					r1 += scoring[c]
					valid = false
					break
				}
				char = char[:len(char)-1]
			}
		}
		if !valid {
			continue
		}
		score := 0
		for i := len(char) - 1; i >= 0; i-- {
			s := char[i]
			score *= 5
			score += scoring[s]
		}
		r2res = append(r2res, score)
	}
	sort.Ints(r2res)
	r2 = r2res[len(r2res)/2]
	return
}
