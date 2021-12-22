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

var inputFile = flag.String("inputFile", "inputs/day22.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
	input2 := parseInput()
	p2(input2)
}

type cube struct {
	x int
	y int
	z int
}

type step struct {
	on bool
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

func p1(steps []step) {
	reactor := make(map[cube]bool)

	for _, step := range steps {
		if !inRange(step) {
			continue
		}
		for i := step.x1; i <= step.x2; i++ {
			for j := step.y1; j <= step.y2; j++ {
				for k := step.z1; k <= step.z2; k++ {
					if step.on {
						reactor[cube{x: i, y: j, z: k}] = true
					} else {
						delete(reactor, cube{x: i, y: j, z: k})
					}
				}
			}
		}
	}
	fmt.Println(len(reactor))
}

func p2(steps []step) {
	var X, Y, Z []int
	for _, s := range steps {
		X = append(X, s.x1)
		X = append(X, s.x2+1)
		Y = append(Y, s.y1)
		Y = append(Y, s.y2+1)
		Z = append(Z, s.z1)
		Z = append(Z, s.z2+1)
	}
	sort.Ints(X)
	sort.Ints(Y)
	sort.Ints(Z)
	N := len(X)
	visited := make([][][]bool, N)

	for i := range visited {
		visited[i] = make([][]bool, N)
		for j := range visited[i] {
			visited[i][j] = make([]bool, N)
		}
	}

	getIndex := func(arr []int, v int) int {
		return lowerBound(arr, v)
	}

	for _, s := range steps {
		x1 := getIndex(X, s.x1)
		x2 := getIndex(X, s.x2+1)
		y1 := getIndex(Y, s.y1)
		y2 := getIndex(Y, s.y2+1)
		z1 := getIndex(Z, s.z1)
		z2 := getIndex(Z, s.z2+1)
		for i := x1; i < x2; i++ {
			for j := y1; j < y2; j++ {
				for k := z1; k < z2; k++ {
					visited[i][j][k] = s.on
				}
			}
		}
	}
	var ans int
	for x := 0; x < N-1; x++ {
		for y := 0; y < N-1; y++ {
			for z := 0; z < N-1; z++ {
				ans += boolToInt(visited[x][y][z]) * (X[x+1] - X[x]) * (Y[y+1] - Y[y]) * (Z[z+1] - Z[z])
			}
		}
	}
	fmt.Println(ans)
}

func inRange(s step) bool {
	return s.x1 >= -50 && s.x2 <= 50 && s.y1 >= -50 && s.y2 <= 50 && s.z1 >= -50 && s.z2 <= 50
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func lowerBound(a []int, x int) int {
	lo := 0
	hi := len(a)
	for lo < hi {
		mid := (lo + hi) >> 1
		if a[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func parseInput() []step {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	si := strings.Split(string(b), "\n")
	var input []step
	for _, v := range si {
		if v == "" {
			continue
		}
		input = append(input, parseStep(v))
	}

	return input
}

func parseStep(v string) step {
	ss := strings.Split(v, " ")
	sc := strings.Split(ss[1], ",")
	x1, x2 := parseRange(sc[0])
	y1, y2 := parseRange(sc[1])
	z1, z2 := parseRange(sc[2])

	return step{
		on: ss[0] == "on",
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
		z1: z1,
		z2: z2,
	}
}

func parseRange(v string) (int, int) {
	si := strings.Split(v[2:], "..")
	x1, _ := strconv.Atoi(si[0])
	x2, _ := strconv.Atoi(si[1])
	return x1, x2
}
