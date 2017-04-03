package livesort

type graphNode struct {
	value    interface{}
	parents  map[*graphNode]struct{}
	children map[*graphNode]struct{}
}

// Delete removes the node and returns all of the newly
// exposed root nodes.
func (g *graphNode) Delete() []*graphNode {
	for parent := range g.parents {
		delete(parent.children, g)
	}
	var res []*graphNode
	for child := range g.children {
		delete(child.parents, g)
		if len(child.parents) == 0 {
			res = append(res, child)
		}
	}
	return res
}

type graph map[*graphNode]struct{}

func (g graph) Delete(node *graphNode) []*graphNode {
	res := node.Delete()
	delete(g, node)
	return res
}

func (g graph) Roots() []*graphNode {
	var res []*graphNode
	for node := range g {
		if len(node.parents) == 0 {
			res = append(res, node)
		}
	}
	return res
}

func (g graph) Cyclic() bool {
	g1 := graph{}
	for k, v := range g {
		g1[k] = v
	}
	roots := g1.Roots()
	for len(g1) > 0 {
		if len(roots) == 0 {
			return true
		}
		var newRoots []*graphNode
		for _, root := range roots {
			newRoots = append(newRoots, g1.Delete(root)...)
		}
		roots = newRoots
	}
	return false
}
