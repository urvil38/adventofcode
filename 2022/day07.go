package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day07.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	dirTree := parseInput()
	totalSize := dirTree.updateSize()
	fmt.Println(dirTree.Sum(100000, 0))

	freeSpaceNeeded := 30000000
	totalSpace := 70000000
	currentlyFreeSpace := totalSpace - totalSize
	moreSpaceNeeded := freeSpaceNeeded - currentlyFreeSpace
	fmt.Println(dirTree.DirToBeDeleted(nil, moreSpaceNeeded).size)
}

type directory struct {
	parent *directory
	name   string
	size   int
	isDir  bool
	dirs   []*directory
}

func (root *directory) Add(d *directory) {
	root.dirs = append(root.dirs, d)
}

func (root *directory) updateSize() int {
	if !root.isDir {
		return root.size
	}

	dirSize := 0
	for _, dir := range root.dirs {
		dirSize += dir.updateSize()
	}

	root.size = dirSize
	return dirSize
}

func (root *directory) Sum(thresold, totalSize int) int {
	if !root.isDir {
		return totalSize
	}

	for _, d := range root.dirs {
		totalSize = d.Sum(thresold, totalSize)
	}

	if root.size <= thresold {
		return root.size + totalSize
	}
	return totalSize
}

func (root *directory) DirToBeDeleted(candidateNode *directory, targetSize int) *directory {
	if !root.isDir {
		return candidateNode
	}

	for _, dir := range root.dirs {
		candidateNode = dir.DirToBeDeleted(candidateNode, targetSize)
	}

	if root.size < targetSize {
		return candidateNode
	}

	if candidateNode == nil || root.size <= candidateNode.size {
		return root
	}

	return candidateNode
}

func parseInput() *directory {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	root := &directory{
		parent: nil,
		name:   "/",
		isDir:  true,
		dirs:   make([]*directory, 0),
	}

	curr := root

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "$ cd") {
			d := strings.Split(line, " ")[2]
			if d == ".." {
				curr = curr.parent
			} else {
				for _, dir := range curr.dirs {
					if dir.name == d {
						curr = dir
					}
				}
			}
		} else {
			var d *directory
			if startWith(line, "dir") {
				dirName := strings.Split(line, " ")[1]
				d = &directory{
					parent: curr,
					name:   dirName,
					isDir:  true,
					dirs:   make([]*directory, 0),
				}
			} else {
				s := strings.Split(line, " ")
				fz, _ := strconv.Atoi(s[0])
				name := s[1]
				d = &directory{
					parent: curr,
					name:   name,
					size:   fz,
				}
			}
			curr.Add(d)
		}
	}
	return root
}

func startWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}
