package graph

import (
	"container/heap"
	"fmt"
	"time"
)

type CostedNode[T] struct {
	Node T
	Cost int
}

type THeap[T] []CostedNode[T]

func (t THeap[T]) Len() int { return len(t) }
func (t THeap[T]) Less(i, j int) bool { return t[i].Cost < t[j].Cost }
func (t THeap[T]) Swap(i, j int) {t[i], t[j] = t[j], t[i]}

func (t *THeap[T]) Push(x interface{}) {
	*t = append(*t, x.(CostedNode[T]))
}

func (t *THeap[T]) Pop() interface{} {
	old := *t
	n := len(old)
	x := old[n-1]
	*t = old[0 : n-1]
	return x
}

type Traversable[T comparable] interface {
	Neighbors(node T) []CostedNode[T]
	String(node T) string
}

func Dijkstra[T comparable](graph Traversable[T], start T, stop T) (int, []CostedNode[T]) {
	costs := make(map[T]int)
	prev := make(map[T]T)

	fringe := THeap[T]{CostedNode[T]{start, 0}}
	heap.Init(&fringe)

	startT := time.Now()
	n := 0
	for len(fringe) > 0 {
		cur := heap.Pop(&fringe).(CostedNode[T])
		node, cost := cur.Node, cur.Cost
		if node == stop {
			path := []CostedNode[T]{{Node: node, Cost: cost}}
			for path[len(path)-1].Node != start {
				prevNode, ok := prev[node]
				if !ok {
					panic(node)
				}
				node = prevNode
				path = append(path, CostedNode[T]{Node: node, Cost: costs[node]})
			}
			for i, j := 0, len(path) - 1; i < j; i, j = i + 1, j - 1 {
				path[i], path[j] = path[j], path[i]
			}
			return cost, path
		}

		n++
		if n % 100000 == 0 {
			fmt.Printf("%d\n", n)
		}
		if n % 1000000 == 0 {
			fmt.Printf("%d\n", n)
			fmt.Printf("%s\n", time.Since(startT))
			startT = time.Now()
		}

		for _, n := range graph.Neighbors(node) {
			next := n.Node
			nextCost := cost + n.Cost
			if _, ok := costs[next]; !ok || nextCost < costs[next] {
				costs[next] = nextCost
				prev[next] = node
				heap.Push(&fringe, CostedNode[T]{next, nextCost})
			}
		}
	}

	return -1, nil
}