func isValid(s string) bool {
    if s[0] == '0' {
        return false
    }
    if s[0] == '1' {
        return true
    }
    if s[0] == '2' && (s[1] >= '0' && s[1] <= '6') {
        return true
    }
    return false
}

func numDecodings(s string) int {
    n := len(s)
    dp := make([]int, n+1)
    
    dp[n] = 1
    for i := n-1; i >= 0; i-- {
        if s[i] == '0' {
            dp[i] = 0
        } else if i == n-1 {
            dp[i] = 1
        } else {
            var tmp string = s[i:i+2]
            cnt := 0
            if isValid(tmp) {
                cnt += dp[i+2] 
            }
            cnt += dp[i+1]
            dp[i] = cnt
        }
    }
    
    return dp[0]
}