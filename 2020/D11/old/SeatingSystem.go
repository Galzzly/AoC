package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	seats := map[image.Point]rune{}
	for y, s := range lines {
		for x, r := range s {
			seats[image.Point{x, y}] = r
		}
	}
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", getSeats(seats, 4, func(p, d image.Point) image.Point { return p.Add(d) }), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", getSeats(seats, 5, func(p, d image.Point) image.Point {
		for seats[p.Add(d)] == '.' {
			p = p.Add(d)
		}
		return p.Add(d)
	}), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
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
			diff = diff || next[p] != seats[p]
		}
		seats = next
	}
	return
}
