func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    sum := make([]int, 2) //0 -> alice, 1 -> bob
    arr := [2][]int{aliceSizes, bobSizes}
    freq := make([]map[int]int, 2)
    
    for i := 0; i < 2; i++ {
        sum[i] = 0
        freq[i] = map[int]int{}
        for j := 0; j < len(arr[i]); j++ {
            freq[i][arr[i][j]]++
            sum[i] += arr[i][j]
        }
    }
    
    diff := sum[0] - sum[1]
    ans := make([]int, 2)
    for i := 0; i < len(arr[0]); i++ {
        if diff % 2 == 1 {
            continue
        }
        x := arr[0][i]
        y := x - diff / 2
        if freq[1][y] > 0 {
            ans[0] = x; ans[1] = y
            break
        }
    }
    
    return ans
}