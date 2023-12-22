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
	grid := parseInput()
	fmt.Println(p1(grid))
	fmt.Println(p2(grid))
}

type Point struct {
	r, c int
}

func p1(grid [][]string) int {
	set := make(map[Point]bool)
	for r, row := range grid {
		for c, _ := range row {
			if grid[r][c] == "." || isDigit(grid[r][c]) {
				continue
			}

			for _, i := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					ri := r + i
					cj := c + j
					if ri >= 0 && ri < len(grid) && cj >= 0 && cj < len(row) && isDigit(grid[ri][cj]) {
						for cj > 0 && isDigit(grid[ri][cj-1]) {
							cj--
						}
						_, ok := set[Point{ri, cj}]
						if !ok {
							set[Point{ri, cj}] = true
						}
					}
				}
			}
		}
	}

	sum := 0

	for k := range set {
		r := k.r
		c := k.c
		s := ""
		for c < len(grid[r]) && isDigit(grid[r][c]) {
			s += grid[r][c]
			c++
		}
		i, _ := strconv.Atoi(s)
		sum += i
	}

	return sum
}

func p2(grid [][]string) int {
	set := make(map[Point]bool)
	count := make(map[Point][]Point)
	for r, row := range grid {
		for c, _ := range row {
			if grid[r][c] == "." || isDigit(grid[r][c]) || grid[r][c] != "*" {
				continue
			}

			for _, i := range []int{-1, 0, 1} {
				for _, j := range []int{-1, 0, 1} {
					ri := r + i
					cj := c + j
					if ri >= 0 && ri < len(grid) && cj >= 0 && cj < len(row) && isDigit(grid[ri][cj]) {
						for cj > 0 && isDigit(grid[ri][cj-1]) {
							cj--
						}
						_, ok := set[Point{ri, cj}]
						if !ok {
							set[Point{ri, cj}] = true
							count[Point{r, c}] = append(count[Point{r, c}], Point{ri, cj})
						}
					}
				}
			}
		}
	}
	sum := 0
	for _, vv := range count {
		if len(vv) == 2 {
			mul := 1
			for _, k := range vv {
				r := k.r
				c := k.c
				s := ""
				for c < len(grid[r]) && isDigit(grid[r][c]) {
					s += grid[r][c]
					c++
				}
				i, _ := strconv.Atoi(s)
				mul *= i
			}
			sum += mul
		}
	}

	return sum
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func parseInput() [][]string {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	ss := strings.Split(string(bytes), "\n")
	var grid [][]string

	for _, s := range ss {
		grid = append(grid, strings.Split(s, ""))
	}
	return grid
}
