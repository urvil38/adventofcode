package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day09.input", "Relative file path to use as input.")

var dirs = [][]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, -1}, {1, -1}, {1, 1}, {-1, 1}}

func main() {
	flag.Parse()
	moves := parseInput()
	p(moves, 2)
	p(moves, 10)
}

func p(moves []move, knots int) {
	tailPosHistory := make(map[Pos]bool)

	rope := make([]Pos, knots)

	tailPosHistory[Pos{x: 0, y: 0}] = true

	for _, m := range moves {

		for step := 0; step < m.steps; step++ {
			switch m.dir {
			case "R":
				rope[0].y++
			case "L":
				rope[0].y--
			case "U":
				rope[0].x++
			case "D":
				rope[0].x--
			}

			for i := 1; i < knots; i++ {
				headPos := i - 1
				tailPos := i
				head := rope[headPos]
				tail := rope[tailPos]

				shouldUpdatePos := true
				for _, d := range dirs {
					var t Pos
					t.x = head.x + d[0]
					t.y = head.y + d[1]
					if t.equal(tail) {
						shouldUpdatePos = false
					}
				}

				if shouldUpdatePos {
					newTail := tail
					if head.y == tail.y {
						if head.x < tail.x {
							newTail.x--
						} else {
							newTail.x++
						}
					} else if head.x == tail.x {
						if head.y < tail.y {
							newTail.y--
						} else {
							newTail.y++
						}
					} else {
						if head.x > tail.x && head.y > tail.y {
							newTail.x++
							newTail.y++
						}
						if head.x > tail.x && head.y < tail.y {
							newTail.x++
							newTail.y--
						}
						if head.x < tail.x && head.y < tail.y {
							newTail.x--
							newTail.y--
						}
						if head.x < tail.x && head.y > tail.y {
							newTail.x--
							newTail.y++
						}
					}
					rope[tailPos] = newTail
					if tailPos == knots-1 {
						tailPosHistory[newTail] = true
					}
				}
			}
		}
	}
	fmt.Println(len(tailPosHistory))
}

func (p Pos) equal(p1 Pos) bool {
	return p.x == p1.x && p.y == p1.y
}

type Pos struct {
	x int
	y int
}

type move struct {
	dir   string
	steps int
}

func parseInput() []move {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var input []move
	si := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, v := range si {
		if v == "" {
			continue
		}
		input = append(input, parseMove(v))
	}

	return input
}

func parseMove(s string) move {
	ss := strings.Split(s, " ")
	var m move
	m.dir = ss[0]
	m.steps, _ = strconv.Atoi(ss[1])
	return m
}
