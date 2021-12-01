package main

import (
	"fmt"
	"image"
	"os"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	seats := utils.MakeImagePointMap(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", getSeats(seats, 4, func(p, d image.Point) image.Point { return p.Add(d) }), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", getSeats(seats, 5, func(p, d image.Point) image.Point {
		for seats[p.Add(d)] == '.' {
			p = p.Add(d)
		}
		return p.Add(d)
	}), "Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
}

func getSeats(seats map[image.Point]rune, maxAdj int, adj func(p, d image.Point) image.Point) (occ int) {
	delta := []image.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for diff := true; diff; {
		occ = 0
		diff = false
		next := map[image.Point]rune{}
		for p, r := range seats {
			sum := 0
			for _, d := range delta {
				if seats[adj(p, d)] == '#' {
					sum++
				}
			}

			if r == '#' && sum >= maxAdj {
				r = 'L'
			} else if r == 'L' && sum == 0 || r == '#' {
				r = '#'
				occ++
			}
			next[p] = r
			diff = diff || r != seats[p]
		}
		seats = next
	}
	return
}
