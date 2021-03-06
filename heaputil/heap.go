package heaputil

import (
	"container/heap"
)

// Item is something we manage in a priority queue.
type Item struct {
	Value    string // The value of the item; arbitrary.
	Priority int    // The priority of the item in the queue.
	start    int
	end      int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// if we want Pop to give us the lowest, not highest priority so we use < here.
	// If we want the highest priority to come first, we will be using > symbol here.
	return pq[i].Priority > pq[j].Priority

}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push pushes an item into priority queue. Items are ordered in a priority queue always.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop pops an item from priority queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}

// Peek returns the top item from priority queue
func (pq *PriorityQueue) Peek() interface{} {
	curr := *pq
	n := len(curr)
	if n > 0 {
		return curr[0]
	}

	return nil
}
