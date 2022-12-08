package main

import (
	"fmt"
	"path"
	"strings"
	"time"
	"unicode"

	"github.com/Galzzly/AoC/utils"
)

type FS map[string]int

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)

	filesystem := getFs(lines)
	p1, p2 := filesystem.solve(100000, 30000000)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Took:", time.Since(start))
}

func (f FS) solve(m1, m2 int) (r1, r2 int) {
	r2 = f["/"]
	for _, s := range f {
		if s <= m1 {
			r1 += s
		}
		if s+70000000-f["/"] >= m2 && s < r2 {
			r2 = s
		}
	}
	return
}

func getFs(lines []string) (res FS) {
	res = make(FS)
	var p string
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd") {
			p = path.Join(p, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var size int
			fmt.Sscanf(line, "%d", &size)
			for d := p; d != "/"; d = path.Dir(d) {
				res[d] += size
			}
			res["/"] += size
		}
	}

	return
}
