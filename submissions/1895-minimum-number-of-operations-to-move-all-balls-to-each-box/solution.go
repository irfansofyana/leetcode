func minOperations(boxes string) []int {
    n := len(boxes); cnt := make([]int, n+1); sum := make([]int, n+1)
    ans := make([]int, 0)
    
    for i := 1; i <= n; i++ {
        if boxes[i-1] == '1' {
            cnt[i] = cnt[i-1] + 1
            sum[i] = sum[i-1] + i
        } else {
            cnt[i] = cnt[i-1]
            sum[i] = sum[i-1]
        }     
    }
    
    for i := 1; i <= n; i++ {
        cntLeft := cnt[i]; sumLeft := sum[i]
        cntRight := cnt[n] - cnt[i]; sumRight := sum[n] - sum[i]
        numOperations := cntLeft * i - sumLeft + sumRight - cntRight * i
        ans = append(ans, numOperations)
    }
    
    return ans
}
