package main

import (
	"testing"

	"github.com/jeyabalajis/goalgos/stringutil"
)

func TestWordBreak(t *testing.T) {
	s := "pineapplepenapple"
	wordDict := stringutil.WordDictionary{"apple", "pen", "applepen", "pine", "pineapple"}
	outSentences := stringutil.BreakSentenceByDict(s, wordDict)

	if len(outSentences) != 3 {
		t.Errorf("Expected 3 valid sentences but got %d", len(outSentences))
	}

	s1 := "catsanddog"
	wordDict1 := stringutil.WordDictionary{"cat", "cats", "and", "sand", "dog"}
	outSentences1 := stringutil.BreakSentenceByDict(s1, wordDict1)

	if len(outSentences1) != 2 {
		t.Errorf("Expected 2 valid sentences but got %d", len(outSentences1))
	}

}
