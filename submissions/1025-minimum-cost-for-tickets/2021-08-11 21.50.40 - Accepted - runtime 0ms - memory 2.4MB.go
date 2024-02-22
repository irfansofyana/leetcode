func mini(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func mincostTickets(days []int, costs []int) int {
    n := len(days); duration := [3]int{1, 7, 30}
    
    dp := make([]int, n + 1)
    dp[1] = mini(costs[0], mini(costs[1], costs[2]))
    
    for i := 2; i <= n; i++ {
        var minimum int = 2000000000
        for j, dur := range duration {
            for k := i; k >= 1; k-- {
                if days[i-1]-days[k-1]+1 > dur {
                    break
                }
                minimum = mini(minimum, dp[k-1] + costs[j])
            }
        }
        dp[i] = minimum
    }
    
    return dp[n]
}