func sqrt(n int) int {
    lo := 0; hi := n
    
    ans := -1
    for lo <= hi {
        mid := (lo + hi) / 2
        if mid * mid == n {
            ans = mid
            break
        } else if mid * mid > n {
            hi = mid - 1
        } else {
            ans = mid
            lo = mid + 1
        }
    }
    
    return ans
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n == 2 {
        return true
    }
    for i := 2; i <= sqrt(n); i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func isThree(n int) bool {
    sqrtN := sqrt(n)
    if sqrtN * sqrtN == n && isPrime(sqrtN) {
        return true
    }
    return false
}