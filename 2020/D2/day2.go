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
	var count int
	for i := 0; i < len(lines); i++ {
		vals := strings.Fields(lines[i])
		n := strings.Split(vals[0], "-")
		mm := make([]int, 0, len(n))
		for _, a := range n {
			m, err := strconv.Atoi(a)
			utils.Check(err)
			mm = append(mm, m)
		}
		char := vals[1][:len(vals[1])-1]
		res := strings.Count(vals[2], char)
		if (res >= mm[0]) && (res <= mm[1]) {
			count++
		}
	}
	return count
}

func part2(lines []string) int {
	var count int
	for i := 0; i < len(lines); i++ {
		vals := strings.Fields(lines[i])
		n := strings.Split(vals[0], "-")
		p := make([]int, 0, len(n))
		for _, a := range n {
			m, err := strconv.Atoi(a)
			utils.Check(err)
			p = append(p, m)
		}
		char := vals[1][:len(vals[1])-1]
		pass := string(vals[2])
		r := make([]string, 0, len(p))
		for _, a := range p {
			c := string(pass[a-1])
			r = append(r, c)
		}
		if (r[0] == char) && (r[1] == char) {
			continue
		} else if (r[0] == char) || (r[1] == char) {
			count++
		}
	}
	return count
}
