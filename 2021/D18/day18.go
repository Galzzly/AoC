package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type visitfunc func(s *snailfish, d int) error

type snailfish struct {
	id     *int
	parent *snailfish
	left   *snailfish
	right  *snailfish
}

func main() {
	start := time.Now()
	lines := utils.ReadFileLineByLine("input")
	snailfish := parseSnailfish(lines)
	t2 := time.Now()
	p2 := part2(snailfish)
	s2 := time.Since(t2)
	t1 := time.Now()
	fmt.Printf("Part 1: %d in %s\n", part1(snailfish), time.Since(t1))
	fmt.Printf("Part 2: %d in %s\n", p2, s2)
	fmt.Printf("Total: %s\n", time.Since(start))
}

func part1(s []*snailfish) int64 {
	var result *snailfish = s[0]
	for i := 1; i < len(s); i++ {
		result = addSnailfish(result, s[i])
	}
	return magnitude(result)
}

func part2(s []*snailfish) (res int64) {
	res = math.MinInt64
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			prev, cur := copyTree(s[i], nil), copyTree(s[j], nil)
			ns := addSnailfish(prev, cur)
			m := magnitude(ns)
			if m > res {
				res = m
			}
		}
	}
	return
}

func copyTree(s *snailfish, p *snailfish) *snailfish {
	if s == nil {
		return nil
	}
	r := &snailfish{id: s.id, parent: p}
	r.left = copyTree(s.left, r)
	r.right = copyTree(s.right, r)
	return r
}

func magnitude(s *snailfish) int64 {
	m := int64(0)
	if s.left != nil {
		if s.left.id != nil {
			m += 3 * int64(*s.left.id)
		} else {
			m += 3 * magnitude(s.left)
		}
	}
	if s.right != nil {
		if s.right.id != nil {
			m += 2 * int64(*s.right.id)
		} else {
			m += 2 * magnitude(s.right)
		}
	}

	return m
}

func addSnailfish(a, b *snailfish) *snailfish {
	sn := &snailfish{left: a, right: b}
	a.parent = sn
	b.parent = sn
	sn = reduce(sn)
	return sn
}

func reduce(s *snailfish) *snailfish {
	for {
		reduced := false
		exp := toExplode(s)
		for exp != nil {
			reduced = true
			explode(exp)
			exp = toExplode(s)
		}

		sp := toSplit(s)
		if sp != nil {
			reduced = true
			split(sp)
		}

		if !reduced {
			break
		}
	}

	return s
}

func explode(s *snailfish) {
	nextLeft := getNextLeft(s)
	if nextLeft != nil {
		if nextLeft.id == nil {
			nextLeft = rightLeaf(nextLeft)
		}
		lid := addIds(s.left, nextLeft)
		nextLeft.id = lid
	}

	nextRight := getNextRight(s)
	if nextRight != nil {
		if nextRight.id == nil {
			nextRight = leftLeaf(nextRight)
		}
		rid := addIds(s.right, nextRight)
		nextRight.id = rid
	}

	newid := 0
	ns := &snailfish{id: &newid, parent: s.parent}
	if s == s.parent.right {
		ns.parent.right = ns
	} else if s == s.parent.left {
		ns.parent.left = ns
	}

	removeNode(s.left)
	removeNode(s.right)
	removeNode(s)
}

func removeNode(s *snailfish) {
	if s == nil {
		return
	}
	if s == s.parent.left {
		s.parent.left = nil
	}
	if s == s.parent.right {
		s.parent.right = nil
	}
	s.parent = nil
	s.left = nil
	s.right = nil
}

func split(s *snailfish) {
	val := *s.id

	fval := float64(val) / 2
	v1, v2 := int(math.Floor(fval)), int(math.Ceil(fval))

	left := &snailfish{id: &v1, parent: s}
	right := &snailfish{id: &v2, parent: s}

	s.id = nil
	s.left = left
	s.right = right
}

func leftLeaf(p *snailfish) *snailfish {
	for p.left != nil {
		p = p.left
	}
	return p
}

func rightLeaf(p *snailfish) *snailfish {
	for p.right != nil {
		p = p.right
	}
	return p
}

func getNextLeft(p *snailfish) *snailfish {
	orig := p
	for p != nil {
		p = p.parent
		if p != nil && p.left != nil && p.left != orig {
			return p.left
		}
		orig = p
	}
	return nil
}

func getNextRight(p *snailfish) *snailfish {
	orig := p
	for p != nil {
		p = p.parent
		if p != nil && p.right != nil && p.right != orig {
			return p.right
		}
		orig = p
	}
	return nil
}

func addIds(a, b *snailfish) *int {
	if a == nil || b == nil || a.id == nil || b.id == nil {
		return nil
	}
	s := *a.id + *b.id
	return &s
}

func toSplit(s *snailfish) *snailfish {
	var sp *snailfish
	traverse(s, func(c *snailfish, d int) error {
		if c.id != nil && *c.id > 9 && sp == nil {
			sp = c
			return fmt.Errorf("traverse no more")
		}
		return nil
	}, 0)
	return sp
}

func toExplode(s *snailfish) *snailfish {
	var exp *snailfish
	traverse(s, func(c *snailfish, d int) error {
		if d >= 4 && c.left != nil && c.left.id != nil && c.right != nil && c.right.id != nil {
			exp = c
			return fmt.Errorf("traverse no more")
		}
		return nil
	}, 0)

	return exp
}

func traverse(s *snailfish, v visitfunc, depth int) error {
	err := v(s, depth)
	if err != nil {
		return err
	}
	if s.left != nil {
		err = traverse(s.left, v, depth+1)
		if err != nil {
			return err
		}
	}
	if s.right != nil {
		err = traverse(s.right, v, depth+1)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseSnailfish(lines []string) (res []*snailfish) {
	for _, line := range lines {
		res = append(res, getSnailfish(line))
	}
	return
}

func getSnailfish(line string) (res *snailfish) {
	res = &snailfish{}
	cur := res
	for _, c := range line {
		switch c {
		case '[':
			n := &snailfish{parent: cur}
			cur.left = n
			cur = n
		case ']':
			cur = cur.parent
		case ',':
			n := &snailfish{parent: cur}
			cur.right = n
			cur = n
		default:
			id := utils.Atoi(string(c))
			cur.id = &id
			cur = cur.parent
		}
	}
	return
}
