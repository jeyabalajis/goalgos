# Schedule Courses

## Problem Statement

There are n different online courses numbered from 1 to n. Each course has some duration(course length) t and closed on dth day. A course should be taken continuously for t days and must be finished before or on the dth day. You will start at the 1st day.

Given n online courses represented by pairs (t,d), your task is to find the maximal number of courses that can be taken.

## Example:

- Input: [[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]
- Output: 3

__Explanation:__ 
There're totally 4 courses, but you can take 3 courses at most:

- First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
- Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day. 
- Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day. 
- The 4th course _cannot_ be taken now, since you will finish it on the 3300th day, which exceeds the closed date.

Difficulty: :moneybag: :moneybag: :moneybag:

[Leet Code Link](https://leetcode.com/problems/course-schedule-iii/)

## Solution

The key concept behind attacking this problem is this:

> It is always profitable to take the course with a smaller end day prior to a course with a larger end day. This is because, the course with a smaller duration, if can be taken, can surely be taken only if it is taken prior to a course with a larger end day.

So, the first thing we do is to __order courses by their end day in ascending order__.

Secondly, when we try to fit in the courses, if we encounter a course that does not fit, we can try to swap out this course
with the previous one (already put into the schedule), provided the current course duration is lesser than the one we are trying to swap out.

> For us to be able to do this, the courses being scheduled must be stored in a Max Heap (by Duration) for an easy swap.

## Source Code

```
func ScheduleItems(items []ExternalItem) []ScheduledItem {
	// First sort the items to be scheduled in ascending order of end duration
	sort.Sort(ByEndingTime(items))

	// Now go through the courses and put them into (Max) Heap.
	// If a course cannot fit, see if it can by popping the max item from the heap
	pq := &PriorityQueue{}
	heap.Init(pq)
	var start = 0
	for _, item := range items {

		var internalItem = Item{value: item.Value, priority: item.Duration, start: start}
		if start+item.Duration <= item.EndingTime {
			internalItem.end = start + item.Duration
			heap.Push(pq, &internalItem)
			start = internalItem.end + 1
		} else {
			maxHeapTop, _ := heap.Pop(pq).(*Item)
			if maxHeapTop.priority > item.Duration && maxHeapTop.start+item.Duration <= item.EndingTime {
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
			ScheduledItem{Value: item.value, Start: item.start, Duration: item.priority, End: item.end},
		)
	}
	sort.Sort(ByEndTime(scheduledItems))
	return scheduledItems
}
```