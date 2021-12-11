package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day11.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	g1 := parseInput()
	p1(&g1, 100)
	g2 := parseInput()
	p2(&g2)
}

type grid [10][10]int

func p1(g *grid, step int) {
	totalFlash := 0
	for i := 0; i < step; i++ {
		totalFlash += g.update()
	}
	fmt.Println(totalFlash)
}

func p2(g *grid) {
	for i := 1; ; i++ {
		flash := g.update()
		if flash == 100 {
			fmt.Println(i)
			return
		}
	}
}

func (g *grid) update() int {
	flash := 0
	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			g[row][col]++
		}
	}

	for row := 0; row < 10; row++ {
		for col := 0; col < 10; col++ {
			if g[row][col] > 9 {
				g.flash(row, col, &flash)
			}
		}
	}

	return flash
}

func (g *grid) flash(row, col int, flash *int) {
	if g[row][col] <= 9 {
		return
	}

	g[row][col] = 0
	*flash++

	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	for _, d := range dirs {
		dx := row + d[0]
		dy := col + d[1]

		if isValid(dx, dy) && g[dx][dy] > 0 {
			g[dx][dy] += 1
			g.flash(dx, dy, flash)
		}
	}
}

func isValid(row, col int) bool {
	return row >= 0 && col >= 0 && row < 10 && col < 10
}

func parseInput() grid {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var input grid
	for r, line := range lines {
		for c, d := range line {
			vi, _ := strconv.Atoi(string(d))
			input[r][c] = vi
		}
	}
	return input
}
