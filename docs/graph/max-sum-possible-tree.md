# Maximum Sum Possible in a Tree

## Problem Statement:

Given a __non-empty__ binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain __at least one node__ and does not need to go through the root.

### Example: 
![Max Path Sum Illustration](https://jeyabalajis.github.io/goalgos/images/max-path-sum.png)

Difficulty: :moneybag: :moneybag: :moneybag:

[Leet Code Link](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

## Solution

> TLDR; Treat each node as an independent tree on it's own and compute max path sum per node. The max of all such sums yields the result.

Fundamentally, this is a tree traversal problem. The following are the facts provided in the problem statement:

1. A path must pass through a node. 
2. A path need not have to go through root node for it to be considered.

This means that the maximum sum path could either be a single node on it's own (or) it could be a small sub-tree contained within a node (and) it can be at level level of the tree.

To solve this, 

> Each node must be treated as a __first class citizen__ on it's own. I.e. Keep a Hash at a particular node level.

> Whenever a new node is encountered, a new hash is formed and any children traversals are accumulated at this hash level

> The algorithm must keep track of traversal __from each node__ along it's children; left and right separately.

Once the problem is broken down as above, the solution is pretty easy.

- Against each node (hash), keep track of the maximum sum on the left traversal and right traversal separately.
- The total sum for a hash (i.e. against a node) would be the node's value + left maximum + right maximum.
- Every new node traversal (including root) will result in two go routines: 
    - one go routine that's a continuation of it's parent
    - one go routine with this node as a first class citizen for the traversals below this node.
- The accumulated value uptil all the iterations for a Hash shall be passed as a parameter to the recursive function
- Hash Map updates (_left maximum_ and _right maximum_) will happen __only when the current accumulated value is greater than what is already available__.

For the diagram illustrated above, the Hash data structure would look as follows:
```json
{
    "0-6-ROOT": {
        "left_max": 6,
        "right_max": 16
    },
    "1-4-ROOT": {
        "left_max": -5,
        "right_max": 2
    },
    "1-3-ROOT": {
        "left_max": 1,
        "right_max": 13
    },
    "2--5-ROOT": {
        "left_max": 0,
        "right_max": 0
    },
    "2-2-ROOT": {
        "left_max": 0,
        "right_max": 0
    },
    "2-1-ROOT": {
        "left_max": 0,
        "right_max": 0
    },
    "2-13-ROOT": {
        "left_max": 0,
        "right_max": 0
    }
}
```

The maximum sum out of all these Hash Maps is the maximum sum path.

For the example mentioned above, this is against the root node and the value is 28.

The core logic is as follows:
```golang
// traverse each node and build up hash map by each node
func traverse(
    t *tree.Tree, 
    hm *TreePath, 
    depth int, 
    nodeQualifier string, 
    hashKey string, 
    acc int
    ) {
	if t == nil {
		wg.Done()
		return
	}

	if nodeQualifier == "ROOT" {
        hashKey = strconv.Itoa(depth) 
            + "-" + strconv.Itoa(t.Value) 
            + "-" + nodeQualifier
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
```





