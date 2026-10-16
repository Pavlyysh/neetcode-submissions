func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, word := range strs {
		var key [26]int
		for _, ch := range word {
			key[ch-'a']++
		}

		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0, len(groups))

	for _, words := range groups {
		result = append(result, words)
	}

	return result
}
