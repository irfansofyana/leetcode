func countBits(n int) []int {
    var ans = make([]int, 0)
    for i := 0; i <= n; i++ {
        tmp := i
        cnt := 0
        for tmp > 0 {
            if tmp%2 == 1 {
                cnt += 1
            }
            tmp /= 2
        }
        ans = append(ans, cnt)
    }
    return ans
}
