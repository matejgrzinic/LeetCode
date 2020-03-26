// https://leetcode.com/problems/longest-substring-without-repeating-characters

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 3.8 MB, less than 23.33% of Go online submissions for Longest Substring Without Repeating Characters.

package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcb"))  // 3
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring("cdd"))      // 2
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	seen := map[string]int{}
	best := 0
	l, r := 0, 0
	for r < len(s) {
		key := string(s[r])
		for !keyIsZero(key, seen) {
			seen[string(s[l])]--
			l++
		}
		seen[key] = 1
		if r-l > best {
			best = r - l
		}
		r++
	}
	return best + 1
}

func keyIsZero(key string, seen map[string]int) bool {
	if val, ok := seen[key]; ok {
		if val > 0 {
			return false
		}
	}
	return true
}
