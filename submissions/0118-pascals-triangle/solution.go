func generate(numRows int) [][]int {
    var ans = make([][]int, 0)
    
    for i := 0; i < numRows; i++ {
        var row = make([]int, 0)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                row = append(row, 1)
            } else {
                row = append(row, ans[i-1][j-1] + ans[i-1][j])
            }
        }
        ans = append(ans, row)
    }
    
    return ans
}
