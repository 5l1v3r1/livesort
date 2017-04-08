package livesort

import (
	"math/rand"
	"reflect"
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

func TestSortRandom(t *testing.T) {
	// Create a list of integers, which is almost sorted.
	// The edges connecting 0 and 9 are deleted, so that they
	// can be anywhere in the sorted sequence.

	sorter := &Sorter{}
	for i := 0; i < 10; i++ {
		sorter.Elements = append(sorter.Elements, i)
		if i > 0 {
			sorter.Comparisons = append(sorter.Comparisons, &Comparison{
				Lesser:  i - 1,
				Greater: i,
			})
		}
	}

	// Remove first and last edge.
	sorter.Comparisons = sorter.Comparisons[1 : len(sorter.Comparisons)-1]

	sortResults := [][]interface{}{}
	addResult := func(res []interface{}) {
		for _, x := range sortResults {
			if reflect.DeepEqual(x, res) {
				return
			}
		}
		sortResults = append(sortResults, res)
	}

	for i := 0; i < 50000; i++ {
		addResult(sorter.Sort())
	}

	// Two nodes can be anywhere, so (10 choose 2) * 2!.
	expectedNum := 10 * 9

	if len(sortResults) != expectedNum {
		t.Errorf("expected %d results but got %d", expectedNum, len(sortResults))
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
