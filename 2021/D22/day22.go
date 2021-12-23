package main

import (
	"fmt"
	"time"
	"sort"

	"github.com/Galzzly/AoC/utils"
)

// type cube struct {
// 	x, y, z int
// }

// var steps map[int]map[cube]bool

// var cuboids map[cube]bool

func main() {
	start := time.Now()
	lines := utils.ReadFileLineByLine("input")
	// steps = map[string][]cube{}
	// getSteps(lines)
	// cuboids = map[cube]bool{}
	// fmt.Println(lines)
	r1, r2 := reboot(lines)
	fmt.Printf("Part 1: %d\n", r1)
	fmt.Printf("Part 2: %d\n", r2)
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

// func part1(lines []string) (res int) {
// 	for _, line := range lines {
// 		// fmt.Println(i)
// 		var status string
// 		var minx, maxx, miny, maxy, minz, maxz int
// 		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &status, &minx, &maxx, &miny, &maxy, &minz, &maxz)
// 		if minx < -50 || maxx > 50 || miny < -50 || maxy > 50 || minz < -50 || maxz > 50 {
// 			continue
// 		}
// 		for x := minx; x <= maxx; x++ {
// 			for y := miny; y <= maxy; y++ {
// 				for z := minz; z <= maxz; z++ {
// 					c := cube{x, y, z}
// 					cuboids[c] = status == "on"

// 				}
// 			}
// 		}
// 	}
// 	for _, s := range cuboids {
// 		if s {
// 			res++
// 		}
// 	}
// 	// res = len(cuboids)
// 	return
// }

type Range struct {
	min int
	max int
}

type cuberange struct {
	x Range
	y Range
	z Range
}

type tracking struct {
	on bool
	c cuberange
}

func reboot(lines []string) (r1, r2 int64) {
	var tracker []tracking
	x, y, z := []int{}, []int{}, []int{}
	for _, line := range lines {
		var tr tracking
		var status string
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &status, &tr.c.x.min, &tr.c.x.max, &tr.c.y.min, &tr.c.y.max, &tr.c.z.min, &tr.c.z.max)
		tr.on = status == "on"
		tr.c.x.max++
		tr.c.y.max++
		tr.c.z.max++
		tracker = append(tracker, tr)
		x = append(x, tr.c.x.min, tr.c.x.max)
		y = append(y, tr.c.y.min, tr.c.y.max)
		z = append(z, tr.c.z.min, tr.c.z.max)
	}
	// x := append(make([]int, 0, 2*len(tracker)+2), -50, 51)
	// y := append(make([]int, 0, 2*len(tracker)+2), -50, 51)
	// z := append(make([]int, 0, 2*len(tracker)+2), -50, 51)

	// for _, tr := range tracker {
	// 	x = append(x, tr.c.x.min, tr.c.x.max)
	// 	y = append(y, tr.c.y.min, tr.c.y.max)
	// 	z = append(z, tr.c.z.min, tr.c.z.max)
	// }
	x = utils.SortUniqInts(x)
	y = utils.SortUniqInts(y)
	z = utils.SortUniqInts(z)

	buf := make([]bool, len(x)*len(y)*len(z))

	for _, tr := range tracker {
		minz := sort.SearchInts(z, tr.c.z.min)
		maxz := sort.SearchInts(z, tr.c.z.max)
		miny := sort.SearchInts(y, tr.c.y.min)
		maxy := sort.SearchInts(y, tr.c.y.max)
		minx := sort.SearchInts(x, tr.c.x.min)
		maxx := sort.SearchInts(x, tr.c.x.max)
		for trz := minz; trz < maxz; trz ++ {
			off := trz * len(x) * len(y)
			for try := miny; try < maxy; try++ {
				off := try*len(x) + off
				for trx := minx; trx < maxx; trx++ {
					buf[off+trx] = tr.on
				}
			}
		}
	}

	for trz := 0; trz < len(z)-1; trz++ {
		off := trz * len(x) * len(y)
		r1test := z[trz] > -51 && z[trz] < 51
		cubes := int64(z[trz+1] - z[trz])
		for try := 0; try < len(y)-1; try++ {
			off := try * len(x)+off
			r1test := r1test && y[try] > -51 && y[try] < 51
			cubes := cubes * int64(y[try+1] - y[try])
			for trx := 0; trx < len(x)-1;trx++ {
				if !buf[off+trx] {
					continue
				}
				r1test := r1test && x[trx] > -51 && x[trx] < 51
				cubes := cubes * int64(x[trx+1] - x[trx])
				r2 += cubes
				if r1test {
					r1 += cubes
				}
			}
		}
	}

	return
}
