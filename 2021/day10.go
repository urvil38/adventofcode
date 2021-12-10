package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day10.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	legalLines := p1(input)
	p2(legalLines)
}

func p1(input []string) []string {
	var legalLines []string
	illegalChar := make(map[rune]int)
	for _, str := range input {
		s := NewStack()
		isLegal := true
		for _, c := range str {
			if startBrace(c) {
				s.push(c)
			} else {
				if !s.empty() {
					top := s.pop()
					if matchingBraces(top, c) {
						continue
					}

					isLegal = false
					illegalChar[c]++
				}
			}
		}
		if isLegal {
			legalLines = append(legalLines, str)
		}
	}
	fmt.Println(cost(illegalChar))
	return legalLines
}

func p2(lines []string) {
	var scores []int
	for _, str := range lines {
		s := NewStack()
		for _, c := range str {
			if startBrace(c) {
				s.push(c)
			} else {
				if !s.empty() {
					s.pop()
				}
			}
		}
		si := score(s)
		if si != 0 {
			scores = append(scores, si)
		}
	}
	sort.Ints(scores)
	mid := len(scores) / 2
	fmt.Println(scores[mid])
}

func matchingBraces(top, c rune) bool {
	if top == '<' && c == '>' || top == '[' && c == ']' || top == '{' && c == '}' || top == '(' && c == ')' {
		return true
	}
	return false
}

func startBrace(c rune) bool {
	if c == '{' || c == '[' || c == '<' || c == '(' {
		return true
	}
	return false
}

func score(s *stack) int {
	ss := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	score := 0
	for !s.empty() {
		c := s.pop()
		score = ss[c] + (score * 5)
	}
	return score
}

func cost(illegalChar map[rune]int) int {
	ss := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	c := 0
	for k, v := range illegalChar {
		c += ss[k] * v
	}
	return c
}

type stack struct {
	s []rune
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) push(x rune) rune {
	s.s = append(s.s, x)
	return x
}

func (s *stack) pop() rune {
	if len(s.s) == 0 {
		var v rune
		return v
	}

	e := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return e
}

func (s stack) top() rune {
	if len(s.s) == 0 {
		var v rune
		return v
	}

	return s.s[len(s.s)-1]
}

func (s stack) empty() bool {
	if len(s.s) == 0 {
		return true
	}
	return false
}

func (s stack) print() {
	for _, v := range s.s {
		fmt.Printf("[%c] ", v)
	}
	fmt.Println()
}

func parseInput() []string {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(b), "\n")
}
