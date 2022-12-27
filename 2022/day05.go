package main

import (
	"adventofcode/lib"
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
	input := parseInput()
	p1(input)
	p2(input)
}

/*
            [L] [M]         [M]
        [D] [R] [Z]         [C] [L]
        [C] [S] [T] [G]     [V] [M]
[R]     [L] [Q] [B] [B]     [D] [F]
[H] [B] [G] [D] [Q] [Z]     [T] [J]
[M] [J] [H] [M] [P] [S] [V] [L] [N]
[P] [C] [N] [T] [S] [F] [R] [G] [Q]
[Z] [P] [S] [F] [F] [T] [N] [P] [W]
 1   2   3   4   5   6   7   8   9
*/

func stacks() []*lib.Stack[string] {
	var ss []*lib.Stack[string]
	ss = append(ss, newStack([]string{"R", "H", "M", "P", "Z"}))
	ss = append(ss, newStack([]string{"B", "J", "C", "P"}))
	ss = append(ss, newStack([]string{"D", "C", "L", "G", "H", "N", "S"}))
	ss = append(ss, newStack([]string{"L", "R", "S", "Q", "D", "M", "T", "F"}))
	ss = append(ss, newStack([]string{"M", "Z", "T", "B", "Q", "P", "S", "F"}))
	ss = append(ss, newStack([]string{"G", "B", "Z", "S", "F", "T"}))
	ss = append(ss, newStack([]string{"V", "R", "N"}))
	ss = append(ss, newStack([]string{"M", "C", "V", "D", "T", "L", "G", "P"}))
	ss = append(ss, newStack([]string{"L", "M", "F", "J", "N", "Q", "W"}))
	return ss
}

func p1(ops []Op) {
	ss := stacks()
	for _, op := range ops {
		from := ss[op.from-1]
		to := ss[op.to-1]
		for i := 0; i < op.count; i++ {
			e := from.Pop()
			to.Push(e)
		}
	}

	for _, s := range ss {
		fmt.Print(s.Top())
	}
	fmt.Println()
}

func p2(ops []Op) {
	ss := stacks()
	for _, op := range ops {
		from := ss[op.from-1]
		to := ss[op.to-1]
		var vals []string
		for i := 0; i < op.count; i++ {
			e := from.Pop()
			vals = append(vals, e)
		}
		for i := len(vals) - 1; i >= 0; i-- {
			to.Push(vals[i])
		}
	}

	for _, s := range ss {
		fmt.Print(s.Top())
	}
	fmt.Println()
}

func newStack(vals []string) *lib.Stack[string] {
	s := lib.NewStack[string]()
	for i := len(vals) - 1; i >= 0; i-- {
		s.Push(vals[i])
	}
	return s
}

type Op struct {
	from  int
	to    int
	count int
}

func parseInput() []Op {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var ops []Op

	var start bool
	for _, line := range lines {
		if line == "" {
			start = true
		}

		if start {
			if line == "" {
				continue
			}
			ss := strings.Split(line, " ")
			count, from, to := atoi(ss[1]), atoi(ss[3]), atoi(ss[5])
			ops = append(ops, Op{count: count, from: from, to: to})
		}

	}
	return ops
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
