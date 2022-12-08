package main

import (
	"fmt"
	"os"
	"strings"

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
		if checkPair(line) && checkRepeat(line) {
			res++
		}
	}
	return
}

func checkRepeat(line string) bool {
	for i := range line[:len(line)-2] {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func checkPair(line string) bool {
	for i := range line[:len(line)-1] {
		if strings.Count(line, line[i:i+2]) > 1 {

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
