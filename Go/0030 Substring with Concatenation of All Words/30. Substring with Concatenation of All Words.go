package main

// Runtime: 204 ms, faster than 30.12% of Go online submissions for Substring with Concatenation of All Words.
// Memory Usage: 3.2 MB, less than 100.00% of Go online submissions for Substring with Concatenation of All Words.

import "fmt"

func main() {
	//fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"})) // [0, 9]
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"})) // []
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","good"})) // [8]
}

func findSubstring(s string, words []string) []int {
	result := []int{}

	if len(s) == 0 || len(words) == 0{
		return result
	}
	
	totaln := 0
	for _, word := range words {
		totaln += len(word)
	}

	for index := range s[:len(s) - totaln + 1]{
		if helper(s[index:], words){
			result = append(result, index)
		}
	}	
	return result
}

func helper(s string, words []string) bool{
	if len(words) == 0{
		return true
	}
	for index, word := range words {
		n := len(word)
		if s[:n] == word{
			return helper(s[n:], makeSlice(words, index))
			break
		}		
	}	
	return false
}

func makeSlice(words []string, index int) []string{
	n := len(words) - 1	
	
	if index >= n {
		return words[:n]
	}
	words[index], words[n] = words[n], words[index]
	return words[:n]
}