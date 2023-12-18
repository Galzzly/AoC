package main

import (
	"fmt"
	"image"
	"strconv"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Instr struct {
	Dir  string
	Dist int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	for i, b := range []bool{false, true} {
		timer := time.Now()
		fmt.Printf("Part %d: %d Took: %v\n", i+1, solve(lines, b), time.Since(timer))
	}
	fmt.Println("Total Time:", time.Since(start))
}

type Points struct {
	a, b image.Point
}

func solve(lines []string, part2 bool) (res int) {
	var minX, maxX, minY, maxY int
	P := []Points{}
	for _, line := range lines {
		var D, H string
		var N int
		fmt.Sscanf(line, "%s %d (#%s)", &D, &N, &H)
		if part2 {
			switch H[5] {
			case '0':
				D = "R"
			case '1':
				D = "D"
			case '2':
				D = "L"
			case '3':
				D = "U"
			}
			val, _ := strconv.ParseInt(H[:5], 16, 64)
			N = int(val)
		}
		switch D {
		case "R":
			maxX = minX + N
		case "L":
			maxX = minX - N
		case "D":
			maxY = minY + N
		case "U":
			maxY = minY - N
		}
		P = append(P, Points{image.Point{minX, minY}, image.Point{maxX, maxY}})
		minX, minY = maxX, maxY
	}
	for _, p := range P {
		res += (p.a.Y + p.b.Y) * (p.a.X - p.b.X)
		if p.a.X == p.b.X {
			res += utils.Abs(p.a.Y - p.b.Y)
			continue
		}
		res += utils.Abs(p.a.X - p.b.X)
	}

	res = res/2 + 1
	return
}
