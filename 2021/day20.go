package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day20.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	algo, image := parseInput()
	p1(algo, image, 2)
	p1(algo, image, 50)
}

type Pixel struct {
	x int
	y int
}

func print(pixels map[Pixel]bool) {
	var maxX, maxY, minX, minY int
	for p, _ := range pixels {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if pixels[Pixel{x: x, y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}

func p1(algo string, image Image, steps int) {
	pixels := make(map[Pixel]bool)
	for r, row := range image {
		for c, col := range row {
			if col == '#' {
				pixels[Pixel{x: c, y: r}] = true
			}
		}
	}
	for i := 0; i < steps; i++ {
		pixels = enhance(pixels, i%2 == 1, algo)
	}
	fmt.Println(len(pixels))
}

func enhance(pixels map[Pixel]bool, outsideOn bool, algo string) map[Pixel]bool {
	newPixels := make(map[Pixel]bool)

	var minX, minY, maxX, maxY int
	dir := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

	for p, _ := range pixels {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}

	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			var s string
			for _, d := range dir {
				dx := x + d[0]
				dy := y + d[1]

				if pixels[Pixel{x: dx, y: dy}] {
					s += "1"
				} else if outsideOn && (dx > maxX || dx < minX || dy > maxY || dy < minY) {
					s += "1"
				} else {
					s += "0"
				}
			}
			if binaryToIndexValue(algo, s) == "#" {
				newPixels[Pixel{x, y}] = true
			}
		}
	}
	return newPixels
}

func binaryToIndexValue(algo, s string) string {
	v, _ := strconv.ParseInt(s, 2, 64)
	vs := string(algo[int(v)])
	return vs
}

type Image []string

func parseInput() (string, Image) {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var algo string
	si := strings.Split(string(b), "\n")
	algo = si[0]
	var image Image
	for _, row := range si[1:] {
		if row == "" {
			continue
		}
		image = append(image, row)
	}

	return algo, image
}
