package main

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var (
	delta = []image.Point{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
)

func main() {
	start := time.Now()
	lines := utils.ReadFileLineByLine(os.Args[1])
	octopuses, rect := utils.MakeIntImagePointMap(lines)
	tb := time.Now()
	r1, r2 := both(octopuses, rect)
	fmt.Println("Part 1:", r1, "\nPart 2:", r2, "\nTook:", time.Since(tb))
	fmt.Println("Total Time:", time.Since(start))
}

func both(octopuses map[image.Point]int, rect image.Rectangle) (r1, r2 int) {
	current := octopuses
	var count, r1step int
	var r1b, r2b = false, false
	for !(r2b && r1b) {
		r1step++
		if !r2b {
			r2++
		}
		current, count = step(current, rect)
		if count == 100 {
			r2b = true
		}
		r1 += count
		if r1step == 100 {
			r1b = true
		}
	}
	return
}

func step(current map[image.Point]int, rect image.Rectangle) (next map[image.Point]int, count int) {
	next = make(map[image.Point]int)
	flashed := make(map[image.Point]bool)

	for p := range current {
		next[p] = current[p] + 1
	}
	for p := range next {
		if flashed[p] {
			continue
		}
		if next[p] > 9 {
			next[p] = 0
			flashed[p] = true
			count++
			toCheck := []image.Point{}
			for _, d := range delta {
				adjpoint := adj(p, d)
				toCheck = append(toCheck, adjpoint)
			}
			var point image.Point
			for len(toCheck) > 0 {
				point, toCheck = toCheck[0], toCheck[1:]
				if flashed[point] {
					continue
				}
				if !point.In(rect) {
					continue
				}
				next[point]++
				if next[point] > 9 {
					next[point] = 0
					count++
					flashed[point] = true
					for _, d := range delta {
						adjpoint := adj(point, d)
						toCheck = append(toCheck, adjpoint)
					}
				}
			}
		}
	}
	return
}

func adj(p, d image.Point) image.Point {
	return p.Add(d)
}
