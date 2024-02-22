func decode(encoded []int, first int) []int {
    ans := make([]int, 0); ans = append(ans, first); n := len(encoded)
    for i := 0; i < n; i++ {
        ans = append(ans, encoded[i]^ans[i])
    }
    return ans
}