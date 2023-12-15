package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type slot struct {
	label string
	val   int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.FileLineByComma(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(lines []string) (res int) {
	for _, line := range lines {
		res += getVal(line)
	}
	return
}

func part2(lines []string) (res int) {
	boxes := make([][]slot, 256)
	for _, line := range lines {
		if line[len(line)-1] == '-' {
			L := line[:len(line)-1]
			H := getVal(L)
			i := Contains(boxes[H], L)
			if i > -1 {
				boxes[H] = append(boxes[H][:i], boxes[H][i+1:]...)
			}
			continue
		}
		L := line[:len(line)-2]
		V := utils.Atoi(string(line[len(line)-1]))
		H := getVal(L)
		i := Contains(boxes[H], L)
		if i == -1 {
			boxes[H] = append(boxes[H], slot{L, V})
			continue
		}
		boxes[H][i].val = V
	}
	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		for j, L := range box {
			res += (i + 1) * (j + 1) * L.val
		}
	}
	return
}

func Contains(slots []slot, L string) int {
	for i, s := range slots {
		if s.label == L {
			return i
		}
	}
	return -1
}

func getVal(line string) (res int) {
	for _, c := range line {
		res += int(c)
		res *= 17
		res = res % 256
	}
	return
}
