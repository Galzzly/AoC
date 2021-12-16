package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/Galzzly/AoC/utils"
)

var HTB = [16][]byte{
	{0, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 0}, {0, 0, 1, 1}, {0, 1, 0, 0}, {0, 1, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1},
	{1, 0, 0, 0}, {1, 0, 0, 1}, {1, 0, 1, 0}, {1, 0, 1, 1}, {1, 1, 0, 0}, {1, 1, 0, 1}, {1, 1, 1, 0}, {1, 1, 1, 1},
}

func main() {
	line := utils.ReadFileLineByLine("input")[0]
	// fmt.Println(line)
	hexBits := HexToBits(line)
	start := time.Now()
	_, part1, part2 := parse(hexBits, 0)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	fmt.Printf("Time: %v\n", time.Since(start))

}

func parse(hexBits []byte, start int) (newstart int, p1res, p2res int64) {
	p1res, _ = ReadBits(hexBits, start, 3)
	typeid, _ := ReadBits(hexBits, start+3, 3)
	var count int
	if typeid == 4 {
		newstart = start + 6
		p2res, count = ReadNumber(hexBits, newstart)
		newstart += count
		return
	} else {
		newstart = start + 6
		lengthid, count := ReadBits(hexBits, newstart, 1)
		newstart += count
		var results []int64
		if lengthid == 0 {
			length, count := ReadBits(hexBits, newstart, 15)
			newstart += count
			stop := newstart + int(length)
			var addC, res int64
			for {
				newstart, addC, res = parse(hexBits, newstart)
				p1res += addC
				results = append(results, res)
				if newstart >= stop {
					break
				}
			}
		} else {
			switch typeid {

			}
			length, count := ReadBits(hexBits, newstart, 11)
			newstart += count
			var addC, res int64
			for i := int64(0); i < length; i++ {
				newstart, addC, res = parse(hexBits, newstart)
				p1res += addC
				results = append(results, res)
			}
		}
		switch typeid {
		case 0:
			for _, v := range results {
				p2res += v
			}
		case 1:
			p2res = 1
			for _, v := range results {
				p2res *= v
			}
		case 2:
			p2res = math.MaxInt
			for _, v := range results {
				if v < p2res {
					p2res = v
				}
			}
		case 3:
			for _, v := range results {
				if v > p2res {
					p2res = v
				}
			}
		case 5:
			if results[0] > results[1] {
				p2res = 1
			}
		case 6:
			if results[0] < results[1] {
				p2res = 1
			}
		case 7:
			if results[0] == results[1] {
				p2res = 1
			}
		}
	}
	return
}

func HexToBits(hexString string) (res []byte) {
	for _, c := range hexString {
		v, _ := strconv.ParseInt(string(c), 16, 8)
		res = append(res, HTB[v]...)
	}
	return
}

func ReadBits(data []byte, start, count int) (res int64, c int) {
	for _, b := range data[start : start+count] {
		res <<= 1
		res |= int64(b)
	}
	return res, count
}

func ReadNumber(data []byte, start int) (res int64, count int) {
	for {
		p, _ := ReadBits(data, start, 5)
		res <<= 4
		res |= int64(p & 0x0f)
		start += 5
		count += 5
		if p&0x10 == 0 {
			break
		}

	}
	return
}
