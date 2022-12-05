package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/Galzzly/AoC/utils"
)

type rule struct {
	num      int
	from, to rune
}

func main() {
	start := time.Now()
	f := "input.txt"
	sects := utils.ReadFileDoubleLineNoTrim(f)
	s1, s2, k := getStack(sects[0])
	rules := getRules(sects[1])

	p1, p2 := solve(s1, s2, k, rules)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Time Taken:", time.Since(start))
}

func solve(s1, s2 map[rune]string, k string, rules []rule) (r1, r2 string) {
	for _, r := range rules {
		for i := 0; i < r.num; i++ {
			s1[r.to] += string(s1[r.from][len(s1[r.from])-1])
			s1[r.from] = s1[r.from][:len(s1[r.from])-1]
		}
		s2[r.to] += string(s2[r.from][len(s2[r.from])-r.num:])
		s2[r.from] = s2[r.from][:len(s2[r.from])-r.num]
	}
	for _, i := range k {
		r1 += string(s1[i][len(s1[i])-1])
		r2 += string(s2[i][len(s2[i])-1])
	}
	return
}

func getStack(stack string) (s1, s2 map[rune]string, k string) {
	crates := strings.Split(stack, "\n")
	keys := crates[len(crates)-1]
	s1, s2 = map[rune]string{}, map[rune]string{}
	k = strings.ReplaceAll(keys, " ", "")
	for i := len(crates) - 2; i >= 0; i-- {
		for j, c := range crates[i] {
			if unicode.IsLetter(c) {
				s1[[]rune(keys)[j]] += string(c)
				s2[[]rune(keys)[j]] += string(c)
			}
		}
	}
	return
}

func getRules(rules string) (res []rule) {
	for _, r := range strings.Split(strings.TrimSpace(rules), "\n") {
		var a int
		var b, c rune
		fmt.Sscanf(r, "move %d from %c to %c", &a, &b, &c)
		res = append(res, rule{a, b, c})
	}
	return
}
