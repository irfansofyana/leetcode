func consecutiveNumbersSum(n int) int {
    n *= 2

    var ans int
    var k int64 = int64(1)
    for k * k <= int64(n) {
        if int64(n) % k == 0 {
            other := int64(n) / k
            if (other + 1 - k) % 2 == 0 {
                x := (other + 1 - k) / 2
                if x > 0 {
                    ans++
                }
            }
        }
        
        k++
    }

    return ans
}

// x, x+1, ..., x+k-1
// k*x + (k-1)*k/2
// 2kx + k(k-1) = 2n
// k (2x + k - 1) = 2n

// 2x + k - 1 = 2n/k
