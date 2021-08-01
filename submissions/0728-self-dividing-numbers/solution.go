func isSelfDividing(num int) bool {
    tmp := num
    isFulfill := true
    
    for tmp > 0 && isFulfill {
        digit := tmp % 10
        if digit == 0 || num % digit != 0 {
            isFulfill = false
        }
        tmp /= 10
    }
    
    return isFulfill
}

func selfDividingNumbers(left int, right int) []int {
    result := make([]int, 0)
    
    for i := left; i <= right; i++ {
        if isSelfDividing(i) {
            result = append(result, i)
        }    
    }
    
    return result
}
