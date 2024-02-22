func matrixReshape(mat [][]int, r int, c int) [][]int {
    n := len(mat)
    m := len(mat[0])
    if n * m != r * c {
        return mat
    }
    
    ans := make([][]int, r)
    iRow := 0; iCol := 0
    
    for i := 0; i < r; i++ {
        for j := 0; j < c; j++ {
            if iCol == m {
                iCol = 0
                iRow++
            }
            ans[i] = append(ans[i], mat[iRow][iCol])
            iCol++
        }
    }
    
    return ans
}