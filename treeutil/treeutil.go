package treeutil

import (
	"strconv"
	"sync"

	"github.com/jeyabalajis/goalgos/stringutil"
	"github.com/jeyabalajis/goalgos/tree"
)

var wg sync.WaitGroup

// Walk walks the tree t sending all values
// from the tree to the channel ch.
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
