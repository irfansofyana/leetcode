func maxOperations(nums []int, k int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    tot := 0
    for _, num := range nums {
        if k - num < 0 {
            continue
        }

        if k == 2 * num {
            tot += freq[num] / 2
            freq[num] = 0
        } else {
            curr := freq[num]
            mini := curr
            if freq[k-num] < mini {
                mini = freq[k-num]
            }
            tot += mini

            freq[num] -= mini
            freq[k-num] -= mini
        }
    }

    return tot
}