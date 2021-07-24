func uniquePaths(m int, n int) int {
    ways := make([][]int, 0)
    
    for i := 0; i <= n + m -2; i++ {
        tmp := make([]int, 0)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                tmp = append(tmp, 1)
            } else {
                tmp = append(tmp, ways[i-1][j] + ways[i-1][j-1])
            }
        }
        ways = append(ways, tmp)
    }
    
    return ways[n+m-2][n-1]
}
