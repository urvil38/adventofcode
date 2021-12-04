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
	fmt.Println(p1(input, boards))
	fmt.Println(p2(input, boards))
}

type board struct {
	values  [5][5]int
	visited [5][5]int
}

func p1(input []int, boards []board) int {
	var winnerBoard, currInput int
out:
	for _, v := range input {
		for i := 0; i < len(boards); i++ {
			set(&boards[i], v)
			if bingo(boards[i]) {
				winnerBoard = i
				currInput = v
				break out
			}
		}
	}

	return currInput * sumNotVisited(boards[winnerBoard])

}

func p2(input []int, boards []board) int {
	var winnerBoard, currInput int
	wmap := make(map[int]bool)
out:
	for _, v := range input {
		for i := 0; i < len(boards); i++ {
			set(&boards[i], v)
			if !wmap[i] {
				if bingo(boards[i]) {
					wmap[i] = true
				}
			}

			if len(wmap) == len(boards) {
				winnerBoard = i
				currInput = v
				break out
			}
		}
	}

	return currInput * sumNotVisited(boards[winnerBoard])

}

func sumNotVisited(b board) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.visited[i][j] == 0 {
				sum += b.values[i][j]
			}
		}
	}
	return sum
}

func set(b *board, v int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.values[i][j] == v {
				b.visited[i][j] = 1
			}
		}
	}
}

func bingo(b board) bool {
	var rSum, cSum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			rSum += b.visited[i][j]
			cSum += b.visited[j][i]
		}

		if rSum == 5 || cSum == 5 {
			return true
		}
		rSum, cSum = 0, 0
	}
	return false
}

func parseInput() ([]int, []board) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var input []int
	var boards []board

	for _, v := range strings.Split(lines[0], ",") {
		vi, _ := strconv.Atoi(v)
		input = append(input, vi)
	}

	for i := 2; i < len(lines); i += 6 {
		boards = append(boards, parseBoard(lines[i:i+5]))
	}

	return input, boards
}

func parseBoard(lines []string) board {
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
	return b
}
