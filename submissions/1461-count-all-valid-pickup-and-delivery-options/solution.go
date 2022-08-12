func countOrders(n int) int {
    mod := 1000000007
    
    ans := 1
    for i :=  2*n; i >= 2; i -= 2 {
        ans = (ans * nc2(i)) % mod
    }
    
    return ans
}

func nc2(n int) int {
    return n*(n - 1) / 2
}
