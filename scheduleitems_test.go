package main

import (
	"testing"

	"github.com/jeyabalajis/goalgos/heaputil"
)

func TestScheduleItems(t *testing.T) {
	myCourses := []heaputil.ExternalItem{
		heaputil.ExternalItem{Value: "1", Duration: 100, EndingTime: 200},
		heaputil.ExternalItem{Value: "2", Duration: 200, EndingTime: 1300},
		heaputil.ExternalItem{Value: "3", Duration: 1000, EndingTime: 1250},
		heaputil.ExternalItem{Value: "4", Duration: 2000, EndingTime: 3200},
	}

	var scheduledCourses = heaputil.ScheduleItems(myCourses)

	for index, scheduledItem := range scheduledCourses {
		if index == 0 && scheduledItem.Value != "1" {
			t.Errorf("Expected 1 got %s", scheduledItem.Value)
		}
	}
}
