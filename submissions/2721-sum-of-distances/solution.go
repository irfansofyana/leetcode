func distance(nums []int) []int64 {
    ansLeft := iterate(0, len(nums), 1, nums)
    ansRight := iterate(len(nums)-1, -1, -1, nums)
    for i := 0; i < len(ansLeft); i++ {
        ansLeft[i] += ansRight[i]
    }

    return ansLeft
}

func abs(a, b int64) int64 {
    if a < b {
        return b-a
    }

    return a-b
}

func iterate(start, end, inc int, nums []int) []int64{
    cnt := make(map[int]int)
    sum := make(map[int]int64)
    
    ans := make([]int64, len(nums))
    for i := start; i != end; i += inc {
        num := nums[i]
        ans[i] = abs(int64(i) * int64(cnt[num]), sum[num])
        cnt[num]++
        sum[num] += int64(i)
    }

    return ans
}
