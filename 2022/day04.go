package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

type pair struct {
	start int
	end   int
}

func p1(input []pair) int {
	var count int
	for i := 0; i < len(input); i = i + 2 {
		p1, p2 := input[i], input[i+1]
		if p1.contains(p2) || p2.contains(p1) {
			count++
		}
	}
	return count
}

func p2(input []pair) int {
	var count int
	for i := 0; i < len(input); i = i + 2 {
		p1, p2 := input[i], input[i+1]
		if p1.overlap(p2) && p2.overlap(p1) {
			count++
		}
	}
	return count
}

func (p pair) contains(p2 pair) bool {
	if p.start <= p2.start && p.end >= p2.end {
		return true
	}
	return false
}

func (p pair) overlap(p2 pair) bool {
	if p.start < p2.end && p.end < p2.start {
		return false
	}
	return true
}

func (p pair) String() string {
	var s []string
	for i := 0; i <= p.end; i++ {
		s = append(s, ".")
	}
	for i := p.start; i <= p.end; i++ {
		s[i] = strconv.Itoa(i)
	}
	fmt.Println(fmt.Sprintf("%v-%v", p.start, p.end))
	return strings.Join(s, "")
}

func parseInput() []pair {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var pairs []pair

	for _, line := range lines {
		if line == "" {
			continue
		}
		for _, v := range strings.Split(line, ",") {
			pairs = append(pairs, parsePair(v))
		}
	}

	return pairs

}

func parsePair(p string) pair {
	ss := strings.Split(p, "-")
	start, _ := strconv.Atoi(ss[0])
	end, _ := strconv.Atoi(ss[1])
	return pair{start: start, end: end}
}
