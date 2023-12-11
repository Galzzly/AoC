package main

import (
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Universe [][]bool

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	universe := makeUniverse(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(universe, 2), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(universe, 1000000), time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func solve(universe Universe, n int) (res int) {
	galaxies := findGalaxies(universe)
	blankCol, blankRow := getEmpty(universe)
	var wg sync.WaitGroup
	ch := make(chan int, len(galaxies)-1)
	wg.Add(len(galaxies) - 1)
	for i, G1 := range galaxies[:len(galaxies)-1] {
		go func(r chan int, i int, G1 utils.Point) {
			defer wg.Done()
			result := 0
			for _, G2 := range galaxies[i+1:] {
				A := utils.Abs(G1.X - G2.X)
				B := utils.Abs(G1.Y - G2.Y)
				C := universe.checkBlank(G1.X, G2.X, blankCol)
				C += universe.checkBlank(G1.Y, G2.Y, blankRow)
				C *= n - 1
				result += A + B + C
			}
			r <- result
		}(ch, i, G1)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		res += r
	}
	return
}

func (u Universe) checkBlank(a, b int, c []int) (res int) {
	if a < b {
		return u.Check(a, b, c)
	}
	return u.Check(b, a, c)
}

func (u Universe) Check(a, b int, c []int) (res int) {
	for _, v := range c {
		if a < v && v < b {
			res++
		}
	}
	return
}

func findGalaxies(universe Universe) (res []utils.Point) {
	for Y := range universe {
		for X, gal := range universe[Y] {
			if gal {
				res = append(res, utils.Point{X: X, Y: Y})
			}
		}
	}
	return
}

func getEmpty(universe Universe) (col, row []int) {
	for R := range universe {
		if !slices.Contains(universe[R], true) {
			row = append(row, R)
		}
	}
	for C := 0; C < len(universe[0]); C++ {
		found := false
		for R := 0; R < len(universe); R++ {
			if universe[R][C] {
				found = true
				break
			}
		}
		if !found {
			col = append(col, C)
		}
	}
	return
}

func makeUniverse(lines []string) (universe Universe) {
	universe = make(Universe, len(lines))
	for R, line := range lines {
		universe[R] = make([]bool, len(line))
		for C, c := range line {
			if c == '.' {
				universe[R][C] = false
				continue
			}
			universe[R][C] = true
		}
	}
	return
}
