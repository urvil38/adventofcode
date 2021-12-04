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

func p1(input []int, boards [][][]int) int {
	visited := initVisited(len(boards))
	var winnerBoard, currInput int
out:
	for _, v := range input {
		for i := 0; i < len(boards); i++ {
			set(boards[i], visited[i], v)
			if dowehaveWinner(visited[i]) {
				winnerBoard = i
				currInput = v
				break out
			}
		}
	}

	return currInput * sumNotVisited(boards[winnerBoard], visited[winnerBoard])

}

func p2(input []int, boards [][][]int) int {
	visited := initVisited(len(boards))
	var winnerBoard, currInput int
	wmap := make(map[int]bool)
out:
	for _, v := range input {
		for i := 0; i < len(boards); i++ {
			set(boards[i], visited[i], v)
			if !wmap[i] {
				if dowehaveWinner(visited[i]) {
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

	return currInput * sumNotVisited(boards[winnerBoard], visited[winnerBoard])

}

func initVisited(totalBoards int) [][][]int {
	var visited [][][]int
	for i := 0; i < totalBoards; i++ {
		var board [][]int
		for k := 0; k < 5; k++ {
			var line []int
			for j := 0; j < 5; j++ {
				line = append(line, 0)
			}
			board = append(board, line)
		}
		visited = append(visited, board)
	}
	return visited
}

func sumNotVisited(board, visited [][]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if visited[i][j] == 0 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func set(board, visited [][]int, v int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == v {
				visited[i][j] = 1
			}
		}
	}
}

func dowehaveWinner(board [][]int) bool {
	var rSum, cSum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			rSum += board[i][j]
			cSum += board[j][i]
		}

		if rSum == 5 || cSum == 5 {
			return true
		}
		rSum, cSum = 0, 0
	}
	return false
}

func parseInput() ([]int, [][][]int) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input []int
	var boardLines [][]int
	for i, v := range si {
		if v == "" {
			continue
		}
		if i == 0 {
			vs := strings.Split(v, ",")
			for _, v := range vs {
				vi, _ := strconv.Atoi(v)
				input = append(input, vi)
			}
			continue
		}

		vs := strings.Split(v, " ")
		var line []int
		for _, v := range vs {
			if v == "" {
				continue
			}
			vi, _ := strconv.Atoi(v)
			line = append(line, vi)
		}
		boardLines = append(boardLines, line)
	}
	var boards [][][]int
	k := 0
	var bs [][]int
	for _, line := range boardLines {
		bs = append(bs, line)
		k++
		if k == 5 {
			k = 0
			boards = append(boards, bs)
			bs = [][]int{}
		}
	}

	return input, boards
}
