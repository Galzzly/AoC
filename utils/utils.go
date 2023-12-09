package utils

import (
	"image"
	"os"
	"sort"
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
	f, err := os.ReadFile(file)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	return lines
}

func ReadFileSingleLine(file string) (line string) {
	f, err := os.ReadFile(file)
	Check(err)
	line = strings.Split(strings.TrimSpace(string(f)), "\n")[0]
	return line
}

func ReadFileDoubleLine(file string) []string {
	f, err := os.ReadFile(file)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	return lines
}

func ReadFileDoubleLineNoTrim(file string) []string {
	f, err := os.ReadFile(file)
	Check(err)
	lines := strings.Split(string(f), "\n\n")
	return lines
}

func ReadIntsByLine(file string) (nums []int) {
	f, err := os.ReadFile(file)
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
	f, err := os.ReadFile(file)
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

func ReadRowIntsByLine(file string) (nums [][]int) {
	f, err := os.ReadFile(file)
	Check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	nums = make([][]int, len(lines))
	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}
		s := strings.Split(lines[l], " ")
		nums[l] = make([]int, 0, len(s))
		for i := range s {
			nums[l] = append(nums[l], Atoi(s[i]))
		}
	}
	return
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

func GetLargest(n []int) (key, result int) {
	for k, v := range n {
		if v > result {
			result = v
			key = k
		}
	}
	return
}

func MakeImagePointSquareBool(max int) (mapping map[image.Point]bool) {
	mapping = make(map[image.Point]bool)
	for x := 0; x < max; x++ {
		for y := 0; y < max; y++ {
			mapping[image.Point{x, y}] = false
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
func MakeImagePointMapRect(lines []string) (mapping map[image.Point]rune, rect image.Rectangle) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
		}
	}
	rect = image.Rect(0, 0, len(lines[0])-1, len(lines)-1)
	return
}

func MakeIntImagePointMap(lines []string) (mapping map[image.Point]int, rect image.Rectangle) {
	mapping = MakeIntImagePoint(lines)
	rect = image.Rect(0, 0, len(lines[0])-1, len(lines)-1)
	return
}

func MakeIntImagePoint(lines []string) (mapping map[image.Point]int) {
	mapping = make(map[image.Point]int)
	for y, s := range lines {
		for x, r := range strings.Split(s, "") {
			mapping[image.Point{x, y}] = Atoi(r)
		}
	}
	return
}

func Adj(p, d image.Point) image.Point {
	return p.Add(d)
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

func SortUniqInts(s []int) []int {
	sort.Ints(s)
	j := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}
		s[j] = s[i]
		j++
	}
	return s[:j]
}

func SumArray(n []int) (res int) {
	for _, v := range n {
		res += v
	}
	return
}

func MultiplyArray(n []int) (res int) {
	res = 1
	for _, v := range n {
		res *= v
	}
	return
}

func Biggest(a, b int) (res int) {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) (res int) {
	if a < b {
		return a
	}
	return b
}

func Ter[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Select[T any](in []T, f func(i T) bool) (res []T) {
	res = make([]T, 0)
	for _, v := range in {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}

func CopyMap[K comparable, V any](m map[K]V) (res map[K]V) {
	res = make(map[K]V)
	for k, v := range m {
		res[k] = v
	}
	return
}

func CopySlice[V any](s []V) (res []V) {
	res = make([]V, len(s))
	for k, v := range s {
		res[k] = v
	}
	return
}

func Combo[T any](iterable []T, r int) chan []T {
	ch := make(chan []T)

	go func() {
		l := len(iterable)
		for combo := range GenCombo(l, r) {
			res := make([]T, r)
			for i, v := range combo {
				res[i] = iterable[v]
			}
			ch <- res
		}
		close(ch)
	}()
	return ch
}

func GenCombo(n, r int) <-chan []int {
	if r > n {
		panic("invalid argument")
	}
	ch := make(chan []int)

	go func() {
		res := make([]int, r)
		for i := range res {
			res[i] = i
		}
		t := make([]int, r)
		copy(t, res)
		ch <- t
		for {
			for i := r - 1; i >= 0; i-- {
				if res[i] < i+n-r {
					res[i]++
					for j := 1; j < r-i; j++ {
						res[i+j] = res[i] + j
					}
					t := make([]int, r)
					copy(t, res)
					ch <- t
					break
				}
			}
			if res[0] >= n-r {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func ChunkSlice[T any](input []T, size int) [][]T {
	out := make([][]T, 0, len(input)/size)
	for i := 0; i < len(input); i += size {
		chunk := make([]T, 0, size)
		for j := 0; j < size; j++ {
			chunk = append(chunk, input[i+j])
		}
		out = append(out, chunk)
	}
	return out
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
