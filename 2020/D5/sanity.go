package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s: ", elapsed)
}

func getRow(r string) int {
	row := make([]int, 127)
	for i := range row {
		row[i] = i
	}
	//fmt.Printf("%d\n", len(r))
	for i := 0; i < len(r); i++ {
		fmt.Printf("%s\n", r[i])
	}
	return 0
}

func main() {
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	for _, l := range lines {
		//fmt.Printf("%s\n", l[7:])
		getRow(l[:7])
	}
}
