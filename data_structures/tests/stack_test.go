package data_structures_tests

import (
	ds "assalielmehdi/scratches/data_structures"
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	var st ds.Stack[int] = ds.NewList[int]()

	st.Push(1).Push(2)

	v := 2
	exp := &v
	var expErr error = nil
	gvn, err := st.Pop()

	AssertDeepEquals(t, *gvn, *exp)
	AssertDeepEquals(t, err, expErr)

	v = 1
	gvn, err = st.Pop()

	AssertDeepEquals(t, *gvn, *exp)
	AssertDeepEquals(t, err, expErr)

	exp = nil
	expErr = errors.New("Empty")
	gvn, err = st.Pop()

	AssertDeepEquals(t, gvn, exp)
	AssertDeepEquals(t, err, expErr)
}
