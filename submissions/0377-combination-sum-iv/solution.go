func combinationSum4(nums []int, target int) int {
    var dp = make([]int, target+1)
    n := len(nums)
    
    dp[0] = 1
    for i:= 1; i <= target; i++ {
        for j := 0; j < n; j++ {
            if i >= nums[j] {
                dp[i] += dp[i-nums[j]]
            }
        }
    }
    return dp[target]
}
