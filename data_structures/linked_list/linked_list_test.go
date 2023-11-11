package linked_list_test

import (
	"scratchs/data_structures/linked_list"
	"testing"
)

func TestShouldReturnNewList(t *testing.T) {
	list := linked_list.NewList[int]()

	if list.Size != 0 {
		t.Fatalf("Incorrect size. Wanted=0, Found=%v.", list.Size)
	}

	if list.Head != nil {
		t.Fatalf("Head should be nil")
	}
}

func TestShouldAddFirstWhenEmptyList(t *testing.T) {
	value := 1
	list := linked_list.NewList[int]()

	list = linked_list.AddFirst(list, value)

	if list.Size != 1 {
		t.Fatalf("Incorrect size. Wanted=1, Found=%v.", list.Size)
	}

	if list.Head == nil {
		t.Fatalf("Head should not be nil")
	}

	if list.Head.Value != value {
		t.Fatalf("Incorrect value. Wanted=%v, Found=%v.", value, list.Head.Value)
	}
}

func TestShouldAddFirstWhenNotEmptyList(t *testing.T) {
	value := 1
	list := linked_list.NewList[int]()

	list = linked_list.AddFirst(list, 0)
	list = linked_list.AddFirst(list, value)

	if list.Size != 2 {
		t.Fatalf("Incorrect size. Wanted=2, Found=%v.", list.Size)
	}

	if list.Head == nil {
		t.Fatalf("Head should not be nil")
	}

	if list.Head.Value != value {
		t.Fatalf("Incorrect value. Wanted=%v, Found=%v.", value, list.Head.Value)
	}
}
