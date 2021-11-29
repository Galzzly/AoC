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
	mlines := utils.ReadFileDoubleLine(f)
	// convert to single line string
	lines := make([]string, len(mlines))
	for _, line := range mlines {
		single := strings.ReplaceAll(line, "\n", " ")
		lines = append(lines, single)
	}
	t1 := time.Now()
	t1res, valid := part1(lines)
	fmt.Println("Part 1:", t1res, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(valid), "Took:", time.Since(t2))
	fmt.Println("Total Time: ", time.Since(start))
}

func part1(lines []string) (result int, valid []string) {
	for i := 0; i < len(lines); i++ {
		s := strings.Split(lines[i], " ")
		if len(s) == 8 {
			result++
			valid = append(valid, lines[i])
			continue
		}
		if len(s) == 7 && !utils.FoundString(s, "cid") {
			result++
			valid = append(valid, lines[i])
			continue
		}
	}
	return
}

func part2(lines []string) (result int) {
	for i := 0; i < len(lines); i++ {
		if checkPassport(lines[i]) {
			result++
		}
	}
	return
}

func checkPassport(s string) bool {
	parts := strings.Split(s, " ")
	for _, part := range parts {
		field := strings.Split(part, ":")
		switch field[0] {
		case "byr":
			if len(field[1]) == 4 && 1920 <= utils.Atoi(field[1]) && utils.Atoi(field[1]) <= 2002 {
				continue
			}
			return false
		case "iyr":
			if len(field[1]) == 4 && 2010 <= utils.Atoi(field[1]) && utils.Atoi(field[1]) <= 2020 {
				continue
			}
			return false
		case "eyr":
			if len(field[1]) == 4 && 2020 <= utils.Atoi(field[1]) && utils.Atoi(field[1]) <= 2030 {
				continue
			}
			return false
		case "hgt":
			if (strings.HasSuffix(field[1], "cm") && utils.Atoi(field[1][:len(field[1])-2]) >= 150 && utils.Atoi(field[1][:len(field[1])-2]) <= 193) ||
				(strings.HasSuffix(field[1], "in") && utils.Atoi(field[1][:len(field[1])-2]) >= 59 && utils.Atoi(field[1][:len(field[1])-2]) <= 76) {
				continue
			}
			return false
		case "hcl":
			if strings.HasPrefix(field[1], "#") && len(field[1]) == 7 {
				continue
			}
			return false
		case "ecl":
			if field[1] == "amb" || field[1] == "blu" || field[1] == "brn" || field[1] == "gry" || field[1] == "grn" || field[1] == "hzl" || field[1] == "oth" {
				continue
			}
			return false
		case "pid":
			if len(field[1]) == 9 {
				continue
			}
			return false
		}
	}
	return true
}
