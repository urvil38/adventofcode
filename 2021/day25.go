package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day25.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
}

func p1(grid [][]int) {
	step := 1
	for {
		moves := update(grid)
		if moves == 0 {
			break
		}
		step++
	}
	fmt.Println(step)
}

type pos struct {
	r, c int
}

func update(grid [][]int) int {
	var moves int
	visited1 := make(map[pos]bool)
	w, h := len(grid[0]), len(grid)
	// update east facing (1) cucumbers
	for r := 0; r < h; r++ {
		for c := 0; c < w; c++ {
			if grid[r][c] == 1 && !visited1[pos{r: r, c: c}] {
				ci := (c + 1) % w
				if grid[r][ci] == 0 && !visited1[pos{r: r, c: ci}] {
					grid[r][c] = 0
					grid[r][ci] = 1
					visited1[pos{r: r, c: c}] = true
					visited1[pos{r: r, c: ci}] = true
					moves++
				}
			}
		}
	}
	visited2 := make(map[pos]bool)
	// update down facing (2) cucumbers
	for c := 0; c < w; c++ {
		for r := 0; r < h; r++ {
			if grid[r][c] == 2 && !visited2[pos{r: r, c: c}] {
				ri := (r + 1) % h
				if grid[ri][c] == 0 && !visited2[pos{r: ri, c: c}] {
					grid[r][c] = 0
					grid[ri][c] = 2
					visited2[pos{r: r, c: c}] = true
					visited2[pos{r: ri, c: c}] = true
					moves++
				}
			}
		}
	}
	return moves
}

func print(grid [][]int) {
	for _, row := range grid {
		for _, col := range row {
			if col == 1 {
				fmt.Print(">")
			} else if col == 2 {
				fmt.Print("v")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func parseInput() [][]int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input [][]int
	for _, v := range si {
		if v == "" {
			continue
		}
		var row []int
		for _, c := range v {
			if c == '>' {
				row = append(row, 1)
			} else if c == 'v' {
				row = append(row, 2)
			} else {
				row = append(row, 0)
			}
		}
		input = append(input, row)
	}

	return input
}
