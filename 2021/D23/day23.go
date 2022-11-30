package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/graph"
	"github.com/Galzzly/AoC/utils"
)

type Amphipods struct {
	x, y int
	kind string
}

type State struct {
	amphipods [16]Amphipods
	cost      int
	num       int
}

type Graph struct{}

var final1 = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########`

var final2 = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########`

func (s State) String() string {
	var rows [5][]string
	rows[0] = strings.Split("#...........#", "")
	rows[1] = strings.Split("###.#.#.#.###", "")
	rows[2] = strings.Split("  #.#.#.#.#", "")
	rows[3] = strings.Split("  #.#.#.#.#", "")
	rows[4] = strings.Split("  #.#.#.#.#", "")

	for _, a := range s.amphipods[:s.num] {
		rows[a.y][a.x+1] = a.kind
	}

	result := fmt.Sprintf(
		"#############\n%s\n%s\n%s\n",
		strings.Join(rows[0], ""),
		strings.Join(rows[1], ""),
		strings.Join(rows[2], ""),
		strings.Join(rows[3], ""),
		strings.Join(rows[4], ""),
		s.cost,
	)

	if s.num == 16 {
		result += fmt.Sprintf(
			"%s\n%s\n",
			strings.Join(rows[3], ""),
			strings.Join(rows[4], ""),
		)

	}

	return result + "  #########"
}

func main() {
	start := time.Now()
	lines := utils.ReadFileLineByLine("sample")
	t1 := time.Now()
	g := Graph{}
	t1cost, path := graph.Dijkstra[string](g, lines, final1)
	fmt.Printf("Part 1: %d in %v\n", part1(lines), time.Since(t1))
	fmt.Printf("Total time: %v\n", time.Since(start))
}
