package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

var inputFile = flag.String("inputFile", "inputs/day12.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	input2 := parseInput()
	p2(input2)
}

func p1(graph CaveMap) {
	findPathP1(graph)
}

func p2(graph CaveMap) {
	findPathP2(graph)
}

type Path []string

type CaveMap map[string][]string

type PathWithRepeat struct {
	repeatCave string
	visited    Path
}

func (p Path) isVisited(v string) bool {
	if !isLower(v) {
		return false
	}
	for _, c := range p {
		if v == c {
			return true
		}
	}
	return false
}

func findPathP1(graph CaveMap) {
	q := NewQueue[Path]()
	var pathCount int
	var path Path
	path = append(path, "start")
	q.enqueue(path)
	for !q.empty() {
		path = q.dequeue()

		last := path[len(path)-1]
		if last == "end" {
			pathCount++
			continue
		}

		for _, v := range graph[last] {
			if !path.isVisited(v) {
				var newPath Path
				newPath = append(newPath, path...)
				newPath = append(newPath, v)
				q.enqueue(newPath)
			}
		}
	}
	fmt.Println(pathCount)
}

func (p PathWithRepeat) isVisited(c string) (bool, bool) {
	if !isLower(c) {
		return false, false
	}
	if p.repeatCave == c || c == "start" {
		return true, false
	}

	for _, v := range p.visited {
		if c == v {
			if p.repeatCave == "" {
				// allow to visit but needs to mark the cave as repeated
				return false, true
			} else {
				// already visited and don't needs to repeat again!
				return true, false
			}
		}
	}

	return false, false
}

func findPathP2(graph CaveMap) {
	var pathCount int
	q := NewQueue[PathWithRepeat]()
	var pathR PathWithRepeat
	pathR.visited = append(pathR.visited, "start")
	q.enqueue(pathR)
	for !q.empty() {
		item := q.dequeue()
		path := item.visited
		last := path[len(path)-1]
		if last == "end" {
			pathCount++
			continue
		}
		for _, v := range graph[last] {
			visited, repeat := item.isVisited(v)
			if !visited {
				var newPath PathWithRepeat
				newPath.repeatCave = item.repeatCave
				newPath.visited = append(newPath.visited, item.visited...)
				newPath.visited = append(newPath.visited, v)
				if repeat {
					newPath.repeatCave = v
				}
				q.enqueue(newPath)
			}
		}
	}
	fmt.Println(pathCount)
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

func parseInput() CaveMap {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	cMap := make(CaveMap)
	si := strings.Split(string(b), "\n")
	for _, v := range si {
		if v == "" {
			continue
		}
		c1, c2 := parseCave(v)
		cMap[c1] = append(cMap[c1], c2)
		cMap[c2] = append(cMap[c2], c1)
	}

	return cMap
}

func parseCave(input string) (string, string) {
	cs := strings.Split(input, "-")
	return cs[0], cs[1]
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
