func findMaxAverage(nums []int, k int) float64 {
    n := len(nums)
    
    for i := 1; i < n; i++ {
        nums[i] += nums[i-1]
    }
    
    maks := 0
    for i := 0; i + k - 1 < n; i++ {
        var sum int
        if i == 0 {
            sum = nums[i + k - 1]
        } else {
            sum = nums[i + k - 1] - nums[i-1]
        }
        if i == 0 {
            maks = sum
        } else if sum > maks {
            maks = sum
        }
    }
    
    return float64(maks) / float64(k)
}