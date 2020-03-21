# https://leetcode.com/problems/container-with-most-water/

# Runtime: 112 ms, faster than 99.64% of Python3 online submissions for Container With Most Water.
# Memory Usage: 14.4 MB, less than 81.05% of Python3 online submissions for Container With Most Water.

class Solution:
    def maxArea(self, height):
        left, right = 0, len(height) - 1
        best = 0
        while left < right:           
            #size = 0
            if height[left] < height[right]:
                size = height[left] * (right - left)
                left += 1
            else:
                size = height[right] * (right - left)
                right -= 1
            if size > best:
                best = size
        return best
 
    
print(Solution().maxArea([1,8,6,2,5,4,8,3,7]))