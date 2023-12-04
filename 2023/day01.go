package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day01.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += findNumber(line)
	}
	return sum
}

func p2(lines []string) int {
	sum := 0

	for _, line := range lines {
		line = convertSplledDigitsToDigit(line)
		sum += findNumber(line)
	}

	return sum
}

func findNumber(line string) int {
	first := ""
	last := ""

	for i := 0; i < len(line); i++ {
		if line[i] >= '1' && line[i] <= '9' {
			first = string(line[i])
			break
		}
	}

	for j := len(line) - 1; j >= 0; j-- {
		if line[j] >= '1' && line[j] <= '9' {
			last = string(line[j])
			break
		}
	}

	num, _ := strconv.Atoi(first + last)
	return num
}

func convertSplledDigitsToDigit(s string) string {
	re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine")
	digits := map[string]string{
		"one":   "on1e",
		"two":   "tw2o",
		"three": "thr3ee",
		"four":  "fo4ur",
		"five":  "fi5ve",
		"six":   "si6x",
		"seven": "sev7en",
		"eight": "eig8ht",
		"nine":  "ni9ne",
	}
	for {
		match := re.FindStringIndex(s)
		if match == nil {
			break
		}
		s = s[:match[0]] + digits[s[match[0]:match[1]]] + s[match[1]:]
	}
	return s
}

func parseInput() []string {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	return si
}