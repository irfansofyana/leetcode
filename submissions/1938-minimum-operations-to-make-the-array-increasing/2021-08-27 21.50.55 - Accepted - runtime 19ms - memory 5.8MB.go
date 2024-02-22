func minOperations(nums []int) int {
    n := len(nums); 
    cost := 0
    for i := 0; i < n; i++ {
        if i == 0 {
            continue
        }
        if nums[i] <= nums[i-1] {
            cost += nums[i-1] + 1 - nums[i]
            nums[i] = nums[i-1] + 1
        }
    }
    return cost
}