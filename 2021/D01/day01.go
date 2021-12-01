package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	nums := utils.ReadIntsByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(nums), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(nums), "Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
}

func part1(nums []int) (res int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			res++
		}
	}
	return
}

func part2(nums []int) (res int) {
	for i := 1; i < len(nums)-2; i++ {
		if (nums[i-1] + nums[i] + nums[i+1]) < (nums[i] + nums[i+1] + nums[i+2]) {
			res++
		}
	}
	return
}
