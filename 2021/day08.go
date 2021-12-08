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

var inputFile = flag.String("inputFile", "inputs/day08.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	p2(input)
}

type signalEntry struct {
	signals []string
	digits  []string
}

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
				d[sortStr(s)] = "1"
			case 3:
				l[3] = append(l[3], s)
				d[sortStr(s)] = "7"
			case 4:
				l[4] = append(l[4], s)
				d[sortStr(s)] = "4"
			case 5:
				l[5] = append(l[5], s)
			case 6:
				l[6] = append(l[6], s)
			case 7:
				l[7] = append(l[7], s)
				d[sortStr(s)] = "8"
			}
		}

		d4 := l[4][0]
		d7 := l[3][0]
		var d9 string
		for _, v := range l[6] {
			if contain(v, d4) && contain(v, d7) {
				d[sortStr(v)] = "9"
				d9 = v
			} else if contain(v, d7) {
				d[sortStr(v)] = "0"
			} else {
				d[sortStr(v)] = "6"
			}
		}

		d2 := l[2][0]
		for _, v := range l[5] {
			if !contain(d9, v) {
				d[sortStr(v)] = "2"
			} else if contain(v, d2) {
				d[sortStr(v)] = "3"
			} else {
				d[sortStr(v)] = "5"
			}
		}

		var digit string
		for _, ds := range se.digits {
			digit += d[sortStr(ds)]
		}

		si, _ := strconv.Atoi(digit)
		count += si
	}

	fmt.Println(count)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortStr(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
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
