package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day08.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	p2(input)
}

func p1(grid [][]int) {
	w, h := len(grid[0]), len(grid)
	treeCount := 0
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if onEdge(row, col, w, h) || isVisible(grid, row, col) {
				treeCount++
			}
		}
	}
	fmt.Println(treeCount)
}

func p2(grid [][]int) {
	w, h := len(grid[0]), len(grid)
	var maxScenicScore int
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			score := getScenicScore(grid, row, col)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	fmt.Println(maxScenicScore)
}

func isVisible(grid [][]int, r, c int) bool {
	w := len(grid[0])
	h := len(grid)
	treeHeight := grid[r][c]
	top, bottom, right, left := true, true, true, true
	// top dir
	for i := r - 1; i >= 0; i-- {
		if grid[i][c] >= treeHeight {
			top = false
		}
	}
	// right dir
	for i := c + 1; i < w; i++ {
		if grid[r][i] >= treeHeight {
			right = false
		}
	}
	// bottom dir
	for i := r + 1; i < h; i++ {
		if grid[i][c] >= treeHeight {
			bottom = false
		}
	}
	// left dir
	for i := c - 1; i >= 0; i-- {
		if grid[r][i] >= treeHeight {
			left = false
		}
	}
	return top || bottom || right || left
}

func getScenicScore(grid [][]int, r, c int) int {
	w := len(grid[0])
	h := len(grid)
	treeHeight := grid[r][c]
	var top, bottom, right, left int
	// top dir
	for i := r - 1; i >= 0; i-- {
		if grid[i][c] < treeHeight {
			top++
		} else if grid[i][c] >= treeHeight {
			top++
			break
		}
	}
	// right dir
	for i := c + 1; i < w; i++ {
		if grid[r][i] < treeHeight {
			right++
		} else if grid[r][i] >= treeHeight {
			right++
			break
		}
	}
	// bottom dir
	for i := r + 1; i < h; i++ {
		if grid[i][c] < treeHeight {
			bottom++
		} else if grid[i][c] >= treeHeight {
			bottom++
			break
		}
	}
	// left dir
	for i := c - 1; i >= 0; i-- {
		if grid[r][i] < treeHeight {
			left++
		} else if grid[r][i] >= treeHeight {
			break
		}
	}
	return top * bottom * right * left
}

func onEdge(r, c, w, h int) bool {
	return r == 0 || r == h-1 || c == 0 || c == w-1
}

func parseInput() [][]int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var input [][]int
	si := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, v := range si {
		if v == "" {
			continue
		}
		input = append(input, parseLine(v))
	}

	return input
}

func parseLine(s string) []int {
	var vals []int
	for _, v := range s {
		vi, _ := strconv.Atoi(string(v))
		vals = append(vals, vi)
	}
	return vals
}

func startWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
