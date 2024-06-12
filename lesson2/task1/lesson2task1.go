package main

import "fmt"

// Стек

type stack struct {
	s    []any
	size int
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, 0, size),
		size: size,
	}
}

func (s *stack) push(v any) {
	if len(s.s) < s.size {
		s.s = append(s.s, v)
	} else {
		fmt.Println("Стек переполнен")
	}
}

func (s *stack) pop() any {
	if len(s.s) == 0 {
		return nil
	}
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}

func (s *stack) peek() any {
	if len(s.s) == 0 {
		return nil
	}
	return s.s[len(s.s)-1]
}

// Очередь

type queue struct {
	s    []any
	size int
}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, 0, size),
		size: size,
	}
}

func (q *queue) push(v any) {
	if len(q.s) < q.size {
		q.s = append(q.s, v)
	} else {
		fmt.Println("Очередь переполнена")
	}
}

func (q *queue) pop() any {
	if len(q.s) == 0 {
		return nil
	}
	v := q.s[0]
	q.s = q.s[1:]
	return v
}

// Односвязанный список

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

func (l *singlyLinkedList) add(v any) {
	newItem := &item{v, nil}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

func (l *singlyLinkedList) get(idx int) any {
	if idx < 0 || idx >= l.size {
		return nil
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

func (l *singlyLinkedList) remove(idx int) {
	if idx < 0 || idx >= l.size {
		return
	}
	if idx == 0 {
		l.first = l.first.next
		if l.first == nil {
			l.last = nil
		}
	} else {
		prev := l.first
		for i := 0; i < idx-1; i++ {
			prev = prev.next
		}
		prev.next = prev.next.next
		if prev.next == nil {
			l.last = prev
		}
	}
	l.size--
}

func (l *singlyLinkedList) values() []any {
	values := make([]any, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		values[i] = current.v
		current = current.next
	}
	return values
}

// Демонстрация

func main() {
	fmt.Println("Выберите структуру данных для демонстрации:")
	fmt.Println("1. Стек")
	fmt.Println("2. Очередь")
	fmt.Println("3. Односвязанный список")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		s := newStack(5)
		s.push(1)
		s.push(2)
		s.push(3)
		fmt.Println("Стек после добавления элементов:", s.s)
		fmt.Println("Вершина стека:", s.peek())
		fmt.Println("Извлеченный элемент:", s.pop())
		fmt.Println("Стек после извлечения элемента:", s.s)
	case 2:
		q := newQueue(5)
		q.push(1)
		q.push(2)
		q.push(3)
		fmt.Println("Очередь после добавления элементов:", q.s)
		fmt.Println("Извлеченный элемент:", q.pop())
		fmt.Println("Очередь после извлечения элемента:", q.s)
	case 3:
		l := newSinglyLinkedList()
		l.add(1)
		l.add(2)
		l.add(3)
		fmt.Println("Односвязанный список после добавления элементов:", l.values())
		fmt.Println("Элемент по индексу 1:", l.get(1))
		l.remove(1)
		fmt.Println("Односвязанный список после удаления элемента по индексу 1:", l.values())
	default:
		fmt.Println("Некорректный выбор")
	}
}
