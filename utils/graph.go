package utils

import (
	"image"
)

type Queue[T any] []T

type Graph interface {
	Neighbours(p image.Point) []image.Point
}

type Grid[T any] struct {
	x, y      int
	state     map[image.Point]T
	movements []image.Point
}

func NewGrid[T any](x, y int, movements []image.Point) *Grid[T] {
	state := make(map[image.Point]T)
	return &Grid[T]{
		x:         x,
		y:         y,
		state:     state,
		movements: movements,
	}
}

func (g *Grid[T]) IsValid(x, y int) bool {
	switch {
	case x < 0, x >= g.x, y < 0, y >= g.y:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) SetState(x, y int, state T) {
	if g.IsValid(x, y) {
		g.state[image.Point{x, y}] = state
	}
}

func (g *Grid[T]) GetState(p image.Point) T {
	return g.state[p]
}

func (g *Grid[T]) Neighbours(p image.Point) (res []image.Point) {
	for _, m := range g.movements {
		np := p.Add(m)
		if g.IsValid(np.X, np.Y) {
			res = append(res, np)
		}
	}
	return
}

func (q *Queue[T]) Put(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Get() T {
	ret := (*q)[0]
	*q = (*q)[1:]

	return ret
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func Search(g Graph, s, e image.Point) (res []image.Point) {
	var queue Queue[image.Point]
	queue.Put(s)

	from := map[image.Point]*image.Point{}
	from[s] = nil

	for !queue.Empty() {
		current := queue.Get()
		if current == e {
			break
		}
		for _, p := range g.Neighbours(current) {
			if _, ok := from[p]; !ok {
				queue.Put(p)
				from[p] = &current
			}
		}
	}

	res = []image.Point{e}
	for p := from[e]; p != nil; p = from[*p] {
		res = append(res, *p)
	}
	return
}
