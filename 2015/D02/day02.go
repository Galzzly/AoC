package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		s := strings.Split(strings.TrimSpace(line), "x")
		l := utils.Atoi(s[0])
		w := utils.Atoi(s[1])
		h := utils.Atoi(s[2])
		res += (2*(l*w) + 2*(w*h) + 2*(h*l)) + int(math.Min(math.Min(float64(l*w), float64(w*h)), float64(h*l)))
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		sides := []int{0, 0, 0}
		s := strings.Split(strings.TrimSpace(line), "x")
		sides[0] = utils.Atoi(s[0])
		sides[1] = utils.Atoi(s[1])
		sides[2] = utils.Atoi(s[2])
		sort.Ints(sides)
		// res += (2*(l*w) + 2*(w*h) + 2*(h*l)) + (l * w * h)
		res += ((2 * sides[0]) + (2 * sides[1])) + (sides[0] * sides[1] * sides[2])
	}
	return
}
