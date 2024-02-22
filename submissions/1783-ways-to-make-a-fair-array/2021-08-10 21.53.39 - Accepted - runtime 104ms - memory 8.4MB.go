func waysToMakeFair(nums []int) int {
    n := len(nums)
    evenSum := make([]int, n+2); oddSum := make([]int, n+2)
    
    for i := 0; i < n; i++ {
        if i == 0 {
            evenSum[i] = nums[i]
            oddSum[i] = 0
        } else if i % 2 == 0 {
            evenSum[i] = evenSum[i-1] + nums[i]
            oddSum[i] = oddSum[i-1]
        } else {
            evenSum[i] = evenSum[i-1]
            oddSum[i] = oddSum[i-1] + nums[i]
        }
    }
    
    count := 0
    for i := 0; i < n; i++ {
        sumEvenIdx := 0; sumOddIdx := 0
        
        if i > 0 {
            sumEvenIdx += evenSum[i-1]
            sumOddIdx += oddSum[i-1]
        }
        sumEvenIdx += oddSum[n-1]-oddSum[i]
        sumOddIdx += evenSum[n-1]-evenSum[i]
        
        if sumEvenIdx == sumOddIdx {
            count += 1
        }
    }
    
    return count
}