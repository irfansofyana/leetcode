func maxSubArray(nums []int) int {
    max := -2000000000
    curr := 0
    for _, num := range nums {
        curr += num
        if curr > max {
            max = curr
        }
        if curr < 0 {
            curr = 0
        }
    }
    return max
}
