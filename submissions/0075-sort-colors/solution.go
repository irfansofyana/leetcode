func sortColors(nums []int)  {
    cnt := make([]int, 3)

    for _, num := range nums {
        cnt[num]++
    }

    idx := 0
    for i := 0; i < 3; i++ {
        for j := 0; j < cnt[i]; j++ {
            nums[idx] = i
            idx++
        }
    }
}
