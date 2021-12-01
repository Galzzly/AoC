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
	nums := []int{}
	for _, v := range utils.FileLineByComma(f) {
		nums = append(nums, utils.Atoi(v))
	}
	t1 := time.Now()
	fmt.Println("Part 1:", getNum(nums, 2020), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", getNum(nums, 30000000), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func getNum(nums []int, target int) (res int) {
	m := make(map[int]int)
	v := 0
	for i := 0; i < len(nums); i++ {
		v = nums[i]
		m[v] = i
	}
	v = nextNum(m, v, len(nums)-1)
	for i := len(nums); i < target; i++ {
		nums = append(nums, v)
		next := nextNum(m, v, i)
		m[v] = i
		v = next
	}
	res = nums[len(nums)-1]
	return
}

func nextNum(m map[int]int, n, r int) (res int) {
	lastSeen, ok := m[n]
	if ok {
		res = r - lastSeen
	} else {
		res = 0
	}
	return
}
