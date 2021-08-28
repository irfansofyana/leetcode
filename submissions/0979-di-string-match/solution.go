func diStringMatch(s string) []int {
    n := len(s)
    used := make([]bool, n+1)
    
    currI := 0
    currD := n
    
    ans := make([]int, 0)
    for i := 0; i < n; i++ {
        if s[i] == 'I' {
            for currI <= n + 1 && used[currI] {
                currI++
            }
            ans = append(ans, currI)
            used[currI] = true
            currI++
        } else {
             for currD >= 0 && used[currD] {
                currD--
            }
            ans = append(ans, currD)
            used[currD] = true
            currD--
        }
    }
    ans = append(ans, currD)
    
    return ans
}
