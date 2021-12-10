package main

import (
	"fmt"
	"image"
	"os"
	"sort"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	points, rect := utils.MakeIntImagePointMap(lines)
	t1 := time.Now()
	r1, lowPoints := part1(points, rect, func(p, d image.Point) image.Point { return p.Add(d) })
	fmt.Println("Part 1:", r1, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lowPoints, points, rect, func(p, d image.Point) image.Point { return p.Add(d) }), "Took:", time.Since(t2))
	fmt.Println("Total:", time.Since(start))
}

func part1(points map[image.Point]int, rect image.Rectangle, adj func(p, d image.Point) image.Point) (res int, lowPoints []image.Point) {
	delta := []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for p, r := range points {
		adjacents := []int{}
		for _, d := range delta {
			adjpoint := adj(p, d)
			if adjpoint.In(rect) {
				adjacents = append(adjacents, points[adj(p, d)])
			}

		}
		if pointlower(r, adjacents) {
			lowPoints = append(lowPoints, p)
			res += r
		}
	}
	res += len(lowPoints)
	return
}

func part2(lowPoints []image.Point, points map[image.Point]int, rect image.Rectangle, adj func(p, d image.Point) image.Point) (res int) {
	counts := []int{}
	for _, p := range lowPoints {
		count := getBasin(p, points, rect, adj)
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	res = 1
	for i := 0; i < 3; i++ {
		res *= counts[i]
	}
	return
}

func getBasin(p image.Point, points map[image.Point]int, rect image.Rectangle, adj func(p, d image.Point) image.Point) (res int) {
	res = 1
	seen := map[image.Point]struct{}{}
	delta := []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	toCheck := []image.Point{}
	for _, d := range delta {
		adjpoint := adj(p, d)
		toCheck = append(toCheck, adjpoint)
	}
	seen[p] = struct{}{}
	var point image.Point
	for len(toCheck) > 0 {
		point, toCheck = toCheck[0], toCheck[1:]
		if _, ok := seen[point]; ok {
			continue
		}
		if !point.In(rect) {
			continue
		}
		if points[point] == 9 {
			continue
		}
		seen[point] = struct{}{}
		res++
		for _, d := range delta {
			adjpoint := adj(point, d)
			toCheck = append(toCheck, adjpoint)
		}
	}
	return
}

func pointlower(p int, adj []int) bool {
	sum := 0
	for _, a := range adj {
		if a > p {
			sum++
		}
	}
	return sum == len(adj)
}
