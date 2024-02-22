func minNum(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func maxNum(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minimumDeletions(nums []int) int {
    max := -1000000; min := 1000000
    maxIdx := -1; minIdx := -1
    n := len(nums)
    
    for i, num := range nums {
        if num > max {
            max = num
            maxIdx = i
        }
        if num < min {
            min = num
            minIdx = i
        }
    }
    
    if minIdx == maxIdx {
        return minNum(minIdx+1, n-minIdx)
    }
    
    ans := maxNum(minIdx, maxIdx) + 1
    ans = minNum(ans, n-minNum(minIdx, maxIdx))
    
    if minIdx < maxIdx {
        ans = minNum(ans, n-maxIdx+minIdx+1)
    }
    if maxIdx < minIdx {
        ans = minNum(ans, n-minIdx+maxIdx+1)
    }
    
    return ans
}