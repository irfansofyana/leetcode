func subsetXORSum(nums []int) int {
    n := len(nums)
    
    ans := 0
    for i := 0; i < (1 << n); i++ {
        xor := 0
        for j := 0; j < n; j++ {
            if (i&(1<<j)) > 0 {
                xor ^= nums[j]
            }
        }
        ans += xor
    }
    
    return ans
}
