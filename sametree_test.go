package main

import (
	"testing"

	"github.com/jeyabalajis/goalgos/tree"
	"github.com/jeyabalajis/goalgos/treeutil"
)

func TestSameTree(t *testing.T) {
	sameTree := treeutil.Same(tree.New(1), tree.New(1))

	if !sameTree {
		t.Errorf("Expected true  got %t", sameTree)
	}
}

func TestDifferentTree(t *testing.T) {
	sameTree := treeutil.Same(tree.New(2), tree.New(1))

	if sameTree {
		t.Errorf("Expected false  got %t", sameTree)
	}
}
