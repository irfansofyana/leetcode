func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n+2)
    
    dp[0] = 0; dp[1] = 0
    for i := 2; i < n+1; i++ {
        mini := dp[i-1] + cost[i-1]
        if dp[i-2] + cost[i-2] < mini {
            mini = dp[i-2] + cost[i-2]
        }
        dp[i] = mini
    }
    
    return dp[n]
}