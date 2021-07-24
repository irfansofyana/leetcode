func lengthOfLIS(nums []int) int {
    n := len(nums)
    var dp = make([]int, n+1)
    ans := 0
    
    for i := 0; i < n; i++ {
        maks := 0
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] && dp[j] > maks {
                maks = dp[j]
            }
        }
        dp[i] = maks + 1
        if dp[i] > ans {
            ans = dp[i]
        }
    }
    return ans
}
