package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Step struct {
	Pairs map[string]int
	Count map[string]int
}

type Out struct {
	P1   string
	P2   string
	NewC string
}

func main() {
	start := time.Now()
	lines := utils.ReadFileDoubleLine(os.Args[1])
	poly := step(lines[0])
	rules := makeRules(strings.Split((lines[1]), "\n"))
	t1 := time.Now()
	fmt.Printf("Part 1: %d, Took: %s\n", count(poly, rules, 10), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d, Took: %s\n", count(poly, rules, 40), time.Since(t2))
	fmt.Printf("Total: %s\n", time.Since(start))
}

func count(poly *Step, rules map[string]Out, steps int) (res int) {
	s := poly
	for i := 0; i < steps; i++ {
		s = next(s, rules)
	}

	min, max := math.MaxInt, 0
	for _, v := range s.Count {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	res = max - min
	return
}

func next(st *Step, rules map[string]Out) (res *Step) {
	res = &Step{
		Pairs: map[string]int{},
		Count: map[string]int{},
	}

	for k, v := range st.Count {
		res.Count[k] = v
	}

	for pair, count := range st.Pairs {
		t := rules[pair]
		res.Pairs[t.P1] += count
		res.Pairs[t.P2] += count
		res.Count[t.NewC] += count
	}
	return
}

func step(s string) (res *Step) {
	res = &Step{
		Pairs: map[string]int{},
		Count: map[string]int{},
	}

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		res.Pairs[pair]++
	}

	for _, c := range s {
		res.Count[string(c)]++
	}
	return
}

func makeRules(lines []string) (res map[string]Out) {
	res = map[string]Out{}
	for _, line := range lines {
		var in, out string
		fmt.Sscanf(line, "%s -> %s", &in, &out)
		res[in] = Out{
			P1:   string(in[0]) + out,
			P2:   out + string(in[1]),
			NewC: out,
		}
	}
	return
}
