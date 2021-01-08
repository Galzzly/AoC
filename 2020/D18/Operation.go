package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", part1(lines), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(lines), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func part2(l []string) int {
	res := 0
	re := regexp.MustCompile(`\([^\(\)]+\)`)
	for _, v := range l {
		res += r(v, re, func(s string) int {
			return r(s, regexp.MustCompile(`\d+ \+ \d+`), e)
		})
	}
	return res
}

func part1(l []string) int {
	res := 0
	re := regexp.MustCompile(`\([^\(\)]+\)`)
	for _, v := range l {
		res += r(v, re, e)
	}
	return res
}

func r(s string, re *regexp.Regexp, e func(string) int) int {
	for re.MatchString(s) {
		s = re.ReplaceAllStringFunc(s, func(s string) string {
			return strconv.Itoa(e(s))
		})
	}
	return e(s)
}

func e(s string) int {
	f := strings.Fields(strings.Trim(s, "()"))
	a, _ := strconv.Atoi(f[0])
	for i := 1; i < len(f); i += 2 {
		switch n, _ := strconv.Atoi(f[i+1]); f[i] {
		case "+":
			a += n
		case "*":
			a *= n
		}
	}
	return a
}
