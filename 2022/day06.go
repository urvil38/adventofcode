package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day06.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(decoder(input, 4))
	fmt.Println(decoder(input, 14))
}

func decoder(signal string, window int) int {
	for i := 0; i < len(signal)-window; i++ {
		if !isSignalRepeating(signal[i : i+window]) {
			return i + window
		}
	}
	return 0
}

func isSignalRepeating(s string) bool {
	set := make(map[string]int)
	for _, v := range s {
		set[string(v)]++
	}
	for _, v := range set {
		if v > 1 {
			return true
		}
	}
	return false
}

func parseInput() string {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var signal string

	for _, line := range lines {
		if line == "" {
			continue
		}

		return line
	}
	return signal
}
