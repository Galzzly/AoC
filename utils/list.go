package utils

type Node struct {
	value any
	prev  *Node
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList() List {
	return List{}
}

func (l *List) Populate(input []any) {
	for i := 0; i < len(input); i++ {
		l.Insert(input[i])
	}
	l.head.prev = l.tail
	l.tail.next = l.head
}

func (l *List) Insert(n any) {
	node := &Node{value: n, prev: nil, next: nil}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		node.prev = p
		p.next = node
		l.tail = node
	}
}
