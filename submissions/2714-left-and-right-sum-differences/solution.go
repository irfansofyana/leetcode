func leftRightDifference(nums []int) []int {
    left := 0; right := 0; n := len(nums)

    ans := make([]int, 0)
    for i := 0; i < n; i++ {
        ans = append(ans, left)
        left += nums[i]
    }
    for i := n-1; i >= 0; i-- {
        ans[i] = abs(ans[i], right)
        right += nums[i]
    }

    return ans
}

func abs(a, b int) int {
    if a < b {
        return b - a
    }

    return a - b
}
