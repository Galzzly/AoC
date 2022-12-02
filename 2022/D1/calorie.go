package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]

	lines := utils.ReadFileDoubleLine(f)
	elves := getElves(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", elves[len(elves)-1], "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", utils.SumArray(elves[len(elves)-3:]), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func getElves(lines []string) []int {
	var res []int
	for _, l := range lines {
		var cal int
		for _, n := range strings.Split(strings.TrimSpace(l), "\n") {
			cal += utils.Atoi(n)
		}
		res = append(res, cal)
	}
	sort.Ints(res)
	return res
}
