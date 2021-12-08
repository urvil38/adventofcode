package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day08.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	p2(input)
}

// 1 2 cf
// 4 4 bcdf
// 7 3 acf
// 8 7 abcdefg

type signalEntry struct {
	signals []string
	digits  []string
}

// abcdeg => 0
// bcdefg => 6
// abcdef => 9
func p1(ses []signalEntry) {
	count := 0
	for _, se := range ses {
		for _, d := range se.digits {
			if len(d) == 2 || len(d) == 3 || len(d) == 4 || len(d) == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func p2(ses []signalEntry) {
	var count int
	for _, se := range ses {
		d := make(map[string]string)
		l := make(map[int][]string)
		for _, s := range se.signals {
			switch len(s) {
			case 2:
				l[2] = append(l[2], s)
				d[s] = "1"
			case 3:
				l[3] = append(l[3], s)
				d[s] = "7"
			case 4:
				l[4] = append(l[4], s)
				d[s] = "4"
			case 5:
				l[5] = append(l[5], s)
			case 6:
				l[6] = append(l[6], s)
			case 7:
				l[7] = append(l[7], s)
				d[s] = "8"
			}
		}

		s4 := l[4][0]
		s3 := l[3][0]
		var d9, d6 string
		for _, v := range l[6] {
			if contain(v, s4) {
				d[v] = "9"
				d9 = v
			} else if contain(v, s3) {
				d[v] = "0"
			} else {
				d[v] = "6"
				d6 = v
			}
		}

		s2 := l[2][0]
		for _, v := range l[5] {
			if contain(d9, v) && contain(d6, v) {
				d[v] = "5"
			} else if contain(v, s2) {
				d[v] = "3"
			} else {
				d[v] = "2"
			}
		}

		var digit string
		for _, ds := range se.digits {
			for k, v := range d {
				if len(k) == len(ds) {
					if contain(k, ds) {
						digit += v
					}
				}
			}
		}

		si, _ := strconv.Atoi(digit)
		count += si
	}

	fmt.Println(count)
}

func contain(s, v string) bool {
	for _, c := range v {
		if !strings.Contains(s, string(c)) {
			return false
		}
	}
	return true
}

func parseInput() []signalEntry {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var input []signalEntry
	si := strings.Split(strings.TrimSpace(string(b)), "\n")
	for _, v := range si {
		if v == "" {
			continue
		}
		input = append(input, parseSignalEntry(v))
	}

	return input
}

func parseSignalEntry(s string) signalEntry {
	se := strings.Split(s, " | ")
	return signalEntry{
		signals: strings.Split(se[0], " "),
		digits:  strings.Split(se[1], " "),
	}
}

/*

  c c
a     g
a     g
  e e
e	  f
  b	b f

*/
