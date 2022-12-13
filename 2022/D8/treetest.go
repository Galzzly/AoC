package main

import (
	"fmt"
	"image"
	"sync"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Tree struct {
	height                int
	up, down, right, left bool
	isvis                 bool
}

type Treemapper map[image.Point]Tree

func main() {
	start := time.Now()
	f := "input.txt"
	lines := utils.ReadFileLineByLine(f)
	treemap := buildMap(lines)
	fmt.Println(treemap)
	fmt.Println("Took:", time.Since(start))
}

func buildMap(lines []string) Treemapper {
	treemap := Treemapper{}
	// Map the heights
	for y, line := range lines {
		for x, c := range line {
			treemap[image.Point{x, y}] = Tree{height: int(c), isvis: false}
		}
	}

	fmt.Println("here")

	// Build the visibilities
	var wg sync.WaitGroup
	wg.Add(4)
	go treemap.visUp(&wg)
	go treemap.visDown(&wg)
	go treemap.visLeft(&wg)
	go treemap.visDown(&wg)
	wg.Wait()
	return treemap
}

func (t Treemapper) visUp(wg *sync.WaitGroup) {
	defer wg.Done()
	dir := image.Point{0, -1}
	for p, tree := range t {
		for np := p.Add(dir); ; np = np.Add(dir) {
			if _, ok := t[np]; !ok {
				tree.up = true
			}
		}
	}
}

func (t Treemapper) visDown(wg *sync.WaitGroup) {
	defer wg.Done()
	dir := image.Point{0, 1}
	for p, tree := range t {
		for np := p.Add(dir); ; np = np.Add(dir) {
			if _, ok := t[np]; !ok {
				tree.down = true
			}
		}
	}
}

func (t Treemapper) visLeft(wg *sync.WaitGroup) {
	defer wg.Done()
	dir := image.Point{-1, 0}
	for p, tree := range t {
		for np := p.Add(dir); ; np = np.Add(dir) {
			if _, ok := t[np]; !ok {
				tree.left = true
			}
		}
	}
}

func (t *Treemapper) visRight(wg *sync.WaitGroup) {
	defer wg.Done()
	dir := image.Point{1, 0}
	for p, tree := range *t {
		for np := p.Add(dir); ; np = np.Add(dir) {
			if _, ok := (*t)[np]; !ok {
				tree.right = true
			}
		}
	}
}
