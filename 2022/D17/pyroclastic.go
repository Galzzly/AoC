package main

import (
	"fmt"
	"image"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var Move = map[rune]int{
	'<': -1,
	'>': 1,
}

type Key struct {
	s int
	m int
	d [7]int
}

type State struct {
	h, stopped int
}

type Screen []uint8

type Shape struct {
	p      image.Point
	points []image.Point
	h, w   int
}

var Shapes = []Shape{
	{points: []image.Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}}, h: 1, w: 4},
	{points: []image.Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}, h: 3, w: 3},
	{points: []image.Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, h: 3, w: 3},
	{points: []image.Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}}, h: 4, w: 3},
	{points: []image.Point{{0, 0}, {1, 0}, {0, 1}, {1, 1}}, h: 2, w: 2},
}

func main() {
	start := time.Now()

	f := "input.txt"
	// f = "test"

	input := []rune(utils.ReadFileSingleLine(f))
	t1 := time.Now()
	fmt.Println("Part 1:", solve(input, 2022), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(input, 1000000000000), " - Took:", time.Since(t2))
	fmt.Println(time.Since(start))
}

func solve(input []rune, lim int) (res int) {
	checker := true
	count := lim
	highPoint, cycleHeight := 0, 0
	shapeI, moveI := 0, 0
	lenShape := len(Shapes)
	lenMoves := len(input)
	s := make(Screen, 1)
	s[0] = (1 << 7) - 1
	seen := make(map[Key]State)
	for {
		if count == 0 {
			res = highPoint + cycleHeight
			break
		}
		shape := Shapes[shapeI]
		shapeI = (shapeI + 1) % lenShape
		shape.p.X = 2
		shape.p.Y = highPoint + 4
		if len(s) < shape.p.Y+shape.h+1 {
			nextScreen := make([]uint8, (shape.p.Y+shape.h+1)*100000)
			copy(nextScreen, s)
			s = nextScreen
		}

		for {
			move := Move[input[moveI]]
			moveI = (moveI + 1) % lenMoves

			if shape.canMove(image.Point{move, 0}, s) {
				shape.p.X += move
			}

			if shape.canMove(image.Point{0, -1}, s) {
				shape.p.Y--
			} else {
				s.draw(shape)
				sHigh := shape.p.Y + shape.h - 1
				highPoint = utils.Biggest(sHigh, highPoint)
				count--
				// count++
				if checker {
					d := s.depth(highPoint)
					k := Key{shapeI, moveI, d}
					if prev, ok := seen[k]; ok {
						diff := highPoint - prev.h
						cLen := prev.stopped - count
						cRem := count / cLen
						cycleHeight = diff * cRem
						count = count % cLen
						checker = false
					} else {
						seen[k] = State{highPoint, count}
					}
				}
				break
			}
		}
	}

	return
}

func (s Screen) depth(h int) (res [7]int) {
	res = [7]int{}
	for i := range res {
		m := uint8(1 << i)
		for j := h; j >= 0; j-- {
			if s[j]&m != 0 {
				res[i] = h - j
				break
			}
		}
	}
	return
}

func (s *Screen) draw(shape Shape) {
	for _, p := range shape.points {
		x := shape.p.X + p.X
		y := shape.p.Y + p.Y
		(*s)[y] |= 1 << x
	}
}

func (s *Shape) canMove(m image.Point, scr []uint8) bool {
	px := s.p.X + m.X
	py := s.p.Y + m.Y
	for _, i := range s.points {
		x := px + i.X
		y := py + i.Y
		switch {
		case x < 0, x >= 7, scr[y]&(1<<x) != 0:
			return false
		}
	}
	return true
}
