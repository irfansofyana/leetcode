func maximumWealth(accounts [][]int) int {
    maks := 0; m := len(accounts); n := len(accounts[0])
    for i := 0; i < m; i++ {
        sum := 0
        for j := 0; j < n; j++ {
            sum += accounts[i][j]
        }
        if sum > maks {
            maks = sum
        }
    }
    return maks
}