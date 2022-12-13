package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type instr struct {
	power              string
	minpoint, maxpoint utils.Point
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	// instr := getInstr(lines)
	p1, p2 := lights(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Took:", time.Since(start))
}

func lights(lines []string) (r1, r2 int) {
	g1 := make(map[utils.Point]bool)
	g2 := make(map[utils.Point]int)
	for _, line := range lines {
		var minx, miny, maxx, maxy int
		if f, _ := fmt.Sscanf(line, "turn on %d,%d through %d,%d", &minx, &miny, &maxx, &maxy); f == 4 {
			for x := minx; x <= maxx; x++ {
				for y := miny; y <= maxy; y++ {
					g1[utils.Point{x, y}] = true
					g2[utils.Point{x, y}]++
				}
			}
		} else if f, _ := fmt.Sscanf(line, "turn off %d,%d through %d,%d", &minx, &miny, &maxx, &maxy); f == 4 {
			for x := minx; x <= maxx; x++ {
				for y := miny; y <= maxy; y++ {
					g1[utils.Point{x, y}] = false
					g2[utils.Point{x, y}]--
					if g2[utils.Point{x, y}] < 0 {
						g2[utils.Point{x, y}] = 0
					}
				}
			}
		} else if f, _ := fmt.Sscanf(line, "toggle %d,%d through %d,%d", &minx, &miny, &maxx, &maxy); f == 4 {
			for x := minx; x <= maxx; x++ {
				for y := miny; y <= maxy; y++ {
					g1[utils.Point{x, y}] = !g1[utils.Point{x, y}]
					g2[utils.Point{x, y}] += 2
				}
			}
		}
	}
	for _, v := range g1 {
		if v {
			r1++
		}
	}
	for _, v := range g2 {
		r2 += v
	}
	return
}
