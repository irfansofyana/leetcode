func numberOfArithmeticSlices(nums []int) int {
    var n int = len(nums)
    
    if n < 3 {
        return 0
    }
    
    var ans int = 0
    var idxCurr int = 0
    var occur int = 0
    
    for idxCurr < n {
        if idxCurr + 1 >= n {
            break
        }
        var diff int = nums[idxCurr + 1] - nums[idxCurr]
        var idxNext int = idxCurr + 2
        if idxNext >= n {
            break
        }
        for idxNext < n && nums[idxNext] - nums[idxNext - 1] == diff {
            idxNext += 1
        }
        occur = idxNext - idxCurr
        ans += (occur - 2) * (occur - 1) / 2
        idxCurr = idxNext - 1
    }
    
    return ans
}
