func pivotIndex(nums []int) int {
    n := len(nums)
    pSum := make([]int, n)
    
    for i := 0; i < n; i++ {
        if i == 0 {
            pSum[i] = nums[i]
            continue
        }
        
        pSum[i] = pSum[i-1] + nums[i]
    }
    
    for i := 0; i < n; i++ {
        sumLeft := 0
        if i != 0 {
            sumLeft = pSum[i-1]
        }
        sumRight := pSum[n-1]-pSum[i]
        
        if sumLeft == sumRight {
            return i
        }
    }
    
    return -1
}