package main

import (
	"fmt"

	"github.com/jeyabalajis/goalgos/tree"
	"github.com/jeyabalajis/goalgos/treeutil"
)

func main() {
	var level3LLeft1 = tree.Tree{Value: 120}
	var level3LRight1 = tree.Tree{Value: 1}

	var level3LLeft2 = tree.Tree{Value: 3}
	var level3LRight2 = tree.Tree{Value: 5}

	var level2LLeft = tree.Tree{Value: -5, Left: &level3LLeft1, Right: &level3LRight1}
	var level2LRight = tree.Tree{Value: 2, Left: &level3LLeft2, Right: &level3LRight2}
	var level1LLeft = tree.Tree{Left: &level2LLeft, Value: 4, Right: &level2LRight}

	var l3l = tree.Tree{Value: 145}
	var l3r = tree.Tree{Value: 6}
	var level2RLeft = tree.Tree{Value: 1}
	var level2RRight = tree.Tree{Value: 13, Left: &l3l, Right: &l3r}
	var level1RRight = tree.Tree{Left: &level2RLeft, Value: 3, Right: &level2RRight}

	var myTree = tree.Tree{Left: &level1LLeft, Value: 6, Right: &level1RRight}
	fmt.Println(myTree)
	maxValue := treeutil.MaxSumPath(&myTree)

	fmt.Println(maxValue)

}
