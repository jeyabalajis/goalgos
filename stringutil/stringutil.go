package stringutil

import (
	"fmt"
	"hash/fnv"
	"sync"
)

var wg sync.WaitGroup

// Hash returns a numeric representation of a string. Useful while comparing strings.
func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func traverse(s string, wordDict WordDictionary, hm *SentenceMap, prefix string, threadNo int) {
	/*
		incrementally accumulate each character of string & validate
		whether the substring at any step matches any dictonary word.

		if there is a match, collect this as a string and fire a thread to check for other parts of the string.
	*/

	i := 1
	for i <= len(s) {
		if len(s[0:i]) >= wordDict.minLength() {

			if wordDict.matched(s[0:i]) {
				currentStr := prefix + " " + s[0:i]

				// spawn a new thread with rest of string
				restOfStr := s[i:len(s)]
				if len(restOfStr) == 0 {
					fmt.Println("completed", currentStr)
					hm.update(currentStr)
				} else {
					wg.Add(1)
					go traverse(restOfStr, wordDict, hm, currentStr, threadNo+1)
				}

			}
		}
		i++
	}
	wg.Done()
}

// BreakSentenceByDict looks for all possible combinations of word breaks based on dictionary sent
// and returns all possible sentences with a space after the words (present in the dictionary)
func BreakSentenceByDict(s string, wordDict WordDictionary) []string {
	/*
		Iterate through the incoming string, one letter at a time.
		Compare the letter with each of the dict (Hash Compare)
		If there is no match, keep accumulating one letter after another
		If there is a match, keep collecting the words & call same function with reminder of string
		Continue accumulating letters to ensure no other matches are not missed out.
		if all characters are exhausted, flag it up as it is a valid output.

		example:
		s = "pineapplepenapple"
		wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]

		pine is the first match, push into an array with index 0
		[pine]
			call recursive function with applepenapple, index 0
				apple is the first match
				[pine, apple]
					call recursive function with penapple, index 0
					pen is a match
					[pine, apple, pen]
						call recursive function with apple
						apple is a match
						[pine, apple, pen, apple], reminder 0 chars VALID OUTPUT
				applepen is the second match
				[pine, applepen]
						call recursive function with apple, index 0
						apple is a match
						[pine, applepen, apple], reminder 0 chars VALID OUTPUT

		pineapple is the second match, push into array with index 1
		[pineapple]
			call recursive function with penapple, index 1
			pen is the first match
			[pineapple, pen]
				call recursive function with apple, index 1
				apple is a match
				[pineapple, pen, apple], reminder 0 chars VALID OUTPUT

		VALID OUTPUTS
		pine apple pen apple
		pine applepen apple
		pineapple pen apple
	*/

	var sentenceMap SentenceMap = make(SentenceMap)

	wg.Add(1)
	go traverse(s, wordDict, &sentenceMap, "", 0)

	wg.Wait()

	fmt.Println("done")
	return sentenceMap.toSlice()
}
