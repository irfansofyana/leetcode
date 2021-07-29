func mySqrt(x int) int {
    lo := 1; hi := x
    ans := 0
    
    if x == 0 {
        return 0
    }
    
    for lo <= hi {
        mid := (lo + hi) / 2
        if mid == x / mid {
            ans = mid
            break
        } else if mid > x / mid {
            hi = mid - 1
        } else {
            ans = mid
            lo = mid + 1
        }
    }
    return ans
}
