package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func part1(instr []string) (res int) {
	p := utils.Point{0, 0}
	d := utils.Point{1, 0}
	for _, i := range instr {
		a := i[:1]
		b := utils.Atoi(i[1:])
		switch a {
		case "F":
			p.X += d.X * b
			p.Y += d.Y * b
		case "R":
			for i := 0; i < (b / 90); i++ {
				d.X, d.Y = d.Y, -d.X
			}
		case "L":
			for i := 0; i < (b / 90); i++ {
				d.X, d.Y = -d.Y, d.X
			}
		case "W":
			p.X -= b
		case "E":
			p.X += b
		case "S":
			p.Y -= b
		case "N":
			p.Y += b
		}
	}
	res = utils.Abs(p.X) + utils.Abs(p.Y)
	return
}

func part2(instr []string) (res int) {
	p := utils.Point{0, 0}
	d := utils.Point{10, 1}
	for _, i := range instr {
		a := i[:1]
		b := utils.Atoi(i[1:])
		switch a {
		case "F":
			p.X += d.X * b
			p.Y += d.Y * b
		case "R":
			for i := 0; i < (b / 90); i++ {
				d.X, d.Y = d.Y, -d.X
			}
		case "L":
			for i := 0; i < (b / 90); i++ {
				d.X, d.Y = -d.Y, d.X
			}
		case "W":
			d.X -= b
		case "E":
			d.X += b
		case "S":
			d.Y -= b
		case "N":
			d.Y += b
		}
	}
	res = utils.Abs(p.X) + utils.Abs(p.Y)
	return
}
