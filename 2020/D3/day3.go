package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(lines []string) int {
	return slopes(lines, 3, 1)
}

func part2(lines []string) int {
	var total int
	trav := []string{"1 1", "3 1", "5 1", "7 1", "1 2"}
	res := make([]int, 0)
	for _, a := range trav {
		s := strings.Split(a, " ")
		h, _ := strconv.Atoi(s[0])
		v, _ := strconv.Atoi(s[1])
		res = append(res, slopes(lines, h, v))
	}
	total = res[0]
	for i := 1; i < len(res); i++ {
		total = total * res[i]
	}
	return total
}

func slopes(lines []string, h, v int) int {
	x, y := 0, 0
	count := 0

	for y < len(lines)-v {
		y += v
		x += h
		if x >= len(lines[y]) {
			x = x - len(lines[y])
		}
		c := string(lines[y])
		r := string(c[x])
		if r == "#" {
			count++
		}
	}
	return count
}
