package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Modules map[string]Module
type Module struct {
	Type  string
	State bool
	LastP map[string]bool
	Dest  []string
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	modules := getModules(lines)
	p1, p2 := solve(modules, "rx")
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
	fmt.Println("Total Time:", time.Since(start))
}

type Queue struct {
	M, From string
	Pulse   bool
}

func solve(modules Modules, end string) (r1, r2 int) {
	var watch []string
	var parent string
	from := map[string]map[string]bool{}
	for i, m := range modules {
		if slices.Contains(m.Dest, end) {
			parent = i
		}
		for _, d := range m.Dest {
			if modules[d].Type == "&" {
				if _, ok := from[d]; !ok {
					from[d] = map[string]bool{}
				}
				from[d][i] = false
			}
		}
	}
	for i, m := range modules {
		if slices.Contains(m.Dest, parent) {
			watch = append(watch, i)
		}
	}
	t := 0
	prev := map[string]int{}
	count := map[string]int{}
	on := []string{}
	nums := []int{}
	var low, high int
	for {
		t++
		cur := []Queue{{"broadcaster", "button", false}}
		for len(cur) > 0 {
			next := []Queue{}
			for _, Q := range cur {
				if !Q.Pulse {
					if _, ok := prev[Q.M]; ok && count[Q.M] == 2 && slices.Contains(watch, Q.M) {
						nums = append(nums, t-prev[Q.M])
					}
					prev[Q.M] = t
					count[Q.M] += 1
				}
				if len(nums) == len(watch) {
					return r1, utils.LCM(1, nums[0], nums[1:]...)
				}

				if !Q.Pulse {
					low++
				} else {
					high++
				}

				if _, ok := modules[Q.M]; !ok {
					continue
				}
				if Q.M == "broadcaster" {
					for _, nextM := range modules[Q.M].Dest {
						next = append(next, Queue{nextM, Q.M, Q.Pulse})
					}
				} else if modules[Q.M].Type == "%" {
					if Q.Pulse {
						continue
					}
					var newPulse bool
					if !slices.Contains[[]string](on, Q.M) {
						on = append(on, Q.M)
						newPulse = true
					} else {
						idx := slices.Index[[]string](on, Q.M)
						on = append(on[:idx], on[idx+1:]...)
						newPulse = false
					}
					for _, nextM := range modules[Q.M].Dest {
						next = append(next, Queue{nextM, Q.M, newPulse})
					}
				} else if modules[Q.M].Type == "&" {
					from[Q.M][Q.From] = Q.Pulse
					newPulse := allHigh(from[Q.M])
					for _, nextM := range modules[Q.M].Dest {
						next = append(next, Queue{nextM, Q.M, newPulse})
					}
				}
			}
			cur = next
		}
		if t == 1000 {
			r1 = low * high
		}
	}
}

func allHigh(P map[string]bool) bool {
	for _, v := range P {
		if !v {
			return !v
		}
	}
	return false
}

func getModules(lines []string) (modules Modules) {
	modules = make(Modules, len(lines))
	for _, line := range lines {
		s := strings.Split(line, " -> ")
		if s[0] == "broadcaster" {
			modules[s[0]] = Module{State: false, Dest: strings.Split(s[1], ", ")}
			continue
		}
		T := string(s[0][0])
		M := Module{
			Type:  T,
			State: utils.Ter(T == "&", true, false),
			LastP: map[string]bool{},
			Dest:  strings.Split(s[1], ", "),
		}
		modules[string(s[0][1:])] = M
	}
	for name, m := range modules {
		for _, d := range m.Dest {
			if modules[d].Type == "&" {
				state := modules[d]
				state.LastP[name] = false
				modules[d] = state
			}
		}
	}
	return
}
