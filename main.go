package main

import (
	"fmt"

	"github.com/jeyabalajis/goalgos/heaputil"
)

func main() {

	myCourses := []heaputil.ExternalItem{
		heaputil.ExternalItem{Value: "1", Duration: 100, EndingTime: 200},
		heaputil.ExternalItem{Value: "2", Duration: 200, EndingTime: 1300},
		heaputil.ExternalItem{Value: "3", Duration: 1000, EndingTime: 1250},
		heaputil.ExternalItem{Value: "4", Duration: 2000, EndingTime: 3200},
	}

	var scheduledCourses = heaputil.ScheduleItems(myCourses)

	fmt.Println(scheduledCourses)

}
