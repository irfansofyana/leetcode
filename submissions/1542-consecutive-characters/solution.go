func maxPower(s string) int {
    idx := 0; n := len(s)
    ans := 0
    
    for idx < n {
        tidx := idx + 1
        curr := 1
        for tidx < n && s[tidx] == s[idx] {
            tidx++
            curr++
        }
        if curr > ans {
            ans = curr
        }
        idx = tidx
    }
    
    return ans
}
