func findShortestSubArray(nums []int) int {
    freq := map[int]int{}; degree := 0
    for _, num := range nums {
        freq[num] += 1
        if freq[num] > degree {
            degree = freq[num]
        }
    }
    
    firstIdx := map[int]int{}; freq = map[int]int{}; ans := 1000000000
    for i, num := range nums {
        _, ok := firstIdx[num]
        if !ok {
            firstIdx[num] = i
        }
        freq[num] += 1
        if freq[num] == degree {
            if (i - firstIdx[num] + 1 < ans) {
                ans = i - firstIdx[num] + 1
            }
        }
    }
    
    return ans
}