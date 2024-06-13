func distinctDifferenceArray(nums []int) []int {
    n := len(nums); mep := make(map[int]bool)
    ans := make([]int, n)

    distinct := 0
    for i := 0; i < n; i++ {
        if !mep[nums[i]] {
            mep[nums[i]] = true
            distinct++
        }

        ans[i] += distinct;
    }

    distinct = 0; clear(mep);
    for i := n-1; i >= 1; i-- {
        if !mep[nums[i]] {
            mep[nums[i]] = true
            distinct++
        }

        ans[i-1] -= distinct
    }

    return ans
}
