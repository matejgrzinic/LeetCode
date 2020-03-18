class Solution:
    def longestPalindrome(self, s: str) -> str:
        best = ""
        n = len(s)
        if n == 1:
            return s[0]
        end_of_loop = n - 1
        for char_index in range(end_of_loop):
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
                end_of_loop = n - len(best)
            
        return best
    
    
        
s = Solution()
print(s.longestPalindrome("caacc"))
# print("Expected: bab|aba, got:", s.longestPalindrome("babad"))
#print("Expected: bb, got:", s.longestPalindrome("cbbd"))
#print(s.longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy"))

