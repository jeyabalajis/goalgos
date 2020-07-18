package treeutil

import (
	"strconv"
	"sync"

	"github.com/jeyabalajis/goalgos/stringutil"
	"github.com/jeyabalajis/goalgos/tree"
)

var wg sync.WaitGroup

// Walk walks the tree t sending all values from the tree to the channel ch.
func walk(t *tree.Tree, ch *chan int) {
	if t == nil {
		return
	}

	defer wg.Done()

	if t.Left != nil {
		wg.Add(1)
		go walk(t.Left, ch)
	}

	if t.Right != nil {
		wg.Add(1)
		go walk(t.Right, ch)
	}

	*ch <- t.Value
}

// traverse each node and build up hash map by each node
func traverse(t *tree.Tree, hm *TreePath, depth int, nodeQualifier string, hashKey string, acc int) {
	if t == nil {
		wg.Done()
		return
	}

	if nodeQualifier == "ROOT" {
		hashKey = strconv.Itoa(depth) + "-" + strconv.Itoa(t.Value) + "-" + nodeQualifier
		hm.createHash(hashKey, t.Value)
	} else {
		var left bool = false
		if nodeQualifier == "LEFT" {
			left = true
		}

		hm.putHashValue(hashKey, left, acc+t.Value)
	}

	var _nodeQualifier = nodeQualifier
	var _acc = 0
	if t.Left != nil {

		//  if it is root, qualify as left or right, else, pass on the qualifier
		if nodeQualifier == "ROOT" {
			_nodeQualifier = "LEFT"
		} else {
			_acc = acc + t.Value
		}

		wg.Add(1)
		go traverse(t.Left, hm, depth+1, _nodeQualifier, hashKey, _acc)

		wg.Add(1)
		go traverse(t.Left, hm, depth+1, "ROOT", "", 0)

	}

	if t.Right != nil {

		//  if it is root, qualify as left or right, else, pass on the qualifier
		if nodeQualifier == "ROOT" {
			_nodeQualifier = "RIGHT"
		} else {
			_acc = acc + t.Value
		}

		wg.Add(1)
		go traverse(t.Right, hm, depth+1, _nodeQualifier, hashKey, _acc)

		wg.Add(1)
		go traverse(t.Right, hm, depth+1, "ROOT", "", 0)
	}
	wg.Done()
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	wg.Add(1)

	go walk(t1, &ch1)

	go func(wg *sync.WaitGroup, c chan int) {
		wg.Wait()
		close(c)
	}(&wg, ch1)

	var treeHash1 uint32 = 0

	for i := range ch1 {
		treeHash1 += stringutil.Hash(strconv.Itoa(i))
	}

	ch2 := make(chan int)
	wg.Add(1)
	go walk(t2, &ch2)

	go func(wg *sync.WaitGroup, c chan int) {
		wg.Wait()
		close(c)
	}(&wg, ch2)

	var treeHash2 uint32 = 0
	for i := range ch2 {
		treeHash2 += stringutil.Hash(strconv.Itoa(i))
	}

	if treeHash1 != treeHash2 {
		return false
	}

	return true
}

// MaxSumPath provides the maximum sum possible in a tree and also the path that corresponds
// to this maximum sum
func MaxSumPath(t *tree.Tree) (maxValue int) {
	/*
		(1) Keep each node in a tree and it's level as a unique hashMap.
		(2) Traverse the tree. With each traversal, perform the following on the HashKey
			- collect the node value
			- collect the node left, call traversal recursively, but with the HashKey & the current list
				fire a new tree walk for the left node as a Hash Key
			- collect the node right, call traversal recursively, but with the HashKey & the current list
				fire a new tree walk for the left node as a Hash Key
		(3) Once all the traversals are done, sum up the numbers under each list of all keys under  hash map and find max across the HashMap

		Example: Let us consider the following tree:
		{
			level 0: 6->4, 6->3
			level 1(R): 3->13, 3-> 1
			level 1(L): 4
		}
		The HashMap will look as follows:
		{
			6(level 0)Root: [6, [6,4], [6,3], [6,4,3], [6,4,3,13], [6,4,3,1]
			4(level 1)Left: 4
			3(level 1)Right: [[3,13], [3,1], [3,13,1]]
			13(level 2)Left: [13]
			3(level 2)Right: [1]
		}
	*/

	var treePaths TreePath = make(TreePath)

	wg.Add(1)
	go traverse(t, &treePaths, 0, "ROOT", "", 0)

	wg.Wait()

	treePaths.Print()
	maxValue = treePaths.getMaxValue()
	return maxValue
}
