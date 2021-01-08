package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

type Rule struct {
	lit     bool
	content string
	rNum    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()
	file := os.Args[1]
	f, err := ioutil.ReadFile(file)
	check(err)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n")
	// Make a map of the rules
	rules := make(map[string]*Rule)
	for _, v := range strings.Split(lines[0], "\n") {
		s := strings.Split(v, ": ")
		rules[s[0]] = getRules(v)
	}
	// Make a slice of the messages
	m := []string{}
	for _, v := range strings.Split(lines[1], "\n") {
		m = append(m, v)
	}
	t1 := time.Now()
	fmt.Printf("Part 1: %d (%s)\n", part1(rules, m), time.Since(t1))
	t2 := time.Now()
	fmt.Printf("Part 2: %d (%s)\n", part2(rules, m), time.Since(t2))
	fmt.Printf("Total Time: %s\n", time.Since(start))
}

func getRules(rStr string) *Rule {
	split := strings.Split(rStr, ": ")
	s := split[1]
	if s[0] == '"' {
		return &Rule{
			rNum:    split[0],
			content: s[1 : len(s)-1],
			lit:     true,
		}
	}
	return &Rule{
		rNum:    split[0],
		content: s,
	}
}

func part1(rules map[string]*Rule, m []string) int {
	rules1 := make(map[string]*Rule)
	rules1 = rules
	rStr := rules1["0"].resolve(rules1, false)
	re := pcre.MustCompile("^"+rStr+"$", 0)
	count := 0
	for _, v := range m {
		if re.MatcherString(v, 0).Matches() {
			count++
		}
	}
	return count
}
func part2(rules map[string]*Rule, m []string) int {
	rules2 := make(map[string]*Rule)
	for k := range rules {
		rules2[k] = rules[k]
		if k == "8" {
			rules2[k].content = "42"
		} else if k == "11" {
			rules2[k].content = "42 31"
		}
	}
	rStr := rules2["0"].resolve(rules2, true)
	re := pcre.MustCompile("^"+rStr+"$", 0)
	count := 0
	for _, v := range m {
		if re.MatcherString(v, 0).Matches() {
			count++
		}
	}
	return count
}

func (r *Rule) resolve(rules map[string]*Rule, p2 bool) string {
	if r == nil {
		return ""
	}
	if r.lit {
		return r.content
	}
	out := strings.Builder{}
	out.WriteString("(?:")
	s := strings.Split(r.content, " | ")
	orR := []string{}
	for _, rS := range s {
		s := strings.Split(rS, " ")
		cRules := []string{}
		for _, rNum := range s {
			if p2 {
				switch rNum {
				case "8":
					this := rules["42"].resolve(rules, p2)
					cRules = append(cRules, fmt.Sprintf("%s+", this))
					continue
				case "11":
					this := rules["42"].resolve(rules, p2)
					thenThis := rules["31"].resolve(rules, p2)
					cRules = append(cRules, fmt.Sprintf("(?<eleven>(%s%s|%[1]s(?&eleven)%[2]s))", this, thenThis))
					continue
				}
			}
			toAdd := rules[rNum].resolve(rules, p2)
			cRules = append(cRules, toAdd)
		}
		orR = append(orR, strings.Join(cRules, ""))
	}
	out.WriteString(strings.Join(orR, "|"))
	out.WriteRune(')')
	return out.String()
}
