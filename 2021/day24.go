package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var inputFile = flag.String("inputFile", "inputs/day24.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	blocks := parseInput()
	p1(blocks)
}

func p1(blocks []block) {
	visited := make(map[answer]int)
	ans := find(visited, blocks, 0, 0)
	fmt.Println(ans)
}

func getReg(s string) int {
	switch s {
	case "w":
		return 0
	case "x":
		return 1
	case "y":
		return 2
	case "z":
		return 3
	}
	return 0
}

func find(dp map[answer]int, blocks []block, block, z int) int {
	if block == len(blocks) {
		if z == 0 {
			return 0
		} else {
			return -1
		}
	}

	if ans, ok := dp[answer{block: block, z: z}]; ok && ans != -1 {
		return ans
	}

	for digit := 1; digit < 10; digit++ {
		reg := [4]int{digit, 0, 0, z}
		for _, in := range blocks[block] {
			if in[0] == "add" {
				if isNumber(in[2]) {
					v, _ := strconv.Atoi(in[2])
					reg[getReg(in[1])] += v
				} else {
					reg[getReg(in[1])] += reg[getReg(in[2])]
				}
			} else if in[0] == "mul" {
				if isNumber(in[2]) {
					v, _ := strconv.Atoi(in[2])
					reg[getReg(in[1])] *= v
				} else {
					reg[getReg(in[1])] *= reg[getReg(in[2])]
				}
			} else if in[0] == "div" {
				if isNumber(in[2]) {
					v, _ := strconv.Atoi(in[2])
					reg[getReg(in[1])] /= v
				} else {
					reg[getReg(in[1])] /= reg[getReg(in[2])]
				}
			} else if in[0] == "mod" {
				if isNumber(in[2]) {
					v, _ := strconv.Atoi(in[2])
					reg[getReg(in[1])] %= v
				} else {
					reg[getReg(in[1])] %= reg[getReg(in[2])]
				}
			} else if in[0] == "eql" {
				if isNumber(in[2]) {
					v, _ := strconv.Atoi(in[2])
					if reg[getReg(in[1])] == v {
						reg[getReg(in[1])] = 1
					} else {
						reg[getReg(in[1])] = 0
					}
				} else {
					if reg[getReg(in[1])] == reg[getReg(in[2])] {
						reg[getReg(in[1])] = 1
					} else {
						reg[getReg(in[1])] = 0
					}
				}
			}

		}
		z := reg[3]
		ans := find(dp, blocks, block+1, z)
		if ans != -1 {
			dp[answer{block: block, z: z}] = (ans*10 + digit)
			return (ans*10 + digit)
		}
	}
	return -1
}

type answer struct {
	block int
	z     int
}

func isNumber(s string) bool {
	return unicode.IsNumber(rune(s[0]))
}

type instruction []string
type block []instruction

func parseInput() []block {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var ins []instruction
	for _, v := range si {
		if v == "" {
			continue
		}
		si := strings.Split(v, " ")
		ins = append(ins, si)
	}
	var blocks []block
	for i := 0; i < len(ins); i += 18 {
		blocks = append(blocks, ins[i+1:i+18])
	}
	fmt.Println(blocks)
	return blocks
}
