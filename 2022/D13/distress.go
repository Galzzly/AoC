package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileDoubleLine(f)
	i1, i2 := getInput(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(i1), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(i2), "- Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(input [][]any) (res int) {
	for k, pair := range input {
		if r := compare(pair[0], pair[1]); r >= 0 {
			res += k + 1
		}
	}
	return
}

func part2(input []any) (res int) {
	sort.Slice(input, func(i, j int) bool {
		a := input[i]
		b := input[j]
		return compare(a, b) == 1
	})
	res = 1
	for i, p := range input {
		if fmt.Sprintf("%v", p) == "[[[2]]]" || fmt.Sprintf("%v", p) == "[[[6]]]" {
			res *= i + 1
		}
	}
	return
}

func compare(l any, r any) int {
	lInt, lIsInt := l.(int)
	rInt, rIsInt := r.(int)
	if lIsInt && rIsInt {
		switch {
		case lInt < rInt:
			return 1
		case lInt > rInt:
			return -1
		default:
			return 0
		}
	}
	lList, lIsList := l.([]any)
	rList, rIsList := r.([]any)
	if !lIsList {
		lList = []any{lInt}
	}
	if !rIsList {
		rList = []any{rInt}
	}

	max := utils.Biggest(len(lList), len(rList))
	for i := 0; i < max; i++ {
		if i >= len(lList) {
			return 1
		}
		if i >= len(rList) {
			return -1
		}
		if sub := compare(lList[i], rList[i]); sub != 0 {
			return sub
		}
	}
	return 0
}

func getInput(lines []string) (i1 [][]any, i2 []any) {
	i1 = make([][]any, 0, len(lines))
	i2 = make([]any, 0, len(lines)*2)
	for _, line := range lines {
		s := strings.Split(line, "\n")
		pair := getPair(s)
		i2 = append(i2, pair...)
		i1 = append(i1, pair)
	}
	pair := getPair([]string{"[[2]]", "[[6]]"})
	i2 = append(i2, pair...)
	return
}

func getPair(lines []string) (pair []any) {
	pair = make([]any, 0, len(lines))
	for _, line := range lines {
		np, _ := parseLine(line)
		pair = append(pair, np)
	}
	return
}

func parseLine(line string) (any, int) {
	out := make([]any, 0)
	nC := make([]rune, 0)
	var i int
	for i = 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case '[':
			o, k := parseLine(line[i+1:])
			out = append(out, o)
			i += k
		case ']':
			if len(nC) > 0 {
				n := utils.Atoi(string(nC))
				out = append(out, n)
			}
			return out, i + 1
		case ',':
			if len(nC) > 0 {
				n := utils.Atoi(string(nC))
				out = append(out, n)
				nC = make([]rune, 0)
			}
		default:
			nC = append(nC, rune(c))
		}
	}
	return out, i
}
