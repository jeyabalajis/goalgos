package queue

import (
	"container/list"
)

// Queue implements queue data structure
type Queue struct {
	queue *list.List
}

// New returns an empty list
func New() *Queue {
	return &Queue{queue: list.New()}
}

// Push pushes a value into the queue
func (myQueue *Queue) Push(value interface{}) {
	myQueue.queue.PushBack(value)
}

// Pop takes out first value out of the queue
func (myQueue *Queue) Pop() interface{} {
	e := myQueue.queue.Front() // First element
	myQueue.queue.Remove(e)    // Dequeue
	i := e.Value
	return i
}

// Empty evaluates whether the queue is empty
func (myQueue *Queue) Empty() bool {
	return (myQueue.queue.Len() == 0)
}
