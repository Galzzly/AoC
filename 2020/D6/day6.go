package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileDoubleLine(f)
	t1res, t2res := checkAnswers(lines)
	fmt.Println("Part 1:", t1res)
	fmt.Println("Part 2:", t2res)
	fmt.Println("Total time: ", time.Since(start))
}

func checkAnswers(lines []string) (p1, p2 int) {
	for _, line := range lines {
		q := map[string]int{}
		people := strings.Split(line, "\n")
		for _, p := range people {
			for _, ans := range strings.Split(p, "") {
				q[ans]++
			}
		}

		p1 += len(q)
		for _, v := range q {
			if v == len(people) {
				p2++
			}
		}
	}
	return p1, p2 - 1
}
