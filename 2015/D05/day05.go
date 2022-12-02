package main

import (
	"fmt"
	"os"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		if checkDouble(line) && checkVowel(line) && checkBad(line) {
			res++
		}
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		if checkTwoPair(line) {
			res++
		}
	}
}

func checkTwoPair(s string) bool {
	if len(s) < 4 {
		return false
	}
	for i := 0; i < len(s)-2; i++ {
		pair := s[i : i+2]
		var c int

		if c == 2 {
			return true
		}
	}
	return false
}

func checkDouble(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func checkVowel(s string) bool {
	var count int
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count > 2
}

func checkBad(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		// fmt.Println(s[i : i+2])
		switch s[i : i+2] {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}
	return true
}
