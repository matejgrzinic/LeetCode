// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

package main

import (
	"fmt"
	"time"
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Letter Combinations of a Phone Number.

// func main() {
// 	t1 := time.Now().UnixNano()
// 	letterCombinations("2345678923445")
// 	t2 := time.Now().UnixNano()
// 	fmt.Println((t2 - t1) / 1000000)
// 	t1 = time.Now().UnixNano()
// 	letterCombinationsSecond("2345678923445")
// 	t2 = time.Now().UnixNano()
// 	fmt.Println((t2 - t1) / 1000000)
// 	//fmt.Println(letterCombinations("234567892345")) // [ad ae af bd be bf cd ce cf]
// }

func letterCombinations(digits string) []string {
	result := []string{}

	keymap := map[string][]string{"2": []string{"a", "b", "c"}, "3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"}, "5": []string{"j", "k", "l"}, "6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"}, "8": []string{"t", "u", "v"}, "9": []string{"w", "x", "y", "z"}}

	if len(digits) == 0 {
		return result
	}
	addLetter("", digits, keymap, &result)
	return result
}

func addLetter(current string, next string, keymap map[string][]string, result *[]string) {
	if len(next) == 0 {
		*result = append(*result, current)
	} else {
		for _, key := range keymap[string(next[0])] {
			addLetter(current+string(key), next[1:], keymap, result)
		}
	}
}

func letterCombinationsSecond(digits string) []string {

	// Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
	// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Letter Combinations of a Phone Number.

	result := []string{}

	keymap := map[string][]string{"2": []string{"a", "b", "c"}, "3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"}, "5": []string{"j", "k", "l"}, "6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"}, "8": []string{"t", "u", "v"}, "9": []string{"w", "x", "y", "z"}}

	if len(digits) == 0 {
		return result
	}

	if len(digits) == 1 {
		return keymap[string(digits[0])]
	}

	for _, key := range keymap[string(digits[0])] {
		r2 := letterCombinationsSecond(digits[1:])
		for index := range r2 {
			r2[index] = string(key) + r2[index]
			result = append(result, r2[index])
		}
	}

	return result
}
