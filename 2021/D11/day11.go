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
	t1 := time.Now()
	fmt.Println("Part 1:", part1(octopuses, rect), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(octopuses, rect), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(octopusus map[image.Point]int, rect image.Rectangle) (res int) {
	current := octopusus
	var count int
	for i := 0; i < 100; i++ {
		current, count = step(current, rect)
		res += count
	}
	return
}

func part2(octopuses map[image.Point]int, rect image.Rectangle) (res int) {
	current := octopuses
	var count int
	for {
		res++
		current, count = step(current, rect)
		if count == 100 {
			break
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
