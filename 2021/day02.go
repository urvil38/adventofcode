package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(dirs [][]string) int {
	xPos := 0
	yPos := 0

	for _, d := range dirs {
		dist, _ := strconv.Atoi(d[1])
		switch d[0] {
		case "forward":
			xPos += dist
		case "up":
			yPos -= dist
		case "down":
			yPos += dist
		}
	}

	return xPos * yPos
}

func p2(dirs [][]string) int {
	xPos := 0
	yPos := 0
	aim := 0

	for _, d := range dirs {
		dist, _ := strconv.Atoi(d[1])
		switch d[0] {
		case "forward":
			xPos += dist
			yPos += aim * dist
		case "up":
			aim -= dist
		case "down":
			aim += dist
		}
	}

	return xPos * yPos
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
