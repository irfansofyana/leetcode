func minFallingPathSum(matrix [][]int) int {
    n := len(matrix); dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        dp[0][i] = matrix[0][i]
    }
    
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            minimum := dp[i-1][j]
            for k := -1; k <= 1; k++ {
                if j+k >= 0 && j+k < n && dp[i-1][j+k] < minimum {
                    minimum = dp[i-1][j+k]
                }
            }
            dp[i][j] = minimum + matrix[i][j]
        }
    }
    
    ans := 2000000000
    for i := 0; i < n; i++ {
        if dp[n-1][i] < ans {
            ans = dp[n-1][i]
        }
    }
    
    return ans
}
