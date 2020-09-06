package main

import (
	"fmt"

	"github.com/jeyabalajis/goalgos/graphutil"
)

func main() {

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

	fmt.Println(minCost)

}
