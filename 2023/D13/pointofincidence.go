package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Blocks []Block
type Block [][]rune

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileDoubleLine(f)
	blocks := makeBlocks(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(blocks, false), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(blocks, true), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func solve(blocks Blocks, part2 bool) int {
	var C, R int
	for _, block := range blocks {
		C += block.countcolumn(part2)
		R += block.countrow(part2)
	}
	return C + (100 * R)
}

func (b Block) countcolumn(part2 bool) (res int) {
	n := utils.Ter(part2, 1, 0)
	for c := 0; c < len(b[0])-1; c++ {
		howbad := 0
		for dc := range b[0] {
			L, R := c-dc, c+1+dc
			if 0 <= L && L < R && R < len(b[0]) {
				for r := 0; r < len(b); r++ {
					if b[r][L] != b[r][R] {
						howbad += 1
					}
				}
			}
		}
		if howbad == n {
			res += c + 1
		}
	}
	return
}

func (b Block) countrow(part2 bool) (res int) {
	n := utils.Ter(part2, 1, 0)
	for r := 0; r < len(b)-1; r++ {
		howbad := 0
		for dr := range b {
			U, D := r-dr, r+1+dr
			if 0 <= U && U < D && D < len(b) {
				for c := range b[0] {
					if b[U][c] != b[D][c] {
						howbad += 1
					}
				}
			}
		}
		if howbad == n {
			res += r + 1
		}
	}
	return
}

func makeBlocks(lines []string) (res Blocks) {
	res = make(Blocks, len(lines))
	for i, line := range lines {
		s := strings.Split(line, "\n")
		block := make(Block, len(s))
		for j, S := range s {
			block[j] = []rune(S)
		}
		res[i] = block
	}
	return
}
