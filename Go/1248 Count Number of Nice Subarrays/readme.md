# 1248. Count Number of Nice Subarrays

https://leetcode.com/problems/count-number-of-nice-subarrays/

```
Runtime: 112 ms, faster than 100.00% of Go online submissions for Count Number of Nice Subarrays.
Memory Usage: 7.5 MB, less than 20.00% of Go online submissions for Count Number of Nice Subarrays.
```

Given an array of integers nums and an integer k. A continuous subarray is called nice if there are k odd numbers on it.

Return the number of nice sub-arrays.

**Example 1:**

```
Input: nums = [1,1,2,1,1], k = 3
Output: 2
Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].
```

**Example 2:**

```
Input: nums = [2,4,6], k = 1
Output: 0
Explanation: There is no odd numbers in the array.
```

**_Example 3:_**

```
Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
Output: 16
```

**Constraints:**

- <= nums.length <= 50000
- <= nums[i] <= 10^5
- <= k <= nums.length
