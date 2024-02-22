func calculate(n int, isProduct bool) int {
    var res int = 0
    if isProduct {
        res = 1
    }
    for n > 0 {
        digit := n % 10
        if isProduct {
            res *= digit
        } else {
            res += digit
        }
        n /= 10
    }
    
    return res
}


func subtractProductAndSum(n int) int {
    return calculate(n, true) - calculate(n, false)
}