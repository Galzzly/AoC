	package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

func main() {
	start := time.Now()
	f := "input.txt"

	stream := utils.ReadFileSingleLine(f)
	t1 := time.Now()
	fmt.Println("Part 1:", streamcheck(stream, 4), "- Took", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", streamcheck(stream, 14), "Took:", time.Since(t2))
	fmt.Println(time.Since(start))
}

func streamcheck(stream string, length int) (res int) {
	for i := 0; i < len(stream)-length; i++ {
		if checkUnique(stream[i : i+length]) {
			return i + length
		}
	}
	return
}

func checkUnique(str string) bool {
	for _, c := range str {
		if strings.Count(str, string(c)) > 1 {
			return false
		}
	}
	return true
}
