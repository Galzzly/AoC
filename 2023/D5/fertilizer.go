package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Seeds []int
type Maps []Map
type Map []mapper
type mapper struct {
	dest, src, size int
}
type Ranges [][]int

var (
	seeds Seeds
	maps  Maps
)

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileDoubleLine(f)
	seeds.populateSeeds(lines[0])
	maps.populateMaps(lines[1:])
	t1 := time.Now()
	fmt.Println("Part 1:", part1(), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1() int {
	S := []int{}
	for _, seed := range seeds {
	nextmap:
		for _, tMap := range maps {
			for _, ttMap := range tMap {
				if ttMap.src <= seed && seed < ttMap.src+ttMap.size {
					seed = ttMap.dest + seed - ttMap.src
					continue nextmap
				}
			}
		}
		S = append(S, seed)
	}
	r, _ := utils.MinMax(S)
	return r
}

func part2() int {
	pairs := utils.ChunkSlice[int](seeds, 2)
	var R Ranges
	for _, pair := range pairs {
		R = append(R, []int{pair[0], pair[0] + pair[1] - 1})
	}
	r, _ := utils.MinMax(apply_range(R))
	return r
}

func apply_range(input Ranges) []int {
	A := []int{}
	for _, pair := range input {
		R := [][]int{pair}
		for _, tMap := range maps {
			AR := [][]int{}
			for _, ttMap := range tMap {
				src_end := ttMap.src + ttMap.size - 1
				NR := [][]int{}
				for _, v := range R {
					// Look for the range of values that are within the test range (inter)
					// Anything before and after can be passed on to the next test as is
					// Anything within the range can be modified correctly.
					// [start                                                           end]
					//                  [src_start        src_end]
					// [before         ][inter                   ][after                   ]
					// start is the lowest value of the lowest test value, and the start of
					// the input
					// End is the highest value of the largest test value and the end of the
					// input
					// Inter is the range between what is higher between the lowest test value
					// and start of the input
					before := []int{v[0], utils.Min(v[1], ttMap.src-1)}
					inter := []int{utils.Biggest(v[0], ttMap.src), utils.Min(v[1], src_end)}
					after := []int{utils.Biggest(v[0], src_end), v[1]}
					// If there are items prior to the lower test range, then pass them on
					if before[1] > before[0] {
						NR = append(NR, before)
					}
					// Transform the items within the range
					if inter[1] > inter[0] {
						diff := ttMap.dest - ttMap.src
						AR = append(AR, []int{diff + inter[0], diff + inter[1]})
					}
					// If there are items after the upper test range, then pass them on
					if after[1] > after[0] {
						NR = append(NR, after)
					}
				}
				R = NR
			}
			R = append(AR, R...)
		}
		for _, pair := range R {
			A = append(A, pair...)
		}
	}
	return A
}

func (s Seeds) populateSeeds(line string) {
	for _, seed := range strings.Fields(line)[1:] {
		seeds = append(seeds, utils.Atoi(seed))
	}
}

func (m Maps) populateMaps(input []string) {
	for _, lines := range input {
		tMap := []mapper{}
		for _, line := range strings.Split(lines, "\n")[1:] {
			s := strings.Fields(line)
			tMap = append(tMap, mapper{utils.Atoi(s[0]), utils.Atoi(s[1]), utils.Atoi(s[2])})
		}
		maps = append(maps, tMap)
	}
}
