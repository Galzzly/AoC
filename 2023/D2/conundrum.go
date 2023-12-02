package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type game struct {
	id  int
	rgb []RGB
}

type RGB struct {
	red   int
	green int
	blue  int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	total, games := parseInput(lines)
	fmt.Println(total)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(total, games), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(games), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func part1(total int, games []game) int {
	var res int
	red := 12
	green := 13
	blue := 14
	res = total
nextgame:
	for _, game := range games {
		for _, rgb := range game.rgb {
			if rgb.red > red ||
				rgb.green > green ||
				rgb.blue > blue {
				res -= game.id
				continue nextgame
			}
		}
	}

	return res
}

func part2(games []game) int {
	var res int
	for _, game := range games {
		var red, green, blue int
		for _, rgb := range game.rgb {
			if rgb.red > red {
				red = rgb.red
			}
			if rgb.green > green {
				green = rgb.green
			}
			if rgb.blue > blue {
				blue = rgb.blue
			}
		}
		res += red * green * blue
	}
	return res
}

func parseInput(lines []string) (int, []game) {
	var result []game
	var total int
	for _, line := range lines {
		var rgb []RGB
		var id int
		var restofline string
		s := strings.Split(line, ": ")
		id = utils.Atoi(strings.Split(s[0], " ")[1])
		restofline = s[1]
		subset := strings.Split(restofline, "; ")
		for _, sub := range subset {
			var colours RGB
			s := strings.Split(sub, ", ")
			for _, s := range s {
				var num int
				var colour string
				fmt.Sscanf(s, "%d %s", &num, &colour)
				switch colour {
				case "red":
					colours.red += num
				case "green":
					colours.green += num
				case "blue":
					colours.blue += num
				}
				rgb = append(rgb, colours)
			}
		}
		total += id
		result = append(result, game{id, rgb})
	}
	return total, result
}
