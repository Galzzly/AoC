package main

import (
	"fmt"
	"image"
	"time"
	"unicode"

	"github.com/Galzzly/AoC/utils"
)

var Delta = []int{-1, 0, 1}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	G, R := utils.MakeImagePointMapRect(lines)
	t1 := time.Now()
	nums, p1 := solve(G, R)
	fmt.Println("Part 1:", p1, "Took:", time.Since(t1))
	t2 := time.Now()
	p2 := 0
	for _, v := range nums {
		if len(v) == 2 {
			p2 += v[0] * v[1]
		}
	}
	fmt.Println("Part 2:", p2, "Took:", time.Since(t2))
	fmt.Println("Total time:", time.Since(start))
}

func solve(G map[image.Point]rune, R image.Rectangle) (nums map[image.Point][]int, r1 int) {
	Row := R.Max.Y
	Col := R.Max.X
	nums = make(map[image.Point][]int, 0)
	for Y := R.Min.Y; Y <= Row; Y++ {
		gears := map[image.Point]bool{}
		n := 0
		has_part := false
		for X := R.Min.X; X <= Col+1; X++ {
			if unicode.IsDigit(G[image.Point{X, Y}]) {
				n = n*10 + utils.Atoi(string(G[image.Point{X, Y}]))
				for _, YY := range Delta {
					for _, XX := range Delta {
						if (0 <= Y+YY && Y+YY <= Row) && (0 <= X+XX && X+XX <= Col) {
							p := image.Point{X + XX, Y + YY}
							c := G[p]
							if !unicode.IsDigit(c) && c != '.' {
								has_part = true
							}
							if c == '*' {
								gears[p] = true
							}
						}
					}
				}
			} else if n > 0 {
				if has_part {
					r1 += n
				}
				for gear := range gears {
					nums[gear] = append(nums[gear], n)
				}
				n = 0
				has_part = false
				gears = map[image.Point]bool{}
			}
		}
	}

	return
}

// type schematicmap map[image.Point]rune

var FullDelta = []image.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

// var AroundLeft = []image.Point{
// 	{-1, -1}, {0, -1},
// 	{-1, 0},
// 	{-1, 1}, {0, 1},
// }

// var AroundRight = []image.Point{
// 	{0, -1}, {1, -1},
// 	{1, 0},
// 	{0, 1}, {1, 1},
// }

// var TopBottom = []image.Point{
// 	{0, -1},
// 	{0, 1},
// }

// func main() {
// 	f := "input.txt"
// 	// f := "sample"
// 	schematic := utils.MakeImagePointMap(utils.ReadFileLineByLine(f))
// 	t1 := time.Now()
// 	fmt.Println("Part 1:", part1(schematic), "Took:", time.Since(t1))
// 	t2 := time.Now()
// 	fmt.Println("Part 2:", part2(schematic), "Took:", time.Since(t2))
// }

// func part1(grid schematicmap) (res int) {
// 	var wg sync.WaitGroup
// 	nums := make(chan int, len(grid))
// 	wg.Add(len(grid))
// 	go func() {
// 		for k, v := range grid {
// 			if !unicode.IsDigit(v) {
// 				wg.Done()
// 				continue
// 			}
// 			grid.Lookaround(k, v, nums, &wg)
// 		}
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(nums)
// 	}()

// 	for num := range nums {
// 		res += num
// 	}
// 	return
// }

// func part2(grid schematicmap) (res int) {
// 	var wg sync.WaitGroup
// 	nums := make(chan int, len(grid))
// 	wg.Add(len(grid))
// 	go func() {
// 		for k, v := range grid {
// 			if v != '*' {
// 				wg.Done()
// 				continue
// 			}
// 			grid.Dualgears(k, v, nums, &wg)
// 		}
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(nums)
// 	}()

// 	for num := range nums {
// 		res += num
// 	}
// 	return
// }

// func (s schematicmap) Dualgears(k image.Point, v rune, num chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	nums := []int{}
// 	np := k.Add(image.Point{-1, 0})
// 	if val, ok := s[np]; ok && unicode.IsDigit(val) {
// 		n := string(val)
// 		for p := np.Add(image.Point{-1, 0}); ; p = p.Add(image.Point{-1, 0}) {
// 			if val, ok := s[p]; ok && unicode.IsDigit(val) {
// 				n = string(val) + n
// 			} else {
// 				break
// 			}
// 		}
// 		nums = append(nums, utils.Atoi(n))
// 	}

// 	np = k.Add(image.Point{1, 0})
// 	if val, ok := s[np]; ok && unicode.IsDigit(val) {
// 		n := string(val)
// 		for p := np.Add(image.Point{1, 0}); ; p = p.Add(image.Point{1, 0}) {
// 			if val, ok := s[p]; ok && unicode.IsDigit(val) {
// 				n += string(val)
// 			} else {
// 				break
// 			}
// 		}
// 		nums = append(nums, utils.Atoi(n))
// 	}
// 	for _, y := range []int{-1, 1} {
// 		for _, x := range []int{-1, 0, 1} {
// 			np = k.Add(image.Point{x, y})
// 			if val, ok := s[np]; ok && unicode.IsDigit(val) {
// 				// Go left...
// 				start := np.X
// 				for i := start; i >= 0; i-- {
// 					if val, ok := s[image.Point{i, np.Y}]; ok && unicode.IsDigit(val) {
// 						start = i
// 						if i == 0 {
// 							break
// 						}
// 					} else {
// 						break
// 					}
// 				}

// 				// Then go right...
// 				p := image.Point{start, np.Y}
// 				n := string(s[p])
// 				for p = p.Add(image.Point{1, 0}); ; p = p.Add(image.Point{1, 0}) {
// 					if val, ok := s[p]; ok && unicode.IsDigit(val) {
// 						n += string(val)
// 					} else {
// 						break
// 					}
// 				}
// 				nums = append(nums, utils.Atoi(n))
// 				if start+len(n) > np.X {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	if len(nums) != 2 {
// 		num <- 0
// 		return
// 	}
// 	num <- nums[0] * nums[1]
// }

// func (s schematicmap) GetNum(k image.Point) int {
// 	n := string(s[k])
// 	// Go left
// 	for np, num := k.Add(image.Point{-1, 0}), true; num; np = np.Add(image.Point{-1, 0}) {
// 		if val, ok := s[np]; ok && unicode.IsDigit(val) {
// 			n = string(val) + n
// 			continue
// 		}
// 		num = false
// 	}
// 	// Go right
// 	for np, num := k.Add(image.Point{1, 0}), true; num; np = np.Add(image.Point{1, 0}) {
// 		if val, ok := s[np]; ok && unicode.IsDigit(val) {
// 			n += string(val)
// 			continue
// 		}
// 		num = false
// 	}
// 	return utils.Atoi(n)
// }

// func (s schematicmap) Lookaround(k image.Point, v rune, num chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	var ret int
// 	var symbol bool
// 	np := k.Add(image.Point{-1, 0})
// 	if val, ok := s[np]; !ok || !unicode.IsDigit(val) {
// 		var n string
// 		n += string(v)
// 		symbol = s.Search(k, AroundLeft)
// 		for np, num := k.Add(image.Point{1, 0}), true; num; np = np.Add(image.Point{1, 0}) {
// 			if val, ok := s[np]; ok && unicode.IsDigit(val) {
// 				n += string(val)
// 				if !symbol {
// 					symbol = s.Search(np, TopBottom)
// 				}
// 				continue
// 			}
// 			num = false
// 			if !symbol {
// 				symbol = s.Search(np.Add(image.Point{-1, 0}), AroundRight)
// 			}
// 		}
// 		if symbol {
// 			ret = utils.Atoi(n)
// 		}
// 	}
// 	num <- ret
// }

// func (s schematicmap) Search(k image.Point, delta []image.Point) bool {
// 	for _, p := range delta {
// 		np := k.Add(p)
// 		if val, ok := s[np]; ok && val != '.' && !unicode.IsDigit(val) {
// 			return true
// 		}
// 	}
// 	return false
// }
