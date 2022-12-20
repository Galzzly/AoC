package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
	"github.com/gammazero/deque"
)

type Blueprint struct {
	ID       int
	Ore      int
	Clay     int
	Obsidian [2]int
	Geode    [2]int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	blueprints := getBlueprints(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(blueprints), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve2(blueprints[:3]), "- Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func solve2(blueprints []Blueprint) (res int) {
	res = 1
	var wg sync.WaitGroup
	wg.Add(len(blueprints))
	resp := make(chan int, len(blueprints))
	go func() {
		for _, b := range blueprints {
			go func(b Blueprint) {
				defer wg.Done()
				resp <- solver(b.Ore, b.Clay, b.Obsidian[0], b.Obsidian[1], b.Geode[0], b.Geode[1], 32)
			}(b)
		}
		wg.Wait()
		close(resp)
	}()

	for r := range resp {
		res *= r
	}
	return
}

func solve(blueprints []Blueprint) (res int) {
	var wg sync.WaitGroup
	wg.Add(len(blueprints))
	resp := make(chan int, len(blueprints))
	go func() {
		for _, b := range blueprints {
			go func(b Blueprint) {
				defer wg.Done()
				s1 := solver(b.Ore, b.Clay, b.Obsidian[0], b.Obsidian[1], b.Geode[0], b.Geode[1], 24)
				resp <- b.ID * s1
			}(b)
		}
		wg.Wait()
		close(resp)
	}()

	for r := range resp {
		res += r
	}
	return
}

func solver(Co, Cc, Co1, Co2, Cg1, Cg2, T int) (res int) {
	S := [9]int{0, 0, 0, 0, 1, 0, 0, 0, T}
	var Q deque.Deque[[9]int]
	Q.PushBack(S)
	SEEN := map[[9]int]bool{}
	for Q.Len() != 0 {
		state := Q.PopFront()
		o, c, ob, g, r1, r2, r3, r4, t := state[0], state[1], state[2], state[3], state[4], state[5], state[6], state[7], state[8]
		res = utils.Biggest(res, g)
		if t == 0 {
			continue
		}

		_, Core := utils.MinMax([]int{Co, Cc, Co1, Cg1})
		if r1 >= Core {
			r1 = Core
		}
		if r2 >= Co2 {
			r2 = Co2
		}
		if r3 >= Cg2 {
			r3 = Cg2
		}
		if o >= t*Core-r1*(t-1) {
			o = t*Core - r1*(t-1)
		}
		if c >= t*Co2-r2*(t-1) {
			c = t*Co2 - r2*(t-1)
		}
		if ob >= t*Cg2-r3*(t-1) {
			ob = t*Cg2 - r3*(t-1)
		}
		state = [9]int{o, c, ob, g, r1, r2, r3, r4, t}
		if SEEN[state] {
			continue
		}
		SEEN[state] = true
		Q.PushBack([9]int{o + r1, c + r2, ob + r3, g + r4, r1, r2, r3, r4, t - 1})
		if o >= Co {
			Q.PushBack([9]int{o - Co + r1, c + r2, ob + r3, g + r4, r1 + 1, r2, r3, r4, t - 1})
		}
		if o >= Cc {
			Q.PushBack([9]int{o - Cc + r1, c + r2, ob + r3, g + r4, r1, r2 + 1, r3, r4, t - 1})
		}
		if o >= Co1 && c >= Co2 {
			Q.PushBack([9]int{o - Co1 + r1, c - Co2 + r2, ob + r3, g + r4, r1, r2, r3 + 1, r4, t - 1})
		}
		if o >= Cg1 && ob >= Cg2 {
			Q.PushBack([9]int{o - Cg1 + r1, c + r2, ob - Cg2 + r3, g + r4, r1, r2, r3, r4 + 1, t - 1})
		}
	}
	return
}

func getBlueprints(lines []string) (blueprints []Blueprint) {
	blueprints = make([]Blueprint, len(lines))
	for i, line := range lines {
		s := strings.Split(line, ". ")
		id := utils.Atoi(strings.TrimSuffix(strings.Fields(s[0])[1], ":"))
		ore := utils.Atoi(strings.Fields(s[0])[6])
		clay := utils.Atoi(strings.Fields(s[1])[4])
		var ob1, ob2, g1, g2 int
		fmt.Sscanf(s[2], "Each obsidian robot costs %d ore and %d clay", &ob1, &ob2)
		fmt.Sscanf(s[3], "Each geode robot costs %d ore and %d obsidian", &g1, &g2)
		blueprints[i] = Blueprint{id, ore, clay, [2]int{ob1, ob2}, [2]int{g1, g2}}
	}
	return
}
