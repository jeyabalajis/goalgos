package main

import (
	"fmt"

	"github.com/jeyabalajis/goalgos/stringutil"
)

func main() {

	s := "pineapplepenapple"
	wordDict := stringutil.WordDictionary{"apple", "pen", "applepen", "pine", "pineapple"}
	outSentences := stringutil.BreakSentenceByDict(s, wordDict)

	fmt.Println(outSentences)

	s1 := "catsandog"
	wordDict1 := stringutil.WordDictionary{"cats", "dog", "sand", "and", "cat"}
	outSentences1 := stringutil.BreakSentenceByDict(s1, wordDict1)

	fmt.Println(outSentences1)

}
