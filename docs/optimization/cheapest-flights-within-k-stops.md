# Cheapest flights within k stops

[Problem Link](https://leetcode.com/problems/cheapest-flights-within-k-stops/)


## Solution

This is a node traversal problem. This can be solved in two ways:

- This can be a depth first search (DFS) and can be attacked with recursion and parallelism.
- This can be a breadth first search. In this case, a Hash Map is maintained as follows:
    - Against each node, the shortest distance from source & the previous node to this path is maintained

### Depth First Search
- Keep a common variable for the cheapest cost
- Traverse nodes recursively and accumulate cost.
- For each traversal, validate whether the destination is reached and if yes, push to cheapest cost store
- The cheapest cost store shall accept the path &amp; the cost only if the incoming cost is cheaper than the existing
- Once all the go routines complete, return the cheapest cost and the path

#### Source Code
```
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
```