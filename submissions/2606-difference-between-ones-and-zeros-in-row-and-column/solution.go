func onesMinusZeros(grid [][]int) [][]int {
    m := len(grid); n := len(grid[0])
    onesRow := make([]int, m); onesCol := make([]int, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                onesRow[i]++
                onesCol[j]++
            }
        }
    }

    ans := make([][]int, m)
    for i := 0; i < m; i++ {
        ans[i] = make([]int, n)
        for j := 0; j < n; j++ {
            ans[i][j] = onesRow[i] + onesCol[j] - (m - onesRow[i]) - (n - onesCol[j])
        }
    }

    return ans
}
