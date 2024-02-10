func findMatrix(nums []int) [][]int {
    mep := make([]int, len(nums)+1)
    for _, num := range nums {
        mep[num]++
    }

    ans := make([][]int, 0)
    for i := 1; i <= len(nums); i++ {
        freq := mep[i]
        ansLength := len(ans)

        for j := 0; j < min(freq, ansLength); j++ {
            ans[j] = append(ans[j], i)
        }
        
        for j := 0; j < freq - ansLength; j++ {
            ans = append(ans, []int{i})
        }
    }

    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}
