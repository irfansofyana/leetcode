func containsNearbyDuplicate(nums []int, k int) bool {
    lastIdx := map[int]int{}
    n := len(nums)
    ans := false
    
    for i := 0; i < n; i++ {
        if val, ok := lastIdx[nums[i]]; ok {
            if i-val <= k {
                ans = true
                break
            }
        }
        lastIdx[nums[i]] = i
    }
    
    return ans
}
