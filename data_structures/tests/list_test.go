package data_structures_tests

import (
	ds "assalielmehdi/scratches/data_structures"
	"errors"
	"testing"
)

func TestAddFirst(t *testing.T) {
	l := ds.NewList[int]()

	l.AddFirst(3).AddFirst(2).AddFirst(1)

	exp := []int{1, 2, 3}
	gvn := l.ToSlice()

	AssertDeepEquals(t, gvn, exp)
}

func TestAddLast(t *testing.T) {
	l := ds.NewList[int]()

	l.AddLast(1).AddLast(2).AddLast(3)

	exp := []int{1, 2, 3}
	gvn := l.ToSlice()

	AssertDeepEquals(t, gvn, exp)
}

func TestRemoveFirst(t *testing.T) {
	l := ds.NewList[int]()

	l.AddLast(1).AddLast(2)

	v := 1
	exp := &v
	var expErr error = nil
	gvn, err := l.RemoveFirst()

	AssertDeepEquals(t, *gvn, *exp)
	AssertDeepEquals(t, err, expErr)

	v = 2
	gvn, err = l.RemoveFirst()

	AssertDeepEquals(t, *gvn, *exp)
	AssertDeepEquals(t, err, expErr)

	exp = nil
	expErr = errors.New("Empty")
	gvn, err = l.RemoveFirst()

	AssertDeepEquals(t, gvn, exp)
	AssertDeepEquals(t, err, expErr)
}

func TestRemoveLast(t *testing.T) {
	l := ds.NewList[int]()

	l.AddLast(1).AddLast(2)

	v := 2
	exp := &v
	var expErr error = nil
	gvn, err := l.RemoveLast()

	AssertDeepEquals(t, *gvn, *exp)
	AssertDeepEquals(t, err, expErr)

	v = 1
	gvn, err = l.RemoveLast()

	AssertDeepEquals(t, *gvn, *exp)
	AssertDeepEquals(t, err, expErr)

	exp = nil
	expErr = errors.New("Empty")
	gvn, err = l.RemoveLast()

	AssertDeepEquals(t, gvn, exp)
	AssertDeepEquals(t, err, expErr)
}
