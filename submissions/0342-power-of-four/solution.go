func isPowerOfFour(n int) bool {
    if n <= 0 {
        return false
    }
    
    cnt := 0
    for n % 2 == 0{
        cnt += 1
        n /= 2
    }
    return cnt%2 == 0 && n == 1
}
