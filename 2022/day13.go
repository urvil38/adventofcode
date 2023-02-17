package main

import (
	"adventofcode/lib"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

var inputFile = flag.String("inputFile", "inputs/day13.input", "Relative file path to use as input.")

func main() {
	flag.Parse()
	signals := parseInput()
	p(signals)
}

func p(signals []*signal) {
	var n int
	var j int
	for i := 0; i < len(signals); i += 2 {
		s1 := signals[i]
		s2 := signals[i+1]
		if cmp(*s1, *s2) <= 0 {
			n += j + 1
		}
		j++
	}
	fmt.Println(n)
	d1, d2 := d(2), d(6)
	signals = append(signals, d1, d2)
	slices.SortFunc(signals, func(a, b *signal) bool { return cmp(*a, *b) < 0 })
	i1 := slices.Index(signals, d1) + 1
	i2 := slices.Index(signals, d2) + 1
	fmt.Println(i1 * i2)
}

type signal struct {
	val    int
	childs []*signal
}

func parseSignal(s string) *signal {
	stack := lib.NewStack[*signal]()
	stack.Push(&signal{})
	for len(s) > 0 {
		switch s[0] {
		case '[':
			sig := signal{childs: []*signal{}}
			stack.Push(&sig)
			s = s[1:]
		case ',':
			s = s[1:]
		case ']':
			top := stack.Pop()
			parent := stack.Pop()
			parent.childs = append(parent.childs, top)
			stack.Push(parent)
			s = s[1:]
		default:
			i := strings.IndexAny(s, "[],")
			if i < 0 {
				i = len(s)
			}
			n, _ := strconv.Atoi(s[:i])
			t := stack.Pop()
			t.childs = append(t.childs, &signal{val: n})
			stack.Push(t)
			s = s[i:]
		}
	}
	top := stack.Pop()
	return top.childs[0]
}

func cmp(l, r signal) (v int) {
	if l.childs == nil && r.childs == nil {
		return Cmp(l.val, r.val)
	}
	if l.childs != nil && r.childs != nil {
		i := 0
		for ; i < len(l.childs) && i < len(r.childs); i++ {
			switch cmp(*l.childs[i], *r.childs[i]) {
			case -1:
				return -1
			case 1:
				return 1
			}
		}
		if i < len(l.childs) {
			return 1
		}
		if i < len(r.childs) {
			return -1
		}
		return 0
	}
	if l.childs != nil {
		return cmp(l, signal{childs: []*signal{{val: r.val}}})
	}
	return cmp(signal{childs: []*signal{{val: l.val}}}, r)
}

func (t signal) String() string {
	if t.childs == nil {
		return strconv.Itoa(t.val)
	}
	var parts []string
	for _, c := range t.childs {
		parts = append(parts, c.String())
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func parseInput() []*signal {
	b, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	b = append(b, '\n')
	var signals []*signal
	si := strings.Split(string(b), "\n")
	for _, line := range si {
		if line == "" {
			continue
		}
		signals = append(signals, parseSignal(line))
	}
	return signals
}

type TotallyOrdered interface {
	constraints.Integer | ~string
}

func Cmp[T TotallyOrdered](a, b T) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func d(i int) *signal {
	return &signal{childs: []*signal{
		{childs: []*signal{
			{val: i},
		}},
	}}
}
