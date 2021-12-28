package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Grid [][]string

func (g Grid) String() string {
	var sb strings.Builder
	for _, row := range g {
		for _, c := range row {
			if c == "" {
				sb.WriteByte('.')
			} else {
				sb.WriteString(c)
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func main() {
	start := time.Now()
	lines := utils.ReadFileLineByLine("input")
	grid := buildGrid(lines)
	t1 := time.Now()
	fmt.Printf("Part 1: %d in %s\n", part1(grid), time.Since(t1))
	fmt.Printf("Total Time: %s", time.Since(start))
}

func part1(grid Grid) (res int) {
	var count int
	// current := grid
	for ok := true; ok; grid, ok = step(grid) {
		count++
		fmt.Println(count)
		fmt.Println(grid)
	}

	res = count
	return
}

func step(g Grid) (next Grid, movement bool) {
	next = make([][]string, len(g))
	for y, row := range g {
		next[y] = make([]string, len(row))
	}

	for y, row := range g {
		for x, c := range row {
			switch c {
			case "":
				continue
			case "v":
				next[y][x] = "v"
			case ">":
				idx := (x + 1) % len(row)
				if g[y][idx] == "" {
					next[y][idx] = ">"
					movement = true
				} else {
					next[y][x] = ">"
				}
			}
		}
	}

	g = next
	next = make([][]string, len(g))
	for y, row := range g {
		next[y] = make([]string, len(row))
	}

	for y, row := range g {
		for x, c := range row {
			switch c {
			case "":
				continue
			case "v":
				idy := (y + 1) % len(g)
				if g[idy][x] == "" {
					next[idy][x] = "v"
					movement = true
				} else {
					next[y][x] = "v"
				}
			case ">":
				next[y][x] = ">"
			}
		}
	}

	return
}

func buildGrid(lines []string) Grid {
	grid := make([][]string, len(lines))
	for y, line := range lines {
		grid[y] = make([]string, len(line))
		for x, c := range line {
			if c != '.' {
				grid[y][x] = string(c)
			}
		}
	}
	return grid
}
