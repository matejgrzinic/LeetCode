class Solution:
    def longestPalindrome(self, s: str) -> str:
        best = ""
        n = len(s)
        if n == 1:
            return s[0]
        for i in range(n - 1):
            p = s[i]
            
            start, end = i, i            
            for k in range(1, n):
                if i + k < n and s[i] == s[i +k]:
                    p += s[i]
                    end += 1
                    if i - k >= 0 and s[i] == s[i - k]:
                        p += s[i]
                        start -= 1
                    else:
                        break   
                else:
                    break                                    
            
            if len(p) > len(best):
                best = p               
                       
        
            for index in range(1, n):
                si = start - index
                ei = end + index
                if si >= 0 and ei < n and s[si] == s[ei]:
                    p = s[si] + p + s[si]
                    if len(p) > len(best):
                        best = p 
                else:
                    break                   
           
            
        return best
    
    def isPalindrome(self, s):
        for index in range(len(s) // 2):
            if s[index] != s[-index - 1]:
                return False
        return True
        
s = Solution()
print(s.longestPalindrome("caacc"))
# print("Expected: bab|aba, got:", s.longestPalindrome("babad"))
#print("Expected: bb, got:", s.longestPalindrome("cbbd"))
#print(s.longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy"))

