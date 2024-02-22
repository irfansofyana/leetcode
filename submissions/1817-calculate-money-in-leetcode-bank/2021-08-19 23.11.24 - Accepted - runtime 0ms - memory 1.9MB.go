func totalMoney(n int) int {
    ans := 0
    for i := 1; i <= n; i++ {
        if i % 7 == 0 {
            ans += 6 + (i / 7)
        } else {
            ans += (i/7) + (i%7)   
        }
    }
    return ans
}