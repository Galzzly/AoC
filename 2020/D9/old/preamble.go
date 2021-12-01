//package day9

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := ioutil.ReadFile("input")
	check(err)
	split := strings.Split(strings.TrimSpace(string(f)), "\n")

	nums := make([]int, len(split))
	for i, s := range split {
		nums[i], _ = strconv.Atoi(s)
	}

	invalid := 0

findNum:
	for i := 25; i < len(nums); i++ {
		for j := i - 25; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if nums[j]+nums[k] == nums[i] {
					continue findNum
				}
			}
		}
		invalid = nums[i]
		break
	}
	fmt.Println(invalid)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := 0
			for _, v := range nums[i : j+1] {
				sum += v
			}
			if sum == invalid {
				sort.Ints(nums[i : j+1])
				fmt.Println(nums[i] + nums[j])
				return
			}
		}
	}
}
