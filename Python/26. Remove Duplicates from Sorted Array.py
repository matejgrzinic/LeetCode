# https://leetcode.com/problems/remove-duplicates-from-sorted-array

# Runtime: 84 ms, faster than 76.60% of Python3 online submissions for Remove Duplicates from Sorted Array.
# Memory Usage: 14.7 MB, less than 86.07% of Python3 online submissions for Remove Duplicates from Sorted Array.

class Solution:
    def removeDuplicates(self, nums):        
        unique = 0
        for index in range(1, len(nums)):
            if nums[unique] != nums[index]:
                unique += 1
                nums[unique] = nums[index]
        
        return unique + 1
    
    
nums = [1,1,2]
print(Solution().removeDuplicates(nums))
print(nums)