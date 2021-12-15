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
	findPathP1(input)
	ei := extendInput(input)
	findPathP2(ei)
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
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
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

func findPathP1(graph [100][100]int) {
	q := make(PriorityQueue, 1)
	q[0] = &Item{x: 0, y: 0, cost: 0, index: 0}
	heap.Init(&q)
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	dist := make(map[Pos]int)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			dist[Pos{x: i, y: j}] = 99999
		}
	}
	dist[Pos{x: 0, y: 0}] = 0
	for q.Len() > 0 {
		e := heap.Pop(&q).(*Item)

		if e.x == 99 && e.y == 99 {
			fmt.Println(dist[Pos{x: 99, y: 99}])
			return
		}

		for _, d := range dirs {
			dx := e.x + d[0]
			dy := e.y + d[1]

			if isValid(dx, dy, 100, 100) {
				next := Item{x: dx, y: dy, cost: e.cost + graph[dx][dy]}
				if next.cost < dist[Pos{x: dx, y: dy}] {
					heap.Push(&q, &next)
					dist[Pos{x: next.x, y: next.y}] = next.cost
				}
			}
		}
	}

}

func findPathP2(graph [500][500]int) {
	q := make(PriorityQueue, 1)
	q[0] = &Item{x: 0, y: 0, cost: 0, index: 0}
	heap.Init(&q)
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	dist := make(map[Pos]int)
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			dist[Pos{x: i, y: j}] = 99999
		}
	}
	dist[Pos{x: 0, y: 0}] = 0
	for q.Len() > 0 {
		e := heap.Pop(&q).(*Item)

		if e.x == 499 && e.y == 499 {
			fmt.Println(dist[Pos{x: 499, y: 499}])
			return
		}

		for _, d := range dirs {
			dx := e.x + d[0]
			dy := e.y + d[1]

			if isValid(dx, dy, 500, 500) {
				next := Item{x: dx, y: dy, cost: e.cost + graph[dx][dy]}
				if next.cost < dist[Pos{x: dx, y: dy}] {
					heap.Push(&q, &next)
					dist[Pos{x: next.x, y: next.y}] = next.cost
				}
			}
		}
	}

}

func extend(input [100][100]int) [][100][100]int {
	var extend [][100][100]int

	for k := 0; k < 5; k++ {
		var arr [100][100]int
		for i := 0; i < 100; i++ {
			for j := 0; j < 100; j++ {
				v := input[i][j] + k
				if v > 9 {
					v = v % 9
				}
				arr[i][j] = v
			}
		}
		extend = append(extend, arr)
	}

	return extend
}

func extendInput(input [100][100]int) [500][500]int {
	arrs := extend(input)
	var final [][100][100]int
	var output [500][500]int
	final = append(final, arrs...)
	for i := 1; i < 5; i++ {
		final = append(final, extend(arrs[i])...)
	}
	k := 0
	for i := 0; i < 500; i += 100 {
		for j := 0; j < 500; j += 100 {
			v1 := final[k]
			for r, row := range v1 {
				for c, col := range row {
					output[r+i][c+j] = col
				}
			}
			k++
		}
	}
	return output
}

func isValid(x, y, h, w int) bool {
	return x >= 0 && y >= 0 && x < w && y < h
}

func parseInput() [100][100]int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var input [100][100]int
	si := strings.Split(string(b), "\n")
	for r, row := range si {
		for c, v := range row {
			vi, _ := strconv.Atoi(string(v))
			input[r][c] = vi
		}
	}

	return input
}
