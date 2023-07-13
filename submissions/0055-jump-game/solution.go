func canJump(nums []int) bool {
    n := len(nums)
    can := make([]bool, n)
    can[n-1] = true

    for i := n-2; i >= 0; i-- {
        jump := nums[i]
        for j := 1; j <= jump; j++ {
            if i + j <= n-1 && can[i+j] {
                can[i] = true
                break
            }
        }
    }

    return can[0]
}

