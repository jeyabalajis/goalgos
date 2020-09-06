package graphutil

import (
	"container/heap"
	"fmt"
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

// CostType is a structure to hold shortest distance from a node and the previous vertex
type CostType struct {
	minCost        int
	previousVertex int
}

// MinCostMap is a hashmap of keys that stores the lowest cost from a source node
type MinCostMap map[int]CostType

func (minCostMap MinCostMap) initKey(key int, source bool) {
	_, ok := minCostMap[key]

	if !ok {
		if source {
			minCostMap[key] = CostType{minCost: 0, previousVertex: -1}
		} else {
			minCostMap[key] = CostType{minCost: -1, previousVertex: -1}
		}
	}
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
	pq := &PriorityQueue{}

	graph.marked = true
	heap.Init(pq)

	item := Item{Value: graph.Value, Priority: 0, Node: graph, Depth: 0}
	heap.Push(pq, &item)

	minCostMap := make(MinCostMap)

	for pq.Len() > 0 {
		parentNode := heap.Pop(pq).(*Item)

		source := false

		if src == parentNode.Value {
			source = true
		}
		minCostMap.initKey(parentNode.Value, source)

		for _, edge := range parentNode.Node.Children {
			childNode := edge.Node

			source := false
			if src == childNode.Value {
				source = true
			}

			minCostMap.initKey(childNode.Value, source)
			currMinCost := minCostMap[childNode.Value].minCost
			newCost := minCostMap[parentNode.Value].minCost + edge.Cost
			item := Item{Value: childNode.Value, Priority: newCost, Node: childNode, Depth: parentNode.Depth + 1}

			// Visited first time, so update min cost and previous vertex
			if currMinCost == -1 {
				minCostMap[childNode.Value] = CostType{minCost: newCost, previousVertex: parentNode.Value}
				if item.Depth <= maxStops {
					heap.Push(pq, &item)
				}
			}

			// Update distance and previous vertex if the new cost is smaller than the old cost
			if currMinCost > newCost {
				minCostMap[childNode.Value] = CostType{minCost: newCost, previousVertex: parentNode.Value}
				pq.update(&item, item.Value, item.Priority, item.Depth)
			}
			fmt.Println(minCostMap)
		}
	}
	return minCostMap[dest].minCost
}

// MinCostRouteDfs finds the minimum cost between src to dest with max Stops constraint, using Depth First Search
func MinCostRouteDfs(graph Node, src int, dest int, maxStops int) (minCost int) {

	/* Perform a depth first search on the graph
	   On each iteration, first locate source.
	   Then check if destination is reached.
	   Along the way, keep accumulating cost and also keep an eye on max Stops constrint
	*/
	return searchGraphDfs(graph, src, dest, maxStops, 0, minCost)
}

// MinCostRouteBfs finds the minimum cost between src to dest with max Stops constraint, using Breadth First Search
func MinCostRouteBfs(graph Node, src int, dest int, maxStops int) (minCost int) {

	return searchGraphBfs(graph, src, dest, maxStops)
}
