package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInputAsInts()
	fmt.Printf("p1: %v\n", p1(input))
	fmt.Printf("p2: %v\n", p2(input))
}

func p1(input []int) int {
	count := 0

	for i := 0; i < len(input)-1; i++ {
		if input[i+1] > input[i] {
			count++
		}
	}
	return count
}

func p2(input []int) int {
	count := 0

	for i := 0; i < len(input)-3; i++ {
		v1 := input[i] + input[i+1] + input[i+2]
		v2 := input[i+1] + input[i+2] + input[i+3]
		if v2 > v1 {
			count++
		}
	}

	return count
}

func parseInputAsInts() []int {
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
		vi, _ := strconv.Atoi(v)
		input = append(input, vi)
	}

	return input
}
