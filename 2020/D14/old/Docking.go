package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	p1, p2 := memory(lines)
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func memory(l []string) (int, int) {
	m1, m2 := map[int]int{}, map[int]int{}
	r1, r2 := 0, 0
	var mask string
	for _, s := range l {
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
			r2, m2[a] = r2+v-m2[a], v
		}
		v = apply(mask, v)
		r1, m1[a] = r1+v-m1[a], v
	}
	return r1, r2
}

func apply(mask string, value int) int {
	and, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 0)
	or, _ := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 0)
	return value&int(and) | int(or)
}
