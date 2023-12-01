package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/Galzzly/AoC/utils"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Digits [2]digit

type digit struct {
	id  int
	num string
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	calcstart := time.Now()
	p1, p2 := calculate(lines)
	calctime := time.Since(calcstart)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Calc time:", calctime)
	fmt.Println("Total time:", time.Since(start))
}

func calculate(lines []string) (r1, r2 int) {
	for _, line := range lines {
		digits := findNums(line)
		r1 += utils.Atoi(digits[0].num + digits[1].num)

		words := findWords(line)
		if digits[0].id == -1 || (words[0].id != -1 && words[0].id < digits[0].id) {
			digits[0] = words[0]
		}
		if words[1].id != -1 && words[1].id > digits[1].id {
			digits[1] = words[1]
		}
		r2 += utils.Atoi(digits[0].num + digits[1].num)
	}
	return
}

func findNums(line string) Digits {
	result := [2]digit{{-1, ""}, {-1, ""}}
	for i, c := range line {
		if unicode.IsDigit(c) {
			if result[0].id == -1 {
				result[0] = digit{i, string(c)}
				result[1] = digit{i, string(c)}
				continue
			}
			result[1] = digit{i, string(c)}
		}
	}
	return result
}

func findWords(line string) Digits {
	result := [2]digit{{-1, ""}, {-1, ""}}
	charMap := make(map[int]string, 0)
	var charIdx []int

	for word, num := range digitMap {
		first := strings.Index(line, word)
		last := strings.LastIndex(line, word)
		if first != -1 {
			charIdx = append(charIdx, first)
			charMap[first] = num
		}
		if last != -1 {
			charIdx = append(charIdx, last)
			charMap[last] = num
		}
	}

	if len(charIdx) > 0 {
		sort.Ints(charIdx)
		result[0] = digit{charIdx[0], charMap[charIdx[0]]}
		result[1] = digit{charIdx[len(charIdx)-1], charMap[charIdx[len(charIdx)-1]]}
	}
	return result
}
