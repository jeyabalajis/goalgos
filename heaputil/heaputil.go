package heaputil

import (
	"container/heap"
)

// ExternalItem is sent from the caller for scheduling
type ExternalItem struct {
	Value      string // The value of the item; arbitrary.
	Duration   int
	EndingTime int // The priority of the item in the queue.
}

// ScheduledItem is the scheduled item post processing
type ScheduledItem struct {
	Value    string
	Start    int
	Duration int
	End      int
}

// ScheduleItems provides a schedule of maximum items that can be taken contiguously
// The "item" can be a course or an event
func ScheduleItems(items []ExternalItem) []ScheduledItem {
	/*
		Only one item (course or program or event) is allowed at a time
		The items must be contiguous, i.e. there cannot be a gap between one item and the next
	*/
	pq := make(PriorityQueue, len(items))
	i := 0
	for _, item := range items {
		var internalItem = Item{value: item.Value, priority: item.EndingTime, duration: item.Duration}
		pq[i] = &internalItem
		i++
	}
	heap.Init(&pq)

	var scheduledItems []ScheduledItem
	var start = 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if start+item.duration <= item.priority {
			scheduledItems = append(
				scheduledItems,
				ScheduledItem{Value: item.value, Start: start, Duration: item.duration, End: start + item.duration},
			)
			start += item.duration + 1
		}
	}
	return scheduledItems
}
