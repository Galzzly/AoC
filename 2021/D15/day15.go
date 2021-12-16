package main

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var Delta = []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func main() {
	start := time.Now()

	risk, rect := utils.MakeIntImagePointMap(utils.ReadFileLineByLine("input"))
	fmt.Println("P1:", both(risk, rect.Max.X, rect.Max.Y, rect, false))
	fmt.Println("Part 1:", part1(risk, rect))
	fmt.Println("Part 2:", part2(risk, rect))
	fmt.Printf("Total: %s\n", time.Since(start))
}

func both(input map[image.Point]int, maxX, maxY int, rect image.Rectangle, big bool) (res int) {
	risk := map[image.Point]int{}
	visited := map[image.Point]bool{}
	risk[image.Point{0, 0}] = 0
	newrect := rect
	if big {
		maxX *= 5
		maxY *= 5
		newrect = image.Rect(0, 0, maxX, maxY)
	}
	current := image.Point{0, 0}
	for {
		for _, d := range Delta {
			adjpoint := adj(current, d)
			if visited[adjpoint] || !adjpoint.In(newrect) {
				continue
			}

			y := adjpoint.Y % rect.Max.Y
			x := adjpoint.X % rect.Max.X
			val := input[image.Point{x, y}]
			val += adjpoint.Y/rect.Max.Y + adjpoint.X/rect.Max.X
			if val > 9 {
				val -= 9
			}
			newRisk := risk[current] + val
			if _, ok := risk[adjpoint]; !ok {
				risk[adjpoint] = newRisk
			} else if newRisk < risk[adjpoint] {
				risk[adjpoint] = newRisk
			}
		}
		visited[current] = true
		if visited[image.Point{maxX - 1, maxY - 1}] {
			break
		}
		res = math.MaxInt
		current = image.Point{maxX, maxY}
		for p, r := range risk {
			if !visited[p] && r < res {
				res = r
				current = p
			}
		}
	}
	return
}

func part1(input map[image.Point]int, rect image.Rectangle) (res int) {
	risk := map[image.Point]int{}
	visited := map[image.Point]bool{}
	risk[image.Point{0, 0}] = 0

	current := image.Point{0, 0}

	for {
		for _, d := range Delta {
			adjpoint := adj(current, d)
			if visited[adjpoint] || !adjpoint.In(rect) {
				continue
			}

			// if _, ok := input[adjpoint]; ok {
			newRisk := risk[current] + input[adjpoint]
			if _, ok := risk[adjpoint]; !ok {
				risk[adjpoint] = newRisk
			} else if newRisk < risk[adjpoint] {
				risk[adjpoint] = newRisk
			}
			// }
		}
		visited[current] = true
		if visited[image.Point{rect.Max.X - 1, rect.Max.Y - 1}] {
			break
		}
		res = math.MaxInt
		current = image.Point{rect.Max.X, rect.Max.Y}
		for p, r := range risk {
			if !visited[p] && r < res {
				res = r
				current = p
			}
		}
	}
	return
}

func part2(input map[image.Point]int, rect image.Rectangle) (res int) {
	risk := map[image.Point]int{}
	visited := map[image.Point]bool{}
	maxX, maxY := rect.Max.X, rect.Max.Y
	maxX *= 5
	maxY *= 5
	newrect := image.Rect(0, 0, maxX, maxY)
	risk[image.Point{0, 0}] = 0
	current := image.Point{0, 0}
	for {
		for _, d := range Delta {
			adjpoint := adj(current, d)
			if visited[adjpoint] || !adjpoint.In(newrect) {
				continue
			}

			// if _, ok := risk[adjpoint]; ok {
			y := adjpoint.Y % rect.Max.Y
			x := adjpoint.X % rect.Max.X
			val := input[image.Point{x, y}]
			val += adjpoint.Y/rect.Max.Y + adjpoint.X/rect.Max.X
			if val > 9 {
				val -= 9
			}
			newRisk := risk[current] + val
			if _, ok := risk[adjpoint]; !ok {
				risk[adjpoint] = newRisk
			} else if newRisk < risk[adjpoint] {
				risk[adjpoint] = newRisk
			}
			// }
		}
		visited[current] = true
		if visited[image.Point{maxX - 1, maxY - 1}] {
			break
		}
		res = math.MaxInt
		current = image.Point{maxX, maxY}
		for p, r := range risk {
			if !visited[p] && r < res {
				res = r
				current = p
			}
		}
	}
	return
}

func adj(p, d image.Point) image.Point {
	return p.Add(d)
}
