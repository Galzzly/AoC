package main

import (
	"fmt"
	"os"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	line := utils.ReadFileLineByLine(f)[0]
	fmt.Println("Part 1:", part1(line))
	fmt.Println("Part 2:", part2(line))
}

func part1(s string) (res int) {
	var homes = make(map[utils.Point]bool)
	coords := utils.Point{0, 0}
	homes[coords] = true
	for _, c := range s {
		switch c {
		case '^':
			coords = utils.Point{coords.X, coords.Y + 1}
		case '>':
			coords = utils.Point{coords.X + 1, coords.Y}
		case 'v':
			coords = utils.Point{coords.X, coords.Y - 1}
		case '<':
			coords = utils.Point{coords.X - 1, coords.Y}
		}
		homes[coords] = true
	}
	res = len(homes)

	return
}

func part2(s string) (res int) {
	var homes = make(map[utils.Point]bool)
	//var rhomes = make(map[utils.Point]bool)
	scoord := utils.Point{0, 0}
	rcoord := utils.Point{0, 0}
	homes[scoord] = true
	//rhomes[rcoord] = true
	for i, c := range s {
		if i%2 == 0 {
			switch c {
			case '^':
				scoord = utils.Point{scoord.X, scoord.Y + 1}
			case '>':
				scoord = utils.Point{scoord.X + 1, scoord.Y}
			case 'v':
				scoord = utils.Point{scoord.X, scoord.Y - 1}
			case '<':
				scoord = utils.Point{scoord.X - 1, scoord.Y}
			}
			homes[scoord] = true
		} else {
			switch c {
			case '^':
				rcoord = utils.Point{rcoord.X, rcoord.Y + 1}
			case '>':
				rcoord = utils.Point{rcoord.X + 1, rcoord.Y}
			case 'v':
				rcoord = utils.Point{rcoord.X, rcoord.Y - 1}
			case '<':
				rcoord = utils.Point{rcoord.X - 1, rcoord.Y}
			}
			homes[rcoord] = true
		}
	}
	res = len(homes)
	return
}
