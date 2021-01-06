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
	b := make([]int, 0)
	for _, v := range strings.Split(lines[1], ",") {
		if v == "x" {
			continue
		}
		n, _ := strconv.Atoi(v)
		b = append(b, n)
	}
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", part1(t, b), time.Since(t1))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func part1(t int, b []int) int {
	minT, minB := 0, -1
	for i, bus := range b {
		next := bus - (t % bus)
		if i == 0 || next < minT {
			minT = next
			minB = bus
		}
	}
	fmt.Println(minT, minB)
	return minT * minB
}

func part2(t int, b []int) int {

}
