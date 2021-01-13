package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	tiles := getTiles(lines)
	t1 := time.Now()
	p1, tEdge := part1(tiles)
	fmt.Printf("Part 1: %d (%s)\n", p1, time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(tiles, tEdge), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
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

func (t *Tile) allOrient() []Tile {
	ret := make([]Tile, 12)
	p := *t
	for i := 0; i < 4; i++ {
		ret[i*3] = p
		ret[i*3+1] = p.flipLR()
		ret[i*3+2] = p.flipTB()

		p = p.turn()
	}
	return ret
}

func (t *Tile) orient(left, top string, tEdge map[string]int) *Tile {
	for _, o := range t.allOrient() {
		if ((left == "" && tEdge[o.left()] == 1) || left == o.left()) && ((top == "" && tEdge[o.top()] == 1) || top == o.top()) {
			return &o
		}
	}
	return nil
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

func getTiles(t []string) TileSet {
	tiles := make(TileSet, len(t))
	for _, v := range t {
		split := strings.SplitN(v, "\n", 2)
		var tID int
		fmt.Sscanf(split[0], "Tile %d:", &tID)
		s := strings.Split(split[1], "\n")
		size := len(s)
		tile := Tile{id: tID, grid: make([][]rune, size), size: size}
		for i, l := range s {
			tile.grid[i] = make([]rune, size)
			for j, c := range l {
				tile.grid[i][j] = c
			}
		}
		tiles[tID] = tile
	}
	return tiles
}

func (t *Tile) left() string {
	var left strings.Builder
	for i := 0; i < len(t.grid[0]); i++ {
		left.WriteRune(t.grid[i][0])
	}
	return left.String()
}

func (t *Tile) right() string {
	var right strings.Builder
	for i := 0; i < t.size; i++ {
		right.WriteRune(t.grid[i][t.size-1])
	}
	return right.String()
}

func (t *Tile) top() string {
	return string(t.grid[0])
}

func (t *Tile) bottom() string {
	return string(t.grid[t.size-1])
}

func reverse(s string) string {
	var ret strings.Builder
	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		ret.WriteRune(r[i])
	}
	return ret.String()
}

func (t *Tile) allEdge() []string {
	edges := []string{
		/*top*/ t.top(),
		/*bottom*/ t.bottom(),
		/*left*/ t.left(),
		/*right*/ t.right(),
	}

	return []string{
		edges[0],
		reverse(edges[0]),
		edges[1],
		reverse(edges[1]),
		edges[2],
		reverse(edges[2]),
		edges[3],
		reverse(edges[3]),
	}
}

func (t *Tile) flipTB() Tile {
	grid := make([][]rune, t.size)
	for y := range t.grid {
		grid[y] = t.grid[t.size-y-1]
	}
	return Tile{
		id:   t.id,
		grid: grid,
		size: t.size,
	}
}

func (t *Tile) flipLR() Tile {
	grid := make([][]rune, t.size)
	for y := range t.grid {
		grid[y] = make([]rune, t.size)
		for x, c := range t.grid[y] {
			grid[y][t.size-x-1] = c
		}
	}
	return Tile{
		id:   t.id,
		grid: grid,
		size: t.size,
	}
}

func (t *Tile) turn() Tile {
	grid := make([][]rune, t.size)
	for y := range t.grid {
		grid[y] = make([]rune, t.size)
	}
	for y := range t.grid {
		for x := range t.grid[y] {
			grid[y][t.size-x-1] = t.grid[x][y]
		}
	}
	return Tile{
		id:   t.id,
		grid: grid,
		size: t.size,
	}
}
