func longestPalindrome(s string) int {
    freq := make([]int, 52)
    n := len(s)
    
    for i := 0; i < n; i++ {
        var idx int
        if s[i] >= 'a' && s[i] <= 'z' {
            idx = int(s[i] - 'a')
        } else {
            idx = int(s[i] - 'A') + 26
        }
        freq[idx] += 1
    }
    
    ans := 0
    maxOdd := 0
    
    for i := 0; i < 52; i++ {
        if freq[i] % 2 == 0 {
            ans += freq[i]
        } else if freq[i] % 2 == 1  && freq[i] > maxOdd {
            if maxOdd > 0 {
                ans += maxOdd - 1
            }
            maxOdd = freq[i]
        } else if freq[i] % 2 == 1 {
            ans += freq[i] - 1
        }
    }
    ans += maxOdd
    
    return ans
}