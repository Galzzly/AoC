package main

import (
	"container/heap"
	"fmt"
	"image"
	"math"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type pqi[T any] struct {
	v T
	p int
}

type PQ[T any] []pqi[T]

func (q PQ[_]) Len() int           { return len(q) }
func (q PQ[_]) Less(i, j int) bool { return q[i].p < q[j].p }
func (q PQ[_]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *PQ[T]) Push(x any)        { *q = append(*q, x.(pqi[T])) }
func (q *PQ[_]) Pop() (x any)      { x, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]; return x }
func (q *PQ[T]) GPush(v T, p int)  { heap.Push(q, pqi[T]{v, p}) }
func (q *PQ[T]) GPop() (T, int)    { x := heap.Pop(q).(pqi[T]); return x.v, x.p }

type State struct {
	Pos image.Point
	Dir image.Point
}

func main() {
	start := time.Now()
	f := "input.txt"
	grid, rect := utils.MakeIntImagePointMap(utils.ReadFileLineByLine(f))
	t1 := time.Now()
	fmt.Println("Part 1:", solve(1, 3, grid, rect), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(4, 10, grid, rect), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func solve(min, max int, G map[image.Point]int, R image.Rectangle) (res int) {
	Q := PQ[State]{}
	Seen := map[State]bool{}
	Q.GPush(State{image.Point{0, 0}, image.Point{1, 0}}, 0)
	Q.GPush(State{image.Point{0, 0}, image.Point{0, 1}}, 0)

	for len(Q) > 0 {
		state, H := Q.GPop()
		if state.Pos == R.Max {
			return H
		}
		if _, ok := Seen[state]; ok {
			continue
		}
		Seen[state] = true
		for i := -max; i <= max; i++ {
			np := state.Pos.Add(state.Dir.Mul(i))
			if _, ok := G[np]; !ok || i > -min && i < min {
				continue
			}
			var h int
			s := int(math.Copysign(1, float64(i)))
			for j := s; j != i+s; j += s {
				h += G[state.Pos.Add(state.Dir.Mul(j))]
			}
			Q.GPush(State{np, image.Point{state.Dir.Y, state.Dir.X}}, H+h)
		}
	}
	return -1
}
