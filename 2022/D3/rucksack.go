package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var values = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9,
	'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17,
	'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25,
	'z': 26,
	'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35,
	'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43,
	'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51,
	'Z': 52,
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), ", Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), ", Took:", time.Since(t2))
	fmt.Println(time.Since(start))
}

func part1(lines []string) (res int) {
bag:
	for _, r := range getRucksacks(lines) {
		for _, c := range r[0] {
			if strings.ContainsRune(r[1], c) {
				res += values[c]
				continue bag
			}
		}
	}
	return
}

func part2(rucksacks []string) (res int) {
bags:
	for i := 0; i < len(rucksacks); i += 3 {
		for _, c := range rucksacks[i] {
			if strings.ContainsRune(rucksacks[i+1], c) && strings.ContainsRune(rucksacks[i+2], c) {
				res += values[c]
				continue bags
			}
		}
	}
	return
}

func getRucksacks(lines []string) (res [][]string) {
	for _, line := range lines {
		l := len(line) / 2
		res = append(res, []string{line[:l], line[l:]})
	}
	return
}
