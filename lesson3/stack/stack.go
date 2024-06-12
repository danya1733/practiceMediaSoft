package stack

import "fmt"

type Stack[T any] struct {
	s    []T
	size int
}

func New[T any](size int) *Stack[T] {
	return &Stack[T]{
		s:    make([]T, 0, size),
		size: size,
	}
}

func (s *Stack[T]) Push(v T) {
	if len(s.s) < s.size {
		s.s = append(s.s, v)
	} else {
		panic("Stack overflow")
	}
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.s) == 0 {
		var zero T
		return zero, false
	}
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.s) == 0 {
		var zero T
		return zero, false
	}
	return s.s[len(s.s)-1], true
}

func (s *Stack[T]) String() string {
	return fmt.Sprint(s.s)
}
