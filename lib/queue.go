package lib

type Queue[T any] struct {
	q []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(x T) T {
	q.q = append(q.q, x)
	return x
}

func (q *Queue[T]) Dequeue() T {
	if len(q.q) == 0 {
		var v T
		return v
	}

	e := q.q[0]
	q.q = q.q[1:]
	return e
}

func (q Queue[T]) Front() T {
	if len(q.q) == 0 {
		var v T
		return v
	}

	return q.q[0]
}

func (q Queue[T]) Empty() bool {
	return len(q.q) == 0
}
