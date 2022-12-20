package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Boulders map[utils.Cube]bool

type Cube struct {
	size     utils.Cube
	boulders Boulders
}

var Delta = []utils.Cube{
	{-1, 0, 0}, {1, 0, 0},
	{0, -1, 0}, {0, 1, 0},
	{0, 0, -1}, {0, 0, 1},
}

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileLineByLine(f)
	cube := getBoulders(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(cube.boulders), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(cube), "- Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(boulders Boulders) (res int) {
	var count int
	for b := range boulders {
		for _, d := range Delta {
			if _, ok := boulders[utils.Cube{b.X + d.X, b.Y + d.Y, b.Z + d.Z}]; ok {
				count++
			}
		}
	}
	res = (len(boulders) * 6) - count
	return
}

func part2(cube Cube) (res int) {
	exterior := map[utils.Cube]bool{}
	res = markExterior(utils.Cube{0, 0, 0}, cube, exterior)
	return
}

func markExterior(coord utils.Cube, cube Cube, exterior map[utils.Cube]bool) (res int) {
	if exterior[coord] {
		return 0
	}
	if coord.X < -1 || coord.X > cube.size.X+1 ||
		coord.Y < -1 || coord.Y > cube.size.Y+1 ||
		coord.Z < -1 || coord.Z > cube.size.Z+1 {
		return 0
	}
	if cube.boulders[coord] {
		return 1
	}
	exterior[coord] = true
	for _, d := range Delta {
		newcoord := utils.Cube{coord.X + d.X, coord.Y + d.Y, coord.Z + d.Z}
		res += markExterior(newcoord, cube, exterior)
	}
	return
}

func getBoulders(lines []string) Cube {
	size := utils.Cube{0, 0, 0}
	boulders := Boulders{}
	for _, line := range lines {
		var b utils.Cube
		fmt.Sscanf(line, "%d,%d,%d", &b.X, &b.Y, &b.Z)
		boulders[b] = true
		if b.X > size.X {
			size.X = b.X
		}
		if b.Y > size.Y {
			size.Y = b.Y
		}
		if b.Z > size.Z {
			size.Z = b.Z
		}
	}
	return Cube{size, boulders}
}
