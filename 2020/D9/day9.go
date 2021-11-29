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
	nums := utils.ReadIntsByLine(f)
	t1 := time.Now()
	t1Ans := part1(nums)
	fmt.Println("Part 1:", t1Ans, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(nums, t1Ans), "Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
}

func part1(nums []int) (res int) {
	for i := 25; i < len(nums); i++ {
		if !sumRange(nums[i-25:i], nums[i]) {
			return nums[i]
		}
	}
	return
}

func part2(nums []int, inv int) (res int) {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			res = 0
			for _, v := range nums[i : j+1] {
				res += v
			}
			if res == inv {
				sort.Ints(nums[i : j+1])
				return nums[i] + nums[j]
			}
		}
	}

	return
}

func sumRange(nums []int, target int) bool {
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}
