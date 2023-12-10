package main

import (
	"fmt"
	"image"
	"slices"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type PipeMap map[image.Point]rune

func main() {
	start := time.Now()
	f := "input.txt"
	pipemap := utils.MakeImagePointMap(utils.ReadFileLineByLine(f))
	P, _ := utils.MapKey(pipemap, 'S')
	t1 := time.Now()
	p1, p2map := part1(pipemap, P)
	fmt.Println("Part 1:", p1, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(p2map, pipemap), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(mapper PipeMap, startpoint image.Point) (int, []image.Point) {
	PointsInLoop := []image.Point{startpoint}
	P := startpoint
	Dir := '>'
	for {
		var NP image.Point
		switch mapper[P] {
		case 'S': // this is the start point.
			NP = P.Add(image.Point{1, 0})
		case '|': // Vertical Pipe
			if Dir == 'U' {
				NP = P.Add(image.Point{0, -1})
			} else {
				NP = P.Add(image.Point{0, 1})
			}
		case '-': // Horizontal Pipe
			if Dir == '>' {
				NP = P.Add(image.Point{1, 0})
			} else {
				NP = P.Add(image.Point{-1, 0})
			}
		case 'L': // North to East
			if Dir == 'D' {
				NP = P.Add(image.Point{1, 0})
				Dir = '>'
			} else {
				NP = P.Add(image.Point{0, -1})
				Dir = 'U'
			}
		case 'J': // North to West
			if Dir == 'D' {
				NP = P.Add(image.Point{-1, 0})
				Dir = '<'
			} else {
				NP = P.Add(image.Point{0, -1})
				Dir = 'U'
			}
		case '7': // South to West
			if Dir == 'U' {
				NP = P.Add(image.Point{-1, 0})
				Dir = '<'
			} else {
				NP = P.Add(image.Point{0, 1})
				Dir = 'D'
			}
		case 'F': // South to East
			if Dir == '<' {
				NP = P.Add(image.Point{0, 1})
				Dir = 'D'
			} else {
				NP = P.Add(image.Point{1, 0})
				Dir = '>'
			}
		case '.': // Ground, no pipe
		}
		if NP == startpoint {
			break
		}
		P = NP
		PointsInLoop = append(PointsInLoop, P)
	}
	return len(PointsInLoop) / 2, PointsInLoop
}

func part2(loop []image.Point, pipemap PipeMap) (res int) {
	pf := utils.NewPolyfence(loop)
	for P := range pipemap {
		if !slices.Contains[[]image.Point](loop, P) && pf.Inside(P) {
			res++
		}
	}
	return
}
