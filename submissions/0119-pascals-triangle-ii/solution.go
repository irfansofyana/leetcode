func getRow(rowIndex int) []int {
    pascal := make([][]int, 0)
    
    ans := make([]int, 0)
    for i := 0; i <= rowIndex; i++ {
        tmp := make([]int, 0)
        for j := 0; j <= i; j++ {
            if j == 0 || j == i {
                tmp = append(tmp, 1)
            } else {
                tmp = append(tmp, pascal[i-1][j] + pascal[i-1][j-1])
            }
        }
        pascal = append(pascal, tmp)
        if i == rowIndex {
           ans = tmp 
        }
    }
    
    return ans
}
