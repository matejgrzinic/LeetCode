// https://leetcode.com/problems/find-the-difference/

// Runtime: 169 ms, faster than 12.05% of C# online submissions for Find the Difference.
// Memory Usage: 37.6 MB, less than 96.39% of C# online submissions for Find the Difference.

public class Solution
{
    public char FindTheDifference(string s, string t)
    {
        var seen = new Dictionary<char, int>();
        foreach (char c in s)
        {
            if (!seen.ContainsKey(c))
                seen[c] = 0;
            seen[c] += 1;
        }
        foreach (char c in t)
        {
            if (!seen.ContainsKey(c))
                seen[c] = 0;
            seen[c] -= 1;
        }
        foreach (var item in seen)
        {
            if (item.Value != 0)
                return item.Key;
        }
        return t[0];
    }

    public static void Main()
    {
        Solution x = new();
        Console.WriteLine(x.FindTheDifference("abcd", "abcde"));
    }
}