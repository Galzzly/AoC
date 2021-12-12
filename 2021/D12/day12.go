package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Cave struct {
	Name   string
	Small  bool
	Others []string
}

type Path struct {
	Path       []string
	Visited    map[string]bool
	VisitSmall bool
}

func main() {
	start := time.Now()
	lines := utils.ReadFileLineByLine(os.Args[1])
	caves := makeCaves(lines)
	t1 := time.Now()
	fmt.Println("Part 1", part1(caves), "Took", time.Since(t1))
	t2 := time.Now()
	fmt.Println("Part 2:", part2(caves), "Took", time.Since(t2))
	fmt.Println("Total Time:", time.Since(start))
}

func copyMap(m map[string]bool) (out map[string]bool) {
	out = map[string]bool{}
	for k, v := range m {
		out[k] = v
	}
	return
}

func part1(caves map[string]*Cave) (res int) {
	queue := []Path{{Path: []string{"start"}, Visited: map[string]bool{}}}
	paths := []string{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		cave := caves[current.Path[len(current.Path)-1]]
		if cave.Name == "end" {
			paths = append(paths, strings.Join(current.Path, ","))
		}

		newVisited := copyMap(current.Visited)
		if cave.Small {
			newVisited[cave.Name] = true
		}
		for _, cave := range cave.Others {
			if current.Visited[cave] {
				continue
			}

			newPath := make([]string, len(current.Path))
			copy(newPath, current.Path)
			newPath = append(newPath, cave)
			queue = append(queue, Path{Path: newPath, Visited: newVisited})
		}
	}
	res = len(paths)
	return
}

func part2(caves map[string]*Cave) (res int) {
	queue := []Path{{[]string{"start"}, map[string]bool{}, false}}
	paths := []string{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		cave := caves[current.Path[len(current.Path)-1]]
		if cave.Name == "end" {
			paths = append(paths, strings.Join(current.Path, ","))
			continue
		}

		newVisited := copyMap(current.Visited)
		if cave.Small {
			newVisited[cave.Name] = true
		}
		for _, cave := range cave.Others {
			VisitSmall := current.VisitSmall
			if current.Visited[cave] {
				if cave == "start" || current.VisitSmall {
					continue
				} else {
					VisitSmall = true
				}
			}

			newPath := make([]string, len(current.Path))
			copy(newPath, current.Path)
			newPath = append(newPath, cave)
			queue = append(queue, Path{newPath, newVisited, VisitSmall})
		}
	}
	res = len(paths)
	return
}

func makeCaves(lines []string) (caves map[string]*Cave) {
	caves = map[string]*Cave{}
	for _, line := range lines {
		sp := strings.Split(line, "-")
		a, b := sp[0], sp[1]
		if _, ok := caves[a]; !ok {
			caves[a] = &Cave{
				Name:   a,
				Small:  a[0] >= 'a',
				Others: []string{b},
			}
		} else {
			caves[a].Others = append(caves[a].Others, b)
		}

		a, b = b, a
		if _, ok := caves[a]; !ok {
			caves[a] = &Cave{
				Name:   a,
				Small:  a[0] >= 'a',
				Others: []string{b},
			}
		} else {
			caves[a].Others = append(caves[a].Others, b)
		}
	}
	return
}
