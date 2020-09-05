# Word Break II

## Problem Statement:

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

- The same word in the dictionary may be reused multiple times in the segmentation.
- You may assume the dictionary does not contain duplicate words.

## Example:

```
Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.
```

Difficulty: :moneybag: :moneybag: :moneybag:

[Leet Code Link](https://leetcode.com/problems/word-break-ii/)

## Solution:

This problem is an excellent candidate for combining recursion with parallelism. The pseudo-code can be organized as follows:

1. Start with the whole string as an input and an empty string as a prefix
2. Traverse through the characters & check for match against a word in word dictionary
3. If a match is obtained, collect this word as a prefix and fire a __recursive thread__ with rest of the string as input and the match as prefix.

That's it!

Let us take the input __pineapplepenapple__

```
1. First match: pine (Rest of string: applepenapple)
    - 1.1 First match: apple (Rest of string: penapple)
        - 1.1.1 penapple is a match. __Result: pine apple penapple__
    - 1.2 Second match: applepen (Rest of string: apple)
        - 1.2.1 apple is a match. __Result: pine applepen apple__
2. Second match: pineapple (Rest of string: penapple)
    - 2.1 First match: pen (Rest of string: apple)
        - 2.1.1 apple is a match. __Result: pineapple pen apple__

Output:

- pine apple penapple
- pine applepen apple
- pineapple pen apple
```

### Code
```
func traverse(
    s string, 
    wordDict WordDictionary, 
    hm *SentenceMap, 
    prefix string, 
    threadNo int
    ) {
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
```