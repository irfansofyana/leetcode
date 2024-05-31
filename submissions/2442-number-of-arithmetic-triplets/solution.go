func arithmeticTriplets(nums []int, diff int) int {
    exist := make(map[int]bool)

    for _, num := range nums {
        exist[num] = true
    }

    cnt := 0
    for _, num := range nums {
        if exist[num+diff] && exist[num+2*diff] {
            cnt++
        }
    }

    return cnt
}
