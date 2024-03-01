/*
Leet Code #242

Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

package arrays

func IsAnagram(s string, t string) bool {

	if len(s) != (len(t)) {
		return false
	}

	letterCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		letterCount[s[i]] = letterCount[s[i]] + 1
		letterCount[t[i]] = letterCount[t[i]] - 1
	}

	for _, v := range letterCount {
		if v != 0 {
			return false
		}
	}

	return true

}
