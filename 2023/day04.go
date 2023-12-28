package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day04.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	cards := parseInput()
	fmt.Println(p1(cards))
	fmt.Println(p2(cards))
}
func p1(cards []Card) int {
	sum := 0
	for _, card := range cards {
		if card.match > 0 {
			sum += (1 << (card.match - 1))
		}
	}
	return sum
}

func p2(cards []Card) int {
	mm := make(map[int]int)
	total := 0
	for i, card := range cards {
		_, ok := mm[i]
		if !ok {
			mm[i] = 1
		}
		for j := i + 1; j < i+1+card.match; j++ {
			v, ok := mm[j]
			if !ok {
				mm[j] = 1 + mm[i]
			} else {
				mm[j] = v + mm[i]
			}
		}
	}

	for _, v := range mm {
		total += v
	}
	return total
}

type Card struct {
	nums  []int
	wins  []int
	set   map[int]bool
	match int
}

func parseInput() []Card {
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	cs := strings.Split(string(bytes), "\n")
	var cards []Card
	for i, card := range cs {
		c := parseCard(card, i)
		cards = append(cards, c)
	}
	return cards
}

func parseCard(card string, i int) Card {
	c := Card{
		nums: make([]int, 0),
		wins: make([]int, 0),
		set:  make(map[int]bool),
	}
	_, a, _ := strings.Cut(card, ":")
	win, all, _ := strings.Cut(a, " | ")
	for _, s := range strings.Split(win, " ") {
		if s == "" {
			continue
		}
		v, _ := strconv.Atoi(s)
		c.wins = append(c.wins, v)
	}

	for _, s := range strings.Split(all, " ") {
		if s == "" {
			continue
		}
		v, _ := strconv.Atoi(s)
		c.nums = append(c.nums, v)
		_, ok := c.set[v]
		if !ok {
			c.set[v] = true
		}
	}

	for _, v := range c.wins {
		_, ok := c.set[v]
		if ok {
			c.match += 1
		}
	}

	return c
}
