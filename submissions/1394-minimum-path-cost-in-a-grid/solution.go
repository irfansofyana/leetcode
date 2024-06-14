import "fmt"

func minPathCost(grid [][]int, moveCost [][]int) int {
    n := len(grid); m := len(grid[0])
    dp := initDp(n, m)

    for i := 1; i < n; i++ {
        for j := 0; j < m; j++ {
            for k := 0; k < m; k++ {
                valGrid := grid[i-1][k]
                cost := moveCost[valGrid][j]
                dp[i][j] = min(dp[i][j], valGrid+cost+dp[i-1][k])
            }
        }
    }

    minimum := 2000000000
    for i := 0; i < m; i++ {
        minimum = min(minimum, dp[n-1][i] + grid[n-1][i])
    }

    return minimum
}

func initDp(n, m int) [][]int {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
        for j := 0; j < m; j++ {
            dp[i][j] = 2000000000
            if i == 0 {
                dp[i][j] = 0
            }
        }
    }
    return dp
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
