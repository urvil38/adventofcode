package main

import (
	"fmt"
	"strings"
)

type Board [11 + 4*2]Pod

type Pod uint

const (
	Empty Pod = iota
	A
	B
	C
	D
)

func (p Pod) String() string {
	return ".ABCD"[p:p+1]
}

const EmptyBoard = `
#############
#...........#
###.#.#.#.###
  #.#.#.#.#
  #########
`

func (b Board) String() string {
	var args []any
	for _, v := range b {
		args = append(args, v)
	}
	return fmt.Sprintf(strings.ReplaceAll(EmptyBoard, ".", "%v"), args...)
}

func main() {
	b := Board{11:A, 13:B, 15:C, 17:D}
	fmt.Println(b)
}
