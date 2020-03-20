# https://leetcode.com/problems/zigzag-conversion/

# Runtime: 36 ms, faster than 99.98% of Python3 online submissions for ZigZag Conversion.
# Memory Usage: 13 MB, less than 98.57% of Python3 online submissions for ZigZag Conversion  

class Solution:
    def convert(self, s: str, numRows: int) -> str:             
        if numRows == 1:
            return s
        else:           
            rows = [''] * numRows

            going_up = True
            index_row = 0        

            for char in s:
                rows[index_row] += char           
                index_row = index_row + 1 if going_up else index_row - 1
                if index_row == 0 or index_row == numRows - 1:
                    going_up = not going_up                   
            return "".join(rows)
        

s = Solution()
print(s.convert("AB", 2))


