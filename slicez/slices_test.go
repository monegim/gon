package slicez

import "testing"

func TestEqualFunc(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 0, 3}
	res := EqualFunc(s1, s2, func(a,b int) bool {return a==b })
	t.Log(res)
}
