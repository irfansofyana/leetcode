func generateParenthesis(n int) []string {
    ans := make([]string, 0)
    
    n *= 2
    for i := 0; i < (1 << n); i++ {
        sum := 0
        st := ""
        isValid := true
        for j := 0; j < n && isValid; j++ {
            if (i&(1<<j)) > 0 {
                sum += 1
                st = st + "("
            } else {
                sum -= 1
                st = st + ")"
            }
            if sum < 0 {
                isValid = false
            }
        }
        if sum != 0 {
            isValid = false
        }
        if isValid {
            ans = append(ans, st)
        }
    }
    
    return ans
}