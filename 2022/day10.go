package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day10.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	ins := parseInput()
	p(ins)
}

func p(ins []instruction) {
	var pc, strength int
	var grid [40 * 6]string

	cycle := 1
	xReg := 1

	m := make(map[int]bool)
	for _, v := range []int{20, 60, 100, 140, 180, 220} {
		m[v] = true
	}

	for {
		if cycle > 240 {
			break
		}
		pos := (cycle - 1) % 40
		if pos == xReg-1 || pos == xReg || pos == xReg+1 {
			grid[cycle-1] = "#"
		}
		in := ins[pc]
		if m[cycle] {
			strength += cycle * xReg
		}
		if in.op == "noop" {
			pc++
		} else {
			in.iter++
			ins[pc] = in
			if in.iter == 2 {
				xReg += in.val
				pc++
			}
		}
		cycle++
	}

	fmt.Println(strength)

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			val := grid[(i*40)+j]
			if val == "" {
				fmt.Print(". ")
			} else {
				fmt.Print(val + " ")
			}
		}
		fmt.Println()
	}
}

type instruction struct {
	op   string
	val  int
	iter int
}

func parseInput() []instruction {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var input []instruction
	si := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, v := range si {
		if v == "" {
			continue
		}
		input = append(input, parseInstruction(v))
	}

	return input
}

func parseInstruction(s string) instruction {
	ss := strings.Split(s, " ")
	if ss[0] == "noop" {
		return instruction{op: ss[0], val: 0}
	}
	val, _ := strconv.Atoi(ss[1])
	return instruction{op: ss[0], val: val}
}
