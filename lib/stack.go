package lib

import "fmt"

type Stack[T any] struct {
	val []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.val = append(s.val, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.val) == 0 {
		var v T
		return v
	}

	val := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return val
}

func (s *Stack[T]) Top() T {
	if len(s.val) == 0 {
		var v T
		return v
	}
	return s.val[len(s.val)-1]
}

func (s *Stack[T]) Empty() bool {
	return len(s.val) == 0
}

func (s *Stack[T]) Print() {
	for i := len(s.val)-1; i >=0 ; i-- {
		fmt.Print(s.val[i])
	}
	fmt.Println()
}
