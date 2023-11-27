package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Monkey struct {
	items []int
	op    func(int) int
	test  func(int) int
}

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileDoubleLine(f)
	monkeys, trans := getMonkeys(lines)
	newt1 := time.Now()
	fmt.Println(solve(monkeys, func(i int) int { return i / 3 }, 20), time.Since(newt1))
	newt2 := time.Now()
	fmt.Println(solve(monkeys, func(i int) int { return i % trans }, 10000), time.Since(newt2))
	fmt.Println("Took", time.Since(start))
}

func solve(monkeys []Monkey, trans func(i int) int, rounds int) (res int) {
	inspects := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].items {
				item = trans(monkeys[m].op(item))
				nm := monkeys[m].test(item)
				monkeys[nm].items = append(monkeys[nm].items, item)
				inspects[m]++
			}
			monkeys[m].items = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))
	res = inspects[0] * inspects[1]
	return
}

func getMonkeys(lines []string) (monkeys []Monkey, trans int) {
	monkeys = make([]Monkey, len(lines))
	trans = 1
	for m, line := range lines {
		l := strings.Split(line, "\n")
		monkeys[m].items = getItems(l[1])
		oper := strings.Fields((strings.ReplaceAll(strings.Split(l[2], " = ")[1], "* old", "^ 2")))
		switch oper[1] {
		case "+":
			monkeys[m].op = func(i int) int { return i + utils.Atoi(oper[2]) }
		case "*":
			monkeys[m].op = func(i int) int { return i * utils.Atoi(oper[2]) }
		case "^":
			monkeys[m].op = func(i int) int { return i * i }
		}
		test := utils.Atoi(strings.Fields(l[3])[3])
		monkeys[m].test = func(i int) int {
			if i%test == 0 {
				return utils.Atoi(strings.Fields(l[4])[5])
			}
			return utils.Atoi(strings.Fields(l[5])[5])
		}
		trans *= test
	}
	return
}

func getItems(line string) (res []int) {
	res = []int{}
	l := strings.Split(line, ": ")[1]
	for _, i := range strings.Split(l, ", ") {
		res = append(res, utils.Atoi(i))
	}
	return
}
