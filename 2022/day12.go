package main

import (
	"adventofcode/lib"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

var inputFile = flag.String("inputFile", "inputs/day12.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	grid := parseInput()
	p1(grid, true)
	p2(grid, true)
}

type Pos struct {
	x, y int
}

func p1(grid [][]byte, debug bool) {
	for i, row := range grid {
		for j, col := range row {
			if col == 'S' {
				fmt.Println(findMinCost(grid, Pos{x: i, y: j}, debug))
				return
			}
		}
	}
}

func p2(grid [][]byte, debug bool) {
	min := math.MaxInt

	for i, row := range grid {
		if row[0] == 'a' {
			cost := findMinCost(grid, Pos{x: i, y: 0}, debug)
			if cost != 0 && cost <= min {
				min = cost
			}
		}
	}

	fmt.Println(min)
}

func parseInput() [][]byte {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	return bytes.Fields(b)
}

func findMinCost(grid [][]byte, start Pos, debug bool) int {
	h := len(grid)
	w := len(grid[0])

	wq := lib.NewWorkQueue[Pos]()
	wq.Add(start, start, 0)

	add := func(prev, p Pos, cost int) {
		if p.x < 0 || p.y < 0 || p.x >= h || p.y >= w {
			return
		}

		pv := grid[prev.x][prev.y]
		cv := grid[p.x][p.y]

		if pv == 'S' {
			pv = 'a'
		}
		if cv == 'E' {
			cv = 'z'
		}

		if pv+1 == cv || pv >= cv {
			wq.Add(prev, p, cost)
		}
	}

	visit := func(p Pos, c int) {
		add(p, Pos{x: p.x, y: p.y + 1}, c+1)
		add(p, Pos{x: p.x, y: p.y - 1}, c+1)
		add(p, Pos{x: p.x + 1, y: p.y}, c+1)
		add(p, Pos{x: p.x - 1, y: p.y}, c+1)
	}

	for !wq.Empty() {
		p, cost := wq.Next()
		if grid[p.x][p.y] == 'E' {
			if debug {
				ss := make([][]string, h)
				for i := range ss {
					ss[i] = make([]string, w)
					for j := range ss[i] {
						ss[i][j] = "."
					}
				}
				for _, d := range wq.Path(p) {
					ss[d.x][d.y] = string(grid[d.x][d.y])
				}
				for _, row := range ss {
					for _, col := range row {
						fmt.Print(col)
					}
					fmt.Println()
				}
			}
			return cost
		}
		visit(p, cost)
	}
	return 0
}
