package main

import (
	"fmt"
	"image"
	"math"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Grid struct {
	*utils.Grid[rune]
}

func (g Grid) Neighbours(p image.Point) (res []image.Point) {
	val := g.GetState(p)
	res = utils.Select(g.Grid.Neighbours(p), func(x image.Point) bool {
		return g.GetState(x) <= val+1
	})
	return
}

var Delta = []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	grid, s, e, a := buildGrid(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(grid, s, e), "- took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(grid, a, e), "- took:", time.Since(t2))
	fmt.Println("Took:", time.Since(start))
}

func part1(g Grid, s, e image.Point) (res int) {
	res = len(utils.Search(g, s, e)) - 1
	return
}

func part2(g Grid, a []image.Point, e image.Point) (res int) {
	res = math.MaxInt
	distance := make(chan int, len(a))
	var wg sync.WaitGroup
	for _, p := range a {
		wg.Add(1)
		go func(p image.Point) {
			defer wg.Done()
			d := len(utils.Search(g, p, e)) - 1
			distance <- d
		}(p)
	}
	wg.Wait()
	close(distance)
	for d := range distance {
		if d < res && d > 1 {
			res = d
		}
	}
	return
}

func buildGrid(lines []string) (grid Grid, s, e image.Point, a []image.Point) {
	grid.Grid = utils.NewGrid[rune](len(lines[0]), len(lines), Delta)
	for y, line := range lines {
		for x, c := range line {
			switch c {
			case 'S':
				s = image.Point{x, y}
				grid.SetState(x, y, 'a')
			case 'E':
				e = image.Point{x, y}
				grid.SetState(x, y, 'z')
			case 'a':
				a = append(a, image.Point{x, y})
				grid.SetState(x, y, 'a')
			default:
				grid.SetState(x, y, c)
			}

		}
	}
	return
}
