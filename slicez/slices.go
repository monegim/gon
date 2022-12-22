package slicez

import "github.com/monegim/gon/compare"

func Equal[A comparable](s1, s2 []A) bool {
	return EqualFunc(s1, s2, compare.Equal[A])
}
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		if !eq(v1, s2[i]) {
			return false
		}
	}
	return true
}
