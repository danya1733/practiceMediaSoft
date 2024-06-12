package queue

import "fmt"

type Queue[T any] struct {
	s    []T
	size int
}

func New[T any](size int) *Queue[T] {
	return &Queue[T]{
		s:    make([]T, 0, size),
		size: size,
	}
}

func (q *Queue[T]) Push(v T) {
	if len(q.s) < q.size {
		q.s = append(q.s, v)
	} else {
		panic("Queue overflow")
	}
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.s) == 0 {
		var zero T
		return zero, false
	}
	v := q.s[0]
	q.s = q.s[1:]
	return v, true
}

func (q *Queue[T]) String() string {
	return fmt.Sprint(q.s)
}
