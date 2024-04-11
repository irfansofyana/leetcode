func commonFactors(a int, b int) int {
    cnt := 0
    for i := 1; i <= min(a, b); i++ {
        if a % i == 0 && b % i == 0 {
            cnt++
        }
    }

    return cnt
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
