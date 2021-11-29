package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var bags = make(map[string]map[string]int)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	bags = getBags(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1() (res int) {
	for bag := range bags {
		if findGoldBag(bag) {
			res++
		}
	}
	return
}

func part2() (res int) {
	return countBag("shiny gold")
}

func getBags(lines []string) map[string]map[string]int {
	for _, line := range lines {
		s := strings.Split(line, "bags contain ")
		bag := strings.TrimSpace(s[0])
		bags[bag] = map[string]int{}
		for _, c := range strings.Split(s[1], ", ") {
			if c == "no other bags." {
				continue
			}
			p := strings.Split(c, " ")
			bags[bag][p[1]+" "+p[2]] = utils.Atoi(p[0])
		}

	}
	return bags
}

func findGoldBag(bag string) bool {
	if _, found := bags[bag]["shiny gold"]; found {
		return true
	}
	for k := range bags[bag] {
		if findGoldBag(k) {
			return true
		}
	}

	return false
}

func countBag(bag string) (res int) {
	res = 1

	for b, count := range bags[bag] {
		res += count * countBag(b)
	}
	return
}
