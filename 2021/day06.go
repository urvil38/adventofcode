package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day06.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(p1(input, 80))
	fmt.Println(p1(input, 256))
}

type fish struct {
	k int
	v int
}

func fishes(lifetime map[int]int) []fish {
	var fs []fish
	for k, v := range lifetime {
		fs = append(fs, fish{k: k, v: v})
	}
	return fs
}

func p1(population []int, totalDays int) int {
	lifetime := make(map[int]int)

	for _, v := range population {
		lifetime[v]++
	}

	for day := 0; day < totalDays; day++ {
		fs := fishes(lifetime)
		for _, f := range fs {
			if f.k == 0 {
				lifetime[0] -= f.v
				lifetime[6] += f.v
				lifetime[8] += f.v
			} else {
				lifetime[f.k] -= f.v
				lifetime[f.k-1] += f.v
			}
		}
	}

	count := 0
	for _, v := range lifetime {
		count += v
	}

	return count
}

func parseInput() []int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(strings.TrimSpace(string(b)), ",")
	var input []int
	for _, v := range si {
		if v == "" {
			continue
		}
		vi, _ := strconv.Atoi(v)
		input = append(input, vi)
	}
	return input
}
