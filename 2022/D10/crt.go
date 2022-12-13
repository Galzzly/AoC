package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileLineByLine(f)
	p1, p2 := solve(lines)
	fmt.Println(p1)
	fmt.Println("Part 2:")
	for _, row := range p2 {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println(time.Since(start))
}

func solve(lines []string) (r1 int, r2 [][]string) {
	x := 1
	var cycle, c int
	row := []string{}
	for _, line := range lines {
		var instr string
		var num int
		fmt.Sscanf(line, "%s %d", &instr, &num)
		switch instr {
		case "noop":
			fmt.Println(cycle, x)
			if c >= x-1 && c <= x+1 {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
			cycle++
			c++
			switch cycle {
			case 20, 60, 100, 140, 180, 220:
				r1 += cycle * x
			case 40, 80, 120, 160, 200, 240:

				r2 = append(r2, row)
				row = []string{}
				c = 0
			}
		case "addx":
			for i := 0; i < 2; i++ {
				if c >= x-1 && c <= x+1 {
					row = append(row, "#")
				} else {
					row = append(row, ".")
				}
				fmt.Println(row)
				cycle++
				c++
				switch cycle {
				case 20, 60, 100, 140, 180, 220:
					r1 += cycle * x
				case 40, 80, 120, 160, 200, 240:
					r2 = append(r2, row)
					row = []string{}
					c = 0
				}
			}

			x += num
		}

	}
	return
}
