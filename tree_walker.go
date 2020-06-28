package main

import "golang.org/x/tour/tree"
import "fmt"
import "sync"

var wg sync.WaitGroup

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch *chan int) {
	if t == nil {
		return
	}
	
	defer wg.Done()
	
	if t.Left != nil {
		wg.Add(1)
		go Walk(t.Left, ch)
	}
	
	if t.Right != nil {
		wg.Add(1)
		go Walk(t.Right, ch)
	}
	
	*ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go Walk(tree.New(2), &ch)
	
	go func(wg *sync.WaitGroup, c chan int) {
		wg.Wait()
		close(c)
	}(&wg, ch)
	
	for i := range ch {
        fmt.Println(i)
    }
}