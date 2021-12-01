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
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	t, _ := strconv.Atoi(lines[0])
	b := make(map[int]int, 0)
	//	idx := make([]int, 0)
	for i, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		n, _ := strconv.Atoi(v)
		b[i] = n
		//		idx = append(idx, i)
	}
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", part1(t, b), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(t, b), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func part1(t int, b map[int]int) int {
	minT, minB := 10000, -1
	for _, bus := range b {
		next := bus - (t % bus)
		if next < minT {
			minT = next
			minB = bus
		}
	}
	return minT * minB
}

func part2(t int, b map[int]int) int {
	res, step := 0, 1
	for i, bus := range b {
		for (res+i)%bus != 0 {
			res += step
		}
		step *= bus
	}
	return res
}
