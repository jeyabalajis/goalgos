package main

import (
	"jeyabalajis/goalgos/tree"
	"jeyabalajis/goalgos/treeutil"
	"testing"
)

func TestSameTree(t *testing.T) {
	sameTree := treeutil.Same(tree.New(1), tree.New(1))

	if !sameTree {
		t.Errorf("Expected true  got %t", sameTree)
	}
}
