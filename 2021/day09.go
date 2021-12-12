package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day09.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input, w, h := parseInput()
	lowPoints := p1(input, w, h)
	p2(lowPoints, w, h)
}

func p1(input [][]int, w, h int) [][]int {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var res []int
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			var hi []int
			for _, d := range dirs {
				dx := row + d[0]
				dy := col + d[1]
				if isValid(dx, dy, w, h) {
					hi = append(hi, input[dx][dy])
				}
			}
			if min(hi, input[row][col]) {
				res = append(res, input[row][col])
				input[row][col] = -1
			}
		}
	}
	fmt.Println(riskLevel(res))
	return input
}

func p2(input [][]int, w, h int) {
	var visited [][]bool
	for i := 0; i < h; i++ {
		var r []bool
		for j := 0; j < w; j++ {
			r = append(r, false)
		}
		visited = append(visited, r)
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	var res []int
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if !visited[row][col] && input[row][col] == -1 {
				q := NewQueue[Pos]()
				q.enqueue(Pos{row: row, col: col})
				visited[row][col] = true
				var count int
				for !q.empty() {
					count++
					e := q.dequeue()
					for _, d := range dirs {
						dx := e.row + d[0]
						dy := e.col + d[1]

						if isValid(dx, dy, w, h) && !visited[dx][dy] && input[dx][dy] != 9 {
							q.enqueue(Pos{row: dx, col: dy})
							visited[dx][dy] = true
						}
					}
				}
				res = append(res, count)
			}
		}
	}
	sort.Ints(res)
	fmt.Println(res[len(res)-1] * res[len(res)-2] * res[len(res)-3])

}

func riskLevel(heights []int) int {
	risk := 0
	for _, h := range heights {
		risk += h + 1
	}
	return risk
}

func min(a []int, v int) bool {
	for _, a1 := range a {
		if a1 <= v {
			return false
		}
	}
	return true
}

func isValid(row, col, w, h int) bool {
	return row >= 0 && col >= 0 && row < h && col < w
}

type Pos struct {
	row int
	col int
}

type queue[T any] struct {
	q []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) enqueue(x T) T {
	q.q = append(q.q, x)
	return x
}

func (q *queue[T]) dequeue() T {
	if len(q.q) == 0 {
		var v T
		return v
	}

	e := q.q[0]
	q.q = q.q[1:]
	return e
}

func (q queue[T]) front() T {
	if len(q.q) == 0 {
		var v T
		return v
	}

	return q.q[0]
}

func (q queue[T]) empty() bool {
	if len(q.q) == 0 {
		return true
	}
	return false
}

func parseInput() ([][]int, int, int) {
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
		input = append(input, parseSignalEntry(v))
	}

	return input, len(input[0]), len(input)
}

func parseSignalEntry(s string) []int {
	var vals []int
	for _, v := range s {
		vi, _ := strconv.Atoi(string(v))
		vals = append(vals, vi)
	}
	return vals
}
