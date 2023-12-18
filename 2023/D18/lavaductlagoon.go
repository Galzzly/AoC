package main

import (
	"fmt"
	"image"
	"strconv"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Instr struct {
	Dir  string
	Dist int
}

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	for i, b := range []bool{false, true} {
		timer := time.Now()
		fmt.Printf("Part %d: %d Took: %v\n", i+1, solve(lines, b), time.Since(timer))
	}
	fmt.Println("Total Time:", time.Since(start))
}

type Points struct {
	a, b image.Point
}

func solve(lines []string, part2 bool) (res int) {
	var minX, maxX, minY, maxY int
	P := []Points{}
	for _, line := range lines {
		var D, H string
		var N int
		fmt.Sscanf(line, "%s %d (#%s)", &D, &N, &H)
		if part2 {
			switch H[5] {
			case '0':
				D = "R"
			case '1':
				D = "D"
			case '2':
				D = "L"
			case '3':
				D = "U"
			}
			val, _ := strconv.ParseInt(H[:5], 16, 64)
			N = int(val)
		}
		switch D {
		case "R":
			maxX = minX + N
		case "L":
			maxX = minX - N
		case "D":
			maxY = minY + N
		case "U":
			maxY = minY - N
		}
		P = append(P, Points{image.Point{minX, minY}, image.Point{maxX, maxY}})
		minX, minY = maxX, maxY
	}
	for i := range P {
		res += (P[i].a.Y + P[i].b.Y) * (P[i].a.X - P[i].b.X)
	}
	for _, p := range P {
		if p.a.X == p.b.X {
			res += utils.Abs(p.a.Y - p.b.Y)
			continue
		}
		res += utils.Abs(p.a.X - p.b.X)
	}
	res = res/2 + 1
	return
}

// func solve(lines []string, part2 bool) (res int) {
// 	var instr []Instr
// 	var values = map[string]int{}
// 	P := image.Point{0, 0}

// 	for _, line := range lines {
// 		var D, H string
// 		var N int

// 		fmt.Sscanf(line, "%s %d (#%s)", &D, &N, &H)
// 		if part2 {
// 			var dir string
// 			switch H[5] {
// 			case '0':
// 				dir = "R"
// 			case '1':
// 				dir = "D"
// 			case '2':
// 				dir = "L"
// 			case '3':
// 				dir = "U"
// 			}
// 			val, _ := strconv.ParseInt(H[:5], 16, 64)
// 			values[dir] = int(val)
// 			instr = append(instr, Instr{dir, int(val)})
// 			fmt.Println(dir, int(val), int(val)/N)
// 			continue
// 		}
// 		values[D] += N
// 		instr = append(instr, Instr{D, N})
// 	}

// 	maxY := (values["U"] + values["D"]) * 2
// 	maxX := (values["L"] + values["R"]) * 2
// 	// R := image.Rect(0, 0, maxX, maxY)

// 	if values["U"] > values["D"] {
// 		P.Y = (maxY / 2) + (values["U"] - values["D"])
// 	} else if values["U"] < values["D"] {
// 		P.Y = (maxY / 2) - (values["D"] - values["U"])
// 	} else {
// 		P.Y = 0
// 	}
// 	if values["L"] > values["R"] {
// 		P.X = (maxX / 2) + (values["L"] - values["R"])
// 	} else if values["L"] < values["R"] {
// 		P.X = (maxX / 2) - (values["R"] - values["L"])
// 	} else {
// 		P.X = 0
// 	}
// 	var edge = []image.Point{P}
// 	min, max := image.Point{maxX, maxY}, image.Point{0, 0}
// 	for _, ins := range instr {
// 		switch ins.Dir {
// 		case "U":
// 			edge = append(edge, PointToAdd(P, image.Point{0, -1}, ins.Dist)...)
// 		case "R":
// 			edge = append(edge, PointToAdd(P, image.Point{1, 0}, ins.Dist)...)
// 		case "L":
// 			edge = append(edge, PointToAdd(P, image.Point{-1, 0}, ins.Dist)...)
// 		case "D":
// 			edge = append(edge, PointToAdd(P, image.Point{0, 1}, ins.Dist)...)
// 		}
// 		res += ins.Dist
// 		P = edge[len(edge)-1]
// 		X, Y := P.X, P.Y
// 		min.X = utils.Ter(X < min.X, X, min.X)
// 		max.X = utils.Ter(X > max.X, X, max.X)
// 		min.Y = utils.Ter(Y < min.Y, Y, min.Y)
// 		max.Y = utils.Ter(Y > max.Y, Y, max.Y)
// 	}
// 	fmt.Println("EDGE IS DONE")
// 	pf := utils.NewPolyfence(edge)
// 	var wg sync.WaitGroup
// 	ch := make(chan int)
// 	for Y := min.Y; Y < max.Y; Y++ {
// 		for X := min.X; X < max.X; X++ {
// 			p := image.Point{X, Y}
// 			wg.Add(1)
// 			go func(edge []image.Point, p image.Point, pf *utils.Polyfence, wg *sync.WaitGroup, ch chan int) {
// 				defer wg.Done()
// 				if !slices.Contains(edge, p) && pf.Inside(p) {
// 					ch <- 1
// 					return
// 				}
// 				ch <- 0
// 			}(edge, p, pf, &wg, ch)

// 		}
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for n := range ch {
// 		res += n
// 	}
// 	return
// }

// func PointToAdd(P image.Point, M image.Point, D int) []image.Point {
// 	res := []image.Point{}
// 	for i, np := 0, P.Add(M); i < D; i, np = i+1, np.Add(M) {
// 		res = append(res, np)
// 	}
// 	return res
// }
