func countVowelStrings(n int) int {
    dp := make([][5]int, n + 1)
    
    for i := 0; i < 5; i++ {
        dp[0][i] = 1
    }
    
    for i := 1; i < n; i++ {
        for j := 0; j < 5; j++ {
            for k := j; k < 5; k++ {
                dp[i][j] += dp[i-1][k]
            }
        }
    }
    
    ans := 0
    for i := 0; i < 5; i++ {
        ans += dp[n-1][i]
    }
    
    return ans
}