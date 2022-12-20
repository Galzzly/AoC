package main

import (
	"fmt"
	"image"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Sensor image.Point

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileLineByLine(f)
	sensors, beacons := getSensors(lines)
	// fmt.Println(part1(sensors, 10))
	// fmt.Println(part2(sensors, 20))
	t1 := time.Now()
	// fmt.Println("Part 1:", part1(sensors, beacons, 2000000), "Took:", time.Since(t1))
	fmt.Println("Part 1:", part1(sensors, beacons, 99), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(sensors, 4000000), "Took:", time.Since(t2))
	// fmt.Println("Part 2:", part2(sensors, 104), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(sensors map[Sensor]int, beacons map[Sensor]bool, row int) (res int) {
	counter := map[image.Point]bool{}
	for p := range sensors {
		if row >= p.Y-sensors[p] && row <= p.Y+sensors[p] {
			delta := sensors[p] - utils.Abs(p.Y-row)
			for x := p.X - delta; x <= p.X+delta; x++ {
				if _, ok := sensors[Sensor{x, row}]; !ok {
					if _, ok := beacons[Sensor{x, row}]; !ok {
						counter[image.Point{x, row}] = true
					}
				}
			}
		}
	}
	res = len(counter)
	return
}

func part2(sensors map[Sensor]int, max int) (res int) {
	unseen := findUnseenPoints(sensors, image.Point{0, 0}, image.Point{max, max}, 0)
	res = (unseen.X * 4000000) + unseen.Y
	return
}

func findUnseenPoints(sensors map[Sensor]int, min, max image.Point, count int) image.Point {
	count++
	fmt.Println("P2", count)
	if min == max {
		return min
	}

	mid := image.Point{(min.X + max.X) / 2, (min.Y + max.Y) / 2}
	quads := make([][]image.Point, 4)
	quads[0] = []image.Point{min, mid}
	quads[1] = []image.Point{{mid.X + 1, min.Y}, {max.X, mid.Y}}
	quads[2] = []image.Point{{min.X, mid.Y + 1}, {mid.X, max.Y}}
	quads[3] = []image.Point{{mid.X + 1, mid.Y + 1}, max}

	for _, q := range quads {
		if q[0].X > q[1].X || q[0].Y > q[1].Y {
			continue
		}
		allPairsCanContain := true
		for s, m := range sensors {
			if !s.canContain(q[0], q[1], m) {
				allPairsCanContain = false
				break
			}
		}
		if allPairsCanContain {
			k := findUnseenPoints(sensors, q[0], q[1], count)
			if k.X != -1 || k.Y != -1 {
				return k
			}
		}
	}
	return image.Point{-1, -1}
}

func (s Sensor) canContain(min, max image.Point, d int) bool {
	corners := []image.Point{min, {min.X, max.Y}, {max.X, min.Y}, max}
	for _, c := range corners {
		if utils.Abs(s.X-c.X)+utils.Abs(s.Y-c.Y) > d {
			return true
		}
	}
	return false
}

func getSensors(lines []string) (sensors map[Sensor]int, beacons map[Sensor]bool) {
	sensors = make(map[Sensor]int, len(lines))
	beacons = make(map[Sensor]bool, len(lines))
	for _, line := range lines {
		var s, b Sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.X, &s.Y, &b.X, &b.Y)
		sensors[s] = utils.Abs(s.X-b.X) + utils.Abs(s.Y-b.Y)
		beacons[b] = true
	}
	return
}
