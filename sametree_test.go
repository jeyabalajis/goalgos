package main

import (
	"testing"

	"github.com/jeyabalajis/goalgos/tree"
	"github.com/jeyabalajis/goalgos/treeutil"
)

func TestSameTree(t *testing.T) {

	tree1 := tree.Tree{Left: &tree.Tree{Value: 2}, Right: &tree.Tree{Value: 1}, Value: 1}
	tree2 := tree.Tree{Left: &tree.Tree{Value: 2}, Right: &tree.Tree{Value: 1}, Value: 1}

	sameTree := treeutil.Same(&tree1, &tree2)

	if !sameTree {
		t.Errorf("Expected true  got %t", sameTree)
	}
}

func TestDifferentTree(t *testing.T) {
	tree1 := tree.Tree{Left: &tree.Tree{Value: 1}, Right: &tree.Tree{Value: 1}, Value: 2}
	tree2 := tree.Tree{Left: &tree.Tree{Value: 2}, Right: &tree.Tree{Value: 1}, Value: 1}

	sameTree := treeutil.Same(&tree1, &tree2)

	if sameTree {
		t.Errorf("Expected false  got %t", sameTree)
	}
}
