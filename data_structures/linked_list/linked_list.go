package linked_list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Size int
	Head *Node[T]
	Tail *Node[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{
		Size: 0,
		Head: nil,
	}
}

func NodeAtIndex[T any](list *List[T], index int) (*Node[T], bool) {
	if index < 0 || index >= list.Size {
		return nil, false
	}

	node := list.Head

	for inc := 0; inc < index; inc++ {
		node = node.Next
	}

	return node, true
}

func AddLast[T any](list *List[T], value T) *List[T] {
	node := Node[T]{
		Value: value,
		Next:  nil,
	}

	if list.Size == 0 {
		list.Head = &node
		list.Tail = &node
	} else {
		list.Tail.Next = &node
	}

	list.Size++

	return list
}

func AddFirst[T any](list *List[T], value T) *List[T] {
	node := Node[T]{
		Value: value,
		Next:  list.Head,
	}

	if list.Size == 0 {
		list.Head = &node
		list.Tail = &node
	} else {
		list.Head = &node
	}

	list.Size++

	return list
}
