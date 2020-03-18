class Solution:
    def longestPalindrome(self, s: str) -> str:
        # Top
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
    
    
        
    def longestPalindromeMine(self, s: str) -> str:
        # Runtime: 296 ms, faster than 93.42% of Python3 online submissions for Longest Palindromic Substring.
        # Memory Usage: 12.9 MB, less than 100.00% of Python3 online submissions for Longest Palindromic Substring.
        best = ""
        n = len(s)
        if len(s) < 2 or s ==s[::-1]:
            return s        
        for char_index in range(n - 1):
            string = s[char_index]            
            start, end = char_index, char_index            
            for offset in range(1, n):
                if char_index + offset < n and s[char_index] == s[char_index + offset]:
                    string += s[char_index]
                    end += 1
                    if char_index - offset >= 0 and s[char_index] == s[char_index - offset]:
                        string += s[char_index]
                        start -= 1
                    else:
                        break   
                else:
                    break                                                  
        
            for index in range(1, n):
                start_index = start - index
                end_index = end + index
                if start_index >= 0 and end_index < n and s[start_index] == s[end_index]:
                    string = s[start_index] + string + s[start_index]
                    if len(string) > len(best):
                        best = string 
                else:
                    break                   
            
            if len(string) > len(best):
                best = string                  
            
        return best
    
    
        
s = Solution()
print(s.longestPalindrome("caacc"))
# print("Expected: bab|aba, got:", s.longestPalindrome("babad"))
#print("Expected: bb, got:", s.longestPalindrome("cbbd"))
#print(s.longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy"))

