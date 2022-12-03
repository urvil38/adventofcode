package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var mapping = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var win = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
	"X": "C",
	"Y": "A",
	"Z": "B",
}

var lose = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
	"X": "B",
	"Y": "C",
	"Z": "A",
}

var draw = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
	"X": "A",
	"Y": "B",
	"Z": "C",
}
var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input [][]string) int {
	var score int
	for _, moves := range input {
		op, me := moves[0], moves[1]
		if win[me] == op {
			score += 6
			score += mapping[me]
		} else if draw[me] == op {
			score += 3
			score += mapping[me]
		} else {
			score += mapping[me]
		}
	}
	return score
}

func p2(input [][]string) int {
	var score int
	for _, moves := range input {
		op, me := moves[0], moves[1]
		if me == "X" {
			score += mapping[lose[op]]
		} else if me == "Y" {
			score += 3
			score += mapping[draw[op]]
		} else {
			score += 6
			score += mapping[win[op]]
		}
	}
	return score
}

func parseInput() [][]string {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input [][]string
	for _, v := range si {
		if v == "" {
			continue
		}
		input = append(input, strings.Split(v, " "))
	}

	return input
}
