package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Hotsprings []Hotspring
type Hotspring struct {
	spring      []rune
	arrangement []int
}
type Tested map[[3]int]int

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	hotsprings := getHotsprings(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(hotsprings, false), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(hotsprings, true), "Took:", time.Since(t2))

	fmt.Println("Total Time:", time.Since(start))
}

func solve(hotsprings Hotsprings, part2 bool) (res int) {
	var wg sync.WaitGroup
	wg.Add(len(hotsprings))
	arrangement := make(chan int, len(hotsprings))
	for _, hs := range hotsprings {
		if part2 {
			s := hs.spring
			a := hs.arrangement
			for i := 0; i < 4; i++ {
				hs.spring = append(hs.spring, '?')
				hs.spring = append(hs.spring, s...)
				hs.arrangement = append(hs.arrangement, a...)
			}
		}
		go hs.getArrangement(&wg, arrangement, 0, 0, 0)
	}

	go func() {
		wg.Wait()
		close(arrangement)
	}()

	for A := range arrangement {
		res += A
	}
	return
}

func (h Hotspring) getArrangement(wg *sync.WaitGroup, arr chan int, si, ai, curr int) {
	defer wg.Done()
	tested := Tested{}
	arr <- tested.testArrangement(h.spring, h.arrangement, 0, 0, 0)
}

func (t Tested) testArrangement(spring []rune, arr []int, si, ai, curr int) (res int) {
	k := [3]int{si, ai, curr}
	if val, ok := t[k]; ok {
		return val
	}
	if si == len(spring) {
		if (ai == len(arr) && curr == 0) || (ai == len(arr)-1 && arr[ai] == curr) {
			return 1
		}
		return 0
	}
	for _, c := range []rune{'.', '#'} {
		if spring[si] == c || spring[si] == '?' {
			if c == '.' && curr == 0 {
				res += t.testArrangement(spring, arr, si+1, ai, 0)
			} else if c == '.' && curr > 0 && ai < len(arr) && arr[ai] == curr {
				res += t.testArrangement(spring, arr, si+1, ai+1, 0)
			} else if c == '#' {
				res += t.testArrangement(spring, arr, si+1, ai, curr+1)
			}
		}
	}
	t[k] = res
	return res
}

func getHotsprings(lines []string) (hotsprings Hotsprings) {
	hotsprings = make(Hotsprings, 0, len(lines))
	for _, line := range lines {
		var hs, ar string
		fmt.Sscanf(line, "%s %s", &hs, &ar)
		s := strings.Split(ar, ",")
		arr := make([]int, len(s))
		for i, v := range s {
			arr[i] = utils.Atoi(v)
		}
		hotsprings = append(hotsprings, Hotspring{[]rune(hs), arr})
	}
	return
}
