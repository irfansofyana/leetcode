func missingNumber(nums []int) int {
    n := len(nums)
    var sum int = 0
    
    for _, num := range nums {
        sum += num
    }
    
    return n*(n+1)/2 - sum
}
