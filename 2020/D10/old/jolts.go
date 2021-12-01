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

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s: ", elapsed)
}

func readInts(lines []string) (nums []int, err error) {
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

func p1(num []int) (int, int) {
	defer timeTrack(time.Now())
	sort.Ints(num)
	target := num[len(num)-1] + 3
	num = append(num, target)
	jolt, jolts, waysTo := 0, [4]int{}, map[int]int{0: 1}
	for _, v := range num {
		waysTo[v] = waysTo[v-1] + waysTo[v-2] + waysTo[v-3]
		jolts[v-jolt]++
		jolt = v
	}
	res := jolts[1] * jolts[3]
	return res, waysTo[target]
}

func main() {
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	num, err := readInts(lines)
	check(err)
	part1, part2 := p1(num)
	fmt.Printf("Part 1: %d, Part2: %d\n", part1, part2)
}
