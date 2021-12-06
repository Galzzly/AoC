package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	lines := utils.ReadFileDoubleLine(f)
	numbers := []int{}
	nums := strings.Split(lines[0], ",")
	for _, num := range nums {
		numbers = append(numbers, utils.Atoi(num))
	}
	boards := makeBoard(lines[1:])
	results := bingo(numbers, boards)
	fmt.Println("Part 1:", results[0])
	fmt.Println("Part 2:", results[len(results)-1])
}

func bingo(numbers []int, boards map[int][][]int) (res []int) {
	for i := range numbers {
		for board := range boards {
			if win(boards[board], numbers[:i]) {
				set := make(map[int]bool)
				for _, num := range numbers[:i] {
					set[num] = true
				}

				var total int
				for _, r := range boards[board] {
					for _, c := range r {
						if !set[c] {
							total += c
						}
					}
				}
				res = append(res, total*numbers[i-1])
				delete(boards, board)
			}
		}
	}
	return
}

func makeBoard(lines []string) (boards map[int][][]int) {
	boards = make(map[int][][]int)
	for i, line := range lines {
		rows := strings.Split(strings.TrimSpace(line), "\n")
		boards[i] = make([][]int, len(rows))
		for j, row := range rows {
			r := strings.Fields(strings.TrimSpace(row))
			boards[i][j] = make([]int, len(r))
			for k, num := range r {
				boards[i][j][k] = utils.Atoi(num)
			}
		}
	}
	return
}

func win(board [][]int, numbers []int) bool {
	played := make(map[int]bool)
	for _, n := range numbers {
		played[n] = true
	}

	for i := 0; i < len(board); i++ {
		var rMatch, cMatch int
		for j := 0; j < len(board[i]); j++ {
			if played[board[i][j]] {
				rMatch++
			}
			if played[board[j][i]] {
				cMatch++
			}
		}
		if rMatch == len(board) || cMatch == len(board) {
			return true
		}
	}
	return false
}
