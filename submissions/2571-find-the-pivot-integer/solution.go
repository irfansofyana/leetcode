func pivotInteger(n int) int {
    ans := -1
    sum := 0
    for i := 1; i <= n && ans == -1; i++ {
        sum += i
        if sum == n*(n+1)/2 - sum + i {
            ans = i
        }
    }

    return ans
}
