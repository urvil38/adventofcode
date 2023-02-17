package lib

import "container/heap"

type Work[T comparable] struct {
	heap []heapEntry[T]
	pos  map[T]int
	prev map[T]T
}

type heapEntry[T comparable] struct {
	b    T
	cost int
}

func (w *Work[T]) Add(prev, curr T, c int) {
	if i, ok := w.pos[curr]; ok {
		if i < 0 || w.heap[i].cost <= c {
			return
		}
		w.heap[i].cost = c
		heap.Fix((*byCost[T])(w), i)
	} else {
		heap.Push((*byCost[T])(w), heapEntry[T]{curr, c})
	}
	w.prev[curr] = prev
}

func (w *Work[T]) Next() (T, int) {
	e := heap.Pop((*byCost[T])(w)).(heapEntry[T])
	return e.b, e.cost
}

func (w *Work[T]) Empty() bool {
	return len(w.heap) == 0
}

func (w *Work[T]) Path(b T) []T {
	prev := w.prev[b]
	if prev == b {
		return []T{b}
	}
	return append(w.Path(prev), b)
}

func NewWorkQueue[T comparable]() *Work[T] {
	return &Work[T]{
		pos: make(map[T]int),
		prev: make(map[T]T),
	}
}

type byCost[T comparable] Work[T]

func (w *byCost[T]) Len() int { return len(w.heap) }

func (w *byCost[T]) Less(i, j int) bool { return w.heap[i].cost < w.heap[j].cost }

func (w *byCost[T]) fix(i int) {
	w.pos[w.heap[i].b] = i
}

func (w *byCost[T]) Swap(i, j int) {
	w.heap[i], w.heap[j] = w.heap[j], w.heap[i]
	w.fix(i)
	w.fix(j)
}

func (w *byCost[T]) Push(x any) {
	w.heap = append(w.heap, x.(heapEntry[T]))
	w.fix(len(w.heap)-1)
}

func (w *byCost[T]) Pop() any {
	x := w.heap[len(w.heap)-1]
	w.heap = w.heap[:len(w.heap)-1]
	w.pos[x.b] = -1
	return x
}
