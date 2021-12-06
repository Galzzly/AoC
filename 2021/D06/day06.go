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
	nums := utils.FileIntsLineByComma(f)
	t1 := time.Now()
	fmt.Println("Part 1:", fish(nums, 80), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", fish(nums, 256), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func fish(nums []int, days int) (res int) {
	currentstate := getCounts(nums)
	for i := 0; i < days; i++ {
		nextstate := make(map[int]int)
		for j := len(currentstate); j > 0; j-- {
			nextstate[j-1] = currentstate[j]
		}
		nextstate[6] += currentstate[0]
		nextstate[8] += currentstate[0]
		currentstate = nextstate
	}
	for _, v := range currentstate {
		res += v
	}
	return
}

func getCounts(nums []int) (res map[int]int) {
	res = make(map[int]int)
	for _, n := range nums {
		res[n]++
	}
	return
}
