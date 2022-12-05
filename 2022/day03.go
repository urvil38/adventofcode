package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day03.input", "Relative file path to use as input.")

var priorities = make(map[string]int)

func main() {
	flag.Parse()
	input := parseInput()
	var v int
	for i := 'a'; i <= 'z'; i++ {
		v++
		priorities[string(i)] = v
	}
	for i := 'A'; i <= 'Z'; i++ {
		v++
		priorities[string(i)] = v
	}
	fmt.Println(p1(input))
	input2 := parseInputP2()
	fmt.Println(p2(input2))
}

func p1(rucksacks []rucksack) int {
	var total int
	for _, r := range rucksacks {
		c1, c2 := r.compartments[0], r.compartments[1]
		for _, c := range c1 {
			if c2.contains(c) {
				total += priorities[string(c)]
				break
			}
		}
	}
	return total
}

func p2(rucksacks []rucksack) int {
	var total int
	for _, r := range rucksacks {
		items := make(map[string]bool)
		var mapping []map[string]bool
		for _, c := range r.compartments {
			cc := make(map[string]bool)
			for _, item := range c {
				items[string(item)] = true
				cc[string(item)] = true
			}
			mapping = append(mapping, cc)
		}
		for i, _ := range items {
			if mapping[0][i] && mapping[1][i] && mapping[2][i] {
				total += priorities[i]
				break
			}
		}
	}
	return total
}

type rucksack struct {
	compartments []compartment
}

type compartment string

func (c compartment) contains(i rune) bool {
	return strings.ContainsRune(string(c), i)
}

func parseInput() []rucksack {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input []rucksack

	for _, v := range si {
		var r rucksack
		r.compartments = append(r.compartments, compartment(v[:len(v)/2]))
		r.compartments = append(r.compartments, compartment(v[len(v)/2:]))
		input = append(input, r)
	}

	return input
}

func parseInputP2() []rucksack {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input []rucksack

	for i := 0; i < len(si)-3; i += 3 {
		var r rucksack
		r.compartments = append(r.compartments, compartment(si[i]))
		r.compartments = append(r.compartments, compartment(si[i+1]))
		r.compartments = append(r.compartments, compartment(si[i+2]))
		input = append(input, r)
	}
	return input
}
