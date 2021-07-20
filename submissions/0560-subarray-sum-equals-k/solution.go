func subarraySum(nums []int, k int) int {
    var n int = len(nums)
    var ans int = 0
    var pSum = make([]int, n + 2)
    var freq = map[int]int{}
    
    for i := 1; i <= n; i++ {
        pSum[i] = pSum[i-1] + nums[i-1]
        freq[pSum[i]] += 1
    }
    
    for i := 1; i <= n; i++ {
        var target int = k + pSum[i-1]
        ans += freq[target]
        freq[pSum[i]] -= 1
    }
    
    return ans
}
