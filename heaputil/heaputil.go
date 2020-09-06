package heaputil

import (
	"container/heap"
	"sort"
)

// ExternalItem is sent from the caller for scheduling.
type ExternalItem struct {
	Value      string // The value of the item; arbitrary.
	Duration   int    // The priority of the item in the Max Heap.
	EndingTime int    // Used for sorting items by lowest ending item
}

// ByEndingTime implements sort.Interface based on the EndingTime
type ByEndingTime []ExternalItem

func (a ByEndingTime) Len() int           { return len(a) }
func (a ByEndingTime) Less(i, j int) bool { return a[i].EndingTime < a[j].EndingTime }
func (a ByEndingTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ScheduledItem is the scheduled item post processing
type ScheduledItem struct {
	Value    string
	Start    int
	Duration int
	End      int
}

// ByEndTime implements sort.Interface based on the EndTime
type ByEndTime []ScheduledItem

func (a ByEndTime) Len() int           { return len(a) }
func (a ByEndTime) Less(i, j int) bool { return a[i].End < a[j].End }
func (a ByEndTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ScheduleItems provides a schedule of maximum items that can be taken contiguously
// The "item" can be a course or an event
func ScheduleItems(items []ExternalItem) []ScheduledItem {
	/*
		Only one item (course or program or event) is allowed at a time
		The items must be contiguous, i.e. there cannot be a gap between one item and the next
	*/

	// First sort the items to be scheduled in ascending order of end duration
	sort.Sort(ByEndingTime(items))

	// Now go through the courses and put them into (Max) Heap.
	// If a course cannot fit, see if it can by popping the max item from the heap
	pq := &PriorityQueue{}
	heap.Init(pq)
	var start = 0
	for _, item := range items {

		var internalItem = Item{Value: item.Value, Priority: item.Duration, start: start}
		if start+item.Duration <= item.EndingTime {
			internalItem.end = start + item.Duration
			heap.Push(pq, &internalItem)
			start = internalItem.end + 1
		} else {
			maxHeapTop, _ := heap.Pop(pq).(*Item)
			if maxHeapTop.Priority > item.Duration && maxHeapTop.start+item.Duration <= item.EndingTime {
				internalItem.start = maxHeapTop.start
				internalItem.end = maxHeapTop.start + item.Duration

				heap.Push(pq, &internalItem)
				start = internalItem.end + 1
			} else {
				heap.Push(pq, &maxHeapTop)
			}
		}
	}

	var scheduledItems []ScheduledItem
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		scheduledItems = append(
			scheduledItems,
			ScheduledItem{Value: item.Value, Start: item.start, Duration: item.Priority, End: item.end},
		)
	}
	sort.Sort(ByEndTime(scheduledItems))
	return scheduledItems
}
