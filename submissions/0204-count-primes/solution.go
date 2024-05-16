func countPrimes(n int) int {
    // true -> not prime, false prime
    sieves := make([]bool, n+2)
    
    cnt := 0
    for i := 2; i < n; i++ {
        if !sieves[i] {
            for j := int64(i) * int64(i); j <= int64(n); j+= int64(i) {
                sieves[j] = true
            }
        }

        if (!sieves[i]) {
            cnt++
        }
    }    

    return cnt
}

