package lib

import "container/heap"

type Work[State comparable] struct {
	heap []heapEntry[State]
	pos  map[State]int
	prev map[State]State
}

type heapEntry[State comparable] struct {
	b    State
	cost int
}

func (w *Work[State]) Add(prev, b State, c int) {
	if i, ok := w.pos[b]; ok {
		if i < 0 || w.heap[i].cost <= c {
			return
		}
		w.heap[i].cost = c
		heap.Fix((*byCost[State])(w), i)
	} else {
		heap.Push((*byCost[State])(w), heapEntry[State]{b, c})
	}
	w.prev[b] = prev
}

func (w *Work[State]) Next() (State, int) {
	e := heap.Pop((*byCost[State])(w)).(heapEntry[State])
	return e.b, e.cost
}

func (w *Work[State]) Empty() bool {
	return len(w.heap) == 0
}

func (w *Work[State]) Path(b State) []State {
	prev := w.prev[b]
	if prev == b {
		return []State{b}
	}
	return append(w.Path(prev), b)
}

func NewWorkQueue[State comparable]() *Work[State] {
	return &Work[State]{
		pos: make(map[State]int),
		prev: make(map[State]State),
	}
}

type byCost[State comparable] Work[State]

func (w *byCost[State]) Len() int { return len(w.heap) }

func (w *byCost[State]) Less(i, j int) bool { return w.heap[i].cost < w.heap[j].cost }

func (w *byCost[State]) fix(i int) {
	w.pos[w.heap[i].b] = i
}

func (w *byCost[State]) Swap(i, j int) {
	w.heap[i], w.heap[j] = w.heap[j], w.heap[i]
	w.fix(i)
	w.fix(j)
}

func (w *byCost[State]) Push(x any) {
	w.heap = append(w.heap, x.(heapEntry[State]))
	w.fix(len(w.heap)-1)
}

func (w *byCost[State]) Pop() any {
	x := w.heap[len(w.heap)-1]
	w.heap = w.heap[:len(w.heap)-1]
	w.pos[x.b] = -1
	return x
}