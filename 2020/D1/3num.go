package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInts(f string) (nums []int, err error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}
		n, err := strconv.Atoi(lines[l])
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func main() {
	start := time.Now()
	f := os.Args[1]
	nums, err := ReadInts(f)
	if err != nil {
		panic(err)
	}

	sort.Ints(nums)
	for _, num1 := range nums {
		for _, num2 := range nums {
			target := 2020 - num1 - num2
			i := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
			if i < len(nums) && nums[i] == target {
				fmt.Printf("Found %d, %d an %d, answer is %d\n", num1, num2, target, num1*num2*target)
			}
		}
	}
	fmt.Println(time.Since(start))
}
