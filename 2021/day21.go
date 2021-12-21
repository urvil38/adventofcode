package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day21.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	input2 := parseInput()
	p2(input2)
}

type Player struct {
	pos   int
	score int
}

type Dice struct {
	r int
}

func p1(ps []*Player) {
	d := &Dice{}
	for {
		v1 := d.roll() + d.roll() + d.roll()
		ps[0].forward(v1)
		if ps[0].score >= 1000 {
			fmt.Println(d.r * ps[1].score)
			break
		}

		v2 := d.roll() + d.roll() + d.roll()
		ps[1].forward(v2)
		if ps[1].score >= 1000 {
			fmt.Println(d.r * ps[0].score)
			break
		}
	}
}

type gameState struct {
	p1 Player
	p2 Player
}

type win struct {
	p1 int
	p2 int
}

func p2(ps []*Player) {
	dp := make(map[gameState]win)
	winner := countWin(*ps[0], *ps[1], dp)
	fmt.Println(max(winner.p1, winner.p2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countWin(p1, p2 Player, dp map[gameState]win) win {
	if p1.score >= 21 {
		return win{1, 0}
	}
	if p2.score >= 21 {
		return win{0, 1}
	}
	gs := gameState{p1: p1, p2: p2}
	_, ok := dp[gs]
	if ok {
		return dp[gs]
	}
	var winner win
	dice := []int{1, 2, 3}

	for _, d1 := range dice {
		for _, d2 := range dice {
			for _, d3 := range dice {
				var newState Player
				newState.pos = (p1.pos + d1 + d2 + d3) % 10
				newState.score = p1.score + newState.pos + 1

				x := countWin(p2, newState, dp)
				winner.p1 += x.p2
				winner.p2 += x.p1
			}
		}
	}
	dp[gameState{p1: p1, p2: p2}] = winner
	return winner
}

func (d *Dice) roll() int {
	d.r++
	return d.r
}

func (p *Player) forward(v int) {
	p.pos = (p.pos + v) % 10
	p.score += p.pos + 1
}

func NewPlayer(pos int) *Player {
	return &Player{
		pos:   pos - 1,
		score: 0,
	}
}

func parseInput() []*Player {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var ps []*Player
	for _, s := range si {
		if s == "" {
			continue
		}
		ps = append(ps, parsePlayer(s))
	}

	return ps
}

func parsePlayer(s string) *Player {
	i := strings.IndexRune(s, ':')
	sPos := strings.TrimSpace(s[i+1:])
	pos, _ := strconv.Atoi(sPos)
	return NewPlayer(pos)
}
