package main

import (
	"testing"

	"github.com/jeyabalajis/goalgos/graphutil"
)

func TestMinCostDfs(t *testing.T) {

	twoNode := graphutil.Node{Value: 2}
	oneNode := graphutil.Node{
		Value: 1,
		Children: []graphutil.Edge{
			graphutil.Edge{Cost: 100, Node: twoNode},
		},
	}

	zeroNode := graphutil.Node{
		Value: 0,
		Children: []graphutil.Edge{
			graphutil.Edge{Cost: 100, Node: oneNode},
			graphutil.Edge{Cost: 500, Node: twoNode},
		},
	}

	minCost := graphutil.MinCostRouteDfs(zeroNode, 0, 2, 1)
	if minCost != 200 {
		t.Errorf("Expected 200 but got %d", minCost)
	}
}

func TestMinCostBfs(t *testing.T) {

	eNode := graphutil.Node{Value: 5}
	dNode := graphutil.Node{
		Value: 4,
		Children: []graphutil.Edge{
			graphutil.Edge{Cost: 4, Node: eNode},
		},
	}

	bNode := graphutil.Node{
		Value: 2,
		Children: []graphutil.Edge{
			graphutil.Edge{Cost: 4, Node: eNode},
		},
	}

	cNode := graphutil.Node{
		Value: 3,
		Children: []graphutil.Edge{
			graphutil.Edge{Cost: 2, Node: bNode},
			graphutil.Edge{Cost: 4, Node: dNode},
		},
	}

	aNode := graphutil.Node{
		Value: 1,
		Children: []graphutil.Edge{
			graphutil.Edge{Cost: 4, Node: bNode},
			graphutil.Edge{Cost: 1, Node: cNode},
		},
	}

	minCost := graphutil.MinCostRouteBfs(aNode, 1, 5, 2)

	if minCost != 7 {
		t.Errorf("Expected 7 but got %d", minCost)
	}
}
