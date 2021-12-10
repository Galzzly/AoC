package utils

import (
	"image"
	"io/ioutil"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Cube struct {
	X int
	Y int
	Z int
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileLineByLine(file string) []string {
	f, err := ioutil.ReadFile(file)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	return lines
}

func ReadFileDoubleLine(file string) []string {
	f, err := ioutil.ReadFile(file)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	return lines
}

func ReadIntsByLine(file string) (nums []int) {
	f, err := ioutil.ReadFile(file)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	nums = make([]int, 0, len(lines))
	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}
		n, err := strconv.Atoi(lines[l])
		Check(err)
		nums = append(nums, n)
	}
	return
}

func FileLineByComma(file string) []string {
	f, err := ioutil.ReadFile(file)
	Check(err)
	return strings.Split(strings.TrimSpace(string(f)), ",")
}

func FileIntsLineByComma(file string) []int {
	s := FileLineByComma(file)
	n := make([]int, 0, len(s))
	for _, a := range s {
		n = append(n, Atoi(a))
	}
	return n
}

func Reverse(s string) string {
	var ret strings.Builder
	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		ret.WriteRune(r[i])
	}
	return ret.String()
}

func FoundString(a []string, s string) bool {
	for _, v := range a {
		if strings.HasPrefix(v, s) {
			return true
		}
	}
	return false
}

func FoundInt(nums []int, n int) bool {
	for _, i := range nums {
		if i == n {
			return true
		}
	}
	return false
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func GetLargest(n []int) (result int) {
	for _, i := range n {
		if i > result {
			result = i
		}
	}
	return
}

func MakeImagePointMap(lines []string) (mapping map[image.Point]rune) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
		}
	}
	return
}

func MakeIntImagePointMap(lines []string) (mapping map[image.Point]int, rect image.Rectangle) {
	mapping = make(map[image.Point]int)
	for y, s := range lines {
		for x, r := range strings.Split(s, "") {
			mapping[image.Point{x, y}] = Atoi(r)
		}
	}
	rect = image.Rect(0, 0, len(lines[0]), len(lines))
	return
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func MinMax(nums []int) (min, max int) {
	min, max = int(^uint(0)>>1), 0
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}
