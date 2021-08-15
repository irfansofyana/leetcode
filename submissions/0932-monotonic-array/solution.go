func isMonotonic(nums []int) bool {
    if len(nums) == 1 {
        return true
    }
    limit := 0; n := len(nums)
    for limit + 1 < n && nums[limit] == nums[limit+1] {
        limit += 1
    }
    if limit + 1 >= n {
        return true
    }
    var isIncreasing bool = true
    if nums[limit] > nums[limit+1] {
        isIncreasing = false
    }
    for i := limit+1; i + 1 < n; i++ {
        if isIncreasing && nums[i] > nums[i+1] {
            return false
        }
        if !isIncreasing && nums[i] < nums[i+1] {
            return false
        }
    }
    return true
}
