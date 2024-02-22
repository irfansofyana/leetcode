func maximumProduct(nums []int) int {
    sort.Ints(nums);
    n := len(nums)
    
    var ans int
    ans = nums[n-1] * nums[n-2] * nums[n-3]
    if nums[0] < 0 && nums[1] < 0 {
        tmp := nums[0] *  nums[1] * nums[n-1]
        if tmp > ans {
            ans = tmp
        }
    }
    
    return ans
}