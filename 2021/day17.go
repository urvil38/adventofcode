package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day17.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
}

func p1(t Target) {
	targetArea := make(TargetArea)
	if t.start.y < t.end.y {
		t.start.y, t.end.y = t.end.y, t.start.y
	}
	for i := t.start.x; i <= t.end.x; i++ {
		for j := t.end.y; j <= t.start.y; j++ {
			targetArea[Pos{x: i, y: j}] = true
		}
	}

	var prob Probe
	prob.ta = targetArea
	prob.t = t

	var count int
	for i := -t.end.x; i <= t.end.x; i++ {
		for j := t.end.y; j <= -t.end.y; j++ {
			prob.setVelocity(i, j)
			prob.pos = Pos{x: 0, y: 0}
			hitTarget := prob.move()
			if hitTarget {
				count++
			}
		}
	}
	fmt.Println(prob.maxY)
	fmt.Println(count)
}

type TargetArea map[Pos]bool

type Probe struct {
	pos  Pos
	v    Velocity
	ta   TargetArea
	t    Target
	maxY int
}

type Velocity struct {
	x, y int
}

type Pos struct {
	x, y int
}

func (p *Probe) setVelocity(x, y int) {
	p.v.x = x
	p.v.y = y
}

func (p *Probe) move() (hitTarget bool) {
	var localMaxY int

	for {
		p.pos.x += p.v.x
		p.pos.y += p.v.y

		if p.v.x > 0 {
			p.v.x -= 1
		} else if p.v.x < 0 {
			p.v.x += 1
		}

		p.v.y -= 1
		if p.pos.y > localMaxY {
			localMaxY = p.pos.y
		}

		if p.inTarget() {
			if localMaxY > p.maxY {
				p.maxY = localMaxY
			}
			return true
		} else if p.movePastTarget() {
			return false
		}
	}
}

func (p Probe) movePastTarget() bool {
	return p.pos.x > p.t.end.x || p.pos.y < p.t.end.y
}

func (p Probe) inTarget() bool {
	return p.ta[p.pos]
}

type Target struct {
	start Pos
	end   Pos
}

func parseInput() Target {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Trim(string(b), "\n")
	cords := strings.TrimPrefix(si, "target area: ")
	var t Target
	poss := strings.Split(cords, ", ")
	p1 := parsePos(poss[0][2:])
	t.start.x = p1.x
	t.end.x = p1.y

	p2 := parsePos(poss[1][2:])
	t.start.y = p2.x
	t.end.y = p2.y

	return t
}

func parsePos(si string) Pos {
	var p Pos
	s := strings.Split(si, "..")
	x1, _ := strconv.Atoi(s[0])
	x2, _ := strconv.Atoi(s[1])
	p.x = x1
	p.y = x2
	return p
}
