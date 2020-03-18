class Solution:
    def longestPalindrome(self, s: str) -> str:
        if len(s) < 2 or s ==s[::-1]:
            return s
        n = len(s)
        start, maxlen = 0,1
        
        for i in range(n):
            odd = s[i-maxlen-1:i+1]
            even = s[i-maxlen:i+1]
            
            if i - maxlen -1>=0 and odd == odd[::-1]:
                start = i - maxlen - 1
                maxlen +=2                
            elif i - maxlen >=0 and even ==even[::-1]:
                start = i - maxlen 
                maxlen +=1
        return s[start:start+maxlen]
    
    
        
s = Solution()
print(s.longestPalindrome("caacc"))
# print("Expected: bab|aba, got:", s.longestPalindrome("babad"))
#print("Expected: bb, got:", s.longestPalindrome("cbbd"))
#print(s.longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy"))

