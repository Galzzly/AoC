package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Galzzly/AoC/2020/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	nums := utils.ReadIntsByLine(f)
	sort.Ints(nums)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(nums), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(nums), "Took:", time.Since(t2))
	fmt.Println("Total Time: ", time.Since(start))
}

func part1(nums []int) int {
	var res int
	for _, num := range nums {
		target := 2020 - num
		i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
		if i < len(nums) && nums[i] == target {
			res = num * target
			break
		}
	}
	return res
}

func part2(nums []int) int {
	var res int
	for _, num1 := range nums {
		for _, num2 := range nums {
			target := 2020 - num1 - num2
			i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
			if i < len(nums) && nums[i] == target {
				res = num1 * num2 * target
				break
			}
		}
	}
	return res
}
