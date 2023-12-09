package main

import (
	"fmt"
	"slices"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	nums := utils.ReadRowIntsByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(nums), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(nums), "Took:", time.Since(t2))
	fmt.Println("Total tiem:", time.Since(start))
}

func part1(numlist [][]int) int {
	var res int
	for _, nums := range numlist {
		res += extrapolate(nums)
	}
	return res
}

func part2(numlist [][]int) int {
	var res int
	for _, nums := range numlist {
		slices.Reverse(nums)
		res += extrapolate(nums)
	}
	return res
}

func extrapolate(nums []int) int {
	var res int
	extr := make([][]int, 0)
	extr = append(extr, nums)
	for {
		next := []int{}
		for i := 0; i < len(extr[len(extr)-1])-1; i++ {
			next = append(next, extr[len(extr)-1][i+1]-extr[len(extr)-1][i])
		}
		extr = append(extr, next)
		allzero := true
		for _, v := range next {
			if v != 0 {
				allzero = false
			}
		}
		if allzero {
			break
		}
	}

	for i := len(extr) - 2; i >= 0; i-- {
		extr[i] = append(extr[i], extr[i][len(extr[i])-1]+extr[i+1][len(extr[i+1])-1])
	}
	res = extr[0][len(extr[0])-1]
	return res
}
