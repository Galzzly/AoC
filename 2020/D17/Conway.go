package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type cube struct {
	w, x, y, z int
}

type pockDim struct {
	cubes map[cube]bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	// create the initial structure
	p := pockDim{}
	p.cubes = map[cube]bool{}
	for y, v := range lines {
		for x, s := range v {
			c := cube{w: 0, x: x, y: y, z: 0}
			p.cubes[c] = s == '#'
		}
	}
	p1, p2 := p, p
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", p1.runConway(6, false), time.Since(t1))

	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", p2.runConway(6, true), time.Since(t2))

	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func (p *pockDim) runConway(count int, hyper bool) int {
	// Go through the cycles needed
	for i := 0; i < count; i++ {
		nextDim := map[cube]bool{}

		// Get the invisible edge
		cTA := []cube{}
		for c := range p.cubes {
			for _, n := range c.getNeighbours(hyper) {
				_, ok := p.cubes[n]
				if !ok {
					cTA = append(cTA, n)
				}
			}
		}
		for _, c := range cTA {
			p.cubes[c] = false
		}

		// Get the next dimension state
		for c, a := range p.cubes {
			aN := p.getActiveN(c, hyper)
			nextDim[c] = aN == 3 || (a && aN == 2)
		}
		p.cubes = nextDim
	}

	// return the active cubes
	c := 0
	for _, a := range p.cubes {
		if a {
			c++
		}
	}
	return c
}

func (p *pockDim) getActiveN(c cube, hyper bool) int {
	res := 0
	for _, n := range c.getNeighbours(hyper) {
		if p.cubes[n] {
			res++
		}
	}
	return res
}

func (c cube) getNeighbours(hyper bool) []cube {
	res := []cube{}
	for _, v := range c.myNeighbours(hyper) {
		res = append(res, cube{c.w + v.w, c.x + v.x, c.y + v.y, c.z + v.z})
	}
	return res
}

func (c *cube) myNeighbours(hyper bool) []cube {
	if !hyper {
		return []cube{
			cube{w: 0, x: -1, y: -1, z: -1}, cube{w: 0, x: 0, y: -1, z: -1}, cube{w: 0, x: 1, y: -1, z: -1},
			cube{w: 0, x: -1, y: 0, z: -1}, cube{w: 0, x: 0, y: 0, z: -1}, cube{w: 0, x: 1, y: 0, z: -1},
			cube{w: 0, x: -1, y: 1, z: -1}, cube{w: 0, x: 0, y: 1, z: -1}, cube{w: 0, x: 1, y: 1, z: -1},
			cube{w: 0, x: -1, y: -1, z: 0}, cube{w: 0, x: 0, y: -1, z: 0}, cube{w: 0, x: 1, y: -1, z: 0},
			cube{w: 0, x: -1, y: 0, z: 0}, cube{w: 0, x: 1, y: 0, z: 0},
			cube{w: 0, x: -1, y: 1, z: 0}, cube{w: 0, x: 0, y: 1, z: 0}, cube{w: 0, x: 1, y: 1, z: 0},
			cube{w: 0, x: -1, y: -1, z: 1}, cube{w: 0, x: 0, y: -1, z: 1}, cube{w: 0, x: 1, y: -1, z: 1},
			cube{w: 0, x: -1, y: 0, z: 1}, cube{w: 0, x: 0, y: 0, z: 1}, cube{w: 0, x: 1, y: 0, z: 1},
			cube{w: 0, x: -1, y: 1, z: 1}, cube{w: 0, x: 0, y: 1, z: 1}, cube{w: 0, x: 1, y: 1, z: 1},
		}
	} else {
		return []cube{
			// w -1
			cube{x: -1, y: -1, z: -1, w: -1}, cube{x: 0, y: -1, z: -1, w: -1}, cube{x: 1, y: -1, z: -1, w: -1},
			cube{x: -1, y: 0, z: -1, w: -1}, cube{x: 0, y: 0, z: -1, w: -1}, cube{x: 1, y: 0, z: -1, w: -1},
			cube{x: -1, y: 1, z: -1, w: -1}, cube{x: 0, y: 1, z: -1, w: -1}, cube{x: 1, y: 1, z: -1, w: -1},
			cube{x: -1, y: -1, z: 0, w: -1}, cube{x: 0, y: -1, z: 0, w: -1}, cube{x: 1, y: -1, z: 0, w: -1},
			cube{x: -1, y: 0, z: 0, w: -1}, cube{x: 0, y: 0, z: 0, w: -1}, cube{x: 1, y: 0, z: 0, w: -1},
			cube{x: -1, y: 1, z: 0, w: -1}, cube{x: 0, y: 1, z: 0, w: -1}, cube{x: 1, y: 1, z: 0, w: -1},
			cube{x: -1, y: -1, z: 1, w: -1}, cube{x: 0, y: -1, z: 1, w: -1}, cube{x: 1, y: -1, z: 1, w: -1},
			cube{x: -1, y: 0, z: 1, w: -1}, cube{x: 0, y: 0, z: 1, w: -1}, cube{x: 1, y: 0, z: 1, w: -1},
			cube{x: -1, y: 1, z: 1, w: -1}, cube{x: 0, y: 1, z: 1, w: -1}, cube{x: 1, y: 1, z: 1, w: -1},

			cube{x: -1, y: -1, z: -1, w: 0}, cube{x: 0, y: -1, z: -1, w: 0}, cube{x: 1, y: -1, z: -1, w: 0},
			cube{x: -1, y: 0, z: -1, w: 0}, cube{x: 0, y: 0, z: -1, w: 0}, cube{x: 1, y: 0, z: -1, w: 0},
			cube{x: -1, y: 1, z: -1, w: 0}, cube{x: 0, y: 1, z: -1, w: 0}, cube{x: 1, y: 1, z: -1, w: 0},
			cube{x: -1, y: -1, z: 0, w: 0}, cube{x: 0, y: -1, z: 0, w: 0}, cube{x: 1, y: -1, z: 0, w: 0},
			cube{x: -1, y: 0, z: 0, w: 0}, cube{x: 1, y: 0, z: 0, w: 0},
			cube{x: -1, y: 1, z: 0, w: 0}, cube{x: 0, y: 1, z: 0, w: 0}, cube{x: 1, y: 1, z: 0, w: 0},
			cube{x: -1, y: -1, z: 1, w: 0}, cube{x: 0, y: -1, z: 1, w: 0}, cube{x: 1, y: -1, z: 1, w: 0},
			cube{x: -1, y: 0, z: 1, w: 0}, cube{x: 0, y: 0, z: 1, w: 0}, cube{x: 1, y: 0, z: 1, w: 0},
			cube{x: -1, y: 1, z: 1, w: 0}, cube{x: 0, y: 1, z: 1, w: 0}, cube{x: 1, y: 1, z: 1, w: 0},

			cube{x: -1, y: -1, z: -1, w: 1}, cube{x: 0, y: -1, z: -1, w: 1}, cube{x: 1, y: -1, z: -1, w: 1},
			cube{x: -1, y: 0, z: -1, w: 1}, cube{x: 0, y: 0, z: -1, w: 1}, cube{x: 1, y: 0, z: -1, w: 1},
			cube{x: -1, y: 1, z: -1, w: 1}, cube{x: 0, y: 1, z: -1, w: 1}, cube{x: 1, y: 1, z: -1, w: 1},
			cube{x: -1, y: -1, z: 0, w: 1}, cube{x: 0, y: -1, z: 0, w: 1}, cube{x: 1, y: -1, z: 0, w: 1},
			cube{x: -1, y: 0, z: 0, w: 1}, cube{x: 0, y: 0, z: 0, w: 1}, cube{x: 1, y: 0, z: 0, w: 1},
			cube{x: -1, y: 1, z: 0, w: 1}, cube{x: 0, y: 1, z: 0, w: 1}, cube{x: 1, y: 1, z: 0, w: 1},
			cube{x: -1, y: -1, z: 1, w: 1}, cube{x: 0, y: -1, z: 1, w: 1}, cube{x: 1, y: -1, z: 1, w: 1},
			cube{x: -1, y: 0, z: 1, w: 1}, cube{x: 0, y: 0, z: 1, w: 1}, cube{x: 1, y: 0, z: 1, w: 1},
			cube{x: -1, y: 1, z: 1, w: 1}, cube{x: 0, y: 1, z: 1, w: 1}, cube{x: 1, y: 1, z: 1, w: 1},
		}
	}

}
