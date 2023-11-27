package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/Galzzly/AoC/utils"
)

// for Delta and DV Up, Right, Down, Left
// var D = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
var D = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var DV = []int{3, 0, 1, 2}

var R, C int
var G []string
var instr string

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileDoubleLineNoTrim(f)
	// jungle, rect := makeMap(lines[0])
	// // fmt.Println("--", string(jungle[image.Point{0, 0}]), "--")
	// // instr := getInstr(lines[1])
	// // fmt.Println(len(instr))
	// fmt.Println("Part 1:", solve(jungle, rect, lines[1]))

	G = strings.Split(lines[0], "\n")
	instr = strings.TrimSpace(lines[1])
	R = len(G)
	C = len(G[0])
	for i := 0; i < R; i++ {
		for len(G[i]) < C {
			G[i] += " "
		}
	}
	fmt.Println(solve())
	fmt.Println("Total Time:", time.Since(start))
}

func solve() (res int) {
	r, c, d := 0, 0, 1
	for G[r][c] != '.' {
		c += 1
	}
	i := 0
	for i < len(instr) {
		n := 0
		for i < len(instr) && unicode.IsDigit(rune(instr[i])) {
			n = n*10 + int(rune(instr[i]))
			i += 1
		}
		for j := 0; j < n; j++ {
			rr := (r + D[d][0]) % R
			if rr < 0 {
				rr = R - 1
			}
			if rr == R {
				rr = 0
			}
			cc := (c + D[d][1]) % C
			if cc == C {
				cc = 0
			}
			if cc < 0 {
				cc = C - 1
			}
			if G[rr][cc] == ' ' {
				nr, nc, nd := getDest(r, c, d)
				if G[nr][nc] == '#' {
					break
				}
				r, c, d = nr, nc, nd
				continue
			} else if G[rr][cc] == '#' {
				break
			} else {
				r = rr
				c = cc
			}
		}
		if i == len(instr) {
			break
		}
		turn := instr[i]
		if turn == 'L' {
			d = (d + 3) % 4
		}
		if turn == 'R' {
			d = (d + 1) % 4
		}
		i += 1
	}
	fmt.Println(r+1, c+1, d)
	res = ((r+1)*1000 + (c+1)*4 + DV[d])
	return
}

func getDest(r, c, d int) (int, int, int) {
	r = (r + D[d][0]) % R
	if r < 0 {
		r = R - 1
	}
	if r == R {
		r = 0
	}
	c = (c + D[d][1]) % C
	if c < 0 {
		c = C - 1
	}
	if c == C {
		c = 0
	}
	for G[r][c] == ' ' {
		r = (r + D[d][0]) % R
		if r < 0 {
			r = R - 1
		}
		if r == R {
			r = 0
		}
		c = (c + D[d][1]) % C
		if c < 0 {
			c = C - 1
		}
		if c == C {
			c = 0
		}

	}
	return r, c, d
}

// func solve(jungle map[image.Point]rune, rect image.Rectangle, instr string) (res int) {
// 	// Get the starting point
// 	x := 0
// 	y := 0
// 	z := 1
// 	for jungle[image.Point{x, 0}] != '.' {
// 		x++
// 	}
// 	fmt.Println(x)

// 	i := 0
// 	// for k, v := range instr {
// 	for i < len(instr) {
// 		// If it's even, it's a number to move
// 		// if k%2 == 0 {
// 		// 	for n := 0; n < v.(int); n++ {
// 		n := 0
// 		for i < len(instr) && unicode.IsDigit(rune(instr[i])) {
// 			n = n*10 + utils.Atoi(string(instr[i]))
// 			i++
// 		}
// 		for j := 0; j < n; j++ {
// 			xx := (x + Delta[z][0]) % rect.Max.X
// 			yy := (y + Delta[z][1]) % rect.Max.Y
// 			if jungle[image.Point{xx, yy}] == ' ' {
// 				nx, ny := getDest(x, y, z, jungle, rect)
// 				fmt.Println(xx, yy, nx, ny)
// 				if jungle[image.Point{nx, ny}] == '#' {
// 					break
// 				}
// 				x, y = nx, ny
// 				continue
// 			} else if jungle[image.Point{xx, yy}] == '#' {
// 				break
// 			} else {
// 				x = xx
// 				y = yy
// 			}
// 		}
// 		// } else {
// 		if i == len(instr) {
// 			break
// 		}

// 		switch string(instr[i]) {
// 		// switch v.(string) {
// 		case "L":
// 			z = (z + 3) % 4
// 		case "R":
// 			z = (z + 1) % 4
// 		}
// 		i++
// 	}

// 	// }
// 	fmt.Println(y+1, x+1, DV[z])
// 	res = ((y+1)*1000 + (x+1)*4 + DV[z])
// 	return
// }

// func getDest(x, y, z int, jungle map[image.Point]rune, rect image.Rectangle) (int, int) {
// 	xx := (x + Delta[z][0]) % rect.Max.X
// 	yy := (y + Delta[z][1]) % rect.Max.Y
// 	for jungle[image.Point{xx, yy}] == ' ' {
// 		xx = (xx + Delta[z][0]) % rect.Max.X
// 		yy = (yy + Delta[z][1]) % rect.Max.Y
// 	}
// 	return xx, yy
// }

// // func getInstr(input string) []any {
// // 	instr := []any{}
// // 	i := 0
// // 	for i < len(input) {
// // 		if unicode.IsLetter(rune(input[i])) {
// // 			instr = append(instr, string(input[i]))
// // 			i++
// // 		} else {
// // 			n := 0
// // 			for i < len(input) && unicode.IsDigit(rune(input[i])) {
// // 				n = (n * 10) + utils.Atoi(string(input[i]))
// // 				// fmt.Println(n)
// // 				i++
// // 			}
// // 			instr = append(instr, n)
// // 		}
// // 	}
// // 	return instr
// // }

// func makeMap(input string) (map[image.Point]rune, image.Rectangle) {
// 	jungle := make(map[image.Point]rune)
// 	lines := strings.Split(input, "\n")
// 	for y, line := range lines {
// 		for x, c := range line {
// 			jungle[image.Point{x, y}] = c
// 		}
// 	}
// 	return jungle, image.Rect(0, 0, len(lines), len(lines[1]))
// }
