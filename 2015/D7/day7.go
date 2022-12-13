package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Instr map[string]string

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	instr := Instr{}
	p1wires, p2wires := map[string]uint16{}, map[string]uint16{}
	for _, line := range lines {
		s1 := strings.Split(line, " -> ")
		s2 := strings.Split(s1[0], " ")
		if v, err := strconv.Atoi(s2[0]); err == nil && len(s2) == 1 {
			p1wires[s1[1]] = uint16(v)
			p2wires[s1[1]] = uint16(v)
		}
		instr[s1[1]] = s1[0]
	}
	p1 := instr.solve(p1wires)
	p2wires["b"] = p1
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", instr.solve(p2wires))
	fmt.Println("Took:", time.Since(start))
}

func (i Instr) solve(wires map[string]uint16) uint16 {
	for len(wires) < len(i) {
		for k, v := range i {
			if _, ok := wires[k]; ok {
				continue
			}
			s := strings.Split(v, " ")
			switch len(s) {
			case 1:
				if _, ok := wires[s[0]]; ok {
					wires[k] = wires[s[0]]
				}
			case 2:
				if _, ok := wires[s[1]]; ok {
					wires[k] = ^wires[s[1]]
				}
			case 3:
				if _, ok := wires[s[0]]; ok {
					switch s[1] {
					case "AND":
						if _, ok := wires[s[2]]; ok {
							wires[k] = wires[s[0]] & wires[s[2]]
						}
					case "OR":
						if _, ok := wires[s[2]]; ok {
							wires[k] = wires[s[0]] | wires[s[2]]
						}
					case "LSHIFT":
						wires[k] = wires[s[0]] << utils.Atoi(s[2])
					case "RSHIFT":
						wires[k] = wires[s[0]] >> utils.Atoi(s[2])
					}
				}
				if s[0] == "1" {
					switch s[1] {
					case "AND":
						if _, ok := wires[s[2]]; ok {
							wires[k] = uint16(utils.Atoi(s[0])) & wires[s[2]]
						}
					case "OR":
						if _, ok := wires[s[2]]; ok {
							wires[k] = uint16(utils.Atoi(s[0])) | wires[s[2]]
						}
					}
				}
			}
		}
	}
	return wires["a"]
}
