package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/Galzzly/AoC/2020/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	t1res, seats := part1(lines)
	fmt.Println("Part 1:", t1res, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(seats), "Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
}

func part1(lines []string) (int, []int) {
	seats := make([]int, len(lines))
	for _, line := range lines {
		row := getRow(line[0:7])
		col := getCol(line[7:])
		seats = append(seats, row+col)
	}

	sort.Ints(seats)

	return seats[len(seats)-1], seats
}

func part2(seats []int) (myseat int) {
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] != 1 {
			myseat = seats[i] + 1
		}
	}

	return
}

func getRow(s string) int {
	var row []int
	for i := 0; i < 128; i++ {
		row = append(row, i)
	}
	for _, c := range s {
		mid := len(row) / 2
		switch c {
		case 'F':
			row = row[:mid]
		case 'B':
			row = row[mid:]
		}
	}
	return row[0] * 8
}

func getCol(s string) int {
	var col []int
	for i := 0; i < 8; i++ {
		col = append(col, i)
	}
	for _, c := range s {
		mid := len(col) / 2
		switch c {
		case 'L':
			col = col[:mid]
		case 'R':
			col = col[mid:]
		}
	}
	return col[0]
}
