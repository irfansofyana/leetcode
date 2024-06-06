func countKDifference(nums []int, k int) int {
    freq := make(map[int]int)

    for _, num := range nums {
        freq[num]++
    }

    ans := 0
    for i := 1; i <= 100; i++ {
        ans += freq[i] * freq[i+k]
    }

    return ans
}
