package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Galzzly/AoC/utils"
)

type Workflow map[string][]Rule
type Rule struct {
	cat    string
	sign   string
	num    int
	target string
}
type Parts []map[string]int
type Values map[string][2]int

func main() {
	start := time.Now()
	f := "input.txt"
	sections := utils.ReadFileDoubleLine(f)
	workflows, parts := getSections(sections)
	for i, b := range []bool{false, true} {
		timer := time.Now()
		fmt.Printf("Part %d: %d, Took: %v\n", i+1, solve(workflows, parts, b), time.Since(timer))
	}
	fmt.Println("Total Time:", time.Since(start))
}

func solve(workflows Workflow, parts Parts, p2 bool) (res int) {
	if p2 {
		var values = Values{
			"x": {1, 4000},
			"m": {1, 4000},
			"a": {1, 4000},
			"s": {1, 4000},
		}
		return count(workflows, "in", values)
	}
nextpart:
	for _, R := range parts {
		wf := "in"
		for {
			for _, rule := range workflows[wf] {
				if rule.num == -1 {
					if rule.target == "R" {
						continue nextpart
					} else if rule.target == "A" {
						res += R["x"] + R["m"] + R["a"] + R["s"]
						continue nextpart
					}
					wf = rule.target
					break
				}
				if rule.sign == "<" {
					if R[rule.cat] < rule.num {
						if rule.target == "R" {
							continue nextpart
						} else if rule.target == "A" {
							res += R["x"] + R["m"] + R["a"] + R["s"]
							continue nextpart
						}
						wf = rule.target
						break
					}
					continue
				}
				if R[rule.cat] > rule.num {
					if rule.target == "R" {
						continue nextpart
					} else if rule.target == "A" {
						res += R["x"] + R["m"] + R["a"] + R["s"]
						continue nextpart
					}
					wf = rule.target
					break
				}
			}
		}
	}
	return
}

func count(workflows Workflow, wf string, vals Values) (res int) {
	if wf == "R" {
		return 0
	} else if wf == "A" {
		res = 1
		for _, v := range vals {
			res *= (v[1] - v[0] + 1)
		}
		return
	}

	for _, r := range workflows[wf] {
		V := vals[r.cat]
		var good, bad [2]int
		if r.sign == "<" {
			good = [2]int{V[0], r.num - 1}
			bad = [2]int{r.num, V[1]}
		} else if r.sign == ">" {
			good = [2]int{r.num + 1, V[1]}
			bad = [2]int{V[0], r.num}
		} else {
			res += count(workflows, r.target, vals)
			continue
		}

		if good[0] <= good[1] {
			val2 := utils.CopyMap[string, [2]int](vals)
			val2[r.cat] = good
			res += count(workflows, r.target, val2)
		}

		if bad[0] > bad[1] {
			break
		}

		vals[r.cat] = bad
	}

	return
}

func getSections(lines []string) (workflows Workflow, parts Parts) {
	workflows = make(Workflow, 0)
	for _, line := range strings.Split(lines[0], "\n") {
		s := strings.Split(line, "{")
		name := s[0]
		R := strings.Split(strings.TrimSuffix(s[1], "}"), ",")
		rules := []Rule{}
		for _, r := range R {
			if strings.ContainsAny(r, "<>") {
				rs := strings.Split(r, ":")
				rules = append(rules, Rule{string(rs[0][0]), string(rs[0][1]), utils.Atoi(rs[0][2:]), rs[1]})
				continue
			}
			rules = append(rules, Rule{num: -1, target: r})
		}
		workflows[name] = rules
	}

	P := strings.Split(lines[1], "\n")
	parts = make(Parts, len(P))
	for i, line := range P {
		line = strings.TrimPrefix(strings.TrimSuffix(line, "}"), "{")
		s := strings.Split(line, ",")
		parts[i] = make(map[string]int, len(s))
		for _, rating := range s {
			sp := strings.Split(rating, "=")
			parts[i][sp[0]] = utils.Atoi(sp[1])
		}
	}
	return
}
