package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	nums := utils.FileIntsLineByComma(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(nums), '('+time.Since(t1)+')')
	t2 := time.Now()
	fmt.Println("Part 2", part2(nums), '('+time.Since(t2)+')')
	fmt.Println("Total time: ", time.Since(start))
}

func part1(nums []int) (res int) {
	sort.Ints(nums)
	cost := []int{}
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		fuel := 0
		for _, n := range nums {
			fuel += diff(i, n)
		}
		cost = append(cost, fuel)
	}
	sort.Ints(cost)
	res = cost[0]
	return
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func part2(nums []int) (res int) {
	sort.Ints(nums)
	cost := []int{}
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		fuel := 0
		for _, n := range nums {
			steps := diff(i, n)
			fuel += int((steps * (steps + 1)) / 2)
		}
		cost = append(cost, fuel)
	}
	sort.Ints(cost)
	res = cost[0]
	return
}
