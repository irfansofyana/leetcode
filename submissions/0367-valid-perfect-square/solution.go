func isPerfectSquare(num int) bool {
    l := 1; r := num
    var ans bool = false
    for l <= r {
        mid := (l + r) / 2
        if mid * mid == num {
            ans = true
            break
        } else if mid * mid < num {
            l = mid + 1
        } else if mid * mid > num {
            r = mid - 1
        }
    }
    return ans
}
