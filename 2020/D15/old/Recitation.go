package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), ",")
	nums := []int{}
	for _, v := range lines {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", getNum(nums, 2020), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", getNum(nums, 30000000), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func getNum(nums []int, target int) int {
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
	//fmt.Println(m)
	return nums[len(nums)-1]
}
func nextNum(m map[int]int, n int, r int) int {
	lastSeen, ok := m[n]
	var next int
	if ok {
		next = r - lastSeen
	} else {
		next = 0
	}
	return next
}
