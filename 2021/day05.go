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
	var visited [1000][1000]int
	for _, l := range lines {
		drawLine(&visited, l)
	}
	fmt.Println(overlap(visited))
}

func p2(lines []Line) {
	var visited [1000][1000]int
	for _, l := range lines {
		drawLineWithDiagonal(&visited, l)
	}
	fmt.Println(overlap(visited))
}

func drawLine(visited *[1000][1000]int, l Line) {
	xf, yf := l.from.x, l.from.y
	xt, yt := l.to.x, l.to.y
	if xf == xt {
		if yf > yt {
			yf, yt = yt, yf
		}
		for i := yf; i <= yt; i++ {
			visited[i][xf] += 1
		}
	} else if yf == yt {
		if xf > xt {
			xf, xt = xt, xf
		}
		for i := xf; i <= xt; i++ {
			visited[yf][i] += 1
		}
	}
}

func drawLineWithDiagonal(visited *[1000][1000]int, l Line) {
	xf, yf := l.from.x, l.from.y
	xt, yt := l.to.x, l.to.y
	if xf == xt {
		if yf > yt {
			yf, yt = yt, yf
		}
		for i := yf; i <= yt; i++ {
			visited[i][xf] += 1
		}
	} else if yf == yt {
		if xf > xt {
			xf, xt = xt, xf
		}
		for i := xf; i <= xt; i++ {
			visited[yf][i] += 1
		}
	} else {
		if xt < xf {
			xf, yf, xt, yt = xt, yt, xf, yf
		}
		if yf < yt {
			i := xf
			j := yf

			for i <= xt && j <= yt {
				visited[j][i] += 1
				i++
				j++
			}
		} else {
			i := xf
			j := yf

			for i <= xt && j >= yt {
				visited[j][i] += 1
				i++
				j--
			}
		}

	}
}

func overlap(visited [1000][1000]int) int {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if visited[i][j] >= 2 {
				count++
			}
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
	for _, l := range inputlines {
		if l == "" {
			continue
		}
		lines = append(lines, parseLine(l))
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
