func countSquares(matrix [][]int) int {
    m := len(matrix); n := len(matrix[0])
    dp := make([][]int, m+2)
    
    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            dp[i] = append(dp[i], 0)
        }
        if i > 0 {
            for j := 1; j <= n; j++ {
                dp[i][j] = dp[i-1][j] + dp[i][j-1] + matrix[i-1][j-1] - dp[i-1][j-1]
            }   
        }
    }
    
    count := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            for k := 1; i + k - 1 <= m && j + k - 1 <= n; k++ {
                cnt := dp[i+k-1][j+k-1] - dp[i-1][j+k-1] - dp[i+k-1][j-1] + dp[i-1][j-1]
                if cnt == k * k {
                    count += 1
                }
            }
        }
    }
    
    return count
}
