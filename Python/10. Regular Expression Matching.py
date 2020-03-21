# https://leetcode.com/problems/regular-expression-matching/submissions/
# Runtime: 32 ms, faster than 97.85% of Python3 online submissions for Regular Expression Matching.
# Memory Usage: 13.1 MB, less than 100.00% of Python3 online submissions for Regular Expression Matching.

import re

class Solution:    
    def __init__(self):
        self.seen = {}
    
    def isMatchRe(self, s: str, p: str) -> bool:
        # Runtime: 60 ms, faster than 45.31% of Python3 online submissions for Regular Expression Matching.
        # Memory Usage: 13 MB, less than 100.00% of Python3 online submissions for Regular Expression Matching.        
        return True if  re.match(r'{}$'.format(p), s) else False
        
    def isMatch(self, s: str, p: str) -> bool:
        # Runtime: 32 ms, faster than 97.85% of Python3 online submissions for Regular Expression Matching.
        # Memory Usage: 13.1 MB, less than 100.00% of Python3 online submissions for Regular Expression Matching.
        if s+"_"+p in self.seen.keys():
            return False
        else:
            self.seen[s+"_"+p] = True
        
        if not s:
            while '*' in p[:2]:
                p = p[2:]
            if not p:
                return True
            else:
                return False
            
        if not p:
            return False
        
        if '*' in p[:2]:
            if s[0] == p[0] or p[0] == '.':
                if self.isMatch(s[1:], p):                
                    return True           
            
            if self.isMatch(s, p[2:]):
                return True
        else:
            if s[0] == p[0] or p[0] == '.':
                if self.isMatch(s[1:], p[1:]):
                    return True    
            
                            
        return False
            
        
#print(Solution().isMatch("", "c*c*"))
# print(Solution().isMatch("ab", ".*c"))
# print(Solution().isMatch("aa", "a"))
#print(Solution().isMatch("aa", "a*"))
# print(Solution().isMatch("ab", ".*"))
#print(Solution().isMatch("aab", "c*a*b"))
#print(Solution().isMatch("mississippi", "mis*is*p*."))
print(Solution().isMatch("mississippi", "mis*is*ip*."))

