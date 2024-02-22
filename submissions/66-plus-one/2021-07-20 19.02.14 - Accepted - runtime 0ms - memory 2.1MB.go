func plusOne(digits []int) []int {
    var n int = len(digits)
    var carry int = 0
    var ans = make([]int, n + 1)
    
    for i := n; i >= 1; i-- {
        if i == n {
            ans[i] = digits[i-1] + carry + 1
        } else {
            ans[i] = digits[i-1] + carry
        }
        carry = ans[i] / 10
        ans[i] %= 10
    }
    if carry > 0 {
        ans[0] = carry
        return ans[0:n+1]
    }
    return ans[1:n+1]
}