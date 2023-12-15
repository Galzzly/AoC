package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Grid [][]string

var target = 1000000000

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	grid := makeGrid(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(grid, false), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(grid, true), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func solve(grid Grid, part2 bool) (res int) {
	rockCycle := map[string]int{}
	for t := 1; t <= target; t++ {
		for d := 0; d < 4; d++ {
			for Y := 1; Y < len(grid); Y++ {
				for X := 0; X < len(grid[0]); X++ {
					if grid[Y][X] != "O" {
						continue
					}
					for YY := Y - 1; YY >= 0; YY-- {
						if grid[YY][X] != "." {
							break
						}
						grid[YY][X] = "O"
						grid[YY+1][X] = "."
					}
				}
			}
			if t == 1 && d == 0 && !part2 {
				return grid.Calculate()
			}
			grid = grid.Rotate()
		}
		k := ""
		for _, R := range grid {
			k += strings.Join(R, "")
		}
		if val, ok := rockCycle[k]; ok {
			cycLen := t - val
			remaining := target - t
			a := remaining / cycLen
			t += a * cycLen
		}
		rockCycle[k] = t
	}
	return grid.Calculate()
}

func (g Grid) Print() {
	for _, r := range g {
		fmt.Println(strings.Join(r, " "))
	}
	fmt.Println()
}

func (g Grid) Rotate() (ng Grid) {
	ng = make(Grid, len(g[0]))
	for i := range ng {
		ng[i] = make([]string, len(g))
	}
	maxY := len(g)
	for X := 0; X < len(g[0]); X++ {
		for Y := 0; Y < len(g); Y++ {
			ng[X][maxY-1-Y] = g[Y][X]
		}
	}
	return
}

func (g Grid) Calculate() (res int) {
	for Y := 0; Y < len(g); Y++ {
		var C int
		for X := 0; X < len(g[0]); X++ {
			C += utils.Ter(g[Y][X] == "O", 1, 0)
		}
		res += C * (len(g) - Y)
	}
	return
}

func makeGrid(line []string) (grid Grid) {
	grid = make(Grid, len(line))
	for i, line := range line {
		grid[i] = strings.Split(line, "")
	}
	return
}
