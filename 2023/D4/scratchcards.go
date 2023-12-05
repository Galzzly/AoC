package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type scratchcards map[int]int

type scratchcard struct {
	id      int
	winners int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	scratchcards := parseCards(lines)
	t1 := time.Now()
	fmt.Println("Part 1:", part1(scratchcards), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(scratchcards), "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func parseCards(lines []string) scratchcards {
	var wg sync.WaitGroup
	wg.Add(len(lines))
	cards := make(chan scratchcard, len(lines))
	scratchcards := make(map[int]int, len(lines))
	go func() {
		for _, line := range lines {
			parseCard(line, cards, &wg)
		}
	}()

	go func() {
		wg.Wait()
		close(cards)
	}()

	for card := range cards {
		scratchcards[card.id] = card.winners
	}
	return scratchcards
}

func part1(scratchcards scratchcards) (res int) {
	for _, winner := range scratchcards {
		switch winner {
		case 0:
			continue
		case 1:
			res++
		default:
			ret := 1
			for i := 2; i <= winner; i++ {
				ret *= 2
			}
			res += ret
		}
	}
	return
}

func part2(scratchcards scratchcards) (res int) {
	cardpile := make(map[int]int, len(scratchcards))
	for i := 1; i <= len(scratchcards); i++ {
		cardpile[i] = 1
	}
	for i := 1; i <= len(scratchcards); i++ {
		for j := 1; j <= scratchcards[i]; j++ {
			cardpile[i+j] += cardpile[i]
		}
		res += cardpile[i]
		// fmt.Println("card", i, ":", cardpile[i])
	}
	return
}

func parseCard(line string, card chan scratchcard, wg *sync.WaitGroup) {
	defer wg.Done()
	s := strings.Split(line, ": ")
	numbers := strings.Split(s[1], "|")
	winners := getnumbers(numbers[0])
	nums := getnumbers(numbers[1])
	// var found bool
	var ret scratchcard
	ret.id = utils.Atoi(strings.Fields(s[0])[1])
	ret.winners = 0
	for _, n := range nums {
		if checkNum(winners, n) {
			// if found {
			// 	ret *= 2
			// 	continue
			// }
			// ret = 1
			// found = true
			ret.winners++
		}
	}
	// fmt.Println(s[0], ret)
	card <- ret
}

func getnumbers(nums string) []int {
	s := strings.Fields(strings.TrimSpace(nums))
	ret := make([]int, 0)
	for _, n := range s {
		ret = append(ret, utils.Atoi(n))
	}
	return ret
}

func checkNum(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
