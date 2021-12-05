package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day05.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	lines := parseInput()
	p1(lines)
	p2(lines)
}

type Point struct {
	x int
	y int
}

type Line struct {
	from Point
	to   Point
}

func p1(lines []Line) {
	visited := make(map[Point]int)
	for _, l := range lines {
		drawLine(visited, l)
	}
	fmt.Println(overlap(visited))
}

func p2(lines []Line) {
	visited := make(map[Point]int)
	for _, l := range lines {
		drawLineWithDiagonal(visited, l)
	}
	fmt.Println(overlap(visited))
}

func drawLine(visited map[Point]int, l Line) {
	start := l.from
	end := l.to
	if start.x == end.x {
		if end.y < start.y {
			start, end = end, start
		}

		for i := start.y; i <= end.y; i++ {
			visited[Point{start.x, i}]++
		}
	} else if start.y == end.y {
		if end.x < start.x {
			start, end = end, start
		}

		for i := start.x; i <= end.x; i++ {
			visited[Point{i, start.y}]++
		}
	}
}

func drawLineWithDiagonal(visited map[Point]int, l Line) {
	start := l.from
	end := l.to
	if start.x == end.x {
		if end.y < start.y {
			start, end = end, start
		}

		for i := start.y; i <= end.y; i++ {
			visited[Point{start.x, i}]++
		}
	} else if start.y == end.y {
		if end.x < start.x {
			start, end = end, start
		}

		for i := start.x; i <= end.x; i++ {
			visited[Point{i, start.y}]++
		}
	} else {
		if end.x < start.x {
			start, end = end, start
		}

		incY := 1
		if end.y < start.y {
			incY = -1
		}

		j := start.y
		for i := start.x; i <= end.x; i++ {
			visited[Point{i, j}]++
			j += incY
		}
	}
}

func overlap(visited map[Point]int) int {
	count := 0
	for _, v := range visited {
		if v >= 2 {
			count++
		}
	}
	return count
}

func parseInput() []Line {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	inputlines := strings.Split(string(b), "\n")
	var lines []Line
	for _, line := range inputlines {
		if line == "" {
			continue
		}
		lines = append(lines, parseLine(line))
	}

	return lines
}

func parseLine(line string) Line {
	ps := strings.Split(line, " -> ")
	return Line{
		from: parsePoint(ps[0]),
		to:   parsePoint(ps[1]),
	}
}

func parsePoint(line string) Point {
	var p Point
	ps := strings.Split(line, ",")
	px, _ := strconv.Atoi(ps[0])
	py, _ := strconv.Atoi(ps[1])
	p.x = px
	p.y = py
	return p
}
