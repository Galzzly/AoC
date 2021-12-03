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
	// fmt.Println(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total Time: ", time.Since(start))
}

func part1(lines []string) (res int64) {
	var bits = make(map[int][]int)
	for i := 0; i < len(lines[0]); i++ {
		bits[i] = make([]int, 2)
	}
	for _, line := range lines {
		for i, b := range line {
			switch b {
			case '0':
				bits[i] = []int{bits[i][0] + 1, bits[i][1]}
			case '1':
				bits[i] = []int{bits[i][0], bits[i][1] + 1}
			}
		}
	}

	var gamma, epsilon string
	for i := 0; i < len(bits); i++ {
		if bits[i][0] < bits[i][1] {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	res = g * e
	return
}

func part2(lines []string) (res int64) {
	var oxy, co2 string
	oxylines := lines
	co2lines := lines
	i := 0
	for len(oxylines) > 1 {
		oxyres := []string{}
		ocount := []int{0, 0}
		for _, line := range oxylines {
			switch line[i] {
			case '0':
				ocount[0]++
			case '1':
				ocount[1]++
			}
		}
		if ocount[0] < ocount[1] || ocount[0] == ocount[1] {
			oxy += "1"
		} else {
			oxy += "0"
		}
		for _, line := range oxylines {
			if strings.HasPrefix(line, oxy) {
				oxyres = append(oxyres, line)
			}
		}
		oxylines = oxyres
		i++
	}

	i = 0
	for len(co2lines) > 1 {
		co2res := []string{}
		ccount := []int{0, 0}
		for _, line := range co2lines {
			switch line[i] {
			case '0':
				ccount[0]++
			case '1':
				ccount[1]++
			}
		}

		if (ccount[0] < ccount[1]) || (ccount[0] == ccount[1]) {
			co2 += "0"
		} else {
			co2 += "1"
		}
		for _, line := range co2lines {
			if strings.HasPrefix(line, co2) {
				co2res = append(co2res, line)
			}
		}
		co2lines = co2res
		i++
	}

	c, _ := strconv.ParseInt(co2lines[0], 2, 64)
	o, _ := strconv.ParseInt(oxylines[0], 2, 64)
	res = c * o
	return
}
