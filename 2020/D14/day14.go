package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	p1, p2 := memory(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total Time:", time.Since(start))
}

func memory(lines []string) (p1, p2 int) {
	m1, m2 := map[int]int{}, map[int]int{}
	var mask string
	for _, s := range lines {
		if _, e := fmt.Sscanf(s, "mask = %s", &mask); e == nil {
			continue
		}
		var a, v int
		fmt.Sscanf(s, "mem[%d] = %d", &a, &v)
		for i, x := 0, strings.Count(mask, "X"); i < 1<<x; i++ {
			mask := strings.NewReplacer("X", "x", "0", "X").Replace(mask)
			for _, r := range fmt.Sprintf("%0*b", x, i) {
				mask = strings.Replace(mask, "x", string(r), 1)
			}

			a := apply(mask, a)
			p2, m2[a] = p2+v-m2[a], v
		}
		v = apply(mask, v)
		p1, m1[a] = p1+v-m1[a], v
	}
	return
}

func apply(mask string, value int) (res int) {
	and, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 0)
	or, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 0)
	return value&int(and) | int(or)
}
