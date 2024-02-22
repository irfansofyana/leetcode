func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    ans := make([]int, n)
    
    for i := 0; i < n; i++ {
        ans[i] = nums[(i - k + n)%n]
    }
    for i := 0; i < n; i++ {
        nums[i] = ans[i]
    }
}