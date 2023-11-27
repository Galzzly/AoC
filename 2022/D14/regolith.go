package main

import (
	"fmt"
	"image"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Cave struct {
	cave  map[image.Point]rune
	floor int
}

var Delta = []image.Point{{0, 1}, {-1, 1}, {1, 1}}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	cave := buildCave(lines)
	t := time.Now()
	p1, p2 := solve(cave)
	solTime := time.Since(t)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Solution:", solTime)
	fmt.Println("Total time:", time.Since(start))
}

func solve(c Cave) (r1, r2 int) {
	done := false
	start := image.Point{500, 0}
	var point image.Point
	var rest bool
	p := 1
	sand := 0
	for point != start {
		rest, point = sandfall(c, start, p)
		if !rest {
			sand++
		}
		if rest && !done {
			r1 = sand
			done = !done
			c.floor += 2
			p = 2
		}
	}
	r2 = sand
	return
}

func sandfall(c Cave, s image.Point, p int) (bool, image.Point) {
	pD, pDl, pDr := s.Add(Delta[0]), s.Add(Delta[1]), s.Add(Delta[2])
	bD, bDL, bDR := c.blocking(pD, p), c.blocking(pDl, p), c.blocking(pDr, p)
	if bD && bDL && bDR {
		c.cave[s] = 'O'
		return false, s
	} else {
		var nP image.Point
		switch {
		case !bD:
			nP = pD
		case !bDL:
			nP = pDl
		case !bDR:
			nP = pDr
		}
		if p == 1 && nP.Y >= c.floor {
			return true, nP
		}
		return sandfall(c, nP, p)
	}
}

func (c Cave) blocking(p image.Point, n int) bool {
	if n == 2 && p.Y >= c.floor {
		return true
	}
	if v, ok := c.cave[p]; ok {
		return v == 'R' || v == 'O'
	}
	return false
}

func buildCave(lines []string) Cave {
	rocks := make(map[image.Point]rune)
	var maxY int

	for _, line := range lines {
		points := strings.Split(line, " -> ")
		for i, point := range points {
			var cP image.Point
			fmt.Sscanf(point, "%d,%d", &cP.X, &cP.Y)
			rocks[cP] = 'R'
			if cP.Y > maxY {
				maxY = cP.Y
			}
			if i < len(points)-1 {
				var nX, nY int
				fmt.Sscanf(points[i+1], "%d,%d", &nX, &nY)
				if nX != cP.X {
					var d int
					if nX > cP.X {
						d = 1
					} else {
						d = -1
					}
					for x := cP.X; x != nX; x += d {
						rocks[image.Point{x, nY}] = 'R'
					}
				} else if nY != cP.Y {
					var d int
					if nY > cP.Y {
						d = 1
					} else {
						d = -1
					}
					for y := cP.Y; y != nY; y += 1 * d {
						rocks[image.Point{nX, y}] = 'R'
					}
				}
			}
		}
	}
	return Cave{rocks, maxY}
}
