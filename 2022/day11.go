package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day11.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	ms := parseInput()
	p(ms, 20, true)
	ms1 := parseInput()
	p(ms1, 10000, false)
}

func p(ms []*monkey, rounds int, divideByThree bool) {
	numInspected := make([]int, len(ms))
	modulo := 1
	for _, m := range ms {
		modulo *= m.test[0]
	}
	for r := 0; r < rounds; r++ {
		for i, m := range ms {
			items := m.getItems()
			numInspected[i] += len(items)
			for i := 0; i < len(items); i++ {
				item := items[i]
				var worryLvl int

				if divideByThree {
					worryLvl = m.op(item) / 3
				} else {
					worryLvl = m.op(item) % modulo
				}
				if worryLvl%m.test[0] == 0 {
					ms[m.test[1]].putItem(worryLvl)
				} else {
					ms[m.test[2]].putItem(worryLvl)
				}
			}
		}
	}

	sort.Slice(numInspected, func(i, j int) bool {
		return numInspected[i] > numInspected[j]
	})

	fmt.Println(numInspected[0] * numInspected[1])
}

type opFn func(int) int

type monkey struct {
	items        []int
	op           opFn
	test         [3]int
	numInspected int
}

func (m *monkey) getItems() []int {
	items := make([]int, len(m.items))
	copy(items, m.items)
	m.items = m.items[:0]
	return items
}

func (m *monkey) putItem(v int) {
	m.items = append(m.items, v)
}

func parseInput() []*monkey {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	b = append(b, '\n')
	var ms []*monkey
	si := strings.Split(strings.TrimSpace(string(b)), "\n")
	for i := 0; i < len(si); i += 7 {
		ss := si[i : i+6]
		ms = append(ms, parseMonkey(ss))
	}

	return ms
}

func parseMonkey(ss []string) *monkey {
	var items []int
	vals := strings.Split(strings.TrimPrefix(ss[1], "  Starting items: "), ", ")
	for _, val := range vals {
		v, _ := strconv.Atoi(val)
		items = append(items, v)
	}
	op := strings.TrimPrefix(ss[2], "  Operation: new = ")
	op = strings.ReplaceAll(op, "old", "x")
	eq := strings.Split(op, " ")
	//fmt.Println(eq)
	var fn opFn
	switch eq[1] {
	case "*":
		if eq[0] == "x" && eq[2] != "x" {
			v, _ := strconv.Atoi(eq[2])
			fn = func(i int) int { return i * v }
		} else if eq[0] != "x" && eq[2] == "x" {
			v, _ := strconv.Atoi(eq[0])
			fn = func(i int) int { return i * v }
		} else {
			fn = func(i int) int { return i * i }
		}
	case "+":
		if eq[0] == "x" && eq[2] != "x" {
			v, _ := strconv.Atoi(eq[2])
			fn = func(i int) int { return i + v }
		} else if eq[0] != "x" && eq[2] == "x" {
			v, _ := strconv.Atoi(eq[0])
			fn = func(i int) int { return i + v }
		} else {
			fn = func(i int) int { return i + i }
		}
	}

	v1, _ := strconv.Atoi(strings.TrimPrefix(ss[3], "  Test: divisible by "))
	v2, _ := strconv.Atoi(strings.TrimPrefix(ss[4], "    If true: throw to monkey "))
	v3, _ := strconv.Atoi(strings.TrimPrefix(ss[5], "    If false: throw to monkey "))
	test := [3]int{v1, v2, v3}

	return &monkey{
		items: items,
		op:    fn,
		test:  test,
	}
}
