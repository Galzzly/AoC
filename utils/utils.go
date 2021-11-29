package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

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