package main

import (
	"fmt"
	"strings"
)

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

func (t *Tile) orient(left, top string, tEdge map[string]int) *Tile {
	for _, o := range t.allOrient() {
		if ((left == "" && tEdge[o.left()] == 1) || left == o.left()) && ((top == "" && tEdge[o.top()] == 1) || top == o.top()) {
			return &o
		}
	}
	return nil
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
