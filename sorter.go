package livesort

import "math/rand"

// A Comparison represents the fact that Lesser is less
// than Greater.
type Comparison struct {
	Lesser  interface{}
	Greater interface{}
}

// A Sorter maintains the state of a sort.
//
// You should never add a comparison by modifying the
// Comparisons field, since doing so might result in a
// cycle.
// Instead, use the Add() method.
//
// A Sorter only works for elements that can be used as
// map keys.
// So, for example, slices cannot be used.
type Sorter struct {
	Comparisons []*Comparison
	Elements    []interface{}
}

// Request returns a random pair of elements that the
// Sorter cannot relate using the existing comparisons.
//
// If no comparisons are needed (i.e. the list can be
// fully sorted), this returns (nil, nil).
func (s *Sorter) Request() (interface{}, interface{}) {
	g := s.graph()
	roots := g.Roots()
	for len(g) > 0 {
		if len(roots) == 0 {
			return nil, nil
		} else if len(roots) == 1 {
			roots = g.Delete(roots[0])
		} else {
			return randomValuePair(roots)
		}
	}
	return nil, nil
}

// Add adds a comparison to the Sorter.
//
// This returns false if the comparison was rejected,
// which may happen if it would result in cycles.
//
// A comparison will never be rejected if each comparison
// is performed due to a corresponding Request() call.
// In other words, if you always Request() a pair, do a
// comparison, and then Add() the result, no comparisons
// can result in cycles.
func (s *Sorter) Add(lesser, greater interface{}) bool {
	s.Comparisons = append(s.Comparisons, &Comparison{Lesser: lesser, Greater: greater})
	if s.graph().Cyclic() {
		s.Comparisons = s.Comparisons[:len(s.Comparisons)-1]
		return false
	}
	return true
}

// Sort returns a sorted version of s.Elements based on
// the current comparisons.
// If the sort is incomplete, this will return one valid
// (non-deterministic) ordering.
func (s *Sorter) Sort() []interface{} {
	var res []interface{}
	g := s.graph()
	roots := g.Roots()
	for len(g) > 0 {
		if len(roots) == 0 {
			panic("cycle detected")
		}
		var newRoots []*graphNode
		for _, node := range roots {
			res = append(res, node.value)
			newRoots = append(newRoots, g.Delete(node)...)
		}
		roots = newRoots
	}
	return res
}

func (s *Sorter) graph() graph {
	g := graph{}
	entToNode := map[interface{}]*graphNode{}
	for _, value := range s.Elements {
		node := &graphNode{
			value:    value,
			parents:  map[*graphNode]struct{}{},
			children: map[*graphNode]struct{}{},
		}
		g[node] = struct{}{}
		entToNode[value] = node
	}
	for _, comp := range s.Comparisons {
		parent := entToNode[comp.Lesser]
		child := entToNode[comp.Greater]
		if parent == nil || child == nil {
			continue
		}
		parent.children[child] = struct{}{}
		child.parents[parent] = struct{}{}
	}
	return g
}

func randomValuePair(roots []*graphNode) (interface{}, interface{}) {
	idx := rand.Intn(len(roots))
	r1 := roots[idx]
	roots[idx] = roots[len(roots)-1]
	roots = roots[:len(roots)-1]
	r2 := roots[rand.Intn(len(roots))]
	return r1.value, r2.value
}
