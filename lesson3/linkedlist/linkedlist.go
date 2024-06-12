package linkedlist

type Node[T any] struct {
	v    T
	next *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Add(v T) {
	newNode := &Node[T]{v: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList[T]) Get(idx int) (T, bool) {
	if idx < 0 || idx >= l.size {
		var zero T
		return zero, false
	}
	curr := l.head
	for i := 0; i < idx; i++ {
		curr = curr.next
	}
	return curr.v, true
}

func (l *LinkedList[T]) Remove(idx int) bool {
	if idx < 0 || idx >= l.size {
		return false
	}
	if idx == 0 {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
	} else {
		prev := l.head
		for i := 0; i < idx-1; i++ {
			prev = prev.next
		}
		prev.next = prev.next.next
		if prev.next == nil {
			l.tail = prev
		}
	}
	l.size--
	return true
}

func (l *LinkedList[T]) Values() []T {
	values := make([]T, l.size)
	curr := l.head
	for i := 0; i < l.size; i++ {
		values[i] = curr.v
		curr = curr.next
	}
	return values
}
