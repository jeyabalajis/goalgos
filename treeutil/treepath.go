package treeutil

import (
	"fmt"
	"sync"
)

// NodeMax stores the value of a node and max values alongside any path on the left and right sides
type NodeMax struct {
	value    int
	leftMax  int
	rightMax int
}

var mutex = &sync.Mutex{}

// TreePath holds the maximum sum belonging to a particular node of a tree and below (across all possible paths).
// The node value along with the depth is kept as the hash key of the map
type TreePath map[string]NodeMax

func max(currMax int, incoming int) int {
	if incoming > currMax {
		return incoming
	}

	return currMax
}

func (hashMap TreePath) createHash(hashKey string, hashValue int) {
	mutex.Lock()
	hashMap[hashKey] = NodeMax{value: hashValue, leftMax: 0, rightMax: 0}
	mutex.Unlock()
	return
}

func (hashMap TreePath) putHashValue(hashKey string, left bool, hashValue int) {
	// Serialize access to hashmap to avoid data race
	mutex.Lock()

	if nodeMaximum, ok := hashMap[hashKey]; ok {
		maxValue := 0
		if left {
			maxValue = nodeMaximum.leftMax
			nodeMaximum.leftMax = max(maxValue, hashValue)
			hashMap[hashKey] = nodeMaximum
		} else {
			maxValue = nodeMaximum.rightMax
			nodeMaximum.rightMax = max(maxValue, hashValue)
			hashMap[hashKey] = nodeMaximum
		}

	}
	mutex.Unlock()
	return
}

func (hashMap TreePath) getMaxValue() (maxValue int) {
	mutex.Lock()
	maxValue = 0
	for _, nodeMaximum := range hashMap {
		value := nodeMaximum.value + nodeMaximum.leftMax + nodeMaximum.rightMax
		if value > maxValue {
			maxValue = value
		}
	}
	mutex.Unlock()
	return maxValue
}

// Print prints the hashmap
func (hashMap TreePath) Print() {
	mutex.Lock()
	fmt.Println(hashMap)
	mutex.Unlock()
}
