package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var Delta = []utils.Point{
	{X: -1, Y: -1}, {X: -1, Y: 0}, {X: -1, Y: 1},
	{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1},
	{X: 1, Y: -1}, {X: 1, Y: 0}, {X: 1, Y: 1},
}

func main() {
	start := time.Now()
	lines := utils.ReadFileDoubleLine("input")
	enhancement := lines[0]
	inputImage := buildImage(lines[1])
	t1 := time.Now()
	fmt.Printf("Part 1: %d in %s\n", enhance(enhancement, inputImage, 2), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d in %s\n", enhance(enhancement, inputImage, 50), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func enhance(enhancement string, inputImage map[utils.Point]bool, iter int) (res int) {
	for i := 0; i < iter; i++ {
		output := make(map[utils.Point]bool)
		var maxx, maxy int
		var minx, miny = math.MaxInt, math.MaxInt
		for p := range inputImage {
			_, maxx = utils.MinMax([]int{maxx, p.X})
			_, maxy = utils.MinMax([]int{maxy, p.Y})
			minx, _ = utils.MinMax([]int{minx, p.X})
			miny, _ = utils.MinMax([]int{miny, p.Y})
		}
		for x := minx - 1; x <= maxx+1; x++ {
			for y := miny - 1; y <= maxy+1; y++ {
				var arg int
				for _, d := range Delta {
					arg = arg << 1
					if v, ok := inputImage[utils.Point{X: x + d.X, Y: y + d.Y}]; ok {
						if v {
							arg |= 1
						}
						continue
					}
					if i%2 == 1 {
						arg |= 1
					}

				}
				output[utils.Point{X: x, Y: y}] = enhancement[arg] == '#'
			}
		}
		inputImage = output

	}
	res = count(inputImage)
	return
}
func count(in map[utils.Point]bool) (res int) {
	for _, i := range in {
		if i {
			res++
		}
	}
	return
}

func buildImage(line string) (res map[utils.Point]bool) {
	res = map[utils.Point]bool{}
	lines := strings.Split(strings.TrimSpace(line), "\n")

	for x, l := range lines {
		for y, c := range l {
			res[utils.Point{X: x, Y: y}] = rune(c) == '#'
		}
	}
	return
}
