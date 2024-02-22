func arraySign(nums []int) int {
    cntNeg := 0; cntZero := 0;
    
    for _, num := range nums {
        if num < 0 {
            cntNeg += 1
        } else if num == 0 {
            cntZero += 1
        }
    }
    
    if cntZero > 0 {
        return 0
    }
    if cntNeg % 2 == 1 {
        return -1
    }
    return 1
}