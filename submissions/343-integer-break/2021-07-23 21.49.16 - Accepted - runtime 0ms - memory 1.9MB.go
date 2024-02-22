func integerBreak(n int) int {
    maks := 0
    for i := 2; i <= n; i++ {
        var mul int = 1
        var remainder int = n % i
        
        for j := 0; j < i; j++ {
            var multiplier int = n/i
            if remainder > 0 {
                multiplier += 1
                remainder -= 1
            }
            mul *= multiplier
        }
        
        if mul > maks {
            maks = mul
        }
    }
    return maks
}