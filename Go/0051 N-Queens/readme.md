# 51. N-Queens

https://leetcode.com/problems/n-queens/

```
Runtime: 4 ms, faster than 88.37% of Go online submissions for N-Queens.
Memory Usage: 4.8 MB, less than 100.00% of Go online submissions for N-Queens.
```

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

![Solution](https://assets.leetcode.com/uploads/2018/10/12/8-queens.png)

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

**Example:**
```
Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.
```