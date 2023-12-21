package main

import (
	"fmt"
	"image"
	"slices"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Grid map[image.Point]rune

type Inputs struct {
	part2 bool
	steps int
}

var memories = map[TilePos]int{}

type TilePos struct {
	Tile image.Point
	Pos  image.Point
}
type Memory struct {
	Tile image.Point
	Pos  image.Point
	D    int
}

var check = map[Check]int{}

type Check struct {
	D      int
	corner bool
	steps  int
}

var Delta = []image.Point{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	grid, R, S := buildGrid(lines)
	findTilePos(grid, R, S)
	for i, b := range []Inputs{{false, 64}, {true, 26501365}} {
		timer := time.Now()
		fmt.Printf("Part %d: %d, Took: %v\n", i+1, solve(grid, R, S, b.steps, b.part2), time.Since(timer))
	}
	fmt.Println("Total Time:", time.Since(start))
}

func solve(G Grid, R image.Rectangle, S image.Point, steps int, part2 bool) (res int) {
	tDelta := []int{-3, -2, -1, 0, 1, 2, 3}
	min, max := utils.MinMax(tDelta)
	MM := []int{min, max}
	for yp := R.Min; yp.Y <= R.Max.Y; yp = yp.Add(image.Point{0, 1}) {
		for P := yp; P.X <= R.Max.X; P = P.Add(image.Point{1, 0}) {
			if _, ok := memories[TilePos{image.Point{0, 0}, P}]; ok {
				for _, X := range tDelta {
					for _, Y := range tDelta {
						if !part2 && (X != 0 || Y != 0) {
							continue
						}
						D := memories[TilePos{image.Point{X, Y}, P}]
						if (D%2 == steps%2) && (D <= steps) {
							res++
						}
						if slices.Contains(MM, X) && slices.Contains(MM, Y) {
							res += calculate(D, true, steps, R)
						} else if slices.Contains(MM, X) || slices.Contains(MM, Y) {
							res += calculate(D, false, steps, R)
						}
					}
				}
			}
		}
	}
	return
}

func calculate(D int, corner bool, steps int, R image.Rectangle) (res int) {
	if V, ok := check[Check{D, corner, steps}]; ok {
		return V
	}
	M := (steps - D) / (R.Max.Y + 1)
	for i := 1; i <= M; i++ {
		if D+(R.Max.Y+1)*i <= steps && (D+(R.Max.Y+1)*i)%2 == (steps%2) {
			res++
			if corner {
				res += i
				continue
			}
		}
	}
	check[Check{D, corner, steps}] = res
	return
}

func findTilePos(G Grid, R image.Rectangle, S image.Point) {
	Q := []Memory{{image.Point{0, 0}, S, 0}}
	for len(Q) > 0 {
		next := []Memory{}
		for _, q := range Q {
			tp := TilePos{q.Tile, q.Pos}
			if q.Pos.Y < R.Min.Y {
				tp.Tile.Y -= 1
				tp.Pos.Y += (R.Max.Y + 1)
			}
			if q.Pos.Y > R.Max.Y {
				tp.Tile.Y += 1
				tp.Pos.Y -= (R.Max.Y + 1)
			}
			if q.Pos.X < R.Min.X {
				tp.Tile.X -= 1
				tp.Pos.X += (R.Max.X + 1)
			}
			if q.Pos.X > R.Max.X {
				tp.Tile.X += 1
				tp.Pos.X -= (R.Max.X + 1)
			}
			if v, ok := G[tp.Pos]; !ok || v == '#' {
				continue
			}
			if _, ok := memories[tp]; ok {
				continue
			}
			if utils.Abs(tp.Tile.X) > 3 || utils.Abs(tp.Tile.Y) > 3 {
				continue
			}
			memories[tp] = q.D
			for _, p := range Delta {
				next = append(next, Memory{tp.Tile, tp.Pos.Add(p), q.D + 1})
			}
		}
		Q = next
	}
}

func buildGrid(lines []string) (Grid, image.Rectangle, image.Point) {
	grid, R := utils.MakeImagePointMapRect(lines)
	S, _ := utils.MapKey[image.Point, rune](grid, 'S')
	return grid, R, S
}
