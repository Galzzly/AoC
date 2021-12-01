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
	p1, p2 := jolts(nums)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total time: ", time.Since(start))
}

func jolts(nums []int) (part1, part2 int) {
	sort.Ints(nums)
	target := nums[len(nums)-1] + 3
	nums = append(nums, target)
	jolt, jolts, waysTo := 0, [4]int{}, map[int]int{0: 1}
	for _, v := range nums {
		waysTo[v] = waysTo[v-1] + waysTo[v-2] + waysTo[v-3]
		jolts[v-jolt]++
		jolt = v
	}
	part1 = jolts[1] * jolts[3]
	part2 = waysTo[target]
	return
}
