package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var inputFile = flag.String("inputFile", "inputs/day18.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	p1()
	p2()
}

func p1() {
	pairs := parseInput()
	sum := pairs[0]
	for i, p := range pairs {
		if i == 0 {
			continue
		}
		sum = Sum(sum, p)
		sum.reduce()
	}
	fmt.Println(sum.magnitude())
}

func p2() {
	var max int
	pairs := parseInput()
	for i := 0; i < len(pairs); i++ {
		for j := 0; j < len(pairs); j++ {
			if i != j {
				pairs = parseInput()
				newPair := Sum(pairs[i], pairs[j])
				newPair.reduce()
				mag := newPair.magnitude()
				if mag > max {
					max = mag
				}
			}
		}
	}
	fmt.Println(max)
}

type Pair struct {
	left, right           *int
	leftP, rightP, parent *Pair
	depth                 int
}

func Sum(p1, p2 *Pair) *Pair {
	newPair := &Pair{
		left:   nil,
		right:  nil,
		leftP:  p1.expandDepth(),
		rightP: p2.expandDepth(),
		parent: nil,
	}
	newPair.leftP.parent = newPair
	newPair.rightP.parent = newPair

	return newPair
}

func (p Pair) magnitude() int {
	sum := 0
	if p.left != nil {
		sum += 3 * *p.left
	} else if p.leftP != nil {
		sum += 3 * p.leftP.magnitude()
	}

	if p.right != nil {
		sum += 2 * *p.right
	} else if p.rightP != nil {
		sum += 2 * p.rightP.magnitude()
	}

	return sum
}

func (p *Pair) reduce() {
	for {
		acted := p.explode()
		if acted {
			continue
		}
		acted = p.split()
		if !acted {
			break
		}
	}
}

func (p *Pair) stashUpRight(v int) {
	curr := p.parent
	if curr == nil {
		return
	}

	if curr.leftP == p {
		if curr.right != nil {
			*curr.right += v
		} else {
			curr.rightP.stashDownRight(v)
		}
	} else if curr.rightP == p {
		curr.stashUpRight(v)
	}
}
func (p *Pair) stashUpLeft(v int) {
	curr := p.parent
	if curr == nil {
		return
	}

	if curr.rightP == p {
		if curr.left != nil {
			*curr.left += v
		} else {
			curr.leftP.stashDownLeft(v)
		}
	} else if curr.leftP == p {
		curr.stashUpLeft(v)
	}
}
func (p *Pair) stashDownRight(v int) {
	if p.left != nil {
		*p.left += v
	} else {
		p.leftP.stashDownRight(v)
	}
}
func (p *Pair) stashDownLeft(v int) {
	if p.right != nil {
		*p.right += v
	} else {
		p.rightP.stashDownLeft(v)
	}
}

func (p *Pair) explode() bool {
	if p.depth >= 4 && p.left != nil && p.right != nil {
		p.stashUpLeft(*p.left)
		p.stashUpRight(*p.right)
		// we are left child
		if p.parent.leftP == p {
			zero := 0
			p.parent.left = &zero
			p.parent.leftP = nil
		} else if p.parent.rightP == p {
			// we are right child
			zero := 0
			p.parent.right = &zero
			p.parent.rightP = nil
		}
		return true
	}
	if p.leftP != nil {
		exploded := p.leftP.explode()
		if exploded {
			return true
		}
	}
	if p.rightP != nil {
		exploded := p.rightP.explode()
		if exploded {
			return true
		}
	}
	return false
}

func (p *Pair) split() bool {
	if p.left != nil && *p.left > 9 {
		left := *p.left / 2
		right := *p.left - left
		p.leftP = &Pair{
			left:   &left,
			right:  &right,
			parent: p,
			depth:  p.depth + 1,
		}
		p.left = nil
		return true
	}

	if p.leftP != nil {
		if ret := p.leftP.split(); ret {
			return ret
		}
	}

	if p.right != nil && *p.right > 9 {
		left := *p.right / 2
		right := *p.right - left
		p.rightP = &Pair{
			left:   &left,
			right:  &right,
			parent: p,
			depth:  p.depth + 1,
		}
		p.right = nil
		return true
	}

	if p.rightP != nil {
		if ret := p.rightP.split(); ret {
			return ret
		}
	}
	return false
}

func (p *Pair) expandDepth() *Pair {
	p.depth++
	if p.leftP != nil {
		p.leftP.expandDepth()
	}
	if p.rightP != nil {
		p.rightP.expandDepth()
	}
	return p
}

func parse(s string, depth int, parent *Pair) *Pair {
	if s[0] != '[' || s[len(s)-1] != ']' {
		fmt.Printf("invalid string: %s", s)
		return nil
	}

	var p Pair
	p.depth = depth
	p.parent = parent
	var posAfterComma int
	if s[1] == '[' {
		bktCount := 1
		for i := 2; bktCount != 0; i++ {
			switch s[i] {
			case '[':
				bktCount++
			case ']':
				bktCount--
			}
			posAfterComma = i + 2
		}
		p.leftP = parse(s[1:posAfterComma-1], depth+1, &p)
	} else {
		left := int(s[1] - '0')
		p.left = &left
		posAfterComma = 3
	}

	if s[posAfterComma] == '[' {
		p.rightP = parse(s[posAfterComma:len(s)-1], depth+1, &p)
	} else {
		right := int(s[posAfterComma] - '0')
		p.right = &right
	}
	return &p
}

func (p Pair) print() {
	fmt.Print("[")
	if p.left != nil {
		fmt.Print(*p.left)
	} else {
		p.leftP.print()
	}
	fmt.Print(",")
	if p.right != nil {
		fmt.Print(*p.right)
	} else {
		p.rightP.print()
	}
	fmt.Print("]")
}

func parseInput() []*Pair {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	var pairs []*Pair
	si := strings.Split(string(b), "\n")
	for _, s := range si {
		if s == "" {
			continue
		}
		pairs = append(pairs, parse(s, 0, nil))
	}
	return pairs
}
