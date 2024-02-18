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
