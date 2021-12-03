package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input, nl := parseInput()
	fmt.Println(p1(input, nl))
	fmt.Println(p2(input, nl))
}

func p1(input []int, nl int) int {
	var gRate int
	ones := make([]int, nl)
	for _, v := range input {
		for k := nl - 1; k >= 0; k-- {
			if (v>>k)&1 == 1 {
				ones[k]++
			}
		}
	}
	for i, v := range ones {
		if v*2 > len(input) {
			gRate |= (1 << i)
		}
	}
	eRate := ^gRate & ((1 << nl) - 1)
	return gRate * eRate
}

func p2(input []int, nl int) int {
	return rate(input, nl, oxyRate) * rate(input, nl, co2Rate)
}

type computeRate func(a, b []int) []int

func co2Rate(s0, s1 []int) []int {
	if len(s0) <= len(s1) {
		return s0
	}
	return s1
}

func oxyRate(s0, s1 []int) []int {
	if len(s0) <= len(s1) {
		return s1
	}
	return s0
}

func rate(input []int, nl int, rateFn computeRate) int {
	curr := input
	var s0, s1 []int
	for k := nl - 1; k >= 0; k-- {
		if len(curr) == 1 {
			break
		}
		s0 = s0[:0]
		s1 = s1[:0]
		for _, v := range curr {
			if (v>>k)&1 == 1 {
				s1 = append(s1, v)
			} else {
				s0 = append(s0, v)
			}
		}
		curr = rateFn(s0, s1)
	}
	return curr[0]
}

func parseInput() ([]int, int) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input []int
	for _, v := range si {
		if v == "" {
			continue
		}
		vi, _ := strconv.ParseInt(v, 2, 64)
		input = append(input, int(vi))
	}

	return input, len(si[0])
}
