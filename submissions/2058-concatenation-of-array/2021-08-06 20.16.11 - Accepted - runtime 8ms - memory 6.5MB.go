func getConcatenation(nums []int) []int {
    ans := make([]int, 0)
    for i := 0; i < 2 * len(nums); i++ {
        if i >= len(nums) {
            ans = append(ans, nums[i-len(nums)])
        } else {
            ans = append(ans, nums[i])
        }
    }
    return ans
}