package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	lines := utils.ReadFileDoubleLine(os.Args[1])
	paper := makePaper(strings.Split(lines[0], "\n"))
	folds := getFolds(strings.Split(lines[1], "\n"))
	t1 := time.Now()
	fmt.Println("Part 1:", part1(paper, folds), "Took:", time.Since(t1))
	t2 := time.Now()
	p2 := part2(paper, folds)
	fmt.Println("Part 2:")
	for _, line := range p2 {
		for _, v := range line {
			if v == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
	fmt.Println("Took:", time.Since(t2))
	fmt.Println("Total:", time.Since(start))
}

func part1(paper []utils.Point, folds []utils.Point) (res int) {
	var grid = make([][]int, (folds[1].Y*2)+1)
	for i := range grid {
		grid[i] = make([]int, (folds[0].X*2)+1)
	}
	for _, p := range paper {
		var x, y int = p.X, p.Y
		if p.X > folds[0].X {
			x = 2*folds[0].X - p.X
		}
		if grid[y][x] == 0 {
			res++
			grid[y][x] = 1
		}
	}
	return
}

func part2(paper []utils.Point, folds []utils.Point) (res [][]int) {
	res = make([][]int, 6)
	for i := range res {
		res[i] = make([]int, 40)
	}
	for _, p := range paper {
		var x, y int = p.X, p.Y

		for _, f := range folds {
			if f.X > 0 {
				if x > f.X {
					x = 2*f.X - x
				}
			} else {
				if y > f.Y {
					y = 2*f.Y - y
				}
			}
		}
		res[y][x] = 1
	}
	return
}

func makePaper(lines []string) (res []utils.Point) {
	for _, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		res = append(res, utils.Point{X: x, Y: y})
	}
	return
}

func getFolds(lines []string) (res []utils.Point) {
	for _, line := range lines {
		var c rune
		var v int
		fmt.Sscanf(line, "fold along %c=%d", &c, &v)
		if c == 'x' {
			res = append(res, utils.Point{X: v})
		} else {
			res = append(res, utils.Point{Y: v})
		}
	}
	return
}
