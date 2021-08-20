func isPrime(n int) bool {
    if n < 2 {
        return false
    }    
    if n == 2 {
        return true
    }
    for i := 2; i <= n-1; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func numPrimeArrangements(n int) int {
    cntPrime := 0
    fac := make([]int, 0)
    MOD := 1000000007
    
    for i := 0; i <= n; i++ {
        if i == 0 {
            fac = append(fac, 1)
        } else {
            fac = append(fac, (fac[i-1] * i) % MOD)
        }
    }
    for i := 1; i <= n; i++ {
        if isPrime(i) {
            cntPrime += 1
        }
    }
    return (fac[cntPrime] * fac[n-cntPrime]) % MOD
}
