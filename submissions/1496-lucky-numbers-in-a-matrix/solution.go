func minRow(mat [][]int, row int) int {
    m := len(mat[0])
    mini := 1000000000
    for i := 0; i < m; i++ {
        if mat[row][i] < mini {
            mini = mat[row][i]
        }
    }
    return mini
}

func maxCol(mat [][]int, col int) int {
    n := len(mat)
    maks := -1
    for i := 0; i < n; i++ {
        if mat[i][col] > maks {
            maks = mat[i][col]
        }
    }
    return maks
}


func luckyNumbers (matrix [][]int) []int {
    n := len(matrix)
    m := len(matrix[0])
    
    ans := make([]int, 0)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            minr := minRow(matrix, i)
            maxc := maxCol(matrix, j)
            if minr == matrix[i][j] && maxc == matrix[i][j] {
                ans = append(ans, matrix[i][j])
            }
        }
    }
    
    return ans
}
