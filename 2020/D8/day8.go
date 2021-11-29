package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type comp struct {
	instrs []instruction
	index  int
	acc    int
}

type instruction struct {
	op       string
	argument int
}

func main() {
	start := time.Now()
	f := os.Args[1]
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part1:", part1(lines), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part2:", part2(lines), "Took:", time.Since(t2))
	fmt.Println("Total time: ", time.Since(start))
}

func part1(lines []string) (res int) {
	var ran []int
	instr := 0
	for {
		ran = append(ran, instr)
		s := strings.Split(lines[instr], " ")
		switch s[0] {
		case "acc":
			incr, _ := strconv.Atoi(strings.TrimSpace(s[1]))
			res += incr
			instr++
		case "jmp":
			incr, _ := strconv.Atoi(strings.TrimSpace(s[1]))
			instr += incr
		case "nop":
			instr++
		}
		if utils.FoundInt(ran, instr) {
			return
		}
	}
}

func part2(lines []string) (res int) {
	comp := makeComp(lines)
	for i := range comp.instrs {
		newComp := makeComp(lines)
		switch newComp.instrs[i].op {
		case "jmp":
			newComp.instrs[i].op = "nop"
		case "nop":
			newComp.instrs[i].op = "jmp"
		case "acc":
			continue
		}
		if looper, res := isInfinate(newComp); !looper {
			return res
		}
	}

	return
}

func makeComp(lines []string) comp {
	var instr []instruction
	for _, line := range lines {
		s := strings.Split(line, " ")
		inst := s[0]
		arg, _ := strconv.Atoi(strings.TrimSpace(s[1]))
		instr = append(instr, instruction{inst, arg})
	}
	return comp{instr, 0, 0}
}

func isInfinate(c comp) (bool, int) {
	var ran []int
	for c.index < len(c.instrs) {
		ran = append(ran, c.index)
		switch c.instrs[c.index].op {
		case "acc":
			c.acc += c.instrs[c.index].argument
			c.index++
		case "jmp":
			c.index += c.instrs[c.index].argument
		case "nop":
			c.index++
		}
		if utils.FoundInt(ran, c.index) {
			return true, 0
		}
	}
	return false, c.acc
}
