func findMinFibonacciNumbers(k int) int {
    fibo := make([]int64, 55)
    fibo[0] = int64(1); fibo[1] = int64(1)
    
    start := 1
    for i := 2; i < 55; i++ {
        fibo[i] = fibo[i-1] + fibo[i-2]
        if fibo[i] > int64(k) {
            start = i-1
            break
        }
    }

    ans := 0
    for k > 0 && start >= 0 {
        many := k / int(fibo[start])
        ans += many
        k -= many * int(fibo[start])
        start--
    }

    return ans
}
