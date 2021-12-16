package main

import (
	"container/heap"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day15.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	findPath(input)
	ei := extendInput(input)
	findPath(ei)
}

type Pos struct {
	x, y int
}

type Item struct {
	x, y  int
	cost  int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func findPath(graph [][]int) {
	h := len(graph)
	w := len(graph[0])

	var q PriorityQueue
	q = append(q, &Item{x: 0, y: 0, cost: 0, index: 0})
	heap.Init(&q)

	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	dist := make(map[Pos]int)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			dist[Pos{x: i, y: j}] = 99999
		}
	}

	dist[Pos{x: 0, y: 0}] = 0

	for q.Len() > 0 {
		e := heap.Pop(&q).(*Item)

		if e.x == w-1 && e.y == h-1 {
			fmt.Println(dist[Pos{x: w - 1, y: h - 1}])
			return
		}

		for _, d := range dirs {
			dx := e.x + d[0]
			dy := e.y + d[1]

			if isValid(dx, dy, w, h) {
				next := Item{x: dx, y: dy, cost: e.cost + graph[dx][dy]}
				if next.cost < dist[Pos{x: dx, y: dy}] {
					heap.Push(&q, &next)
					dist[Pos{x: next.x, y: next.y}] = next.cost
				}
			}
		}
	}

}

func extendInput(input [][]int) [][]int {
	rows := len(input)
	cols := len(input[0])
	output := make([][]int, rows*5)
	for i := range output {
		output[i] = make([]int, cols*5)
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					inc := i + j
					val := 1 + ((input[row][col] + inc - 1) % 9)
					output[row+i*rows][col+j*cols] = val
				}
			}
		}
	}
	return output
}

func isValid(x, y, h, w int) bool {
	return x >= 0 && y >= 0 && x < w && y < h
}

func parseInput() [][]int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := make([][]int, 100)
	for i := range input {
		input[i] = make([]int, 100)
	}
	si := strings.Split(string(b), "\n")
	for r, row := range si {
		for c, v := range row {
			vi, _ := strconv.Atoi(string(v))
			input[r][c] = vi
		}
	}

	return input
}
