func equalPairs(grid [][]int) int {
    n := len(grid)

    cnt := 0
    for row := 0; row < n; row++{
        for col := 0; col < n; col++ {
            flag := true
            for idx := 0; idx < n; idx++ {
                if grid[row][idx] != grid[idx][col] {
                    flag = false
                    break
                }
            }

            if flag {
                cnt++
            }
        }
    }

    return cnt
}
