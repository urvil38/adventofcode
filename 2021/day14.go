package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day14.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	tmpl, rules := parseInput()
	p1(tmpl, rules, 10)
	p1(tmpl, rules, 40)
}

func p1(tmpl string, rules map[string]string, steps int) {
	elements := make(map[string]int)
	pairs := make(map[string]int)
	for _, c := range tmpl {
		elements[string(c)]++
	}

	for i := 0; i < len(tmpl)-1; i++ {
		pairs[tmpl[i:i+2]]++
	}

	for i := 0; i < steps; i++ {
		newPairs := make(map[string]int)
		for pair, v := range pairs {
			e := rules[pair]
			elements[e] += v
			p1 := string(pair[0]) + string(e)
			p2 := string(e) + string(pair[1])
			newPairs[p1] += v
			newPairs[p2] += v
		}
		pairs = newPairs
	}

	var arr []int
	for _, e := range elements {
		arr = append(arr, e)
	}

	sort.Ints(arr)
	fmt.Println(arr[len(arr)-1] - arr[0])
}

func parseInput() (string, map[string]string) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	tmpl := si[0]
	rules := make(map[string]string)
	for i, v := range si {
		if v == "" || i < 2 {
			continue
		}
		s := strings.Split(v, " -> ")
		rules[s[0]] = s[1]
	}

	return tmpl, rules
}
