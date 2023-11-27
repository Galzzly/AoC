package main

import (
	"fmt"
	"image"
	"reflect"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var D = [][]image.Point{
	{{0, -1}, {1, -1}, {-1, -1}}, // N, NE, NW
	{{0, 1}, {1, 1}, {-1, 1}},    // S, SE, SW
	{{-1, 0}, {-1, -1}, {-1, 1}}, //W, NW, SW
	{{1, 0}, {1, -1}, {1, 1}},    //E, NE, SE
}

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileLineByLine(f)
	elves := mapElves(lines)
	r1, r2 := solve(elves)
	fmt.Println("Part 1:", r1)
	fmt.Println("Part 2:", r2)
	fmt.Println("Total Time:", time.Since(start))
}

func solve(elves map[image.Point]struct{}) (r1, r2 int) {
	currentelves := utils.CopyMap(elves)
	dlen := len(D)
	for i := 0; ; i++ {
		prop := map[image.Point]image.Point{}
		count := map[image.Point]int{}
		for e := range currentelves {
			nbr := map[int]int{}
			for d := range D {
				for _, p := range D[d] {
					if _, ok := currentelves[e.Add(p)]; ok {
						nbr[d]++
					}
				}
			}
			if len(nbr) == 0 {
				continue
			}

			for d := 0; d < dlen; d++ {
				if dir := (i + d) % len(D); nbr[dir] == 0 {
					prop[e] = e.Add(D[dir][0])
					count[prop[e]]++
					break
				}
			}
		}

		nextelves := map[image.Point]struct{}{}
		for e := range currentelves {
			if _, ok := prop[e]; ok && count[prop[e]] == 1 {
				e = prop[e]
			}
			nextelves[e] = struct{}{}
		}

		if i == 9 {
			var r image.Rectangle
			for e := range currentelves {
				r = r.Union(image.Rectangle{e, e.Add(image.Point{1, 1})})
			}
			r1 = r.Dx()*r.Dy() - len(currentelves)
		}

		if reflect.DeepEqual(currentelves, nextelves) {
			i++
			r2 = i
			break
		}

		currentelves = nextelves
	}
	return
}

func mapElves(lines []string) (elves map[image.Point]struct{}) {
	elves = map[image.Point]struct{}{}
	for y, s := range lines {
		for x, r := range s {
			if r == '#' {
				elves[image.Point{x + 10, y + 10}] = struct{}{}
			}
		}
	}
	return
}
