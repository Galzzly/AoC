package main

import (
	"fmt"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	nums := utils.ReadRowIntsByLine(f)
	solvetime := time.Now()
	p1, p2 := solve(nums)
	fmt.Printf("Part 1: %d\nPart 2: %d\nTime Taken: %v\n", p1, p2, time.Since(solvetime))
	fmt.Println("Total tiem:", time.Since(start))
}

func solve(numlist [][]int) (res, res2 int) {
	for _, nums := range numlist {
		r1, r2 := extrapolate(nums)
		res += r1
		res2 += r2
	}
	return
}

func extrapolate(nums []int) (res, res2 int) {
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
		res += extr[i][len(extr[i])-1]
		res2 = extr[i][0] - res2
	}
	return
}
