func minIncrementForUnique(nums []int) int {
    sort.Ints(nums); mep := make(map[int]bool)
    for _, num := range nums {
        mep[num] = true
    }

    idx := -1; ans := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            if idx <= nums[i] {
                idx = nums[i]+1
            }
            for mep[idx] {
                idx++
            }
            mep[idx] = true
            ans += idx - nums[i]
            idx++
        }
    }

    return ans
}
// 1 1 2 2 3 7
// 1 4 2 5 3 7
