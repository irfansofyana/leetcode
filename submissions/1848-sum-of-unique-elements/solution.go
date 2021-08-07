func sumOfUnique(nums []int) int {
    freq := map[int]int{}
    for _, num := range nums {
        freq[num] += 1
    }
    ans := 0
    for k, v := range freq {
        if v == 1 {
            ans += k
        }
    }
    return ans
}
