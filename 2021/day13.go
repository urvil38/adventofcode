package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day13.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	points, folds := parseInput()
	p1(points, folds[0])
	p2(points, folds)
}

type Point struct {
	x, y int
}

type Fold struct {
	dir string
	v   int
}

func p1(points []Point, fold Fold) {
	pm := make(map[Point]int)
	for _, p := range points {
		pm[p]++
	}
	if fold.dir == "x" {
		foldLeft(fold.v, pm)
	} else {
		foldUp(fold.v, pm)
	}
	fmt.Println(len(pm))
}

func p2(points []Point, folds []Fold) {
	pm := make(map[Point]int)
	for _, p := range points {
		pm[p]++
	}
	for _, fold := range folds {
		if fold.dir == "x" {
			foldLeft(fold.v, pm)
		} else {
			foldUp(fold.v, pm)
		}
	}
	printpaper(pm)
}

func printpaper(pm map[Point]int) {
	w, h := dimention(pm)
	var grid [][]string
	for row := 0; row <= h; row++ {
		var str []string
		for col := 0; col <= w; col++ {
			str = append(str, " ")
		}
		grid = append(grid, str)
	}

	for p, _ := range pm {
		grid[p.y][p.x] = "#"
	}

	for _, s := range grid {
		fmt.Println(s)
	}
}

func dimention(pm map[Point]int) (int, int) {
	h := -1
	w := -1
	for p, _ := range pm {
		if p.y > h {
			h = p.y
		}
		if p.x > w {
			w = p.x
		}
	}
	return w, h
}

func foldUp(v int, pm map[Point]int) {
	for p, _ := range pm {
		if p.y > v {
			newP := p
			delete(pm, p)
			newP.y -= (p.y - v) * 2
			pm[newP]++
		}
	}
}

func foldLeft(v int, pm map[Point]int) {
	for p, _ := range pm {
		if p.x > v {
			newP := p
			delete(pm, p)
			newP.x -= (p.x - v) * 2
			pm[newP]++
		}
	}
}
func parseInput() ([]Point, []Fold) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var points []Point
	var folds []Fold
	for _, v := range si {
		if v == "" {
			continue
		}
		if strings.ContainsRune(v, ',') {
			points = append(points, parsePoint(v))
		} else {
			folds = append(folds, parseFold(v))
		}
	}

	return points, folds
}

func parsePoint(s string) Point {
	sp := strings.Split(s, ",")
	sx, _ := strconv.Atoi(sp[0])
	sy, _ := strconv.Atoi(sp[1])
	return Point{x: sx, y: sy}
}

func parseFold(s string) Fold {
	sp := strings.Split(s, "=")
	dirs := sp[0]
	v, _ := strconv.Atoi(sp[1])
	return Fold{v: v, dir: string(dirs[len(dirs)-1])}
}
