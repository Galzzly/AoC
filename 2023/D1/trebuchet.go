package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/Galzzly/AoC/utils"
)

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

type Digits [2]digit

type digit struct {
	id  int
	num string
}

type returns struct {
	r1, r2 int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	c2start := time.Now()
	c2p1, c2p2 := calc2(lines)
	c2time := time.Since(c2start)
	fmt.Printf("calc2, P1: %d, P2: %d, Took: %s\n", c2p1, c2p2, c2time)
	fmt.Println("Total time:", time.Since(start))
}

func calc2(lines []string) (r1, r2 int) {
	var wg sync.WaitGroup
	ch := make(chan returns, len(lines))
	for _, line := range lines {
		wg.Add(1)
		go func(line string, wg *sync.WaitGroup, ch1 chan returns) {
			defer wg.Done()
			var res returns
			digits := findNums(line)
			res.r1 = utils.Atoi(digits[0].num + digits[1].num)

			words := findWords(line)
			if digits[0].id == -1 || (words[0].id != -1 && words[0].id < digits[0].id) {
				digits[0] = words[0]
			}
			if words[1].id != -1 && words[1].id > digits[1].id {
				digits[1] = words[1]
			}
			res.r2 = utils.Atoi(digits[0].num + digits[1].num)
			ch <- res
		}(line, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		r1 += res.r1
		r2 += res.r2
	}

	return
}

func findNums(line string) Digits {
	result := [2]digit{{-1, ""}, {-1, ""}}
	for i, c := range line {
		if unicode.IsDigit(c) {
			if result[0].id == -1 {
				result[0] = digit{i, string(c)}
				result[1] = digit{i, string(c)}
				continue
			}
			result[1] = digit{i, string(c)}
		}
	}
	return result
}

func findWords(line string) Digits {
	result := [2]digit{{-1, ""}, {-1, ""}}
	charMap := make(map[int]string, 0)
	var charIdx []int

	for word, num := range digitMap {
		first := strings.Index(line, word)
		last := strings.LastIndex(line, word)
		if first != -1 {
			charIdx = append(charIdx, first)
			charMap[first] = num
		}
		if last != -1 {
			charIdx = append(charIdx, last)
			charMap[last] = num
		}
	}

	if len(charIdx) > 0 {
		sort.Ints(charIdx)
		result[0] = digit{charIdx[0], charMap[charIdx[0]]}
		result[1] = digit{charIdx[len(charIdx)-1], charMap[charIdx[len(charIdx)-1]]}
	}
	return result
}
