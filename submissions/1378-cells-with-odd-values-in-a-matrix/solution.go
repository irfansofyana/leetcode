func incrementMatrix(m int, n int, mat *[][]int, r int, c int) {
    for i := 0; i < n; i++ {
        (*mat)[r][i] += 1
    }
    for i := 0; i < m; i++ {
        (*mat)[i][c] += 1
    }
}

func countOddCell(m int, n int, mat [][]int) int {
    cnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] % 2 == 1 {
                cnt++
            }
        }
    }
    return cnt
}

func oddCells(m int, n int, indices [][]int) int {
    mat := make([][]int, m)
    for i := 0; i < m; i++ {
        t := make([]int, n)
        mat[i] = t
    }
    for _, ind := range indices {
        r := ind[0]
        c := ind[1]
        incrementMatrix(m, n, &mat, r, c)
    }
    return countOddCell(m, n, mat)
}
