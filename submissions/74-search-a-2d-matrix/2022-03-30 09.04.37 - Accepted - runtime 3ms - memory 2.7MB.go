func getRowCol(idx, numRow, numCol int) (int, int) {
    return idx/numCol, idx%numCol
}

func searchMatrix(matrix [][]int, target int) bool {
    n, m := len(matrix), len(matrix[0])
    l, r := 0, n * m - 1
    
    for l <= r {
        mid := (l + r) / 2
        row, col := getRowCol(mid, n, m)
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] > target {
            r = mid - 1
        } else if matrix[row][col] < target {
            l = mid + 1
        }
    }
    
    return false
}