func trailingZeroes(n int) int {
    var zeros int = 0
    for n > 0 {
        zeros += n / 5
        n /= 5
    }
    return zeros
}