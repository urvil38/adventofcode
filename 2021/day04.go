package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input, boards := parseInput()
	p1(input, boards)
	p2(input, boards)
}

type board struct {
	values [5][5]int
}

func p1(input []int, boards []*board) {
out:
	for _, v := range input {
		for _, b := range boards {
			b.set(v)
			if b.bingo() {
				fmt.Println(b.score(v))
				break out
			}
		}
	}
}

func p2(input []int, boards []*board) {
	won := make(map[int]bool)
out:
	for _, v := range input {
		for i, b := range boards {
			if won[i] {
				continue
			}

			b.set(v)
			if b.bingo() {
				won[i] = true
				if len(won) == len(boards) {
					fmt.Println(b.score(v))
					break out
				}
			}
		}
	}
}

func (b board) score(v int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.values[i][j] != -1 {
				sum += b.values[i][j]
			}
		}
	}
	return sum * v
}

func (b *board) set(v int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.values[i][j] == v {
				b.values[i][j] = -1
			}
		}
	}
}

func (b board) bingo() bool {
	var rc, cc int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.values[i][j] == -1 {
				rc++
			}
			if b.values[j][i] == -1 {
				cc++
			}
		}

		if rc == 5 || cc == 5 {
			return true
		}
		rc, cc = 0, 0
	}
	return false
}

func parseInput() ([]int, []*board) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var input []int
	var boards []*board

	for _, v := range strings.Split(lines[0], ",") {
		vi, _ := strconv.Atoi(v)
		input = append(input, vi)
	}

	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, parseBoard(lines[i:i+5]))
	}

	return input, boards
}

func parseBoard(lines []string) *board {
	var b board
	for r, line := range lines {
		var c int
		for _, v := range strings.Split(line, " ") {
			if v == "" {
				continue
			}
			vi, _ := strconv.Atoi(v)
			b.values[r][c] = vi
			c++
		}
	}
	return &b
}
