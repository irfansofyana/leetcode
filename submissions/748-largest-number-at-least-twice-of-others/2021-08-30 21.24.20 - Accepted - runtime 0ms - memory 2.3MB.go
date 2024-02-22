func dominantIndex(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        tested := nums[i]
        valid := true
        for j := 0; j < n; j++ {
            if j == i {
                continue
            }
            if tested < 2 * nums[j] {
                valid = false
                break
            }
        }
        if valid {
            return i
        }
    }
    return -1
}