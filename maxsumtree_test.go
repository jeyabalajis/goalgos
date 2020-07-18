package main

import (
	"fmt"
	"testing"

	"github.com/jeyabalajis/goalgos/tree"
	"github.com/jeyabalajis/goalgos/treeutil"
)

func TestMaxSumTree(t *testing.T) {

	var myTree = tree.Tree{
		Value: 6,
		Left: &tree.Tree{
			Value: 4,
			Left: &tree.Tree{
				Value: -5,
				Left:  &tree.Tree{Value: 120},
				Right: &tree.Tree{Value: 1},
			},
			Right: &tree.Tree{Value: 2},
		},
		Right: &tree.Tree{
			Value: 3,
			Left:  &tree.Tree{Value: 1},
			Right: &tree.Tree{
				Value: 13,
				Left: &tree.Tree{
					Value: 145,
					Left:  &tree.Tree{Value: -1000},
					Right: &tree.Tree{Value: 200},
				},
				Right: &tree.Tree{Value: 6},
			},
		},
	}

	maxValue := treeutil.MaxSumPath(&myTree)

	fmt.Println(maxValue)
	if maxValue != 486 {
		t.Errorf("Expected 286 got %d", maxValue)
	}

	var myTreeSimple = tree.Tree{
		Value: 6,
		Left:  &tree.Tree{Value: 4},
		Right: &tree.Tree{Value: 3},
	}

	maxValueSimple := treeutil.MaxSumPath(&myTreeSimple)

	fmt.Println(maxValueSimple)
	if maxValueSimple != 13 {
		t.Errorf("Expected 13 got %d", maxValue)
	}

}
