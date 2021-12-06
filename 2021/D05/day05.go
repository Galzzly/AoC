package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	input := getInput(lines)
	fmt.Println("Part 1:", solver(input, false))
	fmt.Println("Part 2:", solver(input, true))
}

func solver(input [][]utils.Point, diag bool) (res int) {
	grid := make(map[int]map[int]int)
	for _, p := range input {
		if !diag && (p[0].X != p[1].X && p[0].Y != p[1].Y) {
			continue
		}
		start := p[0]
		modX, modY := 0, 0
		if p[0].X > p[1].X {
			modX = -1
		}
		if p[0].X < p[1].X {
			modX = 1
		}
		if p[0].Y > p[1].Y {
			modY = -1
		}
		if p[0].Y < p[1].Y {
			modY = 1
		}
		for {
			if _, ok := grid[start.X]; !ok {
				grid[start.X] = make(map[int]int)
			}
			grid[start.X][start.Y]++
			if grid[start.X][start.Y] == 2 {
				res++
			}
			if start.X == p[1].X && start.Y == p[1].Y {
				break
			}
			start.X += modX
			start.Y += modY
		}
	}
	return
}

func getInput(lines []string) (input [][]utils.Point) {
	input = make([][]utils.Point, len(lines))
	for i, line := range lines {
		s := strings.Split(strings.TrimSpace(line), " -> ")
		for _, v := range s {
			p := strings.Split(v, ",")
			x, y := utils.Atoi(p[0]), utils.Atoi(p[1])
			input[i] = append(input[i], utils.Point{X: x, Y: y})
		}
	}
	return
}
