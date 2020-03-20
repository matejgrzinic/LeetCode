# https://leetcode.com/problems/string-to-integer-atoi/

class Solution:    
    def myAtoi(self, str: str) -> int:
        # Runtime: 16 ms, faster than 99.92% of Python3 online submissions for String to Integer (atoi).
        # Memory Usage: 12.9 MB, less than 100.00% of Python3 online submissions for String to Integer (atoi).
        trimmed = str.strip()
        
        f = False
        once = False
        for index, num in enumerate(trimmed):
            if not num.isdigit():
                if index == 0 and num in  '-+':
                    continue
                else:
                    if index > 0 and num == '.':
                        f = True
                        continue
                    trimmed = trimmed[:index + 1]
                    break
            once = True
        
        if not once:
            return 0
        elif not trimmed[-1].isdigit():
            trimmed = trimmed[:-1]        
        
        number = int(float(trimmed)) if f else int(trimmed)        
            
        if number < - 2 ** 31:
            return -2 ** 31
        elif number >= 2 ** 31:
            return 2 ** 31 - 1            
        return number
        
    
    
    def myAtoi2(self, str: str) -> int:        
        trimmed = str.strip()
        
        num = ""
        for char in trimmed:
            if num == "" and char in "+-":
                num += char
            elif char.isdigit():
                num += char    
            else:
                break                        
       
        if len(num) == 0 or num in "+-":
            return 0
        
        number = int(num)        
            
        if number < - 2 ** 31:
            return -2 ** 31
        elif number >= 2 ** 31:
            return 2 ** 31 - 1            
        return number
        
        
        
s = Solution()
print(s.myAtoi(""))