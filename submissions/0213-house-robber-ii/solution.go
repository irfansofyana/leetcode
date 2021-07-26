func rob(nums []int) int {
    n := len(nums)
    
    if n == 1 {
        return nums[0]
    }
    
    var dp = make([][]int, n+1)
    dp[0] = []int{0, 0}
    dp[1] = []int{0, nums[0]}
    
    for i := 2; i <= n; i++ {
        for j := 0; j <= 1; j++ {
            dp[i] = append(dp[i], 0)
            if i < n || (i == n && j == 0) {
                maks := dp[i-2][j] + nums[i-1]
                if dp[i-1][j] > maks {
                    maks = dp[i-1][j]
                }
                dp[i][j] = maks
            }
        }    
    }
    
    ans := dp[n-1][1]
    if dp[n-1][0] > dp[n-1][1] {
        ans = dp[n-1][0]
    }
    if dp[n][0] > ans {
        ans = dp[n][0]
    }
    
    return ans
}
