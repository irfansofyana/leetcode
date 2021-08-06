func buildArray(nums []int) []int {
    ans := make([]int, 0)
    
    for i := 0; i < len(nums); i++ {
        ans = append(ans, nums[nums[i]])
    }
    
    return ans
}
