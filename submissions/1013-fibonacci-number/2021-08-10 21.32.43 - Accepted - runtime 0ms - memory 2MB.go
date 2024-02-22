func fib(n int) int {
    fibo := make([]int, n+2)
    
    fibo[0] = 0; fibo[1] = 1
    for i := 2; i <= n; i++ {
        fibo[i] = fibo[i-1] + fibo[i-2]    
    }
    
    return fibo[n]
}