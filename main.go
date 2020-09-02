package main

import (
	"fmt"

	"github.com/jeyabalajis/goalgos/graphutil"
)

func main() {

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

	minCost := graphutil.MinCostRoute(zeroNode, 0, 2, 1)
	fmt.Println(minCost)

}
