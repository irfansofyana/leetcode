func arrangeCoins(n int) int {
    lo := 1; hi := n
    ans := 1
    
    for lo <= hi {
        mid := (lo + hi) / 2
        if mid == 2 * n / (mid + 1) {
            ans = mid
            break
        } else if mid > 2 * n / (mid + 1) {
            hi = mid - 1
        } else {
            ans = mid
            lo = mid + 1
        }
    }
    return ans
}