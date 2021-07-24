func rob(nums []int) int {
    n := len(nums)
    var dp = make([]int, n+1)
    
    dp[0] = nums[0];
    for i := 1; i < n; i++ {
        maks := dp[i-1]
        if i-2 >= 0 && dp[i-2] + nums[i] > maks {
            maks = dp[i-2] + nums[i]
        } else if nums[i] > maks {
            maks = nums[i]
        }
        dp[i] = maks
    }
    
    return dp[n-1]
}
