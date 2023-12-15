package main

import (
	"fmt"
	"strings"
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
		switch s[i : i+2] {
		case "ab", "cd", "pq", "xy":
			return false
		}
	}
	return true
}
