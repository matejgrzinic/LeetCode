# https://leetcode.com/problems/reverse-integer/

class Solution:
    def reverse(self, x: int) -> int:
        # Runtime: 28 ms, faster than 77.55% of Python3 online submissions for Reverse Integer.
        # Memory Usage: 13 MB, less than 100.00% of Python3 online submissions for Reverse Integer.
        sign = 1
        if x < 0:
            sign = -1
            x = -x
                    
        
        num = int(str(x)[::-1])
        if num < 2147483648: # 2 ** 31
            return num * sign
        return 0       
        
        
    
s = Solution()
print(s.reverse(123))
