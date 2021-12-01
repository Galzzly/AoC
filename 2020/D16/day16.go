package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Arg[1]
	lines := utils.ReadFileDoubleLine(f)
	rules := mapRules(lines[0])
	myTx := getTx(lines[1])
	nearbyTx := getTx(lines[2])
	t1 := time.Now()
	p1, valid := part1(nearbyTx, rules)
	fmt.Println("Part 1", p1, "Took:", time.Since(t1))
	fmt.Println("Total Time: ", time.Since(start))
}

func part1(nearbyTx [][]int, rules map[string]int) (res int, valid [][]int) {

	return
}

func mapRules(lines string) (rules map[string][]int) {
	for _, v := range strings.Split(lines, "\n") {
		r := strings.Split(v, ": ")
		rules[r[0]] = make([]int, 4)
		fmt.Sscanf(r[1], "%d-%d or %d-%d", &rules[r[0]][0], &rules[r[0]][1], &rules[r[0]][2], &rules[r[0]][3])
	}
	return
}

func getTx(lines string) (tx [][]int) {
	for _, s := range strings.Split(strings.TrimSpace(lines), "\n")[1:] {
		nums := []int{}
		for _, v := range strings.Split(s, ",") {
			nums = append(nums, utils.Atoi(v))
		}
		tx = append(tx, nums)
	}
	return
}
