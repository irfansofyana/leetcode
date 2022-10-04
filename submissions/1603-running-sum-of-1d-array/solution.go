func runningSum(nums []int) []int {
    sums := make([]int, len(nums))
    
    for i, num := range nums {
        if i == 0 {
            sums[i] = num
            continue
        }
        sums[i] = sums[i-1] + num
    }
    
    return sums
}
