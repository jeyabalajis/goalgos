package main

import (
	"testing"

	"github.com/jeyabalajis/goalgos/queue"
)

func TestQueue(t *testing.T) {
	myQueue := queue.New()

	myQueue.Push(1)
	myQueue.Push(2)
	myQueue.Push(3)

	pop1 := myQueue.Pop()
	s, _ := pop1.(int)

	if s != 1 {
		t.Errorf("Expected 1 got %d", s)
	}

	pop2 := myQueue.Pop()
	s, _ = pop2.(int)
	if s != 2 {
		t.Errorf("Expected 2 got %d", s)
	}

	pop3 := myQueue.Pop()
	s, _ = pop3.(int)
	if s != 3 {
		t.Errorf("Expected 3 got %d", s)
	}

	if !myQueue.Empty() {
		t.Errorf("Queue must have been empty, but it is not")
	}
}
