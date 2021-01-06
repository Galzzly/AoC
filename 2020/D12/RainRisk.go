package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

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
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", part1(lines), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(lines), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func part1(instr []string) int {
	pos := point{0, 0}
	dir := point{1, 0}

	for _, i := range instr {
		a := i[0:1]
		b, err := strconv.Atoi(i[1:])
		check(err)

		switch a {
		case "F":
			// Move forward by b in the direction facing
			pos.x += dir.x * b
			pos.y += dir.y * b
		case "R":
			// Turn right by b number of degrees
			for i := 0; i < (b / 90); i++ {
				dir.x, dir.y = dir.y, -dir.x
			}
		case "L":
			// Turn left by b number of degrees
			for i := 0; i < (b / 90); i++ {
				dir.x, dir.y = -dir.y, dir.x
			}
		case "W":
			// Move West by b
			pos.x -= b
		case "E":
			// Move East by b
			pos.x += b
		case "S":
			// Move South by b
			pos.y -= b
		case "N":
			// Move North by b
			pos.y += b
		}
	}

	return abs(pos.x) + abs(pos.y)
}

func part2(instr []string) int {
	pos := point{0, 0}
	dir := point{10, 1}

	for _, i := range instr {
		a := i[0:1]
		b, err := strconv.Atoi(i[1:])
		check(err)

		switch a {
		case "F":
			// Move forward by b in the direction facing
			pos.x += dir.x * b
			pos.y += dir.y * b
		case "R":
			// Turn right by b number of degrees
			for i := 0; i < (b / 90); i++ {
				dir.x, dir.y = dir.y, -dir.x
			}
		case "L":
			// Turn left by b number of degrees
			for i := 0; i < (b / 90); i++ {
				dir.x, dir.y = -dir.y, dir.x
			}
		case "W":
			// Move West by b
			dir.x -= b
		case "E":
			// Move East by b
			dir.x += b
		case "S":
			// Move South by b
			dir.y -= b
		case "N":
			// Move North by b
			dir.y += b
		}
	}
	return abs(pos.x) + abs(pos.y)
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
