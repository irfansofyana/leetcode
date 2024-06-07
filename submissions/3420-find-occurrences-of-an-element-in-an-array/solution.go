func occurrencesOfElement(nums []int, queries []int, x int) []int {
    idx := make([]int, 0)

    for i, num := range nums {
        if num == x {
            idx = append(idx, i)
        }
    }

    ans := make([]int, 0)
    for _, q := range queries {
        q--

        var pos int
        if q >= len(idx) {
            pos = -1
        } else {
            pos = idx[q]
        }

        ans = append(ans, pos)
    }

    return ans
}
