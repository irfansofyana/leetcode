func smallerNumbersThanCurrent(nums []int) []int {
    ans := make([]int, 0); n := len(nums)
    for i := 0; i < n; i++ {
        cnt := 0
        for j := 0; j < n; j++ {
            if j != i && nums[j] < nums[i] {
                cnt += 1
            }
        }
        ans = append(ans, cnt)
    }
    return ans
}