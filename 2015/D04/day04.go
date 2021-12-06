package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	f := os.Args[1]
	line := utils.ReadFileLineByLine(f)[0]
	fmt.Println("Part 1:", part1(line))
	fmt.Println("Part 2:", part2(line))
}

func part1(s string) (res int) {
	regex := regexp.MustCompile("\\A00000")
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

func part2(s string) (res int) {
	regex := regexp.MustCompile("\\A000000")
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
