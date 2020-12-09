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

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s : ", elapsed)
}

func D2P1(l []string) int {
	defer timeTrack(time.Now())
	count := 0
	for i := 0; i < len(l); i++ {
		vals := strings.Fields(l[i])
		n := strings.Split(vals[0], "-")
		mm := make([]int, 0, len(n))
		for _, a := range n {
			m, err := strconv.Atoi(a)
			check(err)
			mm = append(mm, m)
		}
		char := vals[1][:len(vals[1])-1]
		res := strings.Count(vals[2], char)
		if (res >= mm[0]) && (res <= mm[1]) {
			count++
		}
	}
	return count
}

func D2P2(l []string) int {
	defer timeTrack(time.Now())
	count := 0
	for i := 0; i < len(l); i++ {
		vals := strings.Fields(l[i])
		n := strings.Split(vals[0], "-")
		p := make([]int, 0, len(n))
		for _, a := range n {
			m, err := strconv.Atoi(a)
			check(err)
			p = append(p, m)
		}
		char := vals[1][:len(vals[1])-1]
		pass := string(vals[2])
		r := make([]string, 0, len(p))
		for _, a := range p {
			c := string(pass[a-1])
			r = append(r, c)
		}
		if (r[0] == char) && (r[1] == char) {
			continue
		} else if (r[0] == char) || (r[1] == char) {
			count++
		}
	}
	return count
}

func main() {
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	fmt.Printf("Part 1: %d\n", D2P1(lines))
	fmt.Printf("Part 2: %d\n", D2P2(lines))
}
