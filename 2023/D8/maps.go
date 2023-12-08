package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Maps map[rune]map[string]string

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileDoubleLine(f)
	instr := []rune(lines[0])
	maps := buildMaps(lines[1])
	t1 := time.Now()
	fmt.Println("Part 1:", solve(instr, maps, false), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 1:", solve(instr, maps, true), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func solve(instr []rune, maps Maps, part2 bool) int {
	P := []string{}
	suffix := utils.Ter(part2, "A", "AAA")
	for p := range maps['L'] {
		if strings.HasSuffix(p, suffix) {
			P = append(P, p)
		}
	}
	I := []int{}
	i := 0
	for {
		NP := []string{}
		for _, p := range P {
			p = maps[instr[i%len(instr)]][p]
			if strings.HasSuffix(p, "Z") {
				I = append(I, i+1)
				if len(I) == len(P) {
					return utils.LCM(1, I[0], I[1:]...)
				}
			}
			NP = append(NP, p)
		}
		P = NP
		i++
	}
}

func buildMaps(line string) Maps {
	lines := strings.Split(line, "\n")
	maps := make(map[rune]map[string]string, 2)
	maps['L'] = make(map[string]string, len(lines))
	maps['R'] = make(map[string]string, len(lines))
	for _, line := range lines {
		s := strings.Split(line, " = ")
		I := s[0]
		lr := strings.Split(s[1], ", ")
		L := lr[0][1:]
		R := lr[1][:3]
		maps['L'][I] = L
		maps['R'][I] = R
	}
	return maps
}
