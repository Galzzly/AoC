package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	jig "github.com/Galzzly/AoC/2020/D20-2"
	"github.com/Galzzly/AoC/utils"
)

var Monster map[jig.Vector]struct{}

func init() {
	Monster = make(map[jig.Vector]struct{})
	pattern := "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "
	for y, l := range strings.Split(pattern, "\n") {
		for x, r := range l {
			if r == '#' {
				Monster[jig.Vector{x: x, y: y}] = struct{}{}
			}
		}
	}
}

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileDoubleLine(f)
	tiles := jig.GetTiles(lines)
	t1 := time.Now()
	p1, tEdge := part1(tiles)
	fmt.Println("Part 1:", p1, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(tiles, tEdge), "Took:", time.Since(t2))

	fmt.Println("Total Time:", time.Since(start))
}

func part1(tiles jig.TileSet) (int, map[string]int) {
	res := 1
	tEdge := make(map[string]int)
	for _, v := range tiles {
		for _, e := range v.AllEdge() {
			tEdge[e]++
		}
	}
	for idx, v := range tiles {
		count := 0
		for _, e := range v.AllEdge() {
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

func part2(tiles jig.TileSet, tEdge map[string]int) int {
	var left, top string
	usedKeys := make(map[int]struct{})
	usedTilePos := make(map[jig.Vector]jig.Tile)
	var stitcher jig.ImageMap
	position := jig.Vector{}

	var firstTile jig.Tile
	for _, t := range tiles {
		count := 0
		for _, e := range t.AllEdge() {
			if tEdge[e] == 2 {
				count++
			}
		}
		if count == 4 {
			firstTile = t
			break
		}
	}

	firstInRow := firstTile.Orient("", "", tEdge)
	usedKeys[firstInRow.id] = struct{}{}
	usedTilePos[position] = *firstInRow
	stitcher = newImage(firstInRow.size-2, len(tiles))
	stitcher.tileAdd(position, firstInRow.Borderless())
}

func newImage(size, count int) jig.ImageMap {
	width := int(math.Sqrt(float64(count)))
	height := width

	content := make([][]rune, size*height)
	for i := 0; i < size*height; i++ {
		content[i] = make([]rune, size*height)
	}

	return jig.ImageMap{
		contents: content,
		size:     size,
		width:    width,
		height:   height,
	}
}

func (i *jig.ImageMap) tileAdd(pos jig.Vector, t jig.Tile) {
	for tY, r := range t.grid {
		iY := (i.size * pos.y) + tY
		for tX, c := range r {
			iX := (i.size * pos.x) + tX
			i.contents[iY][iX] = c
		}
	}
}

func (t *jig.Tile) Borderless() jig.Tile {
	tile := jig.Tile{id: t.id, grid: make([][]rune, t.size-2), size: t.size - 2}
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
