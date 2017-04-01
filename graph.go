package livesort

type graphNode struct {
	value    interface{}
	parents  map[*graphNode]struct{}
	children map[*graphNode]struct{}
}

func (g *graphNode) Delete() {
	for parent := range g.parents {
		delete(parent.children, g)
	}
	for child := range g.children {
		delete(child.parents, g)
	}
}

type graph map[*graphNode]struct{}

func (g graph) Delete(node *graphNode) {
	node.Delete()
	delete(g, node)
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
	for len(g1) > 0 {
		roots := g1.Roots()
		if len(roots) == 0 {
			return true
		}
		for _, root := range roots {
			g1.Delete(root)
		}
	}
	return false
}
