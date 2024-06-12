package main

import (
	"fmt"
	"lesson3/linkedlist"
	"lesson3/queue"
	"lesson3/stack"
)

func main() {
	// Стек
	s := stack.New[int](5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Стек после добавления элементов:", s)
	if v, ok := s.Peek(); ok {
		fmt.Println("Вершина стека:", v)
	}
	if v, ok := s.Pop(); ok {
		fmt.Println("Извлеченный элемент:", v)
	}
	fmt.Println("Стек после извлечения элемента:", s)
	fmt.Println("")
	// Очередь
	q := queue.New[string](5)
	q.Push("A")
	q.Push("B")
	q.Push("C")
	fmt.Println("Очередь после добавления элементов:", q)
	if v, ok := q.Pop(); ok {
		fmt.Println("Извлеченный элемент:", v)
	}
	fmt.Println("Очередь после извлечения элемента:", q)
	fmt.Println("")
	// Односвязанный список
	l := linkedlist.New[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	fmt.Println("Односвязанный список после добавления элементов:", l.Values())
	if v, ok := l.Get(1); ok {
		fmt.Println("Элемент по индексу 1:", v)
	}
	if ok := l.Remove(1); ok {
		fmt.Println("Односвязанный список после удаления элемента по индексу 1:", l.Values())
	}
}
