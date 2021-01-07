package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	// Map the rules
	rules := map[string][]int{}
	for _, v := range strings.Split(lines[0], "\n") {
		r := strings.Split(v, ": ")
		rules[r[0]] = make([]int, 4)
		fmt.Sscanf(r[1], "%d-%d or %d-%d", &rules[r[0]][0], &rules[r[0]][1], &rules[r[0]][2], &rules[r[0]][3])
	}
	myTx := getTx(lines[1])
	nearbyTx := getTx(lines[2])
	t1 := time.Now()
	p1, valid := part1(nearbyTx, rules)
	fmt.Printf("Part 1: %d (%s)\n", p1, time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(myTx[0], valid, rules), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}
func checkNum(n int, r []int) bool {
	if !(n >= r[0] && n <= r[1] || n >= r[2] && n <= r[3]) {
		return false
	}
	return true
}
func checkNums(n []int, r []int) bool {
	for _, v := range n {
		if !checkNum(v, r) {
			return false
		}
	}
	return true
}
func part2(myTx []int, valid [][]int, rules map[string][]int) int {
	ret := 1
	var tIdx [][]int
	for i := 0; i < len(valid[0]); i++ {
		var n []int
		for _, v := range valid {
			n = append(n, v[i])
		}
		tIdx = append(tIdx, n)
	}

	cMap := make(map[string][]int)
	for r, v := range rules {
		for x, y := range tIdx {
			if checkNums(y, v) {
				cMap[r] = append(cMap[r], x)
			}
		}
	}

	more := true
	for more {
		for r := range rules {
			c := cMap[r]
			if len(c) > 1 {
				continue
			} else if len(c) == 0 {
				panic("no C found")
			}

			correct := c[0]
			for k, oCs := range cMap {
				if len(oCs) == 1 || k == r {
					continue
				}
				for i, oC := range oCs {
					if oC == correct {
						cMap[k] = append(oCs[:i], oCs[i+1:]...)
						break
					}
				}
			}
		}
		more = false
		for r := range rules {
			if len(cMap[r]) > 1 {
				more = true
				break
			}
		}
	}

	for i, v := range cMap {
		if strings.HasPrefix(i, "departure") {
			target := v[0]
			ret *= myTx[target]
		}
	}
	return ret
}
func part1(nearbyTx [][]int, rules map[string][]int) (res int, valid [][]int) {
	res = 0
	//valid = make(map[int][]int)
validTx:
	for _, tx := range nearbyTx {
		for _, n := range tx {
			invalid := 0
			for _, v := range rules {
				if !checkNum(n, v) {
					invalid++
				}
			}
			if invalid == len(rules) {
				res += n
				continue validTx
			}
		}
		valid = append(valid, tx)
	}
	return
}
func getTx(l string) [][]int {
	ret := [][]int{}
	for _, s := range strings.Split(strings.TrimSpace(l), "\n")[1:] {
		nums := []int{}
		for _, v := range strings.Split(s, ",") {
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}
		ret = append(ret, nums)
	}
	return ret
}
