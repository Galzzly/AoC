package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func part1(lines []string) int {
	races := getRaces(lines)
	ways := make([]int, len(races))
	for i, r := range races {
		ways[i] = getWays(r)
	}
	return utils.MultiplyArray(ways)
}

func part2(lines []string) int {
	var race Race
	race.Time = utils.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	race.Distance = utils.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))
	return getMax(race) - getMin(race) + 1
}

func getWays(race Race) int {
	max := getMax(race)
	min := getMin(race)
	return max - min + 1
}

func getMin(race Race) int {
	for i := 0; i < race.Time; i++ {
		runtime := race.Time - i
		dist := i * runtime
		if dist > race.Distance {
			return i
		}
	}
	return race.Time
}

func getMax(race Race) int {
	for i := race.Time; i > 0; i-- {
		runtime := race.Time - i
		dist := i * runtime
		if dist > race.Distance {
			return i
		}
	}
	return 0
}

func getRaces(lines []string) []Race {
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	races := make([]Race, len(times))
	for i := range races {
		races[i] = Race{utils.Atoi(times[i]), utils.Atoi(distances[i])}
	}
	return races
}
