package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type instr struct {
	dir string
	val int
}

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	var instrs = make([]instr, len(lines))
	for _, line := range lines {
		s := strings.Split(strings.TrimSpace(line), " ")
		instrs = append(instrs, instr{s[0], utils.Atoi(s[1])})
	}
	t1 := time.Now()
	fmt.Println("Part 1:", part1(instrs), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(instrs), "Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
	tb := time.Now()
	p1, p2 := both(instrs)
	fmt.Println("Part1:", p1, "Part2:", p2, "Took:", time.Since(tb))
}

func part1(instrs []instr) (res int) {
	pos := utils.Point{0, 0}

	for _, instr := range instrs {
		switch instr.dir {
		case "forward":
			pos.X += instr.val
		case "down":
			pos.Y -= instr.val
		case "up":
			pos.Y += instr.val
		}
	}
	res = utils.Abs(pos.Y) * utils.Abs(pos.X)
	return
}

func part2(instrs []instr) (res int) {
	pos := utils.Cube{0, 0, 0}
	for _, instr := range instrs {
		switch instr.dir {
		case "forward":
			pos.X += instr.val
			pos.Y += instr.val * pos.Z
		case "down":
			pos.Z += instr.val
		case "up":
			pos.Z -= instr.val
		}
	}
	res = utils.Abs(pos.Y) * utils.Abs(pos.X)
	return
}

// Added this in for a speed test only.
func both(instrs []instr) (p1, p2 int) {
	pos1 := utils.Point{0, 0}
	pos2 := utils.Cube{0, 0, 0}
	for _, instr := range instrs {
		switch instr.dir {
		case "forward":
			pos1.X += instr.val
			pos2.X += instr.val
			pos2.Y += instr.val * pos2.Z
		case "down":
			pos1.Y -= instr.val
			pos2.Z += instr.val
		case "up":
			pos1.Y += instr.val
			pos2.Z -= instr.val
		}
	}
	p1 = utils.Abs(pos1.Y) * utils.Abs(pos1.X)
	p2 = utils.Abs(pos2.Y) * utils.Abs(pos2.X)
	return
}
