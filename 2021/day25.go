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
	// update east facing (1) cucumbers
	for r, row := range grid {
		for c, col := range row {
			if col == 1 && !visited1[pos{r: r, c: c}] {
				if c+1 < len(row) {
					if grid[r][c+1] == 0 && !visited1[pos{r: r, c: c + 1}] {
						grid[r][c] = 0
						grid[r][c+1] = 1
						visited1[pos{r: r, c: c}] = true
						visited1[pos{r: r, c: c + 1}] = true
						moves++
					}
				} else {
					if grid[r][0] == 0 && !visited1[pos{r: r, c: 0}] {
						grid[r][c] = 0
						grid[r][0] = 1
						visited1[pos{r: r, c: c}] = true
						visited1[pos{r: r, c: 0}] = true
						moves++
					}
				}
			}
		}
	}
	visited2 := make(map[pos]bool)
	// update down facing (2) cucumbers
	for c := 0; c < len(grid[0]); c++ {
		for r := 0; r < len(grid); r++ {
			if grid[r][c] == 2 && !visited2[pos{r: r, c: c}] {
				if r+1 < len(grid) {
					if grid[r+1][c] == 0 && !visited2[pos{r: r + 1, c: c}] {
						grid[r][c] = 0
						grid[r+1][c] = 2
						visited2[pos{r: r, c: c}] = true
						visited2[pos{r: r + 1, c: c}] = true
						moves++
					}
				} else {
					if grid[0][c] == 0 && !visited2[pos{r: 0, c: c}] {
						grid[r][c] = 0
						grid[0][c] = 2
						visited2[pos{r: r, c: c}] = true
						visited2[pos{r: 0, c: c}] = true
						moves++
					}
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
