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
	sort.Ints(nums)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(nums), '('+time.Since(t1)+')')
	t2 := time.Now()
	fmt.Println("Part 2", part2(nums), '('+time.Since(t2)+')')
	fmt.Println("Total time: ", time.Since(start))
	fmt.Println()
	tb := time.Now()
	b1, b2 := both(nums)
	fmt.Println("1:", b1, ", 2:", b2, "Took:", time.Since(tb))
}

func part1(nums []int) (res int) {
	res = -1
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		fuel := 0
		for _, n := range nums {
			fuel += diff(i, n)
		}
		if fuel < res || res == -1 {
			res = fuel
		}
	}
	return
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func part2(nums []int) (res int) {
	res = -1
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		fuel := 0
		for _, n := range nums {
			steps := diff(i, n)
			fuel += (steps * (steps + 1)) / 2
		}
		if fuel < res || res == -1 {
			res = fuel
		}
	}
	return
}

func both(nums []int) (r1, r2 int) {
	r1, r2 = -1, -1
	for i := nums[0]; i <= nums[len(nums)-1]; i++ {
		f1, f2 := 0, 0
		for _, n := range nums {
			steps := diff(i, n)
			f1 += steps
			f2 += (steps * (steps + 1)) / 2
		}
		if f1 < r1 || r1 == -1 {
			r1 = f1
		}
		if f2 < r2 || r2 == -1 {
			r2 = f2
		}
	}
	return
}
