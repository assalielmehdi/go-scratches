package data_structures_tests

import (
	"reflect"
	"testing"
)

func AssertPerdicate[G, E any](t *testing.T, gvn G, exp E, prd func(G, E) bool) {
	if !prd(gvn, exp) {
		t.Fatalf("Given: %v, Expected: %v", gvn, exp)
	}
}

func AssertDeepEquals[G, E any](t *testing.T, gvn G, exp E) {
	AssertPerdicate(t, gvn, exp, func(g G, e E) bool {
		return reflect.DeepEqual(gvn, exp)
	})
}
