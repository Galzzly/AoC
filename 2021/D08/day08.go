package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		ls := strings.Split(line, " | ")
		s := strings.Split(strings.TrimSpace(ls[1]), " ")
		for _, v := range s {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				res++
			}
		}
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		ls := strings.Split(line, " | ")
		patterns := getPatterns(strings.TrimSpace(ls[0]))
		var matched string
		s := strings.Split(strings.TrimSpace(ls[1]), " ")
		for _, v := range s {
			v = sortString(v)
			for k, p := range patterns {
				if v == p {
					matched += k
				}
			}
		}
		res += utils.Atoi(matched)
	}
	return
}

func getPatterns(line string) (res map[string]string) {
	res = make(map[string]string)
	s := strings.Split(line, " ")
	// map out the knowns first
	for _, v := range s {
		v = sortString(v)
		switch len(v) {
		// 1
		case 2:
			res["1"] = v
			// 7
		case 3:
			res["7"] = v
			// 4
		case 4:
			res["4"] = v
			// 8
		case 7:
			res["8"] = v
		}
	}
	for _, v := range s {
		v = sortString(v)
		switch len(v) {
		case 5:
			// 2
			if numContains(v, res["7"]) == 2 && numContains(v, res["1"]) == 1 && numContains(v, res["4"]) == 2 {
				res["2"] = v
			}
			// 3
			if numContains(v, res["1"]) == 2 && numContains(v, res["7"]) == 3 {
				res["3"] = v
			}
			// 5
			if numContains(v, res["4"]) == 3 && numContains(v, res["1"]) == 1 {
				res["5"] = v
			}

		case 6:
			// 9
			if numContains(v, res["7"]) == 3 && numContains(v, res["4"]) == 4 {
				res["9"] = v
			}
			// 6
			if numContains(v, res["1"]) == 1 && numContains(v, res["4"]) == 3 && numContains(v, res["8"]) == 6 {
				res["6"] = v
			}
			// 0
			if numContains(v, res["4"]) == 3 && numContains(v, res["8"]) == 6 && numContains(v, res["1"]) == 2 {
				res["0"] = v
			}
		}
	}
	return
}

func numContains(a, b string) (res int) {
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			res++
		}
	}
	return
}

func sortString(s string) (res string) {
	c := strings.Split(s, "")
	sort.Strings(c)
	res = strings.Join(c, "")
	return
}
