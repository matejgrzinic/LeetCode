package main

import (
	"reflect"

	"./test0049"
)

func main() {
	test0049.Test(groupAnagrams)
}

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	maps := []map[string]int{}
	for _, word := range strs {
		m1 := countChars(word)
		ok := false
		for i, m2 := range maps {
			if len(word) == len(result[i][0]) {
				if reflect.DeepEqual(m1, m2) {
					result[i] = append(result[i], word)
					ok = true
					break
				}
			}
		}
		if !ok {
			maps = append(maps, m1)
			result = append(result, []string{word})
		}
	}
	return result
}

func countChars(word string) map[string]int {
	counts := map[string]int{}
	for _, char := range word {
		if _, ok := counts[string(char)]; ok {
			counts[string(char)]++
		} else {
			counts[string(char)] = 1
		}
	}
	return counts
}
