import "math"

func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = triangle[i]
    }

    for i := 1; i < n; i++ {
        for j := 0; j <= i; j++ {
            if j == 0 {
                dp[i][j] += dp[i-1][j]
                continue
            }

            if j == i {
                dp[i][j] += dp[i-1][j-1]
                continue
            }

            dp[i][j] += min(dp[i-1][j], dp[i-1][j-1])
        }
    }

    mini := math.MaxInt64
    for i := 0; i < n; i++ {
        if dp[n-1][i] < mini {
            mini = dp[n-1][i]
        }
    }

    return mini
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}