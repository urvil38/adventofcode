package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day02.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(games []game) int {
	sum := 0
	for i, g := range games {
		shouldAdd := true
		for _, c := range g.cubes {
			for _, c := range c {
				if c.color == "red" && c.count > 12 {
					shouldAdd = false
				}
				if c.color == "green" && c.count > 13 {
					shouldAdd = false
				}
				if c.color == "blue" && c.count > 14 {
					shouldAdd = false
				}
			}
		}
		if shouldAdd {
			sum += i + 1
		}
	}
	return sum
}

func p2(games []game) int {
	sum := 0
	for _, g := range games {
		red, green, blue := 0, 0, 0
		for _, c := range g.cubes {
			for _, c := range c {
				if c.color == "red" && c.count >= red {
					red = c.count
				}
				if c.color == "green" && c.count >= green {
					green = c.count
				}
				if c.color == "blue" && c.count >= blue {
					blue = c.count
				}

			}
		}
		sum += (red * green * blue)
	}
	return sum
}

type game struct {
	cubes [][]cube
}

type cube struct {
	count int
	color string
}

func parseInput() []game {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")

	games := make([]game, 0)

	for _, s := range si {
		games = append(games, parseGame(s))
	}
	return games
}

func parseGame(s string) game {
	g := game{}
	si := strings.Split(s, ": ")
	sc := strings.Split(si[1], "; ")
	for _, c := range sc {
		g.cubes = append(g.cubes, parseCube(c))
	}
	return g
}

func parseCube(s string) []cube {
	cubes := make([]cube, 0)
	si := strings.Split(s, ", ")
	for _, s := range si {
		c := cube{}
		vi := strings.Split(s, " ")
		v, _ := strconv.Atoi(vi[0])
		c.count = v
		c.color = vi[1]
		cubes = append(cubes, c)
	}
	return cubes
}
