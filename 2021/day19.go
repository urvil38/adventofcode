package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day19.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	input := parseInput()
	p1(input)
}

type Point [3]int

type Final map[Point]bool

type PDir struct {
	scanner int
	dir     int
}

func (p Point) Add(p1 Point) Point {
	return Point{
		p[0] + p1[0],
		p[1] + p1[1],
		p[2] + p1[2],
	}
}

func (p Point) Sub(p1 Point) Point {
	return Point{
		p[0] - p1[0],
		p[1] - p1[1],
		p[2] - p1[2],
	}

}
func p1(beacons [][]Point) {
	N := len(beacons)
	final := make(Final)
	for _, b := range beacons[0] {
		final[b] = true
	}
	pos := make([]Point, N)
	pos[0] = Point{0, 0, 0}

	good := make(map[int]bool)
	good[0] = true

	bad := make(map[int]bool)
	for i := 1; i < N; i++ {
		bad[i] = true
	}

	perms := permutations([]int{0, 1, 2})
	badAdj := make(map[PDir][]Point)
	for s := 0; s < N; s++ {
		for dir := 0; dir < 48; dir++ {
			for _, beacon := range beacons[s] {
				pdir := PDir{scanner: s, dir: dir}
				badAdj[pdir] = append(badAdj[pdir], adjust(beacon, dir, perms))
			}
		}
	}

	for len(bad) > 0 {
		found := -1
		for b, _ := range bad {

			var gScan []Point
			for v, _ := range final {
				gScan = append(gScan, v)
			}

			for bDir := 0; bDir < 48; bDir++ {
				bScan := badAdj[PDir{scanner: b, dir: bDir}]
				vote := make(map[Point]int, len(bScan)*len(gScan))

				for _, bv := range bScan {
					for _, gv := range gScan {
						// b[0] + dx = g[0] => dx = g[0] - b[0]
						np := gv.Sub(bv)
						vote[np] += 1
					}
				}

				for p, v := range vote {
					if v >= 12 {
						pos[b] = p
						for _, be := range bScan {
							final[be.Add(p)] = true
						}
						found = b
						goto end
					}
				}
			}
		}
	end:
		delete(bad, found)
		good[found] = true
	}
	fmt.Println(len(final))

	max := -1
	for i, p1 := range pos {
		for j, p2 := range pos {
			if i != j {
				manhattanDis := abs(p1[0]-p2[0]) + abs(p1[1]-p2[1]) + abs(p1[2]-p2[2])
				if manhattanDis > max {
					max = manhattanDis
				}
			}
		}
	}
	fmt.Println(max)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func adjust(beacon Point, dir int, perms [][]int) Point {
	ret := beacon
	for i, v := range perms {
		if dir>>3 == i {
			ret = Point{ret[v[0]], ret[v[1]], ret[v[2]]}
		}
	}
	if dir&1 == 1 {
		ret[0] *= -1
	}
	if dir>>1&1 == 1 {
		ret[1] *= -1
	}
	if dir>>2&1 == 1 {
		ret[2] *= -1
	}
	return ret
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					arr[i], arr[n-1] = arr[n-1], arr[i]
				} else {
					arr[0], arr[n-1] = arr[n-1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func parseInput() [][]Point {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanners := strings.Split(string(b), "\n\n")
	var B [][]Point
	for _, s := range scanners {
		var beacons []Point
		for _, line := range strings.Split(s, "\n") {
			if strings.HasPrefix(line, "--") || line == "" {
				continue
			}
			beacons = append(beacons, parsePoint(line))
		}
		B = append(B, beacons)
	}

	return B
}

func parsePoint(s string) Point {
	si := strings.Split(s, ",")
	var p Point
	for i, vi := range si {
		v, _ := strconv.Atoi(vi)
		p[i] = v
	}
	return p
}
