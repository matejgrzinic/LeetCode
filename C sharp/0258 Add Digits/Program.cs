// https://leetcode.com/problems/add-digits/submissions/

// Runtime: 25 ms, faster than 86.88% of C# online submissions for Add Digits.
// Memory Usage: 26.7 MB, less than 9.05% of C# online submissions for Add Digits.

public class Solution
{
    public int AddDigits(int num)
    {
        int result = num;
        while (result >= 10)
        {
            num = result;
            result = 0;
            while (num > 0)
            {
                result += num % 10;
                num /= 10;
            }
        }
        return result;
    }

    public static void Main()
    {
        Solution x = new();
        Console.WriteLine(x.AddDigits(38));
    }
}