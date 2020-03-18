class Solution:
    def longestPalindrome(self, s: str) -> str:
        if s == "":
            return s
        results = []
        best = s[0]
        for char in s:
            for index, substring in enumerate(results):
                results[index] += char
                if self.isPalindrome(results[index]) and len(results[index]) > len(best):
                    best = results[index]
            results.append(char)    
        return best
    
    def isPalindrome(self, s):
        for index in range(len(s) // 2):
            if s[index] != s[-index - 1]:
                return False
        return True
        
s = Solution()
print(s.longestPalindrome("babad"))
# print("Expected: bab|aba, got:", s.longestPalindrome("babad"))
# print("Expected: bb, got:", s.longestPalindrome("cbbd"))
# print(s.longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy"))

