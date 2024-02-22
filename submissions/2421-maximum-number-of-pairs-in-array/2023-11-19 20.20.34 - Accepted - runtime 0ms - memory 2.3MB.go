func numberOfPairs(nums []int) []int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    ans := make([]int, 2)
    for _, v := range freq {
        ans[0] += v / 2
        ans[1] += v % 2
    }

    return ans
}