package main

import (
	"sort"

	"./test0049"
)

func main() {
	test0049.Test(groupAnagrams)
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	examples := map[string]int{}
	for _, word := range strs {
		sortedWord := sortString(word)

		if i, ok := examples[sortedWord]; ok {
			result[i] = append(result[i], word)
		} else {
			examples[sortedWord] = len(result)
			result = append(result, []string{word})
		}
	}
	return result
}

func sortString(word string) string {
	s := []byte(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}
