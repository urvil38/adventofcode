package main

import (
	"adventofcode/lib"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day15.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	findPath(input)
	ei := extendInput(input)
	findPath(ei)
}

type Pos struct {
	x, y int
}

func findPath(graph [][]int) {
	h := len(graph)
	w := len(graph[0])

	wq := lib.NewWorkQueue[Pos]()
	wq.Add(Pos{0, 0}, Pos{0, 0}, 0)

	add := func(prev, p Pos, cost int) {
		if p.x < 0 || p.y < 0 || p.x >= h || p.y >= w {
			return
		}
		cost += graph[p.x][p.y]
		wq.Add(prev, p, cost)
	}
	
	visit := func(p Pos, cost int) {
		add(p, Pos{p.x - 1, p.y}, cost)
		add(p, Pos{p.x + 1, p.y}, cost)
		add(p, Pos{p.x, p.y - 1}, cost)
		add(p, Pos{p.x, p.y + 1}, cost)
	}

	for !wq.Empty() {
		e, cost := wq.Next()

		if e.x == w-1 && e.y == h-1 {
			fmt.Println(cost)
			return
		}

		visit(e, cost)
	}

}

func extendInput(input [][]int) [][]int {
	rows := len(input)
	cols := len(input[0])
	output := make([][]int, rows*5)
	for i := range output {
		output[i] = make([]int, cols*5)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					inc := i + j
					val := 1 + ((input[row][col] + inc - 1) % 9)
					output[row+i*rows][col+j*cols] = val
				}
			}
		}
	}
	return output
}

func parseInput() [][]int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := make([][]int, 100)
	for i := range input {
		input[i] = make([]int, 100)
	}
	si := strings.Split(string(b), "\n")
	for r, row := range si {
		for c, v := range row {
			vi, _ := strconv.Atoi(string(v))
			input[r][c] = vi
		}
	}

	return input
}
