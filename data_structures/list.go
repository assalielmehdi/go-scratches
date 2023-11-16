package data_structures

import (
	"errors"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type List[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{
		size: 0,
		head: nil,
	}
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) AddFirst(v T) *List[T] {
	node := Node[T]{
		value: v,
		next:  l.head,
		prev:  nil,
	}

	if l.size == 0 {
		l.head = &node
		l.tail = &node
	} else {
		l.head.prev = &node
		l.head = &node
	}

	l.size++

	return l
}

func (l *List[T]) AddLast(v T) *List[T] {
	node := Node[T]{
		value: v,
		next:  nil,
		prev:  l.tail,
	}

	if l.size == 0 {
		l.head = &node
		l.tail = &node
	} else {
		l.tail.next = &node
		l.tail = &node
	}

	l.size++

	return l
}

func (l *List[T]) RemoveFirst() (*T, error) {
	if l.size == 0 {
		return nil, errors.New("Empty")
	}

	v := l.head.value

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.head.next.prev = nil
		l.head = l.head.next
	}

	l.size--

	return &v, nil
}

func (l *List[T]) RemoveLast() (*T, error) {
	if l.size == 0 {
		return nil, errors.New("Empty")
	}

	v := l.tail.value

	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}

	l.size--

	return &v, nil
}

func (l *List[T]) ToSlice() []T {
	items := make([]T, l.size)
	cur := l.head

	for i := 0; i < l.size; i++ {
		items[i] = cur.value

		cur = cur.next
	}

	return items
}

func (l *List[T]) Push(v T) Stack[T] {
	return l.AddFirst(v)
}

func (l *List[T]) Pop() (*T, error) {
	return l.RemoveFirst()
}
