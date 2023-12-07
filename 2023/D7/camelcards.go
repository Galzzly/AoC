package main

import (
	"fmt"
	"reflect"
	"slices"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Hands []Hand
type Hand struct {
	hand []int
	bid  int
}
type Score struct {
	hand     []int
	strength int
}

var (
	handscores = []Score{
		{[]int{5}, 10}, {[]int{1, 4}, 9}, {[]int{2, 3}, 8}, {[]int{1, 1, 3}, 7},
		{[]int{1, 2, 2}, 6}, {[]int{1, 1, 1, 2}, 5}, {[]int{1, 1, 1, 1, 1}, 4},
	}

	cardscore = map[rune]int{
		'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
	}
)

func main() {
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", solve(lines, false), "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", solve(lines, true), "Took:", time.Since(t2))
}

func solve(lines []string, P2 bool) int {
	hands := getHands(lines, P2)
	slices.SortFunc(hands, func(a, b Hand) int {
		Ca := Counter(a, P2)
		Cb := Counter(b, P2)
		if Ca < Cb {
			return -1
		}
		if Ca > Cb {
			return 1
		}
		// If we get here, they have the same score, and so need to iterate through the cards
		for i := 0; i < len(a.hand); i++ {
			if a.hand[i] < b.hand[i] {
				return -1
			}
			if a.hand[i] > b.hand[i] {
				return 1
			}
		}
		return 0
	})
	res := 0
	for i, H := range hands {
		res += (i + 1) * H.bid
	}
	return res
}

func Counter(hand Hand, P2 bool) int {
	C := map[int]int{}
	for _, c := range hand.hand {
		C[c]++
	}
	if P2 {
		T := reflect.ValueOf(C).MapKeys()[0].Interface()
		for k := range C {
			if k != 1 {
				if C[k] > C[T.(int)] || T == 1 {
					T = k
				}
			}
		}
		if _, ok := C[1]; ok && T.(int) != 1 {
			C[T.(int)] += C[1]
			delete(C, 1)
		}
	}
	res := []int{}
	for _, v := range C {
		res = append(res, v)
	}
	slices.Sort(res)
	for _, S := range handscores {
		if reflect.DeepEqual(res, S.hand) {
			return S.strength
		}
	}
	return 0
}

func getHands(lines []string, P2 bool) Hands {
	var hands Hands
	for _, line := range lines {
		var H string
		var B int
		fmt.Sscanf(line, "%s %d", &H, &B)
		hand := []int{}
		for _, c := range H {
			n := cardscore[c]
			if c == 'J' && P2 {
				n = 1
			}
			hand = append(hand, n)
		}
		hands = append(hands, Hand{hand, B})
	}
	return hands
}
