func sumZero(n int) []int {
    ans := make([]int, 0)
    cnt := 1
    
    for i := 0; i < n/2; i++ {
        ans = append(ans, cnt)
        ans = append(ans, -1 * cnt)
        cnt += 1
    }
    
    if n % 2 == 1 {
        ans = append(ans, 0)
    }
    
    return ans
}
