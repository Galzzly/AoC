package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
	"github.com/RyanCarrier/dijkstra"
)

type Valvematrix map[string]map[string]int

func main() {
	start := time.Now()
	f := "input.txt"
	// f = "test"
	lines := utils.ReadFileLineByLine(f)
	matrix, pressure, tunnels := getValves(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(matrix, pressure, 0, 0, 0, "AA", tunnels, 30), "- Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(matrix, pressure, tunnels), "- Took:", time.Since(t2))
	fmt.Println(time.Since(start))
}

func part2(matrix Valvematrix, pressures map[string]int, tunnels []string) (res int) {
	res = 0
	mu := sync.Mutex{}
	for i := 1; i < len(tunnels)/2; i++ {
		for v := range utils.Combo(tunnels, i) {
			go func(v []string) {
				m1 := solve(matrix, pressures, 0, 0, 0, "AA", v, 26)
				m2 := solve(matrix, pressures, 0, 0, 0, "AA", alt(tunnels, v), 26)
				mu.Lock()
				if m1+m2 > res {
					res = m1 + m2
				}
				mu.Unlock()
			}(v)
		}
	}
	return
}

func alt(all []string, part []string) (res []string) {
	res = []string{}
next:
	for _, v1 := range all {
		for _, v2 := range part {
			if v1 == v2 {
				continue next
			}
		}
		res = append(res, v1)
	}
	return
}

func solve(matrix Valvematrix, pressures map[string]int, curTime, curPress, curFlow int, curTun string, tunnels []string, lim int) (res int) {
	nScore := curPress + (lim-curTime)*curFlow
	res = nScore

	for _, v := range tunnels {
		distanceAndOpen := matrix[curTun][v] + 1
		if curTime+distanceAndOpen < lim {
			newTime := curTime + distanceAndOpen
			newPressure := curPress + distanceAndOpen*curFlow
			newFlow := curFlow + pressures[v]
			possibleScore := solve(matrix, pressures, newTime, newPressure, newFlow, v, removeTunnel(tunnels, v), lim)
			if possibleScore > res {
				res = possibleScore
			}
		}
	}
	return
}

func removeTunnel(t []string, v string) (res []string) {
	res = []string{}
	for _, i := range t {
		if i != v {
			res = append(res, i)
		}
	}
	return
}

func getValves(lines []string) (matrix Valvematrix, pressure map[string]int, tunnels []string) {
	pressure = make(map[string]int, len(lines))
	mappingInt := make(map[string]int, len(lines))
	reachability := map[string][]string{}
	count := 0
	for _, line := range lines {
		s := strings.Split(line, "; ")
		var valve string
		var rate int
		fmt.Sscanf(s[0], "Valve %s has flow rate=%d", &valve, &rate)
		lead := strings.TrimPrefix(strings.TrimPrefix(strings.Split(s[1], "valve")[1], "s"), " ")
		c := strings.Split(lead, ", ")
		reachability[valve] = []string{}
		reachability[valve] = append(reachability[valve], c...)
		pressure[valve] = rate
		mappingInt[valve] = count
		count++
	}

	tunnels = []string{}
	for k, v := range pressure {
		if v != 0 {
			tunnels = append(tunnels, k)
		}
	}
	matrix = calcReachabilityMatrix(reachability, mappingInt)
	return
}

func calcReachabilityMatrix(reachability map[string][]string, mappingInt map[string]int) map[string]map[string]int {
	graph := dijkstra.NewGraph()

	for k := range reachability {
		graph.AddVertex(mappingInt[k])
	}

	for k, v := range reachability {
		for _, l := range v {
			graph.AddArc(mappingInt[k], mappingInt[l], 1)
		}
	}

	matrix := map[string]map[string]int{}
	for k1, v1 := range mappingInt {
		matrix[k1] = map[string]int{}
		for k2, v2 := range mappingInt {
			best, _ := graph.Shortest(v1, v2)
			matrix[k1][k2] = int(best.Distance)
		}
	}

	return matrix
}
