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

func p1(population []int, totalDays int) int {
	var lifetime [9]int

	for _, v := range population {
		lifetime[v]++
	}

	for day := 0; day < totalDays; day++ {
		fs := lifetime
		for k, v := range fs {
			if k == 0 {
				lifetime[0] -= v
				lifetime[6] += v
				lifetime[8] += v
			} else {
				lifetime[k] -= v
				lifetime[k-1] += v
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
