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

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s: ", elapsed)
}

func slopes(l []string, h int, v int) int {
	x := 0
	y := 0
	count := 0

	for y < len(l)-v {
		y += v
		x += h
		if x >= len(l[y]) {
			x = x - len(l[y])
		}
		c := string(l[y])
		r := string(c[x])
		if r == "#" {
			count++
		}
	}
	return count
}
func P1(l []string) int {
	defer timeTrack(time.Now())
	c := slopes(l, 3, 1)
	return c
}
func P2(l []string) int {
	defer timeTrack(time.Now())
	trav := []string{"1 1", "3 1", "5 1", "7 1", "1 2"}
	res := make([]int, 0)
	for _, a := range trav {
		s := strings.Split(a, " ")
		h, err := strconv.Atoi(s[0])
		check(err)
		v, err := strconv.Atoi(s[1])
		check(err)
		res = append(res, slopes(l, h, v))
	}
	total := res[0]
	for i := 1; i < len(res); i++ {
		total = total * res[i]
	}
	return total
}
func main() {
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	fmt.Printf("Part 1: %d\n", P1(lines))
	fmt.Printf("Part 2: %d\n", P2(lines))
}
