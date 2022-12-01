package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
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

func p1(calories [][]int) int {
	maxCalories := math.MinInt

	for _, cal := range calories {
		var total int
		for _, elfCal := range cal {
			total += elfCal
		}
		if total > maxCalories {
			maxCalories = total
		}
	}
	return maxCalories
}

func p2(calories [][]int) int {
	var aggCalories []int

	for _, cal := range calories {
		var total int
		for _, elfCal := range cal {
			total += elfCal
		}
		aggCalories = append(aggCalories, total)
	}

	sort.Slice(aggCalories, func(i, j int) bool {
		return aggCalories[i] > aggCalories[j]
	})
	return aggCalories[0] + aggCalories[1] + aggCalories[2]
}

func parseInput() [][]int {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var calories [][]int
	var elfCalories []int
	for _, v := range si {
		if v == "" {
			calories = append(calories, elfCalories)
			elfCalories = make([]int, 0)
			continue
		}
		vi, _ := strconv.Atoi(v)
		elfCalories = append(elfCalories, vi)
	}

	return calories
}
