package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"
	line := utils.ReadFileLineByLine(f)[0]
	t1 := time.Now()
	p1 := part1(line)
	fmt.Println("Part 1:", p1, "Took:", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(line, p1), "Took:", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func part1(s string) (res int) {
	regex := regexp.MustCompile("^00000")
	hashcheck := ""

	for !regex.MatchString(hashcheck) {
		res++
		hash := md5.New()
		io.WriteString(hash, s)
		io.WriteString(hash, fmt.Sprintf("%d", res))
		hashcheck = fmt.Sprintf("%x", hash.Sum(nil))
	}

	return
}

func part2(s string, n int) (res int) {
	regex := regexp.MustCompile("^000000")
	hashcheck := ""
	res = n
	for !regex.MatchString(hashcheck) {
		res++
		hash := md5.New()
		io.WriteString(hash, s)
		io.WriteString(hash, fmt.Sprintf("%d", res))
		hashcheck = fmt.Sprintf("%x", hash.Sum(nil))
	}
	return
}
