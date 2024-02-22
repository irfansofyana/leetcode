func combine(n int, k int) [][]int {
    ans := make([][]int, 0)
    for i := 0; i < (1 << n); i++ {
        arr := make([]int, 0)
        for j := 0; j < n; j++ {
            if i&(1<<j) > 0 {
                arr = append(arr, j+1)
            }
        }
        if len(arr) == k {
            ans = append(ans, arr)
        }
    }
    return ans
}