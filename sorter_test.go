package livesort

import (
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	initial := make([]interface{}, 100)
	for i, j := range rand.Perm(len(initial)) {
		initial[i] = j
	}
	sorter := &Sorter{Elements: initial}
	for {
		n1, n2 := sorter.Request()
		if n1 == nil {
			break
		}
		var res bool
		if n1.(int) < n2.(int) {
			res = sorter.Add(n1, n2)
		} else {
			res = sorter.Add(n2, n1)
		}
		if !res {
			t.Fatal("comparison rejected")
		}
	}
	sorted := sorter.Sort()
	if len(sorted) != len(initial) {
		t.Errorf("bad output length: %d (expected %d)", len(sorted), len(initial))
	}
	for i, x := range sorted {
		if i != x {
			t.Errorf("bad element %d: %v", i, x)
		}
	}
}

func TestReject(t *testing.T) {
	sorter := &Sorter{Elements: []interface{}{0, 1, 2}}
	sorter.Add(0, 1)
	if sorter.Add(1, 0) {
		t.Fatal("missed basic cycle")
	}
	sorter.Add(1, 2)
	if sorter.Add(2, 0) {
		t.Fatal("missed three-cycle")
	}
	if !sorter.Add(0, 2) {
		t.Fatal("wrongful rejection")
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		initial := make([]interface{}, 100)
		for i, j := range rand.Perm(len(initial)) {
			initial[i] = j
		}
		sorter := &Sorter{Elements: initial}
		for {
			n1, n2 := sorter.Request()
			if n1 == nil {
				break
			}
			var res bool
			if n1.(int) < n2.(int) {
				res = sorter.Add(n1, n2)
			} else {
				res = sorter.Add(n2, n1)
			}
			if !res {
				b.Fatal("comparison rejected")
			}
		}
	}
}
