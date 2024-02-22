func constructRectangle(area int) []int {
    ans := make([]int, 2)
    diff := 1000000000
    
    for L := 1; L * L <= area; L++ {
        if area % L != 0 {
            continue
        }

        W := area / L
        if abs(L, W) < diff {
            diff = abs(L, W)
            ans = []int{L, W}
        }
    }
    
    if ans[0] < ans[1] {
        ans[0], ans[1] = ans[1], ans[0]
    }
    
    return ans
}

func abs(a, b int) int {
    if a > b {
        return a - b
    }

    return b-a
}