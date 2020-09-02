package graphutil

import (
	"fmt"

	"github.com/jeyabalajis/goalgos/queue"
)

// Edge represents a line from one node to another
type Edge struct {
	Cost int
	Node Node
}

// Node is a structure to hold graphs
type Node struct {
	Value    int
	Children []Edge
	marked   bool
}

// Visit prints the current node's value
func (n Node) Visit() {
	fmt.Println(n.Value, "-")
}

func searchGraphDfs(graph Node, src int, dest int, maxStops int, stops int, currCost int) (minCost int) {

	minCost = 99999
	for _, edge := range graph.Children {
		var cost int
		if edge.Node.Value == dest && stops <= maxStops {
			cost = currCost + edge.Cost
		} else {
			cost = searchGraphDfs(edge.Node, src, dest, maxStops, stops+1, currCost+edge.Cost)
		}
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func searchGraphBfs(graph Node, src int, dest int, maxStops int) (minCost int) {
	myQueue := queue.New()

	graph.marked = true
	myQueue.Push(graph)

	for !myQueue.Empty() {
		i := myQueue.Pop()
		q := i.(Node)

		for _, edge := range q.Children {
			edge.Node.Visit()
		}
	}
	return minCost
}

// MinCostRoute finds the minimum cost between src to dest with max Stops constraint
func MinCostRoute(graph Node, src int, dest int, maxStops int) (minCost int) {

	/* Perform a depth first search on the graph
	   On each iteration, first locate source.
	   Then check if destination is reached.
	   Along the way, keep accumulating cost and also keep an eye on max Stops constrint
	*/
	return searchGraphDfs(graph, src, dest, maxStops, 0, minCost)
}
