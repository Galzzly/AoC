package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Tile struct {
	id   int
	grid [][]rune
	size int
}

type TileSet map[int]Tile

type ImageMap struct {
	contents            [][]rune
	size, height, width int
}

type Vector struct {
	x, y int
}

var Monster map[Vector]struct{}

func init() {
	Monster = make(map[Vector]struct{})
	pattern := "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "
	for y, l := range strings.Split(pattern, "\n") {
		for x, r := range l {
			if r == '#' {
				Monster[Vector{x: x, y: y}] = struct{}{}
			}
		}
	}
}

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileDoubleLine(f)
	tiles := getTiles(lines)
	t1 := time.Now()
	p1, tEdge := part1(tiles)
	fmt.Println("Part 1:", p1, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(tiles, tEdge), "Took:", time.Since(t2))

	fmt.Println("Total Time:", time.Since(start))
}

func part1(tiles TileSet) (int, map[string]int) {
	res := 1
	tEdge := make(map[string]int)
	for _, v := range tiles {
		for _, e := range v.allEdge() {
			tEdge[e]++
		}
	}
	for idx, v := range tiles {
		count := 0
		for _, e := range v.allEdge() {
			if tEdge[e] == 2 {
				count++
			}
		}
		if count == 4 {
			res *= idx
		}
	}
	return res, tEdge
}

func part2(tiles TileSet, tEdge map[string]int) int {
	// Got the tiles passed in
	// Stitch the tiles together
	var left, top string
	usedKeys := make(map[int]struct{})
	usedTilePos := make(map[Vector]Tile)
	var stitcher ImageMap
	position := Vector{}

	// Get the first corner
	var firstTile Tile
	for _, t := range tiles {
		count := 0
		for _, e := range t.allEdge() {
			if tEdge[e] == 2 {
				count++
			}
		}
		if count == 4 {
			firstTile = t
			break
		}
	}
	// Make sure that it's the right way around
	firstInRow := firstTile.orient("", "", tEdge)
	usedKeys[firstInRow.id] = struct{}{}
	usedTilePos[position] = *firstInRow
	stitcher = newImage(firstInRow.size-2, len(tiles))
	stitcher.tileAdd(position, firstInRow.borderless())

	// Next tile
	position = position.add(Vector{x: 1})
	left, top = firstInRow.right(), ""

	for len(usedKeys) < len(tiles) {
		found := false
		for idx, tile := range tiles {
			if _, ok := usedKeys[idx]; ok {
				continue
			}
			if c := tile.orient(left, top, tEdge); c != nil {
				usedKeys[c.id] = struct{}{}
				usedTilePos[position] = *c
				stitcher.tileAdd(position, c.borderless())

				position = position.add(Vector{x: 1})
				left = c.right()
				if next, ok := usedTilePos[position.add(Vector{y: -1})]; ok {
					top = string(next.grid[c.size-1])
				}
				found = true

				if firstInRow == nil {
					firstInRow = c
				}
				break
			}
		}

		if !found {
			position = Vector{x: 0, y: position.y + 1}
			left = ""

			top = firstInRow.bottom()
			firstInRow = nil
		}
	}
	//Turn that back into a tile
	stitched := Tile{
		id:   0,
		grid: stitcher.contents,
		size: stitcher.size * stitcher.height,
	}
	// Find the monsters
	found := false
	for _, t := range stitched.allOrient() {
		for y := range t.grid {
			for x := range t.grid[y] {
				match := true
				for m := range Monster {
					if y+m.y >= len(t.grid) {
						match = false
						break
					}

					if x+m.x >= len(t.grid[y+m.y]) {
						match = false
						break
					}

					if t.grid[y+m.y][x+m.x] != '#' {
						match = false
						break
					}
				}
				if !match {
					continue
				}
				found = true
				for m := range Monster {
					stitched.grid[y+m.y][x+m.x] = 'O'
				}
			}
		}
		if found {
			break
		}
	}

	//Count the waves
	w := 0
	for y := range stitched.grid {
		for x := range stitched.grid[y] {
			if stitched.grid[y][x] == '#' {
				w++
			}
		}
	}
	return w
}

func (v *Vector) add(a Vector) Vector {
	return Vector{
		x: v.x + a.x,
		y: v.y + a.y,
	}
}

func (i *ImageMap) tileAdd(pos Vector, t Tile) {
	for tY, r := range t.grid {
		iY := (i.size * pos.y) + tY
		for tX, c := range r {
			iX := (i.size * pos.x) + tX
			i.contents[iY][iX] = c
		}
	}
}

func (t *Tile) borderless() Tile {
	tile := Tile{id: t.id, grid: make([][]rune, t.size-2), size: t.size - 2}
	for y := range t.grid {
		if y == 0 || y == t.size-1 {
			continue
		}
		tile.grid[y-1] = make([]rune, tile.size)
		for x, c := range t.grid[y] {
			if x == 0 || x == t.size-1 {
				continue
			}
			tile.grid[y-1][x-1] = c
		}
	}
	return tile
}

func newImage(size, count int) ImageMap {
	width := int(math.Sqrt(float64(count)))
	height := width

	content := make([][]rune, size*height)
	for i := 0; i < size*height; i++ {
		content[i] = make([]rune, size*height)
	}

	return ImageMap{
		contents: content,
		size:     size,
		width:    width,
		height:   height,
	}
}
