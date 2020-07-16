# 79. Word Search

https://leetcode.com/problems/word-search/

To do:
- optimize

```
Runtime: 24 ms, faster than 15.62% of Go online submissions for Word Search.
Memory Usage: 6.3 MB, less than 23.36% of Go online submissions for Word Search.
```

Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

**Example:**
```
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
``` 

**Constraints:**

- board and word consists only of lowercase and uppercase English letters.
- 1 <= board.length <= 200
- 1 <= board[i].length <= 200
- 1 <= word.length <= 10^3