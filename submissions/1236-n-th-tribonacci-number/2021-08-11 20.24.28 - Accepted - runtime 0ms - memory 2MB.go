func tribonacci(n int) int {
    tribo := make([]int, n+1)
    if n == 0 {
        return 0
    } 
    if n == 1 || n == 2 {
        return 1
    }
    tribo[0] = 0; tribo[1] = 1; tribo[2] = 1;
    
    for i := 3; i <= n; i++ {
        tribo[i] = tribo[i-1] + tribo[i-2] + tribo[i-3]
    }
    
    return tribo[n]
}