package stringutil

import (
	"sync"
)

// Sentence keeps track of all words separated
// and also tells whether it's complete (i.e. all words in dict separated and no letters left)
type Sentence struct {
	Value string
}

// WordDictionary is a slice of dictionary keywords
type WordDictionary []string

// SentenceMap holds the maximum sum belonging to a particular node of a tree and below (across all possible paths).
// The node value along with the depth is kept as the hash key of the map
type SentenceMap map[int]Sentence

var mutex = &sync.Mutex{}

func (wordDict WordDictionary) matched(s string) bool {

	for _, dictWord := range wordDict {
		if s == dictWord {
			return true
		}
	}
	return false
}

func (wordDict WordDictionary) minLength() int {
	minLength := 99999
	for _, dictWord := range wordDict {
		if len(dictWord) < minLength {
			minLength = len(dictWord)
		}
	}
	return minLength
}

func (hm SentenceMap) update(s string) {
	mutex.Lock()

	hashKey := int(Hash(s))
	hm[hashKey] = Sentence{Value: s}

	mutex.Unlock()
}

func (hm SentenceMap) toSlice() []string {
	var slices []string
	for _, item := range hm {
		slices = append(slices, item.Value)
	}
	return slices
}
