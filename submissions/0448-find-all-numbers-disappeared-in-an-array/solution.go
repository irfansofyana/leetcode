func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    freq := make([]int, n+1)
    
    for i := 0; i < n; i++ {
        freq[nums[i]] += 1
    }
    
    ans := make([]int, 0)
    for i := 1; i <= n; i++ {
        if freq[i] == 0 {
            ans = append(ans, i)
        }
    }
    
    return ans
}
