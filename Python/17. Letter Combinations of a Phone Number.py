# https://leetcode.com/problems/letter-combinations-of-a-phone-number/

# Runtime: 28 ms, faster than 66.62% of Python3 online submissions for Letter Combinations of a Phone Number.
# Memory Usage: 12.9 MB, less than 98.53% of Python3 online submissions for Letter Combinations of a Phone Number.

class Solution:
    def __init__(self):
        self.key_map = {"2": ["a", "b", "c"], "3": ["d", "e", "f"], "4": ["g", "h", "i"], "5": ["j", "k", "l"], "6": ["m", "n", "o"],
		"7": ["p", "q", "r", "s"], "8": ["t", "u", "v"], "9": ["w", "x", "y", "z"]}
        
        
    def letterCombinations(self, digits, current=None):        
        result = []
        
        if not digits:
            if current:
                return [current]
            else:
                return []
        
        if not current:
            current =""
        
        for char in self.key_map[digits[0]]:
            result += self.letterCombinations(digits[1:], current + char)
        
        return result
    

# print(Solution().letterCombinations("23456789234567"))
Solution().letterCombinations("234567892345678")