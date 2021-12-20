package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Cube struct {
	X, Y, Z int
}

// type scanner struct {
// 	id      int
// 	coord   Cube
// 	beacons []Cube
// }

type beacons []Cube

func main() {
	start := time.Now()
	lines := utils.ReadFileDoubleLine("input")
	beacons := parseScanners(lines)
	r1, r2 := scanning(beacons)
	fmt.Printf("Part 1: %d\n", r1)
	fmt.Printf("Part 2: %d\n", r2)
	fmt.Printf("Total: %s\n", time.Since(start))
}

func scanning(beacons []beacons) (p1res, p2res int) {
	var beaconMatch = map[Cube]bool{}
	scanners := []Cube{{0, 0, 0}}
	for _, p := range beacons[0] {
		beaconMatch[p] = true
	}
	beacons = beacons[1:]
	for len(beacons) > 0 {
	rescan:
		for i := len(beacons) - 1; i >= 0; i-- {
			for rId := 0; rId < 24; rId++ {
				offsets := map[Cube]int{}
				for known := range beaconMatch {
					for _, p := range beacons[i] {
						offset := p.Rotate(rId).Sub(known)
						offsets[offset]++
					}
				}
				for offset, count := range offsets {
					if count >= 12 {
						scanner := offset.Invert()
						scanners = append(scanners, scanner)

						for _, p := range beacons[i] {
							beaconMatch[p.Rotate(rId).Add(scanner)] = true
						}
						beacons = append(beacons[:i], beacons[i+1:]...)
						//fmt.Println("here", rId, scanners[i].id, count)
						continue rescan
					}
				}
			}
		}
	}
	// carryon:
	p1res = len(beaconMatch)
	for _, s1 := range scanners {
		for _, s2 := range scanners {
			if s1.Dist(s2) > p2res {
				p2res = s1.Dist(s2)
			}
		}
	}
	return
}

func (s1 Cube) Dist(s2 Cube) (res int) {
	res = utils.Abs(s1.X-s2.X) + utils.Abs(s1.Y-s2.Y) + utils.Abs(s1.Z-s2.Z)
	return
}

/**
	Possible Orientations
		Around X: [{x, y, z}, {x, -z, y}, {x, -y, -z}, {x, z, -y},
					{-x, -y, z}, {-x, -z, -y}, {-x, y, -z}, {-x, z, y}]
		Around Y: [{y, x, -z}, {y, -x, -z}, {y, z, x}, {y, -z, -x},
					{-y, x, z}, {-y, -x, -z}, {-y, -z, x}, {-y, z, -x}]
		Around Z: [{z, x, y}, {z, -x, -y}, {z, -y, x}, {z, y, -x},
					{-z, x, -y}, {-z, -x, y}, {-z, -y, -x}, {-z, y, x}]
**/

func (b Cube) Rotate(id int) Cube {
	switch id {
	// X axis
	case 0:
		return Cube{b.X, b.Y, b.Z}
	case 1:
		return Cube{b.X, -b.Z, b.Y}
	case 2:
		return Cube{b.X, -b.Y, -b.Z}
	case 3:
		return Cube{b.X, b.Z, -b.Y}
	case 4:
		return Cube{-b.X, -b.Y, b.Z}
	case 5:
		return Cube{-b.X, -b.Z, -b.Y}
	case 6:
		return Cube{-b.X, b.Y, -b.Z}
	case 7:
		return Cube{-b.X, b.Z, b.Y}
	// Y axis
	case 8:
		return Cube{b.Y, b.X, -b.Z}
	case 9:
		return Cube{b.Y, -b.X, b.Z}
	case 10:
		return Cube{b.Y, b.Z, b.X}
	case 11:
		return Cube{b.Y, -b.Z, -b.X}
	case 12:
		return Cube{-b.Y, b.X, b.Z}
	case 13:
		return Cube{-b.Y, -b.X, -b.Z}
	case 14:
		return Cube{-b.Y, -b.Z, b.X}
	case 15:
		return Cube{-b.Y, b.Z, -b.X}
	// Z axis
	case 16:
		return Cube{b.Z, b.X, b.Y}
	case 17:
		return Cube{b.Z, -b.X, -b.Y}
	case 18:
		return Cube{b.Z, -b.Y, b.X}
	case 19:
		return Cube{b.Z, b.Y, -b.X}
	case 20:
		return Cube{-b.Z, b.X, -b.Y}
	case 21:
		return Cube{-b.Z, -b.X, b.Y}
	case 22:
		return Cube{-b.Z, b.Y, b.X}
	case 23:
		return Cube{-b.Z, -b.Y, -b.X}
	// Force a panic...
	default:
		panic(id)
	}
}

func (b Cube) Sub(c Cube) Cube {
	return Cube{
		X: b.X - c.X,
		Y: b.Y - c.Y,
		Z: b.Z - c.Z,
	}
}

func (b Cube) Add(c Cube) Cube {
	return Cube{
		X: b.X + c.X,
		Y: b.Y + c.Y,
		Z: b.Z + c.Z,
	}
}

func (b Cube) Invert() Cube {
	return Cube{-b.X, -b.Y, -b.Z}
}

func parseScanners(lines []string) (res []beacons) {
	for _, line := range lines {
		sp := strings.Split(line, "\n")
		var id int
		fmt.Sscanf(sp[0], "--- scanner %d:", &id)
		var beacons []Cube
		for _, b := range sp[1:] {
			var x, y, z int
			fmt.Sscanf(b, "%d,%d,%d", &x, &y, &z)
			beacons = append(beacons, Cube{X: x, Y: y, Z: z})
		}
		// if id == 0 {
		res = append(res, beacons)
		// } else {
		// 	res = append(res, &scanner{id: id, beacons: beacons})
		// }
	}
	return
}
