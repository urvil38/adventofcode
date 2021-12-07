package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day07.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input, lowest, highest := parseInput()
	p1(input, lowest, highest)
	p2(input, lowest, highest)
}

type CrabGroup map[int]int

func p1(cg CrabGroup, lo, hi int) {
	tc := math.MaxInt
	for i := lo; i < hi; i++ {
		c := cost(cg, i)
		if tc > c {
			tc = c
		} else {
			break
		}
	}
	fmt.Println(tc)
}

func cost(cg CrabGroup, pos int) int {
	fuel := 0
	for k, v := range cg {
		fuel += abs(k-pos) * v
	}
	return fuel
}

func p2(cg CrabGroup, lo, hi int) {
	tc := math.MaxInt
	for i := lo; i < hi; i++ {
		c := cost1(cg, i)
		if tc > c {
			tc = c
		} else {
			break
		}
	}
	fmt.Println(tc)
}

func cost1(cg CrabGroup, pos int) int {
	fuel := 0
	for k, v := range cg {
		// 1+2+3...+n = n(n+1)/2
		if k > pos {
			fuel += v * (k - pos) * (k - pos + 1) / 2
		} else {
			fuel += v * (pos - k) * (pos - k + 1) / 2
		}
	}
	return fuel
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func parseInput() (CrabGroup, int, int) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lo := math.MaxInt
	hi := math.MinInt
	si := strings.Split(strings.TrimSpace(string(b)), ",")
	cg := make(CrabGroup)
	for _, v := range si {
		if v == "" {
			continue
		}
		vi, _ := strconv.Atoi(v)
		if vi > hi {
			hi = vi
		}
		if vi < lo {
			lo = vi
		}
		cg[vi]++
	}

	return cg, lo, hi
}
