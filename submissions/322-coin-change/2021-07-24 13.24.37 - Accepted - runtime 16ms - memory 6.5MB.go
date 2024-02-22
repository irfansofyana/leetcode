func coinChange(coins []int, amount int) int {
    var dp = make([]int, amount+1)
    const INF = 2000000000
    n := len(coins)
    
    for i := 1; i <= amount; i++ {
        dp[i] = INF
    }
    
    for i := 0; i < n; i++ {
        if coins[i] > amount {
            continue
        }
        
        dp[coins[i]] = 1
        for j := coins[i]; j <= amount; j++ {
            if dp[j-coins[i]] + 1 < dp[j] {
                dp[j] = dp[j-coins[i]] + 1
            }
        }
    }
    
    if dp[amount] == INF {
        return -1
    }
    return dp[amount]
}