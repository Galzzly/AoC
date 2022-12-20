package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Node struct {
	value      int
	prev, next *Node
}

type List struct {
	head *Node
	tail *Node
}

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	nums := utils.ReadIntsByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(nums), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(nums), "- Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func part1(n []int) (res int) {
	nodes := makeNodes(n, 1)
	ret := solve(nodes, 1)
	node := ret
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			node = node.next
		}
		res += node.value
	}
	return
}

func part2(n []int) (res int) {
	nodes := makeNodes(n, 811589153)
	ret := solve(nodes, 10)
	node := ret
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			node = node.next
		}
		res += node.value
	}
	return
}

func solve(l List, r int) (res *Node) {
	var nodes []*Node
	n := l.head
	for {
		nodes = append(nodes, n)
		if n == l.tail {
			break
		}
		n = n.next
	}
	nodeLen := len(nodes) - 1
	for round := 0; round < r; round++ {
		for _, node := range nodes {
			moves := node.value % nodeLen
			if moves > 0 {
				for i := 0; i < moves; i++ {
					r := node.next
					node.prev.next = r
					r.prev = node.prev
					r.next.prev = node
					node.next = r.next
					r.next, node.prev = node, r
				}
			} else if moves < 0 {
				for i := 0; i < -moves; i++ {
					p := node.prev
					node.next.prev = p
					p.next = node.next
					p.prev.next = node
					node.prev = p.prev
					node.next, p.prev = p, node
				}
			} else {
				res = node
			}
		}
	}
	return
}

func makeNodes(nums []int, decrypt int) (nodes List) {
	nodes = List{}
	for i := 0; i < len(nums); i++ {
		nodes.Insert(nums[i] * decrypt)
	}
	nodes.head.prev = nodes.tail
	nodes.tail.next = nodes.head
	return
}

func (l *List) Insert(n int) {
	list := &Node{value: n, prev: nil, next: nil}
	if l.head == nil {
		l.head = list
		l.tail = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		list.prev = p
		p.next = list
		l.tail = list
	}

}

// func insertInt(array []int, value int, index int) []int {
// 	return append(array[:index], append([]int{value}, array[index:]...)...)
// }

// func removeInt(array []int, index int) []int {
// 	return append(array[:index], array[index+1:]...)
// }

// func moveInt(array []int, srcIndex int, dstIndex int) []int {
// 	value := array[srcIndex]
// 	return insertInt(removeInt(array, srcIndex), value, dstIndex)
// }
